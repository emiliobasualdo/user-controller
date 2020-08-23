package persistence

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"massimple.com/wallet-controller/internal/models"
	"massimple.com/wallet-controller/internal/utils"
	"os"
	"strings"
	"testing"
	"time"
)

func setup() {
	// We pick up the env setup
	utils.EnvInit("TEST")
	// we now connect
	dbInit()
}

var dbName string
func dbInit() {
	conUri := viper.GetString("database.connectionUri")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conUri))
	if err != nil {
		panic(err)
	}
	dbName = viper.GetString("database.dbName")
	db = client.Database(dbName)
	// We create the collections
	for _, collection := range collections {
		if err := db.CreateCollection(ctx, collection.Name); err != nil {
			if !strings.Contains(err.Error(), " already exists") {
				panic(err)
			}
		}
		*(collection.reference) = db.Collection(collection.Name)
	}
}

func shutdown() {
	_=db.Drop(ctx)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func cleanUp(){
	// todo efficiency
	// we clean the bd and create a new one
	_=db.Drop(ctx)
}

func insertAccount(acc models.Account) primitive.ObjectID{
	result, err := usersCollection.InsertOne(ctx, acc)
	if err != nil {
		panic(err)
	}
	v, _ := result.InsertedID.(primitive.ObjectID)
	return v
}

func TestGetAccountByPhoneNumberOrCreate_returnsExistent(t *testing.T) {
	// SETUP
	numberExpected := models.PhoneNumber("005491133071114")
	theIdExpected := models.ID(numberExpected.String())
	accountExpected := models.Account{
		ID:          theIdExpected,
		Name:        "To be returned",
		PhoneNumber: numberExpected,
	}
	insertAccount(accountExpected)
	insertAccount(models.Account{
		ID:          "Other ID",
		Name:        "No to be returned",
		PhoneNumber: "005491133333333",
	})
	//EXERCISE
	acc, err := GetAccountByPhoneNumberOrCreate(numberExpected, "")
	// ASSERT
	t.Run("Returns Existent", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
		if acc.PhoneNumber != numberExpected {
			t.Errorf("got %s, want %s", acc.PhoneNumber, numberExpected)
		}
		if acc.ID != theIdExpected {
			t.Errorf("got %s, want %s", acc.ID, theIdExpected)
		}
	})
	t.Cleanup(cleanUp)
}

func TestGetAccountByPhoneNumberOrCreate_creates(t *testing.T) {
	// SETUP
	numberExpected := models.PhoneNumber("005491133071114")
	idExpected := models.ID(numberExpected.String())
	//EXERCISE
	acc , err := GetAccountByPhoneNumberOrCreate(numberExpected, idExpected)
	// ASSERT
	t.Run("Creates a new acc", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
		if acc.PhoneNumber != numberExpected {
			t.Errorf("got %s, want %s", acc.PhoneNumber, numberExpected)
		}
		if acc.MongoID.IsZero() {
			t.Errorf("got %t, want %t", acc.MongoID.IsZero(), false)
		}
		if acc.ID != idExpected {
			t.Errorf("got %s, want %s", acc.ID, idExpected)
		}
	})
	t.Cleanup(cleanUp)
}

func TestGetAccountById_returnsExistent(t *testing.T) {
	// SETUP
	expectedId := models.ID("asldfkjalsd")
	accountExpected := models.Account{ID: expectedId, Name: "To be returned"}
	mongoId := insertAccount(accountExpected).Hex()
	//EXERCISE
	acc, err := GetAccountById(expectedId)
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
	})
	t.Run("Got what expected", func(t *testing.T) {
		if acc.ID != expectedId {
			t.Errorf("got %s, want %s", acc.ID, expectedId)
		}
	})
	t.Run("Got what expected", func(t *testing.T) {
		if acc.MongoID.Hex() != mongoId {
			t.Errorf("got %s, want %s", acc.MongoID.Hex(), mongoId)
		}
	})
	t.Cleanup(cleanUp)
}

func TestGetAccountById_returnsError(t *testing.T) {
	// SETUP
	//EXERCISE
	_, err := GetAccountById("abcd")
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if err == nil {
			t.Errorf("got nil, want error")
		}
	})
	t.Cleanup(cleanUp)
}

func TestGetAccountById_returnsNoSuchAccountError(t *testing.T) {
	// SETUP
	//EXERCISE
	_, err := GetAccountById("abcd523b4c06cced03ec6f10")
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if _, ok := err.(*models.NoSuchAccountError); !ok {
			t.Errorf("got %s", err)
		}
	})
	t.Cleanup(cleanUp)
}

func TestReplaceAccount_returnsError(t *testing.T) {
	// SETUP
	//EXERCISE
	err := ReplaceAccount(models.Account{ID:"abcd"})
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if err == nil {
			t.Errorf("got nil, want error")
		}
	})
	t.Cleanup(cleanUp)
}

func TestReplaceAccount_returnsNoSuchAccountError(t *testing.T) {
	// SETUP
	//EXERCISE
	err := ReplaceAccount(models.Account{ID:"abcd523b4c06cced03ec6f10"})
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if _, ok := err.(*models.NoSuchAccountError); !ok {
			t.Errorf("got %s", err)
		}
	})
	t.Cleanup(cleanUp)
}

func TestReplaceAccount_replaces(t *testing.T) {
	// SETUP
	exptectedId := models.ID("lsdjfalsdk")
	oldAcc := models.Account{
		ID:			exptectedId,
		Name:        "Name",
		LastName:    "Lastname",
		PhoneNumber: "005491133071114",
	}
	mongoId := insertAccount(oldAcc)
	now := time.Now()
	replaceWith := models.Account{
		ID: 		exptectedId,
		LastName:  "Other lastname",
		CreatedAt: now,
	}
	//EXERCISE
	err := ReplaceAccount(replaceWith)
	newAcc := &models.Account{}
	_=usersCollection.FindOne(ctx, bson.M{"id": exptectedId}).Decode(&newAcc)
	// ASSERT
	t.Run("Replaces", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
		if newAcc.MongoID != mongoId {
			t.Errorf("got %s, want %s", newAcc.MongoID, mongoId.Hex())
		}
		if newAcc.Name != "" {
			t.Errorf("got %s, want %s", newAcc.Name, "")
		}
		if newAcc.LastName != "Other lastname" {
			t.Errorf("got %s, want %s", newAcc.Name, "Other lastname")
		}
		if newAcc.ID != exptectedId {
			t.Errorf("got %s, want %s", newAcc.ID, exptectedId)
		}
	})
	t.Cleanup(cleanUp)
}
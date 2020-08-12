package persistence

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"massimple.com/wallet-controller/internal/models"
	"os"
	"strings"
	"testing"
	"time"
)

func env(){
	// set defaults
	viper.SetConfigName("config-test")
	viper.AddConfigPath("./../../")		// search locally in this directory
	err := viper.ReadInConfig() 		// Find and read the config file
	if err != nil {
		panic(err)
	}
}

func setup(m *testing.M) {
	// We pick up the env setup
	env()
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
	setup(m)
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
	toReturn := "005491133071114"
	accountExpected := models.Account{
		Name:        "To be returned",
		PhoneNumber: toReturn,
	}
	insertAccount(accountExpected)
	insertAccount(models.Account{
		Name:        "No to be returned",
		PhoneNumber: "005491133333333",
	})
	//EXERCISE
	acc, err := GetAccountByPhoneNumberOrCreate(accountExpected)
	// ASSERT
	t.Run("Returns Existent", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
		if acc.PhoneNumber != toReturn {
			t.Errorf("got %s, want %s", acc.PhoneNumber, toReturn)
		}
	})
	t.Cleanup(cleanUp)
}

func TestGetAccountByPhoneNumberOrCreate_creates(t *testing.T) {
	// SETUP
	toReturn := "005491133071114"
	toCreate := models.Account{
		Name:        "To be returned",
		PhoneNumber: toReturn,
	}
	//EXERCISE
	acc , err := GetAccountByPhoneNumberOrCreate(toCreate)
	// ASSERT
	t.Run("Creates a new acc", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
		if acc.PhoneNumber != toReturn {
			t.Errorf("got %s, want %s", acc.PhoneNumber, toReturn)
		}
		if acc.MongoID.IsZero() {
			t.Errorf("got %t, want %t", acc.MongoID.IsZero(), false)
		}
	})
	t.Cleanup(cleanUp)
}

func TestGetAccountById_returnsExistent(t *testing.T) {
	// SETUP
	accountExpected := models.Account{Name: "To be returned"}
	mongoId := insertAccount(accountExpected).Hex()
	//EXERCISE
	acc, err := GetAccountById(mongoId)
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
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

func TestGetAccountByPhoneNumberOrCreate_replaces(t *testing.T) {
	// SETUP
	oldAcc := models.Account{
		Name:        "Name",
		LastName:    "Lastname",
		PhoneNumber: "005491133071114",
	}
	mongoId := insertAccount(oldAcc)
	now := time.Now()
	replaceWith := models.Account{
		ID: 		mongoId.Hex(),
		LastName:  "Other lastname",
		CreatedAt: now,
	}
	//EXERCISE
	err := ReplaceAccount(replaceWith)
	newAcc := &models.Account{}
	_=usersCollection.FindOne(ctx, bson.M{"_id": mongoId}).Decode(&newAcc)
	// ASSERT
	t.Run("Replaces", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want nil", err.Error())
		}
		if newAcc.ID != mongoId.Hex() {
			t.Errorf("got %s, want %s", newAcc.ID, mongoId.Hex())
		}
		if newAcc.Name != "" {
			t.Errorf("got %s, want %s", newAcc.Name, "")
		}
		if newAcc.LastName != "Other lastname" {
			t.Errorf("got %s, want %s", newAcc.Name, "Other lastname")
		}
	})
	t.Cleanup(cleanUp)
}
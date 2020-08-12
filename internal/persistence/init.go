package persistence

import (
	"context"
	"github.com/apsdehal/go-logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"
)

// db references
var db *mongo.Database
var client *mongo.Client
var log *logger.Logger

// collection references
var usersCollection *mongo.Collection
var transactionsCollection *mongo.Collection
var ctx = context.Background() // todo estudiar para qué sirve

const TRANSACTIONS = "transactions"
var collections = []struct{
	Name string
	reference **mongo.Collection
}{
	{"users", &usersCollection},
	{TRANSACTIONS, &transactionsCollection},
}

func Init() {
	var err error
	log, err = logger.New("Persistence", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	// we get the basicUri
	conUri := viper.GetString("database.connectionUri")
	// we now connect
	log.InfoF("Connecting to: %s", conUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(conUri))
	if err != nil {
		panic(err)
	}
	dbName := viper.GetString("database.dbName")
	db = client.Database(dbName)
	log.Info("DB connected successfully")
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

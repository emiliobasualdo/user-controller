package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "massimple.com/wallet-controller/internal/models"
)


func GetAccountByPhoneNumberOrCreate(query Account) (Account, error) {
	account, err := getByPhoneNumber(query.PhoneNumber)
	if err == nil {
		return account, nil
	}
	if err != mongo.ErrNoDocuments {
		return Account{}, err
	}
	account = query
	result , err := usersCollection.InsertOne(ctx, account)
	if err != nil {
		return account, err
	} else {
		id, _ := result.InsertedID.(primitive.ObjectID)
		account.MongoID = id
		return account, nil
	}
}

func GetAccountById(id ID) (Account, error) {
	var acc Account
	_id, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return acc, err
	}
	single := usersCollection.FindOne(ctx, bson.M{"_id": _id})
	if single.Err() == mongo.ErrNoDocuments {
		return Account{}, &NoSuchAccountError{Query: id.String()}
	}
	if err := single.Decode(&acc); err != nil{
		return Account{}, err
	}
	acc.MongoID = _id
	return acc, nil
}

func getByPhoneNumber(phoneNumber string) (Account, error) {
	var acc Account
	single := usersCollection.FindOne(ctx, Account{PhoneNumber: phoneNumber})
	if single.Err() == mongo.ErrNoDocuments {
		return Account{}, single.Err()
	}
	if err := single.Decode(&acc); err != nil{
		return Account{}, err
	}
	return acc, nil
}

func ReplaceAccount(new Account) error {
	_id, err := primitive.ObjectIDFromHex(new.ID.String())
	if err != nil {
		return err
	}
	single := usersCollection.FindOneAndReplace(ctx, bson.M{"_id": _id}, new)
	if single.Err() == mongo.ErrNoDocuments {
		return &NoSuchAccountError{Query: new.ID.String()}
	}
	return nil
}

/*
Trasnaction contine todos los datos de la transacción.
Se guarda en gp(sin tanto detalle), en el usuario y en una collection transaction
*/
func SaveTransaction(trans Transaction) error {
	// iniciamos la transacción
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	if err := session.StartTransaction(); err != nil {
		return err
	}
	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		// Le insertamos la transacción a la cuenta
		id, _ := primitive.ObjectIDFromHex(trans.OriginAccountId.String())
		query := bson.M{"_id": id}
		update := bson.M{"$push": bson.M{TRANSACTIONS: trans}}
		if _, err := usersCollection.UpdateOne(ctx, query, update); err != nil {
			if err:= session.AbortTransaction(ctx); err != nil {
				log.ErrorF("While aborting(q) transaction %s", err.Error())
			}
			return err
		}
		// Insertamos la transacción en la colección de transacciones
		_, err = transactionsCollection.InsertOne(ctx, trans)
		if err != nil {
			if err:= session.AbortTransaction(ctx); err != nil {
				log.ErrorF("While aborting(2) transaction %s", err.Error())
			}
			return err
		}
		// Commiteamos
		if err = session.CommitTransaction(sc); err != nil {
			log.ErrorF("While committing transaction %s", err.Error())
			return err
		}
		return nil
	})
	session.EndSession(ctx)
	return err
}

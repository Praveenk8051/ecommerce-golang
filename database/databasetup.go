package database


import(
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DBSet() *mongo.Client{

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err!=nil{
		log.fatal(err)
	}
	ctx, cancel := context.WithTimeOut(context.Background(), 10*time.Second)

	defer.cancel()

	err = client.Connect(ctx)
	if err!=nil{
		log.fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Println("Failed to connect to MongodB")
		nil
	}
	fmt.Println("Successfully connected to the MongoDB")
	return client
}


	var Client *mongo.Client = DBSet()


func UserData(client *mongo.Client, collectionname string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("Ecommerce").collection(collectionName)
	return collection
}


func ProductData(client *mongo.Client, collectionname string) *mongo.Collection{
	var productCollection


}

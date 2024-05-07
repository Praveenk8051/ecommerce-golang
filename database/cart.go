package database

import (
	"errors"
	"log"

	"github.com/praveen/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var (
	ErrCantFindProduct          = errors.New("Cant find the Product")
	ErrCantDecodeProducts       = errors.New("Cant find the Product")
	ErrUserIdIsNotValid         = errors.New("This User is not valid")
	ErrCantUpdateUser           = errors.New("Cannot add this product to the cart")
	ErrCantRemoteRemoveItemCart = errors.New("Cannot remove this item from the cart")
	ErrCantGetItem              = errors.New("Cant get the item")
	ErrCantBuyCartItem          = errors.New("Cant buy the item from the cart")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productCart)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProducts
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productCart}}}}}}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}

func RemoveCartItem(ctx context.Context, prodCollection, userCollection *mongoCollection, productID primitive.ObjectID, userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{"$pull": bson.M{"usercart": bson.M{"_id": productID}}}

	_, err = UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantRemoteRemoveItemCart
	}
	return nil
}

func BuyItemFromCart() {

}

func InstantBuyer() {

}

package database

import (
	"errors"
	"log"
	"time"

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

	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantRemoteRemoveItemCart
	}
	return nil
}

func BuyItemFromCart(ctx context.Context, userCollection *mongoCollection, userID string) error {
	//fetch the cart of the user
	//find the cart total
	//create an order with the items
	// added order to the user Collection
	// added items in the cart to the order list
	//empty the cart
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	var getcartitems models.User
	var ordercart models.Order
	ordercart.Order_ID = primitive.NewObjectID()
	ordercart.Ordered_At = time.Now()
	ordercart.Order_Cart = make([]models.ProductUser, 0)
	ordercart.Payment_Method.COD = true

	unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
	grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
	currentresults, err := userCollection.Aggregrate(ctx, mongo.Pipeline{unwind, grouping})
	ctx.Done()
	if err != nil {
		panic(err)
	}
	var getusercart []bson.M
	currentresults.All(ctx, &getusercart);err!=nil{
		panic(err)
	}
	var total_price int32
	for _, user_item := range getgetusercart{
		price := user_item["total"]
		total_price = price(int32)
	}
	ordercart.Price = int(total_price)

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key:"$push",Value:bson.D{primitive.E{Key:"orders", Value:ordercart}}}}

	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Println(err)
	}
	userCollection.FindOne(ctx, bson.D{primitive.E{Key:"_id",Value:id}}).Decode(&getcartitems)
	if err != nil {
		log.Println(err)
	}
	filter2 := bson.D{primitive.E{Key:"_id",Value:id}}
	update2 := bson.M{"$push":bson.M{"orders.$[].order_list":bson.M{"$each":getcartitems.UserCart}}}
	_, err = userCallection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		log.Println(err)
	}
	usercart_empty := make([]models.ProductUser, 0)

	filter3 := bson.D{primitive.E{Key: "_id", Value: id}}
	update3 := bson.D{{Key:"$set",Value:bson.D{primitive.E{Key:"usercart", Value:usercart_empty}}}}

	_, err = userCollection.UpdateOne(ctx, filter3, update3)
	if err != nil {
		return ErrCantBuyCartItem
	}
	return nil
}

func InstantBuyer(ctx context.Context, prodCollection, userCollection *mongo.Collection, productId primitive.ObjectID, UserID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	var product_details models.ProductUser
	var orders_detail models.Order
	orders_detail.Order_ID = primitive.NewObjectID()
	orders_detail.Ordered_At = time.Now()
	order_detail.Order_Cart = make([]models.ProductUser, 0)
	orders_detail.Payment_Method.COD = true
	err = prodCollection.FindOne(ctx, bson.D{primitive.E{Key:"_id",Value: productID}}).Decode(&product_details)
	if err != nil{
		log.Println(err)
	}
	orders_detail.Price	= product_details.Price

	filter := bson.D{primitive.E{Key:"_id",Value:id}}
	update := bson.D{{Key:"$push",Value:bson.D{primitive.E{Key:"orders", Value:orders_detail}}}}
	_, err = userCollection.UpdateOne(ctx, filter3, update3)
	if err != nil {
		log.Println(err)
	}
	filter2 := bson.D{primitive.E{Key:"_id",Value:id}}
	update2 := bson.M{"$push":bson.M{"orders.$[].order_list":product_details}}
	_, err = userCollection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		log.Println(err)
	}
	return nil 




}

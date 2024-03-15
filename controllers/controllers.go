package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBasRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
	}

	count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil{
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the email"})
		return
	}
	if count > 0{
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	}

	count, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.phone})
	defer cancel()

	if err != nil{
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the phone number"})
	}
	if count > 0{
		c.JSON(http.StatusConflict, gin.H{"error": "Phone Number already in use"})
		return
	}

}
}

func Login() gin.HandleFunc {

}

func ProductViewerAdmin() gin.HandleFunc {

}

func SearchProduct() gin.HandleFunc {

}

func SearchProductByQuery() gin.HandleFunc {

}

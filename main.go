package main

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/orm"
	"github.com/nextwebb/go-gin-pg/models"
	"net/http"
)
var ORM orm.Ormer

//The init function executes before the main function is executed

func init() {
	// makes a new connection the DB
	models.ConnectToDb()
	// /gets the ORM object and stores it in the global variable ORM
	ORM = models.GetOrmObject()
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// fix CORS error -- allow all headers
	// gin middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.POST("api/v1/users/createUser", createUser)
	router.GET("api/v1/users/readUsers", readUsers)
	router.PUT("api/v1/users/updateUser", updateUser)
	router.DELETE("api/v1/users/deleteUser", deleteUser)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":3000")
	// router.Run(":3000") for a hard coded port
}

func createUser(c *gin.Context) {
	// we are creating a variable of type user and binding that with the gin context
	//so that we can parse the user information which we add to the body of the API request,
	var newUser models.Users
	c.BindJSON(&newUser)
	_, err := ORM.Insert(&newUser)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK, 
			"email": newUser.Email,
			"user_name": newUser.UserName, 
			"user_id": newUser.UserId})
	} else {
		c.JSON(http.StatusInternalServerError, 
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user", "message": err})
	} 
}

func readUsers(c *gin.Context) {
	var user []models.Users // struct of users
	_, err := ORM.QueryTable("users").All(&user)
	if(err == nil) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": &user})
	} else {
		c.JSON(http.StatusInternalServerError, 
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the users"})
	}
}

func updateUser(c *gin.Context) {
	var updateUser models.Users
	c.BindJSON(&updateUser)
	_, err := ORM.QueryTable("users").Filter("email", updateUser.Email).Update(
		orm.Params{"user_name": updateUser.UserName, 
			"password": updateUser.Password,})
	if(err == nil) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError, 
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to update the users"})
	}
}

func deleteUser(c *gin.Context) {
	var delUser models.Users
	c.BindJSON(&delUser)
	_, err := ORM.QueryTable("users").Filter("email", delUser.Email).Delete()
	if(err == nil) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError, 
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to delete the users"})
	}
}
package main

import (
	"net/http"
	"os"
	"shopping/controllers/account"
	"shopping/controllers/auth"
	"shopping/controllers/product"
	"shopping/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	mysql.Connect()
	mysql.Migrate()

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Web API Http Server is running...")
	})

	api := router.Group("/api")
	{
		api.POST("/auth/login", auth.Login)
		api.POST("/auth/register", auth.Register)
		api.POST("/auth/updateAccessToken", auth.UpdateAccessToken)
		api.POST("/auth/removeRefreshToken", auth.RemoveRefreshToken)

		api.GET("/account/users", account.Users)

		api.GET("/categories", product.CategoryFindAll)
		api.GET("/categories/:id", product.CategoryFindById)
		api.POST("/categories/save", product.CategorySave)
		api.PUT("/categories/update/:id", product.CategoryUpdate)
		api.DELETE("/categories/remove/:id", product.CategoryRemove)

		api.GET("/products", product.FindAll)
		api.GET("/products/:id", product.FindById)
		api.POST("/products/save", product.Save)
		api.PUT("/products/update/:id", product.Update)
		api.DELETE("/products/remove/:id", product.Remove)
	}

	godotenv.Load()

	serverUri := os.Getenv("SERVER_URI")
	serverPort := os.Getenv("SERVER_PORT")

	router.Run(serverUri + ":" + serverPort)
}

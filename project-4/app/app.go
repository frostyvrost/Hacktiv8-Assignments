package app

import (
	"project-4/auth"
	"project-4/dto"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var PORT = ":8080"

// @Title Toko-Tokoan
// @version 1.0
// @description Project-4 Vormes Gema Merdeka API for E-Commerce
// @host localhost:8080
// @BasePath /
func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", dto.Register)
		userRouter.POST("/login", dto.Login)
		userRouter.Use(auth.Authentication())
		userRouter.PATCH("/topup", dto.UpdateBalance)
	}

	categoriesRouter := router.Group("/categories")
	{
		categoriesRouter.Use(auth.Authentication())
		categoriesRouter.POST("/", auth.AdminAuthorization(), dto.CreateCategory)
		categoriesRouter.GET("/", dto.GetAllCategories)
		categoriesRouter.PATCH("/:categoryId", auth.AdminAuthorization(), auth.CategoryAuthorization(), dto.UpdateCategory)
		categoriesRouter.DELETE("/:categoryId", auth.AdminAuthorization(), auth.CategoryAuthorization(), dto.DeleteCategory)
	}

	productsRouter := router.Group("/products")
	{
		productsRouter.Use(auth.Authentication())
		productsRouter.POST("/", auth.AdminAuthorization(), dto.CreateProduct)
		productsRouter.GET("/", dto.GetAllProducts)
		productsRouter.PUT("/:productId", auth.AdminAuthorization(), auth.ProductAuthorization(), dto.UpdateProduct)
		productsRouter.DELETE("/:productId", auth.AdminAuthorization(), auth.ProductAuthorization(), dto.DeleteProduct)
	}

	transactionHistoryRouter := router.Group("/transactions")
	{
		transactionHistoryRouter.Use(auth.Authentication())
		transactionHistoryRouter.POST("/", dto.CreateTransaction)
		transactionHistoryRouter.GET("/my-transactions", dto.GetTransactionsByUserID)
		transactionHistoryRouter.GET("/user-transactions", auth.AdminAuthorization(), dto.GetAllTransaction)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}

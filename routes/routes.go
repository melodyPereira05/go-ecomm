package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/melodyPereira05/go-ecomm/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/users/search", controllers.SearchProuctByQuery())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/viewproduct", controllers.SearchProduct())
}

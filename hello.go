package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	//"github.com/ydhnwb/golang_api/controller"
	"gorm.io/gorm"
	//"fmt"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	//authController controller.AuthController = controller.NewAuthController()
)

// func main() {
// 	fmt.Println("hello word123!")
// }

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
        authRoutes.POST("/login")
		authRoutes.POST("/register")
	}
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

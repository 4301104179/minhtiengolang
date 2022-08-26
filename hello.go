package main

import (
	"github.com/gin-gonic/gin"
	"github.com/4301104179/minhtiengolang/config"
	"github.com/4301104179/minhtiengolang/controller"
	"github.com/4301104179/minhtiengolang/repository"
	"github.com/4301104179/minhtiengolang/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService)
)

// func main() {
// 	fmt.Println("hello word123!")
// }

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
        authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	bookRoutes := r.Group("api/books")
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

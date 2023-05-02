package routes

import (
	"mini-project/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Routes
	e.GET("/roles", controllers.GetRoles)
	e.GET("/roles/:id", controllers.GetRole)
	e.POST("/roles", controllers.CreateRole)
	e.PUT("/roles", controllers.UpdateRole)
	e.DELETE("/roles", controllers.DeleteRole)

	e.GET("/events", controllers.GetEvents)
	e.GET("/events/:id", controllers.GetEvent)
	e.POST("/events", controllers.CreateEvent)
	e.PUT("/events", controllers.UpdateEvent)
	e.DELETE("/events", controllers.DeleteEvent)

	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	e.GET("/blogs", controllers.GetBlogs)
	e.GET("/blogs/:id", controllers.GetBlog)
	e.POST("/blogs", controllers.CreateBlog)
	e.PUT("/blogs", controllers.UpdateBlog)
	e.DELETE("/blogs", controllers.DeleteBlog)

	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products", controllers.UpdateProduct)
	e.DELETE("/products", controllers.DeleteProduct)

	e.GET("/divisions", controllers.GetDivisions)
	e.GET("/divisions/:id", controllers.GetDivision)
	e.POST("/divisions", controllers.CreateDivision)
	e.PUT("/divisions", controllers.UpdateDivision)
	e.DELETE("/divisions", controllers.DeleteDivision)

	return e
}

package routes

import (
	constant "mini-project/constant"
	"mini-project/controllers"
	"mini-project/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	u := e.Group("/users")
	u.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	u.GET("", controllers.GetUsersController)
	u.GET("/:id", controllers.GetUserController)
	u.PUT("/:id", controllers.UpdateUserController)
	u.DELETE("/:id", controllers.DeleteUserController)

	ua := e.Group("/user-activities")
	ua.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	ua.GET("", controllers.GetUserActivityController)
	ua.GET("/:id", controllers.GetUserActivityByUserController)
	ua.POST("", controllers.CreateUserActivityController)
	ua.PUT("/:id", controllers.UpdateUserActivityController)
	ua.DELETE("/:id", controllers.DeleteUserActivityController)

	r := e.Group("/roles")
	r.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	r.GET("", controllers.GetRolesController)
	r.GET("/:id", controllers.GetRoleController)
	r.POST("", controllers.CreateRoleController)
	r.PUT("/:id", controllers.UpdateRoleController)
	r.DELETE("/:id", controllers.DeleteRoleController)

	p := e.Group("/products")
	p.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	p.GET("", controllers.GetProductsController)
	p.GET("/:id", controllers.GetProductController)
	p.POST("", controllers.CreateProductController)
	p.PUT("/:id", controllers.UpdateProductController)
	p.DELETE("/:id", controllers.DeleteProductController)

	ev := e.Group("/events")
	ev.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	ev.GET("", controllers.GetEventsController)
	ev.GET("/:id", controllers.GetEventController)
	ev.POST("", controllers.CreateEventController)
	ev.PUT("/:id", controllers.UpdateEventController)
	ev.DELETE("/:id", controllers.DeleteEventController)

	d := e.Group("/divisions")
	d.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	d.GET("", controllers.GetDivisionsController)
	d.GET("/:id", controllers.GetDivisionController)
	d.POST("", controllers.CreateDivisionController)
	d.PUT("/:id", controllers.UpdateDivisionController)
	d.DELETE("/:id", controllers.DeleteDivisionController)

	b := e.Group("/blogs")
	b.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	b.GET("", controllers.GetBlogsController)
	b.GET("/:id", controllers.GetBlogController)
	b.POST("", controllers.CreateBlogController)
	b.PUT("/:id", controllers.UpdateBlogController)
	b.DELETE("/:id", controllers.DeleteBlogController)

	return e
}

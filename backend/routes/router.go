package routes

import (
	"github.com/KouT127/attendance-management/backend/database"
	. "github.com/KouT127/attendance-management/backend/handlers"
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/repositories"
	. "github.com/KouT127/attendance-management/backend/usecases"
	"github.com/gin-contrib/cors"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func renderIndex(c *Context) {
	c.HTML(http.StatusOK, "index.html", H{})
}

func defaultRouter(r *Engine) {
	r.LoadHTMLFiles("frontend/build/index.html")
	r.NoRoute(renderIndex)
}

func v1AttendancesRouter(v1 *RouterGroup) {
	handlers := []HandlerFunc{
		middlewares.AuthRequired(),
	}
	r := NewAttendanceRepository()
	u := NewAttendanceUsecase(r)
	c := NewAttendanceHandler(u)

	attendances := v1.Group("/attendances", handlers...)
	attendances.GET("", c.AttendanceListHandler)
	attendances.POST("", c.AttendanceCreateHandler)
	attendances.GET("monthly", c.AttendanceMonthlyHandler)
}

func v1UsersRouter(v1 *RouterGroup) {
	handlers := []HandlerFunc{
		middlewares.AuthRequired(),
	}
	engine := database.NewDB()
	r := NewUserRepository(*engine)
	i := NewUserUsecase(r)
	c := NewUserHandler(i)

	users := v1.Group("/users", handlers...)
	users.GET("/mine", c.UserMineHandler)
	users.PUT("/:id", c.UserUpdateHandler)
}

func v1Router(r *Engine) {
	v1Group := r.Group("/v1")
	v1UsersRouter(v1Group)
	v1AttendancesRouter(v1Group)
}

func Init() {
	r := New()
	r.Use(Logger())
	r.Use(Recovery())
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTION", "GET", "POST", "PUT", "DELETE"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	r.StaticFS("/static", http.Dir("frontend/build/static"))

	v1Router(r)
	defaultRouter(r)
	panic(r.Run(":8080"))
}

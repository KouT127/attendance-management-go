package routes

import (
	"github.com/KouT127/attendance-management/database"
	. "github.com/KouT127/attendance-management/handlers"
	"github.com/KouT127/attendance-management/middlewares"
	. "github.com/KouT127/attendance-management/repositories"
	. "github.com/KouT127/attendance-management/usecases"
	"github.com/gin-contrib/cors"
	. "github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func renderIndex(c *Context) {
	c.HTML(http.StatusOK, "index.html", H{})
}

func defaultRouter(r *Engine) {
	r.GET("/health", func(ctx *Context) {
		ctx.JSON(http.StatusOK, "ok")
		return
	})
	//r.LoadHTMLFiles("frontend/build/index.html")
	//r.NoRoute(renderIndex)
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
	r := Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTION", "GET", "POST", "PUT", "DELETE"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	r.StaticFS("/static", http.Dir("frontend/build/static"))

	v1Router(r)
	defaultRouter(r)
	http.Handle("/", r)
	log.Fatal(r.Run(":" + port))
}

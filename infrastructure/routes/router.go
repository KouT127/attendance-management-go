package routes

import (
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func configureDefaultRouter(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
}

func configureV1Router(r *gin.Engine, store sqlstore.SqlStore) {
	group := r.Group("/v1")
	configureUsersRouter(group, store)
	configureAttendancesRouter(group, store)
}

func InitRouter(store sqlstore.SqlStore) {
	r := gin.Default()
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

	configureV1Router(r, store)
	configureDefaultRouter(r)
	http.Handle("/", r)
	log.Fatal(r.Run(":" + port))
}

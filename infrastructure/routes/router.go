package routes

import (
	"github.com/KouT127/attendance-management/database"
	v1 "github.com/KouT127/attendance-management/infrastructure/routes/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func defaultRouter(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		err := database.Ping()
		if err != nil {
			log.Printf("Don't connection: %v", err)
			ctx.JSON(http.StatusBadRequest, "ng")
			return
		}
		ctx.JSON(http.StatusOK, "ok")
	})
}

func Init() {
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

	v1.Router(r)
	defaultRouter(r)
	http.Handle("/", r)
	log.Fatal(r.Run(":" + port))
}

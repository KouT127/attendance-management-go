package routes

import (
	"github.com/KouT127/attendance-management/api/handler/v1/image"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/infrastructure/uploader"
	"github.com/gin-gonic/gin"
)

func configureImagesRouter(v1 *gin.RouterGroup, store sqlstore.SQLStore, upl uploader.Uploader) {
	//funcs := []gin.HandlerFunc{
	//	middlewares.AuthRequired(),
	//}
	handler := image.NewImageHandler(upl)
	image := v1.Group("/images")
	image.POST("user", handler.UploadUserImageHandler)
}

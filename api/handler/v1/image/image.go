package image

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/KouT127/attendance-management/infrastructure/uploader"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"image"
	"net/http"
)

type Handler interface {
	UploadUserImageHandler(c *gin.Context)
}

type handler struct {
	uploader uploader.Uploader
}

func NewImageHandler(uploader uploader.Uploader) Handler {
	return &handler{
		uploader,
	}
}

const bucketName = "attendance-manament-d"

func (h *handler) UploadUserImageHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	decoded, extension, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return

	}

	uuid := uuid.NewV4().String()
	hash := md5.Sum([]byte(uuid))
	filename := hex.EncodeToString(hash[:])
	path := fmt.Sprintf("photos/users/%s/%s.%s", uuid, filename, extension)
	url, err := h.uploader.UploadFromImage(bucketName, path, decoded)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, url)
}

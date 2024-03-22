package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"image-prep/service"
	"image-prep/service/request"

	"github.com/gin-gonic/gin"
)

func Write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}

	content, _ := json.Marshal(data)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	_, _ = w.Write(content)
}

type imageHandler struct {
	shipmentService service.ImageService
}

func NewImageHandler(
	shipmentService service.ImageService,
) *imageHandler {
	return &imageHandler{
		shipmentService: shipmentService,
	}
}

func (c *imageHandler) ConvertHandler(ctx *gin.Context) {
	form := new(request.ConvertImageRequest)
	c.shipmentService.ConvertImage(ctx, form)
	ctx.IndentedJSON(http.StatusOK, `{"success":true}`)
}

func (c *imageHandler) ResizeHandler(ctx *gin.Context) {
	form := new(request.ResizeImageRequest)
	c.shipmentService.ResizeImage(ctx, form)
	ctx.IndentedJSON(http.StatusOK, "Hello, world ResizeHandler!\n")
}

func (c *imageHandler) CompressHandler(ctx *gin.Context) {
	form := new(request.CompressImageRequest)
	c.shipmentService.CompressImage(ctx, form)
	ctx.IndentedJSON(http.StatusOK, "Hello, world CompressHandler!\n")
}

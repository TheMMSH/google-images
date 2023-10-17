package server

import (
	"github.com/gin-gonic/gin"
	"google-images/service"
	"net/http"
	"strconv"
)

const DefaultCount = 10

type Server struct {
	DownloaderService service.IDownloaderService
	PresenterService  service.IPresenterService
}

func (s *Server) SetRoutes(e *gin.Engine) {
	e.GET("/health", healthCheckHandler)
	e.GET("/download/start", s.startDownload)
	e.GET("/image/:id/view", s.viewImage)
}

func (s *Server) startDownload(c *gin.Context) {
	query := c.GetString("query")
	count := c.GetInt("count")
	if count < 1 {
		count = DefaultCount
	}

	go s.DownloaderService.ProcessImagesConcurrently(query, count)

	c.JSON(http.StatusOK, gin.H{"msg": "downloading task started"})
}

func (s *Server) viewImage(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "please indicate a positive number for image id"})
		c.Abort()
	}

	imageData, err := s.PresenterService.ViewImage(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "cannot view this image", "err": err.Error()})
		c.Abort()
	}

	c.Header("Content-Disposition", "attachment; filename=image_"+idStr+".jpg")
	c.Data(http.StatusOK, "image/png", imageData)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

type (
	FileApi struct{}
)

var (
	allowedExtensions = []string{".txt"}
)

func isValidExtension(filename string) bool {
	ext := filepath.Ext(filename)
	for _, e := range allowedExtensions {
		if strings.EqualFold(ext, e) {
			return true
		}
	}

	return false
}

func (f *FileApi) UpdateFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := file.Filename

	if !isValidExtension(filename) {
		c.JSON(http.StatusBadRequest, &gin.H{"error": "Invalid file extension"})
		return
	}

	fh, err := file.Open()
	if nil != err {
		c.JSON(http.StatusBadRequest, &gin.H{"error": "Invalid file extension Not Open"})
		return
	}

	fileService.UpdateFile(fh)
}

func (f *FileApi) ReLoad(c *gin.Context) {
	fileService.ReLoad()
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

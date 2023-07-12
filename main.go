package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//gin set up
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "upload images",
		})
	})

	r.POST("/", func(c *gin.Context) {

		file, err := c.FormFile("image")
		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "failed to upload the image",
			})
			return
		}

		err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "failed to save the image",
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "assets/uploads/" + file.Filename,
		})
	})

	r.Run()
}

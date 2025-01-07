// GIN FRAMEWORK
// type in terminal to use gin framework": go get -u github.com/gin-gonic/gin
//type in terminal: go mod init sportsSite
// GIN Handles logging

package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Status is OK")

	})

	//TO RETURN THINGS TO SCREEN, MUST USE c.JSON or c.String?

	r.GET("/health", func(c *gin.Context) {
		//format using .Format
		currentTime := time.Now().Format("Sat Jan 02 15:04:05 2006")

		// c.JSON(http.StatusOK, gin.H{
		// 	"date&time": currentTime,
		// })
		c.String(http.StatusOK, currentTime)

	})

	//SHOULD BE DONE WITH GORM
	r.POST("/echo", func(c *gin.Context) {
		var doota map[string]string
		//binds JSON body into the doota variable
		if err := c.ShouldBindJSON(&doota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "400"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"echo": doota["message"],
		})

	})

	r.Run(":8080")

}

func handler(c *gin.Context) {
	c.String(http.StatusOK, "Status is OK 200")

}

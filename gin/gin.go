package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Static("/assets", "./client/dist/assets")
	router.StaticFile("/vite.svg", "./client/dist/vite.svg")
	router.LoadHTMLFiles("./client/dist/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}

package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testapp/models/test"
)

//Init initialize all the routes
func Init() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.GET("/tests", func(c *gin.Context) {
		c.JSON(http.StatusOK, test.All())
	})

	// Listen and Server in http://0.0.0.0:8080
	r.Run()
}

package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/testapp/controllers/optioncontroller"
	"github.com/testapp/controllers/questioncontroller"
	"github.com/testapp/controllers/testcontroller"
)

//Init initialize all the routes
func Init() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.GET("/tests", testcontroller.Tests)
	r.POST("/tests", testcontroller.AddTest)
	r.PUT("/tests/:id", testcontroller.UpdateTest)
	r.DELETE("/tests/:id", testcontroller.DeleteTest)

	r.POST("/question", questioncontroller.AddQuestion)
	r.PUT("/question/:id", questioncontroller.UpdateQuestion)
	r.DELETE("/questions/:id", questioncontroller.DeleteQuestion)

	r.POST("/option", optioncontroller.AddOption)
	r.PUT("/option/:id", optioncontroller.UpdateOption)
	r.DELETE("/option/:id", optioncontroller.DeleteOption)

	// Listen and Server in http://0.0.0.0:8080
	r.Run()
}

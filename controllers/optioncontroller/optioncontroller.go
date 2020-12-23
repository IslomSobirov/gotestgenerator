package optioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testapp/models/test"
)

//AddOption add option to the database
func AddOption(c *gin.Context) {
	var optionJSON test.Option
	c.BindJSON(&optionJSON)

	if optionJSON.OptionName == "" || optionJSON.QuestionID == 0 || optionJSON.TestID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Option name question id and test id is required to create option of the quesion",
		})
	}

	test.CreateOption(
		optionJSON.OptionName,
		optionJSON.TrueOption,
		optionJSON.TestID,
		optionJSON.QuestionID,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Option has been created successfully",
	})

}

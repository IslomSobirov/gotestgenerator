package questioncontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/testapp/models/test"
)

//AddQuestion insert question to database
func AddQuestion(c *gin.Context) {
	var quesionJSON test.Question
	c.BindJSON(&quesionJSON)
	if quesionJSON.TestID == 0 || quesionJSON.QuestionName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "QuestionName and testID is required to create test",
		})
		return
	}

	test.CreateQuestion(quesionJSON.QuestionName, quesionJSON.TestID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Question has been created successfully",
	})

}

//UpdateQuestion update question
func UpdateQuestion(c *gin.Context) {
	var quesionJSON test.Question
	c.BindJSON(&quesionJSON)
	id, _ := strconv.Atoi(c.Param("id"))
	if quesionJSON.TestID == 0 || quesionJSON.QuestionName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "QuestionName and testID is required to create test",
		})
		return
	}

	test.UpdateQuestion(id, quesionJSON.TestID, quesionJSON.QuestionName)

	c.JSON(http.StatusOK, gin.H{
		"message": "test has been updated successfully",
	})
}

//DeleteQuestion delete question and its options
func DeleteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	test.DeleteQuestion(id, false)
	test.DeleteOption(id, true)
}

//QuestionOptions get options of questions
func QuestionOptions(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, test.OptionByQuestionID(id))
}

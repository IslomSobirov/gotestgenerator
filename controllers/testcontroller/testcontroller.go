package testcontroller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/testapp/models/test"
)

//Tests Get all tests
func Tests(c *gin.Context) {
	c.JSON(http.StatusOK, test.All())
}

//AddTest insert test to database
func AddTest(c *gin.Context) {
	var testJSON test.Test
	c.BindJSON(&testJSON)
	if testJSON.TestName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "TestName is required to create test",
		})
		return
	}

	test.CreateTest(testJSON.TestName)

	c.JSON(http.StatusOK, gin.H{
		"message": "test has been created successfully",
	})
}

//UpdateTest update test by id
func UpdateTest(c *gin.Context) {
	var testJSON test.Test
	c.BindJSON(&testJSON)
	id, _ := strconv.Atoi(c.Param("id"))
	if testJSON.TestName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "TestName is required to create test",
		})
		return
	}

	test.UpdateTest(id, testJSON.TestName)

	c.JSON(http.StatusOK, gin.H{
		"message": "test has been updated successfully",
	})
}

//DeleteTest delete test from database
func DeleteTest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	test.DeleteTest(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "test has been deleted successfully",
	})
}

//TestOptions Get Options by test id
func TestOptions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, test.OptionByTestID(id))
}

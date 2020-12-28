package test

import (
	"log"

	"github.com/testapp/db"
	"github.com/testapp/helper"
)

//Test struct for the model
type Test struct {
	ID        int        `json:"id"`
	TestName  string     `json:"testName"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	Questions []Question `json:"questions"`
}

var testTable = "test"

//CreateTest insert test to test table
func CreateTest(testName string) {
	db := db.Connect()
	defer db.Close()
	_, err := db.Query(
		"Insert into "+testTable+" (testName, createdAt) values (?, NOW())",
		testName,
	)

	if err != nil {
		helper.LogError(err)
		log.Fatalf(err.Error())
	}
}

//UpdateTest update test by id
func UpdateTest(id int, testName string) {
	db := db.Connect()
	defer db.Close()
	_, err := db.Query(
		"Update "+testTable+" set testName = ?, updatedAt = NOW()",
		testName,
	)
	if err != nil {
		helper.LogError(err)
		log.Fatalf(err.Error())
	}
}

//DeleteTest delete test from test table
func DeleteTest(id int) {
	db := db.Connect()
	defer db.Close()

	_, err := db.Query(
		"Delete from "+testTable+" where id = ?",
		id,
	)
	DeleteQuestion(id, true)
	if err != nil {
		helper.LogError(err)
		log.Fatalf(err.Error())
	}
}

//All tests
func All() []Test {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query(
		"SELECT id, testName, createdAt, updatedAt from " + testTable,
	)
	if err != nil {
		helper.LogError(err)
		log.Fatalf(err.Error())
	}
	var test Test
	var tests []Test

	for rows.Next() {
		rows.Scan(
			&test.ID,
			&test.TestName,
			&test.CreatedAt,
			&test.UpdatedAt,
		)

		tests = append(tests, test)
		test.Questions = QuestionByTestID(test.ID)
	}

	return tests
}

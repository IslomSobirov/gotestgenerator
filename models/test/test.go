package test

import "github.com/testapp/db"

//Test struct for the model
type Test struct {
	id        int
	testName  string
	createdAt string
	updatedAt string
	questions []Question
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
		panic(err)
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
		panic(err)
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

	if err != nil {
		panic(err)
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
		panic(err)
	}
	var test Test
	var tests []Test

	for rows.Next() {
		rows.Scan(
			&test.id,
			&test.testName,
			&test.createdAt,
			&test.updatedAt,
		)

		test.questions = QuestionByTestID(test.id)
	}

	tests = append(tests, test)

	return tests
}

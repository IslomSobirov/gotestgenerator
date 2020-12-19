package test

import "github.com/testapp/db"

//Question for Test
type Question struct {
	id           int
	questionName string
	testID       int
	createdAt    string
	updatedAt    string
	options      []Option
}

const questionTable = "test_question"

//CreateQuestion insert question to test_question table
func CreateQuestion(questionName string, testID int) {
	db := db.Connect()
	defer db.Close()
	_, err := db.Query(
		"Insert into "+questionTable+" (questionName, testID, createdAt) values (?, ?, NOW())",
		questionName,
		testID,
	)

	if err != nil {
		panic(err)
	}
}

//DeleteQuestion delete quesion by id or quesionID
func DeleteQuestion(id int, testID bool) {
	db := db.Connect()
	defer db.Close()
	var idType string
	if testID {
		idType = "testID"

	} else {
		idType = "id"
	}

	_, err := db.Query(
		"DELETE FROM "+questionTable+" where "+idType+" = ?",
		id,
	)
	if err != nil {
		panic(err)
	}
}

//QuestionByTestID get questions by testID
func QuestionByTestID(testID int) []Question {
	db := db.Connect()
	defer db.Close()
	rows, err := db.Query(
		"SELECT id, questionName, testID, createdAt, updatedAt FROM "+questionTable+" where testID = ?",
		testID,
	)

	if err != nil {
		panic(err)
	}

	var question Question
	var questions []Question

	for rows.Next() {
		rows.Scan(
			&question.id,
			&question.questionName,
			&question.testID,
			&question.createdAt,
			&question.updatedAt,
		)

		question.options = OptionByQuestionID(question.id)
	}

	questions = append(questions, question)

	return questions
}

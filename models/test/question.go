package test

import "github.com/testapp/db"

//Question for Test
type Question struct {
	ID           int      `json:"id"`
	QuestionName string   `json:"questionName"`
	TestID       int      `json:"testID"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
	Options      []Option `json:"options"`
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
			&question.ID,
			&question.QuestionName,
			&question.TestID,
			&question.CreatedAt,
			&question.UpdatedAt,
		)

		question.Options = OptionByQuestionID(question.ID)
		questions = append(questions, question)
	}

	return questions
}

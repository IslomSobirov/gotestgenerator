package test

import "github.com/testapp/db"

//Option of the Question
type Option struct {
	ID         int    `json:"id"`
	OptionName string `json:"optionName"`
	TrueOption bool   `json:"trueOption"`
	TestID     int    `json:"testID"`
	QuestionID int    `json:"questionID"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

const optionTable = "test_option"

//CreateOption insert option to test_option table
func CreateOption(optName string, trueOption bool, testID int, questionID int) {
	db := db.Connect()
	defer db.Close()
	_, err := db.Query("Insert into "+optionTable+
		" (optionName, trueOption, testID, questionID, createdAt) value (?, ?, ?, ?, NOW()) ",
		optName, trueOption, testID, questionID)
	if err != nil {
		panic(err)
	}

}

//UpdateOption update the option of question
func UpdateOption(id int, optName string, trueOption bool, testID int, questionID int) {
	db := db.Connect()
	_, err := db.Query(
		"UPDATE "+optionTable+
			" set optionName = ?, trueOption = ?, testID = ?, questionID = ?, updatedAt = NOW() where id = ?",
		optName,
		trueOption,
		testID,
		questionID,
		id,
	)
	defer db.Close()

	if err != nil {
		panic(err)
	}
}

//DeleteOption delete from test_option
func DeleteOption(id int) {
	db := db.Connect()
	defer db.Close()
	_, err := db.Query(
		"DELETE FROM "+optionTable+" where id = ?",
		id,
	)
	if err != nil {
		panic(err)
	}
}

//OptionByQuestionID get options by question id
func OptionByQuestionID(qID int) []Option {
	db := db.Connect()
	defer db.Close()
	rows, err := db.Query(
		"Select id, optionName, trueOption, testID, questionID, createdAt, updatedAt from "+
			optionTable+" where questionID = ? ",
		qID,
	)
	if err != nil {
		panic(err)
	}
	var Options []Option
	var Opt Option
	for rows.Next() {
		rows.Scan(
			&Opt.ID,
			&Opt.OptionName,
			&Opt.TrueOption,
			&Opt.TestID,
			&Opt.QuestionID,
			&Opt.CreatedAt,
			&Opt.UpdatedAt,
		)
		Options = append(Options, Opt)
	}

	return Options
}

package models

type Question struct {
	Date string `json:"date"`
	Data []struct {
		QuestionID string `json:"question_id"`
		Question   string `json:"question"`
		Answers    []struct {
			AnswerID string `json:"answer_id"`
			Answer   string `json:"answer"`
		} `json:"answers"`
	} `json:"data"`
}
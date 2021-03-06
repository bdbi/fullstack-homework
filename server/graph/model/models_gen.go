// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Answer interface {
	IsAnswer()
}

type Question interface {
	IsQuestion()
}

type AnswerInput struct {
	QuestionID string  `json:"questionID"`
	Text       *string `json:"text"`
	OptionID   *string `json:"optionID"`
}

type ChoiceAnswer struct {
	ID             string `json:"id"`
	QuestionID     string `json:"questionID"`
	SelectedOption string `json:"selectedOption"`
}

func (ChoiceAnswer) IsAnswer() {}

type ChoiceQuestion struct {
	ID      string    `json:"id"`
	Body    string    `json:"body"`
	Weight  float64   `json:"weight"`
	Options []*Option `json:"options"`
}

func (ChoiceQuestion) IsQuestion() {}

type Option struct {
	ID     string  `json:"id"`
	Body   string  `json:"body"`
	Weight float64 `json:"weight"`
}

type TextAnswer struct {
	ID         string `json:"id"`
	QuestionID string `json:"questionID"`
	Text       string `json:"text"`
}

func (TextAnswer) IsAnswer() {}

type TextQuestion struct {
	ID     string  `json:"id"`
	Body   string  `json:"body"`
	Weight float64 `json:"weight"`
}

func (TextQuestion) IsQuestion() {}

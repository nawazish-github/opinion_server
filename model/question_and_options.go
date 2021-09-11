package model

type QuestionAndOptions struct {
	Question Question `json:"question"`
	Options  []Option `json:"options"`
}

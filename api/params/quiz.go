package params

type CreateQuizRequest struct {
	Name   string      `json:"name" validate:"required"`
	Phases []PhaseData `json:"phases" validate:"required,dive"`
}

type PhaseData struct {
	Name      string         `json:"name" validate:"required"`
	Questions []QuestionData `json:"questions" validate:"required,dive"`
}

type QuestionData struct {
	Text             string   `json:"text" validate:"required"`
	Type             string   `json:"type" validate:"required,oneof=img audio"`
	ImgURL           string   `json:"img_url,omitempty"`
	AudioURL         string   `json:"audio_url,omitempty"`
	IsMultipleChoice bool     `json:"is_multiple_choice"`
	Answers          []string `json:"answers" validate:"required,dive"`
	Choices          []string `json:"choices,omitempty"`
}

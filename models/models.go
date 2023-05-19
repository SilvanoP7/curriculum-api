package models

// Book represents data about a book.
type Pong struct {
	Status  string  `json:"status"`
}

type Subjects struct {
	Subject string  `json:"subject"`
        SubjectID string `json:"subjectId"`
        KeyStage int     `json:"keyStage"`
	PurposeOfStudy string `json:"purposeOfStudy"`
}

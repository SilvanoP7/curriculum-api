package models

// Book represents data about a book.
type Pong struct {
	Status  string  `json:"status"`
}

type DbTest struct {
	Status  string  `json:"status"`
}

type Subjects struct {
        SubjectID string `json:"subjectId"`
        Subject string  `json:"subject"`
        KeyStage int     `json:"keyStage"`
	PurposeOfStudy string `json:"purposeOfStudy"`
}

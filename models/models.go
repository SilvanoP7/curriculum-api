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

type SubjectContent struct {
        SubjectID string `json:"subjectId"`
        SubjectContent string `json:"subjectContent"`
        SubjectContentId string `json:"subjectContentId"`
}

type Objectives struct {
        SubjectContentId string `json:"subjectContentId"`
        ObjectiveId      string `json:"objectiveId"`
        Objective        string `json:"objective"`
}

type SubObjectives struct {
        ObjectiveId      string `json:"objectiveId"`
        SubOjectiveId    string `json:"subObjectiveId"`
        SubObjective     string `json:"subObjective"`
}

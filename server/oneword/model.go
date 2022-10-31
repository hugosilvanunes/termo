package oneword

type Word struct {
	ID   int
	Name string
}

type AttemptRequest struct {
	Word string `json:"word,omitempty" binding:"required"`
}

type AttemptStatus string

const (
	CORRECT AttemptStatus = "correct"
	WRONG                 = "wrong"
	SEMI                  = "semi"
)

type AttemptCharsResponse struct {
	Index  int           `json:"index"`
	Char   string        `json:"character,omitempty"`
	Status AttemptStatus `json:"status,omitempty"`
}

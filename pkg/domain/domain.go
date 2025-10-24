package domain

type TaskCreatePayload struct {
	Date    string `db:"date" json:"date"`
	Title   string `db:"title" json:"title"`
	Comment string `db:"comment" json:"comment"`
	Repeat  string `db:"repeat" json:"repeat"`
}

type TaskCreateResponse struct {
	ID    *string `json:"id,omitempty"`
	Error *string `json:"error,omitempty"`
}

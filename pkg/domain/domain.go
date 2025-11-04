package domain

type Task struct {
	Date    string `db:"date" json:"date"`
	Title   string `db:"title" json:"title"`
	Comment string `db:"comment" json:"comment"`
	Repeat  string `db:"repeat" json:"repeat"`
	ID      int    `db:"id" json:"id,string,omitempty"`
}

type TaskListResponse struct {
	Tasks []Task `json:"tasks"`
}

type TaskCreateResponse struct {
	ID    *string `json:"id,omitempty"`
	Error *string `json:"error,omitempty"`
}

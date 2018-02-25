package endpoint

type NewPostRequest struct {
	Author   string   `json:"author"`
	Sitename string   `json:"sitename"`
	Title    string   `json:"title"`
	Date     string   `json:"date"`
	Tags     []string `json:"tags"`
	Summary  string   `json:"summary"`
	Content  string   `json:"content"`
}

type NewPostResponse struct {
	Err error `json:"-"`
}

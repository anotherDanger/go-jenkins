package web

type Response struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

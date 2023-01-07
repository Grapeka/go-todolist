package types

type Todo struct {
	ID          string `json:"id"`
	Owner	   	string `json:"owner"`
	Title       string `json:"title"`
	Completed   bool   `json:"complete"`
	TimeCreated string `json:"time_created"`
}
package dtos

type TodoResponse struct {
	Id    string `json:"id" bson:"_id"`
	Title string `json:"title"`
}

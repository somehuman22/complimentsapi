package models

//Compliment ... the model for a compliment
type Compliment struct {
	ComplimentID int    `json:"complimentId"`
	Sender       string `json:"sender"`
	Receiver     string `json:"receiver"`
	Content      string `json:"content"`
	Public       bool   `json:"isPublic"`
}

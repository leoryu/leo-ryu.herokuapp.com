package model

type Introduction struct {
	Title    string `json:"title" bson:"title"`
	Subject  string `json:"subject" bson:"subject"`
	Abstract string `json:"abstract" bson:"abstract"`
}


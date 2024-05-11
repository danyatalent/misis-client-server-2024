package models

type Question struct {
	ID      string
	Text    string `json:"text" xml:"Question"`
	Answer  string `json:"answer" xml:"Answer"`
	Comment string `json:"comment" xml:"Comment"`
	Author  string `json:"author" xml:"Author"`
}

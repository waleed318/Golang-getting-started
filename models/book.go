package models

type Book struct {
	ID     int    `json:"ID"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Pbdate string `json:"pbdate"`
	// ISBN       string   `isbn`
	// Category   string   `category`
	// Publisher  string   `publisher`
}

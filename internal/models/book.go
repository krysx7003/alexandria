package models

import "time"

type Book struct {
	ID			string		`json:"id"`
	Title		string		`json:"title"`
	Series		Series		`json:"series" `
	Genres		[]Genre		`json:"genres"`
	Authors		[]Author	`json:"author"`
	Publisher	string		`json:"publisher"`
	PublishDate	string		`json:"publish_date"`
	CoverURL	string		`json:"cover_url"`
	MediaURL 	string		`json:"media_url"`
	ISBN		*ISBN		`json:"isbn"`
	Created		time.Time	`json:"created_at"`
	Updated		time.Time	`json:"updated_at"`
}

type Series struct {
	Name 		string		`json:"name"`
	Book 		int			`json:"book"`
	Ongoing 	bool   		`json:"is_completed"`
	TotalBooks	int 		`json:"total_books,omitempty"`
}


type Author struct {
	Name 		string		`json:"name"`
}

type ISBN struct {
    ISBN10	 	string 		`json:"isbn10,omitempty"`
    ISBN13 		string 		`json:"isbn13,omitempty"`
}

type Genre struct {
	Name 		string		`json:"name"`
}

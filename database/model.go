package main

import (
	"time"

	"github.com/google/uuid"
)

type Newsletter struct {
	Nltr_ID                uuid.UUID `db:"Nltr_ID"`
	Nltr_Name              string    `db:"Nltr_Name"`
	Nltr_Description       string    `db:"Nltr_Description"`
	Nltr_Inserted_Datetime time.Time `db:"Nltr_Inserted_Datetime,"`
	Nltr_Updated_Datetime  time.Time `db:"Nltr_Updated_Datetime"`
	Nltr_Author            string    `db:"Nltr_Author"`
	concurrency_stamp      uuid.UUID `db:"concurrency_stamp"`
}

func NewNewsletter(Nltr_Name, Nltr_Description, Nltr_Author string) *Newsletter {
	now := time.Now()
	return &Newsletter{
		Nltr_ID:                uuid.New(),
		Nltr_Name:              Nltr_Name,
		Nltr_Description:       Nltr_Description,
		Nltr_Inserted_Datetime: now,
		Nltr_Updated_Datetime:  now,
		Nltr_Author:            Nltr_Author,
		concurrency_stamp:      uuid.New(),
	}
}

func UpdateNewsletter(Nltr_ID, Nltr_Name, Nltr_Description, Nltr_Author string) *Newsletter {
	now := time.Now()
	return &Newsletter{
		Nltr_ID:                uuid.New(),
		Nltr_Name:              Nltr_Name,
		Nltr_Description:       Nltr_Description,
		Nltr_Inserted_Datetime: now,
		Nltr_Updated_Datetime:  now,
		Nltr_Author:            Nltr_Author,
		concurrency_stamp:      uuid.New(),
	}
}

type Post struct {
	Post_ID                uuid.UUID `db:"Post_ID"`
	Post_Name              string    `db:"Post_Name"`
	Post_Content           string    `db:"Post_Content"`
	Post_Publishing_Date   time.Time `db:"Post_Inserted_Datetime,"`
	Post_Inserted_Datetime time.Time `db:"Post_Inserted_Datetime,"`
	Post_Updated_Datetime  time.Time `db:"Post_Updated_Datetime"`
	Post_Author            string    `db:"Post_Author"`
	Nltr_ID                uuid.UUID `db:"Nltr_ID"`
	concurrency_stamp      uuid.UUID `db:"concurrency_stamp"`
}

func NewPost(Post_Name string, Post_Content string, Post_Author string, Nltr_ID uuid.UUID) *Post {
	now := time.Now()
	return &Post{
		Post_ID:                uuid.New(),
		Post_Name:              Post_Name,
		Post_Content:           Post_Content,
		Post_Publishing_Date:   now,
		Post_Inserted_Datetime: now,
		Post_Updated_Datetime:  now,
		Post_Author:            Post_Author,
		Nltr_ID:                Nltr_ID,
		concurrency_stamp:      uuid.New(),
	}
}

func UpdatePost(Post_ID uuid.UUID, Post_Name string, Post_Content string, Post_Author string, Nltr_ID uuid.UUID) *Post {
	now := time.Now()
	return &Post{
		Post_ID:                uuid.New(),
		Post_Name:              Post_Name,
		Post_Content:           Post_Content,
		Post_Publishing_Date:   now,
		Post_Inserted_Datetime: now,
		Post_Updated_Datetime:  now,
		Post_Author:            Post_Author,
		Nltr_ID:                Nltr_ID,
		concurrency_stamp:      uuid.New(),
	}
}

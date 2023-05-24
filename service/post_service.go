package service

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"newsletter-platform/database/model"
)

type NewsletterService struct {
	db *sqlx.DB
}

func NewNewsletterService(connectionString string) (*NewsletterService, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &NewsletterService{
		db: db,
	}, nil
}

// Create a newsletter in the database
func (s *NewsletterService) CreateNewsletter(newsletter *model.Newsletter) error {
	_, err := s.db.NamedExec(`
		INSERT INTO Newsletter (
			Nltr_ID,
			Nltr_Name,
			Nltr_Description,
			Nltr_Inserted_Datetime,
			Nltr_Updated_Datetime,
			Nltr_Author,
			concurrency_stamp
		) VALUES (
			:Nltr_ID,
			:Nltr_Name,
			:Nltr_Description,
			:Nltr_Inserted_Datetime,
			:Nltr_Updated_Datetime,
			:Nltr_Author,
			:concurrency_stamp
		)`, newsletter)
	if err != nil {
		return err
	}

	return nil
}

// Update a newsletter in the database
func (s *NewsletterService) UpdateNewsletter(newsletter *model.Newsletter) error {
	_, err := s.db.NamedExec(`
		UPDATE Newsletter
		SET Nltr_ID = :Nltr_ID,
			Nltr_Name = :Nltr_Name,
			Nltr_Description = :Nltr_Description,
			Nltr_Inserted_Datetime = :Nltr_Inserted_Datetime,
			Nltr_Updated_Datetime = :Nltr_Updated_Datetime,
			Nltr_Author = :Nltr_Author,
			concurrency_stamp = :concurrency_stamp
		WHERE Nltr_ID = :Nltr_ID
		AND Nltr_Author = :Nltr_Author`, newsletter)
	if err != nil {
		return err
	}

	return nil
}

// Delete a newsletter from the database.
func (s *NewsletterService) DeleteNewsletter(newsletterID uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM Newsletter WHERE Nltr_ID = $1", newsletterID)
	if err != nil {
		return err
	}

	return nil
}

// Create a new post in the database.
func (s *NewsletterService) CreatePost(post *model.Post) error {
	_, err := s.db.NamedExec(`
		INSERT INTO Post (
			Post_ID,
			Post_Title,
			Post_Content,
			Post_Publishing_Date,
			Post_Inserted_Datetime,
			Post_Updated_Datetime,
			Post_Author,
			Nltr_ID,
			concurrency_stamp
		) VALUES (
			:Post_ID,
			:Post_Title,
			:Post_Content,
			:Post_Publishing_Date,
			:Post_Inserted_Datetime,
			:Post_Updated_Datetime,
			:Post_Author,
			:Nltr_ID,
			:concurrency_stamp
		)`, post)
	if err != nil {
		return err
	}

	return nil
}

// Update an existing post in the database.
func (s *NewsletterService) UpdatePost(post *model.Post) error {
	_, err := s.db.NamedExec(`
		UPDATE Post
		SET Post_ID = :Post_ID,
			Post_Title = :Post_Title,
			Post_Content = :Post_Content,
			Post_Publishing_Date = :Post_Publishing_Date,
			Post_Inserted_Datetime = :Post_Inserted_Datetime,
			Post_Updated_Datetime = :Post_Updated_Datetime,
			Post_Author = :Post_Author,
			Nltr_ID = :Nltr_ID,
			concurrency_stamp = :concurrency_stamp
		WHERE Post_ID = :Post_ID
		AND Post_Author = :Post_Author`, post)
	if err != nil {
		return err
	}

	return nil
}

// Delete a post from the database.
func (s *NewsletterService) DeletePost(postID uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM Post WHERE Post_ID = $1", postID)
	if err != nil {
		return err
	}

	return nil
}

package repositories

import (
	"Post-platform/database"
	"Post-platform/database/model"
	"Post-platform/database/query"

	"github.com/google/uuid"
)

type PostRepository struct{}

// Method for adding a new post to the database.
func (PostRepository) AddPost(data *model.Post) (uuid.UUID, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return uuid.Nil, err
	}

	rows, err := ctx.NamedQuery(query.CreatePost, data)
	if err != nil {
		return uuid.Nil, err
	}

	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
	}
	if err != nil {
		return uuid.Nil, err
	}

	PostId, err := uuid.Parse(returnedId)
	if err != nil {
		return uuid.Nil, err
	}

	return PostId, nil
}

// Method for obtaining data about a specific post in the database.
func (PostRepository) GetPost(postId uuid.UUID) (model.PostData, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.PostData{}, err
	}

	var post model.PostData
	if err = ctx.Get(&post, query.GetPost, PostId); err != nil {
		return model.PostData{}, err
	}
	return post, err
}

// Method for updating specific Post data in the database.
func (PostRepository) UpdatePost(data model.UpdatePostData) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.UpdatePost, data); err != nil {
		return err
	}
	return nil
}

// Method for deleting a specific Post in the database.
func (PostRepository) DeletePost(PostId, authorId uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeletePost, PostId, authorId); err != nil {
		return err
	}
	return nil
}

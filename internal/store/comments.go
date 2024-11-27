package store

import (
	"context"
	"database/sql"
	"log"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CommentsSore struct {
	db *sql.DB
}

func (s *CommentsSore) GetByPostID(ctx context.Context, postID int64) (*Post, error) {
	log.Printf("GetByPostID: postID=%d", postID)
}

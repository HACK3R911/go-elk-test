package db

import (
	"database/sql"
	"fmt"
	"github.com/HACK3R911/go-elk-test/models"
)

var (
	ErrNoRecord = fmt.Errorf("no matching record found")
	insertOp    = "insert"
	deleteOp    = "delete"
	updateOp    = "update"
)

func (db Database) SavePost(post *models.Post) error {
	var id int
	query := `INSERT INTO posts(title, body) VALUES (&1, &2) RETURNING id`
	err := db.Conn.QueryRow(query, post.Title, post.Body).Scan(&id)
	if err != nil {
		return err
	}

	logQuery := `INSERT INTO post_logs(post_id, operation) VALUES ($1, $2)`
	post.ID = id
	_, err = db.Conn.Exec(logQuery, post.ID, insertOp)
	if err != nil {
		db.Logger.Err(err).Msg("could not log operation for logstash")
	}
	return nil
}

func (db Database) UpdatePost(postId int, post models.Post) error {
	query := "UPDATE posts SET title=$1, body=$2 WHERE id=$3"
	_, err := db.Conn.Exec(query, post.Title, post.Body, postId)
	if err != nil {
		return err
	}

	post.ID = postId
	logQuery := `INSERT INTO post_logs(post_id, operation) VALUES ($1, $2)`
	if _, err := db.Conn.Exec(logQuery, postId, updateOp); err != nil {
		db.Logger.Err(err).Msg("could not log operation for logstash")
	}
	return nil
}

func (db Database) DeletePost(postId int) error {
	query := "DELETE FROM posts WHERE id=$1"
	if _, err := db.Conn.Exec(query, postId); err != nil {
		if err == sql.ErrNoRows {
			return ErrNoRecord
		}
		return err
	}

	logQuery := `INSERT INTO post_logs(post_id, operation) VALUES ($1, $2)`
	if _, err := db.Conn.Exec(logQuery, postId, deleteOp); err != nil {
		db.Logger.Err(err).Msg("could not log operation for logstash")
	}
	return nil
}

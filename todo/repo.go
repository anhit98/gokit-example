package todo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Repository interface {
	GetAllForUser(username string) ([]Todo, error)
	GetByID(id string) (Todo, error)
	Create(req CreateRequest) (string, error)
	Update(id string, todo Todo) error
	Delete(id string) error
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Create(req CreateRequest) (string, error) {
	id, _ := uuid.NewV4()
	sql := `
		INSERT INTO todo (id, username, text, created_at)
		VALUES ($1, $2, $3, $4)`
	_, err := repo.db.Exec(sql, id.String(), req.Username, req.Text, time.Now().UTC())
	return id.String(), err
}

func (repo *repo) GetAllForUser(username string) (todos []Todo, err error) {

	rows, err := repo.db.Query("SELECT id, text, created_at FROM todo WHERE username=$1", username)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Text, &todo.CreatedAt)

		if err != nil {
			fmt.Println(err)
		}

		todos = append(todos, todo)
	}
	return
}

func (repo *repo) GetByID(id string) (todo Todo, err error) {
	_ = repo.db.QueryRow("SELECT email FROM todo WHERE id=$1", id)
	return
}

func (repo *repo) Update(id string, todo Todo) error {
	_, err := repo.db.Exec("UPDATE todo SET text = $2 where id=$1", id, todo.Text)
	return err
}

func (repo *repo) Delete(id string) error {
	_, err := repo.db.Exec("DELETE from todo where id=$1", id)
	return err
}

package store

import (
	"context"
	"fmt"
	"time"
	td "todo/core"
)

const (
	createQuery = `
	INSERT INTO todo
		(created_at, title, status )
	values
		($1, $2, $3)
	RETURNING
		id, created_at, title, status;
	`

	listQuery = `
	SELECT
		* 
	FROM 
		todo;
	`

	getByIdQuery = `
	SELECT
		* 
	FROM 
		todo
	WHERE 
		id =$1;
	`

	updateQuery = `
	UPDATE 
		todo
	SET 
	    title = $1, created_at = $2
	WHERE
	    id = $3
	RETURNING
		id, created_at, title, status;	
	`

	deleteQuery = `
	UPDATE 
		todo
	SET
		status = $1
	WHERE
		id = $2;
	`
	restoreQuery = `
	UPDATE 
		todo
	SET 
	    status = $1
	WHERE
	    id = $2
	RETURNING
		id, created_at, title, status;	
	`
)

type todo struct {
	Id        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Title     string    `db:"title"`
	Status    string    `db:"status"`
}

func (s PgStore) Create(ctx context.Context, td td.Todo) (string, error) {

	createdTodo := todo{}
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	err = tx.QueryRowxContext(ctx, createQuery, td.CreatedAt, td.Title, td.Status).StructScan(&createdTodo)

	if err != nil {
		return "", err
	}
	if err = tx.Commit(); err != nil {
		return "", err
	}
	return createdTodo.Id, nil
}

func (s PgStore) List(ctx context.Context) (td.Todo, error) {

	listTodo := todo{}
	err := s.db.QueryRowxContext(ctx, listQuery).StructScan(&listTodo)
	if err != nil {
		return td.Todo{}, err
	}
	todo := listTodo.Transform()
	return todo, nil
}

func (s PgStore) GetById(ctx context.Context, id string) (td.Todo, error) {

	getTodo := todo{}
	err := s.db.QueryRowxContext(ctx, getByIdQuery, id).StructScan(&getTodo)
	if err != nil {
		return td.Todo{}, err
	}
	todo := getTodo.Transform()
	return todo, nil
}

func (s PgStore) Update(ctx context.Context, id string, t td.Todo) (td.Todo, error) {

	updatedTodo := todo{}
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return td.Todo{}, err
	}

	defer tx.Rollback()

	err = tx.QueryRowxContext(ctx, updateQuery, t.Title, t.CreatedAt, id).StructScan(&updatedTodo)

	if err != nil {
		fmt.Println("err", err)
		return td.Todo{}, err
	}
	if err = tx.Commit(); err != nil {
		fmt.Println("err", err)
		return td.Todo{}, err
	}

	todo := updatedTodo.Transform()
	return todo, nil
}

func (s PgStore) Delete(ctx context.Context, id string) error {

	_, err := s.db.ExecContext(ctx, deleteQuery, "deleted", id)
	if err != nil {
		return err
	}
	return nil
}

func (s PgStore) Restore(ctx context.Context, id string) (td.Todo, error) {
	restoredTodo := todo{}
	err := s.db.QueryRowxContext(ctx, restoreQuery, "completed", id).StructScan(&restoredTodo)
	if err != nil {
		return td.Todo{}, err
	}
	todo := restoredTodo.Transform()
	return todo, nil
}

func (t *todo) Transform() td.Todo {
	todo := td.Todo{}
	todo.Id = t.Id
	todo.CreatedAt = t.CreatedAt
	todo.Status = t.Status
	todo.Title = t.Title
	return todo
}

package db

import (
	"app/api"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	Connection *sql.DB
}

func (sl *SQLite) AddTodo(todo api.Todo) error {
	sqlStmt := fmt.Sprintf(`INSERT INTO todo 
		(id, content)
		VALUES
		(?, json(?))`)

	tx, err := sl.Connection.Begin()
	if err != nil {
		log.Panic(err)
	}
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	content, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid.New(), content)
	if err != nil {
		log.Panic(err)
	}
	tx.Commit()
	return nil
}

func (sl *SQLite) AllTodos() ([]api.Todo, error) {
	sqlStmt := fmt.Sprintf(`SELECT content FROM todo`)
	rows, err := sl.Connection.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	todos := []api.Todo{}
	for rows.Next() {
		var content string
		err = rows.Scan(&content)
		if err != nil {
			return todos, err
		}

		var todo api.Todo
		err := json.Unmarshal([]byte(content), &todo)
		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}
	return todos, nil
}

func Init() Database {
	// TODO: make db-name a config param
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		log.Panic(err)
	}
	return &SQLite{Connection: db}
}

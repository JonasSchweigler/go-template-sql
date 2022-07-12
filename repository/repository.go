package repository

import (
	"database/sql"
	"go-template/db"
)

type TestRepository interface {
	//Add Handlers in here
}

type testRepository struct {
	c *sql.DB
}

func NewTestRepository(conn db.Connection) TestRepository {
	return &testRepository{c: conn.DB1()}
}

//SQL Handlers for Repo

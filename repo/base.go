package repo

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"math/rand"
	"os"
	"time"
)

var (
	DB     *pg.DB
	random *rand.Rand
)

func ConnectDatabase() *pg.DB {
	DB = pg.Connect(&pg.Options{
		// Addr:     "localhost:5432",
		// User:     "postgres",
		// Password: "123",
		// Database: "postgres",
		Addr:     "localhost:5432",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})

	DB.AddQueryHook(dbLogger{}) //Log query to console

	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

	return DB
}

type dbLogger struct{}

// hook into query and print out sql query
func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// hook ran after query has executed
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	fmt.Println(string(bytes))
	return nil
}

package repo

import (
	"car-rentals-backend/config"
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-pg/pg/v10"
)

var (
	DB     *pg.DB
	random *rand.Rand
)

func ConnectDatabase() *pg.DB {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config file: ", err)
	}
	opt, err := pg.ParseURL(config.DB_SOURCE)
	if err != nil {
		// panic(err)
		log.Fatal("ERR connect DB: ", err)
	}

	DB := pg.Connect(opt)

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

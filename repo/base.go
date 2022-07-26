package repo

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

var (
	DB     *pg.DB
	random *rand.Rand
)

func ConnectDatabase() *pg.DB {
	DBHOST, _ := viper.Get("DB.HOST").(string)
	DBPORT, _ := viper.Get("DB.PORT").(string)
	USERNAME, _ := viper.Get("DB.USERNAME").(string)
	PASSWORD, _ := viper.Get("DB.PASSWORD").(string)
	DBNAME, _ := viper.Get("DB.NAME").(string)
	DB = pg.Connect(&pg.Options{
		Addr:     DBHOST + ":" + DBPORT,
		User:     USERNAME,
		Password: PASSWORD,
		Database: DBNAME,
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

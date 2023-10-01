package app

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DatabaseConn struct {
	Addr     string
	User     string
	Password string
	Database string
}

func (d *DatabaseConn) loadDatabase(ctx context.Context) (*pg.DB, error) {
	con := pg.Connect(&pg.Options{
		Addr:     d.Addr,
		User:     d.User,
		Password: d.Password,
		Database: d.Database,
	})

	err := con.Ping(ctx)
	if err != nil {
		return nil, err
	}

	_, err = con.Exec("SET search_path TO benchmarks")

	if err != nil {
		return nil, err
	}

	return con, nil
}

type QueryLogger struct{}

func (q QueryLogger) BeforeQuery(c context.Context, evt *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (q QueryLogger) AfterQuery(c context.Context, evt *pg.QueryEvent) error {
	query, err := evt.FormattedQuery()
	if err != nil {
		return err
	}
	fmt.Println(string(query))
	return nil
}

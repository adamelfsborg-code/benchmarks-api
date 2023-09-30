package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Database struct {
	Client *pg.DB
}

type DatabaseConn struct {
	Addr     string
	User     string
	Password string
	Database string
}

func (d *DatabaseConn) Build(ctx context.Context) (*Database, error) {
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

	defer func() {
		err := con.Close()
		if err != nil {
			fmt.Println("Failed to close Database", err)
		}
	}()

	_, err = con.Exec("SET search_path TO testapi")

	if err != nil {
		return nil, err
	}

	database := &Database{
		Client: con,
	}

	return database, nil
}

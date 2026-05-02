package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var readPool *pgxpool.Pool
var writePool *pgxpool.Pool

func InitDB(
	readURL string,
	writeURL string,
) error {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	var err error

	readPool, err = pgxpool.New(
		ctx,
		readURL,
	)

	if err != nil {
		return err
	}

	writePool, err = pgxpool.New(
		ctx,
		writeURL,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadConnection() *pgxpool.Pool {
	return readPool
}

func WriteConnection() *pgxpool.Pool {
	return writePool
}

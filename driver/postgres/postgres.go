package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresDriver struct {
	conn *pgx.Conn
}

func (pgd *PostgresDriver) Open(ctx context.Context, connStr string) (*PostgresDriver, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	pgd.conn = conn
	return pgd, nil
}

func (pgd *PostgresDriver) MapTypes() (t string, err error) {
	switch {

	}
	return
}

package postgres

import (
	"chat-server/internal/client/database"
	"chat-server/internal/client/database/prettier"
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type key string

const (
	TxKey key = "tx"
)

type pg struct {
	dbc *pgxpool.Pool
}

func NewDatabase(dbc *pgxpool.Pool) database.DB {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOneContext(ctx context.Context, dst interface{}, q database.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dst, row)
}

func (p *pg) ScanAllContext(ctx context.Context, dst interface{}, q database.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dst, rows)
}

func (p *pg) ExecContext(ctx context.Context, q database.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRow, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRow, args...)
}

func (p *pg) QueryContext(ctx context.Context, q database.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRow, args...)
	}

	return p.dbc.Query(ctx, q.QueryRow, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q database.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRow, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRow, args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func logQuery(ctx context.Context, q database.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRow, prettier.PlaceholderDollar, args...)
	log.Println(ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}

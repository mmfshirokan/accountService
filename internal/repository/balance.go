package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mmfshirokan/accountService/internal/model"
)

type postgres struct {
	conn pgxpool.Conn
}

type Interface interface {
	Create(id uuid.UUID) error
	PayIn(ctx context.Context, payIn model.Balance) error
	PayOut(ctx context.Context, payOut model.Balance) error
	Get(ctx context.Context, id uuid.UUID) (model.Balance, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func New(conn pgxpool.Conn) Interface {
	return &postgres{
		conn: conn,
	}
}

func (p *postgres) Create(id uuid.UUID) error {
	_, err := p.conn.Exec(context.Background(), "INSERT INTO balance (id, balance) VALUES ($1, 0)", id)
	return err
}

func (p *postgres) PayIn(ctx context.Context, payIn model.Balance) error {
	_, err := p.conn.Exec(ctx, "UPDATE balance SET balance = balance + $1 WHERE id = $2", payIn.Amount, payIn.ID)
	return err
}

func (p *postgres) PayOut(ctx context.Context, payOut model.Balance) error {
	_, err := p.conn.Exec(ctx, "UPDATE balance SET balance = balance - $1 WHERE id = $2", payOut.Amount, payOut.ID)
	return err
}

func (p *postgres) Get(ctx context.Context, id uuid.UUID) (balance model.Balance, err error) {
	err = p.conn.QueryRow(ctx, "SELECT id, balance FROM balance WHERE id = $1", id).Scan(&balance.ID, &balance.Amount)
	return balance, err
}

func (p *postgres) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := p.conn.Exec(ctx, "DELETE FROM balance WHERE id = $1", id)
	return err
}

package flspostgres

import (
	"database/sql"
	"time"
)

// Client は MySQL クライアントを表す。
type Client struct {
	db *sql.DB
}

// NewClient は新しい Client を生成する。
func NewClient(
	dataSourceName string,
	maxOpenConns int,
	maxIdleConns int,
	connMaxLifetime time.Duration,
) (*Client, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	return &Client{
		db: db,
	}, nil
}

// Close はデータベースを閉じる。
func (c *Client) Close() error {
	return c.db.Close()
}

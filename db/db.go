package db

import (
	"database/sql"
)

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR(32) NOT NULL DEFAULT '',
	comment TEXT,
	repeat VARCHAR(128) DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`

type SQLSchedulerStore struct {
	db *sql.DB
}

func Init(dbFile string) (*SQLSchedulerStore, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, err
	}

	return &SQLSchedulerStore{db: db}, nil
}

func (s *SQLSchedulerStore) Close() error {
	return s.db.Close()
}

package driver

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDriver struct {
	db *sql.DB
}

func NewPostgresDriver(url string) (*PostgresDriver, error) {
	db, err := sql.Open("postgres", url)
	if err!=nil{
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err!=nil{
		return nil, err
	}
	
	return &PostgresDriver{
		db: db,
	}, nil
}

func (p *PostgresDriver) Init() error{
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version BIGINT PRIMARY KEY,
			name TEXT NOT NULL,
			checksum TEXT NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresDriver) Close() {
	p.db.Close()
}

// To-do 
// Implement GetAppliedMigrations and Apply

func (p *PostgresDriver) GetAppliedMigrations() ([]MigrationRecord, error) {
	var migrationRecord []MigrationRecord

	return migrationRecord, nil
}

func (p *PostgresDriver) Apply(version int64, name, checksum, sql string) error{
	return nil
}

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(databaseURL string) (*sql.DB, error) {
	log.Printf("Connecting to database...")
	
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	log.Println("Database connection established")
	return db, nil
}

func Migrate(db *sql.DB) error {
	log.Println("Running database migrations...")
	
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
		
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			email VARCHAR(255) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			avatar_url TEXT,
			email_verified BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		
		`CREATE TABLE IF NOT EXISTS watchlists (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			is_default BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		
		`CREATE TABLE IF NOT EXISTS watchlist_assets (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			watchlist_id UUID REFERENCES watchlists(id) ON DELETE CASCADE,
			symbol VARCHAR(50) NOT NULL,
			asset_type VARCHAR(20) NOT NULL,
			added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(watchlist_id, symbol)
		);`,
		
		`CREATE TABLE IF NOT EXISTS portfolios (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		
		`CREATE TABLE IF NOT EXISTS holdings (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			portfolio_id UUID REFERENCES portfolios(id) ON DELETE CASCADE,
			symbol VARCHAR(50) NOT NULL,
			asset_type VARCHAR(20) NOT NULL,
			quantity DECIMAL(20, 8) NOT NULL,
			avg_buy_price DECIMAL(20, 2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		
		`CREATE TABLE IF NOT EXISTS alerts (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			symbol VARCHAR(50) NOT NULL,
			asset_type VARCHAR(20) NOT NULL,
			alert_type VARCHAR(20) NOT NULL,
			target_price DECIMAL(20, 2),
			percentage_change DECIMAL(10, 2),
			is_active BOOLEAN DEFAULT true,
			triggered_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for i, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("migration %d failed: %w", i+1, err)
		}
	}
	
	log.Println("All migrations completed successfully")
	return nil
}

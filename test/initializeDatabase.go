package test

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Initialize() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "D:/dev/finance-app-backend/test/test.db")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS Users (
			id UUID PRIMARY KEY,
			nome VARCHAR(150) NOT NULL,
			email VARCHAR(200) NOT NULL,
			photoUrl TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS Income (
			id UUID PRIMARY KEY,
			userid UUID REFERENCES users(id) ON DELETE CASCADE,
			value NUMERIC(20,2),
			createdAt DATE
		);

		CREATE TABLE IF NOT EXISTS ExtraIncome (
			id UUID PRIMARY KEY,
			userid UUID REFERENCES users(id) ON DELETE CASCADE,
			category VARCHAR(100),
			value NUMERIC(20,2),
			date DATE
		);

		CREATE TABLE IF NOT EXISTS Outcome (
			id UUID PRIMARY KEY,
			userid UUID  REFERENCES users(id) ON DELETE CASCADE,
			type VARCHAR(100),
			category VARCHAR(100),
			value NUMERIC(20,2),
			paymentMethod VARCHAR(100),
			notification BOOLEAN,
			expireDate INT,
			createdAt DATE
		);

		CREATE TABLE IF NOT EXISTS Investment (
			id UUID PRIMARY KEY,
			userid UUID  REFERENCES users(id) ON DELETE CASCADE,
			category VARCHAR(100),
			stockCode VARCHAR(80),
			totalQuantity INT,
			buyPrice NUMERIC(20,2),
			sellPrice NUMERIC(20,2),
			percentage NUMERIC(5,2),
			value NUMERIC(20,2),
			profit NUMERIC(20,2),
			buyDate DATE,
			sellDate DATE,
			createdAt DATE,
			lastSupplyDate DATE
		);

		CREATE TABLE IF NOT EXISTS BuyOperation (
			id UUID PRIMARY KEY,
			investimentId UUID  REFERENCES investment(id) ON DELETE CASCADE,
			category VARCHAR(100),
			stockCode VARCHAR(80),
			quantity INT,
			buyPrice NUMERIC(20,2),
			value NUMERIC(20,2),
			createdAt DATE,
			buyDate DATE
		);

		CREATE TABLE IF NOT EXISTS SupplyOperation (
			id UUID PRIMARY KEY,
			investimentId UUID  REFERENCES investment(id) ON DELETE CASCADE,
			category VARCHAR(100),
			stockCode VARCHAR(80),
			quantity INT,
			supplyPrice NUMERIC(20,2),
			value NUMERIC(20,2),
			createdAt DATE,
			supplyDate DATE
		);

		CREATE TABLE IF NOT EXISTS SellOperation (
			id UUID PRIMARY KEY,
			investimentId UUID  REFERENCES investment(id) ON DELETE CASCADE,
			category VARCHAR(100),
			stockCode VARCHAR(80),
			quantity INT,
			sellPrice NUMERIC(20,2),
			value NUMERIC(20,2),
			createdAt DATE,
			sellDate DATE
		);

		CREATE TABLE IF NOT EXISTS Goals (
			id UUID PRIMARY KEY,
			userid UUID  REFERENCES users(id) ON DELETE CASCADE,
			category VARCHAR(100),
			percentage NUMERIC(5,2),
			createdAt DATE
		);
	`

	_, err = db.Exec(sqlQuery)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return db, nil
}

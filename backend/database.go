package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type InfoDatabase struct {
	db *sql.DB
}

func NewInfoDatabase() (*InfoDatabase, error) {
	connStr := "user=username dbname=farmdb sslmode=disable password=password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &InfoDatabase{db: db}, nil
}

func (id *InfoDatabase) GetClimateInfo(region string) (string, error) {
	var info string
	err := id.db.QueryRow("SELECT info FROM climate_data WHERE region = $1", region).Scan(&info)
	if err != nil {
		return "", err
	}
	return info, nil
}

func (id *InfoDatabase) GetPestInfo(pestName string) (string, error) {
	var info string
	err := id.db.QueryRow("SELECT info FROM pest_data WHERE name = $1", pestName).Scan(&info)
	if err != nil {
		return "", err
	}
	return info, nil
}

func (id *InfoDatabase) GetFarmingTechniques(crop string) (string, error) {
	var techniques string
	err := id.db.QueryRow("SELECT techniques FROM farming_techniques WHERE crop = $1", crop).Scan(&techniques)
	if err != nil {
		return "", err
	}
	return techniques, nil
}

func (id *InfoDatabase) GetMarketTrends(crop string) (string, error) {
	var trends string
	err := id.db.QueryRow("SELECT trends FROM market_trends WHERE crop = $1", crop).Scan(&trends)
	if err != nil {
		return "", err
	}
	return trends, nil
}

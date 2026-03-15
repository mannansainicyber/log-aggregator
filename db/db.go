package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Log struct {
	ID        int
	Level     string
	Service   string
	Message   string
	Timestamp string
}

type StatRow struct {
	Name  string
	Count int
}

type Stats struct {
	ByLevel   []StatRow
	ByService []StatRow
}

func getDBPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".log-ag", "logs.db")
}

func Connect() (*sql.DB, error) {
	path := getDBPath()
	os.MkdirAll(filepath.Dir(path), 0755)

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS logs (
		id        INTEGER PRIMARY KEY AUTOINCREMENT,
		level     TEXT,
		service   TEXT,
		message   TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	return db, err
}

func Insert(level, service, message string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO logs (level, service, message) VALUES (?, ?, ?)",
		level, service, message,
	)
	return err
}

func Query(level, service, since string, limit int) ([]Log, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, level, service, message, timestamp FROM logs WHERE 1=1"
	args := []any{}

	if level != "" {
		query += " AND level = ?"
		args = append(args, level)
	}
	if service != "" {
		query += " AND service = ?"
		args = append(args, service)
	}
	if since != "" {
		query += " AND timestamp >= datetime('now', ?)"
		args = append(args, "-"+since)
	}

	query += " ORDER BY timestamp DESC LIMIT ?"
	args = append(args, limit)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []Log
	for rows.Next() {
		var l Log
		rows.Scan(&l.ID, &l.Level, &l.Service, &l.Message, &l.Timestamp)
		logs = append(logs, l)
	}
	return logs, nil
}

func QueryAfter(lastID int) ([]Log, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT id, level, service, message, timestamp FROM logs WHERE id > ? ORDER BY id ASC",
		lastID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []Log
	for rows.Next() {
		var l Log
		rows.Scan(&l.ID, &l.Level, &l.Service, &l.Message, &l.Timestamp)
		logs = append(logs, l)
	}
	return logs, nil
}

func GetStats() (Stats, error) {
	db, err := Connect()
	if err != nil {
		return Stats{}, err
	}
	defer db.Close()

	var stats Stats

	// count by level
	rows, err := db.Query("SELECT level, COUNT(*) FROM logs GROUP BY level ORDER BY COUNT(*) DESC")
	if err != nil {
		return Stats{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var s StatRow
		rows.Scan(&s.Name, &s.Count)
		stats.ByLevel = append(stats.ByLevel, s)
	}

	// count by service
	rows2, err := db.Query("SELECT service, COUNT(*) FROM logs GROUP BY service ORDER BY COUNT(*) DESC")
	if err != nil {
		return Stats{}, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var s StatRow
		rows2.Scan(&s.Name, &s.Count)
		stats.ByService = append(stats.ByService, s)
	}

	return stats, nil
}

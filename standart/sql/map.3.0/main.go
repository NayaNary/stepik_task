package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

// начало решения

// SQLMap представляет карту, которая хранится в SQL-базе данных
type SQLMap struct {
	db         *sql.DB
	stmtGet    *sql.Stmt
	stmtSet    *sql.Stmt
	stmtDelete *sql.Stmt
	timeOut    time.Duration
}

// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	query := `
	    drop table if exists map;
		create table if not exists map(key text primary key, val blob);
	`
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return &SQLMap{}, err
	}
	stmtGet, err := db.PrepareContext(ctx, `select val from map where key = ?`)
	if err != nil {
		return &SQLMap{}, err
	}
	stmtSet, err := db.PrepareContext(ctx, `insert into map(key, val) values (?, ?)
on conflict (key) do update set val = excluded.val`)
	if err != nil {
		return &SQLMap{}, err
	}
	stmtDelete, err := db.PrepareContext(ctx, `delete from map where key = ?`)
	if err != nil {
		return &SQLMap{}, err
	}

	return &SQLMap{
		db:         db,
		stmtGet:    stmtGet,
		stmtSet:    stmtSet,
		stmtDelete: stmtDelete,
		timeOut:    60 * time.Second,
	}, nil
}

// SetTimeout устанавливает максимальное время выполнения
// отдельного метода карты.
func (m *SQLMap) SetTimeout(d time.Duration) {
	m.timeOut = d
}

// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeOut)
	defer cancel()
	row := m.stmtGet.QueryRowContext(ctx, key)
	var val any
	err := row.Scan(&val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// Set устанавливает значение для указанного ключа.
// Если такой ключ уже есть - затирает старое значение (это не считается ошибкой).
func (m *SQLMap) Set(key string, val any) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeOut)
	defer cancel()
	_, err := m.stmtSet.ExecContext(ctx, key, val)
	if err != nil {
		return err
	}
	return nil
}

// SetItems устанавливает значения указанных ключей.
func (m *SQLMap) SetItems(items map[string]any) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeOut)
	defer cancel()

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txStmt := tx.StmtContext(ctx, m.stmtSet)
	for key, val := range items {
		_, err := txStmt.ExecContext(ctx, key, val)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// Delete удаляет запись карты с указанным ключом.
// Если такого ключа нет - ничего не делает (это не считается ошибкой).
func (m *SQLMap) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeOut)
	defer cancel()
	_, err := m.stmtDelete.ExecContext(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

// Close освобождает ресурсы, занятые картой в базе.
func (m *SQLMap) Close() error {
	if err := m.stmtDelete.Close(); err != nil {
		return err
	}
	if err := m.stmtGet.Close(); err != nil {
		return err
	}
	if err := m.stmtSet.Close(); err != nil {
		return err
	}
	return nil
}

// конец решения

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m, err := NewSQLMap(db)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.SetTimeout(10 * time.Millisecond)

	m.Set("name", "Alice")
	m.Get("name")
}

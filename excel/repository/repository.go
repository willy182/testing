package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"willy182/testing/excel/model"
)

// TxFunc tx closure func type
type TxFunc func(*sql.Tx) error

// Repository parent domain for all
type Repository struct {
	db    *sql.DB
	tx    *sql.Tx
	mutex sync.Mutex
}

// NewRepository create new Repository domain
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// SetTx starting transaction session
func (repo *Repository) SetTx(tx *sql.Tx) {
	repo.tx = tx
}

// StartTransaction starting transaction session
func (repo *Repository) StartTransaction() {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}
	repo.mutex.Lock()
	repo.tx = tx
}

// Rollback transaction
func (repo *Repository) Rollback() {
	defer repo.mutex.Unlock()
	if repo.tx != nil {
		repo.tx.Rollback()
	}
	repo.tx = nil
}

// Commit transaction
func (repo *Repository) Commit() {
	defer repo.mutex.Unlock()
	if repo.tx != nil {
		repo.tx.Commit()
	}
	repo.tx = nil
}

// RunTransaction clean closure transaction (or multiple transaction) func with tx parameter in closure func (using in internal packages)
func (repo *Repository) RunTransaction(txFunc TxFunc) (err error) {
	// repo.mutex.Lock()
	// defer repo.mutex.Unlock()

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		if err != nil {
			fmt.Println(err, "run_transaction", "")
		}
	}()

	err = txFunc(tx)
	return
}

// ResultRepository data structure
type ResultRepository struct {
	Result interface{}
	Error  error
}

// MigrasiRepository abstraction interface
type MigrasiRepository interface {
	Save(ctx context.Context, payload *model.PayloadModel) <-chan ResultRepository
}

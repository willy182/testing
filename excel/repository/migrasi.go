package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"willy182/testing/excel/model"
)

// migrasi data structure
type migrasi struct {
	*Repository
}

// NewMigrasiRepository function for creating migrasi repository postgres
func NewMigrasiRepository(repo *Repository) MigrasiRepository {
	return &migrasi{repo}
}

// Save function, for save migrasi object into database
func (r *migrasi) Save(c context.Context, payload *model.PayloadModel) <-chan ResultRepository {
	opName := "MigrasiRepository-Save"
	output := make(chan ResultRepository)

	go func() {
		defer close(output)

		tx, err := r.db.Begin()
		if err != nil {
			fmt.Println(err.Error(), opName, "begin")
			output <- ResultRepository{Error: err}
			return
		}

		var stmt *sql.Stmt

		query := `INSERT INTO migrasi (level_harga_id, "data", provinsi, kota, petugas, no_hp, pasar, tanggal) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

		stmt, err = tx.Prepare(query)

		if err != nil {
			tx.Rollback()
			fmt.Println(err.Error(), opName, "prepare")
			output <- ResultRepository{Error: err}
			return
		}
		defer stmt.Close()

		payloadJSON, _ := json.Marshal(payload.Data)

		_, err = stmt.Exec(
			payload.LevelHargaID, payloadJSON, payload.Provinsi, payload.Kota, payload.Petugas, payload.NoHP, payload.Pasar, payload.Tanggal,
		)

		if err != nil {
			tx.Rollback()
			fmt.Println(err.Error(), opName, "exec")
			output <- ResultRepository{Error: err}
			return
		}

		tx.Commit()

		output <- ResultRepository{Result: payload}

	}()

	return output
}

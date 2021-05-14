package model

type PayloadModel struct {
	LevelHargaID int
	Provinsi     string
	Kota         string
	Pasar        string
	Petugas      string
	NoHP         string
	Tanggal      string
	Data         interface{}
}

type DataModel struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga *int   `json:"harga"`
}

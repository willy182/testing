package main

import (
	"context"
	"fmt"
	"strconv"
	"willy182/testing/excel/config"
	"willy182/testing/excel/model"
	"willy182/testing/excel/repository"

	excelize "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	ctx := context.Background()
	conf := config.Load()
	db := repository.NewRepository(conf.DB)
	repo := repository.NewMigrasiRepository(db)

	// data produsen
	file := "data_panel_produsen_jan_maret.xlsx"
	jmlRow := 28307
	levelHargaID := 1

	// data grosir
	// file := "data_panel_grosir_jan_maret.xlsx"
	// jmlRow := 2523
	// levelHargaID := 2

	// data eceran jan maret
	// file := "data_panel_eceran_jan_maret.xlsx"
	// jmlRow := 2533
	// levelHargaID := 3

	// data eceran 3 jan 15 maret
	// file := "data_panel_eceran_3_jan_15_maret.xlsx"
	// jmlRow := 23482
	// levelHargaID := 3

	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	sheet1Name := "Sheet1"

	var loop int
	for index := 2; index <= jmlRow; index++ {
		provinsi, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", index))
		kota, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", index))
		petugas, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", index))
		noHp, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", index))
		// pasar, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", index))
		var pasar string

		/* start produsen */
		tanggal, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("U%d", index))

		columnE := fmt.Sprintf("E%d", index)
		columnF := fmt.Sprintf("F%d", index)
		columnG := fmt.Sprintf("G%d", index)
		columnH := fmt.Sprintf("H%d", index)
		columnI := fmt.Sprintf("I%d", index)
		columnJ := fmt.Sprintf("J%d", index)
		columnK := fmt.Sprintf("K%d", index)
		columnL := fmt.Sprintf("L%d", index)
		columnM := fmt.Sprintf("M%d", index)
		columnN := fmt.Sprintf("N%d", index)
		columnO := fmt.Sprintf("O%d", index)
		columnP := fmt.Sprintf("P%d", index)
		columnQ := fmt.Sprintf("Q%d", index)
		columnR := fmt.Sprintf("R%d", index)
		columnS := fmt.Sprintf("S%d", index)
		columnT := fmt.Sprintf("T%d", index)
		/* end produsen */

		/* start grosir */
		// tanggal, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", index))

		// columnF := fmt.Sprintf("F%d", index)
		// columnG := fmt.Sprintf("G%d", index)
		// columnH := fmt.Sprintf("H%d", index)
		// columnI := fmt.Sprintf("I%d", index)
		// columnJ := fmt.Sprintf("J%d", index)
		// columnK := fmt.Sprintf("K%d", index)
		// columnL := fmt.Sprintf("L%d", index)
		// columnM := fmt.Sprintf("M%d", index)
		// columnN := fmt.Sprintf("N%d", index)
		// columnO := fmt.Sprintf("O%d", index)
		// columnP := fmt.Sprintf("P%d", index)
		// columnQ := fmt.Sprintf("Q%d", index)
		/* end grosir */

		/* start eceran */
		// tanggal, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", index))

		// columnF := fmt.Sprintf("F%d", index)
		// columnG := fmt.Sprintf("G%d", index)
		// columnH := fmt.Sprintf("H%d", index)
		// columnI := fmt.Sprintf("I%d", index)
		// columnJ := fmt.Sprintf("J%d", index)
		// columnK := fmt.Sprintf("K%d", index)
		// columnL := fmt.Sprintf("L%d", index)
		// columnM := fmt.Sprintf("M%d", index)
		// columnN := fmt.Sprintf("N%d", index)
		// columnO := fmt.Sprintf("O%d", index)
		// columnP := fmt.Sprintf("P%d", index)
		// columnQ := fmt.Sprintf("Q%d", index)
		// columnR := fmt.Sprintf("R%d", index)
		/* end eceran */

		/* start produsen */
		var komoditas = []string{
			"Prosentase Luas Panen Padi", "GKP Tingkat Petani", "GKP Tingkat Penggilingan", "GKG Tingkat Penggilingan",
			"Beras Medium Tingkat Penggilingan", "Beras Premium Tingkat Penggilingan", "Jagung Pipilan Kering Tingkat Petani",
			"Kedelai Biji Kering Tingkat Petani", "Cabai Merah Keriting Tingkat Petani", "Cabai Merah Keriting Tingkat Petani",
			"Bawang Merah Tingkat Petani", "Sapi Hidup (Tingkat Peternak/RPH)", "Stok GKG Tingkat Penggilingan",
			"Stok Beras Tingkat Penggilingan", "Ayam Ras Pedaging", "Telur Ayam Ras"}
		var columnKomoditas = []string{columnE, columnF, columnG, columnH,
			columnI, columnJ, columnK, columnL,
			columnM, columnN, columnO, columnP,
			columnQ, columnR, columnS, columnT}
		/* end produsen */

		/* start grosir */
		// var komoditas = []string{
		// 	"Beras Premium", "Beras Medium", "Kedelai Biji Kering",
		// 	"Bawang Merah", "Bawang Putih (Bonggol)", "Cabai Merah Keriting",
		// 	"Cabai Rawit Merah", "Daging Sapi Murni Tingkat Pemotong/RPH", "Daging Ayam Ras",
		// 	"Telur Ayam Ras", "Gula Pasir Lokal/Curah", "Minyak Goreng Kemasan Sederhana"}
		// var columnKomoditas = []string{columnF, columnG, columnH,
		// 	columnI, columnJ, columnK,
		// 	columnL, columnM, columnN,
		// 	columnO, columnP, columnQ}
		/* end grosir */

		/* start eceran */
		// var komoditas = []string{
		// 	"Beras Premium", "Beras Medium", "Kedelai Biji Kering (Impor)",
		// 	"Bawang Merah", "Bawang Putih Bonggol", "Cabai Merah Keriting",
		// 	"Cabai Rawit Merah", "Daging Sapi Murni", "Daging Ayam Ras",
		// 	"Telur Ayam Ras", "Gula Pasir Curah/Lokal", "Minyak Goreng Kemasan Sederhana", "Tepung Terigu (Curah)"}
		// var columnKomoditas = []string{columnF, columnG, columnH,
		// 	columnI, columnJ, columnK,
		// 	columnL, columnM, columnN,
		// 	columnO, columnP, columnQ, columnR}
		/* end eceran */

		var datas []model.DataModel
		for x, val := range columnKomoditas {
			var data model.DataModel
			data.ID = x + 1
			data.Nama = komoditas[x]

			cellValue, _ := xlsx.GetCellValue(sheet1Name, val)
			if cellValue != "" && cellValue != "0" {
				harga, _ := strconv.Atoi(cellValue)
				data.Harga = &harga
			}

			datas = append(datas, data)
		}

		if noHp[:1] != "0" {
			noHp = fmt.Sprintf("0%s", noHp)
		}

		row := model.PayloadModel{
			LevelHargaID: levelHargaID,
			Provinsi:     provinsi,
			Kota:         kota,
			Pasar:        pasar,
			Petugas:      petugas,
			NoHP:         noHp,
			Tanggal:      tanggal,
			Data:         datas,
		}

		result := <-repo.Save(ctx, &row)
		if result.Error != nil {
			fmt.Println(petugas, noHp, len(noHp), tanggal)
		} else {
			// dataResult := result.Result.(*model.PayloadModel)
			// fmt.Println("success result", dataResult.Petugas)
			loop += 1
		}
	}

	fmt.Println("success", loop)
}

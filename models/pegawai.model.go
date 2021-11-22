package models

import (
	"net/http"

	"github.com/basic-echo-golang/db"
	"github.com/go-playground/validator/v10"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrObj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		// Sesuai urutan nama kolomnya
		err := rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func StorePegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	v := validator.New()

	peg := Pegawai{
		Nama:    nama,
		Alamat:  alamat,
		Telepon: telepon,
	}

	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT pegawai (nama,alamat,telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pegawai where id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

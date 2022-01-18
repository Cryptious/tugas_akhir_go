package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type data_karyawan struct {
	Id     string
	Nama   string
	Umur   int
	Posisi string
}

// tampil Karyawan
func tampil_karyawan() ([]data_karyawan, error) {
	var err error
	var client = &http.Client{}
	var data []data_karyawan

	request, err := http.NewRequest("POST", baseURL+"/karyawan", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//cari karyawan
func cari_karyawan(idkaryawan string) (data_karyawan, error) {
	var err error
	var client = &http.Client{}
	var data data_karyawan

	var param = url.Values{}
	param.Set("Id", idkaryawan)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/cari_karyawan", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)

	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//hapus karyawan
func hapus_karyawan(idkaryawan string) (data_karyawan, error) {
	var err error
	var client = &http.Client{}
	var data data_karyawan

	var param = url.Values{}
	param.Set("Id", idkaryawan)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/hapus_karyawan", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)

	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//Update karyawan
func update_karyawan(idkaryawan string, namakaryawan string) (data_karyawan, error) {
	var err error
	var client = &http.Client{}
	var data data_karyawan

	var param = url.Values{}
	param.Set("Id", idkaryawan)
	param.Set("Nama", namakaryawan)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/update_karyawan", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)

	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func main() {

	// tampil karyawan
	// var karyawan, err = tampil_karyawan() // tampil karyawan
	// if err != nil {
	// 	fmt.Println("Karyawan tidak ada", err.Error())
	// 	return
	// }
	// for _, each := range karyawan {
	// 	fmt.Println("ID: ", each.Id, "Nama: ", each.Nama, "Umur: ", each.Umur, "Posisi: ", each.Posisi) // cari karyawan
	// }

	// tampil cari karyawan
	// var karyawan, err = cari_karyawan("K2") // cari karyawan
	// if err != nil {
	// 	fmt.Println("Karyawan tidak terdaftar", err.Error())
	// 	return
	// }
	// fmt.Println("ID: ", karyawan.Id, "Nama: ", karyawan.Nama, "Umur: ", karyawan.Umur, "Posisi: ", karyawan.Posisi) // cari karyawan

	// Hapus Karyawan
	// hapus_karyawan("K5") // Hapus karyawan
	// fmt.Println("Karyawan Berhasil Dihapus")

	// Update karyawan
	update_karyawan("K5", "Fenny") // Hapus karyawan
	fmt.Println("Karyawan Berhasil Diupdate")

}

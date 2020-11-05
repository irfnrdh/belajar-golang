/*
	Catatan Golang Dasar
	@irfnrdh
*/

// wajib nih
package main

// Import dulu yakk
import "fmt"

/**
# import banyak

	import (
		"bytes"
		"flag"
		"fmt"
		"io/ioutil"
		"os"
		"regexp"

		"github.com/tebeka/snowball"
	)
**/

// permainan dimulai
func main() {

	//buat hello world
	fmt.Println("hello world")

	/**
	-------------------------------------------------------
	# Tipe data

		## Number
		- Integer [ada nilai min-max](
			int8(max 127),
			int16(max 32767),
			int32(max 2147483647), alias rune, int
			int64(max 9223372036854775807))
		- Unsign Integer [hanya nilai max] (
			uint8 (max 255), alias byte
			uint16 (max 65535),
			uint32 (max 4294967295), alias uint
			uint64 (max 18446744073709551615))
		- Ploting Point [comma]
			float32 (min 1.8*10^-38 max 3.4*10^38)
			float64 (min 2.23*10^-308 max 1.80*10^308)
			complex64 (+imaginary parts)
			complex128 (+imaginary parts)

		## Boolean
		-> benar atau salah
		bool (true, false)

		## String
		-> karakter dimulai dari nol - inf
		string ""
		fungsi :
		- len("string")		-> menghitung jumlah karakter
		- "string"[number]  -> mengambil karakter sesuai posisi jadi byte

	**/

	// number
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Delapan Koma Lima = ", 8.5)

	// boolean
	fmt.Println("Benar", true)
	fmt.Println("Salah", false)

	// String
	fmt.Println("Irf")
	fmt.Println("Irfannur")
	fmt.Println("Irfannur Diah")

	fmt.Println(len("Irfannur Diah"))
	fmt.Println("Irfannur Diah"[0])

	/**
	-------------------------------------------------------
	# Variabel

	- Hanya bisa menyimpan tipe data yang sama.
	- membuat kata kunci var tidak begitu wajib diganti dengan
	- membuat :=  lalu data

	Deklarasi :
	// normal variabel
	var <tipedata>
	var String

	// singkat
	nama := "Fuelan"

	// multiple variabel
	var(
		namaDepan = "Irfannur"
		namaBelakang = "Diah"
	)

	// constant
	- nilainya tidak bisa diubah
	- wajib inisialisasi tipe data

	const nama string = "irfan"

	// mutiple constant
	const(
		kota string = "Kualasimpang"
		kodepos int = 24476
	)

	**/

	// inisialisasi variabel
	var nama string

	nama = "Irfannur Diah"
	fmt.Println(nama)

	nama = "Irfan"
	fmt.Println(nama)

	// membuat varibel tanpa diikuti tipe data
	var namaKawan = "budi"
	fmt.Println(namaKawan)

	// Jika ingin menjadi int8
	var umur int8 = 25
	/**
	jika tak dipaggil :
	error 'declared but not used Process exiting'
	**/
	fmt.Println(umur)

	// otomatis akan jadi int/int32
	var umurHidup = 65
	fmt.Println(umurHidup)

	// tanpa menggunakan keyword var untuk deklarasi awal
	username := "irfnrdh"
	fmt.Println(username)

	var (
		namaDepan    = "Irfannur"
		namaBelakang = "Diah"
	)
	fmt.Println(namaDepan + " " + namaBelakang)

	// constanta
	const pi float32 = 3.14
	fmt.Println(pi)

	// multiple constanta
	const (
		kota    string = "Kualasimpang"
		kodepos int16  = 24476
	)
	fmt.Println(kota, kodepos)

	/**
	-------------------------------------------------------
	# Konversi tipe data
	- hati-hati jika variabel tampungan tingkat cukup akan terjadi error overflow.

	var <namavar referensi> <tipedata> = <nilai>
	var <namavar baru> <tipedata konversi> = <tipedata konversi> (namavar referensi)

	// kasus pada string

	**/

	var nilai32 int32 = 24476
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai16)

}

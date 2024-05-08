package main

import "fmt"

/* 21.	Aplikasi Chatting
Deskripsi: Aplikasi digunakan untuk pengiriman pesan. Pengguna aplikasi adalah pengguna yang telah melakukan instalasi aplikasi dan admin aplikasi.
Spesifikasi:
a.	Pengguna bisa melakukan registrasi akun.
b.	Admin bisa melakukan persetujuan/penolakan registrasi akun dan mencetak daftar akun.
c.	Pengguna bisa mengirim pesan pribadi kepada akun lain
d.	Pengguna bisa membuat grup chatting dan menambahkan akun lain untuk masuk ke grup chattingnya.
e.	Pengguna bisa mengirim pesan kepada grup chatting.
f.	Pengguna bisa melihat peserta dalam grup.
*/

const NMAXpengguna int = 1000
const NMAXpesanGrup = 1000

type pengguna struct {
	/* 
	kalau misal isVerified string "pending" berarti menunggu persetujuan admin,
	misal "accepted" udah di verif admin
	kalau "rejected" ditolak sama admin
	*/ 
	username string
	email    string
	noTelp int
	isVerified string
}
type pesan struct {
	pengirim pengguna
	text string
	penerima pengguna
}
type group struct {
	pemilik pengguna
	anggota [10] pengguna
	percakapanGrup  
}

type pesanGrup struct {
	pengirim pengguna
	text string
}

type tab [NMAXpengguna]pengguna

func main() {
	var a tab
	var x, n int
	for {
		fmt.Println("-------------------------")
		fmt.Println("    Welcome to HELLO!    ")
		fmt.Println("Teman chatting setiamu ^^")
		fmt.Println("-------------------------")
		fmt.Println("Apakah kamu Pengguna atau Admin?")
		fmt.Println("1. Pengguna")
		fmt.Println("2. Admin")
		n = 0
		fmt.Scan(&x)
		if x == 1 {
			registrasi(a, n)
		} else {

		}
	}
}

func registrasi(user tab, n *int) {
	user[*n].isVerified = "pending"
	fmt.Println("Isi data berikut untuk melakukan proses registrasi !")
	fmt.Println("Masukkan Username")
	fmt.Scan(&user[*n].username)
	fmt.Println("Masukkan Email")
	fmt.Scan(&user[*n].email)
	fmt.Println("Masukkan Nomor Telepon")
	fmt.Scan(&user[*n].noTelp)
	fmt.Println("Harap tunggu sebentar, menunggu verifikasi...")
	*n++
}

func admin(user tab, n int) {
	var x int
	fmt.Println("Selamat datang !")
	fmt.Println("Menu :")
	fmt.Println("1. Verifikasi Akun")
	fmt.Println("2. Cetak Daftar Akun")
	fmt.Scan(&x)
	if x == 1 {
		verifikasi(user, n)
	}
}

func verifikasi(user tab, n int) {
	
}

func cetakDaftar 
6
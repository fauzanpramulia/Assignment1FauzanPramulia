package function

import (
	"Assignment1/model"
)

func ListTeman() []model.Teman{
	temanKelas := []model.Teman{
		{
			Nama: "Thomas",
			Alamat: "Jl. Contoh No. 123",
			Pekerjaan: "Developer",
			Alasan: "Pandai"},
		{
			Nama:"Maria",
			Alamat: "Jl. Sebelah No. 456",
			Pekerjaan: "Designer",
			Alasan: "Hooo"},
		{
			Nama:"John",
			Alamat:  "Jl. Samping No. 789",
			Pekerjaan: "Engineer",
			Alasan: "Yeah"},
	}

	// Mengecek apakah nama teman ada di data teman-teman kelas
	// teman, found := temanKelas[nama]
	// if found {
	// 	// Menampilkan data teman
	// 	fmt.Printf("Nama     : %s\n", teman.Nama)
	// 	fmt.Printf("Alamat   : %s\n", teman.Alamat)
	// 	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	// 	fmt.Printf("Absen    : %d\n", teman.Absen)
	// } else {
	// 	// Jika tidak ditemukan
	// 	fmt.Println("Teman tidak ditemukan.")
	// }

	return temanKelas
}
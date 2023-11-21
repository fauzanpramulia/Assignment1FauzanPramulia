package main

import (
	"Assignment1/method"
	"Assignment1/model"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go <nama>")
		return
	}

	// Mengambil nama teman dari argumen
	namaTeman := args[1]
	teman := method.CustomTeman{
		Teman: model.Teman{
			Nama: namaTeman,
		},
	}
	// Menampilkan data teman berdasarkan nama
	teman.FindTeman()
}

package method

import (
	"Assignment1/function"
	"Assignment1/model"
	"fmt"
	"strings"
)

type CustomTeman struct{
	model.Teman
}

func (t CustomTeman) FindTeman(){
	var temans = function.ListTeman()

	for key := range temans{
		if strings.ToLower(temans[key].Nama) == strings.ToLower(t.Nama){
			fmt.Println("ID:", key+1)
			fmt.Println("Nama:", t.Nama)
			fmt.Println("Alamat:", t.Alamat)
			fmt.Println("Pekerjaan:", t.Pekerjaan)
			fmt.Println("Alasan:", t.Alasan)
		}
	}
}
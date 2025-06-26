package main

import(
	"fmt"
)

func main(){
	var santri = []map[string]interface{}{
		{"nama": "Althaf", "age": 15, "nilai": 95},
		{"nama": "Hafla", "age": 14, "nilai": []int{90, 95, 100}},
		{"nama": "Rasyad", "age": 16, "nilai": "Belum dinilai"},
	}

	for _, s := range santri {
		fmt.Println("Nama: ", s["nama"])
		fmt.Println("Umur: ", s["age"])

		switch v := s["nilai"].(type){
		case int:
			fmt.Println("Nilai tunggal: ", v)
		case string:
			fmt.Println("Status nilai: ", v)
		case []int:
			fmt.Println("Nilai banyak: ", v)
		default:
			fmt.Println("Tidak ada nilai")
		}

		fmt.Println("======================")
	}
}
package main

import "fmt"

func main() {
	CreateSliceWithIndexArray()

}

func CreateSliceWithIndexArray() {
	var listArray = [5]string{"Dicky1", "Dicky2", "Dicky3", "Dicky4", "Dicky5"}

	//listArray[a:b] --> a harus lebih kecil dari pada b

	var newSlice = listArray[0:2] // mengambil data pada listArray dari index 0 hingga sebelum index 2 

	// 						[:] --> add semua elemen
	// 						[2:] --> semua elemen mulai index dari index 2
	// 						[:2] --> semua elemen hingga sebelum index ke-2

	fmt.Println(newSlice)
}
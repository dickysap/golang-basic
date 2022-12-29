package main

import "fmt"

func main() {
	ArrayWithCapasity()
	ArrayWithoutCapacity()
	NestedArray()
	ArrayWithMake()

}

func ArrayWithCapasity(){
	// names --> sebagai variabel yang menampung array
	// [5] --> sebagai kapasitas array atau slot array 

	var names [5]string
	var arrayWithValue  = [4]string{"dicky", "saputra"}


	// cara menambahkan Array
	names[0] = "Dicky"
	names[1] = "Saputra"
	names[2] = "Dicky"
	names[3] = "Hermawan"
	names[4] = "Dicky"

	fmt.Println(names)
	fmt.Println("Array With Value", arrayWithValue)
}

func ArrayWithoutCapacity(){
	var names = [...]string{"Dicky", "Saputra", "Hermawan"}
	fmt.Println(names)

}

func NestedArray() {
	var names[2][3]string

	names[0][0] = "Dicky"
	names[0][1] = "Saputra"
	names[0][2] = "Saputra"

	names[1][0] = "Novita"
	names[1][1] = "Putri"
	names[1][2] = "Putri"
	
	for i := 0; i < len(names); i++ {

		for j := 0; j < len(names[i]); j++{
			fmt.Println(names[i][j])
		}
		
	}
}

func ArrayWithMake() {
	var buah = make([]string, 2)
	buah[0] = "Stroberi"
	buah[1] = "Anggur"
	
	namaBuah := [...]string{"apel", "Mangga", "Alpukat"}

	for i, _ := range namaBuah {
		// append akan menambahkan sebuah slice kedalam array buah 
		buah = append(buah, namaBuah[i])
	}

	fmt.Println(buah)

}
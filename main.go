package main

import "fmt"

func main() {
	//NGC 1 - Variable 1
	var myNum uint32 = 50
	var myNum2 float32 = 51
	var myNumStr string = "50"

	fmt.Println(myNum)
	fmt.Println(myNum2)
	fmt.Println(myNumStr)

	//NGC 1 - Variable 2
	// var x, y int32 = 5, 10
	// z := x + y
	var x int32 = 5
	var y int32 = 10
	var z int32 = x + y

	fmt.Println(z)

	//NGC 1 - CLI
	var name string
	//input nama
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&name)

	//print value yang dimasukkan
	fmt.Println("Hallo", name)

	//NGC 1 - Array & Slice 1
	// letters := []string{"a", "b", "c", "d"}
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}
	//menghitung panjang slice
	fmt.Println(len(people))
	fmt.Println(people)
	//menambahkan variable ke slice
	people = append(people, "Hank", "Marie")
	fmt.Println(len(people))
	fmt.Println(people)

	//NGC 1 - Array & Slice 2
	person1 := map[string]string{
		"name":   "Hank",
		"gender": "M",
	}
	person2 := map[string]string{
		"name":   "Heisenberg",
		"gender": "M",
	}
	person3 := map[string]string{
		"name":   "Skyler",
		"gender": "F",
	}

	person := [3]map[string]string{}
	person[0] = person1
	person[1] = person2
	person[2] = person3
	fmt.Println("Tanpa Gelar:", person)

	//menanmbahkan Mr. dan Mrs.
	if person[0]["gender"] == "F" {
		person[0]["name"] = "Mrs. " + person[0]["name"]
	} else {
		person[0]["name"] = "Mr. " + person[0]["name"]
	}

	if person[1]["gender"] == "F" {
		person[1]["name"] = "Mrs. " + person[1]["name"]
	} else {
		person[1]["name"] = "Mr. " + person[1]["name"]
	}

	if person[2]["gender"] == "F" {
		person[2]["name"] = "Mrs. " + person[2]["name"]
	} else {
		person[2]["name"] = "Mr. " + person[2]["name"]
	}
	fmt.Println("Dengan Gelar:", person)
}

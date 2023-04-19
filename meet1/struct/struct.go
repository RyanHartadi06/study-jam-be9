package main

import "fmt"

/*
	type NameStruct struct {
		Field1 dataType <- field.Property
		FieldN otherDataType
	}

	func (NameStruct) NameOfMethod()
	func (n *NameStruct) NameOfMethodWithStructAlias() {
		//membutuhkan nilai field1
		var field1 = n.Field1
	}
*/

type Biodata struct {
	Name    string
	Age     int
	Hobbies []string
}

type Biodata4 struct {
	*Biodata
	Address string
}

func (b *Biodata) GetName() string {
	return b.Name
}

func (b *Biodata) GetHobbies() []string {
	return b.Hobbies
}

func (b *Biodata) GetAge() int {
	return b.Age
}

// Periodic
func (b *Biodata) AddHobby(hobby ...string) {
	b.Hobbies = append(b.Hobbies, hobby...)
}

func main() {

	biodata := Biodata{
		Name:    "Ryan Hartadi",
		Age:     20,
		Hobbies: []string{"Musik", "Olahraga"},
	}
	//biodata := Biodata4{
	//	Biodata: &Biodata{
	//		Name:    "Ryan Hartadi",
	//		Age:     20,
	//		Hobbies: []string{"Musik", "Olahraga"},
	//	},
	//	Address: "asd",
	//}

	fmt.Println("Get Name:", biodata.GetName())
	fmt.Println("Get Hobbies:", biodata.GetHobbies())
	fmt.Println("Get Age:", biodata.GetAge())

	biodata.AddHobby("Berenang", "Makan")

	fmt.Println("Get Hobbies:", biodata.GetHobbies())

}

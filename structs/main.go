// Struct: is a data structure, a collection of properties that are related to each other
package main

import "fmt"

// define all the properties a struct should have
// go assigns a zero value to each of the fields if they are not defined on initialization
type person struct {
	firstName string
	lastName  string
	age       int32
	height    float64
	weight    float64
	alive     bool
	handles   socials // embedded struct OR just type the name of the struct i.e socials
}

type socials struct {
	github  string
	twitter string
	email   string
}

func main() {
	// create a value that matches the struct definition
	// declarations...
	me := person{firstName: "Emmanuel", lastName: "Sule", height: 6.25, age: 22, weight: 80, handles: socials{
		email: "tim@mail.com",
	}} // option 1
	me1 := person{"Alicent", "Joshua", 19, 20, 20, true, socials{twitter: "ngbede_", github: "ngbede"}} // option 2
	fmt.Println(me.weight)
	fmt.Println(me.alive)
	fmt.Println("Github: ", me1.handles.github)
	var other person = person{firstName: "Joseph", age: 54} // option 3
	fmt.Println(other.age)

	me.age = 23
	fmt.Println(me)
	fmt.Printf("%+v\n", me) // this prints out the field names in the struct alongside their values

	// pointer stuff
	mePointer := &me // this creates a direct link to the spot in ram where me is stored and saves it in the mePointer
	fmt.Println(&mePointer)
	mePointer.height = 7 // so updates on this now point to the actual me value

	fmt.Println(me.getBmi())
	me.updateName("ngbede")
	fmt.Println(me.firstName)
}

// receiver functions on srtucts
func (p person) getBmi() float64 {
	return p.weight / p.height
}

// a copy is made for this funcion call by GO so regular update won't work
// adding the * symbol to the receiver type, causes go to reference the orignal copy of person
func (p *person) updateName(newName string) { // *person refers to the type description, it means we are working with a pointer to a person
	(*p).firstName = newName // *p refers to an operator, it means we want to manipulate the value stored in the pointer
}

// &variable: gives access to the memory address of the value the variable is pointing to
// *pointer: give me the value this memory address is pointing

// turn 'address' into 'value' via *address
// turn 'value' into 'address' via &value

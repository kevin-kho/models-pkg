package models

type Person struct {
	Name     string
	Age      int
	Chickens []Chicken
}

type Chicken struct {
	Name string
	Age  int
}

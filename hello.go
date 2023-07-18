package main

type Student struct {
	name string
	age  int
}

func (s Student) getFullInfo() string {
	return s.name + s.name
}

func main() {

}

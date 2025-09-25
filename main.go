package main

import "fmt"

type Hero struct {
	Name  string
	Age   int
	Level int
}

func (h Hero) Show() {
	fmt.Println(h.Name, h.Age, h.Level)
	fmt.Println("Age = ", h.Age)
	fmt.Println("Name = ", h.Name)
	fmt.Println("Level = ", h.Level)
}

func (h Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(name string) {
	h.Name = name
}

func (h Hero) GetAge() int {
	return h.Age
}

func (h *Hero) SetAge(age int) {
	h.Age = age
}

func (h Hero) GetLevel() int {
	return h.Level
}

func (h *Hero) SetLevel(level int) {
	h.Level = level
}

func main() {
	hero := Hero{Name: "John", Age: 20, Level: 1}
	hero.Show()
	hero.SetName("Jane")
	hero.SetAge(21)
	hero.SetLevel(2)
	hero.Show()
}

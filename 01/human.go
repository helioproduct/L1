package main

import "fmt"

type Human struct {
	Identity int
	FullName string
	Years    int
}

// Метод для текстового представления структуры Person
func (p Human) ToString() string {
	return fmt.Sprintf("FullName: %s; Years: %d; Identity: %d", p.FullName, p.Years, p.Identity)
}

// Метод для приветствия человека
func (p Human) SayHello() {
	fmt.Printf("Hello, I'm %s, I'm %d years old\n", p.FullName, p.Years)
}

type Action struct {
	Human    // Встраивание Person в структуру Movement
	Identity int
	Activity string
}

// Метод для текстового представления структуры Movement
func (m Action) ToString() string {
	// К вложенной структуре можно обратиться напрямую, написав ее тип.
	// Так как p.Person реализует интерфейс Stringer,
	// то fmt.Printf вызовет метод ToString у структуры Person
	return fmt.Sprintf("Person: (%v); Activity: %s; Identity: %d", m.Human, m.Activity, m.Identity)
}

func main() {
	m := Action{
		Human: Human{
			FullName: "Alex",
			Years:    25,
			Identity: 100,
		},
		Activity: "running",
		Identity: 5,
	}

	// Покажем, что методы вложенных структур можно вызвать напрямую
	m.SayHello() // print "Hello, I'm Alex, I'm 25 years old"
}

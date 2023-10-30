package main

import "fmt"

type Animal struct {
	name    string
	species string
}

type Cage struct {
	number int
	animal *Animal
}

type Zookeeper struct {
	name    string
	current *Cage
}

func (z *Zookeeper) CatchAnimal(a *Animal) {
	z.current.animal = a
	fmt.Printf("Zookeeper catch animal %s\n", a.name)
}

func (z *Zookeeper) MoveToNextCage() { // first cage

	if z.current == nil {
		fmt.Println("Enter first cage")
		z.current = &Cage{number: 1}
		return
	}

	z.current = &Cage{number: z.current.number + 1}
	fmt.Printf("Current cage number is %d\n", z.current.number)

}

const (
	CagesCount = 5
)

func main() {

	cat := Animal{name: "Кіт", species: "Котячі"}
	elephant := Animal{name: "Слон", species: "Хоботні"}
	capybara := Animal{name: "Капібара", species: "Ссавці"}
	zebra := Animal{name: "Зебра", species: "Зеброві"}
	lion := Animal{name: "Лев", species: "Котячі"}

	//animals := []Animal{cat, elephant, capybara, zebra, lion}

	cage1 := Cage{number: 1, animal: &cat}
	//cage2 := Cage{number: 2, animal: &elephant}
	//cage3 := Cage{number: 3, animal: &capybara}
	//cage4 := Cage{number: 4, animal: &zebra}
	//cage5 := Cage{number: 5, animal: &lion}

	zookeeper := Zookeeper{name: "Ловець", current: &cage1} //починає з першої клітки

	zookeeper.CatchAnimal(&cat) // first day of work
	zookeeper.MoveToNextCage()

	zookeeper.CatchAnimal(&elephant) // second day of work
	zookeeper.MoveToNextCage()

	zookeeper.CatchAnimal(&capybara) // third day of work
	zookeeper.MoveToNextCage()

	zookeeper.CatchAnimal(&zebra) // fourth day of work
	zookeeper.MoveToNextCage()

	zookeeper.CatchAnimal(&lion) // fifth day of work
	zookeeper.MoveToNextCage()

}

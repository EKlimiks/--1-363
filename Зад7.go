package main
import "fmt"

type Book struct{
	ID int
	Name string
	Author string
	Year int
	Status string
	PagesCount int
}

func main() {
	B1 := Book{1, "Война и мир", "Толстой", 1869, "доступна", 122}
	B2 := Book{2, "Мертвые души", "Гоголь", 1842, "выдана", 332}
	B3 := Book{3, "Мастер ", "Булгаков", 1967, "доступна", 480}
	B1.Status = "выдана"
	fmt.Println(B1)
	fmt.Println(B2)
	fmt.Println(B3)
}
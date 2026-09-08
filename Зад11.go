package main
import "fmt"

type UserProfile struct{
	Username string
	Age int
	FriendsCount int
	IsVerified bool
	Rating float64
}

func main() {
	i := UserProfile{"Eklimiks", 18, 1000,true, 9.9}
	Friend1 := UserProfile{"PeLMeNY", 18, 140,true, 7.9}
	Friend2 := UserProfile{"Klemetev", 17, 10,false, 3.9}
	fmt.Println(i)
	fmt.Println(Friend1)
	fmt.Println(Friend2)
}
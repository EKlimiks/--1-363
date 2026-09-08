package main

import "fmt"

const (
	Online       = "online"
	Offline      = "offline"
	Away         = "away"
	DoNotDisturb = "dnd"
)
func main() {
	statuses := map[string]string{
		"Егор":   Online,
		"Дима":     Offline,
		"Настя": Away,
		"Оля":   DoNotDisturb,
	}
	statuses["Дима"] = Online
	statuses["Настя"] = Online

	fmt.Println("В сети:")
	for name, status := range statuses {
		if status == Online {
			fmt.Println(name)
		}
	}
}
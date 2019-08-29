package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float32)

	grades["Ronaldo"] = 42
	grades["Romario"] = 92
	grades["Ronaldinho"] = 67

	fmt.Println(grades)


	delete(grades, "Ronaldo")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k,":",v)
	}

}
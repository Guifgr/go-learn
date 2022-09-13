package main

import (
	"fmt"
	"os"
	os/exec"
"time"
)

func main() {
	Clear()
  
 loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	firstDate := time.Date(now.Year(), 1, 1,0, 0, 0, 0, loc)
	diff := ow.Sub(firstDate)
  
	i int(diff.Hours()/24)+1 == 256 {
	fmt.Printf("Feliz dia do programador!")
	} else {
	fmt.Printf("No meu pc funciona!")
	}
	
	fmt.Scanln()
}

func Clear() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
}

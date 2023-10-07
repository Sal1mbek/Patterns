package main
import (
	"fmt"
	"time"
)
func main(){
	start := time.Now()
	
	defer func(){
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Bob", "Alice", "Shreder", "Tony"}
	myChan := make(chan bool)	
	for _, evilNinja := range evilNinjas{
		
		go kill(evilNinja, myChan)
	}

	for i := 0; i<len(evilNinjas) ; i++{
		<- myChan
	}
}


func kill(target string, myChan chan bool ) {
	
	time.Sleep(time.Second)

	fmt.Println("Throwing ninja starts at", target)	
	myChan <- true
}
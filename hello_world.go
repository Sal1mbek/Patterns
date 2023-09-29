package main
import "fmt"
func main(){
	evilNinjas := []string{"Bob", "Alice", "Shreder", "Tony"}
	for _, evilNinja := range evilNinjas{
		kill(evilNinja)
	}
}

func kill(target string) {
	fmt.Println("Throwing ninja starts at", target)
	time.Sleep(time.Second)
}
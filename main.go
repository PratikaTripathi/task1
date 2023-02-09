package packagePrint

import (
	"fmt"

	"time"
)

func Print() {
	fmt.Println("Hello, 世界")

	c := make(chan int)
	for i := 0; i < 5; i++ {
		if i == 0 {
			DOSomething(i, nil)
			fmt.Println(i)
		} else {
			go DOSomething(i, c)
		}
	}

	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
}

func main() {
	print()
}

func DOSomething(times int, c chan int) {
	time.Sleep(1 * time.Second)
	if c != nil {
		c <- times
	}

}

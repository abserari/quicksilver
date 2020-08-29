package main
import (
	"fmt"
	"runtime"
	"time"
)
func main() {
	r := NewReceiveScheduler(100)
	for i := 0; i < 100; i++ {
		r.Add("1","haha",func(x interface{}) {fmt.Println(x); time.Sleep(100 * time.Millisecond)})
	}
	
	fmt.Println("hello")

	runtime.Goexit()
}
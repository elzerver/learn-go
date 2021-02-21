package main

import (
	"fmt"
)

func main() {
	fn := []string{"james","jonas","john", "judas"}
	foo("Todd", fn...)
}

func foo(s string, walalang ...string){
	fmt.Println(s, walalang)
}
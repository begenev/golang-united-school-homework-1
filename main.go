package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func SayHello() string {
	return emoji.Sprintf("Hello :world_map:!")
}

func main() {
	fmt.Println(SayHello())
}

package main

import (
	"github.com/kyokomi/emoji"
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {

	res := SayHello()
	expected := emoji.Sprintf("Hello :world_map:!")
	if res != expected {
		log.Fatalf("Not equal. Expected (%v), actual (%v)", expected, res)
	}

}

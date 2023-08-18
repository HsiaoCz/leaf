package snowflak

import (
	"fmt"
	"github/HsiaoCz/leaf/etc"
	"log"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	err := etc.Init()
	if err != nil {
		t.Fatal(err)
	}
	err = Init()
	if err != nil {
		log.Fatal(err)
	}
	id := GenID()
	fmt.Println(id)
}

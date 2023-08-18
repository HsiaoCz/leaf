package snowflak

import (
	"fmt"
	"log"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	id := GenID()
	fmt.Println(id)
}

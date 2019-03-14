package goPattern

import (
	"testing"
	"fmt"
)

func Test_singleton(t *testing.T){
	con:=NewConfig("127.0.0.1",8080)
	fmt.Println(con)

	con2:=NewConfig("127.0.0.1",80)
	fmt.Println(con2)
}

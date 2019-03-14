package goPattern

import (
	"testing"
	"fmt"
)

func Test_Factory(t *testing.T){
	che:=BuildCheFactory("bike")
	fmt.Println("车速为：",che.speed())

	che=BuildCheFactory("car")
	fmt.Println("车速为：",che.speed())

	che=BuildCheFactory("f1")
	fmt.Println("车速为：",che.speed())

	che=BuildCheFactory("tesila")
	fmt.Println("车速为：",che.speed())
}

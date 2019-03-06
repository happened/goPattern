package goPattern

import (
	"testing"
	"fmt"
)

func Test_builder(t *testing.T){
	myOrder:=NewOrder().withCola().withHamburger()
	meal:=myOrder.confirmOrder()
	fmt.Println(meal.taocanming,"\n价格：",meal.price)
}
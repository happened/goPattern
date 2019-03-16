package goPattern

import (
	"testing"
	"fmt"
)

func TestProxy(t *testing.T){
	buyer1:=new(house)
	fmt.Println("直接去买房子的价格为：",buyer1.GetPrice())

	buyer2:=houseProxy{buyer1}

	fmt.Println("通过中介去买的价格为：",buyer2.GetPrice())


}

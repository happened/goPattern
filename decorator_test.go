package goPattern

import (
	"testing"
	"fmt"
)

func Test_decorator(t *testing.T){
	myfunc:=DecoratorFunc(sum1)
	myfunc2:=DecoratorFunc(sum2)

	fmt.Printf("结果分别为： %d , %d \n",myfunc(-1000,20000000),myfunc2(-1000,20000000))
}

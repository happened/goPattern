package goPattern

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	//成为一名店主
	dianzhu := new(shopKeeper)
	//下雨了 找到策略
	dianzhu.NewFunc("下雨")
	fmt.Println(dianzhu.book.getPrice())

	dianzhu.NewFunc("大太阳")
	fmt.Println(dianzhu.book.getPrice())
}

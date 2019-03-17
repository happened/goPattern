package goPattern

import (
	"testing"
)

func TestObserver(t *testing.T){
	board:=blackboard{
		records:map[observer]struct{}{},
	}

	//注册两个观察者
	board.Register(&customer{name:"sheng"})
	board.Register(&customer{name:"feng"})

	//降价通知
	board.Notify(event{msg:"打八折"})


}

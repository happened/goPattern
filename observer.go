/*
观察者模式的实现
故事：有一件衣服 平时太贵很多人喜欢 但买不起 ，但这些人每天都看着价钱,
店家门口还有个黑板 用来记录这些想要衣服的人，要就把名字写上去，不要划掉就行
有一天 衣服降价了，黑板上也贴出了这个事件，这一事件发生后，
自动通知给这些关注衣服的人，他们接收到之后就会去买
*/
package goPattern

import (
	"fmt"
)

//发生了什么事件
type event struct {
	msg string
}

//观察者接口 作为观察者都会在接收到事件后进行相关行为
type observer interface {
	OnNotify(eve event)
}

//媒介（黑板报之类的） 用来记录观察者 移除观察者 和进行通知
type notiffer interface {
	Register(obs observer)	//添加观察者
	DeRegister(obs observer)//移除观察者
	Notify(eve event)
}

//买衣服的顾客
type customer struct {
	name string
}
//顾客接收到降价通知后 就把这个消息大声说出来
func (buyer *customer) OnNotify(eve event){
	fmt.Println("衣服降价了",eve.msg)
}

//店家门口的黑板报
type blackboard struct {
	//记录有哪些顾客想买
	records map[observer]struct{}
}
//将想买的顾客写入到记录中
func (board *blackboard)Register(obs observer){
	board.records[obs]= struct{}{}
}
//不想买的顾客就把记录删除掉
func (board *blackboard)DeRegister(obs observer){
	delete(board.records,obs)
}
//通知记录中的顾客发生了事件 如降价
func (board *blackboard)Notify(eve event){
	for o:=range board.records{
		o.OnNotify(eve)
	}
}
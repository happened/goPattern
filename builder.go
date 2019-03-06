package goPattern

import "bytes"

//builder设计模式 主要解决多变参数传递问题

//故事：平时去肯德基吃饭，有各种套餐(鸡腿 汉堡 可乐等)
//第一个客人：鸡腿+汉堡
//第二个客人：鸡腿+可乐
//通过不同的点餐配合生成了不同价格的套餐

//套餐结构
type setMeal struct {
	taocanming string//套餐名
	price int	//套餐价格

}

var (
	drumStickPrice  =8	//鸡腿价格
	hamburgerPrice  =5	//汉堡价格
	colaPrice =4	//可乐价格
)


type Order struct{
	WithDrumStick bool //是否加鸡腿
	WithHamburger bool //是否加汉堡
	WithCola bool //可乐
}


//监督那些建造者
func NewOrder() *Order{
	return &Order{}
}

//下面三个是各部件的建造方式 就是建造者

//加鸡腿
func (myOrder *Order) withDrumStick() *Order{
	myOrder.WithDrumStick=true
	return myOrder
}
//加汉堡
func (myOrder *Order) withHamburger() *Order{
	myOrder.WithHamburger=true
	return myOrder
}
//加可乐
func (myOrder *Order) withCola()*Order{
	myOrder.WithCola=true
	return myOrder
}

//建造了一个真正的订单

func (myOrder *Order)confirmOrder() setMeal{
	var buffer bytes.Buffer
	var price int
	price=0

	buffer.WriteString("订单如下：")
	if myOrder.WithDrumStick{
		buffer.WriteString("鸡腿。")
		price+=drumStickPrice
	}
	if myOrder.WithHamburger{
		buffer.WriteString("汉堡。")
		price+=hamburgerPrice
	}
	if myOrder.WithCola{
		buffer.WriteString("可乐。")
		price+=colaPrice
	}

	//返回一个各部分组成的复杂对象
	var meal setMeal
	meal.taocanming=buffer.String()
	meal.price=price
	return meal
}

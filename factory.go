/*
	factory 工厂模式
	故事：有个可以建造不同种类车辆的工厂 车辆速度不同
	想创建自行车 轿车 F1方程式等不同类型
	想知道不同车开起来有多快
	只需要告诉工厂要什么车  工厂就生产出该车型

*/
package goPattern

//重新声明
type cheType string

const (

	CarType cheType="car"
	F1Type cheType="f1"
	TesilaType cheType="tesila"
)

//车辆接口
type cheliang interface {
	speed() int	//车跑起来都有速度
}

// 自行车
type bike struct {
}

//自行车车速10
func (che *bike)speed()int{
	return 10
}
//小汽车
type car struct {
}

//小汽车车速80
func (che *car)speed()int{
	return 80
}

//F1方程式
type f1 struct {
}

//方程式车速120
func (che *f1)speed()int{
	return 120
}

//特斯拉
type tesila struct {
}

//特斯拉车速10000
func (che *tesila) speed()int{
	return 10000
}

func BuildCheFactory( t cheType) cheliang{
	switch t {
	case CarType:
		return new(car)
	case F1Type:
		return new(f1)
	case TesilaType:
		return new(tesila)
	//默认为自行车
	default:
		return new(bike)
	}
}


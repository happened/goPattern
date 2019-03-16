/*
	代理模式
	故事：比如有一个房子 我们直接去买 可能要100万，但是有一个人和房东很熟，可以打八折，
	我们就可以通过这个中介去买房，这个中介能对房价进行一定的减免
*/

package goPattern


//建筑物接口 是个建筑都有价格
type building interface {
	GetPrice() float64
}

//家庭住宅也是房子 也有价钱
type house struct {
	value float64
}

func (h *house) GetPrice() float64{
	h.value=100
	return h.value
}

//一个中介手中有一个房源 房源可以更换
type houseProxy struct {
	goodHouse *house
}
//通过中介手中的房子可以打8折
func (houseP *houseProxy)GetPrice() float64{
	if houseP.goodHouse==nil{
		houseP.goodHouse=new(house)
	}
	houseP.goodHouse.value*=0.8
	return houseP.goodHouse.value
}
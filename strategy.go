/*
策略模式
故事：一家店根据不同的天气情况 修改雨伞的价钱 天晴雨伞就比原价便宜10块 下雨就比原价高10块
*/
package goPattern

//雨伞阴天平时卖100块
const firstPrcie = 100

//不同的策略都为了完成一个目的 伞的价钱
type price interface {
	getPrice() int
}

//大晴天的伞
type sunUmbrella struct {
}

func (su *sunUmbrella) getPrice() int {
	return firstPrcie - 10
}

//下雨天的伞
type rainUmbrella struct {
}

func (ru *rainUmbrella) getPrice() int {
	return firstPrcie + 10
}

//利用工厂模式对不同的天气返回不同的定价伞
//店家有一本神书 记载了所有的定价策略 店主把天气输入后 利用这个书就可以运用对应的策略啦

//说一下工厂模式和策略模式的不同 工厂模式关注创建了一个新的店主 这个店主有根据天气更新获得新的价钱的途径
//但是具体更新的能力不是他关心的  比如 下雨天涨价还是降价 他不知道 他只是告诉神书 下雨了 还是天晴了，然后再从神书获取价格
type shopKeeper struct {
	book price
}

func (sk *shopKeeper) NewFunc(status string) {
	sk.book = nil
	switch status {
	case "下雨":
		sk.book = new(rainUmbrella)
	default:
		sk.book = new(sunUmbrella)
	}
}

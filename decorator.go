/*
	实现装饰器模式
	故事：由两个函数 分别进行不同的计算 记录两个计算的时间  并且满足当更换计算方法时 仍能使用该计算代码
	主要避免实现一些小功能的重复

	在net包中的handlerFunc一般有这样用
*/

package goPattern

import (
	"time"
	"fmt"
	"runtime"
	"reflect"
)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
//不同的计算方法逻辑
func sum1(start ,end int)int{
	if start >end{
		start,end=end,start
	}

	sum:=0
	for ;start<end;start++{
		sum+=start
	}
	return sum
}

func sum2(start ,end int)int{
	if start >end{
		start,end=end,start
	}

	return (end-start+1)*(end+start)/2
}

type sumFunc func(int ,int)int

//计时间函数
func DecoratorFunc(f sumFunc )sumFunc{
	return func(start ,end int)int {

		defer func(t time.Time){
			fmt.Printf("---time elapsed---%s:  %v ---\n",getFunctionName(f),time.Since(t))
		}(time.Now())

		return f(start,end)
	}
}
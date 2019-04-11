package main

import (
	"fmt"
)

var lieCount  = 0
var pre []uint32  //  存放节点的前置节点
var relation []uint8  //  存放节点与前置节点（或根节点）的关系
var n uint32

const (
	R_SAME = 0
	R_EAT = 1
	R_BE_EAT = 2
)

const (
	D_SAME = 1
	D_EAT = 2
)

func main() {
	fmt.Printf("input the maxAnimalCount: ")
	_, err1 := fmt.Scanln(&n)
	if checkError(err1) {
		return
	}
	before()

	fmt.Printf("input the sentence, example: 1 2 3\n")
	for {
		var d, x, y uint32
		_, err :=fmt.Scanln(&d, &x, &y)
		if checkError(err) {
			break
		}
		if !addsSentence(d, x, y) {
			fmt.Printf("lie\n")
			lieCount ++
		}
	}
	fmt.Printf("lieCount: %d\n", lieCount)
}


func before() {
	pre = make([] uint32, n)
	relation = make([] uint8, n)
	for i := uint32(0);i < n;i++ {
		pre[i] = i
		relation[i] = R_SAME
	}
}

/**
添加一条句子, 返回该句子是否为真
 */
func addsSentence(d uint32, x uint32, y uint32) bool {
	if x > n || y >n || (x == y && d == D_EAT) {
		return false
	}
	// 同一连通图内
	if find(x) == find(y) {
		if d == D_SAME && relation[x] != relation [y] {
			return false
		}
		if d == D_EAT && !xEatY(x, y) {
			return false
		}
		return true
	}
	// 不在同一连通图内，合并
	merge(d, x, y)
	return true
}

func find(x uint32) uint32 {
	if x == pre[x] {
		return x
	}
	tPre := pre[x]
	// find 一次之后将直接指向根节点,
	pre[x] = find(pre[x])
	// 关系传递公式 吃+吃 = 被吃， 同类+吃/被吃 = 吃/被吃 吃+被吃 = 同类
	relation[x] = (relation[x] + relation[tPre]) % 3
	return pre[x]
}

func merge(d uint32, x uint32, y uint32)  {
	fx := find(x)
	fy := find(y)
	pre[fy] = fx
	// 此关系公式可由关系传递公式得出
	relation[fy] = (relation[x] - relation[y] + 3 + uint8(d) - 1) % 3
}

func xEatY(x uint32, y uint32) bool {
	//  此公式可画循环图枚举得到
	return 	(relation[x] + 1) % 3 == relation[y]
}

func checkError(err error) bool {
	return err != nil
}


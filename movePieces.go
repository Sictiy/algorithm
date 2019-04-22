package main

import (
	"fmt"
	"math/rand"
)

var pieces []byte
var lastFrom, lastTo int
var lastMoveIsOdd bool

const (
	RED = 0
	BLUE = 1
)

const (
	AUTO = 0
	MANUAL = 1
)
//  是否跟随奇数棋子移动
const OPEN_FOLLOW = false
const FIRST_MOVE  = BLUE

func main() {
	initPieces()
	play()
}

/**
初始化棋子数组
 */
func initPieces()  {
	fmt.Print("input pieces size:")
	var piecesSize int
	_, err := fmt.Scanln(&piecesSize)
	if err != nil || piecesSize <= 0 {
		return
	}
	pieces = make([]byte, piecesSize)
	fmt.Print("input the piece`s index : \n")
	maxIndex := 0
	for {
		var index int
		_, err := fmt.Scanln(&index)
		if err != nil {
			break
		}
		pieces[index] = 1
		if index > maxIndex{
			maxIndex = index
		}
	}
	pieces = pieces[0:maxIndex+1]
	fmt.Printf("final pieces: \n")
	fmt.Print(pieces)
	fmt.Print("\n")
}

func play() {
	curPlayer := FIRST_MOVE
	for playerPlay(curPlayer) {
		fmt.Printf("move success! \n")
		fmt.Print(pieces)
		fmt.Print("\n")
		curPlayer ^= 1
	}
	// 自己先不能动，输
	if curPlayer == FIRST_MOVE {
		fmt.Printf("lose!\n")
	}else {
		fmt.Printf("win!\n")
	}
}

/**
玩家操作，无法操作返回false
 */
func playerPlay(player int) bool {
	// 失败
	if findCanMove(pieces) == 0 {
		return false
	}
	var from, to = 0, 0
	for {
		//  自动移动
		if getPlayerRunMode(player) == AUTO {
			from, to = findMoveStep(pieces)
		} else{
			fmt.Printf("input from a to b:")
			_, err :=fmt.Scanln(&from, &to)
			if err!= nil {
				continue
			}
		}
		fmt.Printf("player %s move: %d -> %d \n",getPlayer(player), from, to)
		if move(from, to) {
			break
		}
	}
	return true
}

func getPlayerRunMode(player int) int {
	switch player {
	case RED:
		return AUTO
	case BLUE:
		return MANUAL
	default:
		return MANUAL
	}
}

func getPlayer(player int) string {
	switch player {
	case RED:
		return "RED"
	case BLUE:
		return "BLUE"
	default:
		return "UNKNOWN"
	}
}

/**
获取第一个可以移动的棋子
 */
func findCanMove(pieces []byte) int {
	curByte := 1
	for k, v := range pieces {
		if v == 0 {
			curByte = 0
		}
		if curByte == 0 && v == 1 {
			return k
		}
	}
	return 0
}

func getXOR(intervalMap map[int]int) int {
	var result  = 0
	for _, v := range intervalMap {
		result ^= v
	}
	return result
}

/**
移动棋子，无法移动返回false
 */
func move (from int, to int) bool {
	if from <= to || from >= len(pieces) || pieces[from] == 0{
		fmt.Printf("can`t move from`s piece, from index: %d \n", from)
		return false
	}
	for i := from - 1; i >= to; i-- {
		if pieces[i] == 1 {
			fmt.Printf("can`t move to index: %d \n", to)
			return false
		}
	}
	lastMoveIsOdd = true
	for i:= len(pieces) - 1; i >= from; i-- {
		if pieces[i] == 1 {
			lastMoveIsOdd = !lastMoveIsOdd
		}
	}
	pieces[from], pieces[to] = 0, 1
	lastFrom, lastTo = from, to
	return true
}

/**
获取棋子的间隔map
 */
func findIntervalMap(pieces []byte)  map[int] int{
	var index = -1
	result := make(map[int] int)
	for i:= len(pieces) - 1; i>= 0; i-- {
		// 有棋子
		if pieces[i] == 1 {
			if index == -1 {
				index = i
			}else {
				result[index] = index - i - 1
				index = -1
			}
		}
	}
	if index != -1 {
		result[index] = index
	}
	return result
}

func findMoveStepFromMap(intervalMap map[int] int) (int, int){
	valueXor := getXOR(intervalMap)
	if valueXor == 0 {
		//  没有必胜策略
		for k, v := range intervalMap {
			if v != 0 {
				return k, k - rand.Intn(v) - 1
			}
		}
		canMove := findCanMove(pieces)
		return canMove, canMove - 1
	}
	var leftMoveTimes uint = 0
	var tempValue = valueXor
	for valueXor != 1 {
		leftMoveTimes++
		valueXor >>= 1
	}
	maxByte := 1 << leftMoveTimes
	for k,v := range intervalMap {
		if v & maxByte != 0 {
			changeTo := v ^ tempValue
			return k, k - v + changeTo
		}
	}
	return 0, 0
}

/**
自动寻找策略，奇数偶数两两组合化为nim游戏
 */
func findMoveStep(pieces []byte) (int, int) {
	// 上次移动的奇数棋子，可以移动对应的偶数棋子
	if OPEN_FOLLOW && lastMoveIsOdd {
		for i := lastFrom; i < len(pieces); i++{
			if pieces[i] == 1 {
				return i, i - (lastFrom - lastTo)
			}
		}
	}
	return findMoveStepFromMap(findIntervalMap(pieces))
}
package woker

import (
	"fmt"
	"poker/pkg/model"
	"poker/poker/util"
	"testing"
	"time"
)

func TestGetMaxSameColorStraight(t *testing.T) {
	poker := "TcJcXnKc9cQcTc"
	maxPoker, priority := FindMaxPokerPriority(poker)
	moreMaxPoker, priority := GetGhostMaxPoker(poker, maxPoker, priority)
	fmt.Printf("\npoker:%s,maxPoker:%s,moreMaxPoker:%s,priority:%d\n", poker, maxPoker, moreMaxPoker, priority)
}

func TestFourRazz(t *testing.T) {
	poker1 := "8s8hXn9c8d8cTd"
	maxPoker, priority := FindMaxPokerPriority(poker1)
	moreMaxPoker, priority := GetGhostMaxPoker(poker1, maxPoker, priority)
	fmt.Printf("\npoker:%s,maxPoker:%s,moreMaxPoker:%s,priority:%d\n", poker1, maxPoker, moreMaxPoker, priority)
}

// 测试获取出现大于等于count次数的颜色
func TestGetColor(t *testing.T) {
	poker := "8d8dXn9d8d8dTd"
	colors := GetColor(poker, 2)
	fmt.Printf("\ncolors:")
	for i := 0; i < len(colors); i++ {
		fmt.Printf("\t %s", string(colors[i]))
	}
}

/**
	获取赖子变成的最大附属牌
 	输入需要排除的牌和返回的附属牌颜色，输出癞子可变的最大的牌
	1. 输入颜色，代表要求的是同花顺的最大附属牌，那么获取的牌必须要和excludes的数不一样
	2. 颜色为0，代表只需要除了excludes之外的最大的牌即可
*/
func TestGetMaxSinglePokerByRazz(t *testing.T) {
	// poker := "2d8sTdAdThQh3h"
	poker := GetMaxSinglePokerByRazz('c', "Td", "Th", "Ac")
	fmt.Printf("\nGetMaxSinglePokerByRazz.poker:%s\n", poker)
}

// 判断是否具备顺子条件
// 返回缺的牌
func TestIsToBeStraight(t *testing.T) {
	// poker := "2d3d4d6dXn3h3h"
	poker := "Kc7hXnJdTc9d8c"
	flag, needNum := IsToBeStraight(poker)
	fmt.Printf("\nIsToBeStraight.poker:%s\n", poker)
	fmt.Printf("flag:%v:=:%s\n", flag, string(needNum))
}

// 判断是否具备变成同花顺条件
func TestToBeSameStraight(t *testing.T) {
	poker := "2d3d4s6dXn3h3c"
	flag, maxPoker, priority := ToBeSameStraight(poker)
	fmt.Printf("\nIsToBeStraight.poker:%s\n", poker)
	fmt.Printf("flag:%v:maxPoker:%s,priority:%d\n", flag, maxPoker, priority)
}

func TestStraightRazz(t *testing.T) {
	poker := "XnKc9s7dTc6d8s"
	maxPoker := "Tc9s8s7d6d"
	priority := 5
	moreMaxPoker, priority2 := StraightRazz(poker, maxPoker, priority)
	fmt.Printf("\nStraightRazz.poker:%s,maxPoker:%s,priority:%d\n", poker, maxPoker, priority)
	fmt.Printf("moreMaxPoker:%s,priority2:%d\n", moreMaxPoker, priority2)
}

func TestHighPoker(t *testing.T) {
	poker := "Kd7d4c9h8hQdXn"
	maxPoker := "KdQd9h8h7d"
	priority := 1
	moreMaxPoker, priority2 := HighPoker(poker, maxPoker, priority)
	fmt.Printf("\nStraightRazz.poker:%s,maxPoker:%s,priority:%d\n", poker, maxPoker, priority)
	fmt.Printf("moreMaxPoker:%s,priority2:%d\n", moreMaxPoker, priority2)
}

/**
输入：1～3张相同大小的牌，不同花色
将三条相扩展为4条
一对 变成 三条
散牌变成一对
*/
func TestGetSameNumExtension(t *testing.T) {
	poker := "AdAsAh"
	maxPoker := GetSameNumExtension(poker)
	fmt.Printf("\nIsToBeStraight.poker:%s\n", poker)
	fmt.Printf("maxPoker:%s\n", maxPoker)
}

/*
	输入4~5张同花，获取最大同花
	4 --> 5
	5 --> 5
*/
func TestGetMaxSameColorPokerByRazz(t *testing.T) {
	poker := "4d3d7dAd8d"
	maxPoker := GetMaxSameColorPokerByRazz(poker)
	fmt.Printf("\nIsToBeStraight.poker:%s\n", poker)
	fmt.Printf("maxPoker:%s\n", maxPoker)
}

func TestSevenPokerWithGhostCompare(t *testing.T) {
	alice := "XnJdTc9d8c6h6c"
	bob := "Kc7hXnJdTc9d8c"
	result := SevenPokerWithGhostCompare(alice, bob)
	fmt.Printf("result:%d\n", result)
}

func SevenWithGhost() {
	var dataAll model.DataJson
	var errResults []model.Result
	util.ReadFile("./src/seven_cards_with_ghost.result.json", &dataAll)
	startT := time.Now()
	count := 0 // 对的
	for _, data := range dataAll.Matches {
		myResult := SevenPokerWithGhostCompare(data.Alice, data.Bob)
		if myResult == data.Result {
			count++
		} else {
			errData := model.Result{
				Alice:       data.Alice,
				Bob:         data.Bob,
				Result:      data.Result,
				ErrorResult: myResult,
			}
			errResults = append(errResults, errData)
		}
	}
	tc := time.Since(startT) // 计算耗时
	errDatas := model.DataJson{
		Matches: errResults,
	}
	util.WriteFile("pkg/seven_cards_with_ghost_err.json", &errDatas)
	rate := float64(count) / float64(len(dataAll.Matches)) * 100
	fmt.Printf("比率是%f", rate)
	fmt.Printf("耗时 = %v\n", tc)
}

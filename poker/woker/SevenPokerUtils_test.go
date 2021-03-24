package woker

import (
	"fmt"
	"testing"
)

func TestIsContainSameColor(t *testing.T) {
	poker := "8sTs9c2sJh6sTs"
	isSameColor, color := IsContainSameColor(poker)
	fmt.Printf("\npoker:%s,isSameColor:%v,color:%s", poker, isSameColor, string(color))
}

func TestFindSameColor(t *testing.T) {
	poker := "8sTs9c2sJh6sTs"
	str := FindSameColor(poker, 's')
	fmt.Printf("\npoker:%s,FindSameColor:%v,color:%s\n", poker, str, string('s'))
}

func TestGetMaxStraight(t *testing.T) {
	poker := "Qs4d3c2h5dAsKh"
	isStragint := IsStraight(poker)
	fmt.Printf("\npoker:%s,isStragint:%v\n", poker, isStragint)
	if isStragint {
		str := GetMaxStraight(poker)
		fmt.Printf("\npoker:%s,GetMaxStraight:%v\n", poker, str)
	}
}

func TestGetMaxPoker(t *testing.T) {
	poker := "AsTsKsQsJs6sTs"
	str := GetMaxPoker(poker, 5)
	fmt.Printf("\npoker:%s,GetMaxPoker:%v\n", poker, str)
}

func TestGetNumsMaxSinglePoker(t *testing.T) {
	poker := "Qs4d3c2h5dAsKh"
	str := GetNumsMaxSinglePoker(poker, 1, "Jh", "Td")
	fmt.Printf("\npoker:%s,GetMaxPoker:%v\n", poker, str)
}

func TestGetSameNumPokerBySort(t *testing.T) {
	poker := "4dKc3hThTd2sQh"
	strs := GetSameNumPokerBySort(poker, 2)
	fmt.Printf("\npoker:%s,GetMaxPoker:%v\n", poker, strs)
}

func TestFindMaxPokerPriority(t *testing.T) {
	poker1 := "TcJc2cQdAc8cKh"
	poker2 := "Kc3sTcJc2cQdAc"
	maxPoker1, result1 := FindMaxPokerPriority(poker1)
	maxPoker2, result2 := FindMaxPokerPriority(poker2)
	fmt.Printf("\npoker1:%s,poker1:%s\n", poker1, poker2)
	fmt.Printf("\nmaxPoker1:%s,Priority:%d,maxPoker2:%s,Priority:%d\n", maxPoker1, result1, maxPoker2, result2)
}

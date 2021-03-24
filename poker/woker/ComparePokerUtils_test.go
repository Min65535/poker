package woker

import (
	"fmt"
	"log"
	"testing"
)

// 测试皇家同花顺

func TestPriorityTen(t *testing.T) {
	/*var dataJson model.DataJson
	ReadFile("../pkg/match.json",&dataJson)
	for _,data := range dataJson.Matches {
		log.Printf("%s=%v\t",data.Alice,IsSameColor(data.Alice))
	}*/
	poker := "AdKdQdJdTd"
	log.Printf("IsSameColor%s=%v\n", poker, IsSameColor(poker))
	log.Printf("IsStraight%s=%v\n", poker, IsStraight(poker))
	log.Printf("IsAKQJT%s=%v\n", poker, IsAKQJT(poker))
}

// 测试四条
func TestPriorityEight(t *testing.T) {
	poker := "TdTh5dTsTc"
	log.Printf("GetMaxSameCount%s=%v\n", poker, GetMaxSameCount(poker))
}

// 测试判断三带二
func TestPrioritySeven(t *testing.T) {
	poker := "QdKhTdQsQc"
	log.Printf("GetMaxSameCount%s=%v\n", poker, GetMaxSameCount(poker))
	log.Printf("IsThreeAndTwo%s=%v\n", poker, IsThreeAndTwo(poker))
}

// 测试判断是否是两对
func TestPriorityThree(t *testing.T) {
	poker := "QdKhKdQsTc"
	log.Printf("GetMaxSameCount%s=%v\n", poker, GetMaxSameCount(poker))
	log.Printf("IsThreeAndTwo%s=%v\n", poker, IsTwoTwice(poker))
}

func TestQuickSort(t *testing.T) {
	poker := "3cTs3s9dTc"
	arr1 := []uint8{poker[0], poker[2], poker[4], poker[6], poker[8]}
	fmt.Printf("排序前：%s", arr1)
	QuickSort(arr1, 0, len(arr1)-1)
	fmt.Printf("排序后：%s", arr1)
}

func TestCompare(t *testing.T) {
	poker1 := "8h9c9h6d8c"
	poker2 := "5s7s7c5cKc"
	pokerCompareTest(poker1, poker2)
}

func pokerCompareTest(poker1 string, poker2 string) {
	fmt.Printf("poker1的优先级：%d", GetPriority(poker1))
	fmt.Printf("poker2的优先级：%d", GetPriority(poker2))
	// fmt.Printf("poker1和poker2比较大小结果:%d",PokerCompare(poker1,poker2))
}

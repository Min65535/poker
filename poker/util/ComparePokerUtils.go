package util

import (
	"strings"
)

func PokerCompare(alice string, alicePriority int, bob string, bobPriority int) int {
	compare := 0
	if alicePriority > bobPriority {
		return 1
	}
	if bobPriority > alicePriority {
		return 2
	}
	// 优先级相等之后的比较
	switch alicePriority {
	case 10:
		compare = 0 // 皇家同花顺比较
		break
	case 9: // 同花顺
		compare = StraightCompare(alice, bob)
		break
	case 8: // 四条
		compare = FourOfAKindCompare(alice, bob)
		break
	case 7: // 满堂彩
		compare = ThreeWithTwoCompare(alice, bob)
		break
	case 6: // 同花
		compare = SortCompare(alice, bob)
		break
	case 5: // 顺子
		compare = StraightCompare(alice, bob)
		break
	case 4: // 三条
		compare = ThreeCompare(alice, bob)
		break
	case 3: // 两对
		compare = TwoTwiceCompare(alice, bob)
		break
	case 2: // 一对
		compare = TwoOnceCompare(alice, bob)
		break
	case 1: // 单张大牌
		compare = SortCompare(alice, bob)
		break
	}
	return compare
}

/**
同花顺、顺子比较
大于1，小于2，等于0
*/
func StraightCompare(poker1 string, poker2 string) int {
	face := "AKQJT98765432A"
	arr1 := []uint8{poker1[0], poker1[2], poker1[4], poker1[6], poker1[8]}
	arr2 := []uint8{poker2[0], poker2[2], poker2[4], poker2[6], poker2[8]}
	// 排序
	QuickSort(arr1, 0, len(arr1)-1)
	QuickSort(arr2, 0, len(arr2)-1)
	str1 := string(arr1)
	str2 := string(arr2)
	if arr1[0] == 'A' && !IsAKQJT(poker1) {
		str1 = "5432A"
	}
	if arr2[0] == 'A' && !IsAKQJT(poker2) {
		str2 = "5432A"
	}
	index1 := strings.Index(face, str1)
	index2 := strings.Index(face, str2)
	if index1 < index2 {
		return 1
	}
	if index2 < index1 {
		return 2
	}
	return 0
	/*maxNum1 := GetMaxNum(poker1)
	maxNum2 := GetMaxNum(poker2)
	return ComparePoker(maxNum1,maxNum2)*/
}

/**
四条比较大小
*/
func FourOfAKindCompare(poker1 string, poker2 string) int {
	var fourNum1 uint8
	var fourNum2 uint8
	var otherNum1 uint8
	var otherNum2 uint8
	for i := 0; i < len(poker1); i += 2 {
		count1 := strings.Count(poker1, string(poker1[i]))
		count2 := strings.Count(poker2, string(poker2[i]))
		if count1 == 4 {
			fourNum1 = poker1[i]
		}
		if count2 == 4 {
			fourNum2 = poker2[i]
		}
		if count1 == 1 {
			otherNum1 = poker1[i]
		}
		if count2 == 1 {
			otherNum2 = poker2[i]
		}
	}
	compare := ComparePoker(fourNum1, fourNum2)
	if compare == 0 {
		compare = ComparePoker(otherNum1, otherNum2)
	}
	return compare
}

/**
三带二比较大小
*/
func ThreeWithTwoCompare(poker1 string, poker2 string) int {
	var threeNum1 uint8
	var threeNum2 uint8
	var twoNum1 uint8
	var twoNum2 uint8
	for i := 0; i < len(poker1); i += 2 {
		count1 := strings.Count(poker1, string(poker1[i]))
		count2 := strings.Count(poker2, string(poker2[i]))
		if count1 == 3 {
			threeNum1 = poker1[i]
		}
		if count2 == 3 {
			threeNum2 = poker2[i]
		}
		if count1 == 2 {
			twoNum1 = poker1[i]
		}
		if count2 == 2 {
			twoNum2 = poker2[i]
		}
	}
	compare := ComparePoker(threeNum1, threeNum2)
	if compare == 0 {
		compare = ComparePoker(twoNum1, twoNum2)
	}
	return compare
}

/**
同花比较、单张牌
*/
func SortCompare(poker1 string, poker2 string) int {
	arr1 := []uint8{poker1[0], poker1[2], poker1[4], poker1[6], poker1[8]}
	arr2 := []uint8{poker2[0], poker2[2], poker2[4], poker2[6], poker2[8]}
	// 排序
	QuickSort(arr1, 0, len(arr1)-1)
	QuickSort(arr2, 0, len(arr2)-1)
	for i := 0; i < len(arr1); i++ {
		compare := ComparePoker(arr1[i], arr2[i])
		if compare != 0 {
			return compare
		}
	}
	return 0
}

/**
三条比较
*/
func ThreeCompare(poker1 string, poker2 string) int {
	var threeNum1 uint8
	var threeNum2 uint8
	for i := 0; i < len(poker1); i += 2 {
		count1 := strings.Count(poker1, string(poker1[i]))
		count2 := strings.Count(poker2, string(poker2[i]))
		if count1 == 3 {
			threeNum1 = poker1[i]
		}
		if count2 == 3 {
			threeNum2 = poker2[i]
		}
	}
	compare := ComparePoker(threeNum1, threeNum2)
	if compare == 0 {
		compare = SortCompare(poker1, poker2)
	}
	return compare
}

/**
两对比较
*/
func TwoTwiceCompare(poker1 string, poker2 string) int {
	arr1 := []uint8{poker1[0], poker1[2], poker1[4], poker1[6], poker1[8]}
	arr2 := []uint8{poker2[0], poker2[2], poker2[4], poker2[6], poker2[8]}
	var otherNum1 uint8
	var otherNum2 uint8
	for i := 0; i < len(poker1); i += 2 {
		count1 := strings.Count(poker1, string(poker1[i]))
		count2 := strings.Count(poker2, string(poker2[i]))
		if count1 == 1 {
			otherNum1 = poker1[i]
		}
		if count2 == 1 {
			otherNum2 = poker2[i]
		}
	}
	// 排序
	QuickSort(arr1, 0, len(arr1)-1)
	QuickSort(arr2, 0, len(arr2)-1)
	compare := 0
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); i, j = i+1, j+1 {
		if arr1[i] == otherNum1 {
			i++
		}
		if arr2[j] == otherNum2 {
			j++
		}
		if i == len(arr1) || j == len(arr2) {
			break
		}
		compare = ComparePoker(arr1[i], arr2[j])
		if compare != 0 {
			return compare
		}
	}
	if compare == 0 {
		compare = ComparePoker(otherNum1, otherNum2)
	}
	return compare
}

/**
一对比较
*/
func TwoOnceCompare(poker1 string, poker2 string) int {
	var twoNum1 uint8
	var twoNum2 uint8
	for i := 0; i < len(poker1); i += 2 {
		count1 := strings.Count(poker1, string(poker1[i]))
		count2 := strings.Count(poker2, string(poker2[i]))
		if count1 == 2 {
			twoNum1 = poker1[i]
		}
		if count2 == 2 {
			twoNum2 = poker2[i]
		}
	}
	compare := ComparePoker(twoNum1, twoNum2)
	if compare == 0 {
		compare = SortCompare(poker1, poker2)
	}
	return compare
}

/**
字符版
快速排序法排序扑克牌,从大到小排序
*/
func QuickSort(arr []uint8, l int, r int) {
	if l >= r {
		return
	}
	x := arr[l]
	i := l
	j := r
	for i < j {
		for i < j && ComparePoker(x, arr[j]) == 1 {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && ComparePoker(arr[i], x) == 1 {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = x
	QuickSort(arr, l, i-1)
	QuickSort(arr, i+1, r)
	return
}

/**
比较两张牌的大小
*/
func ComparePoker(c1 uint8, c2 uint8) int {
	face := "23456789TJQKA"
	index1 := strings.Index(face, string(c1))
	index2 := strings.Index(face, string(c2))
	if index1 > index2 {
		return 1
	}
	if index1 < index2 {
		return 2
	}
	return 0
}

/**
获取一对牌的最大牌
大于1，小于2，等于0
*/
func GetMaxNum(poker string) uint8 {
	face := "23456789TJQKA"
	maxIndex := 0
	maxNum := poker[0]
	for i := 0; i < len(poker); i += 2 {
		tempIndex := strings.Index(face, string(poker[i]))
		if tempIndex > maxIndex {
			maxIndex = tempIndex
			maxNum = poker[i]
		}
	}
	return maxNum
}

package util

import (
	"strings"
)

func SevenPokerCompare(alice string, bob string) int {
	aliceMaxPoker, alicePriority := FindMaxPokerPriority(alice)
	bobMaxPoker, bobPriority := FindMaxPokerPriority(bob)
	return PokerCompare(aliceMaxPoker, alicePriority, bobMaxPoker, bobPriority)
}

/**
获取优先级
10.皇家同花顺
9.同花顺
8.四条
7.满堂彩(三带二)
6.同花
5.顺子
4.三条
3.两对
2.一对
1.单张大牌
*/
// 获取最大poker牌型以及优先级
func FindMaxPokerPriority(poker string) (string, int) {
	maxPoker := ""
	isContainSameColor, color := IsContainSameColor(poker)
	isStraight := false // 用来判断是否是顺子
	// 判断是否含有同花
	if isContainSameColor {
		// 获取同花色的扑克牌
		colorPoker := FindSameColor(poker, color)
		isStraight = IsStraight(colorPoker)
		// 判断是否是顺子
		if isStraight {
			// 在同花中获取最大顺子
			colorStraight := GetMaxStraight(colorPoker)
			// 判断是否是皇室同花顺
			if IsAKQJT(colorStraight) {
				return colorStraight, 10
			}
			// 普通同花顺
			return colorStraight, 9
		}
		// 获取最大同花
		maxPoker = GetMaxPoker(colorPoker, 5) // 保留
	}
	// 获取最大相同数
	maxCount := GetMaxSameCount(poker)
	// 如果是4条
	if maxCount == 4 {
		fourPoker := GetSameNumPokerBySort(poker, 4)[0]
		nums := GetNumsMaxSinglePoker(poker, 1, fourPoker)
		maxPoker = fourPoker + nums
		return maxPoker, 8
	}
	// 如果是三条
	if maxCount == 3 {
		threePokers := GetSameNumPokerBySort(poker, 3)
		if len(threePokers) > 1 { // 如果有两个三条 ，取最大的三条和第二的三条的任意两个
			maxPoker = threePokers[0] + threePokers[1][:len(threePokers[1])-2]
			return maxPoker, 7
		} else {
			twoPokers := GetSameNumPokerBySort(poker, 2) // 获取所有的两对
			// 判断是否是满堂彩
			if len(twoPokers) > 0 {
				maxPoker = threePokers[0] + twoPokers[0]
				return maxPoker, 7
			}
		}
		// 判断是否保留最大同花
		if len(maxPoker) > 0 {
			return maxPoker, 6
		}
		// 保留 三条
		threePoker := GetSameNumPokerBySort(poker, 3)[0]
		maxPoker = threePoker + GetNumsMaxSinglePoker(poker, 2, threePoker)
	} else {
		// 判断是否保留最大同花
		if len(maxPoker) > 0 {
			return maxPoker, 6
		}
	}
	isStraight = IsStraight(poker)
	// 判断是否是顺子
	if isStraight {
		// 获取最大顺子
		maxPoker = GetMaxStraight(poker)
		return maxPoker, 5
	}
	// 判断是否保留三条
	if len(maxPoker) > 0 {
		return maxPoker, 4
	}
	// 判度安是否有 "对"
	if maxCount == 2 {
		twoPokers := GetSameNumPokerBySort(poker, 2) // 获取所有的"对"
		// 判断是否有两对
		if len(twoPokers) >= 2 {
			maxPoker = twoPokers[0] + twoPokers[1] + GetNumsMaxSinglePoker(poker, 1, twoPokers[0], twoPokers[1])
			return maxPoker, 3
		}
		// 一对
		maxPoker = twoPokers[0] + GetNumsMaxSinglePoker(poker, 3, twoPokers[0])
		return maxPoker, 2
	}
	// 散牌
	maxPoker = GetNumsMaxSinglePoker(poker, 5)
	return maxPoker, 1
}

// 判断是否包含同花,并且返回相同的花色
func IsContainSameColor(poker string) (bool, uint8) {
	colors := "csdh"
	for i := 0; i < len(colors); i++ {
		color := colors[i]
		if strings.Count(poker, string(color)) >= 5 {
			return true, color
		}
	}
	return false, 0
}

// 找出同花的扑克牌
func FindSameColor(poker string, color uint8) string {
	var builder strings.Builder
	for i := 1; i < len(poker); i++ {
		if color == poker[i] {
			builder.WriteString(string(poker[i-1]))
			builder.WriteString(string(poker[i]))
		}
	}
	return builder.String()
}

// 已知是含有顺子，获取最大顺子,并且排序
func GetMaxStraight(poker string) string {
	face := "AKQJT98765432A"
	count := 0
	index := 0
	// 循环查找顺子中的最大值的下标index
	for i := 0; i <= len(face)-1; i++ {
		if count == 4 {
			index = i - count
			break
		}
		if strings.Contains(poker, string(face[i])) {
			count++
			continue
		}
		count = 0
	}
	var builder strings.Builder
	for i := index; i < index+5; i++ {
		tempIndex := strings.Index(poker, string(face[i]))
		builder.WriteString(string(poker[tempIndex]))
		builder.WriteString(string(poker[tempIndex+1]))
	}
	return builder.String()
}

/**
从大到小，获取poker中牌面最大的num张牌
*/
func GetMaxPoker(poker string, num int) string {
	var arr []string                     // 用来保存一张牌，包括牌的大小、颜色
	for i := 0; i < len(poker); i += 2 { // 分割每张牌到arr
		arr = append(arr, poker[i:i+2])
	}
	// 从大到小排序扑克牌
	QuickSortWithString(arr, 0, len(arr)-1)
	var builder strings.Builder
	// 获取前5张
	for i := 0; i < num; i++ {
		builder.WriteString(arr[i])
	}
	return builder.String()
}

/**
字符串版，及传入整张牌的数组
快速排序法排序扑克牌,从大到小排序
*/
func QuickSortWithString(arr []string, l int, r int) {
	if l >= r {
		return
	}
	x := arr[l]
	i := l
	j := r
	for i < j {
		for i < j && ComparePoker(x[0], arr[j][0]) == 1 {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && ComparePoker(arr[i][0], x[0]) == 1 {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = x
	QuickSortWithString(arr, l, i-1)
	QuickSortWithString(arr, i+1, r)
	return
}

/**
输入需要的附属牌个数,以及需要排除的牌
返回poker中num个最大附属牌，排除exclude
*/
func GetNumsMaxSinglePoker(poker string, num int, exclude ...string) string {
	var arr []string
	// 获取已经排除exclude后的扑克牌数组
	for i := 0; i < len(poker); i += 2 {
		flag := true
		for j := 0; j < len(exclude); j++ {
			if poker[i] == exclude[j][0] {
				flag = false
				break
			}
		}
		if flag {
			arr = append(arr, poker[i:i+2])
		}
	}
	// 从大到小排序
	QuickSortWithString(arr, 0, len(arr)-1)
	var builder strings.Builder
	for i := 0; i < num; i++ {
		builder.WriteString(arr[i])
	}
	return builder.String()
}

/**
获取相同数为count的字符串数组
返回相同牌由大到小排序
*/
func GetSameNumPokerBySort(poker string, count int) []string {
	var arr []string
	for i := 0; i < len(poker); i += 2 {
		if strings.Count(poker, string(poker[i])) == count {
			arr = append(arr, poker[i:i+2])
		}
	}
	// 排序
	QuickSortWithString(arr, 0, len(arr)-1)
	var arr2 []string
	for i := 0; i < len(arr); i += count {
		var builder strings.Builder
		for j := i; j < i+count; j++ {
			builder.WriteString(arr[j])
		}
		arr2 = append(arr2, builder.String())
	}
	return arr2
}

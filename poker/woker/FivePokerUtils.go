package woker

import "strings"

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
func GetPriority(poker string) int {
	isSameColor, _ := IsContainSameColor(poker)
	isStraight := IsStraight(poker)
	if isSameColor && isStraight {
		if IsAKQJT(poker) {
			return 10
		}
		return 9
	}
	if isSameColor {
		return 6
	}
	if isStraight {
		return 5
	}
	maxCount := GetMaxSameCount(poker) // 获取最大相同牌数
	if maxCount == 4 {
		return 8
	}
	if maxCount == 3 {
		if IsThreeAndTwo(poker) {
			return 7
		}
		return 4
	}
	if maxCount == 2 {
		if IsTwoTwice(poker) {
			return 3
		}
		return 2
	}
	return 1
}

/**
判断是否同花
已过期，请使用七张牌兼用版SevenPokerUtils.IsContainSameColor
*/
func IsSameColor(poker string) bool {
	if strings.Count(poker, "c") == 5 ||
		strings.Count(poker, "s") == 5 ||
		strings.Count(poker, "d") == 5 ||
		strings.Count(poker, "h") == 5 {
		return true
	}
	return false
}

// 判断是否是顺子
func IsStraight(poker string) bool {
	face := "A23456789TJQKA"
	count := 0
	for i := 0; i < len(face); i++ {
		if count == 5 {
			return true
		}
		if strings.Contains(poker, string(face[i])) {
			count++
			continue
		}
		count = 0
	}
	if count == 5 {
		return true
	}
	return false
}

// 判断是否是AKQJT
func IsAKQJT(poker string) bool {
	if strings.Contains(poker, "A") &&
		strings.Contains(poker, "K") &&
		strings.Contains(poker, "Q") &&
		strings.Contains(poker, "J") &&
		strings.Contains(poker, "T") {
		return true
	}
	return false
}

// 判断最大相同数,并返回相同数的值
func GetMaxSameCount(poker string) int {
	max := 1
	for i := 0; i < len(poker); i += 2 {
		temp := strings.Count(poker, string(poker[i]))
		if temp > max {
			max = temp
		}
	}
	return max
}

// 在已知是同三的情况下,判断是否是满堂彩(三带二)
func IsThreeAndTwo(poker string) bool {
	for i := 0; i < len(poker); i += 2 {
		if strings.Count(poker, string(poker[i])) == 2 {
			return true
		}
	}
	return false
}

// 在已知是同二的情况下，判断是否是两对
func IsTwoTwice(poker string) bool {
	var c uint8
	for i := 0; i < len(poker); i += 2 {
		count := strings.Count(poker, string(poker[i]))
		if count == 2 {
			if c == 0 {
				c = poker[i]
			}
			if c != 0 && c != poker[i] {
				return true
			}
		}
	}
	return false
}

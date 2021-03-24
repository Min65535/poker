package woker

import (
	"strings"
)

func SevenPokerWithGhostCompare(alice string, bob string) int {
	aliceMaxPoker, alicePriority := FindMaxPokerPriority(alice)
	bobMaxPoker, bobPriority := FindMaxPokerPriority(bob)
	if HaveRazz(alice) {
		aliceMaxPoker, alicePriority = GetGhostMaxPoker(alice, aliceMaxPoker, alicePriority)
	}
	if HaveRazz(bob) {
		bobMaxPoker, bobPriority = GetGhostMaxPoker(bob, bobMaxPoker, bobPriority)
	}
	return PokerCompare(aliceMaxPoker, alicePriority, bobMaxPoker, bobPriority)
}

/**
参数：原扑克牌+原最大牌+原最大优先级
获取+赖子的最大牌以及优先级
*/
func GetGhostMaxPoker(poker string, maxPoker string, priority int) (string, int) {
	moreMaxPoker := ""
	switch priority {
	case 10:
		// 皇室不用获取最优牌
		break
	case 9:
		// 同花顺
		// 获取同花顺
		moreMaxPoker, priority = SameColorStraightRazz(maxPoker, priority)
		break
	case 8:
		// 四条
		moreMaxPoker, priority = FourRazz(maxPoker, priority)
		break
	case 7:
		// 三带二
		moreMaxPoker, priority = ThreeAndTwoRazz(maxPoker, priority)
		break
	case 6:
		// 同花
		moreMaxPoker, priority = SameColorRazz(poker, maxPoker, priority)
		break
	case 5:
		// 顺子
		moreMaxPoker, priority = StraightRazz(poker, maxPoker, priority)
		break
	case 4:
		// 三条
		moreMaxPoker, priority = ThreeRazz(poker, maxPoker, priority)
		break
	case 3:
		// 两对
		moreMaxPoker, priority = TwoTwiceRazz(poker, maxPoker, priority)
		break
	case 2:
		// 一对
		moreMaxPoker, priority = TwoRazz(poker, maxPoker, priority)
		break
	case 1:
		moreMaxPoker, priority = HighPoker(poker, maxPoker, priority)
		break
	}
	return moreMaxPoker, priority
}

// 同花顺+赖子
// TODO 未测试
func SameColorStraightRazz(maxPoker string, priority int) (string, int) {
	moreMaxPoker := GetMaxSameColorStraight(maxPoker, GetColor(maxPoker, 5)[0])
	if IsAKQJT(moreMaxPoker) {
		priority = 10
	}
	return moreMaxPoker, priority
}

// 四条+赖子
// TODO 未测试
func FourRazz(maxPoker string, priority int) (string, int) {
	// 获取最大附属牌
	sameColorPoker := GetSameNumPokerBySort(maxPoker, 4)[0]
	moreMaxPoker := sameColorPoker + GetMaxSinglePokerByRazz(0, sameColorPoker)
	return moreMaxPoker, priority
}

// TODO 未测试
// 三带二+赖子
func ThreeAndTwoRazz(maxPoker string, priority int) (string, int) {
	moreMaxPoker := maxPoker
	// 转变成4条
	threePoker := GetSameNumPokerBySort(maxPoker, 3)[0]
	moreMaxPoker = GetSameNumExtension(threePoker) + GetNumsMaxSinglePoker(maxPoker, 1, threePoker)
	priority = 8
	return moreMaxPoker, priority
}

// 同花+癞子
// TODO 未测试
func SameColorRazz(poker string, maxPoker string, priority int) (string, int) {
	flag, moreMaxPoker, maxPriority := ToBeSameStraight(maxPoker)
	if flag {
		return moreMaxPoker, maxPriority
	}
	// 不具备同花顺的条件，判断是否具备三同
	if GetMaxSameCount(poker) == 3 {
		moreMaxPoker, maxPriority = ThreeAndTwoRazz(poker, priority)
		return moreMaxPoker, maxPriority
	}
	// 不具备同花顺条件并且不拥有三条
	// 扩展最大花色
	maxPoker = GetMaxSameColorPokerByRazz(maxPoker)
	return maxPoker, priority
}

// 顺子+癞子
func StraightRazz(poker string, maxPoker string, priority int) (string, int) {
	// 1.判断顺子是否可以变成同花顺
	if flag, morePoker, morePriority := ToBeSameStraight(maxPoker); flag {
		return morePoker, morePriority
	}
	// 2.判断所有牌是否有4同花
	arr := GetColor(maxPoker, 4)
	if len(arr) > 0 {
		// 找出同花的扑克牌
		colorPoker := FindSameColor(poker, arr[0])
		maxPoker = GetMaxSameColorPokerByRazz(colorPoker)
		priority = 6
		return maxPoker, priority
	}
	// 3.获取最大顺子
	// 判断是否有4顺子
	_, extentNum := IsToBeStraight(poker)
	// 补充为最大顺子
	maxPoker = GetMaxStraight(poker + string(extentNum) + "d")
	return maxPoker, priority
}

/**
三条+癞子
*/
func ThreeRazz(poker string, maxPoker string, priority int) (string, int) {
	// 1.判断顺子是否可以变成同花顺
	if flag, morePoker, morePriority := ToBeSameStraight(poker); flag {
		return morePoker, morePriority
	}
	// 2.三条变四条
	arr := GetSameNumPokerBySort(maxPoker, 3)
	fourPoker := GetSameNumExtension(arr[0])
	maxPoker = fourPoker + GetNumsMaxSinglePoker(maxPoker, 1, fourPoker[0:2], fourPoker[2:4], fourPoker[4:6], fourPoker[6:8])
	priority = 8
	return maxPoker, priority
}

/**
两对+癞子
*/
func TwoTwiceRazz(poker string, maxPoker string, priority int) (string, int) {
	// 1.判断顺子是否可以变成同花顺
	if flag, morePoker, morePriority := ToBeSameStraight(poker); flag {
		return morePoker, morePriority
	}
	// 2.变成三带二
	arr := GetSameNumPokerBySort(maxPoker, 2)
	threePoker := GetSameNumExtension(arr[0])
	maxPoker = threePoker + arr[1]
	priority = 7
	return maxPoker, priority
}

/**
一对+癞子
*/
func TwoRazz(poker string, maxPoker string, priority int) (string, int) {
	arr := GetColor(poker, 4)
	// 1.判断是否有4同花
	if len(arr) > 0 {
		// 找出同花的扑克牌
		colorPoker := FindSameColor(poker, arr[0])
		// 2.判断是否是同花顺
		if flag, extentNum := IsToBeStraight(colorPoker); flag {
			// 补充为最大顺子
			maxPoker = GetMaxStraight(string(extentNum) + string(colorPoker[1]) + colorPoker)
			if IsAKQJT(maxPoker) { // 判断是否是皇室同花顺
				return maxPoker, 10
			}
			return maxPoker, 9
		}
		// 3. 不是同花顺，补足同花
		maxPoker = GetMaxSameColorPokerByRazz(colorPoker)
		priority = 6
		return maxPoker, priority
	}
	// 4.不能补全为同花，判断是否可以补全为顺子
	if flag, extentNum := IsToBeStraight(poker); flag {
		// 补充为最大顺子
		maxPoker = GetMaxStraight(string(extentNum) + "s" + poker)
		priority = 5
		return maxPoker, priority
	}
	// 5.最后以上全部不可以，则补全为三条
	strArr := GetSameNumPokerBySort(poker, 2)
	threePoker := GetSameNumExtension(strArr[0])
	twoMaxSinglePorker := GetNumsMaxSinglePoker(poker, 2, threePoker[0:2], threePoker[2:4], threePoker[4:6])
	maxPoker = threePoker + twoMaxSinglePorker
	priority = 4
	return maxPoker, priority
}

/**
一对+癞子
*/
func HighPoker(poker string, maxPoker string, priority int) (string, int) {
	arr := GetColor(poker, 4)
	// 1.判断是否有4同花
	if len(arr) > 0 {
		// 找出同花的扑克牌
		colorPoker := FindSameColor(poker, arr[0])
		// 2.判断是否是同花顺
		if flag, extentNum := IsToBeStraight(colorPoker); flag {
			// 补充为最大顺子
			maxPoker = GetMaxStraight(string(extentNum) + string(colorPoker[1]) + colorPoker)
			if IsAKQJT(maxPoker) { // 判断是否是皇室同花顺
				return maxPoker, 10
			}
			return maxPoker, 9
		}
		// 3. 不是同花顺，补足同花
		maxPoker = GetMaxSameColorPokerByRazz(colorPoker)
		priority = 6
		return maxPoker, priority
	}
	// 4.不能补全为同花，判断是否可以补全为顺子
	if flag, extentNum := IsToBeStraight(poker); flag {
		// 补充为最大顺子
		maxPoker = GetMaxStraight(string(extentNum) + "s" + poker)
		priority = 5
		return maxPoker, priority
	}
	// 5.凑齐一对
	fourSinglePoker := GetNumsMaxSinglePoker(maxPoker, 4)
	maxPoker = GetSameNumExtension(fourSinglePoker[0:2]) + fourSinglePoker[2:]
	priority = 2
	return maxPoker, priority
}

/*
	输入4~5张同花，获取最大同花
*/
func GetMaxSameColorPokerByRazz(colorPoker string) string {
	var arr []string
	for i := 0; i < len(colorPoker); i += 2 {
		arr = append(arr, colorPoker[i:i+2])
	}
	QuickSortWithString(arr, 0, len(arr)-1)
	// 获取同花最大附属牌
	extendsPoker := GetMaxSinglePokerByRazz(colorPoker[1], arr[0], arr[1], arr[2], arr[3])
	var builder strings.Builder
	builder.WriteString(extendsPoker)
	for i := 0; i < 4; i++ {
		builder.WriteString(arr[i])
	}
	return builder.String()
}

/**
输入：1～3张相同大小的牌，不同花色
将三条相扩展为4条
一对 变成 三条
散牌变成一对
*/
func GetSameNumExtension(poker string) string {
	colors := "shdc"
	// 遍历获取出现的颜色
	var excludeColors []uint8
	for i := 1; i < len(poker); i += 2 {
		excludeColors = append(excludeColors, poker[i])
	}
	var color uint8
	for i := 0; i < len(colors); i++ {
		flag := true
		for j := 0; j < len(excludeColors); j++ {
			if excludeColors[j] == colors[i] {
				flag = false
				break
			}
		}
		if flag {
			color = colors[i]
			break
		}
	}
	return poker + string(poker[0]) + string(color)
}

// 判断是否具备变成同花顺条件
func ToBeSameStraight(poker string) (bool, string, int) {
	moreMaxPoker := poker
	// 获取次数大于等于4的颜色
	arr := GetColor(poker, 4)
	// 判断是否有4同花
	if len(arr) > 0 {
		// 找出同花的扑克牌
		colorPoker := FindSameColor(poker, arr[0])
		// 判断是否有4顺子
		if flag, extentNum := IsToBeStraight(colorPoker); flag {
			// 补充为最大顺子
			moreMaxPoker = GetMaxStraight(string(extentNum) + string(colorPoker[1]) + colorPoker)
			if IsAKQJT(moreMaxPoker) { // 判断是否是皇室同花顺
				return true, moreMaxPoker, 10
			}
			return true, moreMaxPoker, 9
		}
	}
	return false, "", 0
}

// 判断是否具备顺子条件
// 返回缺的牌
func IsToBeStraight(poker string) (bool, uint8) {
	// face := "A23456789TJQKA"
	face := "AKQJT98765432A"
	count := 0
	lackNum := uint8(0)
	// 先便利到5
	i, j := 0, 0
	for j = 0; j < 5; j++ {
		if strings.Contains(poker, face[j:j+1]) {
			count++
		} else {
			lackNum = face[j]
		}
	}
	j--
	if count == 4 {
		return true, lackNum
	}
	for j < len(face)-1 {
		if strings.Contains(poker, face[j+1:j+2]) {
			count++
		} else {
			lackNum = face[j+1]
		}
		if strings.Contains(poker, face[i:i+1]) {
			count--
		}
		i++
		j++
		if count == 4 {
			break
		}
	}
	if count == 4 {
		return true, lackNum
	}
	return false, uint8(0)
}

// 判断是否含有赖子
func HaveRazz(poker string) bool {
	return strings.Contains(poker, "Xn")
}

/**
	获取赖子变成的最大附属牌
 	输入需要排除的牌和返回的附属牌颜色，输出癞子可变的最大的牌
	1. 输入颜色，代表要求的是同花顺的最大附属牌，那么获取的牌必须要和excludes的数不一样
	2. 颜色为0，代表只需要除了excludes之外的最大的牌即可
*/
func GetMaxSinglePokerByRazz(color uint8, excludes ...string) string {
	face := "AKQJT98765432"
	colorFace := "AsAhAdAcKsKhKdKcQsQhQdQcJsJhJdJcTsThTdTc9h9s9d9c8s8h8d8c7s7h7d7c6s6h6d6c5s5h5d5c4s4h4d4c3s3h3d3c2s2h2d2c"
	if color != 0 {
		extendsNum := uint8(0)
		for i := 0; i < len(face); i++ {
			flag := true
			for j := 0; j < len(excludes); j++ {
				if excludes[j][0] == face[i] {
					flag = false
					break
				}
			}
			if flag {
				extendsNum = face[i]
				break
			}
		}
		return string(extendsNum) + string(color)
	} else {
		extendsPoker := ""
		for i := 0; i < len(colorFace); i += 2 {
			flag := true
			for j := 0; j < len(excludes); j++ {
				if excludes[j] == colorFace[i:i+2] {
					flag = false
					break
				}
			}
			if flag {
				extendsPoker = face[i : i+2]
				break
			}
		}
		return extendsPoker
	}
}

// 输入已有5张顺子，将癞子获取更大的5张顺子
func GetMaxSameColorStraight(poker string, color uint8) string {
	face := "AKQJT98765432"
	// 获取现有最大牌面
	maxNum := GetMaxNum(poker)
	index := strings.Index(face, string(maxNum))
	if index == 0 {
		// 已经是最大的了
		return poker
	}
	maxNum = face[index-1]
	maxSinglePoker := string(maxNum) + "s" // 随机颜色
	if color != 0 {                        // 如果约定花色
		maxSinglePoker = string(maxNum) + string(color)
	}
	if len(poker) == 8 {
		// 获取排序的由大到小排序的4张牌
		poker = GetMaxPoker(poker, 4)
		return maxSinglePoker + poker
	}
	poker = GetMaxPoker(poker, 5)
	return maxSinglePoker + poker[:len(poker)-2]
}

// 获取出现大于等于count次数的颜色
func GetColor(poker string, count int) []uint8 {
	var arr []uint8
	for i := 1; i < len(poker); i += 2 {
		if strings.Count(poker, string(poker[i])) >= count {
			flag := true
			if poker[i] == 'n' { // 癞子不算
				continue
			}
			// 判断是否已经存在
			for j := 0; j < len(arr); j++ {
				if poker[i] == arr[j] {
					flag = false
					break
				}
			}
			if flag {
				arr = append(arr, poker[i])
			}
		}
	}
	return arr
}

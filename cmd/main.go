package main

import (
	"fmt"
	"time"
)

func main() {
	SevenWithGhost()
}
func SevenWithGhost() {
	var dataAll model.DataJson
	var errResults []model.Result
	util.ReadFile("pkg/seven_cards_with_ghost.result.json", &dataAll)
	startT := time.Now()
	count := 0 // 对的
	for _, data := range dataAll.Matches {
		myResult := util.SevenPokerWithGhostCompare(data.Alice, data.Bob)
		if myResult == data.Result {
			count++
		} else {
			errData := Result{
				Alice:       data.Alice,
				Bob:         data.Bob,
				Result:      data.Result,
				ErrorResult: myResult,
			}
			errResults = append(errResults, errData)
		}
	}
	tc := time.Since(startT) // 计算耗时
	errDatas := DataJson{
		Matches: errResults,
	}
	util.WriteFile("pkg/seven_cards_with_ghost_err.json", &errDatas)
	rate := float64(count) / float64(len(dataAll.Matches)) * 100
	fmt.Printf("比率是%f", rate)
	fmt.Printf("耗时 = %v\n", tc)
}

// 七张牌
func Seven() {
	var dataAll DataJson
	var errResults []Result
	util.ReadFile("pkg/seven_cards_with_ghost.json", &dataAll)
	startT := time.Now()
	count := 0 // 对的
	for _, data := range dataAll.Matches {
		myResult := util.SevenPokerCompare(data.Alice, data.Bob)
		if myResult == data.Result {
			count++
		} else {
			errData := Result{
				Alice:       data.Alice,
				Bob:         data.Bob,
				Result:      data.Result,
				ErrorResult: myResult,
			}
			errResults = append(errResults, errData)
		}
	}
	tc := time.Since(startT) // 计算耗时
	errDatas := DataJson{
		Matches: errResults,
	}
	util.WriteFile("pkg/seven_cards_err.json", &errDatas)
	rate := float64(count) / float64(len(dataAll.Matches)) * 100
	fmt.Printf("比率是%f", rate)
	fmt.Printf("耗时 = %v\n", tc)
}

// 五张牌
func Five() {
	var dataAll DataJson
	var result DataJson
	var errResults []Result
	util.ReadFile("pkg/match.json", &dataAll)
	util.ReadFile("pkg/match_result.json", &result)
	startT := time.Now()
	count := 0 // 对的
	for index, data := range dataAll.Matches {
		alicePriority := util.GetPriority(data.Alice)
		bobPriority := util.GetPriority(data.Bob)
		data.Result = util.PokerCompare(data.Alice, alicePriority, data.Bob, bobPriority)
		if data.Result == result.Matches[index].Result {
			count++
		} else {
			errData := Result{
				Alice:       data.Alice,
				Bob:         data.Bob,
				Result:      result.Matches[index].Result,
				ErrorResult: data.Result,
			}
			errResults = append(errResults, errData)
		}
	}
	tc := time.Since(startT) // 计算耗时
	errDatas := DataJson{
		Matches: errResults,
	}
	util.WriteFile("pkg/match_err.json", &errDatas)
	rate := float64(count) / float64(len(dataAll.Matches)) * 100
	fmt.Printf("比率是%f", rate)
	fmt.Printf("耗时 = %v\n", tc)
}

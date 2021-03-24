package util

import (
	"TexasPoker/pkg/model"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(filepath string, data *model.DataJson) {
	filePtr, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Open json.file failed:%v\n", err)
		return
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Print("Decode failed", err.Error())
	} else {
		fmt.Print("Decode success")
	}
}

func WriteFile(filepath string, data *model.DataJson) {
	filePtr, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}
}

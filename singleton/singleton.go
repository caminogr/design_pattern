package singleton

import "fmt"

type singleton struct {
}

// インスタンス保持変数
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
		fmt.Println("instanceを生成しました")
	}

	return instance
}

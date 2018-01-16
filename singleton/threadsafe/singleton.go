package singleton

import (
	"fmt"
	"sync"
	"time"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		time.Sleep(5000)
		instance = &singleton{}
		fmt.Println("instanceを生成しました")
	})
	return instance
}

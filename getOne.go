package main

import (
	"log"
	"net/http"
	"sync"
)

var sum int64 = 0

//预存商品数量
var productNum int64 = 10000

// 互斥锁
var mutex sync.Mutex

// 获取秒杀商品
func GetOneProduct() bool {
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	// 判断数据是否超限
	if sum < productNum {
		sum += 1
		return true
	}
	return false

}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
	}
	w.Write([]byte("false"))
}

func main() {
	http.HandleFunc("/getOne", GetProduct)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("Err:", err)
	}

}

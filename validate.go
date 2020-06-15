package main

import (
	"fmt"
	"imooc-product/common"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	// 执行正常业务逻辑
	fmt.Println("执行check！")
}

// 统一验证拦截器， 每个接口都需要提前验证
func Auth(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	//1. 过滤器
	filter := common.NewFilter()
	// 注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	//2. 启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	// 启动服务
	http.ListenAndServe(":8083", nil)

}

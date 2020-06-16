package main

import (
	"errors"
	"fmt"
	"imooc-product/common"
	"imooc-product/encrypt"
	"net/http"
	"sync"
)

var hostArray = []string{"127.0.0.1", "127.0.0.1"}

var localHost = "127.0.0.1"

var port = "8081"

var hashConsistent *common.Consistent

// 用来存放控制信息
type AccessControl struct {
	//用来存放用户想要存放的信息
	sourcesArray map[int]interface{}
	*sync.RWMutex
}

// 创建全局变量
var accessControl = &AccessControl{sourcesArray: make(map[int]interface{})}

// 获取指定的数据

func (m *AccessControl) GetNewRecord(uid int) interface{} {
	m.RWMutex.RLock()
	defer m.RWMutex.Unlock()
	data := m.sourcesArray[uid]
	return data
}

// 设置记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.RWMutex.Lock()
	m.sourcesArray[uid] = "hello imooc"
	m.RWMutex.Unlock()
}

func (m *AccessControl) GetDistributedRight(req *http.Request) bool {

}

// 执行正常业务逻辑
func Check(w http.ResponseWriter, r *http.Request) {
	// 执行正常业务逻辑
	fmt.Println("执行check！")
}

// 统一验证拦截器， 每个接口都需要提前验证
func Auth(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("执行验证！")
	// 添加基于cookie 的权限验证
	err := CheckUserInfo(r)
	if err != nil {
		return err
	}
	return nil
}

// 身份校验函数
func CheckUserInfo(r *http.Request) error {
	// 获取Uid, cookie
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		return errors.New("用户UID Cookie 获取失败！")
	}
	// 获取用户加密串
	signCookie, err := r.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串 Cookie 获取失败！ ")
	}
	// 对信息进行解密
	signByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		return errors.New("加密串已被篡改！")
	}

	fmt.Println("结果比对")
	fmt.Println("用户ID" + uidCookie.Value)
	fmt.Println("解密后用户ID" + uidCookie.Value)
	if checkInfo(uidCookie.Value, string(signByte)) {
		return nil
	}
	return errors.New("身份校验失败！ ")
}

// 自定义逻辑判断
func checkInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}

func main() {
	// 负载均衡器设置
	// 采用一致性哈希算法
	hashConsistent = common.NewConsistent()

	for _, v := range hostArray {
		hashConsistent.Add(v)

	}
	// 采用一致性hash算法
	//1. 过滤器
	filter := common.NewFilter()
	// 注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	//2. 启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	// 启动服务
	http.ListenAndServe(":8083", nil)

}

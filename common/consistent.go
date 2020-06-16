package common

import (
	"errors"
	"hash/crc32"
	"strconv"
	"sync"
)

// 声明新切片类型
type units []uint32

// 返回切片长度
func (x units) Len() int {
	return len(x)
}

// 比对两个数大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

// 切片中两个值的交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// 当hash环上没有数据时， 提示错误
var errEmpty = errors.New("Hash 环没有数据")

//创建结构体保存一致性hash信息
type Consistent struct {
	// hash环， key为哈希值， 值存放节点的信息
	circle map[uint32]string
	// 已经排序的节点hash切片
	sortedHashes units
	//虚拟节点个数，用来增加hash的平衡性
	VirtualNode int
	//map 读写锁
	sync.RWMutex
}

// 创建一致性hash算法结构体, 设置默认节点数量
func NewConsistent() *Consistent {
	return &Consistent{
		// 初始化变量
		circle: make(map[uint32]string),
		// 设置虚拟节点个数
		VirtualNode: 20,
	}
}

// 自动生成key值
func (c *Consistent) generateKey(element string, index int) string {
	//副本key生成逻辑
	return element + strconv.Itoa(index)
}

// 获取hash位置
func (c *Consistent) hashKey(key string) uint32 {
	if len(key) < 64 {
		// 声明一个数组长度为64
		var srcatch [64]byte
		// 拷贝数据到数组中
		copy(srcatch[:], key)
		// 使用IEEE 多项式返回数据的CRC-32校验和
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) Add(element string) {
	// 循环虚拟节点， 设置副本
	for i := 0; i < c.VirtualNode; i++ {
		c.circle[c.hashKey(c.generateKey(element, i))] = element
	}

}

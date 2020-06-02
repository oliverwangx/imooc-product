package repositories

import "imooc-product/datamodels"

type IOrderRepository interface {
	Conn() error
	Insert()
	Delete()
	Update()
	SelectByKey(int64) (*datamodels.Order, error)
}

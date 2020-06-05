package datamodels

type Order struct {
	ID          int64 `sql:"ID" imooc:"ID"`
	UserId      int64 `sql:"userID" imooc:"UserID"`
	ProductId   int64 `sql:"productID" imooc:"ProductID"`
	OrderStatus int   `sql:"orderStatus" imooc:"OrderStatus"`
}

const (
	OrderWait    = iota
	OrderSuccess //1
	OrderFailed  //2
)

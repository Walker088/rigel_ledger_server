package dao

type Dao[T any] interface {
	GetAll()
	GetById(string)
	UpdateById(string)
	Create(T)
}

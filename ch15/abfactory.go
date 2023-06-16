package ch15

import "fmt"

type IDataOperation[T any] interface {
	Insert(int, *T)
	Get(int) *T
}

type User struct {
	Name string
	Id   int
	Age  int
}

// sqlserver database
type SqlserverUserV1[T any] struct {
	database map[int]*T
}

func NewSqlserverUserV1[T any]() *SqlserverUserV1[T] {
	return &SqlserverUserV1[T]{database: make(map[int]*T)}
}

func (this *SqlserverUserV1[T]) Insert(key int, user *T) {
	this.database[key] = user
	println("sqlserver insert")
}

func (this *SqlserverUserV1[T]) Get(id int) *T {
	println("sqlserver get")
	return this.database[id]
}

// access database
type AccessUserV1[T any] struct {
	database map[int]*T
}

func NewAccessUserV1[T any]() *AccessUserV1[T] {
	return &AccessUserV1[T]{database: make(map[int]*T)}
}

func (this *AccessUserV1[T]) Insert(key int, user *T) {
	this.database[key] = user
	println("access insert")
}

func (this *AccessUserV1[T]) Get(id int) *T {
	println("access get")
	return this.database[id]
}

type IFactory[T any] interface {
	Create() IDataOperation[T]
}

// sqlserver factory
type SqlserverFactoryV1[T any] struct {
}

func NewSqlserverFactoryV1[T any]() *SqlserverFactoryV1[T] {
	return &SqlserverFactoryV1[T]{}
}

func (this *SqlserverFactoryV1[T]) Create() IDataOperation[T] {
	return NewSqlserverUserV1[T]()
}

// access factory
type AccessFactoryV1[T any] struct {
}

func NewAccessFactoryV1[T any]() *AccessFactoryV1[T] {
	return &AccessFactoryV1[T]{}
}

func (this *AccessFactoryV1[T]) Create() IDataOperation[T] {
	return NewAccessUserV1[T]()
}

func AbFactoryMainV1() {
	factory := NewSqlserverFactoryV1[User]()
	user := factory.Create()

	u := &User{
		Name: "Jack",
		Age:  12,
		Id:   1,
	}
	user.Insert(1, u)

	ru := user.Get(1)
	fmt.Printf("result:%v\n", ru)
}

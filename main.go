package main

import (
	"go-practics/config"
	dbsvc "go-practics/internal/db"
	"go-practics/internal/imooc"
)

func init() {
	config.New("config/proxy.yaml", "config/components.yaml")
}

func main() {
	// syntax.MyBase()
	// syntax.MyArray()
	// syntax.MyPointer()
	// syntax.MySlice()
	// syntax.MyMap()
	// syntax.MyString()
	// syntax.MyStruct()
	// syntax.MyMethod()
	// syntax.MyInterface()
	// syntax.MyReflect()
	// syntax.MyConcurrent()
	// syntax.MyLock()
	// syntax.Myselect()
	// syntax.MyConcurrent()
	// syntax.MyChannel()
	// syntax.MyDefer()
	// syntax.MyDefer()

	imooc.NewImooc()

	// external.MyViper()

	// kafka.RunProducer()
	// kafka.RunConsumer()
	dbsvc.NewDbServer()

	// r := router.NewRouter()
	// r.Run(":9990")

}

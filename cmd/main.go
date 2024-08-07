package main

import (
	dbsvc "go-practics/internal/db"
	"go-practics/internal/router"
)

func init() {
	// external.NewViper("config/proxy.yaml")
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
	// syntax.Myselect()
	// syntax.MyConcurrent()
	// syntax.MyChannel()

	// external.MyViper()

	// kafka.RunProducer()
	// kafka.RunConsumer()
	dbsvc.NewDbServer()

	r := router.NewRouter()
	r.Run(":9990")

}

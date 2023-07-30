// 结构体
package syntax

import "fmt"

type Person struct {
	name, gender, addr string
	age                int
}
type Student struct {
	id            int32
	person        Person
	school, major string
	class         int
}

var stu = Student{
	id: 1,
	person: Person{
		name:   "Tom",
		gender: "Man",
		addr:   "Chengdu",
		age:    21,
	},
	school: "Peking University",
	major:  "Physical",
	class:  1,
}

func MyStruct() {
	fmt.Printf("stu: %v\n", stu)
	fmt.Printf("UpdateStuInfo(&stu): %v\n", UpdateStuInfo(&stu))
}

func UpdateStuInfo(stu *Student) *Student {
	stu.class = 2
	stu.person.age = 22
	return stu
}

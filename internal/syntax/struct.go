package syntax

import "fmt"

type Person struct {
	name, gender, addr string
	age                int
}

type Student struct {
	id            int32
	class         int
	person        Person
	school, major string
}

var stu = Student{
	person: Person{
		name:   "Tom",
		gender: "Man",
		addr:   "Chengdu",
		age:    21,
	},
	school: "Peking University",
	major:  "Physical",
	id:     1,
	class:  1,
}

func MyStruct() {
	fmt.Printf("stu: %v\n", stu)                                 // stu: {1 {Tom Man Chengdu 21} Peking University Physical 1}
	fmt.Printf("UpdateStuInfo(&stu): %v\n", UpdateStuInfo(&stu)) // UpdateStuInfo(&stu): &{1 {Tom Man Chengdu 22} Peking University Physical 2}
}

func UpdateStuInfo(stu *Student) *Student {
	stu.class = 2
	stu.person.age = 22
	return stu
}

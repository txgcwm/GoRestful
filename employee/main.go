package main

import "fmt"


const (
    Female = iota
    Male = iota
)

type Employee struct {
	id 		int
	sex 	int
	age 	int
	name 	string
	phone	string
	email	string
}

func printEmployee(staff Employee) {
   fmt.Printf("Id: %d\n", staff.id)
   fmt.Printf("Sex: %d\n", staff.sex)
   fmt.Printf("Age: %d\n", staff.age)
   fmt.Printf("Name: %s\n", staff.name)
   fmt.Printf("Phone: %s\n", staff.phone)
   fmt.Printf("Name: %s\n", staff.email)
   fmt.Printf("\n")
}

func printPointerEmployee(staff *Employee) {
   fmt.Printf("Id: %d\n", staff.id)
   fmt.Printf("Sex: %d\n", staff.sex)
   fmt.Printf("Age: %d\n", staff.age)
   fmt.Printf("Name: %s\n", staff.name)
   fmt.Printf("Phone: %s\n", staff.phone)
   fmt.Printf("Name: %s\n", staff.email)
   fmt.Printf("\n")
}

func initEmployee(id int, sex int, age int, name string, phone string, email string) Employee {
	var staff Employee

	staff.id = id
	staff.sex = sex
	staff.age = age
	staff.name = name
	staff.phone = phone
	staff.email = email

	return staff
}

func (staff *Employee) init_staff_info(id int, sex int, age int, name string, phone string, email string) {
	staff.id = id
	staff.sex = sex
	staff.age = age
	staff.name = name
	staff.phone = phone
	staff.email = email
}

func main() {
	var staff0, staff1, staff3, staff4 Employee
	var staffMap map[int]Employee = make(map[int]Employee)

	staff0 = initEmployee(0, Male, 29, "Tom", "13067553598", "haha@163.com")
	staff1 = initEmployee(1, Female, 25, "Till", "13067553598", "test@gmail.com")
	staff3 = initEmployee(2, Male, 19, "Sam", "13067553598", "haha@gmail.com")
	staff4.init_staff_info(3, Male, 27, "David", "13067553598", "haha@gmail.com")

	staffMap[staff0.id] = staff0
	staffMap[staff1.id] = staff1
	staffMap[staff3.id] = staff3
	staffMap[staff4.id] = staff4

   	for id := range staffMap {
   		printEmployee(staffMap[id])
   	}

   	staff, ok := staffMap[1]
   	if(ok){
      	fmt.Println("Name is ", staff.name)  
   	} else {
      	fmt.Println("not present") 
   	}
}



package main

import "fmt"

func main() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)

	//Adding element in map
	employeeSalary["Akshith"] = 12000
	employeeSalary["x"] = 15000
	employeeSalary["a"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	println()
	employeeSalary2 := map[string]int{
		"steven": 13000,
		"james":  17000,
	}
	employeeSalary2["michael"] = 7000
	fmt.Println("employeeSalary map contents:", employeeSalary2)

	println()

	println("Retrieving value for a key from a map")
	employee := "Akshith"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	println("this print 0")
	fmt.Println("Salary of joe is", employeeSalary["joe"])

	println()
	println("Checking if a key exists")
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	println()
	println("Iterate over all elements in a map")
	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	println()
	println("Deleting items from a map")
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)
}

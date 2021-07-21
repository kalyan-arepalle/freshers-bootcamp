package main

import "fmt"

const noOfDaysinMonth = 28

type Salary interface {
	calculateSalary() int
}

type Fulltime struct{
	basic int
	noOfDays int
}

func (f Fulltime) calculateSalary() int{
	return f.noOfDays*f.basic
}

type Contract struct{
	basic int
	noOfDays int
}

func (c Contract) calculateSalary() int{
	return c.noOfDays*c.basic
}

type Freelancer struct{
	basic int
	noOfHrs int
}

func (f Freelancer) calculateSalary() int{
	return f.noOfHrs*f.basic
}

func main() {
	fullTime := Fulltime{500,noOfDaysinMonth}
	fmt.Println(fullTime.calculateSalary())

	contract := Contract{100, noOfDaysinMonth}
	fmt.Println(contract.calculateSalary())

	freelancer := Freelancer{10,200}
	fmt.Println(freelancer.calculateSalary())
}

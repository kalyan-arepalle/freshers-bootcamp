package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	Rows    int     `json:"Rows"`
	Columns int     `json:"Columns"`
	Val     [][]int `json:"Val"`
}

func (m Matrix) getRows() int{
	return m.Rows
}

func (m Matrix) getColumns() int{
	return m.Columns
}

func (m Matrix) setVal(i,j,k int) {
	m.Val[i][j] = k
}

func (m Matrix) addTwo(n Matrix) Matrix{
	sum :=  Matrix{Rows: m.Rows, Columns: m.Columns}
	sum.Val = make([][]int, sum.Rows)
	for i:=0;i<sum.Rows;i++{
		sum.Val[i] = make([]int,sum.Columns)
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			sum.Val[i][j] = m.Val[i][j] + n.Val[i][j]
		}
	}
	return sum
}

func (m Matrix) printMat(){
	data , err := json.MarshalIndent(&m,""," ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func main() {

	r1,c1 :=  3,3
	val1 := make([][]int, r1)
	for i:=0;i<r1;i++{
		val1[i] = make([]int,c1)
	}
	mat1 := Matrix{r1,c1,val1}

	fmt.Println("Number of Rows: ", mat1.getRows())
	fmt.Println("Number of Columns: ", mat1.getColumns())

	mat1.setVal(0,0,3)
	mat1.setVal(0,2,1)
	mat1.setVal(2,1,5)

	mat2 := Matrix{
		Rows:    3,
		Columns: 3,
		Val: [][]int {{1,2,3},
			{5,6,7},
			{9,10,11}},
	}

	fmt.Printf("%+v\n",mat1.addTwo(mat2))

	mat1.printMat()
}


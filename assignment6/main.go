package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)


type Expression struct {
	Operation string `json:"Operation"`
	Units     Units  `json:"Units"`
}
type Units struct {
	One int `json:"one"`
	Two int `json:"two"`
}


func main() {	
	info, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var theInfo []Expression
	error1 := json.Unmarshal([]byte(info), &theInfo)
	if error1 != nil{
		fmt.Println(error1)
	}

	fileAdd, err := os.OpenFile("info.json", os.O_APPEND|os.O_WRONLY, 0644)
	for _,j := range theInfo {
		operations := j.Operation
		switch operations{
		case "Add":
			sum := j.Units.One + j.Units.Two
			fileAdd.WriteString("Addition: " + strconv.Itoa(sum) + "\n")
			break
		case "Subtract":
			difference := j.Units.One - j.Units.Two
			fileAdd.WriteString("Subtraction: " + strconv.Itoa(difference) + "\n")
			break
		case "Multiply":
			product := j.Units.One * j.Units.Two
			fileAdd.WriteString("Multplication: " + strconv.Itoa(product) + "\n")
			break
		case "Divide":
			quotient := j.Units.One / j.Units.Two
			fileAdd.WriteString("Division: " + strconv.Itoa(quotient) + "\n")
			break
		default:
			fmt.Println("Not a valid operation")
		}
	}
	fileAdd.Close()
}
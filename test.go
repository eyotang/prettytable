package main

import (
	"./prettytable"
)

var table = [][]string{
	[]string{ "Anthony", "22", "red"},
	[]string{ "Chris", "23", "blonde"},
	[]string{ "Kevin", "23", "black"},
	[]string{ "Bharat", "23", "black"},
}

var labels = []string { "姓名", "年龄", "头发颜色" }


func main() {
	prettytable.PrintTable(labels, table)
}

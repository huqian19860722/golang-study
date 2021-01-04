package main

import (
	"fmt"
)

func main() {
	key := "" //用户输入选项
	loop := true //控制是否退出循环
	for {
		fmt.Println("-------------------家庭收支记账软件-------------------")
		fmt.Println("-------------------  1 收支明细")
		fmt.Println("-------------------  2 登记收入")
		fmt.Println("-------------------  3 登记支出")
		fmt.Println("-------------------  4 退出软件")
		fmt.Print("请选择(1-4): ")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("-------------------当前收支明细记录-------------------")
		case "2":
		case "3":
			fmt.Println("登记支出..")
		case "4":
			loop = false
		default:
			fmt.Println("请输入正确的选项..")
		}

		if !loop {
			break
		}
	}

	fmt.Println("你退出家庭记账软件的使用...")

}

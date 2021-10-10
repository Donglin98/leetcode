package main

import "fmt"

type ValNode struct {
	row,col,val int
}

func main(){
	var ewcharr [11][11]int
	ewcharr[1][2]=1
	ewcharr[5][6]=2
	for _, v:=range  ewcharr{
		for _,v1 :=range v{
			fmt.Printf("%d\t",v1)
		}
		fmt.Println()
	}
	// 3. 转成稀疏数组
	// 思路：
	// （1）遍历 chessMap，如果我们发现有一个元素的值不为0，我们就创建一个node结构体
	// （2）将其放入到对应的切片即可
	var sparseArr []ValNode

	// 标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模（行和列，默认值）
	// 创建一个ValNode值结点
	valNode := ValNode {
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr,valNode)
	for i1, v1 := range ewcharr {
		for i2, v2 := range v1 {
			if v2 != 0 {
				// 创建一个ValNode  值结点
				valNode = ValNode {
					row: i1,
					col: i2,
					val: v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}
	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d：%d %d %d \n",i,valNode.row,valNode.col,valNode.val)
	}
	// 将这个稀疏数组，存盘 d:/chessmap.data
	// 如果恢复原始的数组
	// 1. 打开这个存盘的文件：d:/chessmap.data => 恢复原始数组
	// 2. 这里使用稀疏数组恢复
	// 先创建一个原始数组
	var chessMap2 [11][11]int
	// 遍历稀疏数组 sparseArr [遍历文件的每一行]
	for index, valNode := range sparseArr {
		if (index == 0) {
			continue
		}
		chessMap2[valNode.row][valNode.col] = valNode.val
	}
	// 看看chessMap2看看是不是恢复了
	for _,value1 := range chessMap2 {
		for _, value2 := range value1 {
			fmt.Printf("%d \t",value2)
		}
		fmt.Println()
	}
}


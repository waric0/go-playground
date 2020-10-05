package main

import (
	"fmt"
	"time"
)

/*

// forを使ったもの
func loop(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			for k := 0; k < num; k++ {
				time.Sleep(time.Millisecond * 100)
				fmt.Printf("%d %d %d \n", i, j, k)
			}
		}
	}
}

*/

func runLoop() {

	num := 3      // 列の数
	colIndex := 0 // 列のインデックス
	list := make([]int, num, num)
	recursiveLoop(num, &colIndex, &list)
	// loop(num)
}

// 再帰を使ったもの
func recursiveLoop(num int, colIndex *int, list *[]int) {

	for i := 0; i < num; i++ {
		tmp := *list
		tmp[*colIndex] = i
		list = &tmp

		if *colIndex == num-1 {
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("%v \n", *list)
		} else {
			*colIndex = *colIndex + 1
			recursiveLoop(num, colIndex, list)
		}
		if i == num-1 {
			*colIndex = *colIndex - 1
		}
	}

}

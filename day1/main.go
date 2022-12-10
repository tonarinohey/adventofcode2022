package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 64)
	var tmp []byte
	tmpSum := 0
	blockNum := 0
	sumColsTopThree := 0
	var cols []int
	for {
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)

		// 空文字の行検出したらリセット
		if string(tmp) == "" {
			blockNum = blockNum + 1
			cols = append(cols, tmpSum)
			tmpSum = 0
		} else {
			i, _ := strconv.Atoi(string(tmp))
			tmpSum = tmpSum + i
		}

		if !isPrefix {
			tmp = nil
		}
	}

	// colsをソートする
	sort.Sort(sort.Reverse(sort.IntSlice(cols)))
	println(cols)

	// 逆順にsliceをイテレートし最大3要素を取得する
	for i := 0; i <= 2; i++ {
		fmt.Println(cols[i])
		sumColsTopThree = sumColsTopThree + cols[i]
	}

	fmt.Printf("最大カロリー: %d \n", cols[0])
	fmt.Printf("最大カロリーTOP3の和: %d \n", sumColsTopThree)
}

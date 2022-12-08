package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 別に各エルフが運ぶ食品カロリーの総量を再利用したりしないので、slice利用したりせずただ数え上げるだけ。
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
	heaviestCol := 0
	heaviestBlockNum := 0
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
			if heaviestCol < tmpSum {
				heaviestCol = tmpSum
				heaviestBlockNum = blockNum
			}
			tmpSum = 0
		} else {
			i, _ := strconv.Atoi(string(tmp))
			tmpSum = tmpSum + i
		}

		if !isPrefix {
			tmp = nil
		}
	}
	println(heaviestBlockNum)
	fmt.Printf("最大カロリーを抱えるエルフの番号: %d \n", heaviestBlockNum)
	fmt.Printf("最大カロリー: %d \n", heaviestCol)
}

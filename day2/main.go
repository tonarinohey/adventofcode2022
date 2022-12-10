package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//
func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 64)
	var tmp []byte
	tmpSum := 0
	sum := 0

	// 何の手を選んだか　と　勝敗結果　がスコアになる
	for {
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)

		// splitする
		arr := strings.Split(string(tmp), "")

		// 1. 自分のだした手をスコア化
		// X for Rock, Y for Paper, and Z for Scissors
		// 1 for Rock, 2 for Paper, and 3 for Scissors
		// A for Rock, B for Paper, and C for Scissors
		// 0 if you lost, 3 if the round was a draw, and 6 if you won
		if arr[2] == "X" {
			tmpSum = 1
			if arr[0] == "A" {
				tmpSum = tmpSum + 3
			} else if arr[0] == "B" {
				// no add
			} else {
				tmpSum = tmpSum + 6
			}
		} else if arr[2] == "Y" {
			tmpSum = 2
			if arr[0] == "A" {
				tmpSum = tmpSum + 6
			} else if arr[0] == "B" {
				tmpSum = tmpSum + 3
			} else {
				// no add
			}
		} else {
			tmpSum = 3
			if arr[0] == "A" {
				// no add
			} else if arr[0] == "B" {
				tmpSum = tmpSum + 6
			} else {
				tmpSum = tmpSum + 3
			}
		}

		sum = sum + tmpSum

		if !isPrefix {
			tmp = nil
		}
	}
	fmt.Printf("part1 合計スコア: %d \n", sum)
}

func solvePart2() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 64)
	var tmp []byte
	tmpSum := 0
	sum := 0

	// 何の手を選んだか　と　勝敗結果　がスコアになる
	for {
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)

		// splitする
		arr := strings.Split(string(tmp), "")

		// 1. 自分のだした手をスコア化
		// X: 負けろ, Y: 引き分けろ, Z: 勝て.
		// 1 for Rock, 2 for Paper, and 3 for Scissors
		// A for Rock, B for Paper, and C for Scissors
		// 0 if you lost, 3 if the round was a draw, and 6 if you won
		if arr[2] == "X" {
			// 負けろ

			tmpSum = 0
			if arr[0] == "A" {
				// グーに対して負ける必要があるためチョキを出す
				tmpSum = tmpSum + 3
			} else if arr[0] == "B" {
				// パーに対して負ける必要があるためグーを出す
				tmpSum = tmpSum + 1
			} else {
				// チョキに対して負ける必要があるためパーを出す
				tmpSum = tmpSum + 2
			}
		} else if arr[2] == "Y" {
			// 引き分けろ

			tmpSum = 3
			if arr[0] == "A" {
				// グーに対して引き分ける必要があるためグーを出す
				tmpSum = tmpSum + 1
			} else if arr[0] == "B" {
				// パーに対して引き分ける必要があるためパーを出す
				tmpSum = tmpSum + 2
			} else {
				// チョキに対して負ける必要があるためチョキを出す
				tmpSum = tmpSum + 3
			}
		} else {
			// 勝て

			tmpSum = 6
			if arr[0] == "A" {
				// グーに対して勝つ必要があるためパーを出す
				tmpSum = tmpSum + 2
			} else if arr[0] == "B" {
				// パーに対して勝つ必要があるためチョキを出す
				tmpSum = tmpSum + 3
			} else {
				// チョキに対して勝つ必要があるためグーを出す
				tmpSum = tmpSum + 1
			}
		}

		sum = sum + tmpSum

		if !isPrefix {
			tmp = nil
		}
	}
	fmt.Printf("part2 合計スコア: %d \n", sum)
}

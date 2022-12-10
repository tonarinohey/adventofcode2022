package main

import (
	"bufio"

	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

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
	sum := 0
	for {
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)

		stringTmp := string(tmp)
		firstHalf := stringTmp[:len(stringTmp)/2]
		letterHalf := stringTmp[len(stringTmp)/2:]

		// runeではA:65,B:66...a:97,b:98...z:122
		for _, c := range firstHalf {
			if strings.Contains(letterHalf, string(c)) {
				if unicode.IsUpper(c) {
					// 大文字の場合
					sum = sum + int(c) - 65 + 27
				} else {
					// 小文字の場合
					sum = sum + int(c) - 97 + 1
				}
				break
			}
		}

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
	sum := 0

	var candidateChars []rune
	var firstRow, secondRow, thirdRow string

	rowNum := 0
	// 3行単位で処理していく
	//
	for {
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)
		stringTmp := string(tmp)

		rowNum++
		if rowNum%3 == 1 {
			candidateChars = make([]rune, 0)

			fmt.Println("----")
			fmt.Println(rowNum)
			// firstRowに文字列詰める
			firstRow = stringTmp
		} else if rowNum%3 == 2 {
			// firstRowとsecondRowに共通する文字をcandidateCharに詰める
			secondRow = stringTmp
			for _, r := range secondRow {
				if strings.Contains(firstRow, string(r)) {
					candidateChars = append(candidateChars, r)
				}
			}
		} else {
			// candidateCharとthirdRowに共通する文字を見つけ、優先度をsumに加算する
			thirdRow = stringTmp
			for _, r := range candidateChars {
				if strings.Contains(thirdRow, string(r)) {
					fmt.Println(string(r))
					// println(string(r))
					if unicode.IsUpper(r) {
						// 大文字の場合
						sum = sum + int(r) - 65 + 27
						println(int(r) - 65 + 27)
					} else {
						// 小文字の場合
						sum = sum + int(r) - 97 + 1
						println(int(r) - 97 + 1)
					}
					break
				}
			}
		}

		if !isPrefix {
			tmp = nil
		}
	}
	fmt.Printf("part2 合計スコア: %d \n", sum)
}

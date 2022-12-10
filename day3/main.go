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
					fmt.Printf("文字: %s \n", string(c))
					fmt.Printf("スコア: %d \n", (int(c) - 65 + 27))
					sum = sum + int(c) - 65 + 27
				} else {
					fmt.Printf("文字: %s \n", string(c))
					fmt.Printf("スコア: %d \n", (int(c) - 97 + 1))
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

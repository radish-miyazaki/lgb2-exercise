package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("0 で割ることはできません")
	}

	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"2", "+", "three"},
		{"5"},
		{"2", "/", "0"},
	}

	for _, expression := range expressions {
		// 演算子と被演算子の合計個数のチェック
		if len(expression) != 3 {
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}

		// 1番目の被演算子のチェック
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}

		// 演算子のチェック
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}

		// 2番目の被演算子のチェック
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}

		// 実際の計算
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}

		fmt.Print(expression, " → ", result, "\n")
	}
}

func exercise2() {
	len, err := fileLen("main.go")
	if err != nil {
		fmt.Print("ファイルの長さの取得に失敗しました: ", err, "\n")
		return
	}
	fmt.Print("ファイルの長さ: ", len, "\n")
}

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return int(stat.Size()), nil
}

func exercise3() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

func prefixer(prefix string) func(string) string {
	return func(input string) string {
		return fmt.Sprintf("%s %s", prefix, input)
	}
}

package chapter2

import "fmt"

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func Calc(slice []int) (int, error) {
	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	length := len(slice)

	if length == 0 {
		return 0, fmt.Errorf("invalid length=%d", length)
	}

	if length <= 2 {
		if length == 1 {
			return slice[0], nil
		}
		return slice[0] * slice[1], nil
	}

	// 3以上は条件分けなくて良いはず
	var result int
	for _, value := range slice {
		result += value
	}
	return result, nil

}

type Number struct {
	index int
}

// 構造体Numberを3つの要素数から成るスライスにして返却
// 3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
func Numbers() []Number {
	var slice []Number
	slice = append(slice, Number{index: 1}, Number{index: 2}, Number{index: 3})
	return slice
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
// キー「yon」に関しては完全一致すること
func CalcMap(m map[string]int) int {
	var result int
	for key, value := range m {
		if key == "yon" {
			// yon 完全一致で除外
			continue
		}
		result += value
	}
	return result
}

type Model struct {
	Value int
}

// 与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
func Add(models []Model) {
	for i := 0; i < len(models); i++ {
		models[i].Value += 5
	}
}

// 引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
// 順序はスライスに格納されている順番のまま返却すること
// ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
func Unique(slice []int) []int {

	// mapを作成
	m := make(map[int]bool)
	for _, e := range slice {
		// 要素をキーとしてmapに登録（値はダミー）
		m[e] = true
	}

	// mapのキーだけのslice生成
	var result []int
	for key, _ := range m {
		result = append(result, key)
	}
	return result
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func Fibonacci() func() int {

	// 呼び出し回数に応じて要素が増えるスライス
	var fibonacciValues []int

	return func() int {
		length := len(fibonacciValues)

		// 要素数が0のとき
		if length == 0 {
			fibonacciValues = append(fibonacciValues, 0)
			return 0
		}

		// 要素数が1のとき
		if length == 1 {
			fibonacciValues = append(fibonacciValues, 1)
			return 1
		}

		// 要素数が2以上のとき
		sumValue := fibonacciValues[length-1] + fibonacciValues[length-2]
		fibonacciValues = append(fibonacciValues, sumValue)
		return sumValue

	}

}

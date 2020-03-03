package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

var (
	fOpt = flag.Int("f", 0, "表示したいフィールドの番号を指定してください。")
	dOpt = flag.String("d", ",", "区切り文字を指定してください。")
)

// go-cutコマンドを実装しよう
func main() {
	flag.Parse()
	argLen := len(os.Args)

	if argLen <= 3 {
		os.Stderr.WriteString("-f オプション と ファイル名が必要です。\n")
		os.Exit(1)
	}

	if *fOpt <= 0 {
		os.Stderr.WriteString("-f オプションは 0より大きい整数を指定してください。\n")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[argLen-1])
	if err != nil {
		os.Stderr.WriteString("ファイルを開けません。\n")
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		splitStr := strings.Split(scanner.Text(), *dOpt)

		// -fオプションの指定がsplitしたスライスの長さよりも長い場合
		if *fOpt > len(splitStr) {
			os.Stdout.WriteString("\n")
			continue
		}

		os.Stdout.WriteString(splitStr[*fOpt-1] + "\n")
	}

	if scanner.Err() != nil {
		os.Stderr.WriteString("読込中にエラーが発生しました。\n")
		os.Exit(1)
	}
}

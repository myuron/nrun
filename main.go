package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageJSON struct {
	Scripts map[string]string `json:"scripts"`
}

func main() {
	// package.jsonを読み込む
	data, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// json dataを構造体に変換する
	var scripts PackageJSON
	if err := json.Unmarshal(data, &scripts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// script一覧を表示する
	for key, value := range scripts.Scripts {
		fmt.Printf("%s: %s\n", key, value)
	}
}

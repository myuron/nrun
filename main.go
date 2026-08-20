package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
)

type PackageJSON struct {
	Scripts map[string]string `json:"scripts"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run() error {
	// package.jsonを読み込む
	data, err := os.ReadFile("package.json")
	if err != nil {
		return err
	}

	// json dataを構造体に変換する
	var pkg_json PackageJSON
	if err := json.Unmarshal(data, &pkg_json); err != nil {
		return err
	}

	// script一覧を表示する
	keys := slices.Sorted(maps.Keys(pkg_json.Scripts))
	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, pkg_json.Scripts[key])
	}

	return nil
}

package main

import "fmt"

func main() {
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 初期値
	var (
		j      int
		ff64   float64
		str    string // 空文字
		tr, fl bool
	)
	fmt.Println(j, ff64, str, tr, fl)

	// 関数内のみの短縮記法。
	k := 1
	fff64 := 1.2
	fmt.Println(k, fff64)
}

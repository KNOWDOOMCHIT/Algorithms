package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var files [50]string
	for i := 0; i < n; i++ {
		fmt.Scan(&files[i])
	}

	if n == 1 {
		fmt.Println(files[0])
		return
	}

	var result string
	for i := 0; i < len(files[0]); i++ {
		str := files[0][i]
		for j := 0; j < n-1; j++ {
			if files[j][i] != files[j+1][i] {
				str = '?'
			}
		}
		result = fmt.Sprintf("%s%c", result, str)
	}

	fmt.Println(result)
}

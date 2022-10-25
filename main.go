package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

var verbose = flag.Bool("v", false, "verbose")

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("<> ")
		wordBytes, _, _ := reader.ReadLine()
		word := strings.Split(string(wordBytes), " ")[0]
		fmt.Println(suggestions(word))
	}
}

func suggestions(word string) []string {
	arr := [][]int{}
	for i, w := range wordlist {
		arr = append(arr, []int{i, levenshteinDis(word, w)})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	res := make([]string, 4)
	if *verbose {
		fmt.Println(arr[:4])
	}
	for i := 0; i < 4; i++ {
		res[i] = wordlist[arr[i][0]]
	}
	return res
}

func levenshteinDis(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + int(math.Min(float64(dp[i][j-1]), math.Min(float64(dp[i][j-1]), float64(dp[i-1][j-1]))))
			}
		}
	}
	return dp[n][m]
}

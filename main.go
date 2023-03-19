package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/Aksh-Bansal-dev/doyoumean/internal/color"
	"github.com/Aksh-Bansal-dev/doyoumean/internal/cursor"
	"github.com/Aksh-Bansal-dev/doyoumean/internal/wordlist"
)

var (
	resultLen = flag.Int("n", 5, "number of items in result")
	filePath  = flag.String("f", "", "path of file to fuzzy search")
	showScore = flag.Bool("v", false, "show score of each result")
)

func main() {
	flag.Parse()

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	str := strings.Builder{}
	b := make([]byte, 1)
	for {
		fmt.Print(cursor.ClearEntireLine())
		fmt.Print(cursor.MoveLeft(100))
		fmt.Print("<> ", str.String())
		os.Stdin.Read(b)
		if b[0] == 127 {
			s := str.String()
			str.Reset()
			if len(s) > 0 {
				str.WriteString(s[:len(s)-1])
			}
			if len(str.String()) == 0 {
				continue
			}
		} else {
			str.Write(b)
		}
		fmt.Print("\n")
		printInplace(suggestions(str.String()))
		fmt.Print(cursor.MoveUp(1))
	}
}

type Result struct {
	val   string
	index int
	score int
	lcs   string
}

func suggestions(word string) []Result {
	arr := [][]int{}
	var list []string
	if len(*filePath) != 0 {
		data, err := os.ReadFile(*filePath)
		if err != nil {
			log.Fatal(err)
		}
		list = strings.Split(string(data), "\n")
	} else {
		list = wordlist.Wordlist
	}
	lcsList := []string{}
	for i, w := range list {
		if len(*filePath) == 0 {
			arr = append(arr, []int{i, levenshteinDis(word, w)})
		} else {
			lcs := longestCommonSubsequence(word, w)
			lcsList = append(lcsList, lcs)
			arr = append(arr, []int{i, -len(lcs)})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	res := []Result{}
	for i := 0; len(res) < *resultLen && i < len(arr); i++ {
		if len(list[arr[i][0]]) > 0 {
			res = append(res, Result{list[arr[i][0]], arr[i][0], arr[i][1], lcsList[arr[i][0]]})
		}
	}
	return res
}

func printInplace(arr []Result) {
	cnt := 0
	for _, e := range arr {
		if len(e.val) == 0 {
			continue
		}
		cnt++
		fmt.Print(cursor.ClearEntireLine())
		fmt.Print(cursor.MoveLeft(200))
		if *showScore {
			fmt.Printf("[%d] %s \t%s\n", e.index, highlightCommon(e.val, e.lcs), color.FadeColor(fmt.Sprintf("%d", e.score)))
		} else {
			fmt.Printf("[%d] %s\n", e.index, highlightCommon(e.val, e.lcs))
		}
	}
	fmt.Print(cursor.MoveUp(cnt))
}

func highlightCommon(s, lcs string) string {
	idx := 0
	res := strings.Builder{}
	for i := range s {
		if idx < len(lcs) && s[i] == lcs[idx] {
			idx++
			res.WriteString(color.HighlightColor(string(s[i])))
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

func longestCommonSubsequence(s1 string, s2 string) string {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	res := make([]byte, dp[n][m])
	pos := dp[n][m] - 1
	i := n
	j := m
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			res[pos] = s1[i-1]
			i--
			j--
            pos--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(res)
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

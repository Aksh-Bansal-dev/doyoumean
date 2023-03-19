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

	"github.com/Aksh-Bansal-dev/doyoumean/internals/cursor"
	"github.com/Aksh-Bansal-dev/doyoumean/internals/wordlist"
)

var (
	verbose   = flag.Bool("v", false, "verbose")
	resultLen = flag.Int("n", 5, "number of items in result")
	filePath  = flag.String("f", "", "path of file to fuzzy search")
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
            if len(str.String())==0{
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

type Result struct{
    val string
    index int
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
	for i, w := range list {
		splitWords := strings.Split(w, " ")
		for _, ww := range splitWords {
			ld := levenshteinDis(word, ww)
			if len(arr) == i {
				arr = append(arr, []int{i, ld})
			} else {
				if ld < arr[i][1] {
					arr[i][1] = ld
				}
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	res := make([]Result, *resultLen)
	if *verbose {
		fmt.Println(arr[:*resultLen])
	}
	for i := 0; i < *resultLen; i++ {
		res[i] = Result{list[arr[i][0]],arr[i][0]}
	}
	return res
}

func printInplace(arr []Result) {
    cnt := 0
	for _, e := range arr {
        if len(e.val)==0{
            continue
        }
        cnt++
		fmt.Print(cursor.ClearEntireLine())
		fmt.Print(cursor.MoveLeft(200))
		fmt.Printf("[%d] %s\n",e.index, e.val)
	}
	fmt.Print(cursor.MoveUp(cnt))
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

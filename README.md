# doyoumean

A CLI tool to fuzzy search any file or dictionary.

It uses Longest Common Subsequence for fuzzy search and levenshtein distance algorithm for dictionary search.

![ezgif com-crop](https://user-images.githubusercontent.com/63552235/226174195-8ea8a798-75f5-4a78-a48d-50dc3da5e586.gif)

## How to use?

- Run `go run main.go -f <file_to_be_fuzzy_searched>` for fuzzy seach mode
- Run `go run main.go` for dictionary mode

> Note: you can see all available flags using `-h` flag

## Credits

The words in wordlist are from [ef.com](https://www.ef.com/wwen/english-resources/english-vocabulary/top-1000-words/).

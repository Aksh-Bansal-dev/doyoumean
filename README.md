# doyoumean

A CLI tool to fuzzy search any file or dictionary.

It uses Longest Common Subsequence for fuzzy search and levenshtein distance algorithm for dictionary search.

![ezgif com-crop](https://user-images.githubusercontent.com/63552235/226174736-4c0e1820-8f30-4856-bab0-c1940a98f623.gif)

## How to use?

- Run `go run main.go -f <file_to_be_fuzzy_searched>` for fuzzy seach mode
- Run `go run main.go` for dictionary mode

> Note: you can see all available flags using `-h` flag

## Credits

The words in wordlist are from [ef.com](https://www.ef.com/wwen/english-resources/english-vocabulary/top-1000-words/).

package color

import "fmt"

func CustomColor(s string, code int) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", code, s)
}
func HighlightColor(s string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", 75, s)
}
func FadeColor(s string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", 239, s)
}

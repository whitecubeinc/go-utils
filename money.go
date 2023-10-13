package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FormatKoreanCurrency 숫자 -> 한국 금액 표기법으로 변경
func FormatKoreanCurrency(n int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", n)
}

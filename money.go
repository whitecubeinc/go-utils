package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

// FormatKoreanCurrency 숫자 -> 한국 금액 표기법으로 변경
func FormatKoreanCurrency(n int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", n)
}

func CalDailyInterest(orgAmount int, interestRate float64, interestStartDate time.Time, standardDate time.Time) (interest int) {
	dayDiff := standardDate.Sub(interestStartDate).Hours() / 24
	return int(CeilAt(float64(orgAmount)*interestRate*dayDiff, 2))
}

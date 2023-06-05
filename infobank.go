package utils

import (
	"fmt"
	"math"
	"net/http"
)

type MessageFormat struct {
	From         string `json:"from,omitempty"`
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	Destinations []Destination
}

type Token struct {
	Schema      string
	AccessToken string
}

type Destination struct {
	UserId int    `json:"-"`  // info bank 에는 보내지 않고 내부 로직에서만 사용
	To     string `json:"to"` // ex) 821012345678

	// replaceWord -> %CHANGEWORD1% in text
	ReplaceWord1 string `json:"replaceWord1,omitempty"`
	ReplaceWord2 string `json:"replaceWord2,omitempty"`
	ReplaceWord3 string `json:"replaceWord3,omitempty"`
	ReplaceWord4 string `json:"replaceWord4,omitempty"`
	ReplaceWord5 string `json:"replaceWord5,omitempty"`
}

type MessageResult struct {
	TotalCount    int
	SuccessCount  int
	SuccessPhones []string
}

type InfoBank struct {
	authUrl          string
	sendSmsUrl       string
	infoBankSmsId    string
	infoBankPassword string
}

func NewInfoBank(authUrl, sendSmsUrl, infoBankSmsId, infoBankPassword string) *InfoBank {
	return &InfoBank{
		authUrl:          authUrl,
		sendSmsUrl:       sendSmsUrl,
		infoBankSmsId:    infoBankSmsId,
		infoBankPassword: infoBankPassword,
	}
}

func (o *InfoBank) getToken() (Token, error) {
	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("X-IB-Client-Id", o.infoBankSmsId)
	header.Add("X-IB-Client-Passwd", o.infoBankPassword)

	return Post[Token](o.authUrl, nil, header)
}

type sendSmsBody struct {
	From         string `json:"from,omitempty"`
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	Destinations []Destination
	TTL          int
}

func (o *InfoBank) SendSms(format MessageFormat) (result MessageResult, err error) {
	if len(format.Destinations) == 0 {
		return
	}

	token, err := o.getToken()
	if err != nil {
		return
	}

	// 200개씩 나눠서 발송
	size := 200
	interval := int(math.Ceil(float64(len(format.Destinations)) / float64(size)))
	for i := 0; i < interval; i++ {
		startIdx := i * size
		endIdx := (i + 1) * size
		if len(format.Destinations) < endIdx {
			endIdx = len(format.Destinations)
		}

		// FIXME go routine 으로 처리
		body := sendSmsBody{
			From:         format.From,
			Title:        format.Title,
			Body:         format.Body,
			Destinations: format.Destinations[startIdx:endIdx],
			TTL:          300,
		}

		header := make(http.Header)
		header.Add("Authorization", fmt.Sprintf("%s %s", token.Schema, token.AccessToken))
		header.Add("Accept", "application/json")
		res, err := Post[map[string]any](o.sendSmsUrl, body, header)

		if err != nil {
			return
		}

		resultSlice := res["destinations"].([]interface{})
		for _, rs := range resultSlice {
			if rs.(map[string]any)["status"] == "R000" {
				result.SuccessCount++
				result.SuccessPhones = append(result.SuccessPhones, rs.(map[string]any)["to"].(string))
			}
		}
	}

	return
}

/**
 * File              : sms.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 01.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package sms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/jiangytcn/gosms-ng/modem"
)

type SMS struct {
	UUID      string `json:"uuid"`
	CMGLID    int    `json:"cmgl_id"`
	Mobile    string `json:"mobile"`
	Body      string `json:"body"`
	Status    int    `json:"status"`
	Type      string `json:"type"`
	Retries   int    `json:"retries"`
	Device    string `json:"device"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Routes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", GetAllMessages)
	return router
}

var raspModem *modem.GSMModem

func init() {
	_port := "/dev/ttyS0"
	_baud := 115200
	raspModem = modem.New(_port, _baud, "mymodem")
	if err := raspModem.Connect(); err != nil {
		panic(err)
	}
}
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	var smss []SMS
	var err error
	allMsgs := raspModem.ReadSMSs()
	msgsStr := strings.Join(allMsgs, "")
	data := strings.Split(msgsStr, "\n")
	var sanitized_data []string
	// +CMGL: 13,"REC READ","15675180530","","19/08/07,21:54:54+32"
	// Hhhhhh

	for _, val := range data {
		if len(val) >= 5 {
			sanitized_data = append(sanitized_data, val)
		}
	}

	for i := 0; i < len(sanitized_data); i = i + 2 {
		sms := SMS{}
		headers := strings.Split(sanitized_data[i], ",")
		if len(headers) < 5 {
			continue
		}
		sms.CMGLID, _ = strconv.Atoi(strings.Trim((strings.Split(headers[0], ":"))[1], " "))
		sms.Type = cleanStr(headers[1])
		sms.Mobile = cleanStr(headers[2])
		sms.CreatedAt = cleanStr(fmt.Sprintf("%s %s", headers[4], headers[5]))

		if sms.Body, err = decodeHexUTF16(sanitized_data[i+1]); err != nil {
			sms.Body = sanitized_data[i+1]
		}
		smss = append(smss, sms)
	}

	bdata, err := json.Marshal(smss)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(bdata)
	}

}

func cleanStr(text string) string {
	return strings.Trim(strings.TrimSpace(text), "\"")
}

func decodeHexUTF16(text string) (string, error) {
	l := len(text) / 4
	var encodedStr []string
	var pre, last int
	for i := 0; i < l; i++ {
		if i == 0 {
			pre = 0
		} else {
			pre += 4
		}
		last = pre + 4
		if safeSubstring, err := unquoteCodePoint((text[pre:last])); err != nil {
			return "", err
		} else {
			encodedStr = append(encodedStr, safeSubstring)
		}

	}
	return strings.Join(encodedStr, ""), nil
}

func unquoteCodePoint(s string) (string, error) {
	// 16 specifies hex encoding
	// 32 is size in bits of the rune type
	r, err := strconv.ParseInt(s, 16, 32)
	return string(r), err
}

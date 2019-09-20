/**
 * File              : sms.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 19.09.2019
 * Last Modified Date: 20.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package sms

import (
	"strconv"
	"strings"

	"github.com/jiangytcn/gosms-ng/pkg/models"
)

func CleanStr(text string) string {
	return strings.Trim(strings.TrimSpace(text), "\"")
}

func DecodeHexUTF16(text string) (string, error) {
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

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []models.SMS) []models.SMS {
	mb := make(map[string]struct{}, len(b))
	for _, sms := range b {
		mb[sms.CMGLID] = struct{}{}
	}
	var diff []models.SMS
	if len(a) == 0 {
		diff = b
	} else {
		for _, sms := range a {
			if _, found := mb[sms.CMGLID]; !found {
				diff = append(diff, sms)
			}
		}
	}
	return diff
}

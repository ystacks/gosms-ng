/**
 * File              : modem_timed_worker.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 20.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package scheduler

import (
	"fmt"
	"strings"

	"github.com/jiangytcn/gosms-ng/modem"
	"github.com/jiangytcn/gosms-ng/pkg/models"
	"github.com/jiangytcn/gosms-ng/pkg/notifier"
	smsutil "github.com/jiangytcn/gosms-ng/pkg/util/sms"
)

type ModemWorker struct {
	Notification *notifier.NotificationService
	Modem        *modem.GSMModem
	rawMsgCh     chan string
}

func (worker ModemWorker) Run() {
	// obtains sms and parse it
	var err error
	allMsgs := worker.Modem.ReadSMSs()
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
		sms := models.SMS{}
		headers := strings.Split(sanitized_data[i], ",")
		if len(headers) < 5 {
			continue
		}
		sms.CMGLID = strings.Trim((strings.Split(headers[0], ":"))[1], " ")
		sms.ID = sms.CMGLID
		sms.Type = smsutil.CleanStr(headers[1])
		sms.Mobile = smsutil.CleanStr(headers[2])
		sms.CreatedAt = smsutil.CleanStr(fmt.Sprintf("%s %s", headers[4], headers[5]))

		if sms.Body, err = smsutil.DecodeHexUTF16(sanitized_data[i+1]); err != nil {
			sms.Body = sanitized_data[i+1]
		}

		worker.Notification.Enqueue(sms)

	}

}

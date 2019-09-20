/**
 * File              : notifier.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 17.08.2019
 * Last Modified Date: 20.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package notifier

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/jiangytcn/gosms-ng/config"
	. "github.com/jiangytcn/gosms-ng/logger"
	"github.com/jiangytcn/gosms-ng/pkg/models"
	smsutil "github.com/jiangytcn/gosms-ng/pkg/util/sms"
	"go.uber.org/zap"
)

var mailTemplates *template.Template

type Notifier interface {
	WebRequestSend() error
	GenericSend() error
}

type NotificationService struct {
	smsQueue     chan models.SMS
	webhookQueue chan Webhook
	done         chan struct{}
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		smsQueue:     make(chan models.SMS),
		webhookQueue: make(chan Webhook),
		done:         make(chan struct{}),
	}
}

func (ns *NotificationService) Run() error {

	ticker := time.NewTicker(time.Minute * 12)
	var smsData, lastSent []models.SMS

	for {
		select {
		case sms := <-ns.smsQueue:
			smsData = append(smsData, sms)

		case <-ticker.C:
			Logger.Info("Timer UP, SMS Sent started")
			if len(smsData) == 0 {
				Logger.Info("SMS Sent Skipped, Empty body")
				continue
			}
			message, err := ns.buildEmailMessage(smsutil.Difference(smsData, lastSent))
			if err == nil {
				if err = message.GenericSend(); err != nil {
					Logger.Error("SMS Sent failed", zap.Error(err))
				} else {
					Logger.Info("SMS Sent succeed")
					lastSent = smsData
					smsData = nil
				}
			} else if err.Error() == "empty sms" {
				Logger.Info("SMS Sent Skipped, Empty body")
			} else {
				Logger.Error("SMS build failed", zap.Error(err))
			}

		case webhook := <-ns.webhookQueue:
			fmt.Println("got webhook, pending s")
			webhook.WebRequestSend()

		case <-ns.done:
			Logger.Info("Shutting down Notification µ-service")
			break
		}
	}
}

func (ns *NotificationService) Enqueue(payload interface{}) {
	Logger.Info("sms queued")
	ns.smsQueue <- payload.(models.SMS)
	//ns.webhookQueue <- ()payload
}

func (ns *NotificationService) buildEmailMessage(data []models.SMS) (*Message, error) {

	if len(os.Getenv("SENDGRID_API_KEY")) == 0 {
		return nil, fmt.Errorf("cannot get 'SENDGRID_API_KEY' variable")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty sms")
	}

	const letter = `
Hi,
<br/>
<br/>
{{ range $sms := . }}
   *************************************************
   <li><strong>From</strong>: {{ $sms.Mobile }}</li>
   <li><strong>Created AT</strong>: {{ $sms.CreatedAt }}</li>
   <li><strong>CMGLID</strong>: {{ $sms.CMGLID }}</li>
   <li><strong>Body</strong>: {{ $sms.Body }}</li>
{{ end }}

Best wishes,
Yitao
`

	buffer := &bytes.Buffer{}
	var err error
	var subject = "GSM Notifications"

	mailTemplates := template.Must(template.New("letter").Parse(letter))

	err = mailTemplates.Execute(buffer, data)
	if err != nil {
		return nil, err
	}

	return &Message{
		To:      []string{"jiangyt.cn@gmail.com", "jiangytcn@163.com", "495585032@qq.com"},
		Subject: subject,
		Body:    buffer.String(),
		SMTPCfg: config.SMTPConfig{
			Server:      "smtp.163.com",
			Port:        465,
			User:        "willierjyt@163.com",
			Password:    os.Getenv("SMTP_API_KEY"),
			FromName:    "Jiang Yi Tao",
			FromAddress: "willierjyt@163.com",
		},
	}, nil
}

/**
 * File              : notifier.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 17.08.2019
 * Last Modified Date: 19.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package notifier

import (
	"bufio"
	"bytes"
	"fmt"
	"text/template"

	"github.com/ystacks/gosms-ng/config"
	. "github.com/ystacks/gosms-ng/logger"
	"github.com/ystacks/gosms-ng/pkg/models"
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

	for {
		select {
		case sms := <-ns.smsQueue:
			message, err := ns.buildEmailMessage(sms)
			if err == nil {
				if err = message.GenericSend(); err != nil {
					Logger.Error("SMS Sent failed", zap.Error(err))
				}
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
	/*
		notifier_type := reflect.TypeOf(notifier)
		switch notifier_type.Name() {
		case "Message":
		case "Webhook":
		}
	*/
}

func (ns *NotificationService) buildEmailMessage(sms models.SMS) (*Message, error) {

	const letter = `
Hi,

{{.}}

Best wishes,
Yitao
`

	var buffer bytes.Buffer
	var err error
	var subject = "hello message"

	data := sms.Data
	bufWriter := bufio.NewWriter(&buffer)

	mailTemplates := template.Must(template.New("letter").Parse(letter))

	err = mailTemplates.Execute(bufWriter, data)
	fmt.Println(buffer.String())
	if err != nil {
		return nil, err
	}

	return &Message{
		To:      []string{"jiangyt.cn@gmail.com", "jiangytcn@163.com"},
		Subject: subject,
		Body:    buffer.String(),
		SMTPCfg: config.SMTPConfig{
			Server:      "smtp.163.com",
			Port:        465,
			User:        "jiangytcn@163.com",
			Password:    "changeit",
			FromName:    "Jiang Yi Tao",
			FromAddress: "jiangytcn@163.com",
		},
	}, nil
}

/**
 * File              : notifier_test.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 18.08.2019
 * Last Modified Date: 20.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package notifier_test

import (
	"testing"

	"github.com/jiangytcn/gosms-ng/pkg/models"
	. "github.com/jiangytcn/gosms-ng/pkg/notifier"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSMSEmailNotification(t *testing.T) {
	Convey("Initilize the Notification Service", t, func() {
		smsNotification := NewNotificationService()
		go smsNotification.Run()

		Convey("Create a sample sms", func() {
			sms := models.SMS{
				ID:     "1",
				Mobile: "13173748220",
				Body:   "hello, world --from test",
			}

			Convey("Enqueue the sms", func() {
				smsNotification.Enqueue(sms)
			})
		})
	})
}

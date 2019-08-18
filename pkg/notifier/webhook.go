/**
 * File              : webhook.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 17.08.2019
 * Last Modified Date: 18.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package notifier

import "fmt"

type Webhook struct {
	Url         string
	User        string
	Password    string
	Body        string
	HttpMethod  string
	HttpHeader  map[string]string
	ContentType string
}

func (hook Webhook) WebRequestSend() error {
	return fmt.Errorf("not implemented")
}

func (hook Webhook) GenericSend() error {
	return fmt.Errorf("not implemented")
}

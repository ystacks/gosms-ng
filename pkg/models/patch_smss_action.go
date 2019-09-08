/**
 * File              : patch_sms_action.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 08.09.2019
 * Last Modified Date: 08.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package models

type SMSsAction struct {
	Action string   `json:"action"`
	SMSIDS []string `json:"sms_ids"`
}

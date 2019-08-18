/**
 * File              : sms_type.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 18.08.2019
 * Last Modified Date: 18.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package models

type SMS_TYPE string

const (
	SMS_SENT      string   = ""
	SMS_RECV      string   = ""
	SMS_RECV_READ SMS_TYPE = "REC READ"
)

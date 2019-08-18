/**
 * File              : sms.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 18.08.2019
 * Last Modified Date: 18.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package models

type SMS struct {
	ID        int
	Type      SMS_TYPE
	Contact   string
	Data      []byte
	CreatedAT string
}

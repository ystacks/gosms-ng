/**
 * File              : sms.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 18.08.2019
 * Last Modified Date: 19.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package models

type SMS struct {
	ID        string `json:"id"`
	CMGLID    string `json:"cmgl_id"`
	Mobile    string `json:"mobile"`
	Body      string `json:"body"`
	Status    int    `json:"status"`
	Type      string `json:"type"`
	Retries   int    `json:"retries"`
	Device    string `json:"device"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

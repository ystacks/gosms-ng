/**
 * File              : patch_response.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 08.09.2019
 * Last Modified Date: 08.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package models

type PatchError struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}
type PatchResponse struct {
	ID      string       `json:"id"`
	Success bool         `json:"success"`
	Errors  []PatchError `json:"errors"`
}

/**
 * File              : patch_smss_response.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 08.09.2019
 * Last Modified Date: 08.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package models

type PatchSMSsResponse struct {
	SMSS []interface{} `json:"smss"`
}

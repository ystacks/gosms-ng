/**
 * File              : controller.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 11.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package controller

type Controller interface {
	Run(stopCh chan struct{}, threadiness int)
}

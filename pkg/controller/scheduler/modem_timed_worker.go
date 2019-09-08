/**
 * File              : modem_timed_worker.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 18.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package scheduler

import (
	"fmt"

	"github.com/jiangytcn/gosms-ng/modem"
)

type ModemWorker struct {
	modem    *modem.GSMModem
	rawMsgCh chan string
}

func (worker ModemWorker) Run() {
	// obtains sms and parse it
	smss := worker.modem.ReadSMSs()
	fmt.Println(smss)
}

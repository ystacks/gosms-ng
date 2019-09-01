/**
 * File              : sms.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 11.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package modem

import (
	"sync"

	. "github.com/jiangytcn/gosms-ng/logger"
	"go.uber.org/zap"
)

type SMSController struct {
}

func (c *SMSController) Run(stopCh chan struct{}, threadiness int) {

	var waitgroup sync.WaitGroup
	waitgroup.Add(threadiness)
	Logger.Info("Starting SMS Controller")

	for i := 0; i < threadiness; i++ {
		go c.runWorker(&waitgroup, i)
	}
	waitgroup.Wait()
	<-stopCh
	Logger.Info("Shutting down SMS Controller")
}

func (c *SMSController) runWorker(wg *sync.WaitGroup, thread int) {
	Logger.Info("sms controller work runs start", zap.Int("thread", thread))
	wg.Done()
	Logger.Info("sms controller work runs end", zap.Int("thread", thread))
}

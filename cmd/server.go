/**
 * File              : server.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 10.08.2019
 * Last Modified Date: 20.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
/*
Copyright © 2019 Jiang Yitao <jiangyt.cn#gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bamzi/jobrunner"
	. "github.com/jiangytcn/gosms-ng/logger"
	"github.com/jiangytcn/gosms-ng/modem"
	"github.com/jiangytcn/gosms-ng/pkg/api"
	"github.com/jiangytcn/gosms-ng/pkg/controller/scheduler"
	"github.com/jiangytcn/gosms-ng/pkg/notifier"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var raspModem *modem.GSMModem

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		r := api.Routes()

		stopCh := make(chan struct{})

		srv := &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%v", 18080),
			Handler: r,
		}

		notification := notifier.NewNotificationService()
		go notification.Run()

		jobrunner.Start()
		jobrunner.Schedule("@every 10m", scheduler.ModemWorker{
			Modem:        raspModem,
			Notification: notification,
		})
		/*
			var smsCtr controller.Controller
			smsCtr = &modem.SMSController{}
			go smsCtr.Run(stopCh, 10)
		*/

		go func() {
			if err := srv.ListenAndServe(); err != nil {
				Logger.Fatal("failed to start µ-service for gosms-ng", zap.Int("port", 8080))
			}
		}()

		Logger.Info("µ-service for gosms-ng is running", zap.Int("port", 18080))

		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)
		<-sig

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			Logger.Warn(fmt.Sprintf("Server Shutdown: %#v", err))
		} else {
			stopCh <- struct{}{}
			Logger.Info("Shutting down gosms-ng µ-service")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	_port := "/dev/ttyS0"
	_baud := 115200
	raspModem = modem.New(_port, _baud, "mymodem")
	if err := raspModem.Connect(); err != nil {
		panic(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

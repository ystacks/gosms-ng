/**
 * File              : config.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 18.08.2019
 * Last Modified Date: 18.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package config

type SMTPConfig struct {
	Server      string
	Port        int
	User        string
	Password    string
	FromName    string
	FromAddress string
}

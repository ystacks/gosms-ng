/**
 * File              : modem.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 10.08.2019
 * Last Modified Date: 11.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 * Contrib           : https://github.com/haxpax/gosms/blob/master/modem/modem.go
 */
package modem

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jiangytcn/serial"
)

type GSMModem struct {
	ComPort  string
	BaudRate int
	Port     *serial.Port
	DeviceId string
}

func New(ComPort string, BaudRate int, DeviceId string) (modem *GSMModem) {
	modem = &GSMModem{ComPort: ComPort, BaudRate: BaudRate, DeviceId: DeviceId}
	return modem
}

func (m *GSMModem) Connect() (err error) {
	config := &serial.Config{Name: m.ComPort, Baud: m.BaudRate, ReadTimeout: time.Second}
	m.Port, err = serial.OpenPort(config)

	if err == nil {
		m.initModem()
	}

	return err
}

func (m *GSMModem) initModem() {
	m.SendCommand("ATE0\r\n", false)      // echo off
	m.SendCommand("AT+CMEE=1\r\n", false) // useful error messages
	m.SendCommand("AT+WIND=0\r\n", false) // disable notifications
	m.SendCommand("AT+CMGF=1\r\n", false) // switch to TEXT mode
}

func (m *GSMModem) Expect(possibilities []string) (string, error) {
	readMax := 0
	for _, possibility := range possibilities {
		length := len(possibility)
		if length > readMax {
			readMax = length
		}
	}

	readMax = readMax + 2 // we need offset for \r\n sent by modem

	var output []string
	var status string = ""
	buf := make([]byte, readMax)

	for {
		c, _ := m.Port.Read(buf)
		if c > 0 {
			output = append(output, string(buf[:c]))
		} else {
			break
		}
	}

	status = strings.Join(output, "")

	status = strings.Trim(strings.TrimSpace(status), "\"")

	if len(status) > 0 {
		for _, possibility := range possibilities {
			if strings.HasSuffix(status, possibility) {
				log.Println("--- Expect:", m.transposeLog(strings.Join(possibilities, "|")), "Got:", m.transposeLog(status))
				return status, nil
			}
		}
	}

	log.Println("--- Expect:", m.transposeLog(strings.Join(possibilities, "|")), "Got:", m.transposeLog(status), "(match not found!)")
	return status, errors.New("match not found")
}

func (m *GSMModem) Send(command string) {
	log.Println("--- Send:", m.transposeLog(command))
	m.Port.Flush()
	_, err := m.Port.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *GSMModem) Read(n int) []string {
	var output []string
	buf := make([]byte, n)
	for {
		c, _ := m.Port.Read(buf)
		if c > 0 {
			output = append(output, string(buf[:c]))
		} else {
			break
		}
	}

	return output
}

func (m *GSMModem) SendCommand(command string, waitForOk bool) string {
	m.Send(command)

	if waitForOk {
		output, _ := m.Expect([]string{"OK\r\n", "ERROR\r\n", "OK", "ERROR"}) // we will not change api so errors are ignored for now
		return output
	} else {
		return m.Read(1)[0]
	}
}

func (m *GSMModem) SendSMS(mobile string, message string) string {
	log.Println("--- SendSMS ", mobile, message)

	m.Send("AT+CMGS=\"" + mobile + "\"\r") // should return ">"
	m.Read(3)

	// EOM CTRL-Z = 26
	return m.SendCommand(message+string(26), true)
}

func (m *GSMModem) ReadSMSs() []string {
	log.Println("--- Read All SMS start.")

	m.Send("AT+CMGL=\"ALL\"\r\n")
	smsStr := m.Read(1024)
	log.Println("--- Read All SMS end.")
	return smsStr
}

func (m *GSMModem) DeleteSMS(id int) string {
	return m.SendCommand(fmt.Sprintf("AT+CMGD=%d\r\n", id), true)
}

func (m *GSMModem) transposeLog(input string) string {
	output := strings.Replace(input, "\r\n", "\\r\\n", -1)
	return strings.Replace(output, "\r", "\\r", -1)
}

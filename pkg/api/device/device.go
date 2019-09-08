/**
 * File              : device.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 08.09.2019
 * Last Modified Date: 08.09.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package device

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jiangytcn/gosms-ng/modem"
	"github.com/jiangytcn/gosms-ng/pkg/models"
)

func Routes() chi.Router {
	router := chi.NewRouter()
	router.Patch("/{deviceid}/sms", patchDevice)
	return router
}

var raspModem *modem.GSMModem

func init() {
	_port := "/dev/ttyS0"
	_baud := 115200
	raspModem = modem.New(_port, _baud, "mymodem")
	if err := raspModem.Connect(); err != nil {
		panic(err)
	}
}
func patchDevice(w http.ResponseWriter, r *http.Request) {

	var genericResp models.PatchSMSsResponse
	var responses []interface{}
	payload := models.SMSsAction{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	for _, smsid := range payload.SMSIDS {
		gsmResp := models.PatchResponse{ID: smsid}
		intID, err := strconv.Atoi(smsid)
		if err == nil {
			res := raspModem.DeleteSMS(intID)
			if res != "OK" {
				gsmResp.Success = false
				gsmResp.Errors = []models.PatchError{models.PatchError{Message: "invalid response from GSM device " + res}}
			} else {
				gsmResp.Success = true
			}
		} else {
			gsmResp.Success = false
			gsmResp.Errors = []models.PatchError{models.PatchError{Message: err.Error()}}
		}
		responses = append(responses, gsmResp)
	}

	genericResp = models.PatchSMSsResponse{
		SMSS: responses,
	}

	bdata, err := json.Marshal(genericResp)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(bdata)
	}

}

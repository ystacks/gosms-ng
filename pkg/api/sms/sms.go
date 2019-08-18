/**
 * File              : sms.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 11.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package sms

import (
	"net/http"

	"github.com/go-chi/chi"
)

type SMS struct {
	UUID      string `json:"uuid"`
	Mobile    string `json:"mobile"`
	Body      string `json:"body"`
	Status    int    `json:"status"`
	Type      string `json:"type"`
	Retries   int    `json:"retries"`
	Device    string `json:"device"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Routes() *chi.Mux {
	router := chi.NewMux()
	router.Get("/", GetAllMessages)
	return router
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
}

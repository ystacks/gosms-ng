/**
 * File              : config.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 11.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
// Package config provides server config routes
package config

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	router := chi.NewMux()
	return router
}

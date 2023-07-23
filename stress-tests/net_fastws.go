package main

import (
	"net/http"

	"github.com/orzogc/fastws"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fastws.NetUpgrade(func(c *fastws.Conn) {
		var (
			bf  []byte
			err error
			m   fastws.Mode
		)

		c.SetReadTimeout(0)
		c.SetWriteTimeout(0)

		for {
			m, bf, err = c.ReadMessage(bf[:0])
			if err != nil {
				break
			}

			c.WriteMessage(m, bf)
		}

		c.Close()
	})(w, r)
}

func main() {
	http.ListenAndServe(":8081", &handler{})
}

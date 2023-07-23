package main

import (
	"fmt"

	"github.com/orzogc/fastws"
	"github.com/valyala/fasthttp"
)

func main() {
	s := fasthttp.Server{
		Handler: fastws.Upgrade(func(c *fastws.Conn) {
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
		}),
	}
	fmt.Println(s.ListenAndServe(":8081"))
}

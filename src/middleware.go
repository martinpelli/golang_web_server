package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, router *http.Request) {
			flag := false
			fmt.Println("Checking authentication")
			if flag {
				handler(writer, router)
			} else {
				return
			}
		}
	}
}

func Login() Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, router *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(router.URL.Path, time.Since(start))
			}()
			handler(writer, router)
		}
	}
}

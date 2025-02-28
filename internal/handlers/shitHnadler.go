package handlers

import (
	"fmt"
	"net/http"
)

type ShitHandler struct {
}

func (sh *ShitHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie("token"))
}

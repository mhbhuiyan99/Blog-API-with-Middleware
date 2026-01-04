package post

import (
	"blogAPI/database"
	"blogAPI/util"
	"fmt"
	"net/http"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPosts handler called")
	util.SendData(w, database.List(), 200)
}
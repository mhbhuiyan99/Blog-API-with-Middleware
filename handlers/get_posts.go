package handlers

import (
	"net/http"
	"blogAPI/database"
	"blogAPI/util"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.PostList, 200)
}
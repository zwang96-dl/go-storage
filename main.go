package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nicemayi/go-storage/handler"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets/build")))
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHanlder)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get a ping!")
		data, _ := json.Marshal("pong from server")
		time.Sleep(time.Second * 5)
		w.Write(data)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err: %s", err.Error())
	}
}

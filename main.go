package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-storage/handler"
)

type Demo struct {
	Name string `json:"myName1"`
	Age  int    `json:"myAge1"`
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./assets/build")))
	// http.Handle("/static/",
	// 	http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// 静态资源处理
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHanlder)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("file/fastupload", handler.HTTPInterceptor(
		handler.TryFastUploadHandler,
	))
	http.HandleFunc("/file/mpupload/init", handler.HTTPInterceptor(
		handler.InitMultipartUploadHandler,
	))
	http.HandleFunc("/file/mpupload/uppart", handler.HTTPInterceptor(
		handler.UploadPartHandler,
	))
	http.HandleFunc("/file/mpupload/complete", handler.HTTPInterceptor(
		handler.CompleteUploadHandler,
	))
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get a ping!")
		d := Demo{
			Name: "Zhe Wang",
			Age:  34,
		}
		data, _ := json.Marshal(d)
		time.Sleep(time.Second * 5)
		w.Write(data)
	})
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))
	http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(handler.TryFastUploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err: %s", err.Error())
	}
}

package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/nicemayi/go-storage/meta"
	"github.com/nicemayi/go-storage/util"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, error: %s\n", err.Error())
			return
		}
		defer file.Close()

		currPath, _ := os.Getwd()
		filePath := currPath + "/uploaded_file/"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.Mkdir(filePath, os.ModePerm)
		}

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: filePath + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create data, error: %s\n", err.Error())
			return
		}
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data, error: %s\n", err.Error())
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHanlder(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success!")
}

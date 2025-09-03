package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func FirstHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./index.html")
}

func SecondHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "incorrent method", http.StatusInternalServerError)
		return
	}

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "ParseMultipartForm was failed", http.StatusInternalServerError)
		return
	}

	file, fileInfo, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "FormFile was failed", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileCht, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "ReadAll was failed", http.StatusInternalServerError)
		return
	}

	fileConv, err := service.Transform(string(fileCht))
	conv := []byte(fileConv)
	if err != nil {
		http.Error(res, "Conversion was failed", http.StatusInternalServerError)
		return
	}

	fileExt := filepath.Ext(fileInfo.Filename)
	fileName := fmt.Sprintf("%s%s", time.Now().UTC().String(), fileExt)

	if err := os.WriteFile(fileName, conv, 0755); err != nil {
		http.Error(res, "WriteFile was failed", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-type", "text/plain")
	res.WriteHeader(http.StatusOK)
	if _, err := res.Write(conv); err != nil {
		http.Error(res, "Write was failed", http.StatusInternalServerError)
		return
	}

}

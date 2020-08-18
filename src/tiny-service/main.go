package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TinyRes struct {
	Input struct{
		Size int64 `json:"size"`
		Type string `json:"type"`
	} `json:"input"`
	Output struct{
		Size int64 `json:"size"`
		Type string `json:"type"`
		Width int `json:"width"`
		Height int `json:"height"`
		Radio float32 `json:"radio"`
		Url string `json:"url"`
	} `json:"output"`
}
type ImageRes struct {
	*TinyRes
	Name string `json:"name"`
}
const COMPRESS_URL = "https://tinypng.com/web/shrink"
const CONTENT_TYPE = "application/x-www-form-urlencoded"

func doUnZip(zipFile, dest string) int {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println(err)
		return 500
	}
	defer reader.Close()
	for _, file := range reader.File {
		fmt.Println(file.Name)
		err := os.Mkdir(dest, 0755)
		f, err := os.Create(dest+"/"+file.Name)
		if err != nil {
			fmt.Println(err)
			return 500
		}
		rc, _ := file.Open()
		defer rc.Close()
		a, err := io.Copy(f, rc)
		if err != nil {
			fmt.Println(a)
			return 500
		}

	}
	return 200
}

func compressZip(coPath, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer zipFile.Close()
	zw := zip.NewWriter(zipFile)
	defer zw.Close()
	filepath.Walk(coPath, func(path string, info os.FileInfo, err error) error {
		if coPath == path {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
			fmt.Println("open file err", err)
		}
		header, _ := zip.FileInfoHeader(info)
		fmt.Println(header.Name)
		wh, _ := zw.CreateHeader(header)
		a, err := io.Copy(wh, file)
		defer file.Close()
		fmt.Println(a, err)
		return nil
	})
	return nil
}
func downloadFile(url, filename, dest string, c chan int) {
	_, err := os.Stat(dest)
	if err != nil {
		err := os.MkdirAll(dest, 0755)
		if err != nil {
			fmt.Println("downloadFile", err)
		}
	}
	if err != nil {
		os.Create(dest + ".zip")
	}
	resp, err := http.Get(url)
	defer resp.Body.Close()
	file, _ := os.Create(dest + "/" + filename)
	a, err := io.Copy(file, resp.Body)
	if err == nil {
		fmt.Println(a)
		c<-1
	}
}

func uploadFile(path string) (string, error) {
	startTime := time.Now()
	compressPath := strings.Replace(path, "origin", "compress", 1)
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	fileLength := len(files)
	uploadChan := make(chan TinyRes, fileLength)
	downloadChan := make(chan int, fileLength)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, file := range files {
		data, _ := ioutil.ReadFile(path+"/"+file.Name())
		go compress(data, uploadChan)
	}
	for i := 0; i < len(files); i++ {
		compressInfo := <-uploadChan
		fmt.Println("........compressInfo.......", compressInfo)
		go downloadFile(compressInfo.Output.Url, files[i].Name(), compressPath, downloadChan)
	}
	for i := 0; i < len(files); i++ {
		<-downloadChan
	}
	compressAddress := compressPath+".zip"
	compressZip(compressPath, compressAddress)
	elapsedTime := time.Since(startTime) / time.Millisecond
	fmt.Println("Segment finished in %dms", elapsedTime)
	return compressAddress, nil
}

func compress(data []byte, c chan TinyRes) error {
	resp, err := http.Post(COMPRESS_URL, CONTENT_TYPE, bytes.NewReader(data))
	if err != nil {
		fmt.Println("resp.....", err)
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	var res *TinyRes
	if err := json.Unmarshal([]byte(content), &res); err == nil {
		c <- *res
	} else {
		return err
		fmt.Println("jsopn.....", err)
	}
	return nil
}

func uploadImg(res http.ResponseWriter, req *http.Request) {
	var imgRes ImageRes
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Disposition")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	defer req.Body.Close()
	if err := req.ParseForm(); err != nil {
		fmt.Println(err)
		res.Write([]byte("出错了"))
		return
	}
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("出错了"))
		return
	}
	data, _ := ioutil.ReadAll(file)
	var compressChan = make(chan TinyRes)
	go compress(data, compressChan)
	resData := <-compressChan
	imgRes.TinyRes = &resData
	imgRes.Name = fileHeader.Filename
	jsonBody, err := json.Marshal(imgRes)
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("出错了"))
		return
	}
	fmt.Println(string(jsonBody))
	res.Write([]byte(string(jsonBody)))
}
func main() {
	http.HandleFunc("/uploadimg", uploadImg)
	http.HandleFunc("/upload", func(w http.ResponseWriter, f *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Disposition")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		if err := f.ParseForm(); err != nil {
			w.Write([]byte("no FormData"))
			fmt.Println("118........", err)
			return
		}
		file, fileHeader, err := f.FormFile("file")
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("上传错误"))
		}
		fileName := url.QueryEscape(fileHeader.Filename)
		w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
		originFilePath := "./origin/" + fileHeader.Filename
		if err != nil {
			fmt.Println("err", err)
			w.Write([]byte("no FormData"))
		}
		fmt.Println(file)
		fmt.Println(fileHeader.Filename)
		if strings.HasSuffix(fileHeader.Filename, ".zip") {
			_, err := os.Stat("./origin")
			if err != nil {
				os.Mkdir("./origin", 0755)
			}
			cfile, err := os.Create(originFilePath)
			if err != nil {
				fmt.Println("创建文件失败", err)
				w.Write([]byte("创建文件失败"))
			}
			if _, err := io.Copy(cfile, file); err == nil {
				fmt.Println("文件copy成功")
				unzipPath := fileHeader.Filename[0:len(fileHeader.Filename) - 4]
				code := doUnZip(originFilePath, "./origin/" + unzipPath)
				if code == 200 {
					addr, err := uploadFile("./origin/" + unzipPath)
					if err != nil {
						fmt.Println("程序出错")
						w.Write([]byte("程序出错"))
					} else {
						fmt.Println("处理完成")
						res, err := os.Open(addr)
						defer res.Close()
						if err != nil {
							fmt.Println("文件不存在")
							w.Write([]byte("程序出错"))
						} else {
							buff, _ := ioutil.ReadAll(res)
							w.Write(buff)
						}
					}
				}
			}
		} else {
			w.Write([]byte("请上传zip文件"))
		}
	})
	http.ListenAndServe(":7777", nil)
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

)
type transfer struct {
	File          *os.File
	Response      *http.Response
	ContentLength int
	Done          bool
}

func (transfer *transfer) startTransfer() {
	_, err := io.Copy(transfer.File, transfer.Response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		transfer.Done = true
		return
	}
	transfer.Done = true
}

func newTransfer(url string, fileName string) *transfer {

	transfer := new(transfer)

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	transfer.File = out

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	transfer.Response = response
	transfer.ContentLength, _ = strconv.Atoi(response.Header.Get("content-length"))
	transfer.Done = false

	return transfer
}

func (transfer *transfer) bytesTransfered() int {
	info, err := transfer.File.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return int(info.Size())
}

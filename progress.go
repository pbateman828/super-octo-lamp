package main

import (
	"fmt"
	"time"
)
type Progress struct {
	transfer *transfer
}

func (progress *Progress) start() {

	go progress.transfer.startTransfer()
	progress.show()


	progress.transfer.Response.Body.Close()
	progress.transfer.File.Close()
}

func (progress *Progress) show() {
	var downlod int
	totalBytes := int(progress.transfer.ContentLength)
	lastTime := false


	for !progress.transfer.Done || lastTime {

		fmt.Print("\r[")
		bytesDone := progress.transfer.bytesTransfered()
		downlod = 40 * bytesDone / totalBytes


		for i := 0; i < 40; i++ {
			if i < downlod {
				fmt.Print("=")
			} else if i == downlod {
				fmt.Print(">")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("] ")
		fmt.Printf("%d/%dkB", bytesDone/1000, totalBytes/1000)
		time.Sleep(100 * time.Millisecond)


		if progress.transfer.Done && !lastTime {
			lastTime = true
		} else {
			lastTime = false
		}
	}
	fmt.Println()
}

package main

func main() {

  // installs golang for you if you have a Linux machine
  transfer := newTransfer("https://dl.google.com/go/go1.9.3.linux-amd64.tar.gz",
		"go1.9.3.linux-amd64.tar.gz")
	progress := Progress{transfer}
	progress.start()
}

package main

import (
	"flag"
	"fmt"
	"github.com/haski007/files"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"sync"
)

type Input struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func main() {
	file := flag.String("file", "", "Path to file with json array with data to create qr code")
	flag.Parse()

	if *file == "" {
		logrus.Fatalf("flag 'file' is empty err: %s", flag.ErrHelp)
	}

	var inputData []*Input
	if err := files.ReadJson(*file, &inputData); err != nil {
		logrus.Fatalf("read json with input data err: %s", err)
	}

	wg := &sync.WaitGroup{}

	for _, data := range inputData {
		wg.Add(1)
		go genQrCode(data.Name, data.Data, wg)
	}

	wg.Wait()
	logrus.Println("Done!")
}


func genQrCode(name, data string, wg *sync.WaitGroup) {
	if err := qrcode.WriteFile(data, qrcode.Medium, 256, fmt.Sprintf("%s.png", name)); err != nil {
		logrus.Errorf("gen qr code and wrifile name: %s err: %s", name, err)
	}
	logrus.Printf("%s - qr code generated", name)
	wg.Done()
}
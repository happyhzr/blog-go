package util

import (
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	layout = "2006-01-02"
)

var (
	output *os.File
)

func Rotate(name string) {
	outputNew, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	CheckError(err)
	logrus.SetOutput(outputNew)
	if output != nil {
		err := output.Close()
		CheckError(err)
	}
	output = outputNew
	logrus.Infoln("rotate log")
}

func RunRotate() {
	for {
		t := time.Now()
		//name := "" // t.Format(layout)[:layoutPrefixLen]
		//path := "" // fmt.Sprintf("%s/%s.log", conf.LogPath, name)
		Rotate("")

		hourLeft := 24 - 1 - t.Hour()
		minLeft := 60 - 1 - t.Minute()
		secLeft := 60 - t.Second()
		ts := 60*(60*hourLeft+minLeft) + secLeft
		time.Sleep(time.Duration(ts) * time.Second)
	}
}

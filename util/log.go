package util

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/insisthzr/blog-back/conf"
)

const (
	layout       = "2006-01-02 15:04:05 +0800 CST"
	layoutPrefix = "2006-01-02"
)

var (
	layoutPrefixLen = len(layoutPrefix)
	output          *os.File
)

func RotateLogNX(path string) {
	outputNew, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	CheckError(err)
	logrus.SetOutput(outputNew)
	if output != nil {
		err := output.Close()
		CheckError(err)
	}
	output = outputNew
	logrus.Infoln("RotateLogNX")
}

//TODO HOW TO LOCK
func AutoRotateLog() {
	for {
		t := time.Now()
		name := t.Format(layout)[:layoutPrefixLen]
		path := fmt.Sprintf("%s/%s.log", conf.LogPath, name)
		RotateLogNX(path)

		hourLeft := 24 - 1 - t.Hour()
		minLeft := 60 - 1 - t.Minute()
		secLeft := 60 - t.Second()
		ts := 60*(24*hourLeft+minLeft) + secLeft
		time.Sleep(time.Duration(ts) * time.Second)
	}
}

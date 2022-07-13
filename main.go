package main

import (
	"math"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"ntsc.ac.cn/st-pcie-sync/pkg/driver"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	formatter := new(prefixed.TextFormatter)
	logrus.SetFormatter(formatter)
	devicePath := "/dev/txpci"
	d, err := driver.NewCardDriver(devicePath)
	if err != nil {
		panic(err)
	}
	if err := d.Open(); err != nil {
		panic(err)
	}
	for {
		cardTime, err := d.ReadTime()
		if err != nil {
			logrus.Errorf("failed to read card time: %v", err)
			time.Sleep(time.Second * 10)
			continue
		}

		tn := time.Now()
		offset := tn.Sub(cardTime)
		logrus.Debugf("offset: %s", offset)
		offset_f64 := math.Abs(float64(offset))
		conf_f64 := float64(time.Duration(
			time.Microsecond * time.Duration(10)))
		if offset_f64 < conf_f64 {
			time.Sleep(time.Second * 10)
			continue
		}
		logrus.Debugf("local      time: %s", tn.Format(time.RFC3339Nano))
		logrus.Debugf("read  card time: %s", cardTime.UTC().Format(time.RFC3339Nano))
		cardTime = cardTime.Add(time.Millisecond * -2)
		logrus.Debugf("fixed card time: %s", cardTime.UTC().Format(time.RFC3339Nano))
		tstr := cardTime.Format(time.RFC3339Nano)
		cmd := exec.Command("/usr/bin/date", "-s", tstr)
		if err := cmd.Run(); err != nil {
			logrus.Errorf("failed to set system time: %v", err)
		}
		logrus.Debugf("success to set system time: %s", tstr)
		time.Sleep(time.Second * 10)
	}
}

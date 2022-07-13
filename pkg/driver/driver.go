package driver

/*
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <errno.h>
#include <pthread.h>
#define TX_PCI_R_TM         _IOR('Z', 3, long)
typedef struct
{
	int num;
	int year;
	int month;
	int day;
	int hour;
	int minute;
	int second;
	int ns;
}TX_PCI_TIME;

int openDevice (char *path)
{
	int fd;
	fd = open(path,O_RDWR);
	return fd;
}

int tx_read_time(int fd,TX_PCI_TIME *cardTime)
{
        if(ioctl(fd, TX_PCI_R_TM, cardTime))
        {
                return -1;
        }
		return 0;
}

*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

type CardTime struct {
	Num    int32
	Year   int32
	Month  int32
	Day    int32
	Hour   int32
	Minute int32
	Second int32
	Ns     int32
}
type CardDriver struct {
	path string
	fd   int
}

func NewCardDriver(path string) (*CardDriver, error) {
	return &CardDriver{path: path}, nil
}

func (cd *CardDriver) Open() error {
	fd := C.openDevice(C.CString(cd.path))
	if fd < 0 {
		return fmt.Errorf("failed to open device [%s]", cd.path)
	}
	cd.fd = int(fd)
	return nil
}

func (cd *CardDriver) ReadTime() (time.Time, error) {
	var cardTime CardTime
	pt := (*C.TX_PCI_TIME)(unsafe.Pointer(&cardTime))
	r := C.tx_read_time(C.int(cd.fd), pt)
	if r == -1 {
		return time.Time{}, fmt.Errorf("ioctl TX_PCI_R_TM failed")
	}
	ts := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d.%03d%03d%03d+08:00",
		cardTime.Year, cardTime.Month, cardTime.Day,
		cardTime.Hour, cardTime.Minute, cardTime.Second,
		cardTime.Ns/1000000, (cardTime.Ns/1000)%1000, cardTime.Ns%1000)
	tc, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time [%s]: %v", ts, err)
	}
	return tc, nil
}

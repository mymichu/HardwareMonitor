package timeseriesdb

import (
	"os"
	"strconv"

	"../common"
)

type FileSettings struct {
	FileName string
	file     *os.File
	maxSize  int64
}

func (r *FileSettings) initialize() {
	var err error
	r.file, err = os.Create(r.FileName)
	if err != nil {
		r.file = nil
	}
}

func (r *FileSettings) writeCPUStateEntry(
	devID string,
	cpuState common.CpuState,
) error {
	var err error
	var fi os.FileInfo
	fi, err = r.file.Stat()
	if err == nil {
		if fi.Size() < r.maxSize {
			_, err = r.file.WriteString("cpu_state,device_id=" + devID +
				" usage=" + strconv.FormatFloat(float64(cpuState.Usage), 'f', 6, 32) +
				",temp=" + strconv.FormatFloat(float64(cpuState.Temp), 'f', 1, 32) +
				" " + strconv.FormatInt(cpuState.Timestamp.UTC().UnixNano(), 10) +
				"\n")
		}
	}
	return err
}

package timeseriesdb

import (
	"fmt"

	"../common"
)

type TimeSeriesDataBaseConfig struct {
	Address        string
	Username       string
	Password       string
	DataBaseName   string
	LogFileName    string
	EnableDataBase bool
	EnableLogFile  bool
	FileMaxSize    int64
}

type timeSeriesDataBase struct {
	deviceID               string
	fileWriter             FileSettings
	dataBaseWriter         InfluxConnection
	enabledWritingDatabase bool
	enabledWritingFile     bool
}

func (r *timeSeriesDataBase) writeCpuStateData(state <-chan common.CpuState) {
	for val := range state {

		fmt.Println(val)
		if r.enabledWritingFile {
			r.fileWriter.writeCPUStateEntry(r.deviceID, val)
		}
		if r.enabledWritingDatabase {
			error := r.dataBaseWriter.writeCPUStateEntry(r.deviceID, val)
			if error != nil {
				r.enabledWritingDatabase = false
				fmt.Println(error)
			}

		}
	}
}

//WriteData
func WriteData(deviceID string, config TimeSeriesDataBaseConfig, cpuState <-chan common.CpuState) {
	var tsdb timeSeriesDataBase
	tsdb.deviceID = deviceID
	tsdb.dataBaseWriter.Address = config.Address
	tsdb.dataBaseWriter.Username = config.Username
	tsdb.dataBaseWriter.Password = config.Password
	tsdb.dataBaseWriter.DataBaseName = config.DataBaseName
	tsdb.fileWriter.FileName = config.LogFileName
	tsdb.fileWriter.maxSize = config.FileMaxSize
	tsdb.enabledWritingDatabase = config.EnableDataBase
	tsdb.enabledWritingFile = config.EnableLogFile
	tsdb.fileWriter.initialize()
	errorConnection := tsdb.dataBaseWriter.initializeInfluxConnection()
	if errorConnection != nil {
		tsdb.enabledWritingDatabase = false
		fmt.Println(errorConnection)
	}
	go tsdb.writeCpuStateData(cpuState)
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"math/rand"

	"./common"
	"./configurator"
	"./httpserver"
	"./timeseriesdb"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"io/ioutil"
)

var timeSeriesDBConfig timeseriesdb.TimeSeriesDataBaseConfig
var httpServer httpserver.RESTServer

func Init() error {
	var appConfig configurator.Configurator
	error := appConfig.Init("assets/config.json")
	if error != nil {
		fmt.Println(error)
		return error
	}

	timeSeriesConfig := appConfig.GetTimeseriesConfig()
	timeSeriesDBConfig.Address = timeSeriesConfig.DataBase.Address
	timeSeriesDBConfig.Username = timeSeriesConfig.DataBase.Username
	timeSeriesDBConfig.Password = timeSeriesConfig.DataBase.Password
	timeSeriesDBConfig.DataBaseName = timeSeriesConfig.DataBase.Name
	timeSeriesDBConfig.LogFileName = timeSeriesConfig.LogFile.Path
	timeSeriesDBConfig.EnableDataBase = timeSeriesConfig.DataBase.Enable
	timeSeriesDBConfig.EnableLogFile = timeSeriesConfig.LogFile.Enable
	timeSeriesDBConfig.FileMaxSize = timeSeriesConfig.LogFile.MaxSizeByte

	httpServer.InitWeb(":82")
	return error
}

func readCPUStat(dataChannel chan common.CpuState) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	var cpuUsage float32
	if err != nil {
		cpuUsage = stat.CPUStatAll.Idle
	}
	else {
		cpuUsage = float32(rand.Intn(1000) / 10)
	}
	for {
		var data common.CpuState
		data.Usage = cpuUsage
		data.Temp = float32(rand.Intn(1000) / 10)
		data.Timestamp = time.Now()
		dataChannel <- data
		fmt.Println(data)
		time.Sleep(2 * time.Second)
	}
}

func broker(pub chan common.CpuState, sub1 chan common.CpuState, sub2 chan common.CpuState) {
	for val := range pub {
		sub1 <- val
		sub2 <- val
		time.Sleep(1 * time.Second)
	}
}

func main() {
	error := Init()
	if error == nil {
		//Init Turn-Off criteria
		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		cpuDataPub := make(chan common.CpuState)
		cpuDataSub1 := make(chan common.CpuState)
		cpuDataSub2 := make(chan common.CpuState)
		go readCPUStat(cpuDataPub)

		go broker(cpuDataPub, cpuDataSub1, cpuDataSub2)

		timeseriesdb.WriteData("T_1", timeSeriesDBConfig, cpuDataSub1)
		httpServer.UpdateCPUState(cpuDataSub2)

		httpServer.ListenAndServer()
		//runtime.NumActiveGoroutine()
		//Wait till gracefulStop gets a signal in channel
		sig := <-gracefulStop
		httpServer.Shutdown()
		fmt.Println("--- Turn off app ---")
		fmt.Printf("caught sig: %+v", sig)
	}

}

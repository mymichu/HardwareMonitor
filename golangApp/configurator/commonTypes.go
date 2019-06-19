package configurator

type TimeSeriesDataBase struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Enable   bool   `json:"enable"`
	Name     string `json:"name"`
}

type TimeSeriesLog struct {
	Path        string `json:"path"`
	Enable      bool   `json:"enable"`
	MaxSizeByte int64  `json:"maxSizeByte"`
}

type TimeSeries struct {
	DataBase TimeSeriesDataBase `json:"database"`
	LogFile  TimeSeriesLog      `json:"logFile"`
}

type Hosts struct {
	TimeSeries TimeSeries `json:"timeSeries"`
}

type Root struct {
	HostsSettings Hosts `json:"hosts"`
}

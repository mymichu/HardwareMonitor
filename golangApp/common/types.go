package common

import "time"

type CpuState struct {
	Usage     float32
	Temp      float32
	Timestamp time.Time
}

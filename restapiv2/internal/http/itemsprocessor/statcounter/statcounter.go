package statcounter

import (
	"restapiv2/internal/repository/stat"
)

type StatCounter struct {
	statStorage *stat.StatStorageType
}

func NewStatCounter(ss *stat.StatStorageType) *StatCounter{
	return &StatCounter {
		statStorage: ss,
	}
}
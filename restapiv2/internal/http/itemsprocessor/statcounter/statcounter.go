package statcounter

import (
	"restapiv2/internal/repository/statstorage"
)

type StatCounter struct {
	statStorage *statstorage.StatStorageType
}

func NewStatCounter(ss *statstorage.StatStorageType) *StatCounter{
	return &StatCounter {
		statStorage: ss,
	}
}
package statcounter

import (
	"restapiv2/internal/repository/statstorage"
)

type StatCounter struct {
	statStorage *statstorage.StatStorageType
}

func NewStatCounter() *StatCounter{
	return &StatCounter {
		statStorage: &statstorage.StatStorage,
	}
}

func (sc *StatCounter) Update(itemAction string) {
	sc.statStorage.Update(itemAction)
}
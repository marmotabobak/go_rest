package statstorage

import (
	"fmt"
	"restapiv2/pkg/utils"
	"sync"
)

type StatStorageType struct {
	storage map[string]int
	m sync.RWMutex
}

func NewStatStorageType() *StatStorageType {
	return &StatStorageType{
		storage: make(map[string]int),
		m: sync.RWMutex{},
	}
}

func (ss *StatStorageType) Update(itemAction string) {
	ss.m.Lock()
	ss.storage[itemAction]++
	ss.m.Unlock()
}

func (ss *StatStorageType) String() string {
	ss.m.RLock()
	res := fmt.Sprintf("%v", utils.SprintMapStringInt(ss.storage))
	ss.m.RUnlock()
	return res
}

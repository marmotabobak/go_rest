package statstorage

import (
	"fmt"
	"restapiv2/pkg/utils"
	"sync"
)

type StatStorage struct {
	storage map[string]int
	m       sync.RWMutex
}

func NewStatStorage() *StatStorage {
	return &StatStorage{
		storage: make(map[string]int),
		m:       sync.RWMutex{},
	}
}

func (ss *StatStorage) Update(itemAction string) {
	ss.m.Lock()
	ss.storage[itemAction]++
	ss.m.Unlock()
}

func (ss *StatStorage) String() string {
	ss.m.RLock()
	res := fmt.Sprintf("%v", utils.SprintMapStringInt(ss.storage))
	ss.m.RUnlock()
	return res
}

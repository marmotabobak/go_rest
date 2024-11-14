package statstorage

import (
	"fmt"
	"restapiv2/internal/mutex"
	"restapiv2/pkg/utils"
)

type StatStorageType map[string]int

var StatStorage StatStorageType = StatStorageType{}

func (s StatStorageType) Update(itemAction string) {
	mutex.M.Lock()
	s[itemAction]++
	mutex.M.Unlock()
}

func (s StatStorageType) String() string {
	mutex.M.RLock()
	res := fmt.Sprintf("%v", utils.SprintMapStringInt(StatStorage))
	mutex.M.RUnlock()
	return res
}

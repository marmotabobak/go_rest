package stat

import (
	"restapiv2/internal/mutex"
	"restapiv2/pkg/utils"
	"fmt"
)

type StatStorageType map[string]int

var StatStorage StatStorageType

func Init() {
	StatStorage = make(StatStorageType)
}

func (s StatStorageType) Update(itemAction string) {
	mutex.M.Lock()
	s[itemAction]++
	mutex.M.Unlock()
}

func (s StatStorageType) String() string {
	return fmt.Sprintf("%v", utils.SprintMapStringInt(StatStorage))
}




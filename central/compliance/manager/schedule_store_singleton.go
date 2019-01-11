package manager

import (
	"sync"

	"github.com/stackrox/rox/central/globaldb"
)

var (
	scheduleStoreInstance     ScheduleStore
	scheduleStoreInstanceInit sync.Once
)

// ScheduleStoreSingleton returns the singleton instance for the schedule store.
func ScheduleStoreSingleton() ScheduleStore {
	scheduleStoreInstanceInit.Do(func() {
		var err error
		scheduleStoreInstance, err = newScheduleStore(globaldb.GetGlobalDB())
		if err != nil {
			log.Panicf("Could not create schedule store: %v", err)
		}
	})
	return scheduleStoreInstance
}

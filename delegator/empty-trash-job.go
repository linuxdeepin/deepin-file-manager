package delegator

import (
	"pkg.linuxdeepin.com/lib/dbus"
	"pkg.linuxdeepin.com/lib/operations"
	"sync"
)

var (
	_EmptyTrashJobCount     uint64
	_EmptyTrashJobCountLock sync.Mutex
)

// EmptyTrashJob for dbus.
type EmptyTrashJob struct {
	dbusInfo dbus.DBusInfo
	op       *operations.EmptyTrashJob

	Trashing        func(string)
	Deleting        func(string)
	Done            func()
	ProcessedAmount func(int64, uint16)
	Aborted         func()
}

// GetDBusInfo returns dbus information.
func (job *EmptyTrashJob) GetDBusInfo() dbus.DBusInfo {
	return job.dbusInfo
}

// NewEmptyTrashJob creates a new EmptyTrashJob for dbus.
func NewEmptyTrashJob(shouldConfirmTrash bool, uiDelegate IUIDelegate) *EmptyTrashJob {
	_EmptyTrashJobCountLock.Lock()
	defer _EmptyTrashJobCountLock.Unlock()
	job := &EmptyTrashJob{
		dbusInfo: genDBusInfo("EmptyTrashJob", &_EmptyTrashJobCount),
		op:       operations.NewEmptyTrashJob(shouldConfirmTrash, uiDelegate),
	}
	_EmptyTrashJobCount++
	return job
}

// TODO:
func (job *EmptyTrashJob) listenSignals() {
}

func (job *EmptyTrashJob) executeJob() {
	defer dbus.UnInstallObject(job)
	job.op.Execute()
	dbus.Emit(job, "Done")
}

// Execute empty trash job.
func (job *EmptyTrashJob) Execute() {
	job.listenSignals()
	job.executeJob()
}
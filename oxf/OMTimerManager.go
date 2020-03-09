package oxf

type OMTimerManager struct {
	_timeouts OMEventQueue
}

func (t *OMTimerManager) Init() {
	t._timeouts.Init()
}

func (t *OMTimerManager) set(timeout IOxfTimeout) {
	t.setTimeoutDueTime(timeout)
	t._timeouts.add(timeout)
	go timeout.action()
}

func (t *OMTimerManager) action(timeout IOxfTimeout) {
	reactive := timeout.getDestination()
	t._timeouts.remove(timeout)
	if reactive != nil {
		reactive.Send(timeout)
	}
}

func (t *OMTimerManager) remove(timeout IOxfTimeout) {
	t._timeouts.remove(timeout)
}

func (*OMTimerManager) setTimeoutDueTime(timeout IOxfTimeout) {
	if timeout != nil {
		var tmLatency int64 = 0
		timeout.setDueTime(timeout.getDelayTime() + tmLatency)
	}
}

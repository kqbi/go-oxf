package oxf

type IOxfTimeout interface {
	IOxfEvent
	cancel()
	isCanceled() bool
	getDelayTime() int64
	setDelayTime(p_delayTime int64)
	getDueTime() int64
	setDueTime(p_dueTime int64)
	action()
}

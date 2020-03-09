package oxf

const OMTimeoutEventId int16 = -2
const OMReactiveTerminationEventId int16 = -9

type OMEvent struct {
	lId                int16
	destination        IOxfReactive
	deleteAfterConsume bool
	frameworkEvent     bool
}
func (e *OMEvent) Init() {

}

func (e *OMEvent) action() {

}

func (e *OMEvent) cancel() {

}

func (e *OMEvent) isCanceled() bool {
	return true
}

func (e *OMEvent) getDelayTime() int64 {
	return 0
}

func (e *OMEvent) setDelayTime(p_delayTime int64) {

}

func (e *OMEvent) getDueTime() int64 {
	return 0
}

func (e *OMEvent) setDueTime(p_dueTime int64) {

}

func (e *OMEvent) destroy() {

}

func (e *OMEvent) isFrameworkEvent() bool {
	return true
}

func (e *OMEvent) IsTypeOf(eventId int16) bool {
	return true
}

func (e *OMEvent) getId() int16 {
	return e.lId
}

func (e *OMEvent) SetId(p_id int16) {
	e.lId = p_id
}

func (e *OMEvent) getDestination() IOxfReactive {
	return e.destination
}

func (e *OMEvent) setDestination(p_IOxfReactive IOxfReactive) {
	e.destination = p_IOxfReactive
}

func (e *OMEvent) cleanUpRelations() {

}

func (e *OMEvent) setDeleteAfterConsume(p_deleteAfterConsume bool) {
	e.deleteAfterConsume = p_deleteAfterConsume
}

func (e *OMEvent) setFrameworkEvent(p_frameworkEvent bool) {
	e.frameworkEvent = p_frameworkEvent
}

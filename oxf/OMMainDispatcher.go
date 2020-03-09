package oxf

type OMMainDispatcher struct {
}

func (d *OMMainDispatcher) execute(ev IOxfEvent) {
	if ev != nil {
		result := EventConsumed
		result = d.dispatch(ev)

		if result == InstanceReachTerminate {

		}
	}
}
func (*OMMainDispatcher) dispatch(ev IOxfEvent) TakeEventStatus {
	result := EventNotConsumed
	dest := ev.getDestination()
	if dest != nil {
		dest.handleEvent(ev)
	}
	return result
}

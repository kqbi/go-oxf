package oxf

const OMStartBehaviorEventId int16 = -5
const OMStartBehaviorEvent_Events_Services_oxf_Design_id int16 = OMStartBehaviorEventId
const OMTimeoutDelayId int16 = -8

type OMStartBehaviorEvent struct {
	OMEvent
}

func (e *OMStartBehaviorEvent) Init() {
	e.SetId(OMStartBehaviorEvent_Events_Services_oxf_Design_id)
	e.setDeleteAfterConsume(false)
}

func (e *OMStartBehaviorEvent) isTypeOf(id int16) bool {
	return e.lId == id
}

func (e *OMStartBehaviorEvent) reincarnateAsTerminationEvent() {
	e.SetId(OMReactiveTerminationEventId)
	e.setFrameworkEvent(true)
}
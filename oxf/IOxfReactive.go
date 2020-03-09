package oxf

type TakeEventStatus int32

const (
	EventNotConsumed         TakeEventStatus = 0
	EventConsumed            TakeEventStatus = 1
	InstanceUnderDestruction TakeEventStatus = 2
	InstanceReachTerminate   TakeEventStatus = 3
)

type EventNotConsumedReason int32

const (
	StateMachineBusy            EventNotConsumedReason = 0
	EventNotHandledByStatechart EventNotConsumedReason = 1
)

type IOxfReactive interface {
	IOxfEventSender
	cancelTimeout(timeout IOxfTimeout) bool
	Destroy()
	endBehavior()
	handleEvent(ev IOxfEvent) TakeEventStatus
	popNullTransition()
	pushNullTransition()
	StartBehavior() bool
	handleNotConsumed(ev IOxfEvent, reason EventNotConsumedReason)
	handleTrigger(ev IOxfEvent)
	ScheduleTimeout(delay int64, targetStateName []byte) IOxfTimeout
	getActiveContext() IOxfActive
	SetActiveContext(p_IOxfActive IOxfActive)
	GetCurrentEvent() IOxfEvent
	//rootState_processEvent() TakeEventStatus
}

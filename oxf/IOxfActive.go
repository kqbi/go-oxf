package oxf

type IOxfActive interface {
	dispatch(ev IOxfEvent) TakeEventStatus
}

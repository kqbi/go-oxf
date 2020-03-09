package oxf

type IOxfEventSender interface {
	Send(ev IOxfEvent) bool
}

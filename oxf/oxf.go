package oxf

var Dispatcher *OMMainDispatcher
var timerManager *OMTimerManager
func init() {
	Dispatcher = new(OMMainDispatcher)
	timerManager = new(OMTimerManager)
	timerManager.Init()
}

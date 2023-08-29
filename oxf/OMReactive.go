package oxf

import (
	"reflect"
	"sync"
)

type Defaults int32

const (
	DEFAULT_MAX_NULL_STEPS               Defaults = 100
	terminateConnectorReachedStateMask   uint32   = 0x00010000
	deleteOnTerminateStateMask           uint32   = 0x00040000
	destroyEventResentStateMask          uint32   = 0x00200000
	underDestructionStateMask            uint32   = 0x00020000
	behaviorStartedStateMask             uint32   = 0x00100000
	nullTransitionMask                   uint32   = 0x0000FFFF
	shouldCompleteStartBehaviorStateMask uint32   = 0x00080000
)

type OMReactive struct {
	defaultStateMask             uint32
	nullTransitionStateMask      uint32               //## attribute nullTransitionStateMask
	nullTransitionMask           uint32               //## attribute nullTransitionMask
	maxNullSteps                 int32                //## attribute maxNullSteps
	state                        uint32               //## attribute state
	TheStartOrTerminationEvent   OMStartBehaviorEvent //## attribute theStartOrTerminationEvent
	active                       bool                 //## attribute active
	busy                         bool                 //## attribute busy
	supportDirectDeletion        bool                 //## attribute supportDirectDeletion
	globalSupportDirectDeletion  bool                 //## attribute globalSupportDirectDeletion
	supportRestartBehavior       bool                 //## attribute supportRestartBehavior
	globalSupportRestartBehavior bool                 //## attribute globalSupportRestartBehavior
	activeContext                IOxfActive           //## link activeContext
	_currentEvent                IOxfEvent            //## link currentEvent
	_mutex                       sync.Mutex           //## link eventGuard
	Parent                       interface{}
	Event                        chan IOxfEvent
}

//func (r *OMReactive) Init() {
//	r.TheStartOrTerminationEvent.Init()
//}

func (r *OMReactive) IsCurrentEvent(eventId int16) bool {
	return true
}

func (r *OMReactive) Cancel(timeout IOxfTimeout) {
	if timeout != nil {
		timeout.cancel()
		timeout = nil
	}
}

func (r *OMReactive) cancelEvents() {

}

func (r *OMReactive) cancelTimeout(timeout IOxfTimeout) bool {
	return true
}

func (r *OMReactive) Destroy() {
	r.setUnderDestruction()
	r.TheStartOrTerminationEvent.reincarnateAsTerminationEvent()
	r.Send(&r.TheStartOrTerminationEvent)
}

func (r *OMReactive) endBehavior() {

}

func (r *OMReactive) handleEvent(ev IOxfEvent) TakeEventStatus {
	status := EventNotConsumed

	if r.IsUnderDestruction() {
		// in termination process
		status = r.handleEventUnderDestruction(ev)
	} else {
		// check that the behavior should still run
		if r.ShouldTerminate() == true && r.shouldDelete() == true {
			status = EventConsumed
		} else {
			// the event guard is set -
			// use it to set mutual exclusion between Events and Triggered Operations
			// the m_eventGuard is set by the application when there is a guarded triggered operation
			//r._mutex.Lock()
			// actually handle the event
			status = r.processEvent(ev)
			// unlock the event guard
			//r._mutex.Unlock()

			// result with a status which indicates that the item had reached a terminate connector
			if r.ShouldTerminate() {
				status = InstanceReachTerminate
			}
		}
	}
	return status
}

func (r *OMReactive) handleEventNotQueued(ev IOxfEvent) {

}

func (r *OMReactive) popNullTransition() {

}

func (r *OMReactive) pushNullTransition() {

}

func (r *OMReactive) Send(ev IOxfEvent) bool {
	ev.Init()
	retCode := false
	if ev != nil {
		retCode = r.sendEvent(ev)
		if retCode == false {
			r.handleEventNotQueued(ev)
		}
	}
	return retCode
}

func (r *OMReactive) SetActiveContext(context IOxfActive) {
	//if r.context != getActiveContext() {
	//	setActive(activeInstance);
	//	setActiveContext(context);
	//}
	// Make sure we have a context
	if r.getActiveContext() == nil {
		// The fallback is that the object is dispatched by the system thread.
		r.activeContext = context
	}
}

func (r *OMReactive) setShouldTerminate(flag bool) {
	// if (flag) {
	// r.state |= terminateConnectorReachedStateMask
	// } else {
	// r.state &= ~terminateConnectorReachedStateMask
}

func (r *OMReactive) shouldDelete() bool {
	return (r.state & deleteOnTerminateStateMask) != 0
}

func (r *OMReactive) shouldSupportDirectDeletion() bool {
	return true
}

func (r *OMReactive) StartBehavior() {
	r.SetBehaviorStarted()
	// take the default transition
	//r.rootState_entDef()
	c := reflect.ValueOf(r.Parent)
	method := c.MethodByName("RootState_entDef")
	in := make([]reflect.Value, 0)
	method.Call(in)
	/*
		//status := false
		if r.IsUnderDestruction() {
			//	status = false
		} else {
			if r.IsBehaviorStarted() == false ||
				r.RestartBehaviorEnabled() == true {
				r.SetBehaviorStarted()
				// take the default transition
				//r.rootState_entDef()
				c := reflect.ValueOf(r.Parent)
				method := c.MethodByName("RootState_entDef")
				in := make([]reflect.Value, 0)
				method.Call(in)

				// This takes care of transitions without triggering events
				if r.ShouldCompleteRun() {
					// generate a dummy event in case the class doesn't receive any external events
					// this causes the runToCompletion() after the default transition to be taken -
					// in the class own thread (for active classes)
					r.SetCompleteStartBehavior(true)

					r.Send(&r.TheStartOrTerminationEvent)
				}
			}

			toTerminate := r.ShouldTerminate()
			if toTerminate {
				r.Destroy()
			}
			//status = !toTerminate
		}
		//return status
	*/
	for {
		select {
		case ev := <-r.Event:
			Dispatcher.execute(ev)
			toTerminate := r.IsUnderDestruction()
			if toTerminate {
				return
			}
		}
	}

}

func (r *OMReactive) handleNotConsumed(ev IOxfEvent, reason EventNotConsumedReason) {

}

func (r *OMReactive) handleTimeoutSetFailure(timeout IOxfTimeout) {

}

func (r *OMReactive) handleTrigger(ev IOxfEvent) {

}

func (r *OMReactive) hasWaitingNullTransitions() bool {
	return true
}

func (r *OMReactive) IsBehaviorStarted() bool {
	return r.state&behaviorStartedStateMask != 0
}

func (r *OMReactive) IsUnderDestruction() bool {
	return r.state&underDestructionStateMask != 0
}

func (r *OMReactive) processEvent(ev IOxfEvent) TakeEventStatus {
	if r.shouldCompleteStartBehavior() {
		r.SetCompleteStartBehavior(false)
		// protect from recursive Triggered Operation calls
		r.setBusy(true)
		r.runToCompletion()
		// end protection from recursive Triggered Operation calls
		r.setBusy(false)
	}

	// check that this is not the dummy OMStartBehaviorEvent event
	res := EventNotConsumed
	if ev.getId() != OMStartBehaviorEventId {
		if !r.isBusy() {
			// protect from recursive Triggered Operation calls
			r.setBusy(true)
			// Store the event in the OMReactive instance
			r.setCurrentEvent(ev)
			// consume the event

			c := reflect.ValueOf(r.Parent)
			method := c.MethodByName("RootState_processEvent")
			in := make([]reflect.Value, 0)
			res = method.Call(in)[0].Interface().(TakeEventStatus)

			// take null transitions (transitions without triggeres)
			if r.ShouldCompleteRun() {
				r.runToCompletion()
			}
			// notify unconsumed event
			if res == EventNotConsumed {
				r.handleNotConsumed(ev, EventNotHandledByStatechart)
			}
			// done with this event
			//  setCurrentEvent(NULL);
			// end protection from recursive Triggered Operation calls
			r.setBusy(false)
		} else {
			r.handleNotConsumed(ev, StateMachineBusy)
		}
	} else {
		// the start behavior event is consumed by taking the 0 transitions
		// therefore it is always consumed
		res = EventConsumed
	}

	return res
}

func (r *OMReactive) RestartBehaviorEnabled() bool {
	return r.supportRestartBehavior || r.globalSupportRestartBehavior
}

func (r *OMReactive) rootState_entDef() {

}

func (r *OMReactive) rootState_processEvent() TakeEventStatus {
	return 0
}

func (r *OMReactive) runToCompletion() {

}

func (r *OMReactive) ScheduleTimeout(delay int64, targetStateName []byte) IOxfTimeout {
	timeout := new(OMTimeout)
	timeout.Init()
	timeout.setDestination(r)
	timeout.setDelayTime(delay)
	timerManager.set(timeout)
	return timeout
}

func (r *OMReactive) SetBehaviorStarted() {
	r.state |= behaviorStartedStateMask
}

func (r *OMReactive) SetCompleteStartBehavior(flag bool) {
	if flag {
		r.state |= shouldCompleteStartBehaviorStateMask
	} else {
		r.state &= ^shouldCompleteStartBehaviorStateMask
	}
}

func (r *OMReactive) setCurrentEvent(ev IOxfEvent) {
	r._currentEvent = ev
}

func (r *OMReactive) setUnderDestruction() {
	r.state |= underDestructionStateMask
}

func (r *OMReactive) ShouldCompleteRun() bool {
	return r.state&nullTransitionMask != 0
}

func (r *OMReactive) shouldCompleteStartBehavior() bool {
	return (r.state & shouldCompleteStartBehaviorStateMask) != 0
}

func (r *OMReactive) ShouldTerminate() bool {
	return (r.state & terminateConnectorReachedStateMask) != 0
}

func (r *OMReactive) setDestroyEventResent() {
	r.state |= destroyEventResentStateMask
}

func (r *OMReactive) handleEventUnderDestruction(ev IOxfEvent) TakeEventStatus {
	if ev != nil &&
		ev.IsTypeOf(OMReactiveTerminationEventId) {
		if r.shouldDelete() {
			if r.isDestroyEventResent() {
			} else {
				r.Send(&r.TheStartOrTerminationEvent)
				r.setDestroyEventResent()
			}
		}
	}
	return InstanceUnderDestruction
}

func (r *OMReactive) isDestroyEventResent() bool {
	return (r.state & destroyEventResentStateMask) != 0
}

func (r *OMReactive) sendEvent(ev IOxfEvent) bool {
	result := false
	if r.IsUnderDestruction() == true &&
		ev.IsTypeOf(OMReactiveTerminationEventId) == false {
		// Destruction had begun,
		// ignore events
	} else {
		// Set the Receiver of the event
		//context := r.getActiveContext()
		if ev != nil && Dispatcher != nil {
			ev.setDestination(r)
			//boost::asio::post(((OMMainDispatcher*)context)->_ioc,
			//		boost::bind(&OMMainDispatcher::execute,
			//		(OMMainDispatcher*)context,
			//		ev));

			//go Dispatcher.execute(ev)
			r.Event <- ev
			result = true
		}
	}
	return result
}

func (r *OMReactive) getMaxNullSteps() int32 {
	return 0
}
func (r *OMReactive) setMaxNullSteps(p_maxNullSteps int32) {

}

func (r *OMReactive) isActive() {

}

func (r *OMReactive) getSupportDirectDeletion() bool {
	return true
}

func (r *OMReactive) setSupportDirectDeletion(p_supportDirectDeletion bool) {

}

func (r *OMReactive) getGlobalSupportDirectDeletion() bool {
	return true
}

func (r *OMReactive) setGlobalSupportDirectDeletion(p_globalSupportDirectDeletion bool) {

}

func (r *OMReactive) getSupportRestartBehavior() bool {
	return true
}

//## auto_generated
func (r *OMReactive) setSupportRestartBehavior(p_supportRestartBehavior bool) {

}

func (r *OMReactive) getGlobalSupportRestartBehavior() bool {
	return true
}

func (r *OMReactive) setGlobalSupportRestartBehavior(p_globalSupportRestartBehavior bool) {

}

func (r *OMReactive) getActiveContext() IOxfActive {
	return r.activeContext
}

func (r *OMReactive) GetCurrentEvent() IOxfEvent {
	return r._currentEvent
}
func (r *OMReactive) getReactiveInternalState() uint32 {
	return 0
}

func (r *OMReactive) setReactiveInternalState(p_state uint32) {

}

func (r *OMReactive) isBusy() bool {
	return r.busy
}

func (r *OMReactive) setBusy(p_busy bool) {
	r.busy = p_busy
}

func (r *OMReactive) cleanUpRelations() {

}

func (r *OMReactive) setActive(p_active bool) {

}

func (r *OMReactive) IS_EVENT_TYPE_OF(id int16) bool {
	if r.GetCurrentEvent() != nil {
		return (r.GetCurrentEvent()).IsTypeOf(id)
	} else {
		return false
	}
}

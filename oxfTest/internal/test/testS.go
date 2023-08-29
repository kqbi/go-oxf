package test

import (
	"fmt"

	"github.com/kqbi/go-oxf/oxf"
)

type testS_Enum uint16

const (
	OMNonState         int16 = 0
	terminationstate_3 int16 = 1
	state_5            int16 = 2
	Poll               int16 = 3
	Idle               int16 = 4
	End                int16 = 5
)

type TestS struct {
	oxf.OMReactive
	rootState_subState int16
	rootState_active   int16
	rootState_timeout  oxf.IOxfTimeout
	//#]
}

func (t *TestS) initStatechart() {
	t.rootState_subState = OMNonState
	t.rootState_active = OMNonState
	t.rootState_timeout = nil
}

func (t *TestS) Init() {
	//t.omreactive.Init()
	t.Event = make(chan oxf.IOxfEvent, 1)
	t.Parent = t
	t.SetActiveContext(oxf.Dispatcher)
	t.TheStartOrTerminationEvent.Init()
	t.initStatechart()
}

func (t *TestS) cancelTimeouts() {
	t.Cancel(t.rootState_timeout)
}

func (t *TestS) cancelTimeout(arg oxf.IOxfTimeout) bool {
	res := false
	if t.rootState_timeout == arg {
		t.rootState_timeout = nil
		res = true
	}
	return res
}

func (t *TestS) rootState_IN() bool {
	return true
}

func (t *TestS) rootState_isCompleted() bool {
	return t.terminationstate_3_IN()
}

func (t *TestS) terminationstate_3_IN() bool {
	return t.rootState_subState == terminationstate_3
}

func (t *TestS) state_5_IN() bool {
	return t.rootState_subState == state_5
}

func (t *TestS) Poll_IN() bool {
	return t.rootState_subState == Poll
}

func (t *TestS) Idle_IN() bool {
	return t.rootState_subState == Idle
}

func (t *TestS) End_IN() bool {
	return t.rootState_subState == End
}

func (t *TestS) RootState_entDef() {
	{
		t.rootState_subState = Idle
		t.rootState_active = Idle
		//#[ state Idle.(Entry)
		fmt.Print("Idle in\n")
		//#]
	}
}

func (t *TestS) RootState_processEvent() oxf.TakeEventStatus {
	res := oxf.EventNotConsumed
	switch t.rootState_active {
	// State Idle
	case Idle:
		{
			if t.IS_EVENT_TYPE_OF(evPoll_test_id) {
				//#[ state Idle.(Exit)
				fmt.Print("Idle out\n")
				//#]
				t.rootState_subState = Poll
				t.rootState_active = Poll
				//#[ state Poll.(Entry)
				fmt.Print("Poll in\n")
				//#]
				t.rootState_timeout = t.ScheduleTimeout(1000, nil)
				res = oxf.EventConsumed
			}

		}
		break
	// State Poll
	case Poll:
		{
			if t.IS_EVENT_TYPE_OF(oxf.OMTimeoutEventId) {
				if t.GetCurrentEvent() == t.rootState_timeout {
					t.Cancel(t.rootState_timeout)
					//#[ state Poll.(Exit)
					fmt.Print("Poll out\n")
					//#]
					t.rootState_subState = Poll
					t.rootState_active = Poll
					//#[ state Poll.(Entry)
					fmt.Print("Poll in\n")
					//#]
					t.rootState_timeout = t.ScheduleTimeout(1000, nil)
					res = oxf.EventConsumed
				}
			} else if t.IS_EVENT_TYPE_OF(evEnd_test_id) {
				t.Cancel(t.rootState_timeout)
				//#[ state Poll.(Exit)
				fmt.Print("Poll out\n")
				//#]
				t.rootState_subState = End
				t.rootState_active = End
				t.rootState_timeout = t.ScheduleTimeout(5000, nil)
				res = oxf.EventConsumed
			}

		}
		break
	// State End
	case End:
		{
			if t.IS_EVENT_TYPE_OF(oxf.OMTimeoutEventId) {
				if t.GetCurrentEvent() == t.rootState_timeout {
					t.Cancel(t.rootState_timeout)
					t.rootState_subState = state_5
					t.rootState_active = state_5
					res = oxf.EventConsumed
				}
			} else if t.IS_EVENT_TYPE_OF(evDestroy_test_id) {
				t.Cancel(t.rootState_timeout)
				t.rootState_subState = terminationstate_3
				t.rootState_active = terminationstate_3
				res = oxf.EventConsumed
			}

		}
		break
	// State state_5
	case state_5:
		{
			if t.IS_EVENT_TYPE_OF(evDestroy_test_id) {
				t.rootState_subState = terminationstate_3
				t.rootState_active = terminationstate_3
				res = oxf.EventConsumed
			}

		}
		break
	default:
		break
	}
	return res
}

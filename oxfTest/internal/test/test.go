package test

import "github.com/kqbi/go-oxf/oxf"

const (
	evPoll_test_id  int16  = 28001
	evDestroy_test_id int16  = 28002
	evEnd_test_id     int16 = 28003
)

type EvPoll struct {
	oxf.OMEvent
}

func (e *EvPoll) Init() {
	e.SetId(evPoll_test_id)
}

func (e *EvPoll) isTypeOf(id int16) bool {
	return evPoll_test_id == id
}

type EvDestroy struct {
	oxf.OMEvent
}

func (e *EvDestroy) Init() {
	e.SetId(evDestroy_test_id)
}

func (e *EvDestroy) isTypeOf(id int16) bool {
	return evDestroy_test_id == id
}

type EvEnd struct {
	oxf.OMEvent
}

func (e *EvEnd) Init() {
	e.SetId(evEnd_test_id)
}

func (e *EvEnd) isTypeOf(id int16) bool {
	return evEnd_test_id == id
}

/*********************************************************************
	File Path	: ..\..\untitled\test.h
*********************************************************************/

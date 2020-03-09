package oxf

import (
	"github.com/emirpasic/gods/sets/hashset"
	"sync"
)

type OMEventQueue struct {
	_mutex    sync.Mutex
	_theQueue *hashset.Set
}

func (q *OMEventQueue) Init() {
	q._theQueue = hashset.New()
}

func (q *OMEventQueue) Cleanup() {
	//#[ operation cleanup()
	q._mutex.Lock()
	q._theQueue.Clear()
	q._mutex.Unlock()
	//#]
}

func (q *OMEventQueue) get() IOxfTimeout {
	q._mutex.Lock()
	value := q._theQueue.Values()[0]
	q._theQueue.Remove(value)
	q._mutex.Unlock()
	return value.(IOxfTimeout)
}

func (q *OMEventQueue) isEmpty() bool {
	return q._theQueue.Empty()
}

func (q *OMEventQueue) add(ev IOxfTimeout) {
	q._mutex.Lock()
	q._theQueue.Add(ev)
	q._mutex.Unlock()
}

func (q *OMEventQueue) remove(ev IOxfTimeout) {
	q._mutex.Lock()
	q._theQueue.Remove(ev)
	q._mutex.Unlock()
}
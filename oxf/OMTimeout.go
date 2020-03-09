package oxf

import (
	"time"
)

type OMTimeout struct {
	OMEvent
	ti *time.Timer
	canceled  bool
	delayTime int64
	dueTime   int64
//	ch chan int
}

func (t *OMTimeout) Init() {
	t.lId = OMTimeoutEventId
	//t.canceled = false
	//t.ch = make(chan int)
}

func (t *OMTimeout) cancel() {
	t.setDestination(nil)
	t.canceled = true
}

func (t *OMTimeout) setDelayTimeout() {
	t.SetId(OMTimeoutDelayId)
}

func (t *OMTimeout) setRelativeDueTime(now int64) {
	t.dueTime = now + t.delayTime
}

func (t *OMTimeout) isCanceled() bool {
	return t.canceled
}

func (t *OMTimeout) getDelayTime() int64 {
	return t.delayTime
}

func (t *OMTimeout) setDelayTime(p_delayTime int64) {
	t.delayTime = p_delayTime
	t.ti = time.NewTimer(time.Duration(p_delayTime) * time.Millisecond)
}

func (t *OMTimeout) getDueTime() int64 {
	return t.dueTime
}

func (t *OMTimeout) setDueTime(p_dueTime int64) {
	t.dueTime = p_dueTime
}

func (t *OMTimeout) action() {
	//	_timer.expires_from_now(boost::posix_time::milliseconds(delayTime));
	//std::weak_ptr<IOxfTimeout> weakSelf = std::dynamic_pointer_cast<IOxfTimeout>(shared_from_this());
	//	_timer.async_wait([weakSelf](const boost::system::error_code&) {    //使用lambda表达式
	//		auto timeout = weakSelf.lock();
	//		if (timeout) {
	//			if (timeout->isCanceled()) {
	//				(std::dynamic_pointer_cast<OMTimeout>(timeout))->_tm.remove(timeout);
	//			} else {
	//				(std::dynamic_pointer_cast<OMTimeout>(timeout))->_tm.action(timeout);
	//			}
	//		}
	//	});
	defer t.ti.Stop()
	select {
	//case num := <-t.ch: //如果有数据，下面打印。但是有可能ch一直没数据
	case <-t.ti.C: //上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
		if t.isCanceled() {
			timerManager.remove(t)
		} else {
			timerManager.action(t)
		}
	}

}

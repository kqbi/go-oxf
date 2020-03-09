package oxf

//An event id attribute type
//## type ID
type IOxfEvent interface {
	Init()
	destroy()
	isFrameworkEvent() bool
	IsTypeOf(eventId int16) bool
	getId() int16
	SetId(p_id int16)
	getDestination() IOxfReactive
	setDestination(p_IOxfReactive IOxfReactive)
}

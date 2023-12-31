// Code generated by "stringer -type=EventType -trimprefix=Event"; DO NOT EDIT.

package grabbit

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EventUp-0]
	_ = x[EventDown-1]
	_ = x[EventCannotEstablish-2]
	_ = x[EventBlocked-3]
	_ = x[EventUnBlocked-4]
	_ = x[EventClosed-5]
	_ = x[EventMessageReceived-6]
	_ = x[EventMessagePublished-7]
	_ = x[EventMessageReturned-8]
	_ = x[EventConfirm-9]
	_ = x[EventQos-10]
	_ = x[EventConsume-11]
	_ = x[EventDefineTopology-12]
	_ = x[EventDataExhausted-13]
	_ = x[EventDataPartial-14]
}

const _EventType_name = "UpDownCannotEstablishBlockedUnBlockedClosedMessageReceivedMessagePublishedMessageReturnedConfirmQosConsumeDefineTopologyDataExhaustedDataPartial"

var _EventType_index = [...]uint8{0, 2, 6, 21, 28, 37, 43, 58, 74, 89, 96, 99, 106, 120, 133, 144}

func (i EventType) String() string {
	if i < 0 || i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}

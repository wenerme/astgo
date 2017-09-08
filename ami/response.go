package ami

// Raised in response to a Status command.
//
// seealso Status
type StatusCompleteResponse struct {
	// Number of Status events returned
	Items string
}

// Raised in response to a Status command.
//
// seealso Status
type StatusResponse struct {
	// Type of channel
	Type string
	// Dialed number identifier
	DNID string
	// Absolute lifetime of the channel
	TimeToHangup string
	// Identifier of the bridge the channel is in, may be empty if not in one
	BridgeID string
	Linkedid string
	// Application currently executing on the channel
	Application string
	// Data given to the currently executing channel
	Data string
	// Media formats the connected party is willing to send or receive
	Nativeformats string
	// Media formats that frames from the channel are received in
	Readformat string
	// Translation path for media received in native formats
	Readtrans string
	// Media formats that frames to the channel are accepted in
	Writeformat string
	// Translation path for media sent to the connected party
	Writetrans string
	// Configured call group on the channel
	Callgroup string
	// Configured pickup group on the channel
	Pickupgroup string
	// Number of seconds the channel has been active
	Seconds string
}

type CoreShowChannelsCompleteResponse CoreShowChannelsCompleteEvent

type CoreShowChannelResponse CoreShowChannelEvent

// Information about a bridge.
type BridgeInfoCompleteResponse struct {
}

// Information about a channel in a bridge.
type BridgeInfoChannelResponse struct {
}

// Indicates the end of the list the current known extension states.
type ExtensionStateListCompleteResponse struct {
	// Conveys the status of the event list.
	EventList string
	// Conveys the number of statuses reported.
	ListItems string
}

type ExtensionStatusResponse ExtensionStatusEvent

// Indicates the end of the list the current known extension states.
type DeviceStateListCompleteResponse struct {
	// Conveys the status of the event list.
	EventList string
	// Conveys the number of statuses reported.
	ListItems string
}

type DeviceStateChangeResponse DeviceStateChangeEvent

// Indicates the end of the list the current known extension states.
type PresenceStateListCompleteResponse struct {
	// Conveys the status of the event list.
	EventList string
	// Conveys the number of statuses reported.
	ListItems string
}

type PresenceStateChangeResponse PresenceStateChangeEvent

// Provide final information about an endpoint list.
type EndpointListCompleteResponse struct {
	EventList string
	ListItems string
}

type EndpointListResponse EndpointListEvent

// Provide final information about endpoint details.
type EndpointDetailCompleteResponse struct {
	EventList string
	ListItems string
}

type EndpointDetailResponse EndpointDetailEvent

type IdentifyDetailResponse IdentifyDetailEvent

type ContactStatusDetailResponse ContactStatusDetailEvent

type AuthDetailResponse AuthDetailEvent

type TransportDetailResponse TransportDetailEvent

type AorDetailResponse AorDetailEvent

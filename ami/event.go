package ami

// Raised when an alarm is cleared on a DAHDI channel.
type AlarmClearEvent struct {
	// The DAHDI channel on which the alarm was cleared.
	// NOTE: This is not an Asterisk channel identifier.
	DAHDIChannel string
}

// Raised when an alarm is cleared on a DAHDI span.
type SpanAlarmClearEvent struct {
	// The span on which the alarm was cleared.
	Span string
}

// Raised when the Do Not Disturb state is changed on a DAHDI channel.
type DNDStateEvent struct {
	// The DAHDI channel on which DND status changed.
	// NOTE: This is not an Asterisk channel identifier.
	DAHDIChannel string
	//
	// ENUM: enabled,disabled,
	Status string
}

// Raised when an alarm is set on a DAHDI channel.
type AlarmEvent struct {
	// The channel on which the alarm occurred.
	// NOTE: This is not an Asterisk channel identifier.
	DAHDIChannel string
	// A textual description of the alarm that occurred.
	Alarm string
}

// Raised when an alarm is set on a DAHDI span.
type SpanAlarmEvent struct {
	// The span on which the alarm occurred.
	Span string
	// A textual description of the alarm that occurred.
	Alarm string
}

// Raised when a DAHDI channel is created or an underlying technology is associated
// with a DAHDI channel.
type DAHDIChannelEvent struct {
	// The DAHDI span associated with this channel.
	DAHDISpan string
	// The DAHDI channel associated with this channel.
	DAHDIChannel string
}

// Raised when SIPQualifyPeer has finished qualifying the specified peer.
//
// seealso SIPqualifypeer
type SIPQualifyPeerDoneEvent struct {
	// The name of the peer.
	Peer string
}

// Raised when a SIP session times out.
type SessionTimeoutEvent struct {
	// The source of the session timeout.
	//
	// ENUM: RTPTimeout,SIPSessionTimer,
	Source string
}

// Published when a malicious call ID request arrives.
type MCIDEvent struct {
	MCallerIDNumValid       string
	MCallerIDNum            string
	MCallerIDton            string
	MCallerIDNumPlan        string
	MCallerIDNumPres        string
	MCallerIDNameValid      string
	MCallerIDName           string
	MCallerIDNameCharSet    string
	MCallerIDNamePres       string
	MCallerIDSubaddr        string
	MCallerIDSubaddrType    string
	MCallerIDSubaddrOdd     string
	MCallerIDPres           string
	MConnectedIDNumValid    string
	MConnectedIDNum         string
	MConnectedIDton         string
	MConnectedIDNumPlan     string
	MConnectedIDNumPres     string
	MConnectedIDNameValid   string
	MConnectedIDName        string
	MConnectedIDNameCharSet string
	MConnectedIDNamePres    string
	MConnectedIDSubaddr     string
	MConnectedIDSubaddrType string
	MConnectedIDSubaddrOdd  string
	MConnectedIDPres        string
}

// Response event in a series to the Agents AMI action containing
// information about a defined agent.
// The channel snapshot is present if the Status value is `AGENT_IDLE` or `AGENT_ONCALL`.
//
// seealso Agents
type AgentsEvent struct {
	// Agent ID of the agent.
	Agent string
	// User friendly name of the agent.
	Name string
	// Current status of the agent.
	// The valid values are:
	//
	// ENUM: AGENT_LOGGEDOFF,AGENT_IDLE,AGENT_ONCALL,
	Status string
	// `BRIDGEPEER` value on agent channel.
	// Present if Status value is `AGENT_ONCALL`.
	TalkingToChan string
	// Epoche time when the agent started talking with the caller.
	// Present if Status value is `AGENT_ONCALL`.
	CallStarted string
	// Epoche time when the agent logged in.
	// Present if Status value is `AGENT_IDLE` or `AGENT_ONCALL`.
	LoggedInTime string
}

// Final response event in a series of events to the Agents AMI action.
//
// seealso Agents
type AgentsCompleteEvent struct {
}

// Raised when a user joins a MeetMe conference.
//
// seealso MeetmeLeave
// seealso MeetMe
type MeetmeJoinEvent struct {
	// The identifier for the MeetMe conference.
	Meetme string
	// The identifier of the MeetMe user who joined.
	Usernum string
}

// Raised when a user leaves a MeetMe conference.
//
// seealso MeetmeJoin
type MeetmeLeaveEvent struct {
	// The identifier for the MeetMe conference.
	Meetme string
	// The identifier of the MeetMe user who joined.
	Usernum string
	// The length of time in seconds that the Meetme user was in the conference.
	Duration string
}

// Raised when a MeetMe conference ends.
//
// seealso MeetmeJoin
type MeetmeEndEvent struct {
	// The identifier for the MeetMe conference.
	Meetme string
}

// Raised when a MeetMe user has started talking.
type MeetmeTalkRequestEvent struct {
	// The identifier for the MeetMe conference.
	Meetme string
	// The identifier of the MeetMe user who joined.
	Usernum string
	// The length of time in seconds that the Meetme user has been in the conference at the time of this event.
	Duration string
	//
	// ENUM: on,off,
	Status string
}

// Raised when a MeetMe user begins or ends talking.
type MeetmeTalkingEvent struct {
	// The identifier for the MeetMe conference.
	Meetme string
	// The identifier of the MeetMe user who joined.
	Usernum string
	// The length of time in seconds that the Meetme user has been in the conference at the time of this event.
	Duration string
	//
	// ENUM: on,off,
	Status string
}

// Raised when a MeetMe user is muted or unmuted.
type MeetmeMuteEvent struct {
	// The identifier for the MeetMe conference.
	Meetme string
	// The identifier of the MeetMe user who joined.
	Usernum string
	// The length of time in seconds that the Meetme user has been in the conference at the time of this event.
	Duration string
	//
	// ENUM: on,off,
	Status string
}

// Raised when a notification is sent out by a MiniVoiceMail application
type MiniVoiceMailEvent struct {
	// What action was taken. Currently, this will always be `SentNotification`
	Action string
	// The mailbox that the notification was about, specified as `mailbox`@`context`
	Mailbox string
	// A message counter derived from the `MVM_COUNTER` channel variable.
	Counter string
}

// Raised when a Queue member's status has changed.
type QueueMemberStatusEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// Channel technology or location from which to read device state changes.
	StateInterface string
	//
	// ENUM: dynamic,realtime,static,
	Membership string
	// The penalty associated with the queue member.
	Penalty string
	// The number of calls this queue member has serviced.
	CallsTaken string
	// The time this member last took a call, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	LastCall string
	// The time when started last paused the queue member.
	LastPause string
	// Set to 1 if member is in call. Set to 0 after LastCall time is updated.
	//
	// ENUM: 0,1,
	InCall string
	// The numeric device state status of the queue member.
	//
	// ENUM: 0,1,2,3,4,5,6,7,8,
	Status string
	//
	// ENUM: 0,1,
	Paused string
	// If set when paused, the reason the queue member was paused.
	PausedReason string
	//
	// ENUM: 0,1,
	Ringinuse string
}

// Raised when a member is added to the queue.
//
// seealso QueueMemberRemoved
// seealso AddQueueMember
type QueueMemberAddedEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// Channel technology or location from which to read device state changes.
	StateInterface string
	//
	// ENUM: dynamic,realtime,static,
	Membership string
	// The penalty associated with the queue member.
	Penalty string
	// The number of calls this queue member has serviced.
	CallsTaken string
	// The time this member last took a call, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	LastCall string
	// The time when started last paused the queue member.
	LastPause string
	// Set to 1 if member is in call. Set to 0 after LastCall time is updated.
	//
	// ENUM: 0,1,
	InCall string
	// The numeric device state status of the queue member.
	//
	// ENUM: 0,1,2,3,4,5,6,7,8,
	Status string
	//
	// ENUM: 0,1,
	Paused string
	// If set when paused, the reason the queue member was paused.
	PausedReason string
	//
	// ENUM: 0,1,
	Ringinuse string
}

// Raised when a member is removed from the queue.
//
// seealso QueueMemberAdded
// seealso RemoveQueueMember
type QueueMemberRemovedEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// Channel technology or location from which to read device state changes.
	StateInterface string
	//
	// ENUM: dynamic,realtime,static,
	Membership string
	// The penalty associated with the queue member.
	Penalty string
	// The number of calls this queue member has serviced.
	CallsTaken string
	// The time this member last took a call, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	LastCall string
	// The time when started last paused the queue member.
	LastPause string
	// Set to 1 if member is in call. Set to 0 after LastCall time is updated.
	//
	// ENUM: 0,1,
	InCall string
	// The numeric device state status of the queue member.
	//
	// ENUM: 0,1,2,3,4,5,6,7,8,
	Status string
	//
	// ENUM: 0,1,
	Paused string
	// If set when paused, the reason the queue member was paused.
	PausedReason string
	//
	// ENUM: 0,1,
	Ringinuse string
}

// Raised when a member is paused/unpaused in the queue.
//
// seealso PauseQueueMember
// seealso UnPauseQueueMember
type QueueMemberPauseEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// Channel technology or location from which to read device state changes.
	StateInterface string
	//
	// ENUM: dynamic,realtime,static,
	Membership string
	// The penalty associated with the queue member.
	Penalty string
	// The number of calls this queue member has serviced.
	CallsTaken string
	// The time this member last took a call, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	LastCall string
	// The time when started last paused the queue member.
	LastPause string
	// Set to 1 if member is in call. Set to 0 after LastCall time is updated.
	//
	// ENUM: 0,1,
	InCall string
	// The numeric device state status of the queue member.
	//
	// ENUM: 0,1,2,3,4,5,6,7,8,
	Status string
	//
	// ENUM: 0,1,
	Paused string
	// If set when paused, the reason the queue member was paused.
	PausedReason string
	//
	// ENUM: 0,1,
	Ringinuse string
	// The reason a member was paused.
	Reason string
}

// Raised when a member's penalty is changed.
//
// seealso QUEUE_MEMBER
type QueueMemberPenaltyEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// Channel technology or location from which to read device state changes.
	StateInterface string
	//
	// ENUM: dynamic,realtime,static,
	Membership string
	// The penalty associated with the queue member.
	Penalty string
	// The number of calls this queue member has serviced.
	CallsTaken string
	// The time this member last took a call, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	LastCall string
	// The time when started last paused the queue member.
	LastPause string
	// Set to 1 if member is in call. Set to 0 after LastCall time is updated.
	//
	// ENUM: 0,1,
	InCall string
	// The numeric device state status of the queue member.
	//
	// ENUM: 0,1,2,3,4,5,6,7,8,
	Status string
	//
	// ENUM: 0,1,
	Paused string
	// If set when paused, the reason the queue member was paused.
	PausedReason string
	//
	// ENUM: 0,1,
	Ringinuse string
}

// Raised when a member's ringinuse setting is changed.
//
// seealso QUEUE_MEMBER
type QueueMemberRinginuseEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// Channel technology or location from which to read device state changes.
	StateInterface string
	//
	// ENUM: dynamic,realtime,static,
	Membership string
	// The penalty associated with the queue member.
	Penalty string
	// The number of calls this queue member has serviced.
	CallsTaken string
	// The time this member last took a call, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	LastCall string
	// The time when started last paused the queue member.
	LastPause string
	// Set to 1 if member is in call. Set to 0 after LastCall time is updated.
	//
	// ENUM: 0,1,
	InCall string
	// The numeric device state status of the queue member.
	//
	// ENUM: 0,1,2,3,4,5,6,7,8,
	Status string
	//
	// ENUM: 0,1,
	Paused string
	// If set when paused, the reason the queue member was paused.
	PausedReason string
	//
	// ENUM: 0,1,
	Ringinuse string
}

// Raised when a caller joins a Queue.
//
// seealso QueueCallerLeave
// seealso Queue
type QueueCallerJoinEvent struct {
	// The name of the queue.
	Queue string
	// This channel's current position in the queue.
	Position string
	// The total number of channels in the queue.
	Count string
}

// Raised when a caller leaves a Queue.
//
// seealso QueueCallerJoin
type QueueCallerLeaveEvent struct {
	// The name of the queue.
	Queue string
	// The total number of channels in the queue.
	Count string
	// This channel's current position in the queue.
	Position string
}

// Raised when a caller abandons the queue.
type QueueCallerAbandonEvent struct {
	// The name of the queue.
	Queue string
	// This channel's current position in the queue.
	Position string
	// The channel's original position in the queue.
	OriginalPosition string
	// The time the channel was in the queue, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	HoldTime string
}

// Raised when an queue member is notified of a caller in the queue.
//
// seealso AgentRingNoAnswer
// seealso AgentComplete
// seealso AgentConnect
type AgentCalledEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
}

// Raised when a queue member is notified of a caller in the queue and fails to
// answer.
//
// seealso AgentCalled
type AgentRingNoAnswerEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// The time the queue member was rung, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	RingTime string
}

// Raised when a queue member has finished servicing a caller in the queue.
//
// seealso AgentCalled
// seealso AgentConnect
type AgentCompleteEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// The time the channel was in the queue, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	HoldTime string
	// The time the queue member talked with the caller in the queue, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	TalkTime string
	//
	// ENUM: caller,agent,transfer,
	Reason string
}

// Raised when a queue member hangs up on a caller in the queue.
//
// seealso AgentCalled
// seealso AgentConnect
type AgentDumpEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
}

// Raised when a queue member answers and is bridged to a caller in the queue.
//
// seealso AgentCalled
// seealso AgentComplete
// seealso AgentDump
type AgentConnectEvent struct {
	// The name of the queue.
	Queue string
	// The name of the queue member.
	MemberName string
	// The queue member's channel technology or location.
	Interface string
	// The time the queue member was rung, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	RingTime string
	// The time the channel was in the queue, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	HoldTime string
}

// Raised when a variable local to the gosub stack frame is set due to a subroutine
// call.
//
// seealso GoSub
// seealso gosub
// seealso LOCAL
// seealso LOCAL_PEEK
type VarSetEvent struct {
	// The LOCAL variable being set.
	// NOTE: The variable name will always be enclosed with `LOCAL()`
	Variable string
	// The new value of the variable.
	Value string
}

// Raised when a conference starts.
//
// seealso ConfbridgeEnd
// seealso ConfBridge
type ConfbridgeStartEvent struct {
	// The name of the Confbridge conference.
	Conference string
}

// Raised when a conference ends.
//
// seealso ConfbridgeStart
// seealso ConfBridge
type ConfbridgeEndEvent struct {
	// The name of the Confbridge conference.
	Conference string
}

// Raised when a channel joins a Confbridge conference.
//
// seealso ConfbridgeLeave
// seealso ConfBridge
type ConfbridgeJoinEvent struct {
	// The name of the Confbridge conference.
	Conference string
	// Identifies this user as an admin user.
	//
	// ENUM: Yes,No,
	Admin string
}

// Raised when a channel leaves a Confbridge conference.
//
// seealso ConfbridgeJoin
// seealso ConfBridge
type ConfbridgeLeaveEvent struct {
	// The name of the Confbridge conference.
	Conference string
	// Identifies this user as an admin user.
	//
	// ENUM: Yes,No,
	Admin string
}

// Raised when a conference starts recording.
//
// seealso ConfbridgeStopRecord
// seealso ConfBridge
type ConfbridgeRecordEvent struct {
	// The name of the Confbridge conference.
	Conference string
}

// Raised when a conference that was recording stops recording.
//
// seealso ConfbridgeRecord
// seealso ConfBridge
type ConfbridgeStopRecordEvent struct {
	// The name of the Confbridge conference.
	Conference string
}

// Raised when a Confbridge participant mutes.
//
// seealso ConfbridgeUnmute
// seealso ConfBridge
type ConfbridgeMuteEvent struct {
	// The name of the Confbridge conference.
	Conference string
	// Identifies this user as an admin user.
	//
	// ENUM: Yes,No,
	Admin string
}

// Raised when a confbridge participant unmutes.
//
// seealso ConfbridgeMute
// seealso ConfBridge
type ConfbridgeUnmuteEvent struct {
	// The name of the Confbridge conference.
	Conference string
	// Identifies this user as an admin user.
	//
	// ENUM: Yes,No,
	Admin string
}

// Raised when a confbridge participant begins or ends talking.
//
// seealso ConfBridge
type ConfbridgeTalkingEvent struct {
	// The name of the Confbridge conference.
	Conference string
	//
	// ENUM: on,off,
	TalkingStatus string
	// Identifies this user as an admin user.
	//
	// ENUM: Yes,No,
	Admin string
}

// Raised when a CDR is generated.
// The Cdr event is only raised when the cdr_manager backend is loaded and registered with the CDR engine.
// NOTE:  This event can contain additional fields depending on the configuration provided by cdr_manager.conf.
type CdrEvent struct {
	// The account code of the Party A channel.
	AccountCode string
	// The Caller ID number associated with the Party A in the CDR.
	Source string
	// The dialplan extension the Party A was executing.
	Destination string
	// The dialplan context the Party A was executing.
	DestinationContext string
	// The Caller ID name associated with the Party A in the CDR.
	CallerID string
	// The channel name of the Party A.
	Channel string
	// The channel name of the Party B.
	DestinationChannel string
	// The last dialplan application the Party A executed.
	LastApplication string
	// The parameters passed to the last dialplan application the Party A executed.
	LastData string
	// The time the CDR was created.
	StartTime string
	// The earliest of either the time when Party A answered, or the start time of this CDR.
	AnswerTime string
	// The time when the CDR was finished. This occurs when the Party A hangs up or when the bridge between Party A and Party B is broken.
	EndTime string
	// The time, in seconds, of EndTime - StartTime.
	Duration string
	// The time, in seconds, of AnswerTime - StartTime.
	BillableSeconds string
	// The final known disposition of the CDR.
	//
	// ENUM: NO ANSWER,FAILED,BUSY,ANSWERED,CONGESTION,
	Disposition string
	// A flag that informs a billing system how to treat the CDR.
	//
	// ENUM: OMIT,BILLING,DOCUMENTATION,
	AMAFlags string
	// A unique identifier for the Party A channel.
	UniqueID string
	// A user defined field set on the channels. If set on both the Party A and Party B channel, the userfields of both are concatenated and separated by a `;`.
	UserField string
}

// Raised when a Channel Event Log is generated for a channel.
type CELEvent struct {
	// The name of the CEL event being raised. This can include both the system defined CEL events, as well as user defined events.
	// NOTE: All events listed here may not be raised, depending on the configuration in cel.conf.
	//
	// ENUM: CHAN_START,CHAN_END,ANSWER,HANGUP,BRIDGE_ENTER,BRIDGE_EXIT,APP_START,APP_END,PARK_START,PARK_END,BLINDTRANSFER,ATTENDEDTRANSFER,PICKUP,FORWARD,LINKEDID_END,LOCAL_OPTIMIZE,USER_DEFINED,
	EventName string
	// The channel's account code.
	AccountCode string
	// The Caller ID number.
	CallerIDnum string
	// The Caller ID name.
	CallerIDname string
	// The Caller ID Automatic Number Identification.
	CallerIDani string
	// The Caller ID Redirected Dialed Number Identification Service.
	CallerIDrdnis string
	// The Caller ID Dialed Number Identifier.
	CallerIDdnid string
	// The dialplan extension the channel is currently executing in.
	Exten string
	// The dialplan context the channel is currently executing in.
	Context string
	// The dialplan application the channel is currently executing.
	Application string
	// The arguments passed to the dialplan Application.
	AppData string
	// The time the CEL event occurred.
	EventTime string
	// A flag that informs a billing system how to treat the CEL.
	//
	// ENUM: OMIT,BILLING,DOCUMENTATION,
	AMAFlags string
	// The unique ID of the channel.
	UniqueID string
	// The linked ID of the channel, which ties this event to other related channel's events.
	LinkedID string
	// A user defined field set on a channel, containing arbitrary application specific data.
	UserField string
	// If this channel is in a bridge, the channel that it is in a bridge with.
	Peer string
	// If this channel is in a bridge, the accountcode of the channel it is in a bridge with.
	PeerAccount string
	// Some events will have event specific data that accompanies the CEL record. This extra data is JSON encoded, and is dependent on the event in question.
	Extra string
}

// Raised when an Advice of Charge message is sent at the beginning of a call.
//
// seealso AOC-D
// seealso AOC-E
type AOC_SEvent struct {
	Chargeable string
	//
	// ENUM: NotAvailable,Free,FreeFromBeginning,Duration,Flag,Volume,SpecialCode,
	RateType string
	Currency string
	Name     string
	Cost     string
	//
	// ENUM: 1/1000,1/100,1/10,1,10,100,1000,
	Multiplier   string
	ChargingType string
	StepFunction string
	Granularity  string
	Length       string
	Scale        string
	//
	// ENUM: Octect,Segment,Message,
	Unit        string
	SpecialCode string
}

func (AOC_SEvent) OriginalName() string { return "AOC-S" }

// Raised when an Advice of Charge message is sent during a call.
//
// seealso AOCMessage
// seealso AOC-S
// seealso AOC-E
type AOC_DEvent struct {
	Charge string
	//
	// ENUM: NotAvailable,Free,Currency,Units,
	Type string
	//
	// ENUM: Normal,Reverse,CreditCard,CallForwardingUnconditional,CallForwardingBusy,CallForwardingNoReply,CallDeflection,CallTransfer,NotAvailable,
	BillingID string
	//
	// ENUM: SubTotal,Total,
	TotalType string
	Currency  string
	Name      string
	Cost      string
	//
	// ENUM: 1/1000,1/100,1/10,1,10,100,1000,
	Multiplier string
	Units      string
	NumberOf   string
	TypeOf     string
}

func (AOC_DEvent) OriginalName() string { return "AOC-D" }

// Raised when an Advice of Charge message is sent at the end of a call.
//
// seealso AOCMessage
// seealso AOC-S
// seealso AOC-D
type AOC_EEvent struct {
	ChargingAssociation string
	Number              string
	Plan                string
	ID                  string
	Charge              string
	//
	// ENUM: NotAvailable,Free,Currency,Units,
	Type string
	//
	// ENUM: Normal,Reverse,CreditCard,CallForwardingUnconditional,CallForwardingBusy,CallForwardingNoReply,CallDeflection,CallTransfer,NotAvailable,
	BillingID string
	//
	// ENUM: SubTotal,Total,
	TotalType string
	Currency  string
	Name      string
	Cost      string
	//
	// ENUM: 1/1000,1/100,1/10,1,10,100,1000,
	Multiplier string
	Units      string
	NumberOf   string
	TypeOf     string
}

func (AOC_EEvent) OriginalName() string { return "AOC-E" }

// Raised when all Asterisk initialization procedures have finished.
type FullyBootedEvent struct {
	// Informational message
	Status string
	// Seconds since start
	Uptime string
	// Seconds since last reload
	LastReload string
}

// Raised when Asterisk is shutdown or restarted.
type ShutdownEvent struct {
	// Whether the shutdown is proceeding cleanly (all channels were hungup successfully) or uncleanly (channels will be terminated)
	//
	// ENUM: Uncleanly,Cleanly,
	Shutdown string
	// Whether or not a restart will occur.
	//
	// ENUM: True,False,
	Restart string
}

// Raised when the name of a channel is changed.
type RenameEvent struct {
	Channel  string
	Newname  string
	Uniqueid string
}

// Raised when two halves of a Local Channel form a bridge.
type LocalBridgeEvent struct {
	// The context in the dialplan that Channel2 starts in.
	Context string
	// The extension in the dialplan that Channel2 starts in.
	Exten string
	//
	// ENUM: Yes,No,
	LocalOptimization string
}

// Raised when two halves of a Local Channel begin to optimize
// themselves out of the media path.
//
// seealso LocalOptimizationEnd
// seealso LocalOptimizeAway
type LocalOptimizationBeginEvent struct {
	// The unique ID of the bridge into which the local channel is optimizing.
	DestUniqueId string
	// Identification for the optimization operation.
	Id string
}

// Raised when two halves of a Local Channel have finished optimizing
// themselves out of the media path.
//
// seealso LocalOptimizationBegin
// seealso LocalOptimizeAway
type LocalOptimizationEndEvent struct {
	// Indicates whether the local optimization succeeded.
	Success string
	// Identification for the optimization operation. Matches the Id from a previous `LocalOptimizationBegin`
	Id string
}

// Raised when a device state changes
// This differs from the `ExtensionStatus` event because this event is raised for all device state changes, not only for changes that affect dialplan hints.
//
// seealso ExtensionStatus
type DeviceStateChangeEvent struct {
	// The device whose state has changed
	Device string
	// The new state of the device
	State string
}

// Raised when a module has been reloaded in Asterisk.
type ReloadEvent struct {
	// The name of the module that was reloaded, or `All` if all modules were reloaded
	Module string
	// The numeric status code denoting the success or failure of the reload request.
	//
	// ENUM: 0,1,2,3,4,5,6,
	Status string
}

// Raised when a logging channel is re-enabled after a reload operation.
type LogChannelEvent struct {
	// The name of the logging channel.
	Channel string
	Enabled string
}

// Raised when a logging channel is disabled.
type LogChannel_1Event struct {
	// The name of the logging channel.
	Channel string
	Enabled string
	Reason  string
}

func (LogChannel_1Event) OriginalName() string { return "LogChannel" }

type StatusCompleteEvent StatusCompleteResponse

// Raised in response to an Originate command.
//
// seealso Originate
type OriginateResponseEvent struct {
	//
	// ENUM: Failure,Success,
	Response     string
	Channel      string
	Context      string
	Exten        string
	Application  string
	Data         string
	Reason       string
	Uniqueid     string
	CallerIDNum  string
	CallerIDName string
}

// Raised in response to a CoreShowChannels command.
//
// seealso CoreShowChannels
// seealso CoreShowChannelsComplete
type CoreShowChannelEvent struct {
	// Identifier of the bridge the channel is in, may be empty if not in one
	BridgeId string
	// Application currently executing on the channel
	Application string
	// Data given to the currently executing application
	ApplicationData string
	// The amount of time the channel has existed
	Duration string
}

// Raised at the end of the CoreShowChannel list produced by the CoreShowChannels
// command.
//
// seealso CoreShowChannels
// seealso CoreShowChannel
type CoreShowChannelsCompleteEvent struct {
	// Conveys the status of the command reponse list
	EventList string
	// The total number of list items produced
	ListItems string
}

// Raised when a hint changes due to a device state change.
//
// seealso ExtensionState
type ExtensionStatusEvent struct {
	// Name of the extension.
	Exten string
	// Context that owns the extension.
	Context string
	// Hint set for the extension
	Hint string
	// Numerical value of the extension status. Extension status is determined by the combined device state of all items contained in the hint.
	//
	// ENUM: -2,-1,0,1,2,4,8,9,16,17,
	Status string
	// Text representation of `Status`.
	//
	// ENUM: Idle,InUse,Busy,Unavailable,Ringing,InUse&Ringing,Hold,InUse&Hold,Unknown,
	StatusText string
}

// Raised when a hint changes due to a presence state change.
//
// seealso PresenceState
type PresenceStatusEvent struct {
	Exten   string
	Context string
	Hint    string
	Status  string
	Subtype string
	Message string
}

// Raised when a bridge is created.
//
// seealso BridgeDestroy
// seealso BridgeEnter
// seealso BridgeLeave
type BridgeCreateEvent struct {
}

// Raised when a bridge is destroyed.
//
// seealso BridgeCreate
// seealso BridgeEnter
// seealso BridgeLeave
type BridgeDestroyEvent struct {
}

// Raised when a channel enters a bridge.
//
// seealso BridgeCreate
// seealso BridgeDestroy
// seealso BridgeLeave
type BridgeEnterEvent struct {
	// The uniqueid of the channel being swapped out of the bridge
	SwapUniqueid string
}

// Raised when a channel leaves a bridge.
//
// seealso BridgeCreate
// seealso BridgeDestroy
// seealso BridgeEnter
type BridgeLeaveEvent struct {
}

// Raised when the channel that is the source of video in a bridge changes.
//
// seealso BridgeCreate
// seealso BridgeDestroy
type BridgeVideoSourceUpdateEvent struct {
	// The unique ID of the channel that was the video source.
	BridgePreviousVideoSource string
}

// Raised when two bridges are merged.
type BridgeMergeEvent struct {
}

// Raised when a new channel is created.
//
// seealso Newstate
// seealso Hangup
type NewchannelEvent struct {
}

// Raised when a channel's state changes.
//
// seealso Newchannel
// seealso Hangup
type NewstateEvent struct {
}

// Raised when a channel is hung up.
//
// seealso Newchannel
// seealso SoftHangupRequest
// seealso HangupRequest
// seealso Newstate
type HangupEvent struct {
	// A numeric cause code for why the channel was hung up.
	Cause string
	// A description of why the channel was hung up.
	Cause_txt string `astgo:"Cause-txt"`
}

// Raised when a hangup is requested.
//
// seealso SoftHangupRequest
// seealso Hangup
type HangupRequestEvent struct {
	// A numeric cause code for why the channel was hung up.
	Cause string
}

// Raised when a soft hangup is requested with a specific cause code.
//
// seealso HangupRequest
// seealso Hangup
type SoftHangupRequestEvent struct {
	// A numeric cause code for why the channel was hung up.
	Cause string
}

// Raised when a channel enters a new context, extension, priority.
type NewExtenEvent struct {
	// Deprecated in 12, but kept for backward compatability. Please use 'Exten' instead.
	Extension string
	// The application about to be executed.
	Application string
	// The data to be passed to the application.
	AppData string
}

// Raised when a channel receives new Caller ID information.
//
// seealso CALLERID
type NewCalleridEvent struct {
	// A description of the Caller ID presentation.
	CID_CallingPres string `astgo:"CID-CallingPres"`
}

// Raised when a channel's connected line information is changed.
//
// seealso CONNECTEDLINE
type NewConnectedLineEvent struct {
}

// Raised when a Channel's AccountCode is changed.
//
// seealso CHANNEL
type NewAccountCodeEvent struct {
	// The channel's previous account code
	OldAccountCode string
}

// Raised when a dial action has started.
//
// seealso Dial
// seealso Originate
// seealso Originate
// seealso DialEnd
type DialBeginEvent struct {
	// The non-technology specific device being dialed.
	DialString string
}

// Raised when dial status has changed.
type DialStateEvent struct {
	// The new state of the outbound dial attempt.
	//
	// ENUM: RINGING,PROCEEDING,PROGRESS,
	DialStatus string
	// If the call was forwarded, where the call was forwarded to.
	Forward string
}

// Raised when a dial action has completed.
//
// seealso Dial
// seealso Originate
// seealso Originate
// seealso DialBegin
type DialEndEvent struct {
	// The result of the dial operation.
	//
	// ENUM: ABORT,ANSWER,BUSY,CANCEL,CHANUNAVAIL,CONGESTION,CONTINUE,GOTO,NOANSWER,
	DialStatus string
	// If the call was forwarded, where the call was forwarded to.
	Forward string
}

// Raised when a channel goes on hold.
//
// seealso Unhold
type HoldEvent struct {
	// The suggested MusicClass, if provided.
	MusicClass string
}

// Raised when a channel goes off hold.
//
// seealso Hold
type UnholdEvent struct {
}

// Raised when one channel begins spying on another channel.
//
// seealso ChanSpyStop
// seealso ChanSpy
type ChanSpyStartEvent struct {
}

// Raised when a channel has stopped spying.
//
// seealso ChanSpyStart
// seealso ChanSpy
type ChanSpyStopEvent struct {
}

// Raised when a hangup handler is about to be called.
//
// seealso CHANNEL
type HangupHandlerRunEvent struct {
	// Hangup handler parameter string passed to the Gosub application.
	Handler string
}

// Raised when a hangup handler is removed from the handler stack
// by the CHANNEL() function.
//
// seealso HangupHandlerPush
// seealso CHANNEL
type HangupHandlerPopEvent struct {
	// Hangup handler parameter string passed to the Gosub application.
	Handler string
}

// Raised when a hangup handler is added to the handler stack by
// the CHANNEL() function.
//
// seealso HangupHandlerPop
// seealso CHANNEL
type HangupHandlerPushEvent struct {
	// Hangup handler parameter string passed to the Gosub application.
	Handler string
}

// Raised periodically during a fax transmission.
type FAXStatusEvent struct {
	//
	// ENUM: gateway,receive,send,
	Operation string
	// A text message describing the current status of the fax
	Status string
	// The value of the <variable name="">LOCALSTATIONID` channel variable
	LocalStationID string
	// The files being affected by the fax operation
	FileName string
}

// Raised when a receive fax operation has completed.
type ReceiveFAXEvent struct {
	// The value of the `LOCALSTATIONID` channel variable
	LocalStationID string
	// The value of the `REMOTESTATIONID` channel variable
	RemoteStationID string
	// The number of pages that have been transferred
	PagesTransferred string
	// The negotiated resolution
	Resolution string
	// The negotiated transfer rate
	TransferRate string
	// The files being affected by the fax operation
	FileName string
}

// Raised when a send fax operation has completed.
type SendFAXEvent struct {
	// The value of the <variable name="">LOCALSTATIONID` channel variable
	LocalStationID string
	// The value of the <variable name="">REMOTESTATIONID` channel variable
	RemoteStationID string
	// The number of pages that have been transferred
	PagesTransferred string
	// The negotiated resolution
	Resolution string
	// The negotiated transfer rate
	TransferRate string
	// The files being affected by the fax operation
	FileName string
}

// Raised when music on hold has started on a channel.
//
// seealso MusicOnHoldStop
// seealso StartMusicOnHold
// seealso MusicOnHold
type MusicOnHoldStartEvent struct {
	// The class of music being played on the channel
	Class string
}

// Raised when music on hold has stopped on a channel.
//
// seealso MusicOnHoldStart
// seealso StopMusicOnHold
type MusicOnHoldStopEvent struct {
}

// Raised when monitoring has started on a channel.
//
// seealso MonitorStop
// seealso Monitor
// seealso Monitor
type MonitorStartEvent struct {
}

// Raised when monitoring has stopped on a channel.
//
// seealso MonitorStart
// seealso StopMonitor
// seealso StopMonitor
type MonitorStopEvent struct {
}

// Raised when a DTMF digit has started on a channel.
//
// seealso DTMFEnd
type DTMFBeginEvent struct {
	// DTMF digit received or transmitted (0-9, A-E, # or *
	Digit string
	//
	// ENUM: Received,Sent,
	Direction string
}

// Raised when a DTMF digit has ended on a channel.
//
// seealso DTMFBegin
type DTMFEndEvent struct {
	// DTMF digit received or transmitted (0-9, A-E, # or *
	Digit string
	// Duration (in milliseconds) DTMF was sent/received
	DurationMs string
	//
	// ENUM: Received,Sent,
	Direction string
}

// Raised when the state of messages in a voicemail mailbox
// has changed or when a channel has finished interacting with a
// mailbox.
// NOTE: The Channel related parameters are only present if a channel was involved in the manipulation of a mailbox. If no channel is involved, the parameters are not included with the event.
type MessageWaitingEvent struct {
	// The mailbox with the new message, specified as `mailbox`@`context`
	Mailbox string
	// Whether or not the mailbox has messages waiting for it.
	Waiting string
	// The number of new messages.
	New string
	// The number of old messages.
	Old string
}

// Raised when a call pickup occurs.
type PickupEvent struct {
}

// Raised when a presence state changes
// This differs from the `PresenceStatus` event because this event is raised for all presence state changes, not only for changes that affect dialplan hints.
//
// seealso PresenceStatus
type PresenceStateChangeEvent struct {
	// The entity whose presence state has changed
	Presentity string
	// The new status of the presentity
	Status string
	// The new subtype of the presentity
	Subtype string
	// The new message of the presentity
	Message string
}

// Raised when an RTCP packet is sent.
//
// seealso RTCPReceived
type RTCPSentEvent struct {
	// The SSRC identifier for our stream
	SSRC string
	// The type of packet for this RTCP report.
	//
	// ENUM: 200(SR),201(RR),
	PT string
	// The address the report is sent to.
	To string
	// The number of reports that were sent.
	// The report count determines the number of ReportX headers in the message. The X for each set of report headers will range from 0 to `ReportCount - 1`.
	ReportCount string
	// The time the sender generated the report. Only valid when PT is `200(SR)`.
	SentNTP string
	// The sender's last RTP timestamp. Only valid when PT is `200(SR)`.
	SentRTP string
	// The number of packets the sender has sent. Only valid when PT is `200(SR)`.
	SentPackets string
	// The number of bytes the sender has sent. Only valid when PT is `200(SR)`.
	SentOctets string
	// The SSRC for the source of this report block.
	ReportXSourceSSRC string
	// The fraction of RTP data packets from `ReportXSourceSSRC` lost since the previous SR or RR report was sent.
	ReportXFractionLost string
	// The total number of RTP data packets from `ReportXSourceSSRC` lost since the beginning of reception.
	ReportXCumulativeLost string
	// The highest sequence number received in an RTP data packet from `ReportXSourceSSRC`.
	ReportXHighestSequence string
	// The number of sequence number cycles seen for the RTP data received from `ReportXSourceSSRC`.
	ReportXSequenceNumberCycles string
	// An estimate of the statistical variance of the RTP data packet interarrival time, measured in timestamp units.
	ReportXIAJitter string
	// The last SR timestamp received from `ReportXSourceSSRC`. If no SR has been received from `ReportXSourceSSRC`, then 0.
	ReportXLSR string
	// The delay, expressed in units of 1/65536 seconds, between receiving the last SR packet from `ReportXSourceSSRC` and sending this report.
	ReportXDLSR string
}

// Raised when an RTCP packet is received.
//
// seealso RTCPSent
type RTCPReceivedEvent struct {
	// The SSRC identifier for the remote system
	SSRC string
	// The type of packet for this RTCP report.
	//
	// ENUM: 200(SR),201(RR),
	PT string
	// The address the report was received from.
	From string
	// Calculated Round-Trip Time in seconds
	RTT string
	// The number of reports that were received.
	// The report count determines the number of ReportX headers in the message. The X for each set of report headers will range from 0 to `ReportCount - 1`.
	ReportCount string
	// The time the sender generated the report. Only valid when PT is `200(SR)`.
	SentNTP string
	// The sender's last RTP timestamp. Only valid when PT is `200(SR)`.
	SentRTP string
	// The number of packets the sender has sent. Only valid when PT is `200(SR)`.
	SentPackets string
	// The number of bytes the sender has sent. Only valid when PT is `200(SR)`.
	SentOctets string
	// The SSRC for the source of this report block.
	ReportXSourceSSRC string
	// The fraction of RTP data packets from `ReportXSourceSSRC` lost since the previous SR or RR report was sent.
	ReportXFractionLost string
	// The total number of RTP data packets from `ReportXSourceSSRC` lost since the beginning of reception.
	ReportXCumulativeLost string
	// The highest sequence number received in an RTP data packet from `ReportXSourceSSRC`.
	ReportXHighestSequence string
	// The number of sequence number cycles seen for the RTP data received from `ReportXSourceSSRC`.
	ReportXSequenceNumberCycles string
	// An estimate of the statistical variance of the RTP data packet interarrival time, measured in timestamp units.
	ReportXIAJitter string
	// The last SR timestamp received from `ReportXSourceSSRC`. If no SR has been received from `ReportXSourceSSRC`, then 0.
	ReportXLSR string
	// The delay, expressed in units of 1/65536 seconds, between receiving the last SR packet from `ReportXSourceSSRC` and sending this report.
	ReportXDLSR string
}

// Raised when a request violates an ACL check.
type FailedACLEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// If available, the name of the module that raised the event.
	Module string
	// If available, the name of the ACL that failed.
	ACLName string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request fails an authentication check due to an invalid account ID.
type InvalidAccountIDEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request fails due to exceeding the number of allowed concurrent
// sessions for that service.
type SessionLimitEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request fails due to an internal memory allocation failure.
type MemoryLimitEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request fails because a configured load average limit has been
// reached.
type LoadAverageLimitEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request fails due to some aspect of the requested item not being
// supported by the service.
type RequestNotSupportedEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The type of request attempted.
	RequestType string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request is not allowed by the service.
type RequestNotAllowedEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The type of request attempted.
	RequestType string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
	// Parameters provided to the rejected request.
	RequestParams string
}

// Raised when a request used an authentication method not allowed by the service.
type AuthMethodNotAllowedEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The authentication method attempted.
	AuthMethod string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request is received with bad formatting.
type RequestBadFormatEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The account ID associated with the rejected request.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The type of request attempted.
	RequestType string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
	// Parameters provided to the rejected request.
	RequestParams string
}

// Raised when a request successfully authenticates with a service.
type SuccessfulAuthEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// Whether or not the authentication attempt included a password.
	UsingPassword string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request has a different source address then what is expected for a
// session already in progress with a service.
type UnexpectedAddressEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The address that the request was expected to use.
	ExpectedAddress string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request's attempt to authenticate has been challenged, and the request
// failed the authentication challenge.
type ChallengeResponseFailedEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The challenge that was sent.
	Challenge string
	// The response that was received.
	Response string
	// The expected response to the challenge.
	ExpectedResponse string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request provides an invalid password during an authentication
// attempt.
type InvalidPasswordEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
	// The challenge that was sent.
	Challenge string
	// The challenge that was received.
	ReceivedChallenge string
	// The hash that was received.
	RecievedHash string
}

// Raised when an Asterisk service sends an authentication challenge to a request.
type ChallengeSentEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The challenge that was sent.
	Challenge string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// Raised when a request attempts to use a transport not allowed by the Asterisk
// service.
type InvalidTransportEvent struct {
	// The time the event was detected.
	EventTV string
	// A relative severity of the security event.
	//
	// ENUM: Informational,Error,
	Severity string
	// The Asterisk service that raised the security event.
	Service string
	// The version of this event.
	EventVersion string
	// The Service account associated with the security event notification.
	AccountID string
	// A unique identifier for the session in the service that raised the event.
	SessionID string
	// The address of the Asterisk service that raised the security event.
	LocalAddress string
	// The remote address of the entity that caused the security event to be raised.
	RemoteAddress string
	// The transport type that the request attempted to use.
	AttemptedTransport string
	// If available, the name of the module that raised the event.
	Module string
	// The timestamp reported by the session.
	SessionTV string
}

// A user defined event raised from the dialplan.
// Event may contain additional arbitrary parameters in addition to optional bridge and endpoint snapshots. Multiple snapshots of the same type are prefixed with a numeric value.
//
// seealso UserEvent
// seealso UserEvent
type UserEventEvent struct {
	// The event name, as specified in the dialplan.
	UserEvent string
}

// Raised when a blind transfer is complete.
//
// seealso BlindTransfer
type BlindTransferEvent struct {
	// Indicates if the transfer was successful or if it failed.
	// NOTE: A result of `Success` does not necessarily mean that a target was succesfully contacted. It means that a party was succesfully placed into the dialplan at the expected location.
	//
	// ENUM: Fail,Invalid,Not Permitted,Success,
	Result string
	// Indicates if the transfer was performed outside of Asterisk. For instance, a channel protocol native transfer is external. A DTMF transfer is internal.
	//
	// ENUM: Yes,No,
	IsExternal string
	// Destination context for the blind transfer.
	Context string
	// Destination extension for the blind transfer.
	Extension string
}

// Raised when an attended transfer is complete.
// The headers in this event attempt to describe all the major details of the attended transfer. The two transferer channels and the two bridges are determined based on their chronological establishment. So consider that Alice calls Bob, and then Alice transfers the call to Voicemail. The transferer and bridge headers would be arranged as follows:
// OrigTransfererChannel: Alice's channel in the bridge with Bob.
// OrigBridgeUniqueid: The bridge between Alice and Bob.
// SecondTransfererChannel: Alice's channel that called Voicemail.
// SecondBridgeUniqueid: Not present, since a call to Voicemail has no bridge.
// Now consider if the order were reversed; instead of having Alice call Bob and transfer him to Voicemail, Alice instead calls her Voicemail and transfers that to Bob. The transferer and bridge headers would be arranged as follows:
// OrigTransfererChannel: Alice's channel that called Voicemail.
// OrigBridgeUniqueid: Not present, since a call to Voicemail has no bridge.
// SecondTransfererChannel: Alice's channel in the bridge with Bob.
// SecondBridgeUniqueid: The bridge between Alice and Bob.
//
// seealso AtxFer
type AttendedTransferEvent struct {
	// Indicates if the transfer was successful or if it failed.
	// NOTE: A result of `Success` does not necessarily mean that a target was succesfully contacted. It means that a party was succesfully placed into the dialplan at the expected location.
	//
	// ENUM: Fail,Invalid,Not Permitted,Success,
	Result string
	// Indicates the method by which the attended transfer completed.
	//
	// ENUM: Bridge,App,Link,Threeway,Fail,
	DestType string
	// Indicates the surviving bridge when bridges were merged to complete the transfer
	// NOTE:  This header is only present when DestType is `Bridge` or `Threeway`
	DestBridgeUniqueid string
	// Indicates the application that is running when the transfer completes
	// NOTE:  This header is only present when DestType is `App`
	DestApp string
	// The name of the surviving transferer channel when a transfer results in a threeway call
	// NOTE:  This header is only present when DestType is `Threeway`
	DestTransfererChannel string
}

// Raised when an Agent has logged in.
//
// seealso AgentLogin
// seealso AgentLogoff
type AgentLoginEvent struct {
	// Agent ID of the agent.
	Agent string
}

// Raised when an Agent has logged off.
//
// seealso AgentLogin
type AgentLogoffEvent struct {
	// Agent ID of the agent.
	Agent string
	// The number of seconds the agent was logged in.
	Logintime string
}

// Raised when talking is detected on a channel.
//
// seealso TALK_DETECT
// seealso ChannelTalkingStop
type ChannelTalkingStartEvent struct {
}

// Raised when talking is no longer detected on a channel.
//
// seealso TALK_DETECT
// seealso ChannelTalkingStart
type ChannelTalkingStopEvent struct {
	// The length in time, in milliseconds, that talking was detected on the channel.
	Duration string
}

// Raised when the state of a peer changes.
type PeerStatusEvent struct {
	// The channel technology of the peer.
	ChannelType string
	// The name of the peer (including channel technology).
	Peer string
	// New status of the peer.
	//
	// ENUM: Unknown,Registered,Unregistered,Rejected,Lagged,
	PeerStatus string
	// The reason the status has changed.
	Cause string
	// New address of the peer.
	Address string
	// New port for the peer.
	Port string
	// Time it takes to reach the peer and receive a response.
	Time string
}

// Raised when the state of a contact changes.
type ContactStatusEvent struct {
	// This contact's URI.
	URI string
	// New status of the contact.
	//
	// ENUM: Unknown,Unreachable,Reachable,Created,Removed,Updated,
	ContactStatus string
	// The name of the associated aor.
	AOR string
	// The name of the associated endpoint.
	EndpointName string
	// The RTT measured during the last qualify.
	RoundtripUsec string
	// Content of the User-Agent header in REGISTER request
	UserAgent string
	// Absolute time that this contact is no longer valid after
	RegExpire string
	// IP address:port of the last Via header in REGISTER request
	ViaAddress string
	// Content of the Call-ID header in REGISTER request
	CallID string
}

// Raised when an outbound registration completes.
type RegistryEvent struct {
	// The type of channel that was registered (or not).
	ChannelType string
	// The username portion of the registration.
	Username string
	// The address portion of the registration.
	Domain string
	// The status of the registration request.
	//
	// ENUM: Registered,Unregistered,Rejected,Failed,
	Status string
	// What caused the rejection of the request, if available.
	Cause string
}

// Raised when a channel is parked.
type ParkedCallEvent struct {
	// Dial String that can be used to call back the parker on ParkingTimeout.
	ParkerDialString string
	// Name of the parking lot that the parkee is parked in
	Parkinglot string
	// Parking Space that the parkee is parked in
	ParkingSpace string
	// Time remaining until the parkee is forcefully removed from parking in seconds
	ParkingTimeout string
	// Time the parkee has been in the parking bridge (in seconds)
	ParkingDuration string
}

// Raised when a channel leaves a parking lot due to reaching the time limit of being
// parked.
type ParkedCallTimeOutEvent struct {
	// Dial String that can be used to call back the parker on ParkingTimeout.
	ParkerDialString string
	// Name of the parking lot that the parkee is parked in
	Parkinglot string
	// Parking Space that the parkee is parked in
	ParkingSpace string
	// Time remaining until the parkee is forcefully removed from parking in seconds
	ParkingTimeout string
	// Time the parkee has been in the parking bridge (in seconds)
	ParkingDuration string
}

// Raised when a channel leaves a parking lot because it hung up without being
// answered.
type ParkedCallGiveUpEvent struct {
	// Dial String that can be used to call back the parker on ParkingTimeout.
	ParkerDialString string
	// Name of the parking lot that the parkee is parked in
	Parkinglot string
	// Parking Space that the parkee is parked in
	ParkingSpace string
	// Time remaining until the parkee is forcefully removed from parking in seconds
	ParkingTimeout string
	// Time the parkee has been in the parking bridge (in seconds)
	ParkingDuration string
}

// Raised when a channel leaves a parking lot because it was retrieved from the parking
// lot and reconnected.
type UnParkedCallEvent struct {
	// Dial String that can be used to call back the parker on ParkingTimeout.
	ParkerDialString string
	// Name of the parking lot that the parkee is parked in
	Parkinglot string
	// Parking Space that the parkee is parked in
	ParkingSpace string
	// Time remaining until the parkee is forcefully removed from parking in seconds
	ParkingTimeout string
	// Time the parkee has been in the parking bridge (in seconds)
	ParkingDuration string
}

// Raised when a channel takes the place of a previously parked channel
// This event is raised when a channel initially parked in the parking lot is swapped out with a different channel. The most common case for this is when an attended transfer to a parking lot occurs. The Parkee information in the event will indicate the party that was swapped into the parking lot.
type ParkedCallSwapEvent struct {
	// Dial String that can be used to call back the parker on ParkingTimeout.
	ParkerDialString string
	// Name of the parking lot that the parkee is parked in
	Parkinglot string
	// Parking Space that the parkee is parked in
	ParkingSpace string
	// Time remaining until the parkee is forcefully removed from parking in seconds
	ParkingTimeout string
	// Time the parkee has been in the parking bridge (in seconds)
	ParkingDuration string
}

// Raised when a channel starts AsyncAGI command processing.
//
// seealso AsyncAGIEnd
// seealso AsyncAGIExec
// seealso AGI
// seealso AGI
type AsyncAGIStartEvent struct {
	// URL encoded string read from the AsyncAGI server.
	Env string
}

// Raised when a channel stops AsyncAGI command processing.
//
// seealso AsyncAGIStart
// seealso AsyncAGIExec
// seealso AGI
// seealso AGI
type AsyncAGIEndEvent struct {
}

// Raised when AsyncAGI completes an AGI command.
//
// seealso AsyncAGIStart
// seealso AsyncAGIEnd
// seealso AGI
// seealso AGI
type AsyncAGIExecEvent struct {
	// Optional command ID sent by the AsyncAGI server to identify the command.
	CommandID string
	// URL encoded result string from the executed AGI command.
	Result string
}

// Raised when a received AGI command starts processing.
//
// seealso AGIExecEnd
// seealso AGI
type AGIExecStartEvent struct {
	// The AGI command as received from the external source.
	Command string
	// Random identification number assigned to the execution of this command.
	CommandId string
}

// Raised when a received AGI command completes processing.
//
// seealso AGIExecStart
// seealso AGI
type AGIExecEndEvent struct {
	// The AGI command as received from the external source.
	Command string
	// Random identification number assigned to the execution of this command.
	CommandId string
	// The numeric result code from AGI
	ResultCode string
	// The text result reason from AGI
	Result string
}

// A single list item for the FAXSessions AMI command
type FAXSessionsEntryEvent struct {
	// Name of the channel responsible for the FAX session
	Channel string
	// The FAX technology that the FAX session is using
	Technology string
	// The numerical identifier for this particular session
	SessionNumber string
	// FAX session passthru/relay type
	//
	// ENUM: G.711,T.38,
	SessionType string
	// FAX session operation type
	//
	// ENUM: gateway,V.21,send,receive,none,
	Operation string
	// Current state of the FAX session
	//
	// ENUM: Uninitialized,Initialized,Open,Active,Complete,Reserved,Inactive,Unknown,
	State string
	// File or list of files associated with this FAX session
	Files string
}

// Raised when all FAXSession events are completed for a FAXSessions command
type FAXSessionsCompleteEvent struct {
	// Count of FAXSession events sent in response to FAXSessions action
	Total string
}

// Raised in response to FAXSession manager command
type FAXSessionEvent struct {
	// The numerical identifier for this particular session
	SessionNumber string
	// FAX session operation type
	//
	// ENUM: gateway,V.21,send,receive,none,
	Operation string
	// Current state of the FAX session
	//
	// ENUM: Uninitialized,Initialized,Open,Active,Complete,Reserved,Inactive,Unknown,
	State string
	// Whether error correcting mode is enabled for the FAX session. This field is not included when operation is 'V.21 Detect' or if operation is 'gateway' and state is 'Uninitialized'
	//
	// ENUM: yes,no,
	ErrorCorrectionMode string
	// Bit rate of the FAX. This field is not included when operation is 'V.21 Detect' or if operation is 'gateway' and state is 'Uninitialized'.
	DataRate string
	// Resolution of each page of the FAX. Will be in the format of X_RESxY_RES. This field is not included if the operation is anything other than Receive/Transmit.
	ImageResolution string
	// Current number of pages transferred during this FAX session. May change as the FAX progresses. This field is not included when operation is 'V.21 Detect' or if operation is 'gateway' and state is 'Uninitialized'.
	PageNumber string
	// Filename of the image being sent/recieved for this FAX session. This field is not included if Operation isn't 'send' or 'receive'.
	FileName string
	// Total number of pages sent during this session. This field is not included if Operation isn't 'send' or 'receive'. Will always be 0 for 'receive'.
	PagesTransmitted string
	// Total number of pages received during this session. This field is not included if Operation is not 'send' or 'receive'. Will be 0 for 'send'.
	PagesReceived string
	// Total number of bad lines sent/recieved during this session. This field is not included if Operation is not 'send' or 'received'.
	TotalBadLines string
}

// Raised in response to FAXStats manager command
type FAXStatsEvent struct {
	// Number of active FAX sessions
	//
	// reqtured
	CurrentSessions string
	// Number of reserved FAX sessions
	//
	// reqtured
	ReservedSessions string
	// Total FAX sessions for which Asterisk is/was the transmitter
	//
	// reqtured
	TransmitAttempts string
	// Total FAX sessions for which Asterisk is/was the recipient
	//
	// reqtured
	ReceiveAttempts string
	// Total FAX sessions which have been completed successfully
	//
	// reqtured
	CompletedFAXes string
	// Total FAX sessions which failed to complete successfully
	//
	// reqtured
	FailedFAXes string
}

// Raised in response to a MWIGet command.
//
// seealso MWIGet
type MWIGetEvent struct {
	// Specific mailbox ID.
	Mailbox string
	// The number of old messages in the mailbox.
	OldMessages string
	// The number of new messages in the mailbox.
	NewMessages string
}

// Raised in response to a MWIGet command.
//
// seealso MWIGet
type MWIGetCompleteEvent struct {
	EventList string
	// The number of mailboxes reported.
	ListItems string
}

// Provide details about an identify section.
type IdentifyDetailEvent struct {
	// The object's type. This will always be 'identify'.
	ObjectType string
	// The name of this object.
	ObjectName string
	// Name of Endpoint
	Endpoint string
	// IP addresses or networks to match against.
	Match string
	// The name of the endpoint associated with this information.
	EndpointName string
}

// Provide details about an Address of Record (AoR) section.
type AorDetailEvent struct {
	// The object's type. This will always be 'aor'.
	ObjectType string
	// The name of this object.
	ObjectName string
	// Minimum keep alive time for an AoR
	MinimumExpiration string
	// Maximum time to keep an AoR
	MaximumExpiration string
	// Default expiration time in seconds for contacts that are dynamically bound to an AoR.
	DefaultExpiration string
	// Interval at which to qualify an AoR
	QualifyFrequency string
	// Authenticates a qualify request if needed
	AuthenticateQualify string
	// Maximum number of contacts that can bind to an AoR
	MaxContacts string
	// Determines whether new contacts replace existing ones.
	RemoveExisting string
	// Allow subscriptions for the specified mailbox(es)
	Mailboxes string
	// Outbound proxy used when sending OPTIONS request
	OutboundProxy string
	// Enables Path support for REGISTER requests and Route support for other requests.
	SupportPath string
	// The total number of contacts associated with this AoR.
	TotalContacts string
	// The number of non-permanent contacts associated with this AoR.
	ContactsRegistered string
	// The name of the endpoint associated with this information.
	EndpointName string
}

// Provide details about an authentication section.
type AuthDetailEvent struct {
	// The object's type. This will always be 'auth'.
	ObjectType string
	// The name of this object.
	ObjectName string
	// Username to use for account
	Username string
	// Username to use for account
	Password string
	// MD5 Hash used for authentication.
	Md5Cred string
	// SIP realm for endpoint
	Realm string
	// Lifetime of a nonce associated with this authentication config.
	NonceLifetime string
	// Authentication type
	AuthType string
	// The name of the endpoint associated with this information.
	EndpointName string
}

// Provide details about an authentication section.
type TransportDetailEvent struct {
	// The object's type. This will always be 'transport'.
	ObjectType string
	// The name of this object.
	ObjectName string
	// Protocol to use for SIP traffic
	Protocol string
	// IP Address and optional port to bind to for this transport
	Bind string
	// Number of simultaneous Asynchronous Operations
	AsycOperations string
	// File containing a list of certificates to read (TLS ONLY)
	CaListFile string
	// Path to directory containing a list of certificates to read (TLS ONLY)
	CaListPath string
	// Certificate file for endpoint (TLS ONLY)
	CertFile string
	// Private key file (TLS ONLY)
	PrivKeyFile string
	// Password required for transport
	Password string
	// External address for SIP signalling
	ExternalSignalingAddress string
	// External port for SIP signalling
	ExternalSignalingPort string
	// External IP address to use in RTP handling
	ExternalMediaAddress string
	// Domain the transport comes from
	Domain string
	// Require verification of server certificate (TLS ONLY)
	VerifyServer string
	// Require verification of client certificate (TLS ONLY)
	VerifyClient string
	// Require client certificate (TLS ONLY)
	RequireClientCert string
	// Method of SSL transport (TLS ONLY)
	Method string
	// Preferred cryptography cipher names (TLS ONLY)
	Cipher string
	// Network to consider local (used for NAT purposes).
	LocalNet string
	// Enable TOS for the signalling sent over this transport
	Tos string
	// Enable COS for the signalling sent over this transport
	Cos string
	// The timeout (in milliseconds) to set on WebSocket connections.
	WebsocketWriteTimeout string
	// The name of the endpoint associated with this information.
	EndpointName string
}

// Provide details about an endpoint section.
type EndpointDetailEvent struct {
	// The object's type. This will always be 'endpoint'.
	ObjectType string
	// The name of this object.
	ObjectName string
	// Dialplan context for inbound sessions
	Context string
	// Media Codec(s) to disallow
	Disallow string
	// Media Codec(s) to allow
	Allow string
	// DTMF mode
	DtmfMode string
	// Allow use of IPv6 for RTP traffic
	RtpIpv6 string
	// Enforce that RTP must be symmetric
	RtpSymmetric string
	// Enable the ICE mechanism to help traverse NAT
	IceSupport string
	// Use Endpoint's requested packetisation interval
	UsePtime string
	// Force use of return port
	ForceRport string
	// Allow Contact header to be rewritten with the source IP address-port
	RewriteContact string
	// Desired transport configuration
	Transport string
	// Full SIP URI of the outbound proxy used to send requests
	OutboundProxy string
	// Default Music On Hold class
	MohSuggest string
	// Allow support for RFC3262 provisional ACK tags
	A_100rel string `astgo:"100rel"`
	// Session timers for SIP packets
	Timers string
	// Minimum session timers expiration period
	TimersMinSe string
	// Maximum session timer expiration period
	TimersSessExpires string
	// Authentication Object(s) associated with the endpoint
	Auth string
	// Authentication object(s) used for outbound requests
	OutboundAuth string
	// AoR(s) to be used with the endpoint
	Aors string
	// IP address used in SDP for media handling
	MediaAddress string
	// Way(s) for Endpoint to be identified
	IdentifyBy string
	// Determines whether media may flow directly between endpoints.
	DirectMedia string
	// Direct Media method type
	DirectMediaMethod string
	// Connected line method type
	ConnectedLineMethod string
	// Mitigation of direct media (re)INVITE glare
	DirectMediaGlareMitigation string
	// Disable direct media session refreshes when NAT obstructs the media session
	DisableDirectMediaOnNat string
	// CallerID information for the endpoint
	Callerid string
	// Default privacy level
	CalleridPrivacy string
	// Internal id_tag for the endpoint
	CalleridTag string
	// Accept identification information received from this endpoint
	TrustIdInbound string
	// Send private identification details to the endpoint.
	TrustIdOutbound string
	// Send the P-Asserted-Identity header
	SendPai string
	// Send the Remote-Party-ID header
	SendRpid string
	// Send the Diversion header, conveying the diversion information to the called user agent
	SendDiversion string
	// NOTIFY the endpoint when state changes for any of the specified mailboxes
	Mailboxes string
	// Condense MWI notifications into a single NOTIFY.
	AggregateMwi string
	// Determines whether res_pjsip will use and enforce usage of media encryption for this endpoint.
	MediaEncryption string
	// Determines whether encryption should be used if possible but does not terminate the session if not achieved.
	MediaEncryptionOptimistic string
	// Determines whether res_pjsip will use and enforce usage of AVPF for this endpoint.
	UseAvpf string
	// Determines whether res_pjsip will use and enforce usage of AVP, regardless of the RTP profile in use for this endpoint.
	ForceAvp string
	// Determines whether res_pjsip will use the media transport received in the offer SDP in the corresponding answer SDP.
	MediaUseReceivedTransport string
	// Determines whether one-touch recording is allowed for this endpoint.
	OneTouchRecording string
	// Determines whether chan_pjsip will indicate ringing using inband progress.
	InbandProgress string
	// The numeric pickup groups for a channel.
	CallGroup string
	// The numeric pickup groups that a channel can pickup.
	PickupGroup string
	// The named pickup groups for a channel.
	NamedCallGroup string
	// The named pickup groups that a channel can pickup.
	NamedPickupGroup string
	// The number of in-use channels which will cause busy to be returned as device state
	DeviceStateBusyAt string
	// Whether T.38 UDPTL support is enabled or not
	T38Udptl string
	// T.38 UDPTL error correction method
	T38UdptlEc string
	// T.38 UDPTL maximum datagram size
	T38UdptlMaxdatagram string
	// Whether CNG tone detection is enabled
	FaxDetect string
	// Whether NAT support is enabled on UDPTL sessions
	T38UdptlNat string
	// Whether IPv6 is used for UDPTL Sessions
	T38UdptlIpv6 string
	// Set which country's indications to use for channels created for this endpoint.
	ToneZone string
	// Set the default language to use for channels created for this endpoint.
	Language string
	// The feature to enact when one-touch recording is turned on.
	RecordOnFeature string
	// The feature to enact when one-touch recording is turned off.
	RecordOffFeature string
	// Determines whether SIP REFER transfers are allowed for this endpoint
	AllowTransfer string
	// Determines whether a user=phone parameter is placed into the request URI if the user is determined to be a phone number
	UserEqPhone string
	// Determines whether hold and unhold will be passed through using re-INVITEs with recvonly and sendrecv to the remote side
	MohPassthrough string
	// String placed as the username portion of an SDP origin (o=) line.
	SdpOwner string
	// String used for the SDP session (s=) line.
	SdpSession string
	// DSCP TOS bits for audio streams
	TosAudio string
	// DSCP TOS bits for video streams
	TosVideo string
	// Priority for audio streams
	CosAudio string
	// Priority for video streams
	CosVideo string
	// Determines if endpoint is allowed to initiate subscriptions with Asterisk.
	AllowSubscribe string
	// The minimum allowed expiry time for subscriptions initiated by the endpoint.
	SubMinExpiry string
	// Username to use in From header for requests to this endpoint.
	FromUser string
	// Domain to user in From header for requests to this endpoint.
	FromDomain string
	// Username to use in From header for unsolicited MWI NOTIFYs to this endpoint.
	MwiFromUser string
	// Name of the RTP engine to use for channels created for this endpoint
	RtpEngine string
	// Verify that the provided peer certificate is valid
	DtlsVerify string
	// Interval at which to renegotiate the TLS session and rekey the SRTP session
	DtlsRekey string
	// Path to certificate file to present to peer
	DtlsCertFile string
	// Path to private key for certificate file
	DtlsPrivateKey string
	// Cipher to use for DTLS negotiation
	DtlsCipher string
	// Path to certificate authority certificate
	DtlsCaFile string
	// Path to a directory containing certificate authority certificates
	DtlsCaPath string
	// Whether we are willing to accept connections, connect to the other party, or both.
	DtlsSetup string
	// Determines whether 32 byte tags should be used instead of 80 byte tags.
	SrtpTag32 string
	// How redirects received from an endpoint are handled
	RedirectMethod string
	// Variable set on a channel involving the endpoint.
	SetVar string
	// Context to route incoming MESSAGE requests to.
	MessageContext string
	// An accountcode to set automatically on any channels created for this endpoint.
	Accountcode string
	// Respond to a SIP invite with the single most preferred codec rather than advertising all joint codec capabilities. This limits the other side's codec choice to exactly what we prefer.
	PreferredCodecOnly string
	// The aggregate device state for this endpoint.
	DeviceState string
	// The number of active channels associated with this endpoint.
	ActiveChannels string
	// Context for incoming MESSAGE requests.
	SubscribeContext string
	// Enable RFC3578 overlap dialing support.
	Allowoverlap string
}

// Provide details about a contact's status.
type ContactStatusDetailEvent struct {
	// The AoR that owns this contact.
	AOR string
	// This contact's URI.
	URI string
	// This contact's status.
	//
	// ENUM: Reachable,Unreachable,
	Status string
	// The round trip time in microseconds.
	RoundtripUsec string
	// The name of the endpoint associated with this information.
	EndpointName string
	// Content of the User-Agent header in REGISTER request
	UserAgent string
	// Absolute time that this contact is no longer valid after
	RegExpire string
	// IP address:port of the last Via header in REGISTER request. Will only appear in the event if available.
	ViaAddress string
	// Content of the Call-ID header in REGISTER request. Will only appear in the event if available.
	CallID string
	// The sorcery ID of the contact.
	ID string
	// A boolean indicating whether a qualify should be authenticated.
	AuthenticateQualify string
	// The contact's outbound proxy.
	OutboundProxy string
	// The Path header received on the REGISTER.
	Path string
	// The interval in seconds at which the contact will be qualified.
	QualifyFrequency string
	// The elapsed time in decimal seconds after which an OPTIONS message is sent before the contact is considered unavailable.
	QualifyTimeout string
}

// Provide details about a contact's status.
type EndpointListEvent struct {
	// The object's type. This will always be 'endpoint'.
	ObjectType string
	// The name of this object.
	ObjectName string
	// The transport configurations associated with this endpoint.
	Transport string
	// The aor configurations associated with this endpoint.
	Aor string
	// The inbound authentication configurations associated with this endpoint.
	Auths string
	// The outbound authentication configurations associated with this endpoint.
	OutboundAuths string
	// The aggregate device state for this endpoint.
	DeviceState string
	// The number of active channels associated with this endpoint.
	ActiveChannels string
}

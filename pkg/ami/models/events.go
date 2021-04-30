package amimodels

// MCIDEvent Published when a malicious call ID request arrives.
type MCIDEvent struct {
	MCallerIDNumValid string

	MCallerIDNum string

	MCallerIDton string

	MCallerIDNumPlan string

	MCallerIDNumPres string

	MCallerIDNameValid string

	MCallerIDName string

	MCallerIDNameCharSet string

	MCallerIDNamePres string

	MCallerIDSubaddr string

	MCallerIDSubaddrType string

	MCallerIDSubaddrOdd string

	MCallerIDPres string

	MConnectedIDNumValid string

	MConnectedIDNum string

	MConnectedIDton string

	MConnectedIDNumPlan string

	MConnectedIDNumPres string

	MConnectedIDNameValid string

	MConnectedIDName string

	MConnectedIDNameCharSet string

	MConnectedIDNamePres string

	MConnectedIDSubaddr string

	MConnectedIDSubaddrType string

	MConnectedIDSubaddrOdd string

	MConnectedIDPres string
}

func (MCIDEvent) EventTypeName() string {
	return "MCID"
}

// AlarmClearEvent Raised when an alarm is cleared on a DAHDI channel.
type AlarmClearEvent struct {

	// The DAHDI channel on which the alarm was cleared.
	DAHDIChannel string
}

func (AlarmClearEvent) EventTypeName() string {
	return "AlarmClear"
}

// SpanAlarmClearEvent Raised when an alarm is cleared on a DAHDI span.
type SpanAlarmClearEvent struct {

	// The span on which the alarm was cleared.
	Span string
}

func (SpanAlarmClearEvent) EventTypeName() string {
	return "SpanAlarmClear"
}

// DNDStateEvent Raised when the Do Not Disturb state is changed on a DAHDI channel.
type DNDStateEvent struct {

	// The DAHDI channel on which DND status changed.
	DAHDIChannel string

	Status string
}

func (DNDStateEvent) EventTypeName() string {
	return "DNDState"
}

// AlarmEvent Raised when an alarm is set on a DAHDI channel.
type AlarmEvent struct {

	// The channel on which the alarm occurred.
	DAHDIChannel string

	// A textual description of the alarm that occurred.
	Alarm string
}

func (AlarmEvent) EventTypeName() string {
	return "Alarm"
}

// SpanAlarmEvent Raised when an alarm is set on a DAHDI span.
type SpanAlarmEvent struct {

	// The span on which the alarm occurred.
	Span string

	// A textual description of the alarm that occurred.
	Alarm string
}

func (SpanAlarmEvent) EventTypeName() string {
	return "SpanAlarm"
}

// DAHDIChannelEvent Raised when a DAHDI channel is created or an underlying technology is associated with a DAHDI channel.
type DAHDIChannelEvent struct {

	// The DAHDI logical group associated with this channel.
	DAHDIGroup string

	// The DAHDI span associated with this channel.
	DAHDISpan string

	// The DAHDI channel associated with this channel.
	DAHDIChannel string
}

func (DAHDIChannelEvent) EventTypeName() string {
	return "DAHDIChannel"
}

// SIPQualifyPeerDoneEvent Raised when SIPQualifyPeer has finished qualifying the specified peer.
type SIPQualifyPeerDoneEvent struct {

	// The name of the peer.
	Peer string

	// This is only included if an ActionID Header was sent with the action request, in which case it will be
	//               that ActionID.
	ActionID string
}

func (SIPQualifyPeerDoneEvent) EventTypeName() string {
	return "SIPQualifyPeerDone"
}

// SessionTimeoutEvent Raised when a SIP session times out.
type SessionTimeoutEvent struct {

	// The source of the session timeout.
	Source string
}

func (SessionTimeoutEvent) EventTypeName() string {
	return "SessionTimeout"
}

// AgentsEvent Response event in a series to the Agents AMI action containing
//         information about a defined agent.
type AgentsEvent struct {

	// Agent ID of the agent.
	Agent string

	// User friendly name of the agent.
	Name string

	// Current status of the agent.
	//   The valid values are:
	Status string

	// value on agent channel.
	//   Present if Status value is .
	TalkingToChan string

	// Epoche time when the agent started talking with the caller.
	//   Present if Status value is .
	CallStarted string

	// Epoche time when the agent logged in.
	//   Present if Status value is  or .
	LoggedInTime string

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (AgentsEvent) EventTypeName() string {
	return "Agents"
}

// AgentsCompleteEvent Final response event in a series of events to the Agents AMI action.
type AgentsCompleteEvent struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (AgentsCompleteEvent) EventTypeName() string {
	return "AgentsComplete"
}

// ConfbridgeListEvent Raised as part of the ConfbridgeList action response list.
type ConfbridgeListEvent struct {

	// The name of the Confbridge conference.
	Conference string

	// Identifies this user as an admin user.
	Admin string

	// Identifies this user as a marked user.
	MarkedUser string

	// Must this user wait for a marked user to join?
	WaitMarked string

	// Does this user get kicked after the last marked user leaves?
	EndMarked string

	// Is this user waiting for a marked user to join?
	Waiting string

	// The current mute status.
	Muted string

	// Is this user talking?
	Talking string

	// The number of seconds the channel has been up.
	AnsweredTime string
}

func (ConfbridgeListEvent) EventTypeName() string {
	return "ConfbridgeList"
}

// MeetmeJoinEvent Raised when a user joins a MeetMe conference.
type MeetmeJoinEvent struct {

	// The identifier for the MeetMe conference.
	Meetme string

	// The identifier of the MeetMe user who joined.
	User string
}

func (MeetmeJoinEvent) EventTypeName() string {
	return "MeetmeJoin"
}

// MeetmeLeaveEvent Raised when a user leaves a MeetMe conference.
type MeetmeLeaveEvent struct {

	// The identifier for the MeetMe conference.
	Meetme string

	// The identifier of the MeetMe user who joined.
	User string

	// The length of time in seconds that the Meetme user was in the conference.
	Duration string
}

func (MeetmeLeaveEvent) EventTypeName() string {
	return "MeetmeLeave"
}

// MeetmeEndEvent Raised when a MeetMe conference ends.
type MeetmeEndEvent struct {

	// The identifier for the MeetMe conference.
	Meetme string
}

func (MeetmeEndEvent) EventTypeName() string {
	return "MeetmeEnd"
}

// MeetmeTalkRequestEvent Raised when a MeetMe user has started talking.
type MeetmeTalkRequestEvent struct {

	// The identifier for the MeetMe conference.
	Meetme string

	// The identifier of the MeetMe user who joined.
	User string

	// The length of time in seconds that the Meetme user has been in the conference at the time of this
	//               event.
	Duration string

	Status string
}

func (MeetmeTalkRequestEvent) EventTypeName() string {
	return "MeetmeTalkRequest"
}

// MeetmeTalkingEvent Raised when a MeetMe user begins or ends talking.
type MeetmeTalkingEvent struct {

	// The identifier for the MeetMe conference.
	Meetme string

	// The identifier of the MeetMe user who joined.
	User string

	// The length of time in seconds that the Meetme user has been in the conference at the time of this
	//               event.
	Duration string

	Status string
}

func (MeetmeTalkingEvent) EventTypeName() string {
	return "MeetmeTalking"
}

// MeetmeMuteEvent Raised when a MeetMe user is muted or unmuted.
type MeetmeMuteEvent struct {

	// The identifier for the MeetMe conference.
	Meetme string

	// The identifier of the MeetMe user who joined.
	User string

	// The length of time in seconds that the Meetme user has been in the conference at the time of this
	//               event.
	Duration string

	Status string
}

func (MeetmeMuteEvent) EventTypeName() string {
	return "MeetmeMute"
}

// ConfbridgeStartEvent Raised when a conference starts.
type ConfbridgeStartEvent struct {

	// The name of the Confbridge conference.
	Conference string
}

func (ConfbridgeStartEvent) EventTypeName() string {
	return "ConfbridgeStart"
}

// ConfbridgeEndEvent Raised when a conference ends.
type ConfbridgeEndEvent struct {

	// The name of the Confbridge conference.
	Conference string
}

func (ConfbridgeEndEvent) EventTypeName() string {
	return "ConfbridgeEnd"
}

// ConfbridgeJoinEvent Raised when a channel joins a Confbridge conference.
type ConfbridgeJoinEvent struct {

	// The name of the Confbridge conference.
	Conference string

	// Identifies this user as an admin user.
	Admin string

	// The joining mute status.
	Muted string
}

func (ConfbridgeJoinEvent) EventTypeName() string {
	return "ConfbridgeJoin"
}

// ConfbridgeLeaveEvent Raised when a channel leaves a Confbridge conference.
type ConfbridgeLeaveEvent struct {

	// The name of the Confbridge conference.
	Conference string

	// Identifies this user as an admin user.
	Admin string
}

func (ConfbridgeLeaveEvent) EventTypeName() string {
	return "ConfbridgeLeave"
}

// ConfbridgeRecordEvent Raised when a conference starts recording.
type ConfbridgeRecordEvent struct {

	// The name of the Confbridge conference.
	Conference string
}

func (ConfbridgeRecordEvent) EventTypeName() string {
	return "ConfbridgeRecord"
}

// ConfbridgeStopRecordEvent Raised when a conference that was recording stops recording.
type ConfbridgeStopRecordEvent struct {

	// The name of the Confbridge conference.
	Conference string
}

func (ConfbridgeStopRecordEvent) EventTypeName() string {
	return "ConfbridgeStopRecord"
}

// ConfbridgeMuteEvent Raised when a Confbridge participant mutes.
type ConfbridgeMuteEvent struct {

	// The name of the Confbridge conference.
	Conference string

	// Identifies this user as an admin user.
	Admin string
}

func (ConfbridgeMuteEvent) EventTypeName() string {
	return "ConfbridgeMute"
}

// ConfbridgeUnmuteEvent Raised when a confbridge participant unmutes.
type ConfbridgeUnmuteEvent struct {

	// The name of the Confbridge conference.
	Conference string

	// Identifies this user as an admin user.
	Admin string
}

func (ConfbridgeUnmuteEvent) EventTypeName() string {
	return "ConfbridgeUnmute"
}

// ConfbridgeTalkingEvent Raised when a confbridge participant begins or ends talking.
type ConfbridgeTalkingEvent struct {

	// The name of the Confbridge conference.
	Conference string

	TalkingStatus string

	// Identifies this user as an admin user.
	Admin string
}

func (ConfbridgeTalkingEvent) EventTypeName() string {
	return "ConfbridgeTalking"
}

// VarSetEvent Raised when a variable local to the gosub stack frame is set due to a subroutine call.
type VarSetEvent struct {

	// The LOCAL variable being set.
	Variable string

	// The new value of the variable.
	Value string
}

func (VarSetEvent) EventTypeName() string {
	return "VarSet"
}

// QueueMemberStatusEvent Raised when a Queue member's status has changed.
type QueueMemberStatusEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// Channel technology or location from which to read device state changes.
	StateInterface string

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
	InCall string

	// The numeric device state status of the queue member.
	Status string

	Paused string

	// If set when paused, the reason the queue member was paused.
	PausedReason string

	Ringinuse string

	// The Wrapup Time of the queue member. If this value is set will override the wrapup time of queue.
	Wrapuptime string
}

func (QueueMemberStatusEvent) EventTypeName() string {
	return "QueueMemberStatus"
}

// QueueMemberAddedEvent Raised when a member is added to the queue.
type QueueMemberAddedEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// Channel technology or location from which to read device state changes.
	StateInterface string

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
	InCall string

	// The numeric device state status of the queue member.
	Status string

	Paused string

	// If set when paused, the reason the queue member was paused.
	PausedReason string

	Ringinuse string

	// The Wrapup Time of the queue member. If this value is set will override the wrapup time of queue.
	Wrapuptime string
}

func (QueueMemberAddedEvent) EventTypeName() string {
	return "QueueMemberAdded"
}

// QueueMemberRemovedEvent Raised when a member is removed from the queue.
type QueueMemberRemovedEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// Channel technology or location from which to read device state changes.
	StateInterface string

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
	InCall string

	// The numeric device state status of the queue member.
	Status string

	Paused string

	// If set when paused, the reason the queue member was paused.
	PausedReason string

	Ringinuse string

	// The Wrapup Time of the queue member. If this value is set will override the wrapup time of queue.
	Wrapuptime string
}

func (QueueMemberRemovedEvent) EventTypeName() string {
	return "QueueMemberRemoved"
}

// QueueMemberPauseEvent Raised when a member is paused/unpaused in the queue.
type QueueMemberPauseEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// Channel technology or location from which to read device state changes.
	StateInterface string

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
	InCall string

	// The numeric device state status of the queue member.
	Status string

	Paused string

	// If set when paused, the reason the queue member was paused.
	PausedReason string

	Ringinuse string

	// The Wrapup Time of the queue member. If this value is set will override the wrapup time of queue.
	Wrapuptime string
}

func (QueueMemberPauseEvent) EventTypeName() string {
	return "QueueMemberPause"
}

// QueueMemberPenaltyEvent Raised when a member's penalty is changed.
type QueueMemberPenaltyEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// Channel technology or location from which to read device state changes.
	StateInterface string

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
	InCall string

	// The numeric device state status of the queue member.
	Status string

	Paused string

	// If set when paused, the reason the queue member was paused.
	PausedReason string

	Ringinuse string

	// The Wrapup Time of the queue member. If this value is set will override the wrapup time of queue.
	Wrapuptime string
}

func (QueueMemberPenaltyEvent) EventTypeName() string {
	return "QueueMemberPenalty"
}

// QueueMemberRinginuseEvent Raised when a member's ringinuse setting is changed.
type QueueMemberRinginuseEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// Channel technology or location from which to read device state changes.
	StateInterface string

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
	InCall string

	// The numeric device state status of the queue member.
	Status string

	Paused string

	// If set when paused, the reason the queue member was paused.
	PausedReason string

	Ringinuse string

	// The Wrapup Time of the queue member. If this value is set will override the wrapup time of queue.
	Wrapuptime string
}

func (QueueMemberRinginuseEvent) EventTypeName() string {
	return "QueueMemberRinginuse"
}

// QueueCallerJoinEvent Raised when a caller joins a Queue.
type QueueCallerJoinEvent struct {

	// The name of the queue.
	Queue string

	// This channel's current position in the queue.
	Position string

	// The total number of channels in the queue.
	Count string
}

func (QueueCallerJoinEvent) EventTypeName() string {
	return "QueueCallerJoin"
}

// QueueCallerLeaveEvent Raised when a caller leaves a Queue.
type QueueCallerLeaveEvent struct {

	// The name of the queue.
	Queue string

	// The total number of channels in the queue.
	Count string

	// This channel's current position in the queue.
	Position string
}

func (QueueCallerLeaveEvent) EventTypeName() string {
	return "QueueCallerLeave"
}

// QueueCallerAbandonEvent Raised when a caller abandons the queue.
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

func (QueueCallerAbandonEvent) EventTypeName() string {
	return "QueueCallerAbandon"
}

// AgentCalledEvent Raised when an queue member is notified of a caller in the queue.
type AgentCalledEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string
}

func (AgentCalledEvent) EventTypeName() string {
	return "AgentCalled"
}

// AgentRingNoAnswerEvent Raised when a queue member is notified of a caller in the queue and fails to answer.
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

func (AgentRingNoAnswerEvent) EventTypeName() string {
	return "AgentRingNoAnswer"
}

// AgentCompleteEvent Raised when a queue member has finished servicing a caller in the queue.
type AgentCompleteEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string

	// The time the channel was in the queue, expressed in seconds since 00:00, Jan 1, 1970 UTC.
	HoldTime string

	// The time the queue member talked with the caller in the queue, expressed in seconds since 00:00, Jan 1,
	//               1970 UTC.
	TalkTime string

	Reason string
}

func (AgentCompleteEvent) EventTypeName() string {
	return "AgentComplete"
}

// AgentDumpEvent Raised when a queue member hangs up on a caller in the queue.
type AgentDumpEvent struct {

	// The name of the queue.
	Queue string

	// The name of the queue member.
	MemberName string

	// The queue member's channel technology or location.
	Interface string
}

func (AgentDumpEvent) EventTypeName() string {
	return "AgentDump"
}

// AgentConnectEvent Raised when a queue member answers and is bridged to a caller in the queue.
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

func (AgentConnectEvent) EventTypeName() string {
	return "AgentConnect"
}

// MiniVoiceMailEvent Raised when a notification is sent out by a MiniVoiceMail application
type MiniVoiceMailEvent struct {

	// What action was taken. Currently, this will always be
	Action string

	// The mailbox that the notification was about, specified as @
	Mailbox string

	// A message counter derived from the  channel variable.
	Counter string
}

func (MiniVoiceMailEvent) EventTypeName() string {
	return "MiniVoiceMail"
}

// CdrEvent Raised when a CDR is generated.
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

	// The parameters passed to the last dialplan application the
	//               Party A executed.
	LastData string

	// The time the CDR was created.
	StartTime string

	// The earliest of either the time when Party A answered, or
	//               the start time of this CDR.
	AnswerTime string

	// The time when the CDR was finished. This occurs when the
	//               Party A hangs up or when the bridge between Party A and
	//               Party B is broken.
	EndTime string

	// The time, in seconds, of  - .
	Duration string

	// The time, in seconds, of  - .
	BillableSeconds string

	// The final known disposition of the CDR.
	Disposition string

	// A flag that informs a billing system how to treat the CDR.
	AMAFlags string

	// A unique identifier for the Party A channel.
	UniqueID string

	// A user defined field set on the channels. If set on both the Party A
	//               and Party B channel, the userfields of both are concatenated and
	//               separated by a .
	UserField string
}

func (CdrEvent) EventTypeName() string {
	return "Cdr"
}

// CELEvent Raised when a Channel Event Log is generated for a channel.
type CELEvent struct {

	// The name of the CEL event being raised. This can include
	//               both the system defined CEL events, as well as user defined
	//               events.
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

	// The arguments passed to the dialplan .
	AppData string

	// The time the CEL event occurred.
	EventTime string

	// A flag that informs a billing system how to treat the CEL.
	AMAFlags string

	// The unique ID of the channel.
	UniqueID string

	// The linked ID of the channel, which ties this event to other related channel's events.
	LinkedID string

	// A user defined field set on a channel, containing arbitrary
	//               application specific data.
	UserField string

	// If this channel is in a bridge, the channel that it is in
	//               a bridge with.
	Peer string

	// If this channel is in a bridge, the accountcode of the
	//               channel it is in a bridge with.
	PeerAccount string

	// Some events will have event specific data that accompanies the CEL record.
	//               This extra data is JSON encoded, and is dependent on the event in
	//               question.
	Extra string
}

func (CELEvent) EventTypeName() string {
	return "CEL"
}

// FailedACLEvent Raised when a request violates an ACL check.
type FailedACLEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// If available, the name of the module that raised the event.
	Module string

	// If available, the name of the ACL that failed.
	ACLName string

	// The timestamp reported by the session.
	SessionTV string
}

func (FailedACLEvent) EventTypeName() string {
	return "FailedACL"
}

// InvalidAccountIDEvent Raised when a request fails an authentication check due to an invalid account ID.
type InvalidAccountIDEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (InvalidAccountIDEvent) EventTypeName() string {
	return "InvalidAccountID"
}

// SessionLimitEvent Raised when a request fails due to exceeding the number of allowed concurrent sessions for that
//         service.
type SessionLimitEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (SessionLimitEvent) EventTypeName() string {
	return "SessionLimit"
}

// MemoryLimitEvent Raised when a request fails due to an internal memory allocation failure.
type MemoryLimitEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (MemoryLimitEvent) EventTypeName() string {
	return "MemoryLimit"
}

// LoadAverageLimitEvent Raised when a request fails because a configured load average limit has been reached.
type LoadAverageLimitEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (LoadAverageLimitEvent) EventTypeName() string {
	return "LoadAverageLimit"
}

// RequestNotSupportedEvent Raised when a request fails due to some aspect of the requested item not being supported by the
//         service.
type RequestNotSupportedEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// The type of request attempted.
	RequestType string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (RequestNotSupportedEvent) EventTypeName() string {
	return "RequestNotSupported"
}

// RequestNotAllowedEvent Raised when a request is not allowed by the service.
type RequestNotAllowedEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
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

func (RequestNotAllowedEvent) EventTypeName() string {
	return "RequestNotAllowed"
}

// AuthMethodNotAllowedEvent Raised when a request used an authentication method not allowed by the service.
type AuthMethodNotAllowedEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// The authentication method attempted.
	AuthMethod string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (AuthMethodNotAllowedEvent) EventTypeName() string {
	return "AuthMethodNotAllowed"
}

// RequestBadFormatEvent Raised when a request is received with bad formatting.
type RequestBadFormatEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
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

func (RequestBadFormatEvent) EventTypeName() string {
	return "RequestBadFormat"
}

// SuccessfulAuthEvent Raised when a request successfully authenticates with a service.
type SuccessfulAuthEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// Whether or not the authentication attempt included a password.
	UsingPassword string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (SuccessfulAuthEvent) EventTypeName() string {
	return "SuccessfulAuth"
}

// UnexpectedAddressEvent Raised when a request has a different source address then what is expected for a session already in
//         progress with a service.
type UnexpectedAddressEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// The address that the request was expected to use.
	ExpectedAddress string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (UnexpectedAddressEvent) EventTypeName() string {
	return "UnexpectedAddress"
}

// ChallengeResponseFailedEvent Raised when a request's attempt to authenticate has been challenged, and the request failed the
//         authentication challenge.
type ChallengeResponseFailedEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
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

func (ChallengeResponseFailedEvent) EventTypeName() string {
	return "ChallengeResponseFailed"
}

// InvalidPasswordEvent Raised when a request provides an invalid password during an authentication attempt.
type InvalidPasswordEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
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
	ReceivedHash string
}

func (InvalidPasswordEvent) EventTypeName() string {
	return "InvalidPassword"
}

// ChallengeSentEvent Raised when an Asterisk service sends an authentication challenge to a request.
type ChallengeSentEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// The challenge that was sent.
	Challenge string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (ChallengeSentEvent) EventTypeName() string {
	return "ChallengeSent"
}

// InvalidTransportEvent Raised when a request attempts to use a transport not allowed by the Asterisk service.
type InvalidTransportEvent struct {

	// The time the event was detected.
	EventTV string

	// A relative severity of the security event.
	Severity string

	// The Asterisk service that raised the security event.
	Service string

	// The version of this event.
	EventVersion string

	// The Service account associated with the security event
	//               notification.
	AccountID string

	// A unique identifier for the session in the service
	//               that raised the event.
	SessionID string

	// The address of the Asterisk service that raised the
	//               security event.
	LocalAddress string

	// The remote address of the entity that caused the
	//               security event to be raised.
	RemoteAddress string

	// The transport type that the request attempted to use.
	AttemptedTransport string

	// If available, the name of the module that raised the event.
	Module string

	// The timestamp reported by the session.
	SessionTV string
}

func (InvalidTransportEvent) EventTypeName() string {
	return "InvalidTransport"
}

// StatusEvent Raised in response to a Status command.
type StatusEvent struct {
	ActionID string

	// Type of channel
	Type string

	// Dialed number identifier
	DNID string

	EffectiveConnectedLineNum string

	EffectiveConnectedLineName string

	// Absolute lifetime of the channel
	TimeToHangup string

	// Identifier of the bridge the channel is in, may be empty if not in one
	BridgeID string

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

func (StatusEvent) EventTypeName() string {
	return "Status"
}

// StatusCompleteEvent Raised in response to a Status command.
type StatusCompleteEvent struct {

	// Number of Status events returned
	Items string
}

func (StatusCompleteEvent) EventTypeName() string {
	return "StatusComplete"
}

// OriginateResponseEvent Raised in response to an Originate command.
type OriginateResponseEvent struct {
	ActionID string

	Response string

	Channel string

	Context string

	Exten string

	Application string

	Data string

	Reason string

	Uniqueid string

	CallerIDNum string

	CallerIDName string
}

func (OriginateResponseEvent) EventTypeName() string {
	return "OriginateResponse"
}

// CoreShowChannelEvent Raised in response to a CoreShowChannels command.
type CoreShowChannelEvent struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Identifier of the bridge the channel is in, may be empty if not in one
	BridgeId string

	// Application currently executing on the channel
	Application string

	// Data given to the currently executing application
	ApplicationData string

	// The amount of time the channel has existed
	Duration string
}

func (CoreShowChannelEvent) EventTypeName() string {
	return "CoreShowChannel"
}

// CoreShowChannelsCompleteEvent Raised at the end of the CoreShowChannel list produced by the CoreShowChannels command.
type CoreShowChannelsCompleteEvent struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Conveys the status of the command reponse list
	EventList string

	// The total number of list items produced
	ListItems string
}

func (CoreShowChannelsCompleteEvent) EventTypeName() string {
	return "CoreShowChannelsComplete"
}

// ExtensionStatusEvent Raised when a hint changes due to a device state change.
type ExtensionStatusEvent struct {

	// Name of the extension.
	Exten string

	// Context that owns the extension.
	Context string

	// Hint set for the extension
	Hint string

	// Numerical value of the extension status. Extension
	//               status is determined by the combined device state of all items
	//               contained in the hint.
	Status string

	// Text representation of .
	StatusText string
}

func (ExtensionStatusEvent) EventTypeName() string {
	return "ExtensionStatus"
}

// PresenceStatusEvent Raised when a hint changes due to a presence state change.
type PresenceStatusEvent struct {
	Exten string

	Context string

	Hint string

	Status string

	Subtype string

	Message string
}

func (PresenceStatusEvent) EventTypeName() string {
	return "PresenceStatus"
}

// LogChannelEvent Raised when a logging channel is re-enabled after a reload operation.
type LogChannelEvent struct {

	// The name of the logging channel.
	Channel string

	Enabled string

	Reason string
}

func (LogChannelEvent) EventTypeName() string {
	return "LogChannel"
}

// PresenceStateChangeEvent Raised when a presence state changes
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

func (PresenceStateChangeEvent) EventTypeName() string {
	return "PresenceStateChange"
}

// BlindTransferEvent Raised when a blind transfer is complete.
type BlindTransferEvent struct {

	// Indicates if the transfer was successful or if it failed.
	Result string

	// Indicates if the transfer was performed outside of Asterisk. For instance,
	//               a channel protocol native transfer is external. A DTMF transfer is internal.
	IsExternal string

	// Destination context for the blind transfer.
	Context string

	// Destination extension for the blind transfer.
	Extension string
}

func (BlindTransferEvent) EventTypeName() string {
	return "BlindTransfer"
}

// AttendedTransferEvent Raised when an attended transfer is complete.
type AttendedTransferEvent struct {

	// Indicates if the transfer was successful or if it failed.
	Result string

	// Indicates the method by which the attended transfer completed.
	DestType string

	// Indicates the surviving bridge when bridges were merged to complete the transfer
	DestBridgeUniqueid string

	// Indicates the application that is running when the transfer completes
	DestApp string

	// The name of the surviving transferer channel when a transfer results in a threeway call
	DestTransfererChannel string
}

func (AttendedTransferEvent) EventTypeName() string {
	return "AttendedTransfer"
}

// PeerStatusEvent Raised when the state of a peer changes.
type PeerStatusEvent struct {

	// The channel technology of the peer.
	ChannelType string

	// The name of the peer (including channel technology).
	Peer string

	// New status of the peer.
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

func (PeerStatusEvent) EventTypeName() string {
	return "PeerStatus"
}

// ContactStatusEvent Raised when the state of a contact changes.
type ContactStatusEvent struct {

	// This contact's URI.
	URI string

	// New status of the contact.
	ContactStatus string

	// The name of the associated aor.
	AOR string

	// The name of the associated endpoint.
	EndpointName string

	// The RTT measured during the last qualify.
	RoundtripUsec string
}

func (ContactStatusEvent) EventTypeName() string {
	return "ContactStatus"
}

// AgentLoginEvent Raised when an Agent has logged in.
type AgentLoginEvent struct {

	// Agent ID of the agent.
	Agent string
}

func (AgentLoginEvent) EventTypeName() string {
	return "AgentLogin"
}

// AgentLogoffEvent Raised when an Agent has logged off.
type AgentLogoffEvent struct {

	// Agent ID of the agent.
	Agent string

	// The number of seconds the agent was logged in.
	Logintime string
}

func (AgentLogoffEvent) EventTypeName() string {
	return "AgentLogoff"
}

// ChannelTalkingStartEvent Raised when talking is detected on a channel.
type ChannelTalkingStartEvent struct {
}

func (ChannelTalkingStartEvent) EventTypeName() string {
	return "ChannelTalkingStart"
}

// ChannelTalkingStopEvent Raised when talking is no longer detected on a channel.
type ChannelTalkingStopEvent struct {

	// The length in time, in milliseconds, that talking was
	//               detected on the channel.
	Duration string
}

func (ChannelTalkingStopEvent) EventTypeName() string {
	return "ChannelTalkingStop"
}

// ReloadEvent Raised when a module has been reloaded in Asterisk.
type ReloadEvent struct {

	// The name of the module that was reloaded, or
	//
	//               if all modules were reloaded
	Module string

	// The numeric status code denoting the success or failure
	//               of the reload request.
	Status string
}

func (ReloadEvent) EventTypeName() string {
	return "Reload"
}

// LoadEvent Raised when a module has been loaded in Asterisk.
type LoadEvent struct {

	// The name of the module that was loaded
	Module string

	// The result of the load request.
	Status string
}

func (LoadEvent) EventTypeName() string {
	return "Load"
}

// UnloadEvent Raised when a module has been unloaded in Asterisk.
type UnloadEvent struct {

	// The name of the module that was unloaded
	Module string

	// The result of the unload request.
	Status string
}

func (UnloadEvent) EventTypeName() string {
	return "Unload"
}

// UserEventEvent A user defined event raised from the dialplan.
type UserEventEvent struct {

	// The event name, as specified in the dialplan.
	UserEvent string
}

func (UserEventEvent) EventTypeName() string {
	return "UserEvent"
}

// RegistryEvent Raised when an outbound registration completes.
type RegistryEvent struct {

	// The type of channel that was registered (or not).
	ChannelType string

	// The username portion of the registration.
	Username string

	// The address portion of the registration.
	Domain string

	// The status of the registration request.
	Status string

	// What caused the rejection of the request, if available.
	Cause string
}

func (RegistryEvent) EventTypeName() string {
	return "Registry"
}

// DeviceStateChangeEvent Raised when a device state changes
type DeviceStateChangeEvent struct {

	// The device whose state has changed
	Device string

	// The new state of the device
	State string
}

func (DeviceStateChangeEvent) EventTypeName() string {
	return "DeviceStateChange"
}

// MessageWaitingEvent Raised when the state of messages in a voicemail mailbox
//         has changed or when a channel has finished interacting with a
//         mailbox.
type MessageWaitingEvent struct {

	// The mailbox with the new message, specified as @
	Mailbox string

	// Whether or not the mailbox has messages waiting for it.
	Waiting string

	// The number of new messages.
	New string

	// The number of old messages.
	Old string
}

func (MessageWaitingEvent) EventTypeName() string {
	return "MessageWaiting"
}

// FullyBootedEvent Raised when all Asterisk initialization procedures have finished.
type FullyBootedEvent struct {

	// Informational message
	Status string

	// Seconds since start
	Uptime string

	// Seconds since last reload
	LastReload string
}

func (FullyBootedEvent) EventTypeName() string {
	return "FullyBooted"
}

// ShutdownEvent Raised when Asterisk is shutdown or restarted.
type ShutdownEvent struct {

	// Whether the shutdown is proceeding cleanly (all channels
	//               were hungup successfully) or uncleanly (channels will be
	//               terminated)
	Shutdown string

	// Whether or not a restart will occur.
	Restart string
}

func (ShutdownEvent) EventTypeName() string {
	return "Shutdown"
}

// LocalBridgeEvent Raised when two halves of a Local Channel form a bridge.
type LocalBridgeEvent struct {

	// The context in the dialplan that Channel2 starts in.
	Context string

	// The extension in the dialplan that Channel2 starts in.
	Exten string

	LocalOptimization string
}

func (LocalBridgeEvent) EventTypeName() string {
	return "LocalBridge"
}

// LocalOptimizationBeginEvent Raised when two halves of a Local Channel begin to optimize
//         themselves out of the media path.
type LocalOptimizationBeginEvent struct {

	// The unique ID of the bridge into which the local channel is optimizing.
	DestUniqueId string

	// Identification for the optimization operation.
	Id string
}

func (LocalOptimizationBeginEvent) EventTypeName() string {
	return "LocalOptimizationBegin"
}

// LocalOptimizationEndEvent Raised when two halves of a Local Channel have finished optimizing
//         themselves out of the media path.
type LocalOptimizationEndEvent struct {

	// Indicates whether the local optimization succeeded.
	Success string

	// Identification for the optimization operation. Matches the
	//
	//               from a previous
	Id string
}

func (LocalOptimizationEndEvent) EventTypeName() string {
	return "LocalOptimizationEnd"
}

// RenameEvent Raised when the name of a channel is changed.
type RenameEvent struct {
	Channel string

	Newname string

	Uniqueid string
}

func (RenameEvent) EventTypeName() string {
	return "Rename"
}

// RTCPSentEvent Raised when an RTCP packet is sent.
type RTCPSentEvent struct {

	// The SSRC identifier for our stream
	SSRC string

	// The type of packet for this RTCP report.
	PT string

	// The address the report is sent to.
	To string

	// The number of reports that were sent.
	//   The report count determines the number of ReportX headers in
	//               the message. The X for each set of report headers will range from 0 to
	//               .
	ReportCount string

	// The time the sender generated the report. Only valid when
	//               PT is .
	SentNTP string

	// The sender's last RTP timestamp. Only valid when PT is
	//               .
	SentRTP string

	// The number of packets the sender has sent. Only valid when PT
	//               is .
	SentPackets string

	// The number of bytes the sender has sent. Only valid when PT is
	//               .
	SentOctets string

	// The SSRC for the source of this report block.
	ReportXSourceSSRC string

	// The fraction of RTP data packets from
	//
	//               lost since the previous SR or RR report was sent.
	ReportXFractionLost string

	// The total number of RTP data packets from
	//
	//               lost since the beginning of reception.
	ReportXCumulativeLost string

	// The highest sequence number received in an RTP data packet from
	//               .
	ReportXHighestSequence string

	// The number of sequence number cycles seen for the RTP data
	//               received from .
	ReportXSequenceNumberCycles string

	// An estimate of the statistical variance of the RTP data packet
	//               interarrival time, measured in timestamp units.
	ReportXIAJitter string

	// The last SR timestamp received from .
	//               If no SR has been received from ,
	//               then 0.
	ReportXLSR string

	// The delay, expressed in units of 1/65536 seconds, between
	//               receiving the last SR packet from
	//
	//               and sending this report.
	ReportXDLSR string
}

func (RTCPSentEvent) EventTypeName() string {
	return "RTCPSent"
}

// RTCPReceivedEvent Raised when an RTCP packet is received.
type RTCPReceivedEvent struct {

	// The SSRC identifier for the remote system
	SSRC string

	// The type of packet for this RTCP report.
	PT string

	// The address the report was received from.
	From string

	// Calculated Round-Trip Time in seconds
	RTT string

	// The number of reports that were received.
	//   The report count determines the number of ReportX headers in
	//               the message. The X for each set of report headers will range from 0 to
	//               .
	ReportCount string

	// The time the sender generated the report. Only valid when
	//               PT is .
	SentNTP string

	// The sender's last RTP timestamp. Only valid when PT is
	//               .
	SentRTP string

	// The number of packets the sender has sent. Only valid when PT
	//               is .
	SentPackets string

	// The number of bytes the sender has sent. Only valid when PT is
	//               .
	SentOctets string

	// The SSRC for the source of this report block.
	ReportXSourceSSRC string

	// The fraction of RTP data packets from
	//
	//               lost since the previous SR or RR report was sent.
	ReportXFractionLost string

	// The total number of RTP data packets from
	//
	//               lost since the beginning of reception.
	ReportXCumulativeLost string

	// The highest sequence number received in an RTP data packet from
	//               .
	ReportXHighestSequence string

	// The number of sequence number cycles seen for the RTP data
	//               received from .
	ReportXSequenceNumberCycles string

	// An estimate of the statistical variance of the RTP data packet
	//               interarrival time, measured in timestamp units.
	ReportXIAJitter string

	// The last SR timestamp received from .
	//               If no SR has been received from ,
	//               then 0.
	ReportXLSR string

	// The delay, expressed in units of 1/65536 seconds, between
	//               receiving the last SR packet from
	//
	//               and sending this report.
	ReportXDLSR string
}

func (RTCPReceivedEvent) EventTypeName() string {
	return "RTCPReceived"
}

// NewchannelEvent Raised when a new channel is created.
type NewchannelEvent struct {
}

func (NewchannelEvent) EventTypeName() string {
	return "Newchannel"
}

// NewstateEvent Raised when a channel's state changes.
type NewstateEvent struct {
}

func (NewstateEvent) EventTypeName() string {
	return "Newstate"
}

// HangupEvent Raised when a channel is hung up.
type HangupEvent struct {

	// A numeric cause code for why the channel was hung up.
	Cause string

	// A description of why the channel was hung up.
	Causetxt string
}

func (HangupEvent) EventTypeName() string {
	return "Hangup"
}

// HangupRequestEvent Raised when a hangup is requested.
type HangupRequestEvent struct {

	// A numeric cause code for why the channel was hung up.
	Cause string
}

func (HangupRequestEvent) EventTypeName() string {
	return "HangupRequest"
}

// SoftHangupRequestEvent Raised when a soft hangup is requested with a specific cause code.
type SoftHangupRequestEvent struct {

	// A numeric cause code for why the channel was hung up.
	Cause string
}

func (SoftHangupRequestEvent) EventTypeName() string {
	return "SoftHangupRequest"
}

// NewExtenEvent Raised when a channel enters a new context, extension, priority.
type NewExtenEvent struct {

	// Deprecated in 12, but kept for
	//               backward compatability. Please use
	//               'Exten' instead.
	Extension string

	// The application about to be executed.
	Application string

	// The data to be passed to the application.
	AppData string
}

func (NewExtenEvent) EventTypeName() string {
	return "NewExten"
}

// NewCalleridEvent Raised when a channel receives new Caller ID information.
type NewCalleridEvent struct {

	// A description of the Caller ID presentation.
	CIDCallingPres string
}

func (NewCalleridEvent) EventTypeName() string {
	return "NewCallerid"
}

// NewConnectedLineEvent Raised when a channel's connected line information is changed.
type NewConnectedLineEvent struct {
}

func (NewConnectedLineEvent) EventTypeName() string {
	return "NewConnectedLine"
}

// NewAccountCodeEvent Raised when a Channel's AccountCode is changed.
type NewAccountCodeEvent struct {

	// The channel's previous account code
	OldAccountCode string
}

func (NewAccountCodeEvent) EventTypeName() string {
	return "NewAccountCode"
}

// DialBeginEvent Raised when a dial action has started.
type DialBeginEvent struct {

	// The non-technology specific device being dialed.
	DialString string
}

func (DialBeginEvent) EventTypeName() string {
	return "DialBegin"
}

// DialStateEvent Raised when dial status has changed.
type DialStateEvent struct {

	// The new state of the outbound dial attempt.
	DialStatus string

	// If the call was forwarded, where the call was
	//               forwarded to.
	Forward string
}

func (DialStateEvent) EventTypeName() string {
	return "DialState"
}

// DialEndEvent Raised when a dial action has completed.
type DialEndEvent struct {

	// The result of the dial operation.
	DialStatus string

	// If the call was forwarded, where the call was
	//               forwarded to.
	Forward string
}

func (DialEndEvent) EventTypeName() string {
	return "DialEnd"
}

// HoldEvent Raised when a channel goes on hold.
type HoldEvent struct {

	// The suggested MusicClass, if provided.
	MusicClass string
}

func (HoldEvent) EventTypeName() string {
	return "Hold"
}

// UnholdEvent Raised when a channel goes off hold.
type UnholdEvent struct {
}

func (UnholdEvent) EventTypeName() string {
	return "Unhold"
}

// ChanSpyStartEvent Raised when one channel begins spying on another channel.
type ChanSpyStartEvent struct {
}

func (ChanSpyStartEvent) EventTypeName() string {
	return "ChanSpyStart"
}

// ChanSpyStopEvent Raised when a channel has stopped spying.
type ChanSpyStopEvent struct {
}

func (ChanSpyStopEvent) EventTypeName() string {
	return "ChanSpyStop"
}

// HangupHandlerRunEvent Raised when a hangup handler is about to be called.
type HangupHandlerRunEvent struct {

	// Hangup handler parameter string passed to the Gosub application.
	Handler string
}

func (HangupHandlerRunEvent) EventTypeName() string {
	return "HangupHandlerRun"
}

// HangupHandlerPopEvent Raised when a hangup handler is removed from the handler stack
//         by the CHANNEL() function.
type HangupHandlerPopEvent struct {

	// Hangup handler parameter string passed to the Gosub application.
	Handler string
}

func (HangupHandlerPopEvent) EventTypeName() string {
	return "HangupHandlerPop"
}

// HangupHandlerPushEvent Raised when a hangup handler is added to the handler stack by
//         the CHANNEL() function.
type HangupHandlerPushEvent struct {

	// Hangup handler parameter string passed to the Gosub application.
	Handler string
}

func (HangupHandlerPushEvent) EventTypeName() string {
	return "HangupHandlerPush"
}

// FAXStatusEvent Raised periodically during a fax transmission.
type FAXStatusEvent struct {
	Operation string

	// A text message describing the current status of the fax
	Status string

	// The value of the  channel variable
	LocalStationID string

	// The files being affected by the fax operation
	FileName string
}

func (FAXStatusEvent) EventTypeName() string {
	return "FAXStatus"
}

// ReceiveFAXEvent Raised when a receive fax operation has completed.
type ReceiveFAXEvent struct {

	// The value of the  channel variable
	LocalStationID string

	// The value of the  channel variable
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

func (ReceiveFAXEvent) EventTypeName() string {
	return "ReceiveFAX"
}

// SendFAXEvent Raised when a send fax operation has completed.
type SendFAXEvent struct {

	// The value of the  channel variable
	LocalStationID string

	// The value of the  channel variable
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

func (SendFAXEvent) EventTypeName() string {
	return "SendFAX"
}

// MusicOnHoldStartEvent Raised when music on hold has started on a channel.
type MusicOnHoldStartEvent struct {

	// The class of music being played on the channel
	Class string
}

func (MusicOnHoldStartEvent) EventTypeName() string {
	return "MusicOnHoldStart"
}

// MusicOnHoldStopEvent Raised when music on hold has stopped on a channel.
type MusicOnHoldStopEvent struct {
}

func (MusicOnHoldStopEvent) EventTypeName() string {
	return "MusicOnHoldStop"
}

// MonitorStartEvent Raised when monitoring has started on a channel.
type MonitorStartEvent struct {
}

func (MonitorStartEvent) EventTypeName() string {
	return "MonitorStart"
}

// MonitorStopEvent Raised when monitoring has stopped on a channel.
type MonitorStopEvent struct {
}

func (MonitorStopEvent) EventTypeName() string {
	return "MonitorStop"
}

// DTMFBeginEvent Raised when a DTMF digit has started on a channel.
type DTMFBeginEvent struct {

	// DTMF digit received or transmitted (0-9, A-E, # or *
	Digit string

	Direction string
}

func (DTMFBeginEvent) EventTypeName() string {
	return "DTMFBegin"
}

// DTMFEndEvent Raised when a DTMF digit has ended on a channel.
type DTMFEndEvent struct {

	// DTMF digit received or transmitted (0-9, A-E, # or *
	Digit string

	// Duration (in milliseconds) DTMF was sent/received
	DurationMs string

	Direction string
}

func (DTMFEndEvent) EventTypeName() string {
	return "DTMFEnd"
}

// AOCSEvent Raised when an Advice of Charge message is sent at the beginning of a call.
type AOCSEvent struct {
	Chargeable string

	RateType string

	Currency string

	Name string

	Cost string

	Multiplier string

	ChargingType string

	StepFunction string

	Granularity string

	Length string

	Scale string

	Unit string

	SpecialCode string
}

func (AOCSEvent) EventTypeName() string {
	return "AOCS"
}

// AOCDEvent Raised when an Advice of Charge message is sent during a call.
type AOCDEvent struct {
	Charge string

	Type string

	BillingID string

	TotalType string

	Currency string

	Name string

	Cost string

	Multiplier string

	Units string

	NumberOf string

	TypeOf string
}

func (AOCDEvent) EventTypeName() string {
	return "AOCD"
}

// AOCEEvent Raised when an Advice of Charge message is sent at the end of a call.
type AOCEEvent struct {
	ChargingAssociation string

	Number string

	Plan string

	ID string

	Charge string

	Type string

	BillingID string

	TotalType string

	Currency string

	Name string

	Cost string

	Multiplier string

	Units string

	NumberOf string

	TypeOf string
}

func (AOCEEvent) EventTypeName() string {
	return "AOCE"
}

// PickupEvent Raised when a call pickup occurs.
type PickupEvent struct {
}

func (PickupEvent) EventTypeName() string {
	return "Pickup"
}

// BridgeCreateEvent Raised when a bridge is created.
type BridgeCreateEvent struct {
}

func (BridgeCreateEvent) EventTypeName() string {
	return "BridgeCreate"
}

// BridgeDestroyEvent Raised when a bridge is destroyed.
type BridgeDestroyEvent struct {
}

func (BridgeDestroyEvent) EventTypeName() string {
	return "BridgeDestroy"
}

// BridgeEnterEvent Raised when a channel enters a bridge.
type BridgeEnterEvent struct {

	// The uniqueid of the channel being swapped out of the bridge
	SwapUniqueid string
}

func (BridgeEnterEvent) EventTypeName() string {
	return "BridgeEnter"
}

// BridgeLeaveEvent Raised when a channel leaves a bridge.
type BridgeLeaveEvent struct {
}

func (BridgeLeaveEvent) EventTypeName() string {
	return "BridgeLeave"
}

// BridgeVideoSourceUpdateEvent Raised when the channel that is the source of video in a bridge changes.
type BridgeVideoSourceUpdateEvent struct {

	// The unique ID of the channel that was the video source.
	BridgePreviousVideoSource string
}

func (BridgeVideoSourceUpdateEvent) EventTypeName() string {
	return "BridgeVideoSourceUpdate"
}

// BridgeMergeEvent Raised when two bridges are merged.
type BridgeMergeEvent struct {
}

func (BridgeMergeEvent) EventTypeName() string {
	return "BridgeMerge"
}

// ParkedCallEvent Raised when a channel is parked.
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

func (ParkedCallEvent) EventTypeName() string {
	return "ParkedCall"
}

// ParkedCallTimeOutEvent Raised when a channel leaves a parking lot due to reaching the time limit of being parked.
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

func (ParkedCallTimeOutEvent) EventTypeName() string {
	return "ParkedCallTimeOut"
}

// ParkedCallGiveUpEvent Raised when a channel leaves a parking lot because it hung up without being answered.
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

func (ParkedCallGiveUpEvent) EventTypeName() string {
	return "ParkedCallGiveUp"
}

// UnParkedCallEvent Raised when a channel leaves a parking lot because it was retrieved from the parking lot and
//         reconnected.
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

func (UnParkedCallEvent) EventTypeName() string {
	return "UnParkedCall"
}

// ParkedCallSwapEvent Raised when a channel takes the place of a previously parked channel
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

func (ParkedCallSwapEvent) EventTypeName() string {
	return "ParkedCallSwap"
}

// MWIGetEvent Raised in response to a MWIGet command.
type MWIGetEvent struct {
	ActionID string

	// Specific mailbox ID.
	Mailbox string

	// The number of old messages in the mailbox.
	OldMessages string

	// The number of new messages in the mailbox.
	NewMessages string
}

func (MWIGetEvent) EventTypeName() string {
	return "MWIGet"
}

// MWIGetCompleteEvent Raised in response to a MWIGet command.
type MWIGetCompleteEvent struct {
	ActionID string

	EventList string

	// The number of mailboxes reported.
	ListItems string
}

func (MWIGetCompleteEvent) EventTypeName() string {
	return "MWIGetComplete"
}

// IdentifyDetailEvent Provide details about an identify section.
type IdentifyDetailEvent struct {

	// The object's type. This will always be 'identify'.
	ObjectType string

	// The name of this object.
	ObjectName string

	// Name of endpoint identified
	Endpoint string

	// Perform SRV lookups for provided hostnames.
	SrvLookups string

	// IP addresses or networks to match against.
	Match string

	// Header/value pair to match against.
	MatchHeader string

	// The name of the endpoint associated with this information.
	EndpointName string
}

func (IdentifyDetailEvent) EventTypeName() string {
	return "IdentifyDetail"
}

// AorDetailEvent Provide details about an Address of Record (AoR) section.
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

	// Authenticates a qualify challenge response if needed
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

func (AorDetailEvent) EventTypeName() string {
	return "AorDetail"
}

// AuthDetailEvent Provide details about an authentication section.
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

func (AuthDetailEvent) EventTypeName() string {
	return "AuthDetail"
}

// TransportDetailEvent Provide details about an authentication section.
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

	// File containing a list of certificates to read (TLS ONLY, not WSS)
	CaListFile string

	// Path to directory containing a list of certificates to read (TLS ONLY, not WSS)
	CaListPath string

	// Certificate file for endpoint (TLS ONLY, not WSS)
	CertFile string

	// Private key file (TLS ONLY, not WSS)
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

	// Require verification of server certificate (TLS ONLY, not WSS)
	VerifyServer string

	// Require verification of client certificate (TLS ONLY, not WSS)
	VerifyClient string

	// Require client certificate (TLS ONLY, not WSS)
	RequireClientCert string

	// Method of SSL transport (TLS ONLY, not WSS)
	Method string

	// Preferred cryptography cipher names (TLS ONLY, not WSS)
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

func (TransportDetailEvent) EventTypeName() string {
	return "TransportDetail"
}

// EndpointDetailEvent Provide details about an endpoint section.
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

	// Use Endpoint's requested packetization interval
	UsePtime string

	// Force use of return port
	ForceRport string

	// Allow Contact header to be rewritten with the source IP address-port
	RewriteContact string

	// Explicit transport configuration to use
	Transport string

	// Full SIP URI of the outbound proxy used to send requests
	OutboundProxy string

	// Default Music On Hold class
	MohSuggest string

	// Allow support for RFC3262 provisional ACK tags
	Field100rel string

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

	// Way(s) for the endpoint to be identified
	IdentifyBy string

	// Determines whether media may flow directly between endpoints.
	DirectMedia string

	// Direct Media method type
	DirectMediaMethod string

	// Accept Connected Line updates from this endpoint
	TrustConnectedLine string

	// Send Connected Line updates to this endpoint
	SendConnectedLine string

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

	// Send the Diversion header, conveying the diversion
	//               information to the called user agent
	SendDiversion string

	// NOTIFY the endpoint when state changes for any of the specified mailboxes
	Mailboxes string

	// Condense MWI notifications into a single NOTIFY.
	AggregateMwi string

	// Determines whether res_pjsip will use and enforce usage of media encryption
	//               for this endpoint.
	MediaEncryption string

	// Determines whether encryption should be used if possible but does not terminate the
	//               session if not achieved.
	MediaEncryptionOptimistic string

	// Determines whether res_pjsip will use and enforce usage of AVPF for this
	//               endpoint.
	UseAvpf string

	// Determines whether res_pjsip will use and enforce usage of AVP,
	//               regardless of the RTP profile in use for this endpoint.
	ForceAvp string

	// Determines whether res_pjsip will use the media transport received in the
	//               offer SDP in the corresponding answer SDP.
	MediaUseReceivedTransport string

	// Determines whether one-touch recording is allowed for this endpoint.
	OneTouchRecording string

	// Determines whether chan_pjsip will indicate ringing using inband
	//               progress.
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

	// Determines whether a user=phone parameter is placed into the request URI if the user is determined
	//               to be a phone number
	UserEqPhone string

	// Determines whether hold and unhold will be passed through using re-INVITEs with recvonly and
	//               sendrecv to the remote side
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

	// Respond to a SIP invite with the single most preferred codec (DEPRECATED)
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

func (EndpointDetailEvent) EventTypeName() string {
	return "EndpointDetail"
}

// AorListEvent Provide details about an Address of Record (AoR) section.
type AorListEvent struct {

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

	// Authenticates a qualify challenge response if needed
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
}

func (AorListEvent) EventTypeName() string {
	return "AorList"
}

// AuthListEvent Provide details about an Address of Record (Auth) section.
type AuthListEvent struct {

	// The object's type. This will always be 'auth'.
	ObjectType string

	// The name of this object.
	ObjectName string

	// Username to use for account
	Username string

	// MD5 Hash used for authentication.
	Md5Cred string

	// SIP realm for endpoint
	Realm string

	// Authentication type
	AuthType string

	// Plain text password used for authentication.
	Password string

	// Lifetime of a nonce associated with this authentication config.
	NonceLifetime string
}

func (AuthListEvent) EventTypeName() string {
	return "AuthList"
}

// ContactListEvent Provide details about a contact section.
type ContactListEvent struct {

	// The object's type. This will always be 'contact'.
	ObjectType string

	// The name of this object.
	ObjectName string

	// IP address of the last Via header in REGISTER request.
	//               Will only appear in the event if available.
	ViaAddr string

	// Port number of the last Via header in REGISTER request.
	//               Will only appear in the event if available.
	ViaPort string

	// The elapsed time in decimal seconds after which an OPTIONS
	//               message is sent before the contact is considered unavailable.
	QualifyTimeout string

	// Content of the Call-ID header in REGISTER request.
	//               Will only appear in the event if available.
	CallId string

	// Asterisk Server name.
	RegServer string

	// If true delete the contact on Asterisk restart/boot.
	PruneOnBoot string

	// The Path header received on the REGISTER.
	Path string

	// The name of the endpoint associated with this information.
	Endpoint string

	// A boolean indicating whether a qualify should be authenticated.
	AuthenticateQualify string

	// This contact's URI.
	Uri string

	// The interval in seconds at which the contact will be qualified.
	QualifyFrequency string

	// Content of the User-Agent header in REGISTER request
	UserAgent string

	// Absolute time that this contact is no longer valid after
	ExpirationTime string

	// The contact's outbound proxy.
	OutboundProxy string

	// This contact's status.
	Status string

	// The round trip time in microseconds.
	RoundtripUsec string
}

func (ContactListEvent) EventTypeName() string {
	return "ContactList"
}

// ContactStatusDetailEvent Provide details about a contact's status.
type ContactStatusDetailEvent struct {

	// The AoR that owns this contact.
	AOR string

	// This contact's URI.
	URI string

	// This contact's status.
	Status string

	// The round trip time in microseconds.
	RoundtripUsec string

	// The name of the endpoint associated with this information.
	EndpointName string

	// Content of the User-Agent header in REGISTER request
	UserAgent string

	// Absolute time that this contact is no longer valid after
	RegExpire string

	// IP address:port of the last Via header in REGISTER request.
	//               Will only appear in the event if available.
	ViaAddress string

	// Content of the Call-ID header in REGISTER request.
	//               Will only appear in the event if available.
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

	// The elapsed time in decimal seconds after which an OPTIONS
	//               message is sent before the contact is considered unavailable.
	QualifyTimeout string
}

func (ContactStatusDetailEvent) EventTypeName() string {
	return "ContactStatusDetail"
}

// EndpointListEvent Provide details about a contact's status.
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

func (EndpointListEvent) EventTypeName() string {
	return "EndpointList"
}

// FAXSessionsEntryEvent A single list item for the FAXSessions AMI command
type FAXSessionsEntryEvent struct {
	ActionID string

	// Name of the channel responsible for the FAX session
	Channel string

	// The FAX technology that the FAX session is using
	Technology string

	// The numerical identifier for this particular session
	SessionNumber string

	// FAX session passthru/relay type
	SessionType string

	// FAX session operation type
	Operation string

	// Current state of the FAX session
	State string

	// File or list of files associated with this FAX session
	Files string
}

func (FAXSessionsEntryEvent) EventTypeName() string {
	return "FAXSessionsEntry"
}

// FAXSessionsCompleteEvent Raised when all FAXSession events are completed for a FAXSessions command
type FAXSessionsCompleteEvent struct {
	ActionID string

	// Count of FAXSession events sent in response to FAXSessions action
	Total string
}

func (FAXSessionsCompleteEvent) EventTypeName() string {
	return "FAXSessionsComplete"
}

// FAXSessionEvent Raised in response to FAXSession manager command
type FAXSessionEvent struct {
	ActionID string

	// The numerical identifier for this particular session
	SessionNumber string

	// FAX session operation type
	Operation string

	// Current state of the FAX session
	State string

	// Whether error correcting mode is enabled for the FAX session. This field is not
	//               included when operation is 'V.21 Detect' or if operation is 'gateway' and state is
	//               'Uninitialized'
	ErrorCorrectionMode string

	// Bit rate of the FAX. This field is not included when operation is 'V.21 Detect' or
	//               if operation is 'gateway' and state is 'Uninitialized'.
	DataRate string

	// Resolution of each page of the FAX. Will be in the format of X_RESxY_RES. This field
	//               is not included if the operation is anything other than Receive/Transmit.
	ImageResolution string

	// Current number of pages transferred during this FAX session. May change as the FAX
	//               progresses. This field is not included when operation is 'V.21 Detect' or if operation is
	//               'gateway' and state is 'Uninitialized'.
	PageNumber string

	// Filename of the image being sent/received for this FAX session. This field is not
	//               included if Operation isn't 'send' or 'receive'.
	FileName string

	// Total number of pages sent during this session. This field is not included if
	//               Operation isn't 'send' or 'receive'. Will always be 0 for 'receive'.
	PagesTransmitted string

	// Total number of pages received during this session. This field is not included if
	//               Operation is not 'send' or 'receive'. Will be 0 for 'send'.
	PagesReceived string

	// Total number of bad lines sent/received during this session. This field is not
	//               included if Operation is not 'send' or 'received'.
	TotalBadLines string
}

func (FAXSessionEvent) EventTypeName() string {
	return "FAXSession"
}

// FAXStatsEvent Raised in response to FAXStats manager command
type FAXStatsEvent struct {
	ActionID string

	// Number of active FAX sessions
	CurrentSessions string

	// Number of reserved FAX sessions
	ReservedSessions string

	// Total FAX sessions for which Asterisk is/was the transmitter
	TransmitAttempts string

	// Total FAX sessions for which Asterisk is/was the recipient
	ReceiveAttempts string

	// Total FAX sessions which have been completed successfully
	CompletedFAXes string

	// Total FAX sessions which failed to complete successfully
	FailedFAXes string
}

func (FAXStatsEvent) EventTypeName() string {
	return "FAXStats"
}

// AsyncAGIStartEvent Raised when a channel starts AsyncAGI command processing.
type AsyncAGIStartEvent struct {

	// URL encoded string read from the AsyncAGI server.
	Env string
}

func (AsyncAGIStartEvent) EventTypeName() string {
	return "AsyncAGIStart"
}

// AsyncAGIEndEvent Raised when a channel stops AsyncAGI command processing.
type AsyncAGIEndEvent struct {
}

func (AsyncAGIEndEvent) EventTypeName() string {
	return "AsyncAGIEnd"
}

// AsyncAGIExecEvent Raised when AsyncAGI completes an AGI command.
type AsyncAGIExecEvent struct {

	// Optional command ID sent by the AsyncAGI server to identify the command.
	CommandID string

	// URL encoded result string from the executed AGI command.
	Result string
}

func (AsyncAGIExecEvent) EventTypeName() string {
	return "AsyncAGIExec"
}

// AGIExecStartEvent Raised when a received AGI command starts processing.
type AGIExecStartEvent struct {

	// The AGI command as received from the external source.
	Command string

	// Random identification number assigned to the execution of this command.
	CommandId string
}

func (AGIExecStartEvent) EventTypeName() string {
	return "AGIExecStart"
}

// AGIExecEndEvent Raised when a received AGI command completes processing.
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

func (AGIExecEndEvent) EventTypeName() string {
	return "AGIExecEnd"
}

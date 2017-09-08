package ami

// Transfer DAHDI Channel.
// Simulate a flash hook event by the user connected to the channel.
// NOTE: Valid only for analog channels.
type DAHDITransferAction struct {
	// DAHDI channel number to transfer.
	//
	// reqtured
	DAHDIChannel string
}

// Hangup DAHDI Channel.
// Simulate an on-hook event by the user connected to the channel.
// NOTE: Valid only for analog channels.
type DAHDIHangupAction struct {
	// DAHDI channel number to hangup.
	//
	// reqtured
	DAHDIChannel string
}

// Dial over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
type DAHDIDialOffhookAction struct {
	// DAHDI channel number to dial digits.
	//
	// reqtured
	DAHDIChannel string
	// Digits to dial.
	//
	// reqtured
	Number string
}

// Toggle DAHDI channel Do Not Disturb status ON.
// Equivalent to the CLI command "dahdi set dnd `channel` on".
// NOTE: Feature only supported by analog channels.
type DAHDIDNDonAction struct {
	// DAHDI channel number to set DND on.
	//
	// reqtured
	DAHDIChannel string
}

// Toggle DAHDI channel Do Not Disturb status OFF.
// Equivalent to the CLI command "dahdi set dnd `channel` off".
// NOTE: Feature only supported by analog channels.
type DAHDIDNDoffAction struct {
	// DAHDI channel number to set DND off.
	//
	// reqtured
	DAHDIChannel string
}

// Show status of DAHDI channels.
// Similar to the CLI command "dahdi show channels".
type DAHDIShowChannelsAction struct {
	// Specify the specific channel number to show. Show all channels if zero or not present.
	DAHDIChannel string
}

// Fully Restart DAHDI channels (terminates calls).
// Equivalent to the CLI command "dahdi restart".
type DAHDIRestartAction struct {
}

// Show status of PRI spans.
// Similar to the CLI command "pri show spans".
type PRIShowSpansAction struct {
	// Specify the specific span to show. Show all spans if zero or not present.
	Span string
}

// Set PRI debug levels for a span
// Equivalent to the CLI command "pri set debug &lt;level&gt; span &lt;span&gt;".
type PRIDebugSetAction struct {
	// Which span to affect.
	//
	// reqtured
	Span string
	// What debug level to set. May be a numerical value or a text value from the list below
	//
	// reqtured
	// ENUM: off,on,hex,intense,
	Level string
}

// Set the file used for PRI debug message output
// Equivalent to the CLI command "pri set debug file &lt;output-file&gt;"
type PRIDebugFileSetAction struct {
	// Path of file to write debug output.
	//
	// reqtured
	File string
}

// Disables file output for PRI debug messages
type PRIDebugFileUnsetAction struct {
}

// List IAX peers.
type IAXpeersAction struct {
}

// List IAX Peers.
// List all the IAX peers.
type IAXpeerlistAction struct {
}

// Show IAX Netstats.
// Show IAX channels network statistics.
type IAXnetstatsAction struct {
}

// Show IAX registrations.
// Show IAX registrations.
type IAXregistryAction struct {
}

// List SIP peers (text format).
// Lists SIP peers in text format with details on current status. `Peerlist` will follow as separate events, followed by a final event called `PeerlistComplete`.
type SIPpeersAction struct {
}

// show SIP peer (text format).
// Show one SIP peer with details on current status.
type SIPshowpeerAction struct {
	// The peer name you want to check.
	//
	// reqtured
	Peer string
}

// Qualify SIP peers.
// Qualify a SIP peer.
//
// seealso SIPQualifyPeerDone
type SIPqualifypeerAction struct {
	// The peer name you want to qualify.
	//
	// reqtured
	Peer string
}

// Show SIP registrations (text format).
// Lists all registration requests and status. Registrations will follow as separate events followed by a final event called `RegistrationsComplete`.
type SIPshowregistryAction struct {
}

// Send a SIP notify.
// Sends a SIP Notify event.
// All parameters for this event must be specified in the body of this request via multiple `Variable: name=value` sequences.
type SIPnotifyAction struct {
	// Peer to receive the notify.
	//
	// reqtured
	Channel string
	// At least one variable pair must be specified. name=value
	//
	// reqtured
	Variable string
}

// Show the status of one or all of the sip peers.
// Retrieves the status of one or all of the sip peers. If no peer name is specified, status for all of the sip peers will be retrieved.
type SIPpeerstatusAction struct {
	// The peer name you want to check.
	Peer string
}

// List SKINNY devices (text format).
// Lists Skinny devices in text format with details on current status. Devicelist will follow as separate events, followed by a final event called DevicelistComplete.
type SKINNYdevicesAction struct {
}

// Show SKINNY device (text format).
// Show one SKINNY device with details on current status.
type SKINNYshowdeviceAction struct {
	// The device name you want to check.
	//
	// reqtured
	Device string
}

// List SKINNY lines (text format).
// Lists Skinny lines in text format with details on current status. Linelist will follow as separate events, followed by a final event called LinelistComplete.
type SKINNYlinesAction struct {
}

// Show SKINNY line (text format).
// Show one SKINNY line with details on current status.
type SKINNYshowlineAction struct {
	// The line name you want to check.
	//
	// reqtured
	Line string
}

// Add an extension to the dialplan
type DialplanExtensionAddAction struct {
	// Context where the extension will be created. The context will be created if it does not already exist.
	//
	// reqtured
	Context string
	// Name of the extension that will be created (may include callerid match by separating with '/')
	//
	// reqtured
	Extension string
	// Priority being added to this extension. Must be either `hint` or a numerical value.
	//
	// reqtured
	Priority string
	// The application to use for this extension at the requested priority
	//
	// reqtured
	Application string
	// Arguments to the application.
	ApplicationData string
	// If set to 'yes', '1', 'true' or any of the other values we evaluate as true, then if an extension already exists at the requested context, extension, and priority it will be overwritten. Otherwise, the existing extension will remain and the action will fail.
	Replace string
}

// Remove an extension from the dialplan
type DialplanExtensionRemoveAction struct {
	// Context of the extension being removed
	//
	// reqtured
	Context string
	// Name of the extension being removed (may include callerid match by separating with '/')
	//
	// reqtured
	Extension string
	// If provided, only remove this priority from the extension instead of all priorities in the extension.
	Priority string
}

// Lists agents and their status.
// Will list info about all defined agents.
//
// seealso Agents
// seealso AgentsComplete
type AgentsAction struct {
}

// Sets an agent as no longer logged in.
// Sets an agent as no longer logged in.
type AgentLogoffAction struct {
	// Agent ID of the agent to log off.
	//
	// reqtured
	Agent string
	// Set to `true` to not hangup existing calls.
	Soft string
}

// List participants in a conference.
// Lists all users in a particular ConfBridge conference. ConfbridgeList will follow as separate events, followed by a final event called ConfbridgeListComplete.
type ConfbridgeListAction struct {
	// Conference number.
	//
	// reqtured
	Conference string
}

// List active conferences.
// Lists data about all active conferences. ConfbridgeListRooms will follow as separate events, followed by a final event called ConfbridgeListRoomsComplete.
type ConfbridgeListRoomsAction struct {
}

// Mute a Confbridge user.
type ConfbridgeMuteAction struct {
	//
	// reqtured
	Conference string
	// If this parameter is not a complete channel name, the first channel with this prefix will be used.
	// If this parameter is "all", all channels will be muted.
	// If this parameter is "participants", all non-admin channels will be muted.
	//
	// reqtured
	Channel string
}

// Unmute a Confbridge user.
type ConfbridgeUnmuteAction struct {
	//
	// reqtured
	Conference string
	// If this parameter is not a complete channel name, the first channel with this prefix will be used.
	// If this parameter is "all", all channels will be unmuted.
	// If this parameter is "participants", all non-admin channels will be unmuted.
	//
	// reqtured
	Channel string
}

// Kick a Confbridge user.
type ConfbridgeKickAction struct {
	//
	// reqtured
	Conference string
	// If this parameter is "all", all channels will be kicked from the conference.
	// If this parameter is "participants", all non-admin channels will be kicked from the conference.
	//
	// reqtured
	Channel string
}

// Lock a Confbridge conference.
type ConfbridgeLockAction struct {
	//
	// reqtured
	Conference string
}

// Unlock a Confbridge conference.
type ConfbridgeUnlockAction struct {
	//
	// reqtured
	Conference string
}

// Start recording a Confbridge conference.
// Start recording a conference. If recording is already present an error will be returned. If RecordFile is not provided, the default record file specified in the conference's bridge profile will be used, if that is not present either a file will automatically be generated in the monitor directory.
type ConfbridgeStartRecordAction struct {
	//
	// reqtured
	Conference string
	RecordFile string
}

// Stop recording a Confbridge conference.
type ConfbridgeStopRecordAction struct {
	//
	// reqtured
	Conference string
}

// Set a conference user as the single video source distributed to all other participants.
type ConfbridgeSetSingleVideoSrcAction struct {
	//
	// reqtured
	Conference string
	// If this parameter is not a complete channel name, the first channel with this prefix will be used.
	//
	// reqtured
	Channel string
}

// Control the playback of a file being played to a channel.
// Control the operation of a media file being played back to a channel. Note that this AMI action does not initiate playback of media to channel, but rather controls the operation of a media operation that was already initiated on the channel.
// NOTE: The `pause` and `restart` Control options will stop a playback operation if that operation was not initiated from the ControlPlayback application or the control stream file AGI command.
//
// seealso Playback
// seealso ControlPlayback
// seealso stream file
// seealso control stream file
type ControlPlaybackAction struct {
	// The name of the channel that currently has a file being played back to it.
	//
	// reqtured
	Channel string
	//
	// reqtured
	// ENUM: stop,forward,reverse,pause,restart,
	Control string
}

// Mute a Meetme user.
type MeetmeMuteAction struct {
	//
	// reqtured
	Meetme string
	//
	// reqtured
	Usernum string
}

// Unmute a Meetme user.
type MeetmeUnmuteAction struct {
	//
	// reqtured
	Meetme string
	//
	// reqtured
	Usernum string
}

// List participants in a conference.
// Lists all users in a particular MeetMe conference. MeetmeList will follow as separate events, followed by a final event called MeetmeListComplete.
type MeetmeListAction struct {
	// Conference number.
	Conference string
}

// List active conferences.
// Lists data about all active conferences. MeetmeListRooms will follow as separate events, followed by a final event called MeetmeListRoomsComplete.
type MeetmeListRoomsAction struct {
}

// Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
type MixMonitorMuteAction struct {
	// Used to specify the channel to mute.
	//
	// reqtured
	Channel string
	// Which part of the recording to mute: read, write or both (from channel, to channel or both channels).
	Direction string
	// Turn mute on or off : 1 to turn on, 0 to turn off.
	State string
}

// Record a call and mix the audio during the recording. Use of StopMixMonitor is required
// to guarantee the audio file is available for processing during dialplan execution.
// This action records the audio on the current channel to the specified file.
// Vars:Var:MIXMONITOR_FILENAME ->
type MixMonitorAction struct {
	// Used to specify the channel to record.
	//
	// reqtured
	Channel string
	// Is the name of the file created in the monitor spool directory. Defaults to the same name as the channel (with slashes replaced with dashes). This argument is optional if you specify to record unidirectional audio with either the r(filename) or t(filename) options in the options field. If neither MIXMONITOR_FILENAME or this parameter is set, the mixed stream won't be recorded.
	File string
	// Options that apply to the MixMonitor in the same way as they would apply if invoked from the MixMonitor application. For a list of available options, see the documentation for the mixmonitor application.
	Options string `astgo:"options"`
	// Will be executed when the recording is over. Any strings matching `^{X}` will be unescaped to `X`. All variables will be evaluated at the time MixMonitor is called.
	Command string
}

// Stop recording a call through MixMonitor, and free the recording's file handle.
// This action stops the audio recording that was started with the `MixMonitor` action on the current channel.
type StopMixMonitorAction struct {
	// The name of the channel monitored.
	//
	// reqtured
	Channel string
	// If a valid ID is provided, then this command will stop only that specific MixMonitor.
	MixMonitorID string
}

// Show queue status.
// Check the status of one or more queues.
type QueueStatusAction struct {
	// Limit the response to the status of the specified queue.
	Queue string
	// Limit the response to the status of the specified member.
	Member string
}

// Show queue summary.
// Request the manager to send a QueueSummary event.
type QueueSummaryAction struct {
	// Queue for which the summary is requested.
	Queue string
}

// Add interface to queue.
type QueueAddAction struct {
	// Queue's name.
	//
	// reqtured
	Queue string
	// The name of the interface (tech/name) to add to the queue.
	//
	// reqtured
	Interface string
	// A penalty (number) to apply to this member. Asterisk will distribute calls to members with higher penalties only after attempting to distribute calls to those with lower penalty.
	Penalty string
	// To pause or not the member initially (true/false or 1/0).
	Paused string
	// Text alias for the interface.
	MemberName     string
	StateInterface string
}

// Remove interface from queue.
type QueueRemoveAction struct {
	// The name of the queue to take action on.
	//
	// reqtured
	Queue string
	// The interface (tech/name) to remove from queue.
	//
	// reqtured
	Interface string
}

// Makes a queue member temporarily unavailable.
// Pause or unpause a member in a queue.
type QueuePauseAction struct {
	// The name of the interface (tech/name) to pause or unpause.
	//
	// reqtured
	Interface string
	// Pause or unpause the interface. Set to 'true' to pause the member or 'false' to unpause.
	//
	// reqtured
	Paused string
	// The name of the queue in which to pause or unpause this member. If not specified, the member will be paused or unpaused in all the queues it is a member of.
	Queue string
	// Text description, returned in the event QueueMemberPaused.
	Reason string
}

// Adds custom entry in queue_log.
type QueueLogAction struct {
	//
	// reqtured
	Queue string
	//
	// reqtured
	Event     string
	Uniqueid  string
	Interface string
	Message   string
}

// Set the penalty for a queue member.
// Change the penalty of a queue member
type QueuePenaltyAction struct {
	// The interface (tech/name) of the member whose penalty to change.
	//
	// reqtured
	Interface string
	// The new penalty (number) for the member. Must be nonnegative.
	//
	// reqtured
	Penalty string
	// If specified, only set the penalty for the member of this queue. Otherwise, set the penalty for the member in all queues to which the member belongs.
	Queue string
}

// Set the ringinuse value for a queue member.
type QueueMemberRingInUseAction struct {
	//
	// reqtured
	Interface string
	//
	// reqtured
	RingInUse string
	Queue     string
}

// Queue Rules.
// List queue rules defined in queuerules.conf
type QueueRuleAction struct {
	// The name of the rule in queuerules.conf whose contents to list.
	Rule string
}

// Reload a queue, queues, or any sub-section of a queue or queues.
type QueueReloadAction struct {
	// The name of the queue to take action on. If no queue name is specified, then all queues are affected.
	Queue string
	// Whether to reload the queue's members.
	//
	// ENUM: yes,no,
	Members string
	// Whether to reload queuerules.conf
	//
	// ENUM: yes,no,
	Rules string
	// Whether to reload the other queue options.
	//
	// ENUM: yes,no,
	Parameters string
}

// Reset queue statistics.
// Reset the statistics for a queue.
type QueueResetAction struct {
	// The name of the queue on which to reset statistics.
	Queue string
}

// Change priority of a caller on queue.
type QueueChangePriorityCallerAction struct {
	// The name of the queue to take action on.
	//
	// reqtured
	Queue string
	// The caller (channel) to change priority on queue.
	//
	// reqtured
	Caller string
	// Priority value for change for caller on queue.
	//
	// reqtured
	Priority string
}

// Play DTMF signal on a specific channel.
// Plays a dtmf digit on the specified channel.
type PlayDTMFAction struct {
	// Channel name to send digit to.
	//
	// reqtured
	Channel string
	// The DTMF digit to play.
	//
	// reqtured
	Digit string
	// The duration, in milliseconds, of the digit to be played.
	Duration string
}

// List All Voicemail User Information.
type VoicemailUsersListAction struct {
}

// Tell Asterisk to poll mailboxes for a change
// Normally, MWI indicators are only sent when Asterisk itself changes a mailbox. With external programs that modify the content of a mailbox from outside the application, an option exists called `pollmailboxes` that will cause voicemail to continually scan all mailboxes on a system for changes. This can cause a large amount of load on a system. This command allows external applications to signal when a particular mailbox has changed, thus permitting external applications to modify mailboxes and MWI to work without introducing considerable CPU load.
// If Context is not specified, all mailboxes on the system will be polled for changes. If Context is specified, but Mailbox is omitted, then all mailboxes within Context will be polled. Otherwise, only a single mailbox will be polled for changes.
type VoicemailRefreshAction struct {
	Context string
	Mailbox string
}

// List available bridging technologies and their statuses.
// Returns detailed information about the available bridging technologies.
//
// seealso BridgeTechnologySuspend
// seealso BridgeTechnologyUnsuspend
type BridgeTechnologyListAction struct {
}

// Suspend a bridging technology.
// Marks a bridging technology as suspended, which prevents subsequently created bridges from using it.
//
// seealso BridgeTechnologySuspend
// seealso BridgeTechnologyUnsuspend
type BridgeTechnologySuspendAction struct {
	// The name of the bridging technology to suspend.
	//
	// reqtured
	BridgeTechnology string
}

// Unsuspend a bridging technology.
// Clears a previously suspended bridging technology, which allows subsequently created bridges to use it.
//
// seealso BridgeTechnologyList
// seealso BridgeTechnologySuspend
type BridgeTechnologyUnsuspendAction struct {
	// The name of the bridging technology to unsuspend.
	//
	// reqtured
	BridgeTechnology string
}

// Optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away. Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
type LocalOptimizeAwayAction struct {
	// The channel name to optimize away.
	//
	// reqtured
	Channel string
}

// Get DB Entry.
type DBGetAction struct {
	//
	// reqtured
	Family string
	//
	// reqtured
	Key string
}

// Put DB entry.
type DBPutAction struct {
	//
	// reqtured
	Family string
	//
	// reqtured
	Key string
	Val string
}

// Delete DB entry.
type DBDelAction struct {
	//
	// reqtured
	Family string
	//
	// reqtured
	Key string
}

// Delete DB Tree.
type DBDelTreeAction struct {
	//
	// reqtured
	Family string
	Key    string
}

// Bridge two channels already in the PBX.
// Bridge together two channels already in the PBX.
//
// seealso Bridge
// seealso BridgeCreate
// seealso BridgeEnter
// seealso BridgeDestroy
// seealso BridgeInfo
// seealso BridgeKick
// seealso BridgeList
type BridgeAction struct {
	// Channel to Bridge to Channel2.
	//
	// reqtured
	Channel1 string
	// Channel to Bridge to Channel1.
	//
	// reqtured
	Channel2 string
	// Play courtesy tone to Channel 2.
	//
	// ENUM: no,Channel1,Channel2,Both,
	Tone string
}

// Keepalive command.
// A 'Ping' action will ellicit a 'Pong' response. Used to keep the manager connection open.
type PingAction struct {
}

// Control Event Flow.
// Enable/Disable sending of events to this manager client.
type EventsAction struct {
	//
	// reqtured
	// ENUM: on,off,system,call,log,...,
	EventMask string
}

// Logoff Manager.
// Logoff the current manager session.
//
// seealso Login
type LogoffAction struct {
}

// Login Manager.
// Login Manager.
//
// seealso Logoff
type LoginAction struct {
	// Username to login with as specified in manager.conf.
	//
	// reqtured
	Username string
	// Secret to login with as specified in manager.conf.
	Secret string
}

// Generate Challenge for MD5 Auth.
// Generate a challenge for MD5 authentication.
type ChallengeAction struct {
	// Digest algorithm to use in the challenge. Valid values are:
	//
	// reqtured
	// ENUM: MD5,
	AuthType string
}

// Hangup channel.
// Hangup a channel.
type HangupAction struct {
	// The exact channel name to be hungup, or to use a regular expression, set this parameter to: /regex/
	// Example exact channel: SIP/provider-0000012a
	// Example regular expression: /^SIP/provider-.*$/
	//
	// reqtured
	Channel string
	// Numeric hangup cause.
	Cause string
}

// List channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
//
// response StatusComplete
// response-list Status
type StatusAction struct {
	// The name of the channel to query for status.
	Channel string
	// Comma `,` separated list of variable to include.
	Variables string
	// If set to "true", the Status event will include all channel variables for the requested channel(s).
	//
	// ENUM: true,false,
	AllVariables string
}

// Sets a channel variable or function value.
// This command can be used to set the value of channel variables or dialplan functions.
// NOTE:  If a channel name is not provided then the variable is considered global.
//
// seealso Getvar
type SetvarAction struct {
	// Channel to set variable for.
	Channel string
	// Variable name, function or expression.
	//
	// reqtured
	Variable string
	// Variable or function value.
	//
	// reqtured
	Value string
}

// Gets a channel variable or function value.
// Get the value of a channel variable or function return.
// NOTE:  If a channel name is not provided then the variable is considered global.
//
// seealso Setvar
type GetvarAction struct {
	// Channel to read variable from.
	Channel string
	// Variable name, function or expression.
	//
	// reqtured
	Variable string
}

// Retrieve configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only. In the case where a category name is non-unique, a filter may be specified to match only categories with matching variable values.
//
// seealso GetConfigJSON
// seealso UpdateConfig
// seealso CreateConfig
// seealso ListCategories
type GetConfigAction struct {
	// Configuration filename (e.g. foo.conf).
	//
	// reqtured
	Filename string
	// Category in configuration file.
	Category string
	// A comma separated list of name_regex=value_regex expressions which will cause only categories whose variables match all expressions to be considered. The special variable name `TEMPLATES` can be used to control whether templates are included. Passing `include` as the value will include templates along with normal categories. Passing `restrict` as the value will restrict the operation to ONLY templates. Not specifying a `TEMPLATES` expression results in the default behavior which is to not include templates.
	Filter string
}

// Retrieve configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format or optionally by specified category only. This only makes sense to be used using rawman over the HTTP interface. In the case where a category name is non-unique, a filter may be specified to match only categories with matching variable values.
//
// seealso GetConfig
// seealso UpdateConfig
// seealso CreateConfig
// seealso ListCategories
type GetConfigJSONAction struct {
	// Configuration filename (e.g. foo.conf).
	//
	// reqtured
	Filename string
	// Category in configuration file.
	Category string
	// A comma separated list of name_regex=value_regex expressions which will cause only categories whose variables match all expressions to be considered. The special variable name `TEMPLATES` can be used to control whether templates are included. Passing `include` as the value will include templates along with normal categories. Passing `restrict` as the value will restrict the operation to ONLY templates. Not specifying a `TEMPLATES` expression results in the default behavior which is to not include templates.
	Filter string
}

// Update basic configuration.
// This action will modify, create, or delete configuration elements in Asterisk configuration files.
//
// seealso GetConfig
// seealso GetConfigJSON
// seealso CreateConfig
// seealso ListCategories
type UpdateConfigAction struct {
	// Configuration filename to read (e.g. foo.conf).
	//
	// reqtured
	SrcFilename string
	// Configuration filename to write (e.g. foo.conf)
	//
	// reqtured
	DstFilename string
	// Whether or not a reload should take place (or name of specific module).
	Reload string
	// Whether the effective category contents should be preserved on template change. Default is true (pre 13.2 behavior).
	PreserveEffectiveContext string
	// Action to take.
	// 0's represent 6 digit number beginning with 000000.
	//
	// ENUM: NewCat,RenameCat,DelCat,EmptyCat,Update,Delete,Append,Insert,
	Action_000000 string `astgo:"Action-000000"`
	// Category to operate on.
	// 0's represent 6 digit number beginning with 000000.
	Cat_000000 string `astgo:"Cat-000000"`
	// Variable to work on.
	// 0's represent 6 digit number beginning with 000000.
	Var_000000 string `astgo:"Var-000000"`
	// Value to work on.
	// 0's represent 6 digit number beginning with 000000.
	Value_000000 string `astgo:"Value-000000"`
	// Extra match required to match line.
	// 0's represent 6 digit number beginning with 000000.
	Match_000000 string `astgo:"Match-000000"`
	// Line in category to operate on (used with delete and insert actions).
	// 0's represent 6 digit number beginning with 000000.
	Line_000000 string `astgo:"Line-000000"`
	// A comma separated list of action-specific options.
	//
	// The following actions share the same options...
	// 0's represent 6 digit number beginning with 000000.
	//
	// ENUM: NewCat,
	// ENUM: RenameCat,DelCat,EmptyCat,Update,Delete,Append,Insert,
	Options_000000 string `astgo:"Options-000000"`
}

// Creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory. This action is intended to be used before an UpdateConfig action.
//
// seealso GetConfig
// seealso GetConfigJSON
// seealso UpdateConfig
// seealso ListCategories
type CreateConfigAction struct {
	// The configuration filename to create (e.g. foo.conf).
	//
	// reqtured
	Filename string
}

// List categories in configuration file.
// This action will dump the categories in a given file.
//
// seealso GetConfig
// seealso GetConfigJSON
// seealso UpdateConfig
// seealso CreateConfig
type ListCategoriesAction struct {
	// Configuration filename (e.g. foo.conf).
	//
	// reqtured
	Filename string
}

// Redirect (transfer) a call.
// Redirect (transfer) a call.
//
// seealso BlindTransfer
type RedirectAction struct {
	// Channel to redirect.
	//
	// reqtured
	Channel string
	// Second call leg to transfer (optional).
	ExtraChannel string
	// Extension to transfer to.
	//
	// reqtured
	Exten string
	// Extension to transfer extrachannel to (optional).
	ExtraExten string
	// Context to transfer to.
	//
	// reqtured
	Context string
	// Context to transfer extrachannel to (optional).
	ExtraContext string
	// Priority to transfer to.
	//
	// reqtured
	Priority string
	// Priority to transfer extrachannel to (optional).
	ExtraPriority string
}

// Attended transfer.
// Attended transfer.
//
// seealso AttendedTransfer
type AtxferAction struct {
	// Transferer's channel.
	//
	// reqtured
	Channel string
	// Extension to transfer to.
	//
	// reqtured
	Exten string
	// Context to transfer to.
	Context string
}

// Originate a call.
// Generates an outgoing call to a Extension/Context/Priority or Application/Data
//
// seealso OriginateResponse
type OriginateAction struct {
	// Channel name to call.
	//
	// reqtured
	Channel string
	// Extension to use (requires `Context` and `Priority`)
	Exten string
	// Context to use (requires `Exten` and `Priority`)
	Context string
	// Priority to use (requires `Exten` and `Context`)
	Priority string
	// Application to execute.
	Application string
	// Data to use (requires `Application`).
	Data string
	// How long to wait for call to be answered (in ms.).
	Timeout string
	// Caller ID to be set on the outgoing channel.
	CallerID string
	// Channel variable to set, multiple Variable: headers are allowed.
	Variable string
	// Account code.
	Account string
	// Set to `true` to force call bridge on early media..
	EarlyMedia string
	// Set to `true` for fast origination.
	Async string
	// Comma-separated list of codecs to use for this call.
	Codecs string
	// Channel UniqueId to be set on the channel.
	ChannelId string
	// Channel UniqueId to be set on the second local channel.
	OtherChannelId string
}

// Execute Asterisk CLI Command.
// Run a CLI command.
type CommandAction struct {
	// Asterisk CLI command to run.
	//
	// reqtured
	Command string
}

// Check Extension Status.
// Report the extension state for given extension. If the extension has a hint, will use devicestate to check the status of the device connected to the extension.
// Will return an `Extension Status` message. The response will include the hint for the extension and the status.
//
// seealso ExtensionStatus
type ExtensionStateAction struct {
	// Extension to check state on.
	//
	// reqtured
	Exten string
	// Context for extension.
	//
	// reqtured
	Context string
}

// Check Presence State
// Report the presence state for the given presence provider.
// Will return a `Presence State` message. The response will include the presence state and, if set, a presence subtype and custom message.
//
// seealso PresenceStatus
type PresenceStateAction struct {
	// Presence Provider to check the state of
	//
	// reqtured
	Provider string
}

// Set absolute timeout.
// Hangup a channel after a certain time. Acknowledges set time with `Timeout Set` message.
type AbsoluteTimeoutAction struct {
	// Channel name to hangup.
	//
	// reqtured
	Channel string
	// Maximum duration of the call (sec).
	//
	// reqtured
	Timeout string
}

// Check mailbox.
// Checks a voicemail account for status.
// Returns whether there are messages waiting.
// Message: Mailbox Status.
// Mailbox: mailboxid.
// Waiting: `0` if messages waiting, `1` if no messages waiting.
//
// seealso MailboxCount
type MailboxStatusAction struct {
	// Full mailbox ID mailbox@vm-context.
	//
	// reqtured
	Mailbox string
}

// Check Mailbox Message Count.
// Checks a voicemail account for new messages.
// Returns number of urgent, new and old messages.
// Message: Mailbox Message Count
// Mailbox: mailboxid
// UrgentMessages: count
// NewMessages: count
// OldMessages: count
//
// seealso MailboxStatus
type MailboxCountAction struct {
	// Full mailbox ID mailbox@vm-context.
	//
	// reqtured
	Mailbox string
}

// List available manager commands.
// Returns the action name and synopsis for every action that is available to the user.
type ListCommandsAction struct {
}

// Send text message to channel.
// Sends A Text Message to a channel while in a call.
type SendTextAction struct {
	// Channel to send message to.
	//
	// reqtured
	Channel string
	// Message to send.
	//
	// reqtured
	Message string
}

// Send an arbitrary event.
// Send an event to manager sessions.
//
// seealso UserEvent
// seealso UserEvent
type UserEventAction struct {
	// Event string to send.
	//
	// reqtured
	UserEvent string
	// Content1.
	Header1 string
	// ContentN.
	HeaderN string
}

// Wait for an event to occur.
// This action will ellicit a `Success` response. Whenever a manager event is queued. Once WaitEvent has been called on an HTTP manager session, events will be generated and queued.
type WaitEventAction struct {
	// Maximum time (in seconds) to wait for events, `-1` means forever.
	//
	// reqtured
	Timeout string
}

// Show PBX core settings (version etc).
// Query for Core PBX settings.
type CoreSettingsAction struct {
}

// Show PBX core status variables.
// Query for Core PBX status.
type CoreStatusAction struct {
}

// Send a reload event.
// Send a reload event.
//
// seealso ModuleLoad
type Reload_1Action struct {
	// Name of the module to reload.
	Module string
}

func (*Reload_1Action) OriginalName() string { return "Reload" }

// List currently active channels.
// List currently defined channels and some information about them.
//
// response CoreShowChannelsComplete
// response-list CoreShowChannel
type CoreShowChannelsAction struct {
}

// Reload and rotate the Asterisk logger.
// Reload and rotate the logger. Analogous to the CLI command 'logger rotate'.
type LoggerRotateAction struct {
}

// Module management.
// Loads, unloads or reloads an Asterisk module in a running system.
//
// seealso Reload
// seealso ModuleCheck
type ModuleLoadAction struct {
	// Asterisk module name (including .so extension) or subsystem identifier:
	//
	// ENUM: cdr,dnsmgr,extconfig,enum,acl,manager,http,logger,features,dsp,udptl,indications,cel,plc,
	Module string
	// The operation to be done on module. Subsystem identifiers may only be reloaded.
	// If no module is specified for a `reload` loadtype, all modules are reloaded.
	//
	// reqtured
	// ENUM: load,unload,reload,
	LoadType string
}

// Check if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
//
// seealso ModuleLoad
type ModuleCheckAction struct {
	// Asterisk module name (not including extension).
	//
	// reqtured
	Module string
}

// Generate an Advice of Charge message on a channel.
// Generates an AOC-D or AOC-E message on a channel.
//
// seealso AOC-D
// seealso AOC-E
type AOCMessageAction struct {
	// Channel name to generate the AOC message on.
	//
	// reqtured
	Channel string
	// Partial channel prefix. By using this option one can match the beginning part of a channel name without having to put the entire name in. For example if a channel name is SIP/snom-00000001 and this value is set to SIP/snom, then that channel matches and the message will be sent. Note however that only the first matched channel has the message sent on it.
	ChannelPrefix string
	// Defines what type of AOC message to create, AOC-D or AOC-E
	//
	// reqtured
	// ENUM: D,E,
	MsgType string
	// Defines what kind of charge this message represents.
	//
	// reqtured
	// ENUM: NA,FREE,Currency,Unit,
	ChargeType string
	// This represents the amount of units charged. The ETSI AOC standard specifies that this value along with the optional UnitType value are entries in a list. To accommodate this these values take an index value starting at 0 which can be used to generate this list of unit entries. For Example, If two unit entires were required this could be achieved by setting the paramter UnitAmount(0)=1234 and UnitAmount(1)=5678. Note that UnitAmount at index 0 is required when ChargeType=Unit, all other entries in the list are optional.
	UnitAmount_0_ string `astgo:"UnitAmount(0)"`
	// Defines the type of unit. ETSI AOC standard specifies this as an integer value between 1 and 16, but this value is left open to accept any positive integer. Like the UnitAmount parameter, this value represents a list entry and has an index parameter that starts at 0.
	UnitType_0_ string `astgo:"UnitType(0)"`
	// Specifies the currency's name. Note that this value is truncated after 10 characters.
	CurrencyName string
	// Specifies the charge unit amount as a positive integer. This value is required when ChargeType==Currency.
	CurrencyAmount string
	// Specifies the currency multiplier. This value is required when ChargeType==Currency.
	//
	// ENUM: OneThousandth,OneHundredth,OneTenth,One,Ten,Hundred,Thousand,
	CurrencyMultiplier string
	// Defines what kind of AOC-D total is represented.
	//
	// ENUM: Total,SubTotal,
	TotalType string
	// Represents a billing ID associated with an AOC-D or AOC-E message. Note that only the first 3 items of the enum are valid AOC-D billing IDs
	//
	// ENUM: Normal,ReverseCharge,CreditCard,CallFwdUnconditional,CallFwdBusy,CallFwdNoReply,CallDeflection,CallTransfer,
	AOCBillingId string
	// Charging association identifier. This is optional for AOC-E and can be set to any value between -32768 and 32767
	ChargingAssociationId string
	// Represents the charging association party number. This value is optional for AOC-E.
	ChargingAssociationNumber string
	// Integer representing the charging plan associated with the ChargingAssociationNumber. The value is bits 7 through 1 of the Q.931 octet containing the type-of-number and numbering-plan-identification fields.
	ChargingAssociationPlan string
}

// Dynamically add filters for the current manager session.
// The filters added are only used for the current session. Once the connection is closed the filters are removed.
// This comand requires the system permission because this command can be used to create filters that may bypass filters defined in manager.conf
//
// seealso FilterList
type FilterAction struct {
	//
	// ENUM: Add,
	Operation string
	// Filters can be whitelist or blacklist
	// Example whitelist filter: "Event: Newchannel"
	// Example blacklist filter: "!Channel: DAHDI.*"
	// This filter option is used to whitelist or blacklist events per user to be reported with regular expressions and are allowed if both the regex matches and the user has read access as defined in manager.conf. Filters are assumed to be for whitelisting unless preceeded by an exclamation point, which marks it as being black. Evaluation of the filters is as follows:
	// - If no filters are configured all events are reported as normal.
	// - If there are white filters only: implied black all filter processed first, then white filters.
	// - If there are black filters only: implied white all filter processed first, then black filters.
	// - If there are both white and black filters: implied black all filter processed first, then white filters, and lastly black filters.
	Filter string
}

// Show current event filters for this session
// The filters displayed are for the current session. Only those filters defined in manager.conf will be present upon starting a new session.
//
// seealso Filter
type FilterListAction struct {
}

// Blind transfer channel(s) to the given destination
// Redirect all channels currently bridged to the specified channel to the specified destination.
//
// seealso Redirect
// seealso BlindTransfer
type BlindTransferAction struct {
	//
	// reqtured
	Channel string
	Context string
	Exten   string
}

// Get a list of bridges in the system.
// Returns a list of bridges, optionally filtering on a bridge type.
//
// seealso Bridge
// seealso BridgeDestroy
// seealso BridgeInfo
// seealso BridgeKick
type BridgeListAction struct {
	// Optional type for filtering the resulting list of bridges.
	BridgeType string
}

// Get information about a bridge.
// Returns detailed information about a bridge and the channels in it.
//
// seealso Bridge
// seealso BridgeDestroy
// seealso BridgeKick
// seealso BridgeList
// response BridgeInfoComplete
// response-list BridgeInfoChannel
type BridgeInfoAction struct {
	// The unique ID of the bridge about which to retrieve information.
	//
	// reqtured
	BridgeUniqueid string
}

// Destroy a bridge.
// Deletes the bridge, causing channels to continue or hang up.
//
// seealso Bridge
// seealso BridgeInfo
// seealso BridgeKick
// seealso BridgeList
// seealso BridgeDestroy
type BridgeDestroy_1Action struct {
	// The unique ID of the bridge to destroy.
	//
	// reqtured
	BridgeUniqueid string
}

func (*BridgeDestroy_1Action) OriginalName() string { return "BridgeDestroy" }

// Kick a channel from a bridge.
// The channel is removed from the bridge.
//
// seealso Bridge
// seealso BridgeDestroy
// seealso BridgeInfo
// seealso BridgeList
// seealso BridgeLeave
type BridgeKickAction struct {
	// The unique ID of the bridge containing the channel to destroy. This parameter can be omitted, or supplied to insure that the channel is not removed from the wrong bridge.
	BridgeUniqueid string
	// The channel to kick out of a bridge.
	//
	// reqtured
	Channel string
}

// Send an out of call message to an endpoint.
type MessageSendAction struct {
	// The URI the message is to be sent to.
	// INFO: Specifying a prefix of `sip:` will send the message as a SIP MESSAGE request.
	// INFO: Specifying a prefix of `pjsip:` will send the message as a SIP MESSAGE request.
	// INFO: Specifying a prefix of `xmpp:` will send the message as an XMPP chat message.
	//
	// reqtured
	To string
	// A From URI for the message if needed for the message technology being used to send this message.
	// INFO: The `from` parameter can be a configured peer name or in the form of "display-name" &lt;URI&gt;.
	// INFO: The `from` parameter can be a configured endpoint or in the form of "display-name" &lt;URI&gt;.
	// INFO: Specifying a prefix of `xmpp:` will specify the account defined in `xmpp.conf` to send the message from. Note that this field is required for XMPP messages.
	From string
	// The message body text. This must not contain any newlines as that conflicts with the AMI protocol.
	Body string
	// Text bodies requiring the use of newlines have to be base64 encoded in this field. Base64Body will be decoded before being sent out. Base64Body takes precedence over Body.
	Base64Body string
	// Message variable to set, multiple Variable: headers are allowed. The header value is a comma separated list of name=value pairs.
	Variable string
}

// Show dialplan contexts and extensions
// Show dialplan contexts and extensions. Be aware that showing the full dialplan may take a lot of capacity.
type ShowDialPlanAction struct {
	// Show a specific extension.
	Extension string
	// Show a specific context.
	Context string
}

// List the current known extension states.
// This will list out all known extension states in a sequence of ExtensionStatus events. When finished, a ExtensionStateListComplete event will be emitted.
//
// seealso ExtensionState
// seealso HINT
// seealso EXTENSION_STATE
// response ExtensionStateListComplete
// response-list ExtensionStatus
type ExtensionStateListAction struct {
}

// Get a list of parking lots
// List all parking lots as a series of AMI events
type ParkinglotsAction struct {
}

// List parked calls.
// List parked calls.
type ParkedCallsAction struct {
	// If specified, only show parked calls from the parking lot with this name.
	ParkingLot string
}

// Park a channel.
// Park an arbitrary channel with optional arguments for specifying the parking lot used, how long the channel should remain parked, and what dial string to use as the parker if the call times out.
type ParkAction struct {
	// Channel name to park.
	//
	// reqtured
	Channel string
	// Channel name to use when constructing the dial string that will be dialed if the parked channel times out. If `TimeoutChannel` is in a two party bridge with `Channel`, then `TimeoutChannel` will receive an announcement and be treated as having parked `Channel` in the same manner as the Park Call DTMF feature.
	TimeoutChannel string
	// If specified, then this channel will receive an announcement when `Channel` is parked if `AnnounceChannel` is in a state where it can receive announcements (AnnounceChannel must be bridged). `AnnounceChannel` has no bearing on the actual state of the parked call.
	AnnounceChannel string
	// Overrides the timeout of the parking lot for this park action. Specified in milliseconds, but will be converted to seconds. Use a value of 0 to disable the timeout.
	Timeout string
	// The parking lot to use when parking the channel
	Parkinglot string
}

// Add an AGI command to execute by Async AGI.
// Add an AGI command to the execute queue of the channel in Async AGI.
//
// seealso AsyncAGIStart
// seealso AsyncAGIExec
// seealso AsyncAGIEnd
type AGIAction struct {
	// Channel that is currently in Async AGI.
	//
	// reqtured
	Channel string
	// Application to execute.
	//
	// reqtured
	Command string
	// This will be sent back in CommandID header of AsyncAGI exec event notification.
	CommandID string
}

// Lists active FAX sessions
// Will generate a series of FAXSession events with information about each FAXSession. Closes with a FAXSessionsComplete event which includes a count of the included FAX sessions. This action works in the same manner as the CLI command 'fax show sessions'
type FAXSessionsAction struct {
}

// Responds with a detailed description of a single FAX session
// Provides details about a specific FAX session. The response will include a common subset of the output from the CLI command 'fax show session &lt;session_number&gt;' for each technology. If the FAX technolgy used by this session does not include a handler for FAXSession, then this action will fail.
type FAXSessionAction struct {
	// The session ID of the fax the user is interested in.
	//
	// reqtured
	SessionNumber string
}

// Responds with fax statistics
// Provides FAX statistics including the number of active sessions, reserved sessions, completed sessions, failed sessions, and the number of receive/transmit attempts. This command provides all of the non-technology specific information provided by the CLI command 'fax show stats'
type FAXStatsAction struct {
}

// List the current known device states.
// This will list out all known device states in a sequence of DeviceStateChange events. When finished, a DeviceStateListComplete event will be emitted.
//
// seealso DeviceStateChange
// seealso DEVICE_STATE
// response DeviceStateListComplete
// response-list DeviceStateChange
type DeviceStateListAction struct {
}

// List the current known presence states.
// This will list out all known presence states in a sequence of PresenceStateChange events. When finished, a PresenceStateListComplete event will be emitted.
//
// seealso PresenceState
// seealso PresenceStatus
// seealso PRESENCE_STATE
// response PresenceStateListComplete
// response-list PresenceStateChange
type PresenceStateListAction struct {
}

// Monitor a channel.
// This action may be used to record the audio on a specified channel.
type MonitorAction struct {
	// Used to specify the channel to record.
	//
	// reqtured
	Channel string
	// Is the name of the file created in the monitor spool directory. Defaults to the same name as the channel (with slashes replaced with dashes).
	File string
	// Is the audio recording format. Defaults to `wav`.
	Format string
	// Boolean parameter as to whether to mix the input and output channels together after the recording is finished.
	Mix string
}

// Stop monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
type StopMonitorAction struct {
	// The name of the channel monitored.
	//
	// reqtured
	Channel string
}

// Change monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
type ChangeMonitorAction struct {
	// Used to specify the channel to record.
	//
	// reqtured
	Channel string
	// Is the new name of the file created in the monitor spool directory.
	//
	// reqtured
	File string
}

// Pause monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
type PauseMonitorAction struct {
	// Used to specify the channel to record.
	//
	// reqtured
	Channel string
}

// Unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
type UnpauseMonitorAction struct {
	// Used to specify the channel to record.
	//
	// reqtured
	Channel string
}

// Mute an audio stream.
// Mute an incoming or outgoing audio stream on a channel.
type MuteAudioAction struct {
	// The channel you want to mute.
	//
	// reqtured
	Channel string
	//
	// reqtured
	// ENUM: in,out,all,
	Direction string
	//
	// reqtured
	// ENUM: on,off,
	State string
}

// Get selected mailboxes with message counts.
// Get a list of mailboxes with their message counts.
type MWIGetAction struct {
	// Mailbox ID in the form of /regex/ for all mailboxes matching the regular expression. Otherwise it is for a specific mailbox.
	//
	// reqtured
	Mailbox string
}

// Delete selected mailboxes.
// Delete the specified mailboxes.
type MWIDeleteAction struct {
	// Mailbox ID in the form of /regex/ for all mailboxes matching the regular expression. Otherwise it is for a specific mailbox.
	//
	// reqtured
	Mailbox string
}

// Update the mailbox message counts.
// Update the mailbox message counts.
type MWIUpdateAction struct {
	// Specific mailbox ID.
	//
	// reqtured
	Mailbox string
	// The number of old messages in the mailbox. Defaults to zero if missing.
	OldMessages string
	// The number of new messages in the mailbox. Defaults to zero if missing.
	NewMessages string
}

// Qualify a chan_pjsip endpoint.
// Qualify a chan_pjsip endpoint.
type PJSIPQualifyAction struct {
	// The endpoint you want to qualify.
	//
	// reqtured
	Endpoint string
}

// Lists PJSIP endpoints.
// Provides a listing of all endpoints. For each endpoint an `EndpointList` event is raised that contains relevant attributes and status information. Once all endpoints have been listed an `EndpointListComplete` event is issued.
//
// response EndpointListComplete
// response-list EndpointList
type PJSIPShowEndpointsAction struct {
}

// Detail listing of an endpoint and its objects.
// Provides a detailed listing of options for a given endpoint. Events are issued showing the configuration and status of the endpoint and associated objects. These events include `EndpointDetail`, `AorDetail`, `AuthDetail`, `TransportDetail`, and `IdentifyDetail`. Some events may be listed multiple times if multiple objects are associated (for instance AoRs). Once all detail events have been raised a final `EndpointDetailComplete` event is issued.
//
// response EndpointDetailComplete
// response-list EndpointDetail,IdentifyDetail,ContactStatusDetail,AuthDetail,TransportDetail,AorDetail
type PJSIPShowEndpointAction struct {
	// The endpoint to list.
	//
	// reqtured
	Endpoint string
}

// Send a NOTIFY to either an endpoint or an arbitrary URI.
// Sends a NOTIFY to an endpoint or an arbitrary URI.
// All parameters for this event must be specified in the body of this request via multiple `Variable: name=value` sequences.
// NOTE: One (and only one) of `Endpoint` or `URI` must be specified. If `URI` is used, the default outbound endpoint will be used to send the message. If the default outbound endpoint isn't configured, this command can not send to an arbitrary URI.
type PJSIPNotifyAction struct {
	// The endpoint to which to send the NOTIFY.
	Endpoint string
	// Abritrary URI to which to send the NOTIFY.
	URI string
	// Appends variables as headers/content to the NOTIFY. If the variable is named `Content`, then the value will compose the body of the message if another variable sets `Content-Type`. name=value
	//
	// reqtured
	Variable string
}

// Unregister an outbound registration.
// Unregisters the specified (or all) outbound registration(s) and stops future registration attempts. Call PJSIPRegister to start registration and schedule re-registrations according to configuration.
type PJSIPUnregisterAction struct {
	// The outbound registration to unregister or '*all' to unregister them all.
	//
	// reqtured
	Registration string
}

// Register an outbound registration.
// Unregisters the specified (or all) outbound registration(s) then starts registration and schedules re-registrations according to configuration.
type PJSIPRegisterAction struct {
	// The outbound registration to register or '*all' to register them all.
	//
	// reqtured
	Registration string
}

// Lists PJSIP outbound registrations.
// In response `OutboundRegistrationDetail` events showing configuration and status information are raised for each outbound registration object. `AuthDetail` events are raised for each associated auth object as well. Once all events are completed an `OutboundRegistrationDetailComplete` is issued.
type PJSIPShowRegistrationsOutboundAction struct {
}

// Lists subscriptions.
// Provides a listing of all inbound subscriptions. An event `InboundSubscriptionDetail` is issued for each subscription object. Once all detail events are completed an `InboundSubscriptionDetailComplete` event is issued.
type PJSIPShowSubscriptionsInboundAction struct {
}

// Lists subscriptions.
// Provides a listing of all outbound subscriptions. An event `OutboundSubscriptionDetail` is issued for each subscription object. Once all detail events are completed an `OutboundSubscriptionDetailComplete` event is issued.
type PJSIPShowSubscriptionsOutboundAction struct {
}

// Displays settings for configured resource lists.
// Provides a listing of all resource lists. An event `ResourceListDetail` is issued for each resource list object. Once all detail events are completed a `ResourceListDetailComplete` event is issued.
type PJSIPShowResourceListsAction struct {
}

// Lists PJSIP inbound registrations.
// In response, `InboundRegistrationDetail` events showing configuration and status information are raised for all contacts, static or dynamic. Once all events are completed an `InboundRegistrationDetailComplete` is issued.
// WARNING:  This command just dumps all coonfigured AORs with contacts, even if the contact is a permanent one. To really get just inbound registrations, use `PJSIPShowRegistrationInboundContactStatuses`.
//
// seealso PJSIPShowRegistrationInboundContactStatuses
type PJSIPShowRegistrationsInboundAction struct {
}

// Lists ContactStatuses for PJSIP inbound registrations.
// In response, `ContactStatusDetail` events showing status information are raised for each inbound registration (dynamic contact) object. Once all events are completed a `ContactStatusDetailComplete` event is issued.
type PJSIPShowRegistrationInboundContactStatusesAction struct {
}

// Expire (remove) an object from a sorcery memory cache.
// Expires (removes) an object from a sorcery memory cache.
type SorceryMemoryCacheExpireObjectAction struct {
	// The name of the cache to expire the object from.
	//
	// reqtured
	Cache string
	// The name of the object to expire.
	//
	// reqtured
	Object string
}

// Expire (remove) ALL objects from a sorcery memory cache.
// Expires (removes) ALL objects from a sorcery memory cache.
type SorceryMemoryCacheExpireAction struct {
	// The name of the cache to expire all objects from.
	//
	// reqtured
	Cache string
}

// Mark an object in a sorcery memory cache as stale.
// Marks an object as stale within a sorcery memory cache.
type SorceryMemoryCacheStaleObjectAction struct {
	// The name of the cache to mark the object as stale in.
	//
	// reqtured
	Cache string
	// The name of the object to mark as stale.
	//
	// reqtured
	Object string
	// If true, then immediately reload the object from the backend cache instead of waiting for the next retrieval
	Reload string
}

// Marks ALL objects in a sorcery memory cache as stale.
// Marks ALL objects in a sorcery memory cache as stale.
type SorceryMemoryCacheStaleAction struct {
	// The name of the cache to mark all object as stale in.
	//
	// reqtured
	Cache string
}

// Expire all objects from a memory cache and populate it with all objects from the backend.
// Expires all objects from a memory cache and populate it with all objects from the backend.
type SorceryMemoryCachePopulateAction struct {
	// The name of the cache to populate.
	//
	// reqtured
	Cache string
}

// Sends a message to a Jabber Client.
// Sends a message to a Jabber Client.
type JabberSendAction struct {
	// Client or transport Asterisk uses to connect to JABBER.
	//
	// reqtured
	Jabber string
	// XMPP/Jabber JID (Name) of recipient.
	//
	// reqtured
	JID string
	// Message to be sent to the buddy.
	//
	// reqtured
	Message string
}

package amimodels

// SKINNYdevicesAction List SKINNY devices (text format).
type SKINNYdevicesAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (SKINNYdevicesAction) ActionTypeName() string {
	return "SKINNYdevices"
}
func (a *SKINNYdevicesAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYdevicesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SKINNYshowdeviceAction Show SKINNY device (text format).
type SKINNYshowdeviceAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The device name you want to check.
	Device string
}

func (SKINNYshowdeviceAction) ActionTypeName() string {
	return "SKINNYshowdevice"
}
func (a *SKINNYshowdeviceAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYshowdeviceAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SKINNYlinesAction List SKINNY lines (text format).
type SKINNYlinesAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (SKINNYlinesAction) ActionTypeName() string {
	return "SKINNYlines"
}
func (a *SKINNYlinesAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYlinesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SKINNYshowlineAction Show SKINNY line (text format).
type SKINNYshowlineAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The line name you want to check.
	Line string
}

func (SKINNYshowlineAction) ActionTypeName() string {
	return "SKINNYshowline"
}
func (a *SKINNYshowlineAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYshowlineAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// IAXpeersAction List IAX peers.
type IAXpeersAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (IAXpeersAction) ActionTypeName() string {
	return "IAXpeers"
}
func (a *IAXpeersAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXpeersAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// IAXpeerlistAction List IAX Peers.
type IAXpeerlistAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (IAXpeerlistAction) ActionTypeName() string {
	return "IAXpeerlist"
}
func (a *IAXpeerlistAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXpeerlistAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// IAXnetstatsAction Show IAX Netstats.
type IAXnetstatsAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (IAXnetstatsAction) ActionTypeName() string {
	return "IAXnetstats"
}
func (a *IAXnetstatsAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXnetstatsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// IAXregistryAction Show IAX registrations.
type IAXregistryAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (IAXregistryAction) ActionTypeName() string {
	return "IAXregistry"
}
func (a *IAXregistryAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXregistryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDITransferAction Transfer DAHDI Channel.
type DAHDITransferAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// DAHDI channel number to transfer.
	DAHDIChannel string
}

func (DAHDITransferAction) ActionTypeName() string {
	return "DAHDITransfer"
}
func (a *DAHDITransferAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDITransferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDIHangupAction Hangup DAHDI Channel.
type DAHDIHangupAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// DAHDI channel number to hangup.
	DAHDIChannel string
}

func (DAHDIHangupAction) ActionTypeName() string {
	return "DAHDIHangup"
}
func (a *DAHDIHangupAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIHangupAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDIDialOffhookAction Dial over DAHDI channel while offhook.
type DAHDIDialOffhookAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// DAHDI channel number to dial digits.
	DAHDIChannel string

	// Digits to dial.
	Number string
}

func (DAHDIDialOffhookAction) ActionTypeName() string {
	return "DAHDIDialOffhook"
}
func (a *DAHDIDialOffhookAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDialOffhookAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDIDNDonAction Toggle DAHDI channel Do Not Disturb status ON.
type DAHDIDNDonAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// DAHDI channel number to set DND on.
	DAHDIChannel string
}

func (DAHDIDNDonAction) ActionTypeName() string {
	return "DAHDIDNDon"
}
func (a *DAHDIDNDonAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDNDonAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDIDNDoffAction Toggle DAHDI channel Do Not Disturb status OFF.
type DAHDIDNDoffAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// DAHDI channel number to set DND off.
	DAHDIChannel string
}

func (DAHDIDNDoffAction) ActionTypeName() string {
	return "DAHDIDNDoff"
}
func (a *DAHDIDNDoffAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDNDoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDIShowChannelsAction Show status of DAHDI channels.
type DAHDIShowChannelsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Specify the specific channel number to show. Show all channels if zero or not present.
	DAHDIChannel string
}

func (DAHDIShowChannelsAction) ActionTypeName() string {
	return "DAHDIShowChannels"
}
func (a *DAHDIShowChannelsAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIShowChannelsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DAHDIRestartAction Fully Restart DAHDI channels (terminates calls).
type DAHDIRestartAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (DAHDIRestartAction) ActionTypeName() string {
	return "DAHDIRestart"
}
func (a *DAHDIRestartAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIRestartAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PRIShowSpansAction Show status of PRI spans.
type PRIShowSpansAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Specify the specific span to show. Show all spans if zero or not present.
	Span string
}

func (PRIShowSpansAction) ActionTypeName() string {
	return "PRIShowSpans"
}
func (a *PRIShowSpansAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIShowSpansAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PRIDebugSetAction Set PRI debug levels for a span
type PRIDebugSetAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Which span to affect.
	Span string

	// What debug level to set. May be a numerical value or a text value from the list below
	Level string
}

func (PRIDebugSetAction) ActionTypeName() string {
	return "PRIDebugSet"
}
func (a *PRIDebugSetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugSetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PRIDebugFileSetAction Set the file used for PRI debug message output
type PRIDebugFileSetAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Path of file to write debug output.
	File string
}

func (PRIDebugFileSetAction) ActionTypeName() string {
	return "PRIDebugFileSet"
}
func (a *PRIDebugFileSetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugFileSetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PRIDebugFileUnsetAction Disables file output for PRI debug messages
type PRIDebugFileUnsetAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (PRIDebugFileUnsetAction) ActionTypeName() string {
	return "PRIDebugFileUnset"
}
func (a *PRIDebugFileUnsetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugFileUnsetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SIPpeersAction List SIP peers (text format).
type SIPpeersAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (SIPpeersAction) ActionTypeName() string {
	return "SIPpeers"
}
func (a *SIPpeersAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPpeersAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SIPshowpeerAction show SIP peer (text format).
type SIPshowpeerAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The peer name you want to check.
	Peer string
}

func (SIPshowpeerAction) ActionTypeName() string {
	return "SIPshowpeer"
}
func (a *SIPshowpeerAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPshowpeerAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SIPqualifypeerAction Qualify SIP peers.
type SIPqualifypeerAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The peer name you want to qualify.
	Peer string
}

func (SIPqualifypeerAction) ActionTypeName() string {
	return "SIPqualifypeer"
}
func (a *SIPqualifypeerAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPqualifypeerAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SIPshowregistryAction Show SIP registrations (text format).
type SIPshowregistryAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (SIPshowregistryAction) ActionTypeName() string {
	return "SIPshowregistry"
}
func (a *SIPshowregistryAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPshowregistryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SIPnotifyAction Send a SIP notify.
type SIPnotifyAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Peer to receive the notify.
	Channel string

	// At least one variable pair must be specified.
	//             =
	Variable string

	// When specified, SIP notity will be sent as a part of an existing dialog.
	CallID string
}

func (SIPnotifyAction) ActionTypeName() string {
	return "SIPnotify"
}
func (a *SIPnotifyAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPnotifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SIPpeerstatusAction Show the status of one or all of the sip peers.
type SIPpeerstatusAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The peer name you want to check.
	Peer string
}

func (SIPpeerstatusAction) ActionTypeName() string {
	return "SIPpeerstatus"
}
func (a *SIPpeerstatusAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPpeerstatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DialplanExtensionAddAction Add an extension to the dialplan
type DialplanExtensionAddAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Context where the extension will be created. The context will
	//             be created if it does not already exist.
	Context string

	// Name of the extension that will be created (may include callerid match by separating
	//             with '/')
	Extension string

	// Priority being added to this extension. Must be either  or a
	//             numerical value.
	Priority string

	// The application to use for this extension at the requested priority
	Application string

	// Arguments to the application.
	ApplicationData string

	// If set to 'yes', '1', 'true' or any of the other values we evaluate as true, then
	//             if an extension already exists at the requested context, extension, and priority it will
	//             be overwritten. Otherwise, the existing extension will remain and the action will fail.
	Replace string
}

func (DialplanExtensionAddAction) ActionTypeName() string {
	return "DialplanExtensionAdd"
}
func (a *DialplanExtensionAddAction) GetActionID() string {
	return a.ActionID
}
func (a *DialplanExtensionAddAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DialplanExtensionRemoveAction Remove an extension from the dialplan
type DialplanExtensionRemoveAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Context of the extension being removed
	Context string

	// Name of the extension being removed (may include callerid match by separating with '/')
	Extension string

	// If provided, only remove this priority from the extension instead of all
	//             priorities in the extension.
	Priority string
}

func (DialplanExtensionRemoveAction) ActionTypeName() string {
	return "DialplanExtensionRemove"
}
func (a *DialplanExtensionRemoveAction) GetActionID() string {
	return a.ActionID
}
func (a *DialplanExtensionRemoveAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// VoicemailUsersListAction List All Voicemail User Information.
type VoicemailUsersListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (VoicemailUsersListAction) ActionTypeName() string {
	return "VoicemailUsersList"
}
func (a *VoicemailUsersListAction) GetActionID() string {
	return a.ActionID
}
func (a *VoicemailUsersListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// VoicemailUserStatusAction Show the status of given voicemail user's info.
type VoicemailUserStatusAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The context you want to check.
	Context string

	// The mailbox you want to check.
	Mailbox string
}

func (VoicemailUserStatusAction) ActionTypeName() string {
	return "VoicemailUserStatus"
}
func (a *VoicemailUserStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *VoicemailUserStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// VoicemailRefreshAction Tell Asterisk to poll mailboxes for a change
type VoicemailRefreshAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Context string

	Mailbox string
}

func (VoicemailRefreshAction) ActionTypeName() string {
	return "VoicemailRefresh"
}
func (a *VoicemailRefreshAction) GetActionID() string {
	return a.ActionID
}
func (a *VoicemailRefreshAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// AgentsAction Lists agents and their status.
type AgentsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (AgentsAction) ActionTypeName() string {
	return "Agents"
}
func (a *AgentsAction) GetActionID() string {
	return a.ActionID
}
func (a *AgentsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// AgentLogoffAction Sets an agent as no longer logged in.
type AgentLogoffAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Agent ID of the agent to log off.
	Agent string

	// Set to  to not hangup existing calls.
	Soft string
}

func (AgentLogoffAction) ActionTypeName() string {
	return "AgentLogoff"
}
func (a *AgentLogoffAction) GetActionID() string {
	return a.ActionID
}
func (a *AgentLogoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ControlPlaybackAction Control the playback of a file being played to a channel.
type ControlPlaybackAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the channel that currently has a file being played back to it.
	Channel string

	Control string
}

func (ControlPlaybackAction) ActionTypeName() string {
	return "ControlPlayback"
}
func (a *ControlPlaybackAction) GetActionID() string {
	return a.ActionID
}
func (a *ControlPlaybackAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MixMonitorMuteAction Mute / unMute a Mixmonitor recording.
type MixMonitorMuteAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Used to specify the channel to mute.
	Channel string

	// Which part of the recording to mute: read, write or both (from channel, to channel or both channels).
	Direction string

	// Turn mute on or off : 1 to turn on, 0 to turn off.
	State string
}

func (MixMonitorMuteAction) ActionTypeName() string {
	return "MixMonitorMute"
}
func (a *MixMonitorMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MixMonitorMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MixMonitorAction Record a call and mix the audio during the recording. Use of StopMixMonitor is required
//       to guarantee the audio file is available for processing during dialplan execution.
type MixMonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Used to specify the channel to record.
	Channel string

	// Is the name of the file created in the monitor spool directory.
	//             Defaults to the same name as the channel (with slashes replaced with dashes).
	//             This argument is optional if you specify to record unidirectional audio with
	//             either the r(filename) or t(filename) options in the options field. If
	//             neither MIXMONITOR_FILENAME or this parameter is set, the mixed stream won't
	//             be recorded.
	File string

	// Options that apply to the MixMonitor in the same way as they
	//             would apply if invoked from the MixMonitor application. For a list of
	//             available options, see the documentation for the mixmonitor application.
	options string

	// Will be executed when the recording is over.
	//             Any strings matching  will be unescaped to .
	//             All variables will be evaluated at the time MixMonitor is called.
	Command string
}

func (MixMonitorAction) ActionTypeName() string {
	return "MixMonitor"
}
func (a *MixMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *MixMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// StopMixMonitorAction Stop recording a call through MixMonitor, and free the recording's file handle.
type StopMixMonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the channel monitored.
	Channel string

	// If a valid ID is provided, then this command will stop only that specific
	//             MixMonitor.
	MixMonitorID string
}

func (StopMixMonitorAction) ActionTypeName() string {
	return "StopMixMonitor"
}
func (a *StopMixMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *StopMixMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PlayDTMFAction Play DTMF signal on a specific channel.
type PlayDTMFAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel name to send digit to.
	Channel string

	// The DTMF digit to play.
	Digit string

	// The duration, in milliseconds, of the digit to be played.
	Duration string

	// Emulate receiving DTMF on this channel instead of sending it out.
	Receive string
}

func (PlayDTMFAction) ActionTypeName() string {
	return "PlayDTMF"
}
func (a *PlayDTMFAction) GetActionID() string {
	return a.ActionID
}
func (a *PlayDTMFAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeListAction List participants in a conference.
type ConfbridgeListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Conference number.
	Conference string
}

func (ConfbridgeListAction) ActionTypeName() string {
	return "ConfbridgeList"
}
func (a *ConfbridgeListAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeListRoomsAction List active conferences.
type ConfbridgeListRoomsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (ConfbridgeListRoomsAction) ActionTypeName() string {
	return "ConfbridgeListRooms"
}
func (a *ConfbridgeListRoomsAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeListRoomsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeMuteAction Mute a Confbridge user.
type ConfbridgeMuteAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string

	// If this parameter is not a complete channel name, the first channel with this prefix will be used.
	//   If this parameter is "all", all channels will be muted.
	//   If this parameter is "participants", all non-admin channels will be muted.
	Channel string
}

func (ConfbridgeMuteAction) ActionTypeName() string {
	return "ConfbridgeMute"
}
func (a *ConfbridgeMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeUnmuteAction Unmute a Confbridge user.
type ConfbridgeUnmuteAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string

	// If this parameter is not a complete channel name, the first channel with this prefix will be used.
	//   If this parameter is "all", all channels will be unmuted.
	//   If this parameter is "participants", all non-admin channels will be unmuted.
	Channel string
}

func (ConfbridgeUnmuteAction) ActionTypeName() string {
	return "ConfbridgeUnmute"
}
func (a *ConfbridgeUnmuteAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeUnmuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeKickAction Kick a Confbridge user.
type ConfbridgeKickAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string

	// If this parameter is "all", all channels will be kicked from the conference.
	//   If this parameter is "participants", all non-admin channels will be kicked from the conference.
	Channel string
}

func (ConfbridgeKickAction) ActionTypeName() string {
	return "ConfbridgeKick"
}
func (a *ConfbridgeKickAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeKickAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeLockAction Lock a Confbridge conference.
type ConfbridgeLockAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string
}

func (ConfbridgeLockAction) ActionTypeName() string {
	return "ConfbridgeLock"
}
func (a *ConfbridgeLockAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeLockAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeUnlockAction Unlock a Confbridge conference.
type ConfbridgeUnlockAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string
}

func (ConfbridgeUnlockAction) ActionTypeName() string {
	return "ConfbridgeUnlock"
}
func (a *ConfbridgeUnlockAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeUnlockAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeStartRecordAction Start recording a Confbridge conference.
type ConfbridgeStartRecordAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string

	RecordFile string
}

func (ConfbridgeStartRecordAction) ActionTypeName() string {
	return "ConfbridgeStartRecord"
}
func (a *ConfbridgeStartRecordAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeStartRecordAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeStopRecordAction Stop recording a Confbridge conference.
type ConfbridgeStopRecordAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string
}

func (ConfbridgeStopRecordAction) ActionTypeName() string {
	return "ConfbridgeStopRecord"
}
func (a *ConfbridgeStopRecordAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeStopRecordAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ConfbridgeSetSingleVideoSrcAction Set a conference user as the single video source distributed to all other participants.
type ConfbridgeSetSingleVideoSrcAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Conference string

	// If this parameter is not a complete channel name, the first channel with this prefix will be used.
	Channel string
}

func (ConfbridgeSetSingleVideoSrcAction) ActionTypeName() string {
	return "ConfbridgeSetSingleVideoSrc"
}
func (a *ConfbridgeSetSingleVideoSrcAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeSetSingleVideoSrcAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MeetmeMuteAction Mute a Meetme user.
type MeetmeMuteAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Meetme string

	Usernum string
}

func (MeetmeMuteAction) ActionTypeName() string {
	return "MeetmeMute"
}
func (a *MeetmeMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MeetmeUnmuteAction Unmute a Meetme user.
type MeetmeUnmuteAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Meetme string

	Usernum string
}

func (MeetmeUnmuteAction) ActionTypeName() string {
	return "MeetmeUnmute"
}
func (a *MeetmeUnmuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeUnmuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MeetmeListAction List participants in a conference.
type MeetmeListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Conference number.
	Conference string
}

func (MeetmeListAction) ActionTypeName() string {
	return "MeetmeList"
}
func (a *MeetmeListAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MeetmeListRoomsAction List active conferences.
type MeetmeListRoomsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (MeetmeListRoomsAction) ActionTypeName() string {
	return "MeetmeListRooms"
}
func (a *MeetmeListRoomsAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeListRoomsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueStatusAction Show queue status.
type QueueStatusAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Limit the response to the status of the specified queue.
	Queue string

	// Limit the response to the status of the specified member.
	Member string
}

func (QueueStatusAction) ActionTypeName() string {
	return "QueueStatus"
}
func (a *QueueStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueSummaryAction Show queue summary.
type QueueSummaryAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Queue for which the summary is requested.
	Queue string
}

func (QueueSummaryAction) ActionTypeName() string {
	return "QueueSummary"
}
func (a *QueueSummaryAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueSummaryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueAddAction Add interface to queue.
type QueueAddAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Queue's name.
	Queue string

	// The name of the interface (tech/name) to add to the queue.
	Interface string

	// A penalty (number) to apply to this member. Asterisk will distribute calls to members with higher
	//             penalties only after attempting to distribute calls to those with lower penalty.
	Penalty string

	// To pause or not the member initially (true/false or 1/0).
	Paused string

	// Text alias for the interface.
	MemberName string

	StateInterface string
}

func (QueueAddAction) ActionTypeName() string {
	return "QueueAdd"
}
func (a *QueueAddAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueAddAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueRemoveAction Remove interface from queue.
type QueueRemoveAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the queue to take action on.
	Queue string

	// The interface (tech/name) to remove from queue.
	Interface string
}

func (QueueRemoveAction) ActionTypeName() string {
	return "QueueRemove"
}
func (a *QueueRemoveAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueRemoveAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueuePauseAction Makes a queue member temporarily unavailable.
type QueuePauseAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the interface (tech/name) to pause or unpause.
	Interface string

	// Pause or unpause the interface. Set to 'true' to pause the member or 'false' to unpause.
	Paused string

	// The name of the queue in which to pause or unpause this member. If not specified, the member will be
	//             paused or unpaused in all the queues it is a member of.
	Queue string

	// Text description, returned in the event QueueMemberPaused.
	Reason string
}

func (QueuePauseAction) ActionTypeName() string {
	return "QueuePause"
}
func (a *QueuePauseAction) GetActionID() string {
	return a.ActionID
}
func (a *QueuePauseAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueLogAction Adds custom entry in queue_log.
type QueueLogAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Queue string

	Event string

	Uniqueid string

	Interface string

	Message string
}

func (QueueLogAction) ActionTypeName() string {
	return "QueueLog"
}
func (a *QueueLogAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueLogAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueuePenaltyAction Set the penalty for a queue member.
type QueuePenaltyAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The interface (tech/name) of the member whose penalty to change.
	Interface string

	// The new penalty (number) for the member. Must be nonnegative.
	Penalty string

	// If specified, only set the penalty for the member of this queue. Otherwise, set the penalty for the member
	//             in all queues to which the member belongs.
	Queue string
}

func (QueuePenaltyAction) ActionTypeName() string {
	return "QueuePenalty"
}
func (a *QueuePenaltyAction) GetActionID() string {
	return a.ActionID
}
func (a *QueuePenaltyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueMemberRingInUseAction Set the ringinuse value for a queue member.
type QueueMemberRingInUseAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Interface string

	RingInUse string

	Queue string
}

func (QueueMemberRingInUseAction) ActionTypeName() string {
	return "QueueMemberRingInUse"
}
func (a *QueueMemberRingInUseAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueMemberRingInUseAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueRuleAction Queue Rules.
type QueueRuleAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the rule in queuerules.conf whose contents to list.
	Rule string
}

func (QueueRuleAction) ActionTypeName() string {
	return "QueueRule"
}
func (a *QueueRuleAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueRuleAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueReloadAction Reload a queue, queues, or any sub-section of a queue or queues.
type QueueReloadAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the queue to take action on. If no queue name is specified, then all queues are affected.
	Queue string

	// Whether to reload the queue's members.
	Members string

	// Whether to reload queuerules.conf
	Rules string

	// Whether to reload the other queue options.
	Parameters string
}

func (QueueReloadAction) ActionTypeName() string {
	return "QueueReload"
}
func (a *QueueReloadAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueReloadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueResetAction Reset queue statistics.
type QueueResetAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the queue on which to reset statistics.
	Queue string
}

func (QueueResetAction) ActionTypeName() string {
	return "QueueReset"
}
func (a *QueueResetAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueResetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// QueueChangePriorityCallerAction Change priority of a caller on queue.
type QueueChangePriorityCallerAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the queue to take action on.
	Queue string

	// The caller (channel) to change priority on queue.
	Caller string

	// Priority value for change for caller on queue.
	Priority string
}

func (QueueChangePriorityCallerAction) ActionTypeName() string {
	return "QueueChangePriorityCaller"
}
func (a *QueueChangePriorityCallerAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueChangePriorityCallerAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PingAction Keepalive command.
type PingAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (PingAction) ActionTypeName() string {
	return "Ping"
}
func (a *PingAction) GetActionID() string {
	return a.ActionID
}
func (a *PingAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// EventsAction Control Event Flow.
type EventsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	EventMask string
}

func (EventsAction) ActionTypeName() string {
	return "Events"
}
func (a *EventsAction) GetActionID() string {
	return a.ActionID
}
func (a *EventsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// LogoffAction Logoff Manager.
type LogoffAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (LogoffAction) ActionTypeName() string {
	return "Logoff"
}
func (a *LogoffAction) GetActionID() string {
	return a.ActionID
}
func (a *LogoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// LoginAction Login Manager.
type LoginAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Username to login with as specified in manager.conf.
	Username string

	// Secret to login with as specified in manager.conf.
	Secret string
}

func (LoginAction) ActionTypeName() string {
	return "Login"
}
func (a *LoginAction) GetActionID() string {
	return a.ActionID
}
func (a *LoginAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ChallengeAction Generate Challenge for MD5 Auth.
type ChallengeAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Digest algorithm to use in the challenge. Valid values are:
	AuthType string
}

func (ChallengeAction) ActionTypeName() string {
	return "Challenge"
}
func (a *ChallengeAction) GetActionID() string {
	return a.ActionID
}
func (a *ChallengeAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// HangupAction Hangup channel.
type HangupAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The exact channel name to be hungup, or to use a regular expression, set this parameter to: /regex/
	//   Example exact channel: SIP/provider-0000012a
	//   Example regular expression: /^SIP/provider-.*$/
	Channel string

	// Numeric hangup cause.
	Cause string
}

func (HangupAction) ActionTypeName() string {
	return "Hangup"
}
func (a *HangupAction) GetActionID() string {
	return a.ActionID
}
func (a *HangupAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// StatusAction List channel status.
type StatusAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the channel to query for status.
	Channel string

	// Comma  separated list of variable to include.
	Variables string

	// If set to "true", the Status event will include all channel variables for
	//             the requested channel(s).
	AllVariables string
}

func (StatusAction) ActionTypeName() string {
	return "Status"
}
func (a *StatusAction) GetActionID() string {
	return a.ActionID
}
func (a *StatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SetvarAction Sets a channel variable or function value.
type SetvarAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel to set variable for.
	Channel string

	// Variable name, function or expression.
	Variable string

	// Variable or function value.
	Value string
}

func (SetvarAction) ActionTypeName() string {
	return "Setvar"
}
func (a *SetvarAction) GetActionID() string {
	return a.ActionID
}
func (a *SetvarAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// GetvarAction Gets a channel variable or function value.
type GetvarAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel to read variable from.
	Channel string

	// Variable name, function or expression.
	Variable string
}

func (GetvarAction) ActionTypeName() string {
	return "Getvar"
}
func (a *GetvarAction) GetActionID() string {
	return a.ActionID
}
func (a *GetvarAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// GetConfigAction Retrieve configuration.
type GetConfigAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Configuration filename (e.g. ).
	Filename string

	// Category in configuration file.
	Category string

	// A comma separated list of
	//             =
	//
	//             expressions which will cause only categories whose variables match all expressions
	//             to be considered. The special variable name
	//
	//             can be used to control whether templates are included. Passing
	//
	//             as the value will include templates
	//             along with normal categories. Passing
	//
	//             as the value will restrict the operation to
	//             ONLY templates. Not specifying a  expression
	//             results in the default behavior which is to not include templates.
	Filter string
}

func (GetConfigAction) ActionTypeName() string {
	return "GetConfig"
}
func (a *GetConfigAction) GetActionID() string {
	return a.ActionID
}
func (a *GetConfigAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// GetConfigJSONAction Retrieve configuration (JSON format).
type GetConfigJSONAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Configuration filename (e.g. ).
	Filename string

	// Category in configuration file.
	Category string

	// A comma separated list of
	//             =
	//
	//             expressions which will cause only categories whose variables match all expressions
	//             to be considered. The special variable name
	//
	//             can be used to control whether templates are included. Passing
	//
	//             as the value will include templates
	//             along with normal categories. Passing
	//
	//             as the value will restrict the operation to
	//             ONLY templates. Not specifying a  expression
	//             results in the default behavior which is to not include templates.
	Filter string
}

func (GetConfigJSONAction) ActionTypeName() string {
	return "GetConfigJSON"
}
func (a *GetConfigJSONAction) GetActionID() string {
	return a.ActionID
}
func (a *GetConfigJSONAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// UpdateConfigAction Update basic configuration.
type UpdateConfigAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Configuration filename to read (e.g. ).
	SrcFilename string

	// Configuration filename to write (e.g. )
	DstFilename string

	// Whether or not a reload should take place (or name of specific module).
	Reload string

	// Whether the effective category contents should be preserved on template change. Default is true (pre 13.2
	//             behavior).
	PreserveEffectiveContext string

	// Action to take.
	//   0's represent 6 digit number beginning with 000000.
	Action000000 string

	// Category to operate on.
	//   0's represent 6 digit number beginning with 000000.
	Cat000000 string

	// Variable to work on.
	//   0's represent 6 digit number beginning with 000000.
	Var000000 string

	// Value to work on.
	//   0's represent 6 digit number beginning with 000000.
	Value000000 string

	// Extra match required to match line.
	//   0's represent 6 digit number beginning with 000000.
	Match000000 string

	// Line in category to operate on (used with delete and insert actions).
	//   0's represent 6 digit number beginning with 000000.
	Line000000 string

	// A comma separated list of action-specific options.
	//   The following actions share the same options...
	//   0's represent 6 digit number beginning with 000000.
	Options000000 string
}

func (UpdateConfigAction) ActionTypeName() string {
	return "UpdateConfig"
}
func (a *UpdateConfigAction) GetActionID() string {
	return a.ActionID
}
func (a *UpdateConfigAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// CreateConfigAction Creates an empty file in the configuration directory.
type CreateConfigAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The configuration filename to create (e.g. ).
	Filename string
}

func (CreateConfigAction) ActionTypeName() string {
	return "CreateConfig"
}
func (a *CreateConfigAction) GetActionID() string {
	return a.ActionID
}
func (a *CreateConfigAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ListCategoriesAction List categories in configuration file.
type ListCategoriesAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Configuration filename (e.g. ).
	Filename string
}

func (ListCategoriesAction) ActionTypeName() string {
	return "ListCategories"
}
func (a *ListCategoriesAction) GetActionID() string {
	return a.ActionID
}
func (a *ListCategoriesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// RedirectAction Redirect (transfer) a call.
type RedirectAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel to redirect.
	Channel string

	// Second call leg to transfer (optional).
	ExtraChannel string

	// Extension to transfer to.
	Exten string

	// Extension to transfer extrachannel to (optional).
	ExtraExten string

	// Context to transfer to.
	Context string

	// Context to transfer extrachannel to (optional).
	ExtraContext string

	// Priority to transfer to.
	Priority string

	// Priority to transfer extrachannel to (optional).
	ExtraPriority string
}

func (RedirectAction) ActionTypeName() string {
	return "Redirect"
}
func (a *RedirectAction) GetActionID() string {
	return a.ActionID
}
func (a *RedirectAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// AtxferAction Attended transfer.
type AtxferAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Transferer's channel.
	Channel string

	// Extension to transfer to.
	Exten string

	// Context to transfer to.
	Context string
}

func (AtxferAction) ActionTypeName() string {
	return "Atxfer"
}
func (a *AtxferAction) GetActionID() string {
	return a.ActionID
}
func (a *AtxferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// CancelAtxferAction Cancel an attended transfer.
type CancelAtxferAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The transferer channel.
	Channel string
}

func (CancelAtxferAction) ActionTypeName() string {
	return "CancelAtxfer"
}
func (a *CancelAtxferAction) GetActionID() string {
	return a.ActionID
}
func (a *CancelAtxferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// OriginateAction Originate a call.
type OriginateAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel name to call.
	Channel string

	// Extension to use (requires  and
	//             )
	Exten string

	// Context to use (requires  and
	//             )
	Context string

	// Priority to use (requires  and
	//             )
	Priority string

	// Application to execute.
	Application string

	// Data to use (requires ).
	Data string

	// How long to wait for call to be answered (in ms.).
	Timeout string

	// Caller ID to be set on the outgoing channel.
	CallerID string

	// Channel variable to set, multiple Variable: headers are allowed.
	Variable string

	// Account code.
	Account string

	// Set to  to force call bridge on early media..
	EarlyMedia string

	// Set to  for fast origination.
	Async string

	// Comma-separated list of codecs to use for this call.
	Codecs string

	// Channel UniqueId to be set on the channel.
	ChannelId string

	// Channel UniqueId to be set on the second local channel.
	OtherChannelId string
}

func (OriginateAction) ActionTypeName() string {
	return "Originate"
}
func (a *OriginateAction) GetActionID() string {
	return a.ActionID
}
func (a *OriginateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// CommandAction Execute Asterisk CLI Command.
type CommandAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Asterisk CLI command to run.
	Command string
}

func (CommandAction) ActionTypeName() string {
	return "Command"
}
func (a *CommandAction) GetActionID() string {
	return a.ActionID
}
func (a *CommandAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ExtensionStateAction Check Extension Status.
type ExtensionStateAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Extension to check state on.
	Exten string

	// Context for extension.
	Context string
}

func (ExtensionStateAction) ActionTypeName() string {
	return "ExtensionState"
}
func (a *ExtensionStateAction) GetActionID() string {
	return a.ActionID
}
func (a *ExtensionStateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PresenceStateAction Check Presence State
type PresenceStateAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Presence Provider to check the state of
	Provider string
}

func (PresenceStateAction) ActionTypeName() string {
	return "PresenceState"
}
func (a *PresenceStateAction) GetActionID() string {
	return a.ActionID
}
func (a *PresenceStateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// AbsoluteTimeoutAction Set absolute timeout.
type AbsoluteTimeoutAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel name to hangup.
	Channel string

	// Maximum duration of the call (sec).
	Timeout string
}

func (AbsoluteTimeoutAction) ActionTypeName() string {
	return "AbsoluteTimeout"
}
func (a *AbsoluteTimeoutAction) GetActionID() string {
	return a.ActionID
}
func (a *AbsoluteTimeoutAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MailboxStatusAction Check mailbox.
type MailboxStatusAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Full mailbox ID @.
	Mailbox string
}

func (MailboxStatusAction) ActionTypeName() string {
	return "MailboxStatus"
}
func (a *MailboxStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *MailboxStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MailboxCountAction Check Mailbox Message Count.
type MailboxCountAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Full mailbox ID @.
	Mailbox string
}

func (MailboxCountAction) ActionTypeName() string {
	return "MailboxCount"
}
func (a *MailboxCountAction) GetActionID() string {
	return a.ActionID
}
func (a *MailboxCountAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ListCommandsAction List available manager commands.
type ListCommandsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (ListCommandsAction) ActionTypeName() string {
	return "ListCommands"
}
func (a *ListCommandsAction) GetActionID() string {
	return a.ActionID
}
func (a *ListCommandsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SendTextAction Sends a text message to channel. A content type can be optionally specified. If not set
//       it is set to an empty string allowing a custom handler to default it as it sees fit.
type SendTextAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel to send message to.
	Channel string

	// Message to send.
	Message string

	// The type of content in the message
	ContentType string
}

func (SendTextAction) ActionTypeName() string {
	return "SendText"
}
func (a *SendTextAction) GetActionID() string {
	return a.ActionID
}
func (a *SendTextAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// UserEventAction Send an arbitrary event.
type UserEventAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Event string to send.
	UserEvent string

	// Content1.
	Header1 string

	// ContentN.
	HeaderN string
}

func (UserEventAction) ActionTypeName() string {
	return "UserEvent"
}
func (a *UserEventAction) GetActionID() string {
	return a.ActionID
}
func (a *UserEventAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// WaitEventAction Wait for an event to occur.
type WaitEventAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Maximum time (in seconds) to wait for events,  means forever.
	Timeout string
}

func (WaitEventAction) ActionTypeName() string {
	return "WaitEvent"
}
func (a *WaitEventAction) GetActionID() string {
	return a.ActionID
}
func (a *WaitEventAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// CoreSettingsAction Show PBX core settings (version etc).
type CoreSettingsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (CoreSettingsAction) ActionTypeName() string {
	return "CoreSettings"
}
func (a *CoreSettingsAction) GetActionID() string {
	return a.ActionID
}
func (a *CoreSettingsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// CoreStatusAction Show PBX core status variables.
type CoreStatusAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (CoreStatusAction) ActionTypeName() string {
	return "CoreStatus"
}
func (a *CoreStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *CoreStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ReloadAction Send a reload event.
type ReloadAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Name of the module to reload.
	Module string
}

func (ReloadAction) ActionTypeName() string {
	return "Reload"
}
func (a *ReloadAction) GetActionID() string {
	return a.ActionID
}
func (a *ReloadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// CoreShowChannelsAction List currently active channels.
type CoreShowChannelsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (CoreShowChannelsAction) ActionTypeName() string {
	return "CoreShowChannels"
}
func (a *CoreShowChannelsAction) GetActionID() string {
	return a.ActionID
}
func (a *CoreShowChannelsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// LoggerRotateAction Reload and rotate the Asterisk logger.
type LoggerRotateAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (LoggerRotateAction) ActionTypeName() string {
	return "LoggerRotate"
}
func (a *LoggerRotateAction) GetActionID() string {
	return a.ActionID
}
func (a *LoggerRotateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ModuleLoadAction Module management.
type ModuleLoadAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Asterisk module name (including .so extension) or subsystem identifier:
	Module string

	// The operation to be done on module. Subsystem identifiers may only
	//             be reloaded.
	//   If no module is specified for a  loadtype,
	//             all modules are reloaded.
	LoadType string
}

func (ModuleLoadAction) ActionTypeName() string {
	return "ModuleLoad"
}
func (a *ModuleLoadAction) GetActionID() string {
	return a.ActionID
}
func (a *ModuleLoadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ModuleCheckAction Check if module is loaded.
type ModuleCheckAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Asterisk module name (not including extension).
	Module string
}

func (ModuleCheckAction) ActionTypeName() string {
	return "ModuleCheck"
}
func (a *ModuleCheckAction) GetActionID() string {
	return a.ActionID
}
func (a *ModuleCheckAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// AOCMessageAction Generate an Advice of Charge message on a channel.
type AOCMessageAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel name to generate the AOC message on.
	Channel string

	// Partial channel prefix. By using this option one can match the beginning part
	//             of a channel name without having to put the entire name in. For example
	//             if a channel name is SIP/snom-00000001 and this value is set to SIP/snom, then
	//             that channel matches and the message will be sent. Note however that only
	//             the first matched channel has the message sent on it.
	ChannelPrefix string

	// Defines what type of AOC message to create, AOC-D or AOC-E
	MsgType string

	// Defines what kind of charge this message represents.
	ChargeType string

	// This represents the amount of units charged. The ETSI AOC standard specifies that
	//             this value along with the optional UnitType value are entries in a list. To accommodate this
	//             these values take an index value starting at 0 which can be used to generate this list of
	//             unit entries. For Example, If two unit entires were required this could be achieved by setting the
	//             paramter UnitAmount(0)=1234 and UnitAmount(1)=5678. Note that UnitAmount at index 0 is
	//             required when ChargeType=Unit, all other entries in the list are optional.
	UnitAmount string

	// Defines the type of unit. ETSI AOC standard specifies this as an integer
	//             value between 1 and 16, but this value is left open to accept any positive
	//             integer. Like the UnitAmount parameter, this value represents a list entry
	//             and has an index parameter that starts at 0.
	UnitType string

	// Specifies the currency's name. Note that this value is truncated after 10 characters.
	CurrencyName string

	// Specifies the charge unit amount as a positive integer. This value is required
	//             when ChargeType==Currency.
	CurrencyAmount string

	// Specifies the currency multiplier. This value is required when ChargeType==Currency.
	CurrencyMultiplier string

	// Defines what kind of AOC-D total is represented.
	TotalType string

	// Represents a billing ID associated with an AOC-D or AOC-E message. Note
	//             that only the first 3 items of the enum are valid AOC-D billing IDs
	AOCBillingId string

	// Charging association identifier. This is optional for AOC-E and can be
	//             set to any value between -32768 and 32767
	ChargingAssociationId string

	// Represents the charging association party number. This value is optional
	//             for AOC-E.
	ChargingAssociationNumber string

	// Integer representing the charging plan associated with the ChargingAssociationNumber.
	//             The value is bits 7 through 1 of the Q.931 octet containing the type-of-number and
	//             numbering-plan-identification fields.
	ChargingAssociationPlan string
}

func (AOCMessageAction) ActionTypeName() string {
	return "AOCMessage"
}
func (a *AOCMessageAction) GetActionID() string {
	return a.ActionID
}
func (a *AOCMessageAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// FilterAction Dynamically add filters for the current manager session.
type FilterAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Operation string

	// Filters can be whitelist or blacklist
	//   Example whitelist filter: "Event: Newchannel"
	//   Example blacklist filter: "!Channel: DAHDI.*"
	//   This filter option is used to whitelist or blacklist events per user to be
	//             reported with regular expressions and are allowed if both the regex matches
	//             and the user has read access as defined in manager.conf. Filters are assumed to be for whitelisting
	//             unless preceeded by an exclamation point, which marks it as being black.
	//             Evaluation of the filters is as follows:
	//   - If no filters are configured all events are reported as normal.
	//   - If there are white filters only: implied black all filter processed first, then white filters.
	//   - If there are black filters only: implied white all filter processed first, then black filters.
	//   - If there are both white and black filters: implied black all filter processed first, then white
	//             filters, and lastly black filters.
	Filter string
}

func (FilterAction) ActionTypeName() string {
	return "Filter"
}
func (a *FilterAction) GetActionID() string {
	return a.ActionID
}
func (a *FilterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// FilterListAction Show current event filters for this session
type FilterListAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (FilterListAction) ActionTypeName() string {
	return "FilterList"
}
func (a *FilterListAction) GetActionID() string {
	return a.ActionID
}
func (a *FilterListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BlindTransferAction Blind transfer channel(s) to the given destination
type BlindTransferAction struct {
	Channel string

	Context string

	Exten string

	// ActionID for tx - compensate
	ActionID string
}

func (BlindTransferAction) ActionTypeName() string {
	return "BlindTransfer"
}
func (a *BlindTransferAction) GetActionID() string {
	return a.ActionID
}
func (a *BlindTransferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeAction Bridge two channels already in the PBX.
type BridgeAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel to Bridge to Channel2.
	Channel1 string

	// Channel to Bridge to Channel1.
	Channel2 string

	// Play courtesy tone to Channel 2.
	Tone string
}

func (BridgeAction) ActionTypeName() string {
	return "Bridge"
}
func (a *BridgeAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DBGetAction Get DB Entry.
type DBGetAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Family string

	Key string
}

func (DBGetAction) ActionTypeName() string {
	return "DBGet"
}
func (a *DBGetAction) GetActionID() string {
	return a.ActionID
}
func (a *DBGetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DBPutAction Put DB entry.
type DBPutAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Family string

	Key string

	Val string
}

func (DBPutAction) ActionTypeName() string {
	return "DBPut"
}
func (a *DBPutAction) GetActionID() string {
	return a.ActionID
}
func (a *DBPutAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DBDelAction Delete DB entry.
type DBDelAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Family string

	Key string
}

func (DBDelAction) ActionTypeName() string {
	return "DBDel"
}
func (a *DBDelAction) GetActionID() string {
	return a.ActionID
}
func (a *DBDelAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DBDelTreeAction Delete DB Tree.
type DBDelTreeAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	Family string

	Key string
}

func (DBDelTreeAction) ActionTypeName() string {
	return "DBDelTree"
}
func (a *DBDelTreeAction) GetActionID() string {
	return a.ActionID
}
func (a *DBDelTreeAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ShowDialPlanAction Show dialplan contexts and extensions
type ShowDialPlanAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Show a specific extension.
	Extension string

	// Show a specific context.
	Context string
}

func (ShowDialPlanAction) ActionTypeName() string {
	return "ShowDialPlan"
}
func (a *ShowDialPlanAction) GetActionID() string {
	return a.ActionID
}
func (a *ShowDialPlanAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ExtensionStateListAction List the current known extension states.
type ExtensionStateListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (ExtensionStateListAction) ActionTypeName() string {
	return "ExtensionStateList"
}
func (a *ExtensionStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *ExtensionStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MessageSendAction Send an out of call message to an endpoint.
type MessageSendAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The URI the message is to be sent to.
	To string

	// A From URI for the message if needed for the
	//             message technology being used to send this message.
	From string

	// The message body text. This must not contain any newlines as that
	//             conflicts with the AMI protocol.
	Body string

	// Text bodies requiring the use of newlines have to be base64 encoded
	//             in this field. Base64Body will be decoded before being sent out.
	//             Base64Body takes precedence over Body.
	Base64Body string

	// Message variable to set, multiple Variable: headers are
	//             allowed. The header value is a comma separated list of
	//             name=value pairs.
	Variable string
}

func (MessageSendAction) ActionTypeName() string {
	return "MessageSend"
}
func (a *MessageSendAction) GetActionID() string {
	return a.ActionID
}
func (a *MessageSendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// LocalOptimizeAwayAction Optimize away a local channel when possible.
type LocalOptimizeAwayAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The channel name to optimize away.
	Channel string
}

func (LocalOptimizeAwayAction) ActionTypeName() string {
	return "LocalOptimizeAway"
}
func (a *LocalOptimizeAwayAction) GetActionID() string {
	return a.ActionID
}
func (a *LocalOptimizeAwayAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeTechnologyListAction List available bridging technologies and their statuses.
type BridgeTechnologyListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (BridgeTechnologyListAction) ActionTypeName() string {
	return "BridgeTechnologyList"
}
func (a *BridgeTechnologyListAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeTechnologyListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeTechnologySuspendAction Suspend a bridging technology.
type BridgeTechnologySuspendAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the bridging technology to suspend.
	BridgeTechnology string
}

func (BridgeTechnologySuspendAction) ActionTypeName() string {
	return "BridgeTechnologySuspend"
}
func (a *BridgeTechnologySuspendAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeTechnologySuspendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeTechnologyUnsuspendAction Unsuspend a bridging technology.
type BridgeTechnologyUnsuspendAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the bridging technology to unsuspend.
	BridgeTechnology string
}

func (BridgeTechnologyUnsuspendAction) ActionTypeName() string {
	return "BridgeTechnologyUnsuspend"
}
func (a *BridgeTechnologyUnsuspendAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeTechnologyUnsuspendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeListAction Get a list of bridges in the system.
type BridgeListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Optional type for filtering the resulting list of bridges.
	BridgeType string
}

func (BridgeListAction) ActionTypeName() string {
	return "BridgeList"
}
func (a *BridgeListAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeInfoAction Get information about a bridge.
type BridgeInfoAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The unique ID of the bridge about which to retrieve information.
	BridgeUniqueid string
}

func (BridgeInfoAction) ActionTypeName() string {
	return "BridgeInfo"
}
func (a *BridgeInfoAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeInfoAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeDestroyAction Destroy a bridge.
type BridgeDestroyAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The unique ID of the bridge to destroy.
	BridgeUniqueid string
}

func (BridgeDestroyAction) ActionTypeName() string {
	return "BridgeDestroy"
}
func (a *BridgeDestroyAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeDestroyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// BridgeKickAction Kick a channel from a bridge.
type BridgeKickAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The unique ID of the bridge containing the channel to
	//             destroy. This parameter can be omitted, or supplied to insure
	//             that the channel is not removed from the wrong bridge.
	BridgeUniqueid string

	// The channel to kick out of a bridge.
	Channel string
}

func (BridgeKickAction) ActionTypeName() string {
	return "BridgeKick"
}
func (a *BridgeKickAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeKickAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PresenceStateListAction List the current known presence states.
type PresenceStateListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (PresenceStateListAction) ActionTypeName() string {
	return "PresenceStateList"
}
func (a *PresenceStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *PresenceStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ParkinglotsAction Get a list of parking lots
type ParkinglotsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (ParkinglotsAction) ActionTypeName() string {
	return "Parkinglots"
}
func (a *ParkinglotsAction) GetActionID() string {
	return a.ActionID
}
func (a *ParkinglotsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ParkedCallsAction List parked calls.
type ParkedCallsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// If specified, only show parked calls from the parking lot with this name.
	ParkingLot string
}

func (ParkedCallsAction) ActionTypeName() string {
	return "ParkedCalls"
}
func (a *ParkedCallsAction) GetActionID() string {
	return a.ActionID
}
func (a *ParkedCallsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ParkAction Park a channel.
type ParkAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel name to park.
	Channel string

	// Channel name to use when constructing the dial string that will be dialed if the parked channel
	//             times out. If  is in a two party bridge with
	//             , then  will receive an announcement and be
	//             treated as having parked  in the same manner as the Park Call DTMF feature.
	TimeoutChannel string

	// If specified, then this channel will receive an announcement when
	//
	//             is parked if  is in a state where it can receive announcements
	//             (AnnounceChannel must be bridged).  has no bearing on the actual
	//             state of the parked call.
	AnnounceChannel string

	// Overrides the timeout of the parking lot for this park action. Specified in milliseconds, but will be
	//             converted to
	//             seconds. Use a value of 0 to disable the timeout.
	Timeout string

	// The parking lot to use when parking the channel
	Parkinglot string
}

func (ParkAction) ActionTypeName() string {
	return "Park"
}
func (a *ParkAction) GetActionID() string {
	return a.ActionID
}
func (a *ParkAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MWIGetAction Get selected mailboxes with message counts.
type MWIGetAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Mailbox ID in the form of
	//             // for all mailboxes matching the regular
	//             expression. Otherwise it is for a specific mailbox.
	Mailbox string
}

func (MWIGetAction) ActionTypeName() string {
	return "MWIGet"
}
func (a *MWIGetAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIGetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MWIDeleteAction Delete selected mailboxes.
type MWIDeleteAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Mailbox ID in the form of
	//             // for all mailboxes matching the regular
	//             expression. Otherwise it is for a specific mailbox.
	Mailbox string
}

func (MWIDeleteAction) ActionTypeName() string {
	return "MWIDelete"
}
func (a *MWIDeleteAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIDeleteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MWIUpdateAction Update the mailbox message counts.
type MWIUpdateAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Specific mailbox ID.
	Mailbox string

	// The number of old messages in the mailbox. Defaults
	//             to zero if missing.
	OldMessages string

	// The number of new messages in the mailbox. Defaults
	//             to zero if missing.
	NewMessages string
}

func (MWIUpdateAction) ActionTypeName() string {
	return "MWIUpdate"
}
func (a *MWIUpdateAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIUpdateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowRegistrationsInboundAction Lists PJSIP inbound registrations.
type PJSIPShowRegistrationsInboundAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowRegistrationsInboundAction) ActionTypeName() string {
	return "PJSIPShowRegistrationsInbound"
}
func (a *PJSIPShowRegistrationsInboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowRegistrationsInboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowRegistrationInboundContactStatusesAction Lists ContactStatuses for PJSIP inbound registrations.
type PJSIPShowRegistrationInboundContactStatusesAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowRegistrationInboundContactStatusesAction) ActionTypeName() string {
	return "PJSIPShowRegistrationInboundContactStatuses"
}
func (a *PJSIPShowRegistrationInboundContactStatusesAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowRegistrationInboundContactStatusesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPNotifyAction Send a NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
type PJSIPNotifyAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The endpoint to which to send the NOTIFY.
	Endpoint string

	// Abritrary URI to which to send the NOTIFY.
	URI string

	// Channel name to send the NOTIFY. Must be a PJSIP channel.
	channel string

	// Appends variables as headers/content to the NOTIFY. If the variable is
	//             named , then the value will compose the body
	//             of the message if another variable sets .
	//             =
	Variable string
}

func (PJSIPNotifyAction) ActionTypeName() string {
	return "PJSIPNotify"
}
func (a *PJSIPNotifyAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPNotifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPQualifyAction Qualify a chan_pjsip endpoint.
type PJSIPQualifyAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The endpoint you want to qualify.
	Endpoint string
}

func (PJSIPQualifyAction) ActionTypeName() string {
	return "PJSIPQualify"
}
func (a *PJSIPQualifyAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPQualifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowEndpointsAction Lists PJSIP endpoints.
type PJSIPShowEndpointsAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowEndpointsAction) ActionTypeName() string {
	return "PJSIPShowEndpoints"
}
func (a *PJSIPShowEndpointsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowEndpointsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowEndpointAction Detail listing of an endpoint and its objects.
type PJSIPShowEndpointAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The endpoint to list.
	Endpoint string
}

func (PJSIPShowEndpointAction) ActionTypeName() string {
	return "PJSIPShowEndpoint"
}
func (a *PJSIPShowEndpointAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowEndpointAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowAorsAction Lists PJSIP AORs.
type PJSIPShowAorsAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowAorsAction) ActionTypeName() string {
	return "PJSIPShowAors"
}
func (a *PJSIPShowAorsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowAorsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowAuthsAction Lists PJSIP Auths.
type PJSIPShowAuthsAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowAuthsAction) ActionTypeName() string {
	return "PJSIPShowAuths"
}
func (a *PJSIPShowAuthsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowAuthsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowContactsAction Lists PJSIP Contacts.
type PJSIPShowContactsAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowContactsAction) ActionTypeName() string {
	return "PJSIPShowContacts"
}
func (a *PJSIPShowContactsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowContactsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// FAXSessionsAction Lists active FAX sessions
type FAXSessionsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (FAXSessionsAction) ActionTypeName() string {
	return "FAXSessions"
}
func (a *FAXSessionsAction) GetActionID() string {
	return a.ActionID
}
func (a *FAXSessionsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// FAXSessionAction Responds with a detailed description of a single FAX session
type FAXSessionAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The session ID of the fax the user is interested in.
	SessionNumber string
}

func (FAXSessionAction) ActionTypeName() string {
	return "FAXSession"
}
func (a *FAXSessionAction) GetActionID() string {
	return a.ActionID
}
func (a *FAXSessionAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// FAXStatsAction Responds with fax statistics
type FAXStatsAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (FAXStatsAction) ActionTypeName() string {
	return "FAXStats"
}
func (a *FAXStatsAction) GetActionID() string {
	return a.ActionID
}
func (a *FAXStatsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPUnregisterAction Unregister an outbound registration.
type PJSIPUnregisterAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The outbound registration to unregister or '*all' to unregister them all.
	Registration string
}

func (PJSIPUnregisterAction) ActionTypeName() string {
	return "PJSIPUnregister"
}
func (a *PJSIPUnregisterAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPUnregisterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPRegisterAction Register an outbound registration.
type PJSIPRegisterAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The outbound registration to register or '*all' to register them all.
	Registration string
}

func (PJSIPRegisterAction) ActionTypeName() string {
	return "PJSIPRegister"
}
func (a *PJSIPRegisterAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPRegisterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowRegistrationsOutboundAction Lists PJSIP outbound registrations.
type PJSIPShowRegistrationsOutboundAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowRegistrationsOutboundAction) ActionTypeName() string {
	return "PJSIPShowRegistrationsOutbound"
}
func (a *PJSIPShowRegistrationsOutboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowRegistrationsOutboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MuteAudioAction Mute an audio stream.
type MuteAudioAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The channel you want to mute.
	Channel string

	Direction string

	State string
}

func (MuteAudioAction) ActionTypeName() string {
	return "MuteAudio"
}
func (a *MuteAudioAction) GetActionID() string {
	return a.ActionID
}
func (a *MuteAudioAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// MonitorAction Monitor a channel.
type MonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Used to specify the channel to record.
	Channel string

	// Is the name of the file created in the monitor spool directory.
	//             Defaults to the same name as the channel (with slashes replaced with dashes).
	File string

	// Is the audio recording format. Defaults to .
	Format string

	// Boolean parameter as to whether to mix the input and output channels
	//             together after the recording is finished.
	Mix string
}

func (MonitorAction) ActionTypeName() string {
	return "Monitor"
}
func (a *MonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *MonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// StopMonitorAction Stop monitoring a channel.
type StopMonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the channel monitored.
	Channel string
}

func (StopMonitorAction) ActionTypeName() string {
	return "StopMonitor"
}
func (a *StopMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *StopMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// ChangeMonitorAction Change monitoring filename of a channel.
type ChangeMonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Used to specify the channel to record.
	Channel string

	// Is the new name of the file created in the
	//             monitor spool directory.
	File string
}

func (ChangeMonitorAction) ActionTypeName() string {
	return "ChangeMonitor"
}
func (a *ChangeMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *ChangeMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PauseMonitorAction Pause monitoring of a channel.
type PauseMonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Used to specify the channel to record.
	Channel string
}

func (PauseMonitorAction) ActionTypeName() string {
	return "PauseMonitor"
}
func (a *PauseMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *PauseMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// UnpauseMonitorAction Unpause monitoring of a channel.
type UnpauseMonitorAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Used to specify the channel to record.
	Channel string
}

func (UnpauseMonitorAction) ActionTypeName() string {
	return "UnpauseMonitor"
}
func (a *UnpauseMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *UnpauseMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowSubscriptionsInboundAction Lists subscriptions.
type PJSIPShowSubscriptionsInboundAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowSubscriptionsInboundAction) ActionTypeName() string {
	return "PJSIPShowSubscriptionsInbound"
}
func (a *PJSIPShowSubscriptionsInboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowSubscriptionsInboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowSubscriptionsOutboundAction Lists subscriptions.
type PJSIPShowSubscriptionsOutboundAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowSubscriptionsOutboundAction) ActionTypeName() string {
	return "PJSIPShowSubscriptionsOutbound"
}
func (a *PJSIPShowSubscriptionsOutboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowSubscriptionsOutboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// PJSIPShowResourceListsAction Displays settings for configured resource lists.
type PJSIPShowResourceListsAction struct {

	// ActionID for tx - compensate
	ActionID string
}

func (PJSIPShowResourceListsAction) ActionTypeName() string {
	return "PJSIPShowResourceLists"
}
func (a *PJSIPShowResourceListsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowResourceListsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// AGIAction Add an AGI command to execute by Async AGI.
type AGIAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Channel that is currently in Async AGI.
	Channel string

	// Application to execute.
	Command string

	// This will be sent back in CommandID header of AsyncAGI exec
	//             event notification.
	CommandID string
}

func (AGIAction) ActionTypeName() string {
	return "AGI"
}
func (a *AGIAction) GetActionID() string {
	return a.ActionID
}
func (a *AGIAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// DeviceStateListAction List the current known device states.
type DeviceStateListAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string
}

func (DeviceStateListAction) ActionTypeName() string {
	return "DeviceStateList"
}
func (a *DeviceStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *DeviceStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SorceryMemoryCacheExpireObjectAction Expire (remove) an object from a sorcery memory cache.
type SorceryMemoryCacheExpireObjectAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the cache to expire the object from.
	Cache string

	// The name of the object to expire.
	Object string
}

func (SorceryMemoryCacheExpireObjectAction) ActionTypeName() string {
	return "SorceryMemoryCacheExpireObject"
}
func (a *SorceryMemoryCacheExpireObjectAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheExpireObjectAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SorceryMemoryCacheExpireAction Expire (remove) ALL objects from a sorcery memory cache.
type SorceryMemoryCacheExpireAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the cache to expire all objects from.
	Cache string
}

func (SorceryMemoryCacheExpireAction) ActionTypeName() string {
	return "SorceryMemoryCacheExpire"
}
func (a *SorceryMemoryCacheExpireAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheExpireAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SorceryMemoryCacheStaleObjectAction Mark an object in a sorcery memory cache as stale.
type SorceryMemoryCacheStaleObjectAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the cache to mark the object as stale in.
	Cache string

	// The name of the object to mark as stale.
	Object string

	// If true, then immediately reload the object from the backend cache instead of waiting for the next
	//             retrieval
	Reload string
}

func (SorceryMemoryCacheStaleObjectAction) ActionTypeName() string {
	return "SorceryMemoryCacheStaleObject"
}
func (a *SorceryMemoryCacheStaleObjectAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheStaleObjectAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SorceryMemoryCacheStaleAction Marks ALL objects in a sorcery memory cache as stale.
type SorceryMemoryCacheStaleAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the cache to mark all object as stale in.
	Cache string
}

func (SorceryMemoryCacheStaleAction) ActionTypeName() string {
	return "SorceryMemoryCacheStale"
}
func (a *SorceryMemoryCacheStaleAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheStaleAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// SorceryMemoryCachePopulateAction Expire all objects from a memory cache and populate it with all objects from the backend.
type SorceryMemoryCachePopulateAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// The name of the cache to populate.
	Cache string
}

func (SorceryMemoryCachePopulateAction) ActionTypeName() string {
	return "SorceryMemoryCachePopulate"
}
func (a *SorceryMemoryCachePopulateAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCachePopulateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

// JabberSendAction Sends a message to a Jabber Client.
type JabberSendAction struct {

	// ActionID for this transaction. Will be returned.
	ActionID string

	// Client or transport Asterisk uses to connect to JABBER.
	Jabber string

	// XMPP/Jabber JID (Name) of recipient.
	JID string

	// Message to be sent to the buddy.
	Message string
}

func (JabberSendAction) ActionTypeName() string {
	return "JabberSend"
}
func (a *JabberSendAction) GetActionID() string {
	return a.ActionID
}
func (a *JabberSendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

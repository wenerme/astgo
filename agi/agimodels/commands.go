package agimodels

// GosubCommand Cause the channel to execute the specified dialplan subroutine.
//
// Cause the channel to execute the specified dialplan subroutine, returning to the dialplan with execution of a Return().
type GosubCommand struct {
	Context          string
	Extension        string
	Priority         int
	OptionalArgument *string
}

func (cmd GosubCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Context, cmd.Extension, cmd.Priority, cmd.OptionalArgument}
	return joinCommand(s), nil
}
func (cmd GosubCommand) CommandString() string {
	return "GOSUB"
}

func (cmd GosubCommand) SetOptionalArgument(v string) GosubCommand {
	cmd.OptionalArgument = &v
	return cmd
}

func (c *Client) Gosub(context string, extension string, priority int) Response {
	return c.Handler.Command(GosubCommand{
		Context:   context,
		Extension: extension,
		Priority:  priority,
	})
}

// AnswerCommand Answer channel
//
// Answers channel if not already in answer state. Returns `-1` on channel failure, or `0` if successful.
type AnswerCommand struct {
}

func (cmd AnswerCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd AnswerCommand) CommandString() string {
	return "ANSWER"
}

func (c *Client) Answer() Response {
	return c.Handler.Command(AnswerCommand{})
}

// AsyncAGIBreakCommand Interrupts Async AGI
//
// Interrupts expected flow of Async AGI commands and returns control to previous source (typically, the PBX dialplan).
type AsyncAGIBreakCommand struct {
}

func (cmd AsyncAGIBreakCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd AsyncAGIBreakCommand) CommandString() string {
	return "ASYNCAGI BREAK"
}

func (c *Client) AsyncAGIBreak() Response {
	return c.Handler.Command(AsyncAGIBreakCommand{})
}

// ChannelStatusCommand Returns status of the connected channel.
//
// Returns the status of the specified  channelname . If no channel name is given then returns the status of the current channel.
//
// Return values:
type ChannelStatusCommand struct {
	ChannelName *string
}

func (cmd ChannelStatusCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.ChannelName}
	return joinCommand(s), nil
}
func (cmd ChannelStatusCommand) CommandString() string {
	return "CHANNEL STATUS"
}

func (cmd ChannelStatusCommand) SetChannelName(v string) ChannelStatusCommand {
	cmd.ChannelName = &v
	return cmd
}

func (c *Client) ChannelStatus() Response {
	return c.Handler.Command(ChannelStatusCommand{})
}

// ControlStreamFileCommand Sends audio file on channel and allows the listener to control the stream.
//
// Send the given file, allowing playback to be controlled by the given digits, if any. Use double quotes for the digits if you wish none to be permitted. If offsetms is provided then the audio will seek to offsetms before play starts. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed, or `-1` on error or if the channel was disconnected. Returns the position where playback was terminated as endpos.
//
// It sets the following channel variables upon completion:
type ControlStreamFileCommand struct {
	// FileName The file extension must not be included in the filename.
	FileName     string
	EscapeDigits string
	SkipMS       *int
	Ffchar       *string // default to #
	Rewchr       *string // default to *
	Pausechr     *string
	// OffsetMS Offset, in milliseconds, to start the audio playback
	OffsetMS *int
}

func (cmd ControlStreamFileCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.FileName, cmd.EscapeDigits, cmd.SkipMS, cmd.Ffchar, cmd.Rewchr, cmd.Pausechr, cmd.OffsetMS}
	return joinCommand(s), nil
}
func (cmd ControlStreamFileCommand) CommandString() string {
	return "CONTROL STREAM FILE"
}

func (cmd ControlStreamFileCommand) SetSkipMS(v int) ControlStreamFileCommand {
	cmd.SkipMS = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetFfchar(v string) ControlStreamFileCommand {
	cmd.Ffchar = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetRewchr(v string) ControlStreamFileCommand {
	cmd.Rewchr = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetPausechr(v string) ControlStreamFileCommand {
	cmd.Pausechr = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetOffsetMS(v int) ControlStreamFileCommand {
	cmd.OffsetMS = &v
	return cmd
}

func (c *Client) ControlStreamFile(fileName string, escapeDigits string) Response {
	return c.Handler.Command(ControlStreamFileCommand{
		FileName:     fileName,
		EscapeDigits: escapeDigits,
	})
}

// DatabaseDelCommand Removes database key/value
//
// Deletes an entry in the Asterisk database for a given  family  and  key .
//
// Returns `1` if successful, `0` otherwise.
type DatabaseDelCommand struct {
	Family string
	Key    string
}

func (cmd DatabaseDelCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.Key}
	return joinCommand(s), nil
}
func (cmd DatabaseDelCommand) CommandString() string {
	return "DATABASE DEL"
}

func (c *Client) DatabaseDel(family string, key string) Response {
	return c.Handler.Command(DatabaseDelCommand{
		Family: family,
		Key:    key,
	})
}

// DatabaseDelTreeCommand Removes database keytree/value
//
// Deletes a  family  or specific  keytree  within a  family  in the Asterisk database.
//
// Returns `1` if successful, `0` otherwise.
type DatabaseDelTreeCommand struct {
	Family  string
	KeyTree *string
}

func (cmd DatabaseDelTreeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.KeyTree}
	return joinCommand(s), nil
}
func (cmd DatabaseDelTreeCommand) CommandString() string {
	return "DATABASE DELTREE"
}

func (cmd DatabaseDelTreeCommand) SetKeyTree(v string) DatabaseDelTreeCommand {
	cmd.KeyTree = &v
	return cmd
}

func (c *Client) DatabaseDelTree(family string) Response {
	return c.Handler.Command(DatabaseDelTreeCommand{
		Family: family,
	})
}

// DatabaseGetCommand Gets database value
//
// Retrieves an entry in the Asterisk database for a given  family  and  key .
//
// Returns `0` if  key  is not set. Returns `1` if  key  is set and returns the variable in parenthesis.
//
// Example return code: 200 result=1 (testvariable)
type DatabaseGetCommand struct {
	Family string
	Key    string
}

func (cmd DatabaseGetCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.Key}
	return joinCommand(s), nil
}
func (cmd DatabaseGetCommand) CommandString() string {
	return "DATABASE GET"
}

func (c *Client) DatabaseGet(family string, key string) Response {
	return c.Handler.Command(DatabaseGetCommand{
		Family: family,
		Key:    key,
	})
}

// DatabasePutCommand Adds/updates database value
//
// Adds or updates an entry in the Asterisk database for a given  family ,  key , and  value .
//
// Returns `1` if successful, `0` otherwise.
type DatabasePutCommand struct {
	Family string
	Key    string
	Value  string
}

func (cmd DatabasePutCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.Key, cmd.Value}
	return joinCommand(s), nil
}
func (cmd DatabasePutCommand) CommandString() string {
	return "DATABASE PUT"
}

func (c *Client) DatabasePut(family string, key string, value string) Response {
	return c.Handler.Command(DatabasePutCommand{
		Family: family,
		Key:    key,
		Value:  value,
	})
}

// ExecCommand Executes a given Application
//
// Executes  application  with given  options .
//
// Returns whatever the  application  returns, or `-2` on failure to find  application .
type ExecCommand struct {
	Application string
	Options     string
}

func (cmd ExecCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Application, cmd.Options}
	return joinCommand(s), nil
}
func (cmd ExecCommand) CommandString() string {
	return "EXEC"
}

func (c *Client) Exec(application string, options string) Response {
	return c.Handler.Command(ExecCommand{
		Application: application,
		Options:     options,
	})
}

// GetDataCommand Prompts for DTMF on a channel
//
// Stream the given  file , and receive DTMF data.
//
// Returns the digits received from the channel at the other end.
type GetDataCommand struct {
	File      string
	Timeout   *int
	Maxdigits *string
}

func (cmd GetDataCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.File, cmd.Timeout, cmd.Maxdigits}
	return joinCommand(s), nil
}
func (cmd GetDataCommand) CommandString() string {
	return "GET DATA"
}

func (cmd GetDataCommand) SetTimeout(v int) GetDataCommand {
	cmd.Timeout = &v
	return cmd
}
func (cmd GetDataCommand) SetMaxdigits(v string) GetDataCommand {
	cmd.Maxdigits = &v
	return cmd
}

func (c *Client) GetData(file string) Response {
	return c.Handler.Command(GetDataCommand{
		File: file,
	})
}

// GetFullVariableCommand Evaluates a channel expression
//
// Evaluates the given  expression  against the channel specified by  channelname , or the current channel if  channelname  is not provided.
//
// Unlike GET VARIABLE, the  expression  is processed in a manner similar to dialplan evaluation, allowing complex and built-in variables to be accessed, e.g. `The time is ${EPOCH} `
//
// Returns `0` if no channel matching  channelname  exists, `1` otherwise.
//
// Example return code: 200 result=1 (The time is 1578493800)
type GetFullVariableCommand struct {
	Expression  string
	ChannelName *string
}

func (cmd GetFullVariableCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Expression, cmd.ChannelName}
	return joinCommand(s), nil
}
func (cmd GetFullVariableCommand) CommandString() string {
	return "GET FULL VARIABLE"
}

func (cmd GetFullVariableCommand) SetChannelName(v string) GetFullVariableCommand {
	cmd.ChannelName = &v
	return cmd
}

func (c *Client) GetFullVariable(expression string) Response {
	return c.Handler.Command(GetFullVariableCommand{
		Expression: expression,
	})
}

// GetOptionCommand Stream file, prompt for DTMF, with timeout.
//
// Behaves similar to STREAM FILE but used with a timeout option.
type GetOptionCommand struct {
	FileName     string
	EscapeDigits string
	Timeout      *int
}

func (cmd GetOptionCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.FileName, cmd.EscapeDigits, cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd GetOptionCommand) CommandString() string {
	return "GET OPTION"
}

func (cmd GetOptionCommand) SetTimeout(v int) GetOptionCommand {
	cmd.Timeout = &v
	return cmd
}

func (c *Client) GetOption(fileName string, escapeDigits string) Response {
	return c.Handler.Command(GetOptionCommand{
		FileName:     fileName,
		EscapeDigits: escapeDigits,
	})
}

// GetVariableCommand Gets a channel variable.
//
// Returns `0` if  variablename  is not set. Returns `1` if  variablename  is set and returns the variable in parentheses.
//
// Example return code: 200 result=1 (testvariable)
type GetVariableCommand struct {
	VariableName string
}

func (cmd GetVariableCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.VariableName}
	return joinCommand(s), nil
}
func (cmd GetVariableCommand) CommandString() string {
	return "GET VARIABLE"
}

func (c *Client) GetVariable(variableName string) Response {
	return c.Handler.Command(GetVariableCommand{
		VariableName: variableName,
	})
}

// HangupCommand Hangup a channel.
//
// Hangs up the specified channel. If no channel name is given, hangs up the current channel
type HangupCommand struct {
	ChannelName *string
}

func (cmd HangupCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.ChannelName}
	return joinCommand(s), nil
}
func (cmd HangupCommand) CommandString() string {
	return "HANGUP"
}

func (cmd HangupCommand) SetChannelName(v string) HangupCommand {
	cmd.ChannelName = &v
	return cmd
}

func (c *Client) Hangup() Response {
	return c.Handler.Command(HangupCommand{})
}

// NoopCommand Does nothing.
//
// Does nothing.
type NoopCommand struct {
}

func (cmd NoopCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd NoopCommand) CommandString() string {
	return "NOOP"
}

func (c *Client) Noop() Response {
	return c.Handler.Command(NoopCommand{})
}

// ReceiveCharCommand Receives one character from channels supporting it.
//
// Receives a character of text on a channel. Most channels do not support the reception of text. Returns the decimal value of the character if one is received, or `0` if the channel does not support text reception. Returns `-1` only on error/hangup.
type ReceiveCharCommand struct {
	// Timeout The maximum time to wait for input in milliseconds, or `0` for infinite. Most channels
	Timeout int
}

func (cmd ReceiveCharCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd ReceiveCharCommand) CommandString() string {
	return "RECEIVE CHAR"
}

func (c *Client) ReceiveChar(timeout int) Response {
	return c.Handler.Command(ReceiveCharCommand{
		Timeout: timeout,
	})
}

// ReceiveTextCommand Receives text from channels supporting it.
//
// Receives a string of text on a channel. Most channels do not support the reception of text. Returns `-1` for failure or `1` for success, and the string in parenthesis.
type ReceiveTextCommand struct {
	// Timeout The timeout to be the maximum time to wait for input in milliseconds, or `0` for infinite.
	Timeout int
}

func (cmd ReceiveTextCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd ReceiveTextCommand) CommandString() string {
	return "RECEIVE TEXT"
}

func (c *Client) ReceiveText(timeout int) Response {
	return c.Handler.Command(ReceiveTextCommand{
		Timeout: timeout,
	})
}

// RecordFileCommand Records to a given file.
//
// Record to a file until a given dtmf digit in the sequence is received. Returns `-1` on hangup or error. The format will specify what kind of file will be recorded. The  timeout  is the maximum record time in milliseconds, or `-1` for no  timeout .  offset samples  is optional, and, if provided, will seek to the offset without exceeding the end of the file.  beep  can take any value, and causes Asterisk to play a beep to the channel that is about to be recorded.  silence  is the number of seconds of silence allowed before the function returns despite the lack of dtmf digits or reaching  timeout .  silence  value must be preceded by `s=` and is also optional.
type RecordFileCommand struct {
	// FileName The destination filename of the recorded audio.
	FileName string
	// Format The audio format in which to save the resulting file.
	Format string
	// EscapeDigits The DTMF digits that will terminate the recording process.
	EscapeDigits string
	// Timeout The maximum recording time in milliseconds. Set to -1 for no limit.
	Timeout int
	// OffsetSamples Causes the recording to first seek to the specified offset before recording begins.
	OffsetSamples *string
	// Beep Causes Asterisk to play a beep as recording begins. This argument can take any value.
	Beep *string
	// SSilence The number of seconds of silence that are permitted before the recording is terminated, regardless of the  escape_digits  or  timeout  arguments. If specified, this parameter must be preceded by `s=`.
	SSilence *string
}

func (cmd RecordFileCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.FileName, cmd.Format, cmd.EscapeDigits, cmd.Timeout, cmd.OffsetSamples, cmd.Beep, cmd.SSilence}
	return joinCommand(s), nil
}
func (cmd RecordFileCommand) CommandString() string {
	return "RECORD FILE"
}

func (cmd RecordFileCommand) SetOffsetSamples(v string) RecordFileCommand {
	cmd.OffsetSamples = &v
	return cmd
}
func (cmd RecordFileCommand) SetBeep(v string) RecordFileCommand {
	cmd.Beep = &v
	return cmd
}
func (cmd RecordFileCommand) SetSSilence(v string) RecordFileCommand {
	cmd.SSilence = &v
	return cmd
}

func (c *Client) RecordFile(fileName string, format string, escapeDigits string, timeout int) Response {
	return c.Handler.Command(RecordFileCommand{
		FileName:     fileName,
		Format:       format,
		EscapeDigits: escapeDigits,
		Timeout:      timeout,
	})
}

// SayAlphaCommand Says a given character string.
//
// Say a given character string, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed or `-1` on error/hangup.
type SayAlphaCommand struct {
	Number       string
	EscapeDigits string
}

func (cmd SayAlphaCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayAlphaCommand) CommandString() string {
	return "SAY ALPHA"
}

func (c *Client) SayAlpha(number string, escapeDigits string) Response {
	return c.Handler.Command(SayAlphaCommand{
		Number:       number,
		EscapeDigits: escapeDigits,
	})
}

// SayDigitsCommand Says a given digit string.
//
// Say a given digit string, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed or `-1` on error/hangup.
type SayDigitsCommand struct {
	Number       string
	EscapeDigits string
}

func (cmd SayDigitsCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayDigitsCommand) CommandString() string {
	return "SAY DIGITS"
}

func (c *Client) SayDigits(number string, escapeDigits string) Response {
	return c.Handler.Command(SayDigitsCommand{
		Number:       number,
		EscapeDigits: escapeDigits,
	})
}

// SayNumberCommand Says a given number.
//
// Say a given number, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed or `-1` on error/hangup.
type SayNumberCommand struct {
	Number       string
	EscapeDigits string
	Gender       *string
}

func (cmd SayNumberCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number, cmd.EscapeDigits, cmd.Gender}
	return joinCommand(s), nil
}
func (cmd SayNumberCommand) CommandString() string {
	return "SAY NUMBER"
}

func (cmd SayNumberCommand) SetGender(v string) SayNumberCommand {
	cmd.Gender = &v
	return cmd
}

func (c *Client) SayNumber(number string, escapeDigits string) Response {
	return c.Handler.Command(SayNumberCommand{
		Number:       number,
		EscapeDigits: escapeDigits,
	})
}

// SayPhoneticCommand Says a given character string with phonetics.
//
// Say a given character string with phonetics, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit pressed, the ASCII numerical value of the digit if one was pressed, or `-1` on error/hangup.
type SayPhoneticCommand struct {
	String       string
	EscapeDigits string
}

func (cmd SayPhoneticCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.String, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayPhoneticCommand) CommandString() string {
	return "SAY PHONETIC"
}

func (c *Client) SayPhonetic(string string, escapeDigits string) Response {
	return c.Handler.Command(SayPhoneticCommand{
		String:       string,
		EscapeDigits: escapeDigits,
	})
}

// SayDateCommand Says a given date.
//
// Say a given date, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed or `-1` on error/hangup.
type SayDateCommand struct {
	// Date Is number of seconds elapsed since 00:00:00 on January 1, 1970. Coordinated Universal Time (UTC).
	Date         string
	EscapeDigits string
}

func (cmd SayDateCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Date, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayDateCommand) CommandString() string {
	return "SAY DATE"
}

func (c *Client) SayDate(date string, escapeDigits string) Response {
	return c.Handler.Command(SayDateCommand{
		Date:         date,
		EscapeDigits: escapeDigits,
	})
}

// SayTimeCommand Says a given time.
//
// Say a given time, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed or `-1` on error/hangup.
type SayTimeCommand struct {
	// Time Is number of seconds elapsed since 00:00:00 on January 1, 1970. Coordinated Universal Time (UTC).
	Time         float64
	EscapeDigits string
}

func (cmd SayTimeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Time, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayTimeCommand) CommandString() string {
	return "SAY TIME"
}

func (c *Client) SayTime(time float64, escapeDigits string) Response {
	return c.Handler.Command(SayTimeCommand{
		Time:         time,
		EscapeDigits: escapeDigits,
	})
}

// SayDatetimeCommand Says a given time as specified by the format given.
//
// Say a given time, returning early if any of the given DTMF digits are received on the channel. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed or `-1` on error/hangup.
type SayDatetimeCommand struct {
	// Time Is number of seconds elapsed since 00:00:00 on January 1, 1970, Coordinated Universal Time (UTC)
	Time         float64
	EscapeDigits string
	// Format Is the format the time should be said in. See  voicemail.conf  (defaults to `ABdY 'digits/at' IMp`).
	Format *string
	// Timezone Acceptable values can be found in  /usr/share/zoneinfo  Defaults to machine default.
	Timezone *string
}

func (cmd SayDatetimeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Time, cmd.EscapeDigits, cmd.Format, cmd.Timezone}
	return joinCommand(s), nil
}
func (cmd SayDatetimeCommand) CommandString() string {
	return "SAY DATETIME"
}

func (cmd SayDatetimeCommand) SetFormat(v string) SayDatetimeCommand {
	cmd.Format = &v
	return cmd
}
func (cmd SayDatetimeCommand) SetTimezone(v string) SayDatetimeCommand {
	cmd.Timezone = &v
	return cmd
}

func (c *Client) SayDatetime(time float64, escapeDigits string) Response {
	return c.Handler.Command(SayDatetimeCommand{
		Time:         time,
		EscapeDigits: escapeDigits,
	})
}

// SendImageCommand Sends images to channels supporting it.
//
// Sends the given image on a channel. Most channels do not support the transmission of images. Returns `0` if image is sent, or if the channel does not support image transmission. Returns `-1` only on error/hangup. Image names should not include extensions.
type SendImageCommand struct {
	Image string
}

func (cmd SendImageCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Image}
	return joinCommand(s), nil
}
func (cmd SendImageCommand) CommandString() string {
	return "SEND IMAGE"
}

func (c *Client) SendImage(image string) Response {
	return c.Handler.Command(SendImageCommand{
		Image: image,
	})
}

// SendTextCommand Sends text to channels supporting it.
//
// Sends the given text on a channel. Most channels do not support the transmission of text. Returns `0` if text is sent, or if the channel does not support text transmission. Returns `-1` only on error/hangup.
type SendTextCommand struct {
	// TextToSend Text consisting of greater than one word should be placed in quotes since the command only accepts a single argument.
	TextToSend string
}

func (cmd SendTextCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.TextToSend}
	return joinCommand(s), nil
}
func (cmd SendTextCommand) CommandString() string {
	return "SEND TEXT"
}

func (c *Client) SendText(textToSend string) Response {
	return c.Handler.Command(SendTextCommand{
		TextToSend: textToSend,
	})
}

// SetAutoHangupCommand Autohangup channel in some time.
//
// Cause the channel to automatically hangup at  time  seconds in the future. Of course it can be hungup before then as well. Setting to `0` will cause the autohangup feature to be disabled on this channel.
type SetAutoHangupCommand struct {
	Time float64
}

func (cmd SetAutoHangupCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Time}
	return joinCommand(s), nil
}
func (cmd SetAutoHangupCommand) CommandString() string {
	return "SET AUTOHANGUP"
}

func (c *Client) SetAutoHangup(time float64) Response {
	return c.Handler.Command(SetAutoHangupCommand{
		Time: time,
	})
}

// SetCallerIDCommand Sets callerid for the current channel.
//
// Changes the callerid of the current channel.
type SetCallerIDCommand struct {
	Number string
}

func (cmd SetCallerIDCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number}
	return joinCommand(s), nil
}
func (cmd SetCallerIDCommand) CommandString() string {
	return "SET CALLERID"
}

func (c *Client) SetCallerID(number string) Response {
	return c.Handler.Command(SetCallerIDCommand{
		Number: number,
	})
}

// SetContextCommand Sets channel context.
//
// Sets the context for continuation upon exiting the application.
type SetContextCommand struct {
	DesiredContext string
}

func (cmd SetContextCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.DesiredContext}
	return joinCommand(s), nil
}
func (cmd SetContextCommand) CommandString() string {
	return "SET CONTEXT"
}

func (c *Client) SetContext(desiredContext string) Response {
	return c.Handler.Command(SetContextCommand{
		DesiredContext: desiredContext,
	})
}

// SetExtensionCommand Changes channel extension.
//
// Changes the extension for continuation upon exiting the application.
type SetExtensionCommand struct {
	NewExtension string
}

func (cmd SetExtensionCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.NewExtension}
	return joinCommand(s), nil
}
func (cmd SetExtensionCommand) CommandString() string {
	return "SET EXTENSION"
}

func (c *Client) SetExtension(newExtension string) Response {
	return c.Handler.Command(SetExtensionCommand{
		NewExtension: newExtension,
	})
}

// SetMusicCommand Enable/Disable Music on hold generator
//
// Enables/Disables the music on hold generator. If  class  is not specified, then the `default` music on hold class will be used. This generator will be stopped automatically when playing a file.
//
// Always returns `0`.
type SetMusicCommand struct {
	Class string
	// has missing params
}

func (cmd SetMusicCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Class}
	return joinCommand(s), nil
}
func (cmd SetMusicCommand) CommandString() string {
	return "SET MUSIC"
}

func (c *Client) SetMusic(class string) Response {
	return c.Handler.Command(SetMusicCommand{
		Class: class,
	})
}

// SetPriorityCommand Set channel dialplan priority.
//
// Changes the priority for continuation upon exiting the application. The priority must be a valid priority or label.
type SetPriorityCommand struct {
	Priority int
}

func (cmd SetPriorityCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Priority}
	return joinCommand(s), nil
}
func (cmd SetPriorityCommand) CommandString() string {
	return "SET PRIORITY"
}

func (c *Client) SetPriority(priority int) Response {
	return c.Handler.Command(SetPriorityCommand{
		Priority: priority,
	})
}

// SetVariableCommand Sets a channel variable.
//
// Sets a variable to the current channel.
type SetVariableCommand struct {
	VariableName string
	Value        string
}

func (cmd SetVariableCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.VariableName, cmd.Value}
	return joinCommand(s), nil
}
func (cmd SetVariableCommand) CommandString() string {
	return "SET VARIABLE"
}

func (c *Client) SetVariable(variableName string, value string) Response {
	return c.Handler.Command(SetVariableCommand{
		VariableName: variableName,
		Value:        value,
	})
}

// StreamFileCommand Sends audio file on channel.
//
// Send the given file, allowing playback to be interrupted by the given digits, if any. Returns `0` if playback completes without a digit being pressed, or the ASCII numerical value of the digit if one was pressed, or `-1` on error or if the channel was disconnected. If musiconhold is playing before calling stream file it will be automatically stopped and will not be restarted after completion.
//
// It sets the following channel variables upon completion:
type StreamFileCommand struct {
	// FileName File name to play. The file extension must not be included in the  filename .
	FileName string
	// EscapeDigits Use double quotes for the digits if you wish none to be permitted.
	EscapeDigits string
	// SampleOffset If sample offset is provided then the audio will seek to sample offset before play starts.
	SampleOffset *int
}

func (cmd StreamFileCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.FileName, cmd.EscapeDigits, cmd.SampleOffset}
	return joinCommand(s), nil
}
func (cmd StreamFileCommand) CommandString() string {
	return "STREAM FILE"
}

func (cmd StreamFileCommand) SetSampleOffset(v int) StreamFileCommand {
	cmd.SampleOffset = &v
	return cmd
}

func (c *Client) StreamFile(fileName string, escapeDigits string) Response {
	return c.Handler.Command(StreamFileCommand{
		FileName:     fileName,
		EscapeDigits: escapeDigits,
	})
}

// TddModeCommand Toggles TDD mode (for the deaf).
//
// Enable/Disable TDD transmission/reception on a channel. Returns `1` if successful, or `0` if channel is not TDD-capable.
type TddModeCommand struct {
	Boolean string
}

func (cmd TddModeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Boolean}
	return joinCommand(s), nil
}
func (cmd TddModeCommand) CommandString() string {
	return "TDD MODE"
}

func (c *Client) TddMode(boolean string) Response {
	return c.Handler.Command(TddModeCommand{
		Boolean: boolean,
	})
}

// VerboseCommand Logs a message to the asterisk verbose log.
//
// Sends  message  to the console via verbose message system.  level  is the verbose level (1-4). Always returns `1`
type VerboseCommand struct {
	Message string
	Level   string
}

func (cmd VerboseCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Message, cmd.Level}
	return joinCommand(s), nil
}
func (cmd VerboseCommand) CommandString() string {
	return "VERBOSE"
}

func (c *Client) Verbose(message string, level string) Response {
	return c.Handler.Command(VerboseCommand{
		Message: message,
		Level:   level,
	})
}

// WaitForDigitCommand Waits for a digit to be pressed.
//
// Waits up to  timeout  milliseconds for channel to receive a DTMF digit. Returns `-1` on channel failure, `0` if no digit is received in the timeout, or the numerical value of the ascii of the digit if one is received. Use `-1` for the  timeout  value if you desire the call to block indefinitely.
type WaitForDigitCommand struct {
	Timeout int
}

func (cmd WaitForDigitCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd WaitForDigitCommand) CommandString() string {
	return "WAIT FOR DIGIT"
}

func (c *Client) WaitForDigit(timeout int) Response {
	return c.Handler.Command(WaitForDigitCommand{
		Timeout: timeout,
	})
}

// SpeechCreateCommand Creates a speech object.
//
// Create a speech object to be used by the other Speech AGI commands.
type SpeechCreateCommand struct {
	Engine string
}

func (cmd SpeechCreateCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Engine}
	return joinCommand(s), nil
}
func (cmd SpeechCreateCommand) CommandString() string {
	return "SPEECH CREATE"
}

func (c *Client) SpeechCreate(engine string) Response {
	return c.Handler.Command(SpeechCreateCommand{
		Engine: engine,
	})
}

// SpeechSetCommand Sets a speech engine setting.
//
// Set an engine-specific setting.
type SpeechSetCommand struct {
	Name  string
	Value string
}

func (cmd SpeechSetCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Name, cmd.Value}
	return joinCommand(s), nil
}
func (cmd SpeechSetCommand) CommandString() string {
	return "SPEECH SET"
}

func (c *Client) SpeechSet(name string, value string) Response {
	return c.Handler.Command(SpeechSetCommand{
		Name:  name,
		Value: value,
	})
}

// SpeechDestroyCommand Destroys a speech object.
//
// Destroy the speech object created by `SPEECH CREATE`.
type SpeechDestroyCommand struct {
}

func (cmd SpeechDestroyCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd SpeechDestroyCommand) CommandString() string {
	return "SPEECH DESTROY"
}

func (c *Client) SpeechDestroy() Response {
	return c.Handler.Command(SpeechDestroyCommand{})
}

// SpeechLoadGrammarCommand Loads a grammar.
//
// Loads the specified grammar as the specified name.
type SpeechLoadGrammarCommand struct {
	GrammarName   string
	PathToGrammar string
}

func (cmd SpeechLoadGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName, cmd.PathToGrammar}
	return joinCommand(s), nil
}
func (cmd SpeechLoadGrammarCommand) CommandString() string {
	return "SPEECH LOAD GRAMMAR"
}

func (c *Client) SpeechLoadGrammar(grammarName string, pathToGrammar string) Response {
	return c.Handler.Command(SpeechLoadGrammarCommand{
		GrammarName:   grammarName,
		PathToGrammar: pathToGrammar,
	})
}

// SpeechUnloadGrammarCommand Unloads a grammar.
//
// Unloads the specified grammar.
type SpeechUnloadGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechUnloadGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName}
	return joinCommand(s), nil
}
func (cmd SpeechUnloadGrammarCommand) CommandString() string {
	return "SPEECH UNLOAD GRAMMAR"
}

func (c *Client) SpeechUnloadGrammar(grammarName string) Response {
	return c.Handler.Command(SpeechUnloadGrammarCommand{
		GrammarName: grammarName,
	})
}

// SpeechActivateGrammarCommand Activates a grammar.
//
// Activates the specified grammar on the speech object.
type SpeechActivateGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechActivateGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName}
	return joinCommand(s), nil
}
func (cmd SpeechActivateGrammarCommand) CommandString() string {
	return "SPEECH ACTIVATE GRAMMAR"
}

func (c *Client) SpeechActivateGrammar(grammarName string) Response {
	return c.Handler.Command(SpeechActivateGrammarCommand{
		GrammarName: grammarName,
	})
}

// SpeechDeactivateGrammarCommand Deactivates a grammar.
//
// Deactivates the specified grammar on the speech object.
type SpeechDeactivateGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechDeactivateGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName}
	return joinCommand(s), nil
}
func (cmd SpeechDeactivateGrammarCommand) CommandString() string {
	return "SPEECH DEACTIVATE GRAMMAR"
}

func (c *Client) SpeechDeactivateGrammar(grammarName string) Response {
	return c.Handler.Command(SpeechDeactivateGrammarCommand{
		GrammarName: grammarName,
	})
}

// SpeechRecognizeCommand Recognizes speech.
//
// Plays back given  prompt  while listening for speech and dtmf.
type SpeechRecognizeCommand struct {
	Prompt  string
	Timeout int
	Offset  *string
}

func (cmd SpeechRecognizeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Prompt, cmd.Timeout, cmd.Offset}
	return joinCommand(s), nil
}
func (cmd SpeechRecognizeCommand) CommandString() string {
	return "SPEECH RECOGNIZE"
}

func (cmd SpeechRecognizeCommand) SetOffset(v string) SpeechRecognizeCommand {
	cmd.Offset = &v
	return cmd
}

func (c *Client) SpeechRecognize(prompt string, timeout int) Response {
	return c.Handler.Command(SpeechRecognizeCommand{
		Prompt:  prompt,
		Timeout: timeout,
	})
}

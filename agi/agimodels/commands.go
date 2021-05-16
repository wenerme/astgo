package agimodels

// GosubCommand Cause the channel to execute the specified dialplan subroutine.
type GosubCommand struct {
	Context          string // required
	Extension        string // required
	Priority         int    // required
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

// AnswerCommand Answer channel
type AnswerCommand struct {
}

func (cmd AnswerCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd AnswerCommand) CommandString() string {
	return "ANSWER"
}

// AsyncagiBreakCommand Interrupts Async AGI
type AsyncagiBreakCommand struct {
}

func (cmd AsyncagiBreakCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd AsyncagiBreakCommand) CommandString() string {
	return "ASYNCAGI BREAK"
}

// ChannelStatusCommand Returns status of the connected channel.
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

// ControlStreamFileCommand Sends audio file on channel and allows the listener to control the stream.
type ControlStreamFileCommand struct {
	FileName     string // required
	EscapeDigits string // required
	SkipMS       *int
	Ffchar       *string
	Rewchr       *string
	Pausechr     *string
	OffsetMS     *int
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

// DatabaseDelCommand Removes database key/value
type DatabaseDelCommand struct {
	Family string // required
	Key    string // required

}

func (cmd DatabaseDelCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.Key}
	return joinCommand(s), nil
}
func (cmd DatabaseDelCommand) CommandString() string {
	return "DATABASE DEL"
}

// DatabaseDeltreeCommand Removes database keytree/value
type DatabaseDeltreeCommand struct {
	Family  string // required
	KeyTree *string
}

func (cmd DatabaseDeltreeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.KeyTree}
	return joinCommand(s), nil
}
func (cmd DatabaseDeltreeCommand) CommandString() string {
	return "DATABASE DELTREE"
}

func (cmd DatabaseDeltreeCommand) SetKeyTree(v string) DatabaseDeltreeCommand {
	cmd.KeyTree = &v
	return cmd
}

// DatabaseGetCommand Gets database value
type DatabaseGetCommand struct {
	Family string // required
	Key    string // required

}

func (cmd DatabaseGetCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.Key}
	return joinCommand(s), nil
}
func (cmd DatabaseGetCommand) CommandString() string {
	return "DATABASE GET"
}

// DatabasePutCommand Adds/updates database value
type DatabasePutCommand struct {
	Family string // required
	Key    string // required
	Value  string // required

}

func (cmd DatabasePutCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Family, cmd.Key, cmd.Value}
	return joinCommand(s), nil
}
func (cmd DatabasePutCommand) CommandString() string {
	return "DATABASE PUT"
}

// ExecCommand Executes a given Application
type ExecCommand struct {
	Application string // required
	Options     string // required

}

func (cmd ExecCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Application, cmd.Options}
	return joinCommand(s), nil
}
func (cmd ExecCommand) CommandString() string {
	return "EXEC"
}

// GetDataCommand Prompts for DTMF on a channel
type GetDataCommand struct {
	File      string // required
	Timeout   *string
	Maxdigits *string
}

func (cmd GetDataCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.File, cmd.Timeout, cmd.Maxdigits}
	return joinCommand(s), nil
}
func (cmd GetDataCommand) CommandString() string {
	return "GET DATA"
}

func (cmd GetDataCommand) SetTimeout(v string) GetDataCommand {
	cmd.Timeout = &v
	return cmd
}
func (cmd GetDataCommand) SetMaxdigits(v string) GetDataCommand {
	cmd.Maxdigits = &v
	return cmd
}

// GetFullVariableCommand Evaluates a channel expression
type GetFullVariableCommand struct {
	Expression  string // required
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

// GetOptionCommand Stream file, prompt for DTMF, with timeout.
type GetOptionCommand struct {
	FileName     string // required
	EscapeDigits string // required
	Timeout      *string
}

func (cmd GetOptionCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.FileName, cmd.EscapeDigits, cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd GetOptionCommand) CommandString() string {
	return "GET OPTION"
}

func (cmd GetOptionCommand) SetTimeout(v string) GetOptionCommand {
	cmd.Timeout = &v
	return cmd
}

// GetVariableCommand Gets a channel variable.
type GetVariableCommand struct {
	VariableName string // required

}

func (cmd GetVariableCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.VariableName}
	return joinCommand(s), nil
}
func (cmd GetVariableCommand) CommandString() string {
	return "GET VARIABLE"
}

// HangupCommand Hangup a channel.
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

// NoopCommand Does nothing.
type NoopCommand struct {
}

func (cmd NoopCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd NoopCommand) CommandString() string {
	return "NOOP"
}

// ReceiveCharCommand Receives one character from channels supporting it.
type ReceiveCharCommand struct {
	Timeout string // required

}

func (cmd ReceiveCharCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd ReceiveCharCommand) CommandString() string {
	return "RECEIVE CHAR"
}

// ReceiveTextCommand Receives text from channels supporting it.
type ReceiveTextCommand struct {
	Timeout string // required

}

func (cmd ReceiveTextCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd ReceiveTextCommand) CommandString() string {
	return "RECEIVE TEXT"
}

// RecordFileCommand Records to a given file.
type RecordFileCommand struct {
	FileName      string // required
	Format        string // required
	EscapeDigits  string // required
	Timeout       string // required
	OffsetSamples *string
	Beep          *string
	SSilence      *string
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

// SayAlphaCommand Says a given character string.
type SayAlphaCommand struct {
	Number       string // required
	EscapeDigits string // required

}

func (cmd SayAlphaCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayAlphaCommand) CommandString() string {
	return "SAY ALPHA"
}

// SayDigitsCommand Says a given digit string.
type SayDigitsCommand struct {
	Number       string // required
	EscapeDigits string // required

}

func (cmd SayDigitsCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayDigitsCommand) CommandString() string {
	return "SAY DIGITS"
}

// SayNumberCommand Says a given number.
type SayNumberCommand struct {
	Number       string // required
	EscapeDigits string // required
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

// SayPhoneticCommand Says a given character string with phonetics.
type SayPhoneticCommand struct {
	String       string // required
	EscapeDigits string // required

}

func (cmd SayPhoneticCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.String, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayPhoneticCommand) CommandString() string {
	return "SAY PHONETIC"
}

// SayDateCommand Says a given date.
type SayDateCommand struct {
	Date         string // required
	EscapeDigits string // required

}

func (cmd SayDateCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Date, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayDateCommand) CommandString() string {
	return "SAY DATE"
}

// SayTimeCommand Says a given time.
type SayTimeCommand struct {
	Time         string // required
	EscapeDigits string // required

}

func (cmd SayTimeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Time, cmd.EscapeDigits}
	return joinCommand(s), nil
}
func (cmd SayTimeCommand) CommandString() string {
	return "SAY TIME"
}

// SayDatetimeCommand Says a given time as specified by the format given.
type SayDatetimeCommand struct {
	Time         string // required
	EscapeDigits string // required
	Format       *string
	Timezone     *string
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

// SendImageCommand Sends images to channels supporting it.
type SendImageCommand struct {
	Image string // required

}

func (cmd SendImageCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Image}
	return joinCommand(s), nil
}
func (cmd SendImageCommand) CommandString() string {
	return "SEND IMAGE"
}

// SendTextCommand Sends text to channels supporting it.
type SendTextCommand struct {
	TextToSend string // required

}

func (cmd SendTextCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.TextToSend}
	return joinCommand(s), nil
}
func (cmd SendTextCommand) CommandString() string {
	return "SEND TEXT"
}

// SetAutohangupCommand Autohangup channel in some time.
type SetAutohangupCommand struct {
	Time string // required

}

func (cmd SetAutohangupCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Time}
	return joinCommand(s), nil
}
func (cmd SetAutohangupCommand) CommandString() string {
	return "SET AUTOHANGUP"
}

// SetCalleridCommand Sets callerid for the current channel.
type SetCalleridCommand struct {
	Number string // required

}

func (cmd SetCalleridCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Number}
	return joinCommand(s), nil
}
func (cmd SetCalleridCommand) CommandString() string {
	return "SET CALLERID"
}

// SetContextCommand Sets channel context.
type SetContextCommand struct {
	DesiredContext string // required

}

func (cmd SetContextCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.DesiredContext}
	return joinCommand(s), nil
}
func (cmd SetContextCommand) CommandString() string {
	return "SET CONTEXT"
}

// SetExtensionCommand Changes channel extension.
type SetExtensionCommand struct {
	NewExtension string // required

}

func (cmd SetExtensionCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.NewExtension}
	return joinCommand(s), nil
}
func (cmd SetExtensionCommand) CommandString() string {
	return "SET EXTENSION"
}

// SetMusicCommand Enable/Disable Music on hold generator
type SetMusicCommand struct {
	Class string // required
	// has missing params
}

func (cmd SetMusicCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Class}
	return joinCommand(s), nil
}
func (cmd SetMusicCommand) CommandString() string {
	return "SET MUSIC"
}

// SetPriorityCommand Set channel dialplan priority.
type SetPriorityCommand struct {
	Priority int // required

}

func (cmd SetPriorityCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Priority}
	return joinCommand(s), nil
}
func (cmd SetPriorityCommand) CommandString() string {
	return "SET PRIORITY"
}

// SetVariableCommand Sets a channel variable.
type SetVariableCommand struct {
	VariableName string // required
	Value        string // required

}

func (cmd SetVariableCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.VariableName, cmd.Value}
	return joinCommand(s), nil
}
func (cmd SetVariableCommand) CommandString() string {
	return "SET VARIABLE"
}

// StreamFileCommand Sends audio file on channel.
type StreamFileCommand struct {
	FileName     string // required
	EscapeDigits string // required
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

// TddModeCommand Toggles TDD mode (for the deaf).
type TddModeCommand struct {
	Boolean string // required

}

func (cmd TddModeCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Boolean}
	return joinCommand(s), nil
}
func (cmd TddModeCommand) CommandString() string {
	return "TDD MODE"
}

// VerboseCommand Logs a message to the asterisk verbose log.
type VerboseCommand struct {
	Message string // required
	Level   string // required

}

func (cmd VerboseCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Message, cmd.Level}
	return joinCommand(s), nil
}
func (cmd VerboseCommand) CommandString() string {
	return "VERBOSE"
}

// WaitForDigitCommand Waits for a digit to be pressed.
type WaitForDigitCommand struct {
	Timeout string // required

}

func (cmd WaitForDigitCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Timeout}
	return joinCommand(s), nil
}
func (cmd WaitForDigitCommand) CommandString() string {
	return "WAIT FOR DIGIT"
}

// SpeechCreateCommand Creates a speech object.
type SpeechCreateCommand struct {
	Engine string // required

}

func (cmd SpeechCreateCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Engine}
	return joinCommand(s), nil
}
func (cmd SpeechCreateCommand) CommandString() string {
	return "SPEECH CREATE"
}

// SpeechSetCommand Sets a speech engine setting.
type SpeechSetCommand struct {
	Name  string // required
	Value string // required

}

func (cmd SpeechSetCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.Name, cmd.Value}
	return joinCommand(s), nil
}
func (cmd SpeechSetCommand) CommandString() string {
	return "SPEECH SET"
}

// SpeechDestroyCommand Destroys a speech object.
type SpeechDestroyCommand struct {
}

func (cmd SpeechDestroyCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString()}
	return joinCommand(s), nil
}
func (cmd SpeechDestroyCommand) CommandString() string {
	return "SPEECH DESTROY"
}

// SpeechLoadGrammarCommand Loads a grammar.
type SpeechLoadGrammarCommand struct {
	GrammarName   string // required
	PathToGrammar string // required

}

func (cmd SpeechLoadGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName, cmd.PathToGrammar}
	return joinCommand(s), nil
}
func (cmd SpeechLoadGrammarCommand) CommandString() string {
	return "SPEECH LOAD GRAMMAR"
}

// SpeechUnloadGrammarCommand Unloads a grammar.
type SpeechUnloadGrammarCommand struct {
	GrammarName string // required

}

func (cmd SpeechUnloadGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName}
	return joinCommand(s), nil
}
func (cmd SpeechUnloadGrammarCommand) CommandString() string {
	return "SPEECH UNLOAD GRAMMAR"
}

// SpeechActivateGrammarCommand Activates a grammar.
type SpeechActivateGrammarCommand struct {
	GrammarName string // required

}

func (cmd SpeechActivateGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName}
	return joinCommand(s), nil
}
func (cmd SpeechActivateGrammarCommand) CommandString() string {
	return "SPEECH ACTIVATE GRAMMAR"
}

// SpeechDeactivateGrammarCommand Deactivates a grammar.
type SpeechDeactivateGrammarCommand struct {
	GrammarName string // required

}

func (cmd SpeechDeactivateGrammarCommand) Command() (string, error) {
	s := []interface{}{cmd.CommandString(), cmd.GrammarName}
	return joinCommand(s), nil
}
func (cmd SpeechDeactivateGrammarCommand) CommandString() string {
	return "SPEECH DEACTIVATE GRAMMAR"
}

// SpeechRecognizeCommand Recognizes speech.
type SpeechRecognizeCommand struct {
	Prompt  string // required
	Timeout string // required
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

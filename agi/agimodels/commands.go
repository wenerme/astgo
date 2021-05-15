package agimodels

import (
	"context"
	"fmt"
	"strings"
)

// GosubCommand Cause the channel to execute the specified dialplan subroutine.
type GosubCommand struct {
	Context          string
	Extension        string
	Priority         string
	OptionalArgument string
}

func (cmd GosubCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Context == "" {
		return "", fmt.Errorf("field Gosub Context is required")
	}
	s = append(s, cmd.Context)
	if cmd.Extension == "" {
		return "", fmt.Errorf("field Gosub Extension is required")
	}
	s = append(s, cmd.Extension)
	if cmd.Priority == "" {
		return "", fmt.Errorf("field Gosub Priority is required")
	}
	s = append(s, cmd.Priority)
	s = append(s, cmd.OptionalArgument)
	return strings.Join(s, " "), nil
}
func (cmd GosubCommand) CommandString() string {
	return "GOSUB"
}

// AnswerCommand Answer channel
type AnswerCommand struct {
}

func (cmd AnswerCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	return strings.Join(s, " "), nil
}
func (cmd AnswerCommand) CommandString() string {
	return "ANSWER"
}

// AsyncagiBreakCommand Interrupts Async AGI
type AsyncagiBreakCommand struct {
}

func (cmd AsyncagiBreakCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	return strings.Join(s, " "), nil
}
func (cmd AsyncagiBreakCommand) CommandString() string {
	return "ASYNCAGI BREAK"
}

// ChannelStatusCommand Returns status of the connected channel.
type ChannelStatusCommand struct {
	ChannelName string
}

func (cmd ChannelStatusCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	s = append(s, cmd.ChannelName)
	return strings.Join(s, " "), nil
}
func (cmd ChannelStatusCommand) CommandString() string {
	return "CHANNEL STATUS"
}

// ControlStreamFileCommand Sends audio file on channel and allows the listener to control the stream.
type ControlStreamFileCommand struct {
	FileName     string
	EscapeDigits string
	Skipms       string
	Ffchar       string
	Rewchr       string
	Pausechr     string
	OffsetMS     string
}

func (cmd ControlStreamFileCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.FileName == "" {
		return "", fmt.Errorf("field ControlStreamFile FileName is required")
	}
	s = append(s, cmd.FileName)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field ControlStreamFile EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	s = append(s, cmd.Skipms)
	s = append(s, cmd.Ffchar)
	s = append(s, cmd.Rewchr)
	s = append(s, cmd.Pausechr)
	s = append(s, cmd.OffsetMS)
	return strings.Join(s, " "), nil
}
func (cmd ControlStreamFileCommand) CommandString() string {
	return "CONTROL STREAM FILE"
}

// DatabaseDelCommand Removes database key/value
type DatabaseDelCommand struct {
	Family string
	Key    string
}

func (cmd DatabaseDelCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Family == "" {
		return "", fmt.Errorf("field DatabaseDel Family is required")
	}
	s = append(s, cmd.Family)
	if cmd.Key == "" {
		return "", fmt.Errorf("field DatabaseDel Key is required")
	}
	s = append(s, cmd.Key)
	return strings.Join(s, " "), nil
}
func (cmd DatabaseDelCommand) CommandString() string {
	return "DATABASE DEL"
}

// DatabaseDeltreeCommand Removes database keytree/value
type DatabaseDeltreeCommand struct {
	Family  string
	KeyTree string
}

func (cmd DatabaseDeltreeCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Family == "" {
		return "", fmt.Errorf("field DatabaseDeltree Family is required")
	}
	s = append(s, cmd.Family)
	s = append(s, cmd.KeyTree)
	return strings.Join(s, " "), nil
}
func (cmd DatabaseDeltreeCommand) CommandString() string {
	return "DATABASE DELTREE"
}

// DatabaseGetCommand Gets database value
type DatabaseGetCommand struct {
	Family string
	Key    string
}

func (cmd DatabaseGetCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Family == "" {
		return "", fmt.Errorf("field DatabaseGet Family is required")
	}
	s = append(s, cmd.Family)
	if cmd.Key == "" {
		return "", fmt.Errorf("field DatabaseGet Key is required")
	}
	s = append(s, cmd.Key)
	return strings.Join(s, " "), nil
}
func (cmd DatabaseGetCommand) CommandString() string {
	return "DATABASE GET"
}

// DatabasePutCommand Adds/updates database value
type DatabasePutCommand struct {
	Family string
	Key    string
	Value  string
}

func (cmd DatabasePutCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Family == "" {
		return "", fmt.Errorf("field DatabasePut Family is required")
	}
	s = append(s, cmd.Family)
	if cmd.Key == "" {
		return "", fmt.Errorf("field DatabasePut Key is required")
	}
	s = append(s, cmd.Key)
	if cmd.Value == "" {
		return "", fmt.Errorf("field DatabasePut Value is required")
	}
	s = append(s, cmd.Value)
	return strings.Join(s, " "), nil
}
func (cmd DatabasePutCommand) CommandString() string {
	return "DATABASE PUT"
}

// ExecCommand Executes a given Application
type ExecCommand struct {
	Application string
	Options     string
}

func (cmd ExecCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Application == "" {
		return "", fmt.Errorf("field Exec Application is required")
	}
	s = append(s, cmd.Application)
	if cmd.Options == "" {
		return "", fmt.Errorf("field Exec Options is required")
	}
	s = append(s, cmd.Options)
	return strings.Join(s, " "), nil
}
func (cmd ExecCommand) CommandString() string {
	return "EXEC"
}

// GetDataCommand Prompts for DTMF on a channel
type GetDataCommand struct {
	File      string
	Timeout   string
	Maxdigits string
}

func (cmd GetDataCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.File == "" {
		return "", fmt.Errorf("field GetData File is required")
	}
	s = append(s, cmd.File)
	s = append(s, cmd.Timeout)
	s = append(s, cmd.Maxdigits)
	return strings.Join(s, " "), nil
}
func (cmd GetDataCommand) CommandString() string {
	return "GET DATA"
}

// GetFullVariableCommand Evaluates a channel expression
type GetFullVariableCommand struct {
	Expression  string
	ChannelName string
}

func (cmd GetFullVariableCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Expression == "" {
		return "", fmt.Errorf("field GetFullVariable Expression is required")
	}
	s = append(s, cmd.Expression)
	s = append(s, cmd.ChannelName)
	return strings.Join(s, " "), nil
}
func (cmd GetFullVariableCommand) CommandString() string {
	return "GET FULL VARIABLE"
}

// GetOptionCommand Stream file, prompt for DTMF, with timeout.
type GetOptionCommand struct {
	FileName     string
	EscapeDigits string
	Timeout      string
}

func (cmd GetOptionCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.FileName == "" {
		return "", fmt.Errorf("field GetOption FileName is required")
	}
	s = append(s, cmd.FileName)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field GetOption EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	s = append(s, cmd.Timeout)
	return strings.Join(s, " "), nil
}
func (cmd GetOptionCommand) CommandString() string {
	return "GET OPTION"
}

// GetVariableCommand Gets a channel variable.
type GetVariableCommand struct {
	VariableName string
}

func (cmd GetVariableCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.VariableName == "" {
		return "", fmt.Errorf("field GetVariable VariableName is required")
	}
	s = append(s, cmd.VariableName)
	return strings.Join(s, " "), nil
}
func (cmd GetVariableCommand) CommandString() string {
	return "GET VARIABLE"
}

// HangupCommand Hangup a channel.
type HangupCommand struct {
	ChannelName string
}

func (cmd HangupCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	s = append(s, cmd.ChannelName)
	return strings.Join(s, " "), nil
}
func (cmd HangupCommand) CommandString() string {
	return "HANGUP"
}

// NoopCommand Does nothing.
type NoopCommand struct {
}

func (cmd NoopCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	return strings.Join(s, " "), nil
}
func (cmd NoopCommand) CommandString() string {
	return "NOOP"
}

// ReceiveCharCommand Receives one character from channels supporting it.
type ReceiveCharCommand struct {
	Timeout string
}

func (cmd ReceiveCharCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Timeout == "" {
		return "", fmt.Errorf("field ReceiveChar Timeout is required")
	}
	s = append(s, cmd.Timeout)
	return strings.Join(s, " "), nil
}
func (cmd ReceiveCharCommand) CommandString() string {
	return "RECEIVE CHAR"
}

// ReceiveTextCommand Receives text from channels supporting it.
type ReceiveTextCommand struct {
	Timeout string
}

func (cmd ReceiveTextCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Timeout == "" {
		return "", fmt.Errorf("field ReceiveText Timeout is required")
	}
	s = append(s, cmd.Timeout)
	return strings.Join(s, " "), nil
}
func (cmd ReceiveTextCommand) CommandString() string {
	return "RECEIVE TEXT"
}

// RecordFileCommand Records to a given file.
type RecordFileCommand struct {
	FileName      string
	Format        string
	EscapeDigits  string
	Timeout       string
	OffsetSamples string
	Beep          string
	SSilence      string
}

func (cmd RecordFileCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.FileName == "" {
		return "", fmt.Errorf("field RecordFile FileName is required")
	}
	s = append(s, cmd.FileName)
	if cmd.Format == "" {
		return "", fmt.Errorf("field RecordFile Format is required")
	}
	s = append(s, cmd.Format)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field RecordFile EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	if cmd.Timeout == "" {
		return "", fmt.Errorf("field RecordFile Timeout is required")
	}
	s = append(s, cmd.Timeout)
	s = append(s, cmd.OffsetSamples)
	s = append(s, cmd.Beep)
	s = append(s, cmd.SSilence)
	return strings.Join(s, " "), nil
}
func (cmd RecordFileCommand) CommandString() string {
	return "RECORD FILE"
}

// SayAlphaCommand Says a given character string.
type SayAlphaCommand struct {
	Number       string
	EscapeDigits string
}

func (cmd SayAlphaCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Number == "" {
		return "", fmt.Errorf("field SayAlpha Number is required")
	}
	s = append(s, cmd.Number)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayAlpha EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	return strings.Join(s, " "), nil
}
func (cmd SayAlphaCommand) CommandString() string {
	return "SAY ALPHA"
}

// SayDigitsCommand Says a given digit string.
type SayDigitsCommand struct {
	Number       string
	EscapeDigits string
}

func (cmd SayDigitsCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Number == "" {
		return "", fmt.Errorf("field SayDigits Number is required")
	}
	s = append(s, cmd.Number)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayDigits EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	return strings.Join(s, " "), nil
}
func (cmd SayDigitsCommand) CommandString() string {
	return "SAY DIGITS"
}

// SayNumberCommand Says a given number.
type SayNumberCommand struct {
	Number       string
	EscapeDigits string
	Gender       string
}

func (cmd SayNumberCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Number == "" {
		return "", fmt.Errorf("field SayNumber Number is required")
	}
	s = append(s, cmd.Number)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayNumber EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	s = append(s, cmd.Gender)
	return strings.Join(s, " "), nil
}
func (cmd SayNumberCommand) CommandString() string {
	return "SAY NUMBER"
}

// SayPhoneticCommand Says a given character string with phonetics.
type SayPhoneticCommand struct {
	String       string
	EscapeDigits string
}

func (cmd SayPhoneticCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.String == "" {
		return "", fmt.Errorf("field SayPhonetic String is required")
	}
	s = append(s, cmd.String)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayPhonetic EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	return strings.Join(s, " "), nil
}
func (cmd SayPhoneticCommand) CommandString() string {
	return "SAY PHONETIC"
}

// SayDateCommand Says a given date.
type SayDateCommand struct {
	Date         string
	EscapeDigits string
}

func (cmd SayDateCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Date == "" {
		return "", fmt.Errorf("field SayDate Date is required")
	}
	s = append(s, cmd.Date)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayDate EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	return strings.Join(s, " "), nil
}
func (cmd SayDateCommand) CommandString() string {
	return "SAY DATE"
}

// SayTimeCommand Says a given time.
type SayTimeCommand struct {
	Time         string
	EscapeDigits string
}

func (cmd SayTimeCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Time == "" {
		return "", fmt.Errorf("field SayTime Time is required")
	}
	s = append(s, cmd.Time)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayTime EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	return strings.Join(s, " "), nil
}
func (cmd SayTimeCommand) CommandString() string {
	return "SAY TIME"
}

// SayDatetimeCommand Says a given time as specified by the format given.
type SayDatetimeCommand struct {
	Time         string
	EscapeDigits string
	Format       string
	Timezone     string
}

func (cmd SayDatetimeCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Time == "" {
		return "", fmt.Errorf("field SayDatetime Time is required")
	}
	s = append(s, cmd.Time)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field SayDatetime EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	s = append(s, cmd.Format)
	s = append(s, cmd.Timezone)
	return strings.Join(s, " "), nil
}
func (cmd SayDatetimeCommand) CommandString() string {
	return "SAY DATETIME"
}

// SendImageCommand Sends images to channels supporting it.
type SendImageCommand struct {
	Image string
}

func (cmd SendImageCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Image == "" {
		return "", fmt.Errorf("field SendImage Image is required")
	}
	s = append(s, cmd.Image)
	return strings.Join(s, " "), nil
}
func (cmd SendImageCommand) CommandString() string {
	return "SEND IMAGE"
}

// SendTextCommand Sends text to channels supporting it.
type SendTextCommand struct {
	TextToSend string
}

func (cmd SendTextCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.TextToSend == "" {
		return "", fmt.Errorf("field SendText TextToSend is required")
	}
	s = append(s, cmd.TextToSend)
	return strings.Join(s, " "), nil
}
func (cmd SendTextCommand) CommandString() string {
	return "SEND TEXT"
}

// SetAutohangupCommand Autohangup channel in some time.
type SetAutohangupCommand struct {
	Time string
}

func (cmd SetAutohangupCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Time == "" {
		return "", fmt.Errorf("field SetAutohangup Time is required")
	}
	s = append(s, cmd.Time)
	return strings.Join(s, " "), nil
}
func (cmd SetAutohangupCommand) CommandString() string {
	return "SET AUTOHANGUP"
}

// SetCalleridCommand Sets callerid for the current channel.
type SetCalleridCommand struct {
	Number string
}

func (cmd SetCalleridCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Number == "" {
		return "", fmt.Errorf("field SetCallerid Number is required")
	}
	s = append(s, cmd.Number)
	return strings.Join(s, " "), nil
}
func (cmd SetCalleridCommand) CommandString() string {
	return "SET CALLERID"
}

// SetContextCommand Sets channel context.
type SetContextCommand struct {
	DesiredContext string
}

func (cmd SetContextCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.DesiredContext == "" {
		return "", fmt.Errorf("field SetContext DesiredContext is required")
	}
	s = append(s, cmd.DesiredContext)
	return strings.Join(s, " "), nil
}
func (cmd SetContextCommand) CommandString() string {
	return "SET CONTEXT"
}

// SetExtensionCommand Changes channel extension.
type SetExtensionCommand struct {
	NewExtension string
}

func (cmd SetExtensionCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.NewExtension == "" {
		return "", fmt.Errorf("field SetExtension NewExtension is required")
	}
	s = append(s, cmd.NewExtension)
	return strings.Join(s, " "), nil
}
func (cmd SetExtensionCommand) CommandString() string {
	return "SET EXTENSION"
}

// SetMusicCommand Enable/Disable Music on hold generator
type SetMusicCommand struct {
	Class string

	// has missing params

}

func (cmd SetMusicCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Class == "" {
		return "", fmt.Errorf("field SetMusic Class is required")
	}
	s = append(s, cmd.Class)
	return strings.Join(s, " "), nil
}
func (cmd SetMusicCommand) CommandString() string {
	return "SET MUSIC"
}

// SetPriorityCommand Set channel dialplan priority.
type SetPriorityCommand struct {
	Priority string
}

func (cmd SetPriorityCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Priority == "" {
		return "", fmt.Errorf("field SetPriority Priority is required")
	}
	s = append(s, cmd.Priority)
	return strings.Join(s, " "), nil
}
func (cmd SetPriorityCommand) CommandString() string {
	return "SET PRIORITY"
}

// SetVariableCommand Sets a channel variable.
type SetVariableCommand struct {
	VariableName string
	Value        string
}

func (cmd SetVariableCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.VariableName == "" {
		return "", fmt.Errorf("field SetVariable VariableName is required")
	}
	s = append(s, cmd.VariableName)
	if cmd.Value == "" {
		return "", fmt.Errorf("field SetVariable Value is required")
	}
	s = append(s, cmd.Value)
	return strings.Join(s, " "), nil
}
func (cmd SetVariableCommand) CommandString() string {
	return "SET VARIABLE"
}

// StreamFileCommand Sends audio file on channel.
type StreamFileCommand struct {
	FileName     string
	EscapeDigits string
	SampleOffset string
}

func (cmd StreamFileCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.FileName == "" {
		return "", fmt.Errorf("field StreamFile FileName is required")
	}
	s = append(s, cmd.FileName)
	if cmd.EscapeDigits == "" {
		return "", fmt.Errorf("field StreamFile EscapeDigits is required")
	}
	s = append(s, cmd.EscapeDigits)
	s = append(s, cmd.SampleOffset)
	return strings.Join(s, " "), nil
}
func (cmd StreamFileCommand) CommandString() string {
	return "STREAM FILE"
}

// TddModeCommand Toggles TDD mode (for the deaf).
type TddModeCommand struct {
	Boolean string
}

func (cmd TddModeCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Boolean == "" {
		return "", fmt.Errorf("field TddMode Boolean is required")
	}
	s = append(s, cmd.Boolean)
	return strings.Join(s, " "), nil
}
func (cmd TddModeCommand) CommandString() string {
	return "TDD MODE"
}

// VerboseCommand Logs a message to the asterisk verbose log.
type VerboseCommand struct {
	Message string
	Level   string
}

func (cmd VerboseCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Message == "" {
		return "", fmt.Errorf("field Verbose Message is required")
	}
	s = append(s, cmd.Message)
	if cmd.Level == "" {
		return "", fmt.Errorf("field Verbose Level is required")
	}
	s = append(s, cmd.Level)
	return strings.Join(s, " "), nil
}
func (cmd VerboseCommand) CommandString() string {
	return "VERBOSE"
}

// WaitForDigitCommand Waits for a digit to be pressed.
type WaitForDigitCommand struct {
	Timeout string
}

func (cmd WaitForDigitCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Timeout == "" {
		return "", fmt.Errorf("field WaitForDigit Timeout is required")
	}
	s = append(s, cmd.Timeout)
	return strings.Join(s, " "), nil
}
func (cmd WaitForDigitCommand) CommandString() string {
	return "WAIT FOR DIGIT"
}

// SpeechCreateCommand Creates a speech object.
type SpeechCreateCommand struct {
	Engine string
}

func (cmd SpeechCreateCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Engine == "" {
		return "", fmt.Errorf("field SpeechCreate Engine is required")
	}
	s = append(s, cmd.Engine)
	return strings.Join(s, " "), nil
}
func (cmd SpeechCreateCommand) CommandString() string {
	return "SPEECH CREATE"
}

// SpeechSetCommand Sets a speech engine setting.
type SpeechSetCommand struct {
	Name  string
	Value string
}

func (cmd SpeechSetCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Name == "" {
		return "", fmt.Errorf("field SpeechSet Name is required")
	}
	s = append(s, cmd.Name)
	if cmd.Value == "" {
		return "", fmt.Errorf("field SpeechSet Value is required")
	}
	s = append(s, cmd.Value)
	return strings.Join(s, " "), nil
}
func (cmd SpeechSetCommand) CommandString() string {
	return "SPEECH SET"
}

// SpeechDestroyCommand Destroys a speech object.
type SpeechDestroyCommand struct {
}

func (cmd SpeechDestroyCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	return strings.Join(s, " "), nil
}
func (cmd SpeechDestroyCommand) CommandString() string {
	return "SPEECH DESTROY"
}

// SpeechLoadGrammarCommand Loads a grammar.
type SpeechLoadGrammarCommand struct {
	GrammarName   string
	PathToGrammar string
}

func (cmd SpeechLoadGrammarCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.GrammarName == "" {
		return "", fmt.Errorf("field SpeechLoadGrammar GrammarName is required")
	}
	s = append(s, cmd.GrammarName)
	if cmd.PathToGrammar == "" {
		return "", fmt.Errorf("field SpeechLoadGrammar PathToGrammar is required")
	}
	s = append(s, cmd.PathToGrammar)
	return strings.Join(s, " "), nil
}
func (cmd SpeechLoadGrammarCommand) CommandString() string {
	return "SPEECH LOAD GRAMMAR"
}

// SpeechUnloadGrammarCommand Unloads a grammar.
type SpeechUnloadGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechUnloadGrammarCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.GrammarName == "" {
		return "", fmt.Errorf("field SpeechUnloadGrammar GrammarName is required")
	}
	s = append(s, cmd.GrammarName)
	return strings.Join(s, " "), nil
}
func (cmd SpeechUnloadGrammarCommand) CommandString() string {
	return "SPEECH UNLOAD GRAMMAR"
}

// SpeechActivateGrammarCommand Activates a grammar.
type SpeechActivateGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechActivateGrammarCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.GrammarName == "" {
		return "", fmt.Errorf("field SpeechActivateGrammar GrammarName is required")
	}
	s = append(s, cmd.GrammarName)
	return strings.Join(s, " "), nil
}
func (cmd SpeechActivateGrammarCommand) CommandString() string {
	return "SPEECH ACTIVATE GRAMMAR"
}

// SpeechDeactivateGrammarCommand Deactivates a grammar.
type SpeechDeactivateGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechDeactivateGrammarCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.GrammarName == "" {
		return "", fmt.Errorf("field SpeechDeactivateGrammar GrammarName is required")
	}
	s = append(s, cmd.GrammarName)
	return strings.Join(s, " "), nil
}
func (cmd SpeechDeactivateGrammarCommand) CommandString() string {
	return "SPEECH DEACTIVATE GRAMMAR"
}

// SpeechRecognizeCommand Recognizes speech.
type SpeechRecognizeCommand struct {
	Prompt  string
	Timeout string
	Offset  string
}

func (cmd SpeechRecognizeCommand) Command(c context.Context) (string, error) {
	s := []string{cmd.CommandString()}

	if cmd.Prompt == "" {
		return "", fmt.Errorf("field SpeechRecognize Prompt is required")
	}
	s = append(s, cmd.Prompt)
	if cmd.Timeout == "" {
		return "", fmt.Errorf("field SpeechRecognize Timeout is required")
	}
	s = append(s, cmd.Timeout)
	s = append(s, cmd.Offset)
	return strings.Join(s, " "), nil
}
func (cmd SpeechRecognizeCommand) CommandString() string {
	return "SPEECH RECOGNIZE"
}

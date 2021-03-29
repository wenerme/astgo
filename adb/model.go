package adb

import "time"

// https://wiki.asterisk.org/wiki/display/AST/Asterisk+12+CEL+Specification
type CallEventLog struct {
	ID          int64      `gorm:"primary_key"`
	EventType   string     `gorm:"column:eventtype"`
	EventTime   *time.Time `gorm:"column:eventtime"`
	UserDefType string     `gorm:"column:userdeftype"`
	CidName     string
	CidNum      string
	CidAni      string
	CidRdnis    string
	CidDnid     string
	Exten       string
	Context     string
	ChanName    string `gorm:"column:channame"`
	AppName     string `gorm:"column:appname"`
	AppData     string `gorm:"column:appdata"`
	AmaFlags    int    `gorm:"column:amaflags"`
	AccountCode string `gorm:"column:accountcode"`
	PeerAccount string `gorm:"column:peeraccount"`
	UniqueId    string `gorm:"column:uniqueid"`
	LinkedId    string `gorm:"column:linkedid"`
	UserField   string `gorm:"column:userfield"`
	Peer        string `gorm:"column:peer"`
}

func (CallEventLog) TableName() string {
	return "cel"
}

// https://wiki.asterisk.org/wiki/display/AST/Asterisk+12+CDR+Specification
type CallDetailRecord struct {
}

func (CallDetailRecord) TableName() string {
	return "cdr"
}

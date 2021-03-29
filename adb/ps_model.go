package adb

import "database/sql"

type YesNo string

const (
	Yes YesNo = "yes"
	No  YesNo = "no"
)

type PsAor struct {
	ID      string `gorm:"primary_key"`
	Contact sql.NullString
	// 3600
	DefaultExpiration   int
	MinimumExpiration   int
	MaximumExpiration   int
	Mailboxes           sql.NullString
	MaxContacts         int
	QualifyFrequency    int
	QualifyTimeout      float64
	AuthenticateQualify YesNo
	OutboundProxy       sql.NullString
	SupportPath         YesNo
	VoicemailExtension  sql.NullString
}

type PsAuth struct {
	ID string `gorm:"primary_key"`
	// userpass
	AuthType      string
	NonceLifetime sql.NullString
	Md5Cred       sql.NullString
	Username      string
	Password      string
	Realm         sql.NullString
}

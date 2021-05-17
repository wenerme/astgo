package astdb

import "database/sql"

//type YesNo string
//
//const (
//	Yes YesNo = "yes"
//	No  YesNo = "no"
//)
//
//type PjsipTransportProtocol string
//type PjsipTransportMethod string
//
//const (
//	PjsipTransportProtocolUDP PjsipTransportProtocol = "udp"
//	PjsipTransportProtocolTCP PjsipTransportProtocol = "tcp"
//	PjsipTransportProtocolTLS PjsipTransportProtocol = "tls"
//	PjsipTransportProtocolWS  PjsipTransportProtocol = "ws"
//	PjsipTransportProtocolWSS PjsipTransportProtocol = "wss"
//)
//
//const (
//	PjsipTransportMethodDefault     PjsipTransportProtocol = "default"
//	PjsipTransportMethodUnspecified PjsipTransportProtocol = "unspecified"
//	PjsipTransportMethodTLSv1       PjsipTransportProtocol = "tlsv1"
//	PjsipTransportMethodSSLv2       PjsipTransportProtocol = "sslv2"
//	PjsipTransportMethodSSLv3       PjsipTransportProtocol = "sslv3"
//	PjsipTransportMethodSSLv23      PjsipTransportProtocol = "sslv23"
//)

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

type PsContact struct {
	ID               string `gorm:"primary_key"`
	URI              string
	ExpirationTime   string
	QualifyFrequency int
}
type PsDomainAlias struct {
	ID     string `gorm:"primary_key"`
	Domain string
}
type PsEndpointIdIp struct {
	ID       string `gorm:"primary_key"`
	Endpoint string
	Match    string
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
type PsGlobal struct {
	ID                      string `gorm:"primary_key"`
	MaxForwards             *int
	UserAgent               *string
	DefaultOutboundEndpoint *string
}

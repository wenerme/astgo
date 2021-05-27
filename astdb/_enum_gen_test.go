package astdb_test

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
	"github.com/wenerme/astgo/pkg/tools/xmlgen"
)

func TestEnumGen(t *testing.T) {
	// from exported pg sql
	s := `
CREATE TYPE ast_bool_values AS ENUM ('0', '1', 'off', 'on', 'false', 'true', 'no', 'yes');
CREATE TYPE iax_encryption_values AS ENUM ('yes', 'no', 'aes128');
CREATE TYPE iax_requirecalltoken_values AS ENUM ('yes', 'no', 'auto');
CREATE TYPE iax_transfer_values AS ENUM ('yes', 'no', 'mediaonly');
CREATE TYPE moh_mode_values AS ENUM ('custom', 'files', 'mp3nb', 'quietmp3nb', 'quietmp3');
CREATE TYPE moh_mode_values AS ENUM ('custom', 'files', 'mp3nb', 'quietmp3nb', 'quietmp3', 'playlist');
CREATE TYPE pjsip_100rel_values AS ENUM ('no', 'required', 'yes');
CREATE TYPE pjsip_auth_type_values AS ENUM ('md5', 'userpass');
CREATE TYPE pjsip_auth_type_values_v2 AS ENUM ('md5', 'userpass', 'google_oauth');
CREATE TYPE pjsip_cid_privacy_values AS ENUM ('allowed_not_screened', 'allowed_passed_screened', 'allowed_failed_screened', 'allowed', 'prohib_not_screened', 'prohib_passed_screened', 'prohib_failed_screened', 'prohib', 'unavailable');
CREATE TYPE pjsip_connected_line_method_values AS ENUM ('invite', 'reinvite', 'update');
CREATE TYPE pjsip_direct_media_glare_mitigation_values AS ENUM ('none', 'outgoing', 'incoming');
CREATE TYPE pjsip_dtls_setup_values AS ENUM ('active', 'passive', 'actpass');
CREATE TYPE pjsip_dtmf_mode_values AS ENUM ('rfc4733', 'inband', 'info');
CREATE TYPE pjsip_dtmf_mode_values_v2 AS ENUM ('rfc4733', 'inband', 'info', 'auto');
CREATE TYPE pjsip_dtmf_mode_values_v3 AS ENUM ('rfc4733', 'inband', 'info', 'auto', 'auto_info');
CREATE TYPE pjsip_identify_by_values AS ENUM ('username');
CREATE TYPE pjsip_identify_by_values AS ENUM ('username', 'auth_username');
CREATE TYPE pjsip_identify_by_values AS ENUM ('username', 'auth_username', 'ip');
CREATE TYPE pjsip_media_encryption_values AS ENUM ('no', 'sdes', 'dtls');
CREATE TYPE pjsip_redirect_method_values AS ENUM ('user', 'uri_core', 'uri_pjsip');
CREATE TYPE pjsip_t38udptl_ec_values AS ENUM ('none', 'fec', 'redundancy');
CREATE TYPE pjsip_taskprocessor_overload_trigger_values AS ENUM ('none', 'global', 'pjsip_only');
CREATE TYPE pjsip_timer_values AS ENUM ('forced', 'no', 'required', 'yes');
CREATE TYPE pjsip_transport_method_values AS ENUM ('default', 'unspecified', 'tlsv1', 'sslv2', 'sslv3', 'sslv23');
CREATE TYPE pjsip_transport_protocol_values AS ENUM ('udp', 'tcp', 'tls', 'ws', 'wss');
CREATE TYPE pjsip_transport_protocol_values_v2 AS ENUM ('udp', 'tcp', 'tls', 'ws', 'wss', 'flow');
CREATE TYPE queue_autopause_values AS ENUM ('yes', 'no', 'all');
CREATE TYPE queue_strategy_values AS ENUM ('ringall', 'leastrecent', 'fewestcalls', 'random', 'rrmemory', 'linear', 'wrandom', 'rrordered');
CREATE TYPE sha_hash_values AS ENUM ('SHA-1', 'SHA-256');
CREATE TYPE sip_callingpres_values AS ENUM ('allowed_not_screened', 'allowed_passed_screen', 'allowed_failed_screen', 'allowed', 'prohib_not_screened', 'prohib_passed_screen', 'prohib_failed_screen', 'prohib');
CREATE TYPE sip_directmedia_values AS ENUM ('yes', 'no', 'nonat', 'update');
CREATE TYPE sip_directmedia_values_v2 AS ENUM ('yes', 'no', 'nonat', 'update', 'outgoing');
CREATE TYPE sip_dtmfmode_values AS ENUM ('rfc2833', 'info', 'shortinfo', 'inband', 'auto');
CREATE TYPE sip_progressinband_values AS ENUM ('yes', 'no', 'never');
CREATE TYPE sip_session_refresher_values AS ENUM ('uac', 'uas');
CREATE TYPE sip_session_timers_values AS ENUM ('accept', 'refuse', 'originate');
CREATE TYPE sip_transport_values AS ENUM ('udp', 'tcp', 'tls', 'ws', 'wss', 'udp,tcp', 'tcp,udp');
CREATE TYPE type_values AS ENUM ('friend', 'user', 'peer');
CREATE TYPE yes_no_values AS ENUM ('yes', 'no');
CREATE TYPE yesno_values AS ENUM ('yes', 'no');
`

	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	reg := regexp.MustCompile(`CREATE TYPE (\S+) AS ENUM\s\(([^)]+)`)
	enumDedup := map[string]*EnumModel{}
	for _, v := range lines {
		m := reg.FindStringSubmatch(v)
		if len(m) < 1 {
			fmt.Println("Failed Match", v)
			continue
		}
		d := &EnumModel{
			Name: m[1],
		}
		for _, v := range strings.Split(strings.ReplaceAll(m[2], "'", ""), ",") {
			d.Values = append(d.Values, strings.TrimSpace(v))
		}
		d.Name = strings.TrimSuffix(d.Name, "_values")
		d.Name = strings.TrimSuffix(d.Name, "_values_v2")
		d.Name = strings.TrimSuffix(d.Name, "_values_v3")
		switch d.Name {
		case "yesno":
			d.Name = "yes_no"
		case "type":
			d.Name = "sip_peer_type"
		}
		enumDedup[strings.ReplaceAll(d.Name, "_", "")] = d
	}
	var enums []*EnumModel
	for _, def := range enumDedup {
		def.Values = dedup(def.Values)
		enums = append(enums, def)
	}
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})
	for _, d := range enums {
		fmt.Println("Enum", d.Name, d.Values)
	}
	strcase.ConfigureAcronym("TCP", "tcp")
	strcase.ConfigureAcronym("UDP", "udp")
	strcase.ConfigureAcronym("WS", "ws")
	strcase.ConfigureAcronym("WSS", "wss")
	strcase.ConfigureAcronym("TLS", "tls")
	strcase.ConfigureAcronym("SSL", "ttl")

	tpl := template.Must(template.New("enum").Parse(`


{{define "enums"}}
// Code generated by enums_test.go, DO NOT EDIT.

package adb
{{- range $_,$e := .Enums}}
{{template "enum/type" $e}}
{{- end}}

{{- range $_,$e := .Enums}}
{{template "enum/func" $e}}
{{- end}}
{{end}}

{{define "enum/type"}}
type {{.TypeName}} string
{{$e := .}}
const (
{{- range $_,$v := $.Values}}
	{{$.ValueName $v}}  {{$e.TypeName}} = "{{$v}}"
{{- end}}
)
{{end}}

{{define "enum/func"}}
func (e {{.TypeName}}) String() string {
	return string(e)
}
func ({{.TypeName}}) Values() []string {
	return []string{
{{- range $_,$v := $.Values}}
	"{{$v}}",
{{- end}}
	}
}
{{end}}
`))
	buffer := bytes.NewBuffer(nil)
	if assert.NoError(t, tpl.ExecuteTemplate(buffer, "enums", &EnumGenModel{
		Enums: enums,
	})) {
		f := &xmlgen.File{Name: "inmem", Content: buffer.Bytes()}
		if !assert.NoError(t, xmlgen.GoFormatter(f)) {
			xmlgen.ReportFile(f)
		} else {
			assert.NoError(t, os.WriteFile("enums.go", f.Content, 0644))
		}
	}
}

type EnumModel struct {
	Name   string
	Values []string
}

func (e EnumModel) TypeName() string {
	return strcase.ToCamel(e.Name)
}
func (e EnumModel) ValueName(s string) string {
	return e.TypeName() + strcase.ToCamel(s)
}

type EnumGenModel struct {
	Enums []*EnumModel
}

func dedup(s []string) []string {
	m := map[string]bool{}
	for _, v := range s {
		m[v] = true
	}
	var o []string
	for k := range m {
		o = append(o, k)
	}
	sort.Strings(o)
	return o
}

package service

const (
	serviceTemplate = `[Unit]
{{- if .Unit.Description}}
Description={{.Unit.Description}}
{{- end}}
{{- range .Unit.Documentation}}
Documentation={{.}}
{{- end}}
{{- range .Unit.Before}}
Before={{.}}
{{- end}}
{{- range .Unit.After}}
After={{.}}
{{- end}}
{{- range .Unit.Requires}}
Requires={{.}}
{{- end}}
{{- range .Unit.Wants}}
Wants={{.}}
{{- end}}
{{- range .Unit.Conflicts}}
Conflicts={{.}}
{{- end}}
{{- range .Unit.OnFailure}}
OnFailure={{.}}
{{- end}}
{{- range .Unit.BindsTo}}
BindsTo={{.}}
{{- end}}
{{- range .Unit.PartOf}}
PartOf={{.}}
{{- end}}
{{- range .Unit.RequiredBy}}
RequiredBy={{.}}
{{- end}}
{{- range .Unit.WantedBy}}
WantedBy={{.}}
{{- end}}
{{- range .Unit.Alias}}
Alias={{.}}
{{- end}}

[Service]
Type={{.Service.Type}}
ExecStart={{.Service.ExecStart}}
{{- if .Service.ExecStartPre}}
ExecStartPre={{.Service.ExecStartPre}}
{{- end}}
{{- if .Service.ExecStartPost}}
ExecStartPost={{.Service.ExecStartPost}}
{{- end}}
{{- if .Service.ExecReload}}
ExecReload={{.Service.ExecReload}}
{{- end}}
{{- if .Service.ExecStop}}
ExecStop={{.Service.ExecStop}}
{{- end}}
{{- if .Service.ExecStopPost}}
ExecStopPost={{.Service.ExecStopPost}}
{{- end}}
{{- range .Service.ExecCondition}}
ExecCondition={{.}}
{{- end}}
{{- if .Service.Restart}}
Restart={{.Service.Restart}}
{{- end}}
{{- if .Service.RestartSec}}
RestartSec={{.Service.RestartSec}}
{{- end}}
{{- if .Service.TimeoutStartSec}}
TimeoutStartSec={{.Service.TimeoutStartSec}}
{{- end}}
{{- if .Service.TimeoutStopSec}}
TimeoutStopSec={{.Service.TimeoutStopSec}}
{{- end}}
{{- if .Service.PIDFile}}
PIDFile={{.Service.PIDFile}}
{{- end}}
{{- if .Service.User}}
User={{.Service.User}}
{{- end}}
{{- if .Service.Group}}
Group={{.Service.Group}}
{{- end}}
{{- if .Service.WorkingDirectory}}
WorkingDirectory={{.Service.WorkingDirectory}}
{{- end}}
{{- range $key, $value := .Service.Environment}}
Environment="{{$key}}={{$value}}"
{{- end}}
{{- if .Service.EnvironmentFile}}
EnvironmentFile={{.Service.EnvironmentFile}}
{{- end}}
{{- if .Service.Priority}}
Priority={{.Service.Priority}}
{{- end}}
{{- if .Service.Nice}}
Nice={{.Service.Nice}}
{{- end}}
{{- if .Service.AuthorizedKeysFile}}
AuthorizedKeysFile={{.Service.AuthorizedKeysFile}}
{{- end}}
{{- if .Service.RemainAfterExit}}
RemainAfterExit=yes
{{- else}}
RemainAfterExit=no
{{- end}}
{{- if .Service.CacheDirectory}}
CacheDirectory={{.Service.CacheDirectory}}
{{- end}}
{{- range .Service.CacheDirectories}}
CacheDirectory={{.}}
{{- end}}
{{- if .Service.ConfigDirectory}}
ConfigDirectory={{.Service.ConfigDirectory}}
{{- end}}
{{- range .Service.ConfigDirectories}}
ConfigDirectory={{.}}
{{- end}}
{{- if .Service.StateDirectory}}
StateDirectory={{.Service.StateDirectory}}
{{- end}}
{{- range .Service.StateDirectories}}
StateDirectory={{.}}
{{- end}}
{{- if .Service.LogsDirectory}}
LogsDirectory={{.Service.LogsDirectory}}
{{- end}}
{{- range .Service.LogsDirectories}}
LogsDirectory={{.}}
{{- end}}
{{- if .Service.RuntimeDirectory}}
RuntimeDirectory={{.Service.RuntimeDirectory}}
{{- end}}
{{- range .Service.RuntimeDirectories}}
RuntimeDirectory={{.}}
{{- end}}
{{- if .Service.StandardInput}}
StandardInput={{.Service.StandardInput}}
{{- end}}
{{- if .Service.StandardOutput}}
StandardOutput={{.Service.StandardOutput}}
{{- end}}
{{- if .Service.StandardError}}
StandardError={{.Service.StandardError}}
{{- end}}
{{- if .Service.SyslogIdentifier}}
SyslogIdentifier={{.Service.SyslogIdentifier}}
{{- end}}
{{- if .Service.SystemCallArchitectures}}
SystemCallArchitectures={{.Service.SystemCallArchitectures}}
{{- end}}
{{- if .Service.SystemCallFilter}}
SystemCallFilter={{.Service.SystemCallFilter}}
{{- end}}
{{- if .Service.SystemCallErrorNumber}}
SystemCallErrorNumber={{.Service.SystemCallErrorNumber}}
{{- end}}
{{- if .Service.SystemCallLog}}
SystemCallLog={{.Service.SystemCallLog}}
{{- end}}

[Install]
{{- range .Install.Alias}}
Alias={{.}}
{{- end}}
{{- range .Install.RequiredBy}}
RequiredBy={{.}}
{{- end}}
{{- range .Install.WantedBy}}
WantedBy={{.}}
{{- end}}
`
)

// Unit 代表[Unit]部分的配置
type Unit struct {
	Description   string
	Documentation []string
	Before        []string
	After         []string
	Requires      []string
	Wants         []string
	Conflicts     []string
	OnFailure     []string
	BindsTo       []string
	PartOf        []string
	RequiredBy    []string
	WantedBy      []string
	Alias         []string
}

// Service 代表[Service]部分的配置
type Service struct {
	Type                    string
	ExecStart               string
	ExecStartPre            string
	ExecStartPost           string
	ExecReload              string
	ExecStop                string
	ExecStopPost            string
	ExecCondition           []string
	Restart                 string
	RestartSec              int
	TimeoutStartSec         int
	TimeoutStopSec          int
	PIDFile                 string
	User                    string
	Group                   string
	WorkingDirectory        string
	Environment             map[string]string
	EnvironmentFile         string
	Priority                int
	Nice                    int
	AuthorizedKeysFile      string
	RemainAfterExit         bool
	CacheDirectory          string
	CacheDirectories        []string
	ConfigDirectory         string
	ConfigDirectories       []string
	StateDirectory          string
	StateDirectories        []string
	LogsDirectory           string
	LogsDirectories         []string
	RuntimeDirectory        string
	RuntimeDirectories      []string
	StandardInput           string
	StandardOutput          string
	StandardError           string
	SyslogIdentifier        string
	SystemCallArchitectures string
	SystemCallFilter        string
	SystemCallErrorNumber   string
	SystemCallLog           bool
}

// Install 代表[Install]部分的配置
type Install struct {
	Alias      []string
	RequiredBy []string
	WantedBy   []string
}

// ServiceConfig 结合了所有部分的配置
type ServiceConfig struct {
	Unit    Unit
	Service Service
	Install Install
}

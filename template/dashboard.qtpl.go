// This file is automatically generated by qtc from "dashboard.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/dashboard.qtpl:2
package template

//line template/dashboard.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/dashboard.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/dashboard.qtpl:2
type dashboardPage interface {
	//line template/dashboard.qtpl:2
	Title() string
	//line template/dashboard.qtpl:2
	StreamTitle(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteTitle(qq422016 qtio422016.Writer)
	//line template/dashboard.qtpl:2
	Styles() string
	//line template/dashboard.qtpl:2
	StreamStyles(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteStyles(qq422016 qtio422016.Writer)
	//line template/dashboard.qtpl:2
	Body() string
	//line template/dashboard.qtpl:2
	StreamBody(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteBody(qq422016 qtio422016.Writer)
	//line template/dashboard.qtpl:2
	Footer() string
	//line template/dashboard.qtpl:2
	StreamFooter(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteFooter(qq422016 qtio422016.Writer)
	//line template/dashboard.qtpl:2
	Actions() string
	//line template/dashboard.qtpl:2
	StreamActions(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteActions(qq422016 qtio422016.Writer)
	//line template/dashboard.qtpl:2
	Header() string
	//line template/dashboard.qtpl:2
	StreamHeader(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteHeader(qq422016 qtio422016.Writer)
	//line template/dashboard.qtpl:2
	Navigation() string
	//line template/dashboard.qtpl:2
	StreamNavigation(qw422016 *qt422016.Writer)
	//line template/dashboard.qtpl:2
	WriteNavigation(qq422016 qtio422016.Writer)
//line template/dashboard.qtpl:2
}

//line template/dashboard.qtpl:21
type Alert struct {
	Level   level
	Message string
}

type Dashboard struct {
	dashboardPage

	alert Alert

	URI string
}

type level uint8

var (
	NamespacesURI = "(\\/namespaces\\/?|\\/u\\/[_-a-zA-Z0-9\\S.]+\\/[-a-zA-Z0-9\\/\\S]*\\/?[a-z]*)"
	ObjectsURI    = "(^\\/objects\\/?[a-z0-9\\/?]*$)"
	BuildsURI     = "(^\\/$|^\\/builds\\/?[a-z0-9\\/?]*$)"
	SettingsURI   = "\\/settings\\/?"
)

const (
	success level = iota
	warn
	danger
)

func NewDashboard(p dashboardPage, uri string, a Alert) *Dashboard {
	return &Dashboard{
		dashboardPage: p,
		alert:         a,
		URI:           uri,
	}
}

func Danger(msg string) Alert {
	return Alert{
		Level:   danger,
		Message: msg,
	}
}

func Success(msg string) Alert {
	return Alert{
		Level:   success,
		Message: msg,
	}
}

func Warn(msg string) Alert {
	return Alert{
		Level:   warn,
		Message: msg,
	}
}

func (a Alert) IsZero() bool {
	return a.Level == level(0) && a.Message == ""
}

func (l level) String() string {
	switch l {
	case success:
		return "success"
	case warn:
		return "warn"
	case danger:
		return "danger"
	default:
		return ""
	}
}

//line template/dashboard.qtpl:97
func (p *Dashboard) StreamBody(qw422016 *qt422016.Writer) {
	//line template/dashboard.qtpl:97
	qw422016.N().S(` <div class="dashboard"> <div class="dashboard-content"> `)
	//line template/dashboard.qtpl:100
	if !p.alert.IsZero() {
		//line template/dashboard.qtpl:100
		qw422016.N().S(` <div class="alert alert-`)
		//line template/dashboard.qtpl:101
		qw422016.E().S(p.alert.Level.String())
		//line template/dashboard.qtpl:101
		qw422016.N().S(`"> <div class="alert-message">`)
		//line template/dashboard.qtpl:102
		qw422016.E().S(p.alert.Message)
		//line template/dashboard.qtpl:102
		qw422016.N().S(`</div> <a href="`)
		//line template/dashboard.qtpl:103
		qw422016.E().S(p.URI)
		//line template/dashboard.qtpl:103
		qw422016.N().S(`">`)
		//line template/dashboard.qtpl:103
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
		//line template/dashboard.qtpl:103
		qw422016.N().S(`</a> </div> `)
		//line template/dashboard.qtpl:105
	}
	//line template/dashboard.qtpl:105
	qw422016.N().S(` <div class="dashboard-wrap"> <div class="dashboard-header"> <div class="overflow"> <h1>`)
	//line template/dashboard.qtpl:109
	p.dashboardPage.StreamHeader(qw422016)
	//line template/dashboard.qtpl:109
	qw422016.N().S(`</h1> <ul class="dashboard-actions">`)
	//line template/dashboard.qtpl:110
	p.dashboardPage.StreamActions(qw422016)
	//line template/dashboard.qtpl:110
	qw422016.N().S(`</ul> </div> <ul class="dashboard-nav">`)
	//line template/dashboard.qtpl:112
	p.dashboardPage.StreamNavigation(qw422016)
	//line template/dashboard.qtpl:112
	qw422016.N().S(`</ul> </div> <div class="dashboard-body">`)
	//line template/dashboard.qtpl:114
	p.dashboardPage.StreamBody(qw422016)
	//line template/dashboard.qtpl:114
	qw422016.N().S(`</div> </div> </div> <div class="sidebar"> <div class="sidebar-header"> <div class="logo"><div class="left"></div><div class="right"></div></div> <h2>Thrall</h2> </div> <ul class="sidebar-nav"> <li class="sidebar-nav-header">MANAGE</li> <li>`)
	//line template/dashboard.qtpl:124
	StreamRenderLinkPattern(qw422016, "/", BuildsURI, p.URI)
	//line template/dashboard.qtpl:124
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M22.688 18.984c0.422 0.281 0.422 0.984-0.094 1.406l-2.297 2.297c-0.422 0.422-0.984 0.422-1.406 0l-9.094-9.094c-2.297 0.891-4.969 0.422-6.891-1.5-2.016-2.016-2.531-5.016-1.313-7.406l4.406 4.313 3-3-4.313-4.313c2.391-1.078 5.391-0.703 7.406 1.313 1.922 1.922 2.391 4.594 1.5 6.891z"></path>
</svg>
`)
	//line template/dashboard.qtpl:124
	qw422016.N().S(`<span>Builds</span></a></li> <li>`)
	//line template/dashboard.qtpl:125
	StreamRenderLinkPattern(qw422016, "/namespaces", NamespacesURI, p.URI)
	//line template/dashboard.qtpl:125
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9.984 3.984l2.016 2.016h8.016c1.078 0 1.969 0.938 1.969 2.016v9.984c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h6z"></path>
</svg>
`)
	//line template/dashboard.qtpl:125
	qw422016.N().S(`<span>Namespaces</span></a></li> <li>`)
	//line template/dashboard.qtpl:126
	StreamRenderLinkPattern(qw422016, "/objects", ObjectsURI, p.URI)
	//line template/dashboard.qtpl:126
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM9 15.984v-6h-3.984l6.984-6.984 6.984 6.984h-3.984v6h-6z"></path>
</svg>
`)
	//line template/dashboard.qtpl:126
	qw422016.N().S(`<span>Objects</span></a></li> <li><a>`)
	//line template/dashboard.qtpl:127
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
	//line template/dashboard.qtpl:127
	qw422016.N().S(`<span>Variables</span></a></li> <li><a>`)
	//line template/dashboard.qtpl:128
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M6.984 14.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-1.969 0.938-1.969 2.016 0.891 2.016 1.969 2.016zM12.656 9.984h10.359v4.031h-2.016v3.984h-3.984v-3.984h-4.359c-0.797 2.344-3.047 3.984-5.672 3.984-3.328 0-6-2.672-6-6s2.672-6 6-6c2.625 0 4.875 1.641 5.672 3.984z"></path>
</svg>
`)
	//line template/dashboard.qtpl:128
	qw422016.N().S(`<span>SSH Keys</span></a></li> <li class="sidebar-nav-header">ACCOUNT</li> <li>`)
	//line template/dashboard.qtpl:130
	StreamRenderLinkPattern(qw422016, "/settings", SettingsURI, p.URI)
	//line template/dashboard.qtpl:130
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 15.516c1.922 0 3.516-1.594 3.516-3.516s-1.594-3.516-3.516-3.516-3.516 1.594-3.516 3.516 1.594 3.516 3.516 3.516zM19.453 12.984l2.109 1.641c0.188 0.141 0.234 0.422 0.094 0.656l-2.016 3.469c-0.141 0.234-0.375 0.281-0.609 0.188l-2.484-0.984c-0.516 0.375-1.078 0.75-1.688 0.984l-0.375 2.625c-0.047 0.234-0.234 0.422-0.469 0.422h-4.031c-0.234 0-0.422-0.188-0.469-0.422l-0.375-2.625c-0.609-0.234-1.172-0.563-1.688-0.984l-2.484 0.984c-0.234 0.094-0.469 0.047-0.609-0.188l-2.016-3.469c-0.141-0.234-0.094-0.516 0.094-0.656l2.109-1.641c-0.047-0.328-0.047-0.656-0.047-0.984s0-0.656 0.047-0.984l-2.109-1.641c-0.188-0.141-0.234-0.422-0.094-0.656l2.016-3.469c0.141-0.234 0.375-0.281 0.609-0.188l2.484 0.984c0.516-0.375 1.078-0.75 1.688-0.984l0.375-2.625c0.047-0.234 0.234-0.422 0.469-0.422h4.031c0.234 0 0.422 0.188 0.469 0.422l0.375 2.625c0.609 0.234 1.172 0.563 1.688 0.984l2.484-0.984c0.234-0.094 0.469-0.047 0.609 0.188l2.016 3.469c0.141 0.234 0.094 0.516-0.094 0.656l-2.109 1.641c0.047 0.328 0.047 0.656 0.047 0.984s0 0.656-0.047 0.984z"></path>
</svg>
`)
	//line template/dashboard.qtpl:130
	qw422016.N().S(`<span>Settings</span></a></li> <li><form method="POST" action="/logout"><button type="submit">`)
	//line template/dashboard.qtpl:131
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 3c1.078 0 2.016 0.938 2.016 2.016v13.969c0 1.078-0.938 2.016-2.016 2.016h-13.969c-1.125 0-2.016-0.938-2.016-2.016v-3.984h2.016v3.984h13.969v-13.969h-13.969v3.984h-2.016v-3.984c0-1.078 0.891-2.016 2.016-2.016h13.969zM10.078 15.609l2.578-2.625h-9.656v-1.969h9.656l-2.578-2.625 1.406-1.406 5.016 5.016-5.016 5.016z"></path>
</svg>
`)
	//line template/dashboard.qtpl:131
	qw422016.N().S(`<span>Logout</span></button></form></li> </ul> </div> </div> `)
//line template/dashboard.qtpl:135
}

//line template/dashboard.qtpl:135
func (p *Dashboard) WriteBody(qq422016 qtio422016.Writer) {
	//line template/dashboard.qtpl:135
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/dashboard.qtpl:135
	p.StreamBody(qw422016)
	//line template/dashboard.qtpl:135
	qt422016.ReleaseWriter(qw422016)
//line template/dashboard.qtpl:135
}

//line template/dashboard.qtpl:135
func (p *Dashboard) Body() string {
	//line template/dashboard.qtpl:135
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/dashboard.qtpl:135
	p.WriteBody(qb422016)
	//line template/dashboard.qtpl:135
	qs422016 := string(qb422016.B)
	//line template/dashboard.qtpl:135
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/dashboard.qtpl:135
	return qs422016
//line template/dashboard.qtpl:135
}

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
var (
	NamespacesURI = "(\\/namespaces\\/?|\\/u\\/[_-a-zA-Z0-9\\S.]+\\/[-a-zA-Z0-9\\/\\S]*\\/?[a-z]*)"
	BuildsURI     = "(^\\/$|^\\/builds\\/?[a-z0-9]*$)"
	SettingsURI   = "\\/settings\\/?"
)

type Dashboard struct {
	dashboardPage

	URI string
}

func NewDashboard(p dashboardPage, uri string) *Dashboard {
	return &Dashboard{dashboardPage: p, URI: uri}
}

//line template/dashboard.qtpl:39
func (p *Dashboard) StreamBody(qw422016 *qt422016.Writer) {
	//line template/dashboard.qtpl:39
	qw422016.N().S(` <div class="dashboard"> <div class="content"> <div class="header"> <h1>`)
	//line template/dashboard.qtpl:43
	p.dashboardPage.StreamHeader(qw422016)
	//line template/dashboard.qtpl:43
	qw422016.N().S(`</h1> <ul class="actions">`)
	//line template/dashboard.qtpl:44
	p.dashboardPage.StreamActions(qw422016)
	//line template/dashboard.qtpl:44
	qw422016.N().S(`</ul> </div> <div class="navigation">`)
	//line template/dashboard.qtpl:46
	p.dashboardPage.StreamNavigation(qw422016)
	//line template/dashboard.qtpl:46
	qw422016.N().S(`</div> <div class="body">`)
	//line template/dashboard.qtpl:47
	p.dashboardPage.StreamBody(qw422016)
	//line template/dashboard.qtpl:47
	qw422016.N().S(`</div> </div> <div class="sidebar"> <h2> <div class="brand"> <div class="left"></div> <div class="right"></div> </div> Thrall </h2> <ul> <li class="nav-header">MANAGE</li> <li>`)
	//line template/dashboard.qtpl:59
	StreamRenderLinkPattern(qw422016, "/", BuildsURI, p.URI)
	//line template/dashboard.qtpl:59
	qw422016.N().S(`Builds</a></li> <li>`)
	//line template/dashboard.qtpl:60
	StreamRenderLinkPattern(qw422016, "/namespaces", NamespacesURI, p.URI)
	//line template/dashboard.qtpl:60
	qw422016.N().S(`Namespaces</a></li> <li class="nav-header">ACCOUNT</li> <li>`)
	//line template/dashboard.qtpl:62
	StreamRenderLinkPattern(qw422016, "/settings", SettingsURI, p.URI)
	//line template/dashboard.qtpl:62
	qw422016.N().S(`Settings</a></li> <li><form method="POST" action="/logout"><button type="submit">Logout</button></form></li> </ul> </div> </div> `)
//line template/dashboard.qtpl:67
}

//line template/dashboard.qtpl:67
func (p *Dashboard) WriteBody(qq422016 qtio422016.Writer) {
	//line template/dashboard.qtpl:67
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/dashboard.qtpl:67
	p.StreamBody(qw422016)
	//line template/dashboard.qtpl:67
	qt422016.ReleaseWriter(qw422016)
//line template/dashboard.qtpl:67
}

//line template/dashboard.qtpl:67
func (p *Dashboard) Body() string {
	//line template/dashboard.qtpl:67
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/dashboard.qtpl:67
	p.WriteBody(qb422016)
	//line template/dashboard.qtpl:67
	qs422016 := string(qb422016.B)
	//line template/dashboard.qtpl:67
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/dashboard.qtpl:67
	return qs422016
//line template/dashboard.qtpl:67
}

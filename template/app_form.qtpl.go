// Code generated by qtc from "app_form.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/app_form.qtpl:2
package template

//line template/app_form.qtpl:2
import (
	"djinn-ci.com/oauth2"
	"djinn-ci.com/template/form"
)

//line template/app_form.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/app_form.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/app_form.qtpl:9
type AppForm struct {
	*form.Form

	App *oauth2.App
}

func (p *AppForm) action() string {
	if p.App != nil {
		return p.App.Endpoint()
	}
	return "/settings/apps"
}

//line template/app_form.qtpl:24
func (p *AppForm) StreamTitle(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:24
	qw422016.N().S(` `)
//line template/app_form.qtpl:25
	if p.App == nil {
//line template/app_form.qtpl:25
		qw422016.N().S(` New App `)
//line template/app_form.qtpl:27
	} else {
//line template/app_form.qtpl:27
		qw422016.N().S(` Edit App `)
//line template/app_form.qtpl:29
	}
//line template/app_form.qtpl:29
	qw422016.N().S(` `)
//line template/app_form.qtpl:30
}

//line template/app_form.qtpl:30
func (p *AppForm) WriteTitle(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:30
	p.StreamTitle(qw422016)
//line template/app_form.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:30
}

//line template/app_form.qtpl:30
func (p *AppForm) Title() string {
//line template/app_form.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:30
	p.WriteTitle(qb422016)
//line template/app_form.qtpl:30
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:30
	return qs422016
//line template/app_form.qtpl:30
}

//line template/app_form.qtpl:32
func (p *AppForm) StreamHeader(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:32
	qw422016.N().S(` `)
//line template/app_form.qtpl:33
	if p.App == nil {
//line template/app_form.qtpl:33
		qw422016.N().S(` <a href="/settings/apps" class="back">`)
//line template/app_form.qtpl:34
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line template/app_form.qtpl:34
		qw422016.N().S(`</a> Settings - New App `)
//line template/app_form.qtpl:35
	} else {
//line template/app_form.qtpl:35
		qw422016.N().S(` <a href="/settings/apps" class="back">`)
//line template/app_form.qtpl:36
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line template/app_form.qtpl:36
		qw422016.N().S(`</a> Settings - Edit App `)
//line template/app_form.qtpl:37
	}
//line template/app_form.qtpl:37
	qw422016.N().S(` `)
//line template/app_form.qtpl:38
}

//line template/app_form.qtpl:38
func (p *AppForm) WriteHeader(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:38
	p.StreamHeader(qw422016)
//line template/app_form.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:38
}

//line template/app_form.qtpl:38
func (p *AppForm) Header() string {
//line template/app_form.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:38
	p.WriteHeader(qb422016)
//line template/app_form.qtpl:38
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:38
	return qs422016
//line template/app_form.qtpl:38
}

//line template/app_form.qtpl:40
func (p *AppForm) StreamActions(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:40
}

//line template/app_form.qtpl:40
func (p *AppForm) WriteActions(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:40
	p.StreamActions(qw422016)
//line template/app_form.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:40
}

//line template/app_form.qtpl:40
func (p *AppForm) Actions() string {
//line template/app_form.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:40
	p.WriteActions(qb422016)
//line template/app_form.qtpl:40
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:40
	return qs422016
//line template/app_form.qtpl:40
}

//line template/app_form.qtpl:41
func (p *AppForm) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:41
}

//line template/app_form.qtpl:41
func (p *AppForm) WriteNavigation(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:41
	p.StreamNavigation(qw422016)
//line template/app_form.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:41
}

//line template/app_form.qtpl:41
func (p *AppForm) Navigation() string {
//line template/app_form.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:41
	p.WriteNavigation(qb422016)
//line template/app_form.qtpl:41
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:41
	return qs422016
//line template/app_form.qtpl:41
}

//line template/app_form.qtpl:42
func (p *AppForm) StreamFooter(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:42
}

//line template/app_form.qtpl:42
func (p *AppForm) WriteFooter(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:42
	p.StreamFooter(qw422016)
//line template/app_form.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:42
}

//line template/app_form.qtpl:42
func (p *AppForm) Footer() string {
//line template/app_form.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:42
	p.WriteFooter(qb422016)
//line template/app_form.qtpl:42
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:42
	return qs422016
//line template/app_form.qtpl:42
}

//line template/app_form.qtpl:44
func (p *AppForm) streamrenderAppDetails(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:44
	qw422016.N().S(` `)
//line template/app_form.qtpl:45
	if p.App != nil {
//line template/app_form.qtpl:45
		qw422016.N().S(` <div class="panel-body slim"> <h1 class="mb-10">`)
//line template/app_form.qtpl:47
		qw422016.E().S(p.App.Name)
//line template/app_form.qtpl:47
		qw422016.N().S(`</h1> <div class="mb-10"> <strong>Client ID</strong> <br/><code><span class="muted">`)
//line template/app_form.qtpl:50
		qw422016.E().S(p.App.ClientID)
//line template/app_form.qtpl:50
		qw422016.N().S(`</span></code><br/><br/> <strong>Client Secret</strong> <br/><code><span class="muted">`)
//line template/app_form.qtpl:52
		qw422016.E().S(p.App.ClientSecret.String())
//line template/app_form.qtpl:52
		qw422016.N().S(`</span></code></br> </div> <form method="POST" action="`)
//line template/app_form.qtpl:54
		qw422016.E().S(p.App.Endpoint("revoke"))
//line template/app_form.qtpl:54
		qw422016.N().S(`" class="inline-block"> `)
//line template/app_form.qtpl:55
		form.StreamMethod(qw422016, "PATCH")
//line template/app_form.qtpl:55
		qw422016.N().S(` `)
//line template/app_form.qtpl:56
		qw422016.N().V(p.CSRF)
//line template/app_form.qtpl:56
		qw422016.N().S(` <div class="form-field-inline"> <button type="submit" class="btn btn-danger">Revoke Tokens</button> </div> </form> <form method="POST" action="`)
//line template/app_form.qtpl:61
		qw422016.E().S(p.App.Endpoint("reset"))
//line template/app_form.qtpl:61
		qw422016.N().S(`" class="inline-block"> `)
//line template/app_form.qtpl:62
		form.StreamMethod(qw422016, "PATCH")
//line template/app_form.qtpl:62
		qw422016.N().S(` `)
//line template/app_form.qtpl:63
		qw422016.N().V(p.CSRF)
//line template/app_form.qtpl:63
		qw422016.N().S(` <button type="submit" class="btn btn-danger">Reset Secret</button> </form> </div> `)
//line template/app_form.qtpl:67
	}
//line template/app_form.qtpl:67
	qw422016.N().S(` `)
//line template/app_form.qtpl:68
}

//line template/app_form.qtpl:68
func (p *AppForm) writerenderAppDetails(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:68
	p.streamrenderAppDetails(qw422016)
//line template/app_form.qtpl:68
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:68
}

//line template/app_form.qtpl:68
func (p *AppForm) renderAppDetails() string {
//line template/app_form.qtpl:68
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:68
	p.writerenderAppDetails(qb422016)
//line template/app_form.qtpl:68
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:68
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:68
	return qs422016
//line template/app_form.qtpl:68
}

//line template/app_form.qtpl:70
func (p *AppForm) StreamBody(qw422016 *qt422016.Writer) {
//line template/app_form.qtpl:70
	qw422016.N().S(` <div class="panel"> `)
//line template/app_form.qtpl:72
	p.streamrenderAppDetails(qw422016)
//line template/app_form.qtpl:72
	qw422016.N().S(` <form action="`)
//line template/app_form.qtpl:73
	qw422016.E().S(p.action())
//line template/app_form.qtpl:73
	qw422016.N().S(`" class="panel-body slim" method="POST"> `)
//line template/app_form.qtpl:74
	if p.App != nil {
//line template/app_form.qtpl:74
		form.StreamMethod(qw422016, "PATCH")
//line template/app_form.qtpl:74
	}
//line template/app_form.qtpl:74
	qw422016.N().S(` `)
//line template/app_form.qtpl:75
	qw422016.N().V(p.CSRF)
//line template/app_form.qtpl:75
	qw422016.N().S(` `)
//line template/app_form.qtpl:76
	p.StreamField(qw422016, form.Field{
		ID:   "name",
		Name: "Name",
		Type: form.Text,
	})
//line template/app_form.qtpl:80
	qw422016.N().S(` `)
//line template/app_form.qtpl:81
	p.StreamField(qw422016, form.Field{
		ID:       "description",
		Name:     "Description",
		Type:     form.Text,
		Optional: true,
	})
//line template/app_form.qtpl:86
	qw422016.N().S(` `)
//line template/app_form.qtpl:87
	p.StreamField(qw422016, form.Field{
		ID:       "homepage_uri",
		Name:     "Homepage URI",
		Type:     form.Text,
		Optional: true,
	})
//line template/app_form.qtpl:92
	qw422016.N().S(` `)
//line template/app_form.qtpl:93
	p.StreamField(qw422016, form.Field{
		ID:       "redirect_uri",
		Name:     "Redirect URI",
		Type:     form.Text,
		Optional: true,
	})
//line template/app_form.qtpl:98
	qw422016.N().S(` <div class="form-field"> `)
//line template/app_form.qtpl:100
	if p.App == nil {
//line template/app_form.qtpl:100
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line template/app_form.qtpl:102
	} else {
//line template/app_form.qtpl:102
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line template/app_form.qtpl:104
	}
//line template/app_form.qtpl:104
	qw422016.N().S(` </div> </form> </div> `)
//line template/app_form.qtpl:108
}

//line template/app_form.qtpl:108
func (p *AppForm) WriteBody(qq422016 qtio422016.Writer) {
//line template/app_form.qtpl:108
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/app_form.qtpl:108
	p.StreamBody(qw422016)
//line template/app_form.qtpl:108
	qt422016.ReleaseWriter(qw422016)
//line template/app_form.qtpl:108
}

//line template/app_form.qtpl:108
func (p *AppForm) Body() string {
//line template/app_form.qtpl:108
	qb422016 := qt422016.AcquireByteBuffer()
//line template/app_form.qtpl:108
	p.WriteBody(qb422016)
//line template/app_form.qtpl:108
	qs422016 := string(qb422016.B)
//line template/app_form.qtpl:108
	qt422016.ReleaseByteBuffer(qb422016)
//line template/app_form.qtpl:108
	return qs422016
//line template/app_form.qtpl:108
}
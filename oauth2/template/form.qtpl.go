// Code generated by qtc from "form.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line oauth2/template/form.qtpl:2
package template

//line oauth2/template/form.qtpl:2
import (
	"fmt"
	"strings"

	"github.com/andrewpillar/thrall/oauth2"
	"github.com/andrewpillar/thrall/template"
)

//line oauth2/template/form.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line oauth2/template/form.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line oauth2/template/form.qtpl:12
type AppForm struct {
	template.BasePage
	template.Form

	App *oauth2.App
}

type TokenForm struct {
	template.BasePage
	template.Form

	Token  *oauth2.Token
	Scopes map[string]struct{}
}

func (p *TokenForm) field(field string) string {
	old := p.Form.Fields[field]

	if p.Token != nil {
		if old != "" {
			return old
		}

		if field == "name" {
			return p.Token.Name
		}
		return ""
	}
	return old
}

func (p *AppForm) action() string {
	if p.App == nil {
		return "/settings/apps"
	}
	return p.App.Endpoint()
}

func (p *AppForm) field(field string) string {
	old := p.Form.Fields[field]

	if p.App != nil {
		if old != "" {
			return old
		}

		switch field {
		case "name":
			return p.App.Name
		case "description":
			return p.App.Description
		case "homepage_uri":
			return p.App.HomeURI
		case "redirect_uri":
			return p.App.RedirectURI
		default:
			return ""
		}
	}
	return old
}

//line oauth2/template/form.qtpl:76
func (p *AppForm) StreamTitle(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:76
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:77
	if p.App == nil {
//line oauth2/template/form.qtpl:77
		qw422016.N().S(` Settings - New App `)
//line oauth2/template/form.qtpl:79
	} else {
//line oauth2/template/form.qtpl:79
		qw422016.N().S(` Settings - Edit App `)
//line oauth2/template/form.qtpl:81
	}
//line oauth2/template/form.qtpl:81
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:82
}

//line oauth2/template/form.qtpl:82
func (p *AppForm) WriteTitle(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:82
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:82
	p.StreamTitle(qw422016)
//line oauth2/template/form.qtpl:82
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:82
}

//line oauth2/template/form.qtpl:82
func (p *AppForm) Title() string {
//line oauth2/template/form.qtpl:82
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:82
	p.WriteTitle(qb422016)
//line oauth2/template/form.qtpl:82
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:82
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:82
	return qs422016
//line oauth2/template/form.qtpl:82
}

//line oauth2/template/form.qtpl:84
func (p *AppForm) StreamBody(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:84
	qw422016.N().S(` <div class="panel"> `)
//line oauth2/template/form.qtpl:86
	if p.App != nil {
//line oauth2/template/form.qtpl:86
		qw422016.N().S(` <div class="panel-body slim"> <h1 class="mb-10">`)
//line oauth2/template/form.qtpl:88
		qw422016.E().S(p.App.Name)
//line oauth2/template/form.qtpl:88
		qw422016.N().S(`</h1> <div class="mb-10"> <strong>Client ID</strong><br/><span class="muted">`)
//line oauth2/template/form.qtpl:90
		qw422016.E().S(fmt.Sprintf("%x", p.App.ClientID))
//line oauth2/template/form.qtpl:90
		qw422016.N().S(`</span><br/><br/> <strong>Client Secret</strong><br/><span class="muted">`)
//line oauth2/template/form.qtpl:91
		qw422016.E().S(fmt.Sprintf("%x", p.App.ClientSecret))
//line oauth2/template/form.qtpl:91
		qw422016.N().S(`</span></br> </div> <form method="POST" action="`)
//line oauth2/template/form.qtpl:93
		qw422016.E().S(p.App.Endpoint("revoke"))
//line oauth2/template/form.qtpl:93
		qw422016.N().S(`" class="inline-block"> `)
//line oauth2/template/form.qtpl:94
		qw422016.N().S(p.CSRF)
//line oauth2/template/form.qtpl:94
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> <div class="form-field-inline"> <button type="submit" class="btn btn-danger">Revoke Tokens</button> </div> </form> <form method="POST" action="`)
//line oauth2/template/form.qtpl:100
		qw422016.E().S(p.App.Endpoint("reset"))
//line oauth2/template/form.qtpl:100
		qw422016.N().S(`" class="inline-block"> `)
//line oauth2/template/form.qtpl:101
		qw422016.N().S(p.CSRF)
//line oauth2/template/form.qtpl:101
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> <button type="submit" class="btn btn-danger">Reset Secret</button> </form> </div> <form class="panel-body slim" method="POST" action="`)
//line oauth2/template/form.qtpl:106
		qw422016.E().S(p.App.Endpoint())
//line oauth2/template/form.qtpl:106
		qw422016.N().S(`"> `)
//line oauth2/template/form.qtpl:107
		qw422016.N().S(p.CSRF)
//line oauth2/template/form.qtpl:107
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
//line oauth2/template/form.qtpl:109
	} else {
//line oauth2/template/form.qtpl:109
		qw422016.N().S(` <form class="panel-body slim" method="POST" action="/settings/apps"> `)
//line oauth2/template/form.qtpl:111
		qw422016.N().S(p.CSRF)
//line oauth2/template/form.qtpl:111
		qw422016.N().S(` `)
//line oauth2/template/form.qtpl:112
	}
//line oauth2/template/form.qtpl:112
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input type="text" class="form-text" id="name" name="name" value="`)
//line oauth2/template/form.qtpl:115
	qw422016.E().S(p.field("name"))
//line oauth2/template/form.qtpl:115
	qw422016.N().S(`" autocomplete="off"/> `)
//line oauth2/template/form.qtpl:116
	p.StreamError(qw422016, "name")
//line oauth2/template/form.qtpl:116
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="description">Description <small>(optional)</small></label> <textarea class="form-text" id="description" name="description">`)
//line oauth2/template/form.qtpl:120
	qw422016.E().S(p.field("description"))
//line oauth2/template/form.qtpl:120
	qw422016.N().S(`</textarea> </div> <div class="form-field"> <label class="label" for="homepage_uri">Homepage URI</label> <input type="text" class="form-text" id="homepage_uri" name="homepage_uri" value="`)
//line oauth2/template/form.qtpl:124
	qw422016.E().S(p.field("homepage_uri"))
//line oauth2/template/form.qtpl:124
	qw422016.N().S(`" autocomplete="off"/> `)
//line oauth2/template/form.qtpl:125
	p.StreamError(qw422016, "homepage_uri")
//line oauth2/template/form.qtpl:125
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="redirect_uri">Redirect URI</label> <input type="text" class="form-text" id="redirect_uri" name="redirect_uri" value="`)
//line oauth2/template/form.qtpl:129
	qw422016.E().S(p.field("redirect_uri"))
//line oauth2/template/form.qtpl:129
	qw422016.N().S(`" autocomplete="off"/> `)
//line oauth2/template/form.qtpl:130
	p.StreamError(qw422016, "redirect_uri")
//line oauth2/template/form.qtpl:130
	qw422016.N().S(` </div> <div class="form-field"> `)
//line oauth2/template/form.qtpl:133
	if p.App == nil {
//line oauth2/template/form.qtpl:133
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line oauth2/template/form.qtpl:135
	} else {
//line oauth2/template/form.qtpl:135
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line oauth2/template/form.qtpl:137
	}
//line oauth2/template/form.qtpl:137
	qw422016.N().S(` </div> </form> </div> `)
//line oauth2/template/form.qtpl:141
}

//line oauth2/template/form.qtpl:141
func (p *AppForm) WriteBody(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:141
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:141
	p.StreamBody(qw422016)
//line oauth2/template/form.qtpl:141
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:141
}

//line oauth2/template/form.qtpl:141
func (p *AppForm) Body() string {
//line oauth2/template/form.qtpl:141
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:141
	p.WriteBody(qb422016)
//line oauth2/template/form.qtpl:141
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:141
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:141
	return qs422016
//line oauth2/template/form.qtpl:141
}

//line oauth2/template/form.qtpl:143
func (p *AppForm) StreamHeader(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:143
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:144
	if p.App == nil {
//line oauth2/template/form.qtpl:144
		qw422016.N().S(` <a href="/settings/apps" class="back">`)
//line oauth2/template/form.qtpl:145
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line oauth2/template/form.qtpl:145
		qw422016.N().S(`</a> Settings - New App `)
//line oauth2/template/form.qtpl:146
	} else {
//line oauth2/template/form.qtpl:146
		qw422016.N().S(` <a href="/settings/apps" class="back">`)
//line oauth2/template/form.qtpl:147
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line oauth2/template/form.qtpl:147
		qw422016.N().S(`</a> Settings - Edit App `)
//line oauth2/template/form.qtpl:148
	}
//line oauth2/template/form.qtpl:148
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:149
}

//line oauth2/template/form.qtpl:149
func (p *AppForm) WriteHeader(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:149
	p.StreamHeader(qw422016)
//line oauth2/template/form.qtpl:149
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:149
}

//line oauth2/template/form.qtpl:149
func (p *AppForm) Header() string {
//line oauth2/template/form.qtpl:149
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:149
	p.WriteHeader(qb422016)
//line oauth2/template/form.qtpl:149
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:149
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:149
	return qs422016
//line oauth2/template/form.qtpl:149
}

//line oauth2/template/form.qtpl:151
func (p *AppForm) StreamActions(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:151
}

//line oauth2/template/form.qtpl:151
func (p *AppForm) WriteActions(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:151
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:151
	p.StreamActions(qw422016)
//line oauth2/template/form.qtpl:151
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:151
}

//line oauth2/template/form.qtpl:151
func (p *AppForm) Actions() string {
//line oauth2/template/form.qtpl:151
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:151
	p.WriteActions(qb422016)
//line oauth2/template/form.qtpl:151
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:151
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:151
	return qs422016
//line oauth2/template/form.qtpl:151
}

//line oauth2/template/form.qtpl:152
func (p *AppForm) StreamNavigation(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:152
}

//line oauth2/template/form.qtpl:152
func (p *AppForm) WriteNavigation(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:152
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:152
	p.StreamNavigation(qw422016)
//line oauth2/template/form.qtpl:152
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:152
}

//line oauth2/template/form.qtpl:152
func (p *AppForm) Navigation() string {
//line oauth2/template/form.qtpl:152
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:152
	p.WriteNavigation(qb422016)
//line oauth2/template/form.qtpl:152
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:152
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:152
	return qs422016
//line oauth2/template/form.qtpl:152
}

//line oauth2/template/form.qtpl:154
func (p *TokenForm) StreamTitle(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:154
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:155
	if p.Token == nil {
//line oauth2/template/form.qtpl:155
		qw422016.N().S(` Settings - New Token `)
//line oauth2/template/form.qtpl:157
	} else {
//line oauth2/template/form.qtpl:157
		qw422016.N().S(` Settings - Edit Token `)
//line oauth2/template/form.qtpl:159
	}
//line oauth2/template/form.qtpl:159
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:160
}

//line oauth2/template/form.qtpl:160
func (p *TokenForm) WriteTitle(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:160
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:160
	p.StreamTitle(qw422016)
//line oauth2/template/form.qtpl:160
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:160
}

//line oauth2/template/form.qtpl:160
func (p *TokenForm) Title() string {
//line oauth2/template/form.qtpl:160
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:160
	p.WriteTitle(qb422016)
//line oauth2/template/form.qtpl:160
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:160
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:160
	return qs422016
//line oauth2/template/form.qtpl:160
}

//line oauth2/template/form.qtpl:162
func (p *TokenForm) StreamBody(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:162
	qw422016.N().S(` <div class="panel"> `)
//line oauth2/template/form.qtpl:164
	if p.Token != nil {
//line oauth2/template/form.qtpl:164
		qw422016.N().S(` <form method="POST" action="`)
//line oauth2/template/form.qtpl:165
		qw422016.E().S(p.Token.Endpoint("regenerate"))
//line oauth2/template/form.qtpl:165
		qw422016.N().S(`" class="panel-body slim mb-10"> <input type="hidden" name="_method" value="PATCH"/> `)
//line oauth2/template/form.qtpl:167
		qw422016.N().S(p.CSRF)
//line oauth2/template/form.qtpl:167
		qw422016.N().S(` <button type="submit" class="btn btn-danger right">Regenerate</button> </form> <form class="panel-body slim" method="POST" action="`)
//line oauth2/template/form.qtpl:170
		qw422016.E().S(p.Token.Endpoint())
//line oauth2/template/form.qtpl:170
		qw422016.N().S(`"> <input type="hidden" name="_method" value="PATCH"/> `)
//line oauth2/template/form.qtpl:172
	} else {
//line oauth2/template/form.qtpl:172
		qw422016.N().S(` <form class="panel-body slim" method="POST" action="/settings/tokens"> `)
//line oauth2/template/form.qtpl:174
	}
//line oauth2/template/form.qtpl:174
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:175
	qw422016.N().S(p.CSRF)
//line oauth2/template/form.qtpl:175
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input type="text" class="form-text" id="name" name="name" value="`)
//line oauth2/template/form.qtpl:178
	qw422016.E().S(p.field("name"))
//line oauth2/template/form.qtpl:178
	qw422016.N().S(`" autocomplete="off"/> `)
//line oauth2/template/form.qtpl:179
	p.StreamError(qw422016, "name")
//line oauth2/template/form.qtpl:179
	qw422016.N().S(` </div> `)
//line oauth2/template/form.qtpl:181
	for _, res := range oauth2.Resources {
//line oauth2/template/form.qtpl:181
		qw422016.N().S(` <div class="form-field"> <label class="label">`)
//line oauth2/template/form.qtpl:183
		qw422016.E().S(strings.Title(res.String()))
//line oauth2/template/form.qtpl:183
		qw422016.N().S(`</label> `)
//line oauth2/template/form.qtpl:184
		for _, perm := range oauth2.Permissions {
//line oauth2/template/form.qtpl:184
			qw422016.N().S(` `)
//line oauth2/template/form.qtpl:185
			if _, ok := p.Scopes[res.String()+":"+perm.String()]; ok {
//line oauth2/template/form.qtpl:185
				qw422016.N().S(` <label> <input type="checkbox" name="scope[]" value="`)
//line oauth2/template/form.qtpl:187
				qw422016.E().S(res.String())
//line oauth2/template/form.qtpl:187
				qw422016.N().S(`:`)
//line oauth2/template/form.qtpl:187
				qw422016.E().S(perm.String())
//line oauth2/template/form.qtpl:187
				qw422016.N().S(`" checked="true"/> `)
//line oauth2/template/form.qtpl:187
				qw422016.E().S(strings.Title(perm.String()))
//line oauth2/template/form.qtpl:187
				qw422016.N().S(` </label> `)
//line oauth2/template/form.qtpl:189
			} else {
//line oauth2/template/form.qtpl:189
				qw422016.N().S(` <label> <input type="checkbox" name="scope[]" value="`)
//line oauth2/template/form.qtpl:191
				qw422016.E().S(res.String())
//line oauth2/template/form.qtpl:191
				qw422016.N().S(`:`)
//line oauth2/template/form.qtpl:191
				qw422016.E().S(perm.String())
//line oauth2/template/form.qtpl:191
				qw422016.N().S(`"/> `)
//line oauth2/template/form.qtpl:191
				qw422016.E().S(strings.Title(perm.String()))
//line oauth2/template/form.qtpl:191
				qw422016.N().S(` </label> `)
//line oauth2/template/form.qtpl:193
			}
//line oauth2/template/form.qtpl:193
			qw422016.N().S(` `)
//line oauth2/template/form.qtpl:194
		}
//line oauth2/template/form.qtpl:194
		qw422016.N().S(` </div> `)
//line oauth2/template/form.qtpl:196
	}
//line oauth2/template/form.qtpl:196
	qw422016.N().S(` <div class="form-field"> `)
//line oauth2/template/form.qtpl:198
	if p.Token == nil {
//line oauth2/template/form.qtpl:198
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line oauth2/template/form.qtpl:200
	} else {
//line oauth2/template/form.qtpl:200
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line oauth2/template/form.qtpl:202
	}
//line oauth2/template/form.qtpl:202
	qw422016.N().S(` </div> </form> </div> `)
//line oauth2/template/form.qtpl:206
}

//line oauth2/template/form.qtpl:206
func (p *TokenForm) WriteBody(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:206
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:206
	p.StreamBody(qw422016)
//line oauth2/template/form.qtpl:206
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:206
}

//line oauth2/template/form.qtpl:206
func (p *TokenForm) Body() string {
//line oauth2/template/form.qtpl:206
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:206
	p.WriteBody(qb422016)
//line oauth2/template/form.qtpl:206
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:206
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:206
	return qs422016
//line oauth2/template/form.qtpl:206
}

//line oauth2/template/form.qtpl:208
func (p *TokenForm) StreamHeader(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:208
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:209
	if p.Token == nil {
//line oauth2/template/form.qtpl:209
		qw422016.N().S(` <a href="/settings/tokens" class="back">`)
//line oauth2/template/form.qtpl:210
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line oauth2/template/form.qtpl:210
		qw422016.N().S(`</a> Settings - New Token `)
//line oauth2/template/form.qtpl:211
	} else {
//line oauth2/template/form.qtpl:211
		qw422016.N().S(` <a href="/settings/tokens" class="back">`)
//line oauth2/template/form.qtpl:212
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line oauth2/template/form.qtpl:212
		qw422016.N().S(`</a> Settings - Edit Token `)
//line oauth2/template/form.qtpl:213
	}
//line oauth2/template/form.qtpl:213
	qw422016.N().S(` `)
//line oauth2/template/form.qtpl:214
}

//line oauth2/template/form.qtpl:214
func (p *TokenForm) WriteHeader(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:214
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:214
	p.StreamHeader(qw422016)
//line oauth2/template/form.qtpl:214
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:214
}

//line oauth2/template/form.qtpl:214
func (p *TokenForm) Header() string {
//line oauth2/template/form.qtpl:214
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:214
	p.WriteHeader(qb422016)
//line oauth2/template/form.qtpl:214
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:214
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:214
	return qs422016
//line oauth2/template/form.qtpl:214
}

//line oauth2/template/form.qtpl:216
func (p *TokenForm) StreamActions(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:216
}

//line oauth2/template/form.qtpl:216
func (p *TokenForm) WriteActions(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:216
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:216
	p.StreamActions(qw422016)
//line oauth2/template/form.qtpl:216
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:216
}

//line oauth2/template/form.qtpl:216
func (p *TokenForm) Actions() string {
//line oauth2/template/form.qtpl:216
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:216
	p.WriteActions(qb422016)
//line oauth2/template/form.qtpl:216
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:216
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:216
	return qs422016
//line oauth2/template/form.qtpl:216
}

//line oauth2/template/form.qtpl:217
func (p *TokenForm) StreamNavigation(qw422016 *qt422016.Writer) {
//line oauth2/template/form.qtpl:217
}

//line oauth2/template/form.qtpl:217
func (p *TokenForm) WriteNavigation(qq422016 qtio422016.Writer) {
//line oauth2/template/form.qtpl:217
	qw422016 := qt422016.AcquireWriter(qq422016)
//line oauth2/template/form.qtpl:217
	p.StreamNavigation(qw422016)
//line oauth2/template/form.qtpl:217
	qt422016.ReleaseWriter(qw422016)
//line oauth2/template/form.qtpl:217
}

//line oauth2/template/form.qtpl:217
func (p *TokenForm) Navigation() string {
//line oauth2/template/form.qtpl:217
	qb422016 := qt422016.AcquireByteBuffer()
//line oauth2/template/form.qtpl:217
	p.WriteNavigation(qb422016)
//line oauth2/template/form.qtpl:217
	qs422016 := string(qb422016.B)
//line oauth2/template/form.qtpl:217
	qt422016.ReleaseByteBuffer(qb422016)
//line oauth2/template/form.qtpl:217
	return qs422016
//line oauth2/template/form.qtpl:217
}

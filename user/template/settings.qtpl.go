// Code generated by qtc from "settings.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line user/template/settings.qtpl:2
package template

//line user/template/settings.qtpl:2
import (
	"djinn-ci.com/provider"
	"djinn-ci.com/template"
)

//line user/template/settings.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line user/template/settings.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line user/template/settings.qtpl:9
type Settings struct {
	template.BasePage
	template.Form

	Providers []*provider.Provider
	Section   template.Section
}

//line user/template/settings.qtpl:19
func (p *Settings) StreamTitle(qw422016 *qt422016.Writer) {
//line user/template/settings.qtpl:19
	qw422016.N().S(` Settings`)
//line user/template/settings.qtpl:20
	if p.Section != nil {
//line user/template/settings.qtpl:20
		qw422016.N().S(` - `)
//line user/template/settings.qtpl:20
		p.Section.StreamTitle(qw422016)
//line user/template/settings.qtpl:20
	}
//line user/template/settings.qtpl:20
	qw422016.N().S(` - Djinn CI `)
//line user/template/settings.qtpl:21
}

//line user/template/settings.qtpl:21
func (p *Settings) WriteTitle(qq422016 qtio422016.Writer) {
//line user/template/settings.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/settings.qtpl:21
	p.StreamTitle(qw422016)
//line user/template/settings.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line user/template/settings.qtpl:21
}

//line user/template/settings.qtpl:21
func (p *Settings) Title() string {
//line user/template/settings.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/settings.qtpl:21
	p.WriteTitle(qb422016)
//line user/template/settings.qtpl:21
	qs422016 := string(qb422016.B)
//line user/template/settings.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/settings.qtpl:21
	return qs422016
//line user/template/settings.qtpl:21
}

//line user/template/settings.qtpl:23
func (p *Settings) streamrenderProvider(qw422016 *qt422016.Writer, prv *provider.Provider) {
//line user/template/settings.qtpl:23
	qw422016.N().S(` `)
//line user/template/settings.qtpl:24
	if !prv.Connected {
//line user/template/settings.qtpl:24
		qw422016.N().S(` <a href="`)
//line user/template/settings.qtpl:25
		qw422016.E().S(prv.AuthURL)
//line user/template/settings.qtpl:25
		qw422016.N().S(`" class="provider-btn provider-`)
//line user/template/settings.qtpl:25
		qw422016.E().S(prv.Name)
//line user/template/settings.qtpl:25
		qw422016.N().S(`"> `)
//line user/template/settings.qtpl:26
		switch prv.Name {
//line user/template/settings.qtpl:27
		case "github":
//line user/template/settings.qtpl:27
			qw422016.N().S(` `)
//line user/template/settings.qtpl:28
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385 0.6 0.113 0.82-0.258 0.82-0.577 0-0.285-0.010-1.040-0.015-2.040-3.338 0.724-4.042-1.61-4.042-1.61-0.546-1.385-1.335-1.755-1.335-1.755-1.087-0.744 0.084-0.729 0.084-0.729 1.205 0.084 1.838 1.236 1.838 1.236 1.070 1.835 2.809 1.305 3.495 0.998 0.108-0.776 0.417-1.305 0.76-1.605-2.665-0.3-5.466-1.332-5.466-5.93 0-1.31 0.465-2.38 1.235-3.22-0.135-0.303-0.54-1.523 0.105-3.176 0 0 1.005-0.322 3.3 1.23 0.96-0.267 1.98-0.399 3-0.405 1.020 0.006 2.040 0.138 3 0.405 2.28-1.552 3.285-1.23 3.285-1.23 0.645 1.653 0.24 2.873 0.12 3.176 0.765 0.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92 0.42 0.36 0.81 1.096 0.81 2.22 0 1.606-0.015 2.896-0.015 3.286 0 0.315 0.21 0.69 0.825 0.57 4.801-1.574 8.236-6.074 8.236-11.369 0-6.627-5.373-12-12-12z"></path>
</svg>
`)
//line user/template/settings.qtpl:28
			qw422016.N().S(` `)
//line user/template/settings.qtpl:29
		case "gitlab":
//line user/template/settings.qtpl:29
			qw422016.N().S(` `)
//line user/template/settings.qtpl:30
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M23.955 13.587l-1.342-4.135-2.664-8.189c-0.135-0.423-0.73-0.423-0.867 0l-2.664 8.187h-8.836l-2.663-8.187c-0.136-0.423-0.734-0.423-0.869-0.003l-2.664 8.189-1.342 4.138c-0.121 0.375 0.014 0.789 0.331 1.023l11.625 8.444 11.625-8.443c0.318-0.235 0.453-0.647 0.33-1.024z"></path>
</svg>
`)
//line user/template/settings.qtpl:30
			qw422016.N().S(` `)
//line user/template/settings.qtpl:31
		}
//line user/template/settings.qtpl:31
		qw422016.N().S(` <span>Connect</span> </a> `)
//line user/template/settings.qtpl:34
	} else {
//line user/template/settings.qtpl:34
		qw422016.N().S(` <form method="POST" action="/oauth/`)
//line user/template/settings.qtpl:35
		qw422016.E().S(prv.Name)
//line user/template/settings.qtpl:35
		qw422016.N().S(`" class="inline-block"> `)
//line user/template/settings.qtpl:36
		qw422016.N().V(p.CSRF)
//line user/template/settings.qtpl:36
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <button type="submit" class="provider-btn provider-`)
//line user/template/settings.qtpl:38
		qw422016.E().S(prv.Name)
//line user/template/settings.qtpl:38
		qw422016.N().S(`"> `)
//line user/template/settings.qtpl:39
		switch prv.Name {
//line user/template/settings.qtpl:40
		case "github":
//line user/template/settings.qtpl:40
			qw422016.N().S(` `)
//line user/template/settings.qtpl:41
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385 0.6 0.113 0.82-0.258 0.82-0.577 0-0.285-0.010-1.040-0.015-2.040-3.338 0.724-4.042-1.61-4.042-1.61-0.546-1.385-1.335-1.755-1.335-1.755-1.087-0.744 0.084-0.729 0.084-0.729 1.205 0.084 1.838 1.236 1.838 1.236 1.070 1.835 2.809 1.305 3.495 0.998 0.108-0.776 0.417-1.305 0.76-1.605-2.665-0.3-5.466-1.332-5.466-5.93 0-1.31 0.465-2.38 1.235-3.22-0.135-0.303-0.54-1.523 0.105-3.176 0 0 1.005-0.322 3.3 1.23 0.96-0.267 1.98-0.399 3-0.405 1.020 0.006 2.040 0.138 3 0.405 2.28-1.552 3.285-1.23 3.285-1.23 0.645 1.653 0.24 2.873 0.12 3.176 0.765 0.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92 0.42 0.36 0.81 1.096 0.81 2.22 0 1.606-0.015 2.896-0.015 3.286 0 0.315 0.21 0.69 0.825 0.57 4.801-1.574 8.236-6.074 8.236-11.369 0-6.627-5.373-12-12-12z"></path>
</svg>
`)
//line user/template/settings.qtpl:41
			qw422016.N().S(` `)
//line user/template/settings.qtpl:42
		case "gitlab":
//line user/template/settings.qtpl:42
			qw422016.N().S(` `)
//line user/template/settings.qtpl:43
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M23.955 13.587l-1.342-4.135-2.664-8.189c-0.135-0.423-0.73-0.423-0.867 0l-2.664 8.187h-8.836l-2.663-8.187c-0.136-0.423-0.734-0.423-0.869-0.003l-2.664 8.189-1.342 4.138c-0.121 0.375 0.014 0.789 0.331 1.023l11.625 8.444 11.625-8.443c0.318-0.235 0.453-0.647 0.33-1.024z"></path>
</svg>
`)
//line user/template/settings.qtpl:43
			qw422016.N().S(` `)
//line user/template/settings.qtpl:44
		}
//line user/template/settings.qtpl:44
		qw422016.N().S(` <span>Disconnect</span> </button> </form> `)
//line user/template/settings.qtpl:48
	}
//line user/template/settings.qtpl:48
	qw422016.N().S(` `)
//line user/template/settings.qtpl:49
}

//line user/template/settings.qtpl:49
func (p *Settings) writerenderProvider(qq422016 qtio422016.Writer, prv *provider.Provider) {
//line user/template/settings.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/settings.qtpl:49
	p.streamrenderProvider(qw422016, prv)
//line user/template/settings.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line user/template/settings.qtpl:49
}

//line user/template/settings.qtpl:49
func (p *Settings) renderProvider(prv *provider.Provider) string {
//line user/template/settings.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/settings.qtpl:49
	p.writerenderProvider(qb422016, prv)
//line user/template/settings.qtpl:49
	qs422016 := string(qb422016.B)
//line user/template/settings.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/settings.qtpl:49
	return qs422016
//line user/template/settings.qtpl:49
}

//line user/template/settings.qtpl:51
func (p *Settings) StreamBody(qw422016 *qt422016.Writer) {
//line user/template/settings.qtpl:51
	qw422016.N().S(` `)
//line user/template/settings.qtpl:52
	if p.Section != nil {
//line user/template/settings.qtpl:52
		qw422016.N().S(` `)
//line user/template/settings.qtpl:53
		p.Section.StreamBody(qw422016)
//line user/template/settings.qtpl:53
		qw422016.N().S(` `)
//line user/template/settings.qtpl:54
	} else {
//line user/template/settings.qtpl:54
		qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> `)
//line user/template/settings.qtpl:57
		if !p.User.Verified {
//line user/template/settings.qtpl:57
			qw422016.N().S(` <form method="POST" action="/settings/verify"> <h2>Verify account</h2> `)
//line user/template/settings.qtpl:60
			qw422016.N().V(p.CSRF)
//line user/template/settings.qtpl:60
			qw422016.N().S(` <div class="form-field"> <button type="submit" class="btn btn-primary">Send verification email</button> </div> </form> <div class="separator"></div> `)
//line user/template/settings.qtpl:66
		}
//line user/template/settings.qtpl:66
		qw422016.N().S(` `)
//line user/template/settings.qtpl:67
		if len(p.Providers) > 0 {
//line user/template/settings.qtpl:67
			qw422016.N().S(` <h2>Connected accounts</h2> <div class="providers"> `)
//line user/template/settings.qtpl:70
			for _, prv := range p.Providers {
//line user/template/settings.qtpl:70
				qw422016.N().S(` `)
//line user/template/settings.qtpl:71
				p.streamrenderProvider(qw422016, prv)
//line user/template/settings.qtpl:71
				qw422016.N().S(` `)
//line user/template/settings.qtpl:72
			}
//line user/template/settings.qtpl:72
			qw422016.N().S(` </div> `)
//line user/template/settings.qtpl:74
		}
//line user/template/settings.qtpl:74
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="/settings/cleanup"> `)
//line user/template/settings.qtpl:77
		qw422016.N().V(p.CSRF)
//line user/template/settings.qtpl:77
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> <div class="form-field"> <label class="form-option"> `)
//line user/template/settings.qtpl:81
		if p.User.Cleanup {
//line user/template/settings.qtpl:81
			qw422016.N().S(` <input class="form-selector" type="checkbox" name="cleanup" checked="true"/> `)
//line user/template/settings.qtpl:83
		} else {
//line user/template/settings.qtpl:83
			qw422016.N().S(` <input class="form-selector" type="checkbox" name="cleanup"/> `)
//line user/template/settings.qtpl:85
		}
//line user/template/settings.qtpl:85
		qw422016.N().S(` <strong>Cleanup artifacts</strong> <span class="left">`)
//line user/template/settings.qtpl:87
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>bin</title>
<path d="M18.984 3.984v2.016h-13.969v-2.016h3.469l1.031-0.984h4.969l1.031 0.984h3.469zM6 18.984v-12h12v12c0 1.078-0.938 2.016-2.016 2.016h-7.969c-1.078 0-2.016-0.938-2.016-2.016z"></path>
</svg>
`)
//line user/template/settings.qtpl:87
		qw422016.N().S(`</span> <div class="form-desc">Clean up old artifacts</div> </label> </div> <div class="form-field"> <button type="submit" class="btn btn-primary">Update</button> </div> </form> <div class="separator"></div> <form method="POST" action="/settings/email"> <h2>Change email</h2> `)
//line user/template/settings.qtpl:98
		qw422016.N().V(p.CSRF)
//line user/template/settings.qtpl:98
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> <div class="form-field"> <label class="label" for="email">Email</label> <input class="form-text" type="text" id="email" name="email" value="`)
//line user/template/settings.qtpl:102
		qw422016.E().S(p.Fields["email"])
//line user/template/settings.qtpl:102
		qw422016.N().S(`" autocomplete="off"/> `)
//line user/template/settings.qtpl:103
		p.StreamError(qw422016, "email")
//line user/template/settings.qtpl:103
		qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="verify_password">Verify Password</label> <input class="form-text" type="password" id="verify_password" name="verify_password" autocomplete="off"/> `)
//line user/template/settings.qtpl:108
		p.StreamError(qw422016, "email_verify_password")
//line user/template/settings.qtpl:108
		qw422016.N().S(` </div> <div class="form-field"> <button type="submit" class="btn btn-primary">Update</button> </div> </form> <div class="separator"></div> <form method="POST" action="/settings/password"> <h2>Change password</h2> <input type="hidden" name="_method" value="PATCH"/> `)
//line user/template/settings.qtpl:118
		qw422016.N().V(p.CSRF)
//line user/template/settings.qtpl:118
		qw422016.N().S(` <div class="form-field"> <label class="label" for="old_password">Old Password</label> <input class="form-text" type="password" id="old_password" name="old_password" autocomplete="off"/> `)
//line user/template/settings.qtpl:122
		p.StreamError(qw422016, "old_password")
//line user/template/settings.qtpl:122
		qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="new_password">New Password</label> <input class="form-text" type="password" id="new_password" name="new_password" autocomplete="off"/> `)
//line user/template/settings.qtpl:127
		p.StreamError(qw422016, "new_password")
//line user/template/settings.qtpl:127
		qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="verify_password">Verify Password</label> <input class="form-text" type="password" id="verify_password" name="verify_password" autocomplete="off"/> `)
//line user/template/settings.qtpl:132
		p.StreamError(qw422016, "pass_verify_password")
//line user/template/settings.qtpl:132
		qw422016.N().S(` </div> <div class="form-field"> <button type="submit" class="btn btn-primary">Update</button> </div> </form> <div class="separator"></div> <form method="POST" action="/settings/delete"> <h2>Delete account</h2> `)
//line user/template/settings.qtpl:141
		qw422016.N().V(p.CSRF)
//line user/template/settings.qtpl:141
		qw422016.N().S(` <div class="form-field"> <label class="label" for="delete_password">Verify Password</label> <input class="form-text" type="password" id="delete_password" name="delete_password" autocomplete="off"/> `)
//line user/template/settings.qtpl:145
		p.StreamError(qw422016, "delete_password")
//line user/template/settings.qtpl:145
		qw422016.N().S(` </div> <div class="form-field"> <button type="submit" class="btn btn-danger">Delete</button> </div> </form> </div> </div> `)
//line user/template/settings.qtpl:153
	}
//line user/template/settings.qtpl:153
	qw422016.N().S(` `)
//line user/template/settings.qtpl:154
}

//line user/template/settings.qtpl:154
func (p *Settings) WriteBody(qq422016 qtio422016.Writer) {
//line user/template/settings.qtpl:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/settings.qtpl:154
	p.StreamBody(qw422016)
//line user/template/settings.qtpl:154
	qt422016.ReleaseWriter(qw422016)
//line user/template/settings.qtpl:154
}

//line user/template/settings.qtpl:154
func (p *Settings) Body() string {
//line user/template/settings.qtpl:154
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/settings.qtpl:154
	p.WriteBody(qb422016)
//line user/template/settings.qtpl:154
	qs422016 := string(qb422016.B)
//line user/template/settings.qtpl:154
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/settings.qtpl:154
	return qs422016
//line user/template/settings.qtpl:154
}

//line user/template/settings.qtpl:156
func (p *Settings) StreamHeader(qw422016 *qt422016.Writer) {
//line user/template/settings.qtpl:156
	qw422016.N().S(` `)
//line user/template/settings.qtpl:157
	if p.Section != nil {
//line user/template/settings.qtpl:157
		qw422016.N().S(` `)
//line user/template/settings.qtpl:158
		p.Section.StreamHeader(qw422016)
//line user/template/settings.qtpl:158
		qw422016.N().S(` `)
//line user/template/settings.qtpl:159
	} else {
//line user/template/settings.qtpl:159
		qw422016.N().S(` Settings `)
//line user/template/settings.qtpl:161
	}
//line user/template/settings.qtpl:161
	qw422016.N().S(` `)
//line user/template/settings.qtpl:162
}

//line user/template/settings.qtpl:162
func (p *Settings) WriteHeader(qq422016 qtio422016.Writer) {
//line user/template/settings.qtpl:162
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/settings.qtpl:162
	p.StreamHeader(qw422016)
//line user/template/settings.qtpl:162
	qt422016.ReleaseWriter(qw422016)
//line user/template/settings.qtpl:162
}

//line user/template/settings.qtpl:162
func (p *Settings) Header() string {
//line user/template/settings.qtpl:162
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/settings.qtpl:162
	p.WriteHeader(qb422016)
//line user/template/settings.qtpl:162
	qs422016 := string(qb422016.B)
//line user/template/settings.qtpl:162
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/settings.qtpl:162
	return qs422016
//line user/template/settings.qtpl:162
}

//line user/template/settings.qtpl:164
func (p *Settings) StreamActions(qw422016 *qt422016.Writer) {
//line user/template/settings.qtpl:164
	qw422016.N().S(` `)
//line user/template/settings.qtpl:165
	if p.Section != nil {
//line user/template/settings.qtpl:165
		qw422016.N().S(` `)
//line user/template/settings.qtpl:166
		p.Section.StreamActions(qw422016)
//line user/template/settings.qtpl:166
		qw422016.N().S(` `)
//line user/template/settings.qtpl:167
	}
//line user/template/settings.qtpl:167
	qw422016.N().S(` `)
//line user/template/settings.qtpl:168
}

//line user/template/settings.qtpl:168
func (p *Settings) WriteActions(qq422016 qtio422016.Writer) {
//line user/template/settings.qtpl:168
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/settings.qtpl:168
	p.StreamActions(qw422016)
//line user/template/settings.qtpl:168
	qt422016.ReleaseWriter(qw422016)
//line user/template/settings.qtpl:168
}

//line user/template/settings.qtpl:168
func (p *Settings) Actions() string {
//line user/template/settings.qtpl:168
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/settings.qtpl:168
	p.WriteActions(qb422016)
//line user/template/settings.qtpl:168
	qs422016 := string(qb422016.B)
//line user/template/settings.qtpl:168
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/settings.qtpl:168
	return qs422016
//line user/template/settings.qtpl:168
}

//line user/template/settings.qtpl:170
func (p *Settings) StreamNavigation(qw422016 *qt422016.Writer) {
//line user/template/settings.qtpl:170
	qw422016.N().S(` <li> `)
//line user/template/settings.qtpl:172
	if p.URL.Path == "/settings" {
//line user/template/settings.qtpl:172
		qw422016.N().S(` <a href="/settings" class="active">`)
//line user/template/settings.qtpl:173
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>user</title>
<path d="M12 14.016c2.672 0 8.016 1.313 8.016 3.984v2.016h-16.031v-2.016c0-2.672 5.344-3.984 8.016-3.984zM12 12c-2.203 0-3.984-1.781-3.984-3.984s1.781-4.031 3.984-4.031 3.984 1.828 3.984 4.031-1.781 3.984-3.984 3.984z"></path>
</svg>
`)
//line user/template/settings.qtpl:173
		qw422016.N().S(`<span>Account</span></a> `)
//line user/template/settings.qtpl:174
	} else {
//line user/template/settings.qtpl:174
		qw422016.N().S(` <a href="/settings">`)
//line user/template/settings.qtpl:175
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>user</title>
<path d="M12 14.016c2.672 0 8.016 1.313 8.016 3.984v2.016h-16.031v-2.016c0-2.672 5.344-3.984 8.016-3.984zM12 12c-2.203 0-3.984-1.781-3.984-3.984s1.781-4.031 3.984-4.031 3.984 1.828 3.984 4.031-1.781 3.984-3.984 3.984z"></path>
</svg>
`)
//line user/template/settings.qtpl:175
		qw422016.N().S(`<span>Account</span></a> `)
//line user/template/settings.qtpl:176
	}
//line user/template/settings.qtpl:176
	qw422016.N().S(` </li> <li> `)
//line user/template/settings.qtpl:179
	if p.URL.Path == "/settings/apps" {
//line user/template/settings.qtpl:179
		qw422016.N().S(` <a href="/settings/apps" class="active">`)
//line user/template/settings.qtpl:180
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9.984 3.984l2.016 2.016h8.016c1.078 0 1.969 0.938 1.969 2.016v9.984c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h6z"></path>
</svg>
`)
//line user/template/settings.qtpl:180
		qw422016.N().S(`<span>OAuth Apps</span></a> `)
//line user/template/settings.qtpl:181
	} else {
//line user/template/settings.qtpl:181
		qw422016.N().S(` <a href="/settings/apps">`)
//line user/template/settings.qtpl:182
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9.984 3.984l2.016 2.016h8.016c1.078 0 1.969 0.938 1.969 2.016v9.984c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h6z"></path>
</svg>
`)
//line user/template/settings.qtpl:182
		qw422016.N().S(`<span>OAuth Apps</span></a> `)
//line user/template/settings.qtpl:183
	}
//line user/template/settings.qtpl:183
	qw422016.N().S(` </li> <li> `)
//line user/template/settings.qtpl:186
	if p.URL.Path == "/settings/tokens" || p.URL.Path == "/settings/tokens/create" {
//line user/template/settings.qtpl:186
		qw422016.N().S(` <a href="/settings/tokens" class="active">`)
//line user/template/settings.qtpl:187
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
//line user/template/settings.qtpl:187
		qw422016.N().S(`<span>Access Tokens</span></a> `)
//line user/template/settings.qtpl:188
	} else {
//line user/template/settings.qtpl:188
		qw422016.N().S(` <a href="/settings/tokens">`)
//line user/template/settings.qtpl:189
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
//line user/template/settings.qtpl:189
		qw422016.N().S(`<span>Access Tokens</span></a> `)
//line user/template/settings.qtpl:190
	}
//line user/template/settings.qtpl:190
	qw422016.N().S(` </li> <li> `)
//line user/template/settings.qtpl:193
	if p.URL.Path == "/settings/connections" {
//line user/template/settings.qtpl:193
		qw422016.N().S(` <a href="/settings/connections" class="active">`)
//line user/template/settings.qtpl:194
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>apps</title>
<path d="M15.984 20.016v-4.031h4.031v4.031h-4.031zM15.984 14.016v-4.031h4.031v4.031h-4.031zM9.984 8.016v-4.031h4.031v4.031h-4.031zM15.984 3.984h4.031v4.031h-4.031v-4.031zM9.984 14.016v-4.031h4.031v4.031h-4.031zM3.984 14.016v-4.031h4.031v4.031h-4.031zM3.984 20.016v-4.031h4.031v4.031h-4.031zM9.984 20.016v-4.031h4.031v4.031h-4.031zM3.984 8.016v-4.031h4.031v4.031h-4.031z"></path>
</svg>
`)
//line user/template/settings.qtpl:194
		qw422016.N().S(`<span>Authorized OAuth apps</span></a> `)
//line user/template/settings.qtpl:195
	} else {
//line user/template/settings.qtpl:195
		qw422016.N().S(` <a href="/settings/connections">`)
//line user/template/settings.qtpl:196
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>apps</title>
<path d="M15.984 20.016v-4.031h4.031v4.031h-4.031zM15.984 14.016v-4.031h4.031v4.031h-4.031zM9.984 8.016v-4.031h4.031v4.031h-4.031zM15.984 3.984h4.031v4.031h-4.031v-4.031zM9.984 14.016v-4.031h4.031v4.031h-4.031zM3.984 14.016v-4.031h4.031v4.031h-4.031zM3.984 20.016v-4.031h4.031v4.031h-4.031zM9.984 20.016v-4.031h4.031v4.031h-4.031zM3.984 8.016v-4.031h4.031v4.031h-4.031z"></path>
</svg>
`)
//line user/template/settings.qtpl:196
		qw422016.N().S(`<span>Authorized OAuth apps</span></a> `)
//line user/template/settings.qtpl:197
	}
//line user/template/settings.qtpl:197
	qw422016.N().S(` </li> `)
//line user/template/settings.qtpl:199
}

//line user/template/settings.qtpl:199
func (p *Settings) WriteNavigation(qq422016 qtio422016.Writer) {
//line user/template/settings.qtpl:199
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/settings.qtpl:199
	p.StreamNavigation(qw422016)
//line user/template/settings.qtpl:199
	qt422016.ReleaseWriter(qw422016)
//line user/template/settings.qtpl:199
}

//line user/template/settings.qtpl:199
func (p *Settings) Navigation() string {
//line user/template/settings.qtpl:199
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/settings.qtpl:199
	p.WriteNavigation(qb422016)
//line user/template/settings.qtpl:199
	qs422016 := string(qb422016.B)
//line user/template/settings.qtpl:199
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/settings.qtpl:199
	return qs422016
//line user/template/settings.qtpl:199
}

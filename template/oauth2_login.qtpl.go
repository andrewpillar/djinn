// Code generated by qtc from "oauth2_login.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/oauth2_login.qtpl:2
package template

//line template/oauth2_login.qtpl:2
import (
	"bytes"
	"strings"

	"djinn-ci.com/auth"
	"djinn-ci.com/oauth2"
	"djinn-ci.com/template/form"
)

//line template/oauth2_login.qtpl:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/oauth2_login.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/oauth2_login.qtpl:13
type Oauth2Login struct {
	*form.Form

	User   *auth.User
	Author *auth.User

	Name        string
	ClientID    string
	RedirectURI string
	State       string
	Scope       oauth2.Scope
}

func buildPermissionList(perm oauth2.Permission) string {
	var buf bytes.Buffer

	perms := perm.Expand()
	l := len(perms) - 1

	for i, perm := range perms {
		buf.WriteString(strings.Title(perm.String()))

		if i+1 == l {
			buf.WriteString(", and")
		}
		if i != l {
			buf.WriteString(", ")
		}
	}
	return buf.String()
}

//line template/oauth2_login.qtpl:47
func (p *Oauth2Login) StreamTitle(qw422016 *qt422016.Writer) {
//line template/oauth2_login.qtpl:47
	qw422016.N().S(`Authenticate application`)
//line template/oauth2_login.qtpl:47
}

//line template/oauth2_login.qtpl:47
func (p *Oauth2Login) WriteTitle(qq422016 qtio422016.Writer) {
//line template/oauth2_login.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/oauth2_login.qtpl:47
	p.StreamTitle(qw422016)
//line template/oauth2_login.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line template/oauth2_login.qtpl:47
}

//line template/oauth2_login.qtpl:47
func (p *Oauth2Login) Title() string {
//line template/oauth2_login.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
//line template/oauth2_login.qtpl:47
	p.WriteTitle(qb422016)
//line template/oauth2_login.qtpl:47
	qs422016 := string(qb422016.B)
//line template/oauth2_login.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
//line template/oauth2_login.qtpl:47
	return qs422016
//line template/oauth2_login.qtpl:47
}

//line template/oauth2_login.qtpl:49
func (p *Oauth2Login) StreamFooter(qw422016 *qt422016.Writer) {
//line template/oauth2_login.qtpl:49
	qw422016.N().S(` <style type="text/css">`)
//line template/oauth2_login.qtpl:50
	qw422016.N().S(`*{margin:0;padding:0}body{font-family:sans-serif;font-size:14px;background:#383e51;color:#fff}a{text-decoration:none;cursor:pointer}a:hover{text-decoration:underline}button{cursor:pointer}h1{font-weight:400}.btn{border:none;border-radius:3px;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";padding:10px;padding-left:15px;padding-right:15px;color:#fff}.btn svg{fill:#fff}.btn:hover{text-decoration:none}.btn:disabled{cursor:not-allowed;background:#b2b2b2!important}.btn-primary{background:#61a0ea}.btn-primary:hover{background:#5090d9}.btn-danger{background:#de4141}.btn-danger:hover{background:#cd3030}.logo .left{display:inline-block;width:0;height:0;border-style:solid;border-width:25px 0 0 15px;border-color:transparent transparent transparent #fff}.logo .right{display:inline-block;width:0;height:0;border-style:solid;border-width:0 0 25px 15px;border-color:transparent transparent #fff transparent}.auth-page a{color:#66c9ff}.auth-page .auth-form{margin:0 auto;margin-top:150px;max-width:400px;padding:20px}.auth-page .auth-form .auth-header{margin-bottom:20px;text-align:center}.auth-page .auth-form .auth-header .logo{margin:0 auto;margin-bottom:20px;width:0}.auth-page .auth-form .auth-header .logo .handle{margin-left:-20px;border-style:solid;border-width:5px 0 12px 10px;border-color:transparent transparent transparent #fff}.auth-page .auth-form .auth-header .logo .lid{margin-bottom:-30px;margin-left:5px;border-style:solid;border-width:5px 0 12px 10px;border-color:transparent transparent transparent #fff}.auth-page .auth-form .auth-header .logo .lantern{margin-left:-25px;border-style:solid;border-width:25px 25px 75px 0;border-color:transparent #fff transparent transparent}.auth-page .auth-form .form-error{margin-top:3px;display:block;color:#e74848;min-height:17px}.auth-page .auth-form .form-field{margin-top:10px;width:100%}.auth-page .auth-form .form-field label{margin-bottom:3px;display:block}.auth-page .auth-form .form-field .form-text{box-sizing:border-box;width:100%;font-family:sans-serif;font-size:14px;padding:7px;outline:0;border:solid 1px rgba(255,255,255,.3);border-radius:3px;background:rgba(0,0,0,.3);color:#fff}.auth-page .auth-form .form-field .form-text:focus{border:solid 1px rgba(255,255,255,.5)}.auth-page .auth-form .form-field .btn{color:#fff;width:100%;display:block;text-align:center;box-sizing:border-box}.provider-btn svg{margin-right:5px;fill:#fff;vertical-align:middle}.provider-btn span{display:inline-block;vertical-align:middle}.provider-btn:hover{text-decoration:none}.provider-github{background:#24292e}.provider-github:hover{background:#353a3f}.provider-gitlab{background:#fa7035}.provider-gitlab:hover{background:#e65328}.alert{margin-top:15px;overflow:auto;padding:15px;border-radius:3px;text-align:left}.alert .alert-message{float:left;color:rgba(0,0,0,.6)}.alert a.alert-close{float:right;display:inline-block}.alert a.alert-close svg{width:15px;height:15px;fill:rgba(0,0,0,.4)}.alert a:hover svg{fill:rgba(0,0,0,.5)}.alert-success{background:#caf5ca;border-bottom:solid 1px #a0dfa0}.alert-warn{background:#fff3cd;border-bottom:solid 1px #d9c995}.alert-danger{background:#ffd4d4;border-bottom:solid 1px #e19e9e}.scope-list h3{margin-bottom:15px}.scope-list .scope-item{margin-top:15px;overflow:auto;border-top:solid 1px rgba(255,255,255,.4);padding:15px}.scope-list .scope-item svg{display:inline-block;margin-right:15px;float:left;fill:rgba(255,255,255,.4)}.scope-list .scope-item span{display:inline-block}.scope-list .scope-item span strong{display:block}`)
//line template/oauth2_login.qtpl:50
	qw422016.N().S(`</style> `)
//line template/oauth2_login.qtpl:51
}

//line template/oauth2_login.qtpl:51
func (p *Oauth2Login) WriteFooter(qq422016 qtio422016.Writer) {
//line template/oauth2_login.qtpl:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/oauth2_login.qtpl:51
	p.StreamFooter(qw422016)
//line template/oauth2_login.qtpl:51
	qt422016.ReleaseWriter(qw422016)
//line template/oauth2_login.qtpl:51
}

//line template/oauth2_login.qtpl:51
func (p *Oauth2Login) Footer() string {
//line template/oauth2_login.qtpl:51
	qb422016 := qt422016.AcquireByteBuffer()
//line template/oauth2_login.qtpl:51
	p.WriteFooter(qb422016)
//line template/oauth2_login.qtpl:51
	qs422016 := string(qb422016.B)
//line template/oauth2_login.qtpl:51
	qt422016.ReleaseByteBuffer(qb422016)
//line template/oauth2_login.qtpl:51
	return qs422016
//line template/oauth2_login.qtpl:51
}

//line template/oauth2_login.qtpl:53
func streamrenderScopeItem(qw422016 *qt422016.Writer, res oauth2.Resource, perm oauth2.Permission) {
//line template/oauth2_login.qtpl:53
	qw422016.N().S(` <div class="scope-item"> `)
//line template/oauth2_login.qtpl:55
	switch res {
//line template/oauth2_login.qtpl:56
	case oauth2.Build:
//line template/oauth2_login.qtpl:56
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:57
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M22.688 18.984c0.422 0.281 0.422 0.984-0.094 1.406l-2.297 2.297c-0.422 0.422-0.984 0.422-1.406 0l-9.094-9.094c-2.297 0.891-4.969 0.422-6.891-1.5-2.016-2.016-2.531-5.016-1.313-7.406l4.406 4.313 3-3-4.313-4.313c2.391-1.078 5.391-0.703 7.406 1.313 1.922 1.922 2.391 4.594 1.5 6.891z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:57
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:58
	case oauth2.Invite:
//line template/oauth2_login.qtpl:58
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:59
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>mail</title>
<path d="M12 11.016l8.016-5.016h-16.031zM20.016 18v-9.984l-8.016 4.969-8.016-4.969v9.984h16.031zM20.016 3.984c1.078 0 1.969 0.938 1.969 2.016v12c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h16.031z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:59
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:60
	case oauth2.Image:
//line template/oauth2_login.qtpl:60
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:61
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M8.484 13.5l-3.469 4.5h13.969l-4.5-6-3.469 4.5zM21 18.984c0 1.078-0.938 2.016-2.016 2.016h-13.969c-1.078 0-2.016-0.938-2.016-2.016v-13.969c0-1.078 0.938-2.016 2.016-2.016h13.969c1.078 0 2.016 0.938 2.016 2.016v13.969z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:61
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:62
	case oauth2.Namespace:
//line template/oauth2_login.qtpl:62
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:63
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9.984 3.984l2.016 2.016h8.016c1.078 0 1.969 0.938 1.969 2.016v9.984c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h6z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:63
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:64
	case oauth2.Object:
//line template/oauth2_login.qtpl:64
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:65
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM9 15.984v-6h-3.984l6.984-6.984 6.984 6.984h-3.984v6h-6z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:65
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:66
	case oauth2.Variable:
//line template/oauth2_login.qtpl:66
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:67
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:67
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:68
	case oauth2.Key:
//line template/oauth2_login.qtpl:68
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:69
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M6.984 14.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-1.969 0.938-1.969 2.016 0.891 2.016 1.969 2.016zM12.656 9.984h10.359v4.031h-2.016v3.984h-3.984v-3.984h-4.359c-0.797 2.344-3.047 3.984-5.672 3.984-3.328 0-6-2.672-6-6s2.672-6 6-6c2.625 0 4.875 1.641 5.672 3.984z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:69
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:70
	case oauth2.Cron:
//line template/oauth2_login.qtpl:70
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:71
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M21.984 12c0 5.156-3.938 9.422-8.953 9.938v-2.016c3.938-0.516 6.984-3.891 6.984-7.922s-3.047-7.406-6.984-7.922v-2.016c5.016 0.516 8.953 4.781 8.953 9.938zM5.672 19.734l1.406-1.406c1.125 0.844 2.484 1.406 3.938 1.594v2.016c-2.016-0.188-3.844-0.984-5.344-2.203zM4.078 12.984c0.188 1.453 0.75 2.813 1.594 3.891l-1.406 1.453c-1.219-1.5-2.016-3.328-2.203-5.344h2.016zM5.672 7.078c-0.844 1.125-1.406 2.484-1.594 3.938h-2.016c0.188-2.016 0.984-3.844 2.203-5.344zM11.016 4.078c-1.453 0.188-2.813 0.75-3.938 1.594l-1.406-1.406c1.5-1.219 3.328-2.016 5.344-2.203v2.016zM13.031 9.797l2.953 2.203c-2.007 1.493-4.007 2.993-6 4.5z"></path>
</svg>
`)
//line template/oauth2_login.qtpl:71
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:72
	}
//line template/oauth2_login.qtpl:72
	qw422016.N().S(` <span> <strong>`)
//line template/oauth2_login.qtpl:74
	qw422016.E().S(strings.Title(res.String()))
//line template/oauth2_login.qtpl:74
	qw422016.N().S(`</string> `)
//line template/oauth2_login.qtpl:75
	qw422016.E().S(buildPermissionList(perm))
//line template/oauth2_login.qtpl:75
	qw422016.N().S(` </span> </div> `)
//line template/oauth2_login.qtpl:78
}

//line template/oauth2_login.qtpl:78
func writerenderScopeItem(qq422016 qtio422016.Writer, res oauth2.Resource, perm oauth2.Permission) {
//line template/oauth2_login.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/oauth2_login.qtpl:78
	streamrenderScopeItem(qw422016, res, perm)
//line template/oauth2_login.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line template/oauth2_login.qtpl:78
}

//line template/oauth2_login.qtpl:78
func renderScopeItem(res oauth2.Resource, perm oauth2.Permission) string {
//line template/oauth2_login.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
//line template/oauth2_login.qtpl:78
	writerenderScopeItem(qb422016, res, perm)
//line template/oauth2_login.qtpl:78
	qs422016 := string(qb422016.B)
//line template/oauth2_login.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
//line template/oauth2_login.qtpl:78
	return qs422016
//line template/oauth2_login.qtpl:78
}

//line template/oauth2_login.qtpl:80
func (p *Oauth2Login) StreamBody(qw422016 *qt422016.Writer) {
//line template/oauth2_login.qtpl:80
	qw422016.N().S(` <div class="auth-page"> <div class="auth-form"> <div class="auth-header"> `)
//line template/oauth2_login.qtpl:84
	StreamLogo(qw422016)
//line template/oauth2_login.qtpl:84
	qw422016.N().S(` <h1>Authorize `)
//line template/oauth2_login.qtpl:85
	qw422016.E().S(p.Name)
//line template/oauth2_login.qtpl:85
	qw422016.N().S(`</h1> </div> <form action="/login/oauth/authorize" method="POST"> `)
//line template/oauth2_login.qtpl:88
	qw422016.N().V(p.CSRF)
//line template/oauth2_login.qtpl:88
	qw422016.N().S(` <input type="hidden" name="client_id" value="`)
//line template/oauth2_login.qtpl:89
	qw422016.E().S(p.ClientID)
//line template/oauth2_login.qtpl:89
	qw422016.N().S(`"/> <input type="hidden" name="redirect_uri" value="`)
//line template/oauth2_login.qtpl:90
	qw422016.E().S(p.RedirectURI)
//line template/oauth2_login.qtpl:90
	qw422016.N().S(`"/> <input type="hidden" name="scope" value="`)
//line template/oauth2_login.qtpl:91
	qw422016.E().S(p.Scope.String())
//line template/oauth2_login.qtpl:91
	qw422016.N().S(`"/> <input type="hidden" name="state" value="`)
//line template/oauth2_login.qtpl:92
	qw422016.E().S(p.State)
//line template/oauth2_login.qtpl:92
	qw422016.N().S(`"/> <div class="scope-list"> <h3>Requested scopes</h3> <div> <strong>`)
//line template/oauth2_login.qtpl:96
	qw422016.E().S(p.Name)
//line template/oauth2_login.qtpl:96
	qw422016.N().S(`</strong> from <strong>`)
//line template/oauth2_login.qtpl:96
	qw422016.E().S(p.Author.Email)
//line template/oauth2_login.qtpl:96
	qw422016.N().S(`</strong> would like taccess to the following, </div> `)
//line template/oauth2_login.qtpl:98
	for _, sc := range p.Scope {
//line template/oauth2_login.qtpl:98
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:99
		streamrenderScopeItem(qw422016, sc.Resource, sc.Permission)
//line template/oauth2_login.qtpl:99
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:100
	}
//line template/oauth2_login.qtpl:100
	qw422016.N().S(` </div> `)
//line template/oauth2_login.qtpl:102
	if p.User.ID == 0 {
//line template/oauth2_login.qtpl:102
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:103
		p.StreamField(qw422016, form.Field{
			ID:   "handle",
			Name: "Email / Username",
			Type: form.Text,
		})
//line template/oauth2_login.qtpl:107
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:108
		p.StreamField(qw422016, form.Field{
			ID:   "password",
			Name: "Password",
			Type: form.Password,
		})
//line template/oauth2_login.qtpl:112
		qw422016.N().S(` `)
//line template/oauth2_login.qtpl:113
	}
//line template/oauth2_login.qtpl:113
	qw422016.N().S(` <div class="input-field"> <button type="submit" class="btn btn-primary">Authorize</button> </div> </form> </div> </div> `)
//line template/oauth2_login.qtpl:120
}

//line template/oauth2_login.qtpl:120
func (p *Oauth2Login) WriteBody(qq422016 qtio422016.Writer) {
//line template/oauth2_login.qtpl:120
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/oauth2_login.qtpl:120
	p.StreamBody(qw422016)
//line template/oauth2_login.qtpl:120
	qt422016.ReleaseWriter(qw422016)
//line template/oauth2_login.qtpl:120
}

//line template/oauth2_login.qtpl:120
func (p *Oauth2Login) Body() string {
//line template/oauth2_login.qtpl:120
	qb422016 := qt422016.AcquireByteBuffer()
//line template/oauth2_login.qtpl:120
	p.WriteBody(qb422016)
//line template/oauth2_login.qtpl:120
	qs422016 := string(qb422016.B)
//line template/oauth2_login.qtpl:120
	qt422016.ReleaseByteBuffer(qb422016)
//line template/oauth2_login.qtpl:120
	return qs422016
//line template/oauth2_login.qtpl:120
}

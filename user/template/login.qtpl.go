// Code generated by qtc from "login.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line user/template/login.qtpl:2
package template

//line user/template/login.qtpl:2
import (
	"github.com/andrewpillar/djinn/provider"
	"github.com/andrewpillar/djinn/template"
)

//line user/template/login.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line user/template/login.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line user/template/login.qtpl:9
type Login struct {
	template.BasePage
	template.Form

	Alert     template.Alert
	Providers []*provider.Provider
}

var providerNames = map[string]string{
	"github": "GitHub",
	"gitlab": "GitLab",
}

//line user/template/login.qtpl:24
func (p *Login) StreamTitle(qw422016 *qt422016.Writer) {
//line user/template/login.qtpl:24
	qw422016.N().S(` Login - Thrall `)
//line user/template/login.qtpl:26
}

//line user/template/login.qtpl:26
func (p *Login) WriteTitle(qq422016 qtio422016.Writer) {
//line user/template/login.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/login.qtpl:26
	p.StreamTitle(qw422016)
//line user/template/login.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line user/template/login.qtpl:26
}

//line user/template/login.qtpl:26
func (p *Login) Title() string {
//line user/template/login.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/login.qtpl:26
	p.WriteTitle(qb422016)
//line user/template/login.qtpl:26
	qs422016 := string(qb422016.B)
//line user/template/login.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/login.qtpl:26
	return qs422016
//line user/template/login.qtpl:26
}

//line user/template/login.qtpl:28
func (p *Login) StreamFooter(qw422016 *qt422016.Writer) {
//line user/template/login.qtpl:28
	qw422016.N().S(` <style type="text/css">`)
//line user/template/login.qtpl:29
	qw422016.N().S(`*{margin:0;padding:0}body{font-family:sans-serif;font-size:14px;background:#383e51;color:#fff}a{text-decoration:none;cursor:pointer}a:hover{text-decoration:underline}button{cursor:pointer}h1{font-weight:400}.btn{border:none;border-radius:3px;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";padding:10px;padding-left:15px;padding-right:15px;color:#fff}.btn svg{fill:#fff}.btn:hover{text-decoration:none}.btn-primary{background:#61a0ea}.btn-primary:hover{background:#5090d9}.btn-danger{background:#de4141}.btn-danger:hover{background:#cd3030}.logo .left{display:inline-block;width:0;height:0;border-style:solid;border-width:25px 0 0 15px;border-color:transparent transparent transparent #fff}.logo .right{display:inline-block;width:0;height:0;border-style:solid;border-width:0 0 25px 15px;border-color:transparent transparent #fff transparent}.auth-page a{color:#66c9ff}.auth-page .auth-form{margin:0 auto;margin-top:150px;max-width:400px;padding:20px}.auth-page .auth-form .auth-header{margin-bottom:20px;text-align:center}.auth-page .auth-form .auth-header .brand{margin:0 auto}.auth-page .auth-form .auth-header .brand .left{border-width:45px 0 0 25px}.auth-page .auth-form .auth-header .brand .right{border-width:0 0 45px 25px}.auth-page .auth-form .form-error{margin-top:3px;display:block;color:#e74848;min-height:17px}.auth-page .auth-form .input-field{margin-top:10px;width:100%}.auth-page .auth-form .input-field label{margin-bottom:3px;display:block}.auth-page .auth-form .input-field .text{box-sizing:border-box;width:100%;font-family:sans-serif;font-size:14px;padding:7px;outline:0;border:solid 1px rgba(255,255,255,.3);border-radius:3px;background:rgba(0,0,0,.3);color:#fff}.auth-page .auth-form .input-field .text:focus{border:solid 1px rgba(255,255,255,.5)}.auth-page .auth-form .input-field .btn{color:#fff;width:100%;display:block;text-align:center;box-sizing:border-box}.provider-btn svg{margin-right:5px;fill:#fff;vertical-align:middle}.provider-btn span{display:inline-block;vertical-align:middle}.provider-btn:hover{text-decoration:none}.provider-github{background:#24292e}.provider-github:hover{background:#353a3f}.provider-gitlab{background:#fa7035}.provider-gitlab:hover{background:#e65328}.alert{margin-top:15px;overflow:auto;padding:15px;border-radius:3px}.alert .alert-message{float:left;color:rgba(0,0,0,.6)}.alert a{float:right;display:inline-block}.alert a svg{width:15px;height:15px;fill:rgba(0,0,0,.4)}.alert a:hover svg{fill:rgba(0,0,0,.5)}.alert-success{background:#caf5ca;border-bottom:solid 1px #a0dfa0}.alert-warn{background:#fff3cd;border-bottom:solid 1px #d9c995}.alert-danger{background:#ffd4d4;border-bottom:solid 1px #e19e9e}`)
//line user/template/login.qtpl:29
	qw422016.N().S(`</style> `)
//line user/template/login.qtpl:30
}

//line user/template/login.qtpl:30
func (p *Login) WriteFooter(qq422016 qtio422016.Writer) {
//line user/template/login.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/login.qtpl:30
	p.StreamFooter(qw422016)
//line user/template/login.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line user/template/login.qtpl:30
}

//line user/template/login.qtpl:30
func (p *Login) Footer() string {
//line user/template/login.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/login.qtpl:30
	p.WriteFooter(qb422016)
//line user/template/login.qtpl:30
	qs422016 := string(qb422016.B)
//line user/template/login.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/login.qtpl:30
	return qs422016
//line user/template/login.qtpl:30
}

//line user/template/login.qtpl:32
func (p *Login) StreamBody(qw422016 *qt422016.Writer) {
//line user/template/login.qtpl:32
	qw422016.N().S(` <div class="auth-page"> <div class="auth-form"> <div class="auth-header"> <div class="brand"> <div class="left"></div> <div class="right"></div> </div> <h1>Login to Thrall</h1> `)
//line user/template/login.qtpl:41
	template.StreamRenderAlert(qw422016, p.Alert, "")
//line user/template/login.qtpl:41
	qw422016.N().S(` </div> <form method="POST" action="/login"> `)
//line user/template/login.qtpl:44
	qw422016.N().S(string(p.CSRF))
//line user/template/login.qtpl:44
	qw422016.N().S(` <div class="input-field"> <label>Email / Username</label> <input class="text" type="text" name="handle" value="`)
//line user/template/login.qtpl:47
	qw422016.E().S(p.Fields["handle"])
//line user/template/login.qtpl:47
	qw422016.N().S(`" autocomplete="off"/> `)
//line user/template/login.qtpl:48
	p.StreamError(qw422016, "handle")
//line user/template/login.qtpl:48
	qw422016.N().S(` </div> <div class="input-field"> <label>Password</label> <input class="text" type="password" name="password" autocomplete="off"/> `)
//line user/template/login.qtpl:53
	p.StreamError(qw422016, "password")
//line user/template/login.qtpl:53
	qw422016.N().S(` </div> <div class="input-field"> <button type="submit" class="btn btn-primary">Login</button> </div> `)
//line user/template/login.qtpl:58
	for _, prv := range p.Providers {
//line user/template/login.qtpl:58
		qw422016.N().S(` <div class="input-field"> <a href="`)
//line user/template/login.qtpl:60
		qw422016.E().S(prv.AuthURL)
//line user/template/login.qtpl:60
		qw422016.N().S(`" class="btn provider-btn provider-`)
//line user/template/login.qtpl:60
		qw422016.E().S(prv.Name)
//line user/template/login.qtpl:60
		qw422016.N().S(`"> `)
//line user/template/login.qtpl:61
		switch prv.Name {
//line user/template/login.qtpl:62
		case "github":
//line user/template/login.qtpl:62
			qw422016.N().S(` `)
//line user/template/login.qtpl:63
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385 0.6 0.113 0.82-0.258 0.82-0.577 0-0.285-0.010-1.040-0.015-2.040-3.338 0.724-4.042-1.61-4.042-1.61-0.546-1.385-1.335-1.755-1.335-1.755-1.087-0.744 0.084-0.729 0.084-0.729 1.205 0.084 1.838 1.236 1.838 1.236 1.070 1.835 2.809 1.305 3.495 0.998 0.108-0.776 0.417-1.305 0.76-1.605-2.665-0.3-5.466-1.332-5.466-5.93 0-1.31 0.465-2.38 1.235-3.22-0.135-0.303-0.54-1.523 0.105-3.176 0 0 1.005-0.322 3.3 1.23 0.96-0.267 1.98-0.399 3-0.405 1.020 0.006 2.040 0.138 3 0.405 2.28-1.552 3.285-1.23 3.285-1.23 0.645 1.653 0.24 2.873 0.12 3.176 0.765 0.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92 0.42 0.36 0.81 1.096 0.81 2.22 0 1.606-0.015 2.896-0.015 3.286 0 0.315 0.21 0.69 0.825 0.57 4.801-1.574 8.236-6.074 8.236-11.369 0-6.627-5.373-12-12-12z"></path>
</svg>
`)
//line user/template/login.qtpl:63
			qw422016.N().S(` `)
//line user/template/login.qtpl:64
		case "gitlab":
//line user/template/login.qtpl:64
			qw422016.N().S(` `)
//line user/template/login.qtpl:65
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M23.955 13.587l-1.342-4.135-2.664-8.189c-0.135-0.423-0.73-0.423-0.867 0l-2.664 8.187h-8.836l-2.663-8.187c-0.136-0.423-0.734-0.423-0.869-0.003l-2.664 8.189-1.342 4.138c-0.121 0.375 0.014 0.789 0.331 1.023l11.625 8.444 11.625-8.443c0.318-0.235 0.453-0.647 0.33-1.024z"></path>
</svg>
`)
//line user/template/login.qtpl:65
			qw422016.N().S(` `)
//line user/template/login.qtpl:66
		}
//line user/template/login.qtpl:66
		qw422016.N().S(` <span>Login with `)
//line user/template/login.qtpl:67
		qw422016.E().S(providerNames[prv.Name])
//line user/template/login.qtpl:67
		qw422016.N().S(`</span> </a> </div> `)
//line user/template/login.qtpl:70
	}
//line user/template/login.qtpl:70
	qw422016.N().S(` <div class="input-field">Don't have an account? <a href="/register">Register</a></div> <div class="input-field"><a href="/password_reset">Forgot your password?</a></div> </form> </div> </div> `)
//line user/template/login.qtpl:76
}

//line user/template/login.qtpl:76
func (p *Login) WriteBody(qq422016 qtio422016.Writer) {
//line user/template/login.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line user/template/login.qtpl:76
	p.StreamBody(qw422016)
//line user/template/login.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line user/template/login.qtpl:76
}

//line user/template/login.qtpl:76
func (p *Login) Body() string {
//line user/template/login.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line user/template/login.qtpl:76
	p.WriteBody(qb422016)
//line user/template/login.qtpl:76
	qs422016 := string(qb422016.B)
//line user/template/login.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line user/template/login.qtpl:76
	return qs422016
//line user/template/login.qtpl:76
}

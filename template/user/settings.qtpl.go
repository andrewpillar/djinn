// This file is automatically generated by qtc from "settings.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/user/settings.qtpl:2
package user

//line template/user/settings.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/user/settings.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/user/settings.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/user/settings.qtpl:9
type SettingsPage struct {
	template.BasePage

	User *model.User
}

type ShowInvites struct {
	SettingsPage

	CSRF    string
	Invites []*model.Invite
}

//line template/user/settings.qtpl:24
func (p *SettingsPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:24
	qw422016.N().S(` Settings `)
//line template/user/settings.qtpl:26
}

//line template/user/settings.qtpl:26
func (p *SettingsPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:26
	p.StreamTitle(qw422016)
	//line template/user/settings.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:26
}

//line template/user/settings.qtpl:26
func (p *SettingsPage) Title() string {
	//line template/user/settings.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:26
	p.WriteTitle(qb422016)
	//line template/user/settings.qtpl:26
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:26
	return qs422016
//line template/user/settings.qtpl:26
}

//line template/user/settings.qtpl:28
func (p *ShowInvites) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:28
	qw422016.N().S(` Settings - Namespace Invites `)
//line template/user/settings.qtpl:30
}

//line template/user/settings.qtpl:30
func (p *ShowInvites) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:30
	p.StreamTitle(qw422016)
	//line template/user/settings.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:30
}

//line template/user/settings.qtpl:30
func (p *ShowInvites) Title() string {
	//line template/user/settings.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:30
	p.WriteTitle(qb422016)
	//line template/user/settings.qtpl:30
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:30
	return qs422016
//line template/user/settings.qtpl:30
}

//line template/user/settings.qtpl:32
func (p *SettingsPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:32
	qw422016.N().S(` `)
//line template/user/settings.qtpl:33
}

//line template/user/settings.qtpl:33
func (p *SettingsPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:33
	p.StreamBody(qw422016)
	//line template/user/settings.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:33
}

//line template/user/settings.qtpl:33
func (p *SettingsPage) Body() string {
	//line template/user/settings.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:33
	p.WriteBody(qb422016)
	//line template/user/settings.qtpl:33
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:33
	return qs422016
//line template/user/settings.qtpl:33
}

//line template/user/settings.qtpl:35
func (p *ShowInvites) StreamBody(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:35
	qw422016.N().S(` <div class="panel"> `)
	//line template/user/settings.qtpl:37
	if len(p.Invites) == 0 {
		//line template/user/settings.qtpl:37
		qw422016.N().S(` <div class="panel-message muted">No new Namespace invites.</div> `)
		//line template/user/settings.qtpl:39
	} else {
		//line template/user/settings.qtpl:39
		qw422016.N().S(` <table class="table"> <thead> <tr> <th>NAMESPACE</th> <th>INVITED BY</th> <th></th> </tr> </thead> <tbody> `)
		//line template/user/settings.qtpl:49
		for _, i := range p.Invites {
			//line template/user/settings.qtpl:49
			qw422016.N().S(` <tr> <td>`)
			//line template/user/settings.qtpl:51
			qw422016.E().S(i.Namespace.Path)
			//line template/user/settings.qtpl:51
			qw422016.N().S(`</td> <td>`)
			//line template/user/settings.qtpl:52
			qw422016.E().S(i.Inviter.Username)
			//line template/user/settings.qtpl:52
			qw422016.N().S(`</td> <td class="align-right"> <form method="POST" action="`)
			//line template/user/settings.qtpl:54
			qw422016.E().S(i.UIEndpoint())
			//line template/user/settings.qtpl:54
			qw422016.N().S(`" class="inline-block"> `)
			//line template/user/settings.qtpl:55
			qw422016.N().S(p.CSRF)
			//line template/user/settings.qtpl:55
			qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> <button type="submit" class="btn btn-primary">Accept</button> </form> <form method="POST" action="`)
			//line template/user/settings.qtpl:59
			qw422016.E().S(i.UIEndpoint())
			//line template/user/settings.qtpl:59
			qw422016.N().S(`" class="inline-block"> `)
			//line template/user/settings.qtpl:60
			qw422016.N().S(p.CSRF)
			//line template/user/settings.qtpl:60
			qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <button type="submit" class="btn btn-danger">Decline</button> </form> </td> </tr> `)
			//line template/user/settings.qtpl:66
		}
		//line template/user/settings.qtpl:66
		qw422016.N().S(` </tbody> </table> `)
		//line template/user/settings.qtpl:69
	}
	//line template/user/settings.qtpl:69
	qw422016.N().S(` </div> `)
//line template/user/settings.qtpl:71
}

//line template/user/settings.qtpl:71
func (p *ShowInvites) WriteBody(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:71
	p.StreamBody(qw422016)
	//line template/user/settings.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:71
}

//line template/user/settings.qtpl:71
func (p *ShowInvites) Body() string {
	//line template/user/settings.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:71
	p.WriteBody(qb422016)
	//line template/user/settings.qtpl:71
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:71
	return qs422016
//line template/user/settings.qtpl:71
}

//line template/user/settings.qtpl:73
func (p *SettingsPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:73
	qw422016.N().S(` Settings `)
//line template/user/settings.qtpl:75
}

//line template/user/settings.qtpl:75
func (p *SettingsPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:75
	p.StreamHeader(qw422016)
	//line template/user/settings.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:75
}

//line template/user/settings.qtpl:75
func (p *SettingsPage) Header() string {
	//line template/user/settings.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:75
	p.WriteHeader(qb422016)
	//line template/user/settings.qtpl:75
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:75
	return qs422016
//line template/user/settings.qtpl:75
}

//line template/user/settings.qtpl:77
func (p *ShowInvites) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:77
	qw422016.N().S(` Settings - Namespace Invites `)
//line template/user/settings.qtpl:79
}

//line template/user/settings.qtpl:79
func (p *ShowInvites) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:79
	p.StreamHeader(qw422016)
	//line template/user/settings.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:79
}

//line template/user/settings.qtpl:79
func (p *ShowInvites) Header() string {
	//line template/user/settings.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:79
	p.WriteHeader(qb422016)
	//line template/user/settings.qtpl:79
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:79
	return qs422016
//line template/user/settings.qtpl:79
}

//line template/user/settings.qtpl:81
func (p *SettingsPage) StreamActions(qw422016 *qt422016.Writer) {
//line template/user/settings.qtpl:81
}

//line template/user/settings.qtpl:81
func (p *SettingsPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:81
	p.StreamActions(qw422016)
	//line template/user/settings.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:81
}

//line template/user/settings.qtpl:81
func (p *SettingsPage) Actions() string {
	//line template/user/settings.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:81
	p.WriteActions(qb422016)
	//line template/user/settings.qtpl:81
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:81
	return qs422016
//line template/user/settings.qtpl:81
}

//line template/user/settings.qtpl:83
func (p *SettingsPage) StreamNavigation(qw422016 *qt422016.Writer) {
	//line template/user/settings.qtpl:83
	qw422016.N().S(` <li> `)
	//line template/user/settings.qtpl:85
	if p.URI == "/settings" {
		//line template/user/settings.qtpl:85
		qw422016.N().S(` <a href="/settings" class="active">`)
		//line template/user/settings.qtpl:86
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>user</title>
<path d="M12 14.016c2.672 0 8.016 1.313 8.016 3.984v2.016h-16.031v-2.016c0-2.672 5.344-3.984 8.016-3.984zM12 12c-2.203 0-3.984-1.781-3.984-3.984s1.781-4.031 3.984-4.031 3.984 1.828 3.984 4.031-1.781 3.984-3.984 3.984z"></path>
</svg>
`)
		//line template/user/settings.qtpl:86
		qw422016.N().S(`<span>Account</span></a> `)
		//line template/user/settings.qtpl:87
	} else {
		//line template/user/settings.qtpl:87
		qw422016.N().S(` <a href="/settings">`)
		//line template/user/settings.qtpl:88
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<title>user</title>
<path d="M12 14.016c2.672 0 8.016 1.313 8.016 3.984v2.016h-16.031v-2.016c0-2.672 5.344-3.984 8.016-3.984zM12 12c-2.203 0-3.984-1.781-3.984-3.984s1.781-4.031 3.984-4.031 3.984 1.828 3.984 4.031-1.781 3.984-3.984 3.984z"></path>
</svg>
`)
		//line template/user/settings.qtpl:88
		qw422016.N().S(`<span>Account</span></a> `)
		//line template/user/settings.qtpl:89
	}
	//line template/user/settings.qtpl:89
	qw422016.N().S(` </li> <li> `)
	//line template/user/settings.qtpl:92
	if p.URI == "/settings/invites" {
		//line template/user/settings.qtpl:92
		qw422016.N().S(` <a href="/settings/invites" class="active">`)
		//line template/user/settings.qtpl:93
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9.984 3.984l2.016 2.016h8.016c1.078 0 1.969 0.938 1.969 2.016v9.984c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h6z"></path>
</svg>
`)
		//line template/user/settings.qtpl:93
		qw422016.N().S(`<span>Invites</span></a> `)
		//line template/user/settings.qtpl:94
	} else {
		//line template/user/settings.qtpl:94
		qw422016.N().S(` <a href="/settings/invites">`)
		//line template/user/settings.qtpl:95
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9.984 3.984l2.016 2.016h8.016c1.078 0 1.969 0.938 1.969 2.016v9.984c0 1.078-0.891 2.016-1.969 2.016h-16.031c-1.078 0-1.969-0.938-1.969-2.016v-12c0-1.078 0.891-2.016 1.969-2.016h6z"></path>
</svg>
`)
		//line template/user/settings.qtpl:95
		qw422016.N().S(`<span>Invites</span></a> `)
		//line template/user/settings.qtpl:96
	}
	//line template/user/settings.qtpl:96
	qw422016.N().S(` </li> `)
//line template/user/settings.qtpl:98
}

//line template/user/settings.qtpl:98
func (p *SettingsPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/user/settings.qtpl:98
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/user/settings.qtpl:98
	p.StreamNavigation(qw422016)
	//line template/user/settings.qtpl:98
	qt422016.ReleaseWriter(qw422016)
//line template/user/settings.qtpl:98
}

//line template/user/settings.qtpl:98
func (p *SettingsPage) Navigation() string {
	//line template/user/settings.qtpl:98
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/user/settings.qtpl:98
	p.WriteNavigation(qb422016)
	//line template/user/settings.qtpl:98
	qs422016 := string(qb422016.B)
	//line template/user/settings.qtpl:98
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/user/settings.qtpl:98
	return qs422016
//line template/user/settings.qtpl:98
}
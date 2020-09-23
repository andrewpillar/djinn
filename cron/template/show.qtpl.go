// Code generated by qtc from "show.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line cron/template/show.qtpl:2
package template

//line cron/template/show.qtpl:2
import (
	htmltemplate "html/template"
	"strings"

	buildtemplate "github.com/andrewpillar/djinn/build/template"
	"github.com/andrewpillar/djinn/cron"
	"github.com/andrewpillar/djinn/template"
)

//line cron/template/show.qtpl:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line cron/template/show.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line cron/template/show.qtpl:13
type Show struct {
	template.BasePage

	CSRF   htmltemplate.HTML
	Cron   *cron.Cron
	Builds *buildtemplate.Index
}

//line cron/template/show.qtpl:23
func (p *Show) StreamTitle(qw422016 *qt422016.Writer) {
//line cron/template/show.qtpl:23
	qw422016.N().S(` Cron Jobs - `)
//line cron/template/show.qtpl:24
	qw422016.E().S(p.Cron.Name)
//line cron/template/show.qtpl:24
	qw422016.N().S(` - Djinn `)
//line cron/template/show.qtpl:25
}

//line cron/template/show.qtpl:25
func (p *Show) WriteTitle(qq422016 qtio422016.Writer) {
//line cron/template/show.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cron/template/show.qtpl:25
	p.StreamTitle(qw422016)
//line cron/template/show.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line cron/template/show.qtpl:25
}

//line cron/template/show.qtpl:25
func (p *Show) Title() string {
//line cron/template/show.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line cron/template/show.qtpl:25
	p.WriteTitle(qb422016)
//line cron/template/show.qtpl:25
	qs422016 := string(qb422016.B)
//line cron/template/show.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line cron/template/show.qtpl:25
	return qs422016
//line cron/template/show.qtpl:25
}

//line cron/template/show.qtpl:27
func (p *Show) StreamBody(qw422016 *qt422016.Writer) {
//line cron/template/show.qtpl:27
	qw422016.N().S(` <div class="panel"> <table class="table"> <tr> <td>Name</td> <td class="align-right">`)
//line cron/template/show.qtpl:32
	qw422016.E().S(p.Cron.Name)
//line cron/template/show.qtpl:32
	qw422016.N().S(`</td> </tr> <tr> <td>Schedule</td> <td class="align-right">`)
//line cron/template/show.qtpl:36
	qw422016.E().S(strings.ToUpper(p.Cron.Schedule.String()))
//line cron/template/show.qtpl:36
	qw422016.N().S(`</td> </tr> <tr> <td>Previous Run</td> `)
//line cron/template/show.qtpl:40
	if p.Cron.PrevRun.Valid {
//line cron/template/show.qtpl:40
		qw422016.N().S(` <td class="align-right">`)
//line cron/template/show.qtpl:41
		qw422016.E().S(p.Cron.PrevRun.Time.Format("Mon, 2 Jan 2006 15:04"))
//line cron/template/show.qtpl:41
		qw422016.N().S(`</td> `)
//line cron/template/show.qtpl:42
	} else {
//line cron/template/show.qtpl:42
		qw422016.N().S(` <td class="align-right muted">--</td> `)
//line cron/template/show.qtpl:44
	}
//line cron/template/show.qtpl:44
	qw422016.N().S(` </tr> <tr> <td>Next Run</td> <td class="align-right">`)
//line cron/template/show.qtpl:48
	qw422016.E().S(p.Cron.NextRun.Format("Mon, 2 Jan 2006 15:04"))
//line cron/template/show.qtpl:48
	qw422016.N().S(`</td> </tr> </table> </div> `)
//line cron/template/show.qtpl:52
	p.Builds.StreamBody(qw422016)
//line cron/template/show.qtpl:52
	qw422016.N().S(` `)
//line cron/template/show.qtpl:53
}

//line cron/template/show.qtpl:53
func (p *Show) WriteBody(qq422016 qtio422016.Writer) {
//line cron/template/show.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cron/template/show.qtpl:53
	p.StreamBody(qw422016)
//line cron/template/show.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line cron/template/show.qtpl:53
}

//line cron/template/show.qtpl:53
func (p *Show) Body() string {
//line cron/template/show.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
//line cron/template/show.qtpl:53
	p.WriteBody(qb422016)
//line cron/template/show.qtpl:53
	qs422016 := string(qb422016.B)
//line cron/template/show.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
//line cron/template/show.qtpl:53
	return qs422016
//line cron/template/show.qtpl:53
}

//line cron/template/show.qtpl:55
func (p *Show) StreamHeader(qw422016 *qt422016.Writer) {
//line cron/template/show.qtpl:55
	qw422016.N().S(` <a href="/cron" class="back">`)
//line cron/template/show.qtpl:56
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line cron/template/show.qtpl:56
	qw422016.N().S(`</a> `)
//line cron/template/show.qtpl:57
	if !p.Cron.Namespace.IsZero() {
//line cron/template/show.qtpl:57
		qw422016.N().S(` <a href="`)
//line cron/template/show.qtpl:58
		qw422016.E().S(p.Cron.Namespace.Endpoint())
//line cron/template/show.qtpl:58
		qw422016.N().S(`">`)
//line cron/template/show.qtpl:58
		qw422016.E().S(p.Cron.Namespace.Name)
//line cron/template/show.qtpl:58
		qw422016.N().S(`</a> / `)
//line cron/template/show.qtpl:59
	}
//line cron/template/show.qtpl:59
	qw422016.N().S(` `)
//line cron/template/show.qtpl:60
	qw422016.E().S(p.Cron.Name)
//line cron/template/show.qtpl:60
	qw422016.N().S(` `)
//line cron/template/show.qtpl:61
}

//line cron/template/show.qtpl:61
func (p *Show) WriteHeader(qq422016 qtio422016.Writer) {
//line cron/template/show.qtpl:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cron/template/show.qtpl:61
	p.StreamHeader(qw422016)
//line cron/template/show.qtpl:61
	qt422016.ReleaseWriter(qw422016)
//line cron/template/show.qtpl:61
}

//line cron/template/show.qtpl:61
func (p *Show) Header() string {
//line cron/template/show.qtpl:61
	qb422016 := qt422016.AcquireByteBuffer()
//line cron/template/show.qtpl:61
	p.WriteHeader(qb422016)
//line cron/template/show.qtpl:61
	qs422016 := string(qb422016.B)
//line cron/template/show.qtpl:61
	qt422016.ReleaseByteBuffer(qb422016)
//line cron/template/show.qtpl:61
	return qs422016
//line cron/template/show.qtpl:61
}

//line cron/template/show.qtpl:63
func (p *Show) StreamActions(qw422016 *qt422016.Writer) {
//line cron/template/show.qtpl:63
	qw422016.N().S(` <li><a href="`)
//line cron/template/show.qtpl:64
	qw422016.E().S(p.Cron.Endpoint("edit"))
//line cron/template/show.qtpl:64
	qw422016.N().S(`" class="btn btn-primary">Edit</a></li> `)
//line cron/template/show.qtpl:65
	if p.User.ID == p.Cron.UserID {
//line cron/template/show.qtpl:65
		qw422016.N().S(` <li> <form method="POST" action="`)
//line cron/template/show.qtpl:67
		qw422016.E().S(p.Cron.Endpoint())
//line cron/template/show.qtpl:67
		qw422016.N().S(`"> `)
//line cron/template/show.qtpl:68
		qw422016.N().V(p.CSRF)
//line cron/template/show.qtpl:68
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <button type="submit" class="btn btn-danger">Delete</button> </form> </li> `)
//line cron/template/show.qtpl:73
	}
//line cron/template/show.qtpl:73
	qw422016.N().S(` `)
//line cron/template/show.qtpl:74
}

//line cron/template/show.qtpl:74
func (p *Show) WriteActions(qq422016 qtio422016.Writer) {
//line cron/template/show.qtpl:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cron/template/show.qtpl:74
	p.StreamActions(qw422016)
//line cron/template/show.qtpl:74
	qt422016.ReleaseWriter(qw422016)
//line cron/template/show.qtpl:74
}

//line cron/template/show.qtpl:74
func (p *Show) Actions() string {
//line cron/template/show.qtpl:74
	qb422016 := qt422016.AcquireByteBuffer()
//line cron/template/show.qtpl:74
	p.WriteActions(qb422016)
//line cron/template/show.qtpl:74
	qs422016 := string(qb422016.B)
//line cron/template/show.qtpl:74
	qt422016.ReleaseByteBuffer(qb422016)
//line cron/template/show.qtpl:74
	return qs422016
//line cron/template/show.qtpl:74
}

//line cron/template/show.qtpl:76
func (p *Show) StreamNavigation(qw422016 *qt422016.Writer) {
//line cron/template/show.qtpl:76
}

//line cron/template/show.qtpl:76
func (p *Show) WriteNavigation(qq422016 qtio422016.Writer) {
//line cron/template/show.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cron/template/show.qtpl:76
	p.StreamNavigation(qw422016)
//line cron/template/show.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line cron/template/show.qtpl:76
}

//line cron/template/show.qtpl:76
func (p *Show) Navigation() string {
//line cron/template/show.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line cron/template/show.qtpl:76
	p.WriteNavigation(qb422016)
//line cron/template/show.qtpl:76
	qs422016 := string(qb422016.B)
//line cron/template/show.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line cron/template/show.qtpl:76
	return qs422016
//line cron/template/show.qtpl:76
}

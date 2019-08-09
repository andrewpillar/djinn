// This file is automatically generated by qtc from "show.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/show.qtpl:2
package build

//line template/build/show.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/job"
)

//line template/build/show.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/show.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/show.qtpl:10
type ShowPage struct {
	template.BasePage

	Build *model.Build
}

type ShowManifest struct {
	ShowPage
}

type ShowOutput struct {
	ShowPage
}

type ShowObjects struct {
	ShowPage

	Objects []*model.BuildObject
}

type ShowVariables struct {
	ShowPage

	Variables []*model.BuildVariable
}

//line template/build/show.qtpl:38
func (p *ShowPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:38
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:39
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:39
	qw422016.N().S(` - Thrall `)
//line template/build/show.qtpl:40
}

//line template/build/show.qtpl:40
func (p *ShowPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:40
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:40
}

//line template/build/show.qtpl:40
func (p *ShowPage) Title() string {
	//line template/build/show.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:40
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:40
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:40
	return qs422016
//line template/build/show.qtpl:40
}

//line template/build/show.qtpl:42
func (p *ShowManifest) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:42
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:43
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:43
	qw422016.N().S(` - Manifest `)
//line template/build/show.qtpl:44
}

//line template/build/show.qtpl:44
func (p *ShowManifest) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:44
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:44
}

//line template/build/show.qtpl:44
func (p *ShowManifest) Title() string {
	//line template/build/show.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:44
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:44
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:44
	return qs422016
//line template/build/show.qtpl:44
}

//line template/build/show.qtpl:46
func (p *ShowOutput) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:46
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:47
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:47
	qw422016.N().S(` - Output `)
//line template/build/show.qtpl:48
}

//line template/build/show.qtpl:48
func (p *ShowOutput) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:48
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:48
}

//line template/build/show.qtpl:48
func (p *ShowOutput) Title() string {
	//line template/build/show.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:48
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:48
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:48
	return qs422016
//line template/build/show.qtpl:48
}

//line template/build/show.qtpl:50
func (p *ShowObjects) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:50
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:51
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:51
	qw422016.N().S(` - Objects `)
//line template/build/show.qtpl:52
}

//line template/build/show.qtpl:52
func (p *ShowObjects) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:52
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:52
}

//line template/build/show.qtpl:52
func (p *ShowObjects) Title() string {
	//line template/build/show.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:52
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:52
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:52
	return qs422016
//line template/build/show.qtpl:52
}

//line template/build/show.qtpl:54
func (p *ShowVariables) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:54
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:55
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:55
	qw422016.N().S(` - Variables `)
//line template/build/show.qtpl:56
}

//line template/build/show.qtpl:56
func (p *ShowVariables) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:56
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:56
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:56
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:56
}

//line template/build/show.qtpl:56
func (p *ShowVariables) Title() string {
	//line template/build/show.qtpl:56
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:56
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:56
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:56
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:56
	return qs422016
//line template/build/show.qtpl:56
}

//line template/build/show.qtpl:58
func (p *ShowPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:58
	qw422016.N().S(` <div class="mb-10 overflow"> <div class="col-75 pr-5 left"> <div class="panel"> <div class="panel-header"> `)
	//line template/build/show.qtpl:63
	if p.Build.Namespace.IsZero() {
		//line template/build/show.qtpl:63
		qw422016.N().S(` <h3>Submitted by `)
		//line template/build/show.qtpl:64
		qw422016.E().S(p.Build.User.Username)
		//line template/build/show.qtpl:64
		qw422016.N().S(` &lt;`)
		//line template/build/show.qtpl:64
		qw422016.E().S(p.Build.User.Email)
		//line template/build/show.qtpl:64
		qw422016.N().S(`&gt;</h3> `)
		//line template/build/show.qtpl:65
	} else {
		//line template/build/show.qtpl:65
		qw422016.N().S(` <h3>Submitted by `)
		//line template/build/show.qtpl:66
		qw422016.E().S(p.Build.User.Username)
		//line template/build/show.qtpl:66
		qw422016.N().S(` &lt;`)
		//line template/build/show.qtpl:66
		qw422016.E().S(p.Build.User.Email)
		//line template/build/show.qtpl:66
		qw422016.N().S(`&gt; to <a href="`)
		//line template/build/show.qtpl:66
		qw422016.E().S(p.Build.Namespace.UIEndpoint())
		//line template/build/show.qtpl:66
		qw422016.N().S(`">`)
		//line template/build/show.qtpl:66
		qw422016.E().S(p.Build.Namespace.Path)
		//line template/build/show.qtpl:66
		qw422016.N().S(`</a></h3> `)
		//line template/build/show.qtpl:67
	}
	//line template/build/show.qtpl:67
	qw422016.N().S(` </div> <div class="panel-body"> `)
	//line template/build/show.qtpl:70
	if p.Build.Trigger.Comment != "" {
		//line template/build/show.qtpl:70
		qw422016.N().S(` <pre class="code">`)
		//line template/build/show.qtpl:71
		qw422016.E().S(p.Build.Trigger.Comment)
		//line template/build/show.qtpl:71
		qw422016.N().S(`</pre> `)
		//line template/build/show.qtpl:72
	} else {
		//line template/build/show.qtpl:72
		qw422016.N().S(` <em class="muted">No build comment.</em> `)
		//line template/build/show.qtpl:74
	}
	//line template/build/show.qtpl:74
	qw422016.N().S(` </div> <div class="panel-footer"> <div class="mb-10"> `)
	//line template/build/show.qtpl:78
	if len(p.Build.Tags) > 0 {
		//line template/build/show.qtpl:78
		qw422016.N().S(` `)
		//line template/build/show.qtpl:79
		for _, t := range p.Build.Tags {
			//line template/build/show.qtpl:79
			qw422016.N().S(` <a href="/?tag=`)
			//line template/build/show.qtpl:80
			qw422016.E().S(t.Name)
			//line template/build/show.qtpl:80
			qw422016.N().S(`" class="pill pill-light">`)
			//line template/build/show.qtpl:80
			qw422016.E().S(t.Name)
			//line template/build/show.qtpl:80
			qw422016.N().S(`</a> `)
			//line template/build/show.qtpl:81
		}
		//line template/build/show.qtpl:81
		qw422016.N().S(` `)
		//line template/build/show.qtpl:82
	} else {
		//line template/build/show.qtpl:82
		qw422016.N().S(` <em class="muted">No build tags.</em> `)
		//line template/build/show.qtpl:84
	}
	//line template/build/show.qtpl:84
	qw422016.N().S(` </div> </div> </div> </div> <div class="col-25 pl-5 right"> <div class="panel"> <table class="table"> <tr> <td>Started at:</td> <td class="align-right"> `)
	//line template/build/show.qtpl:95
	if p.Build.StartedAt.Valid {
		//line template/build/show.qtpl:95
		qw422016.N().S(` `)
		//line template/build/show.qtpl:96
		qw422016.E().S(p.Build.StartedAt.Time.Format("2006-01-02T15:04:05"))
		//line template/build/show.qtpl:96
		qw422016.N().S(` `)
		//line template/build/show.qtpl:97
	} else {
		//line template/build/show.qtpl:97
		qw422016.N().S(` <span class="muted">--</span> `)
		//line template/build/show.qtpl:99
	}
	//line template/build/show.qtpl:99
	qw422016.N().S(` </td> </tr> <tr> <td>Finished at:</td> <td class="align-right"> `)
	//line template/build/show.qtpl:105
	if p.Build.FinishedAt.Valid {
		//line template/build/show.qtpl:105
		qw422016.N().S(` `)
		//line template/build/show.qtpl:106
		qw422016.E().S(p.Build.FinishedAt.Time.Format("2006-01-02T15:04:05"))
		//line template/build/show.qtpl:106
		qw422016.N().S(` `)
		//line template/build/show.qtpl:107
	} else {
		//line template/build/show.qtpl:107
		qw422016.N().S(` <span class="muted">--</span> `)
		//line template/build/show.qtpl:109
	}
	//line template/build/show.qtpl:109
	qw422016.N().S(` </td> </tr> <tr> <td>Duration:</td> <td class="align-right"> `)
	//line template/build/show.qtpl:115
	if !p.Build.FinishedAt.Valid || !p.Build.StartedAt.Valid {
		//line template/build/show.qtpl:115
		qw422016.N().S(` <span class="muted">--</span> `)
		//line template/build/show.qtpl:117
	} else {
		//line template/build/show.qtpl:117
		qw422016.N().S(` `)
		//line template/build/show.qtpl:118
		qw422016.E().V(p.Build.FinishedAt.Time.Sub(p.Build.StartedAt.Time))
		//line template/build/show.qtpl:118
		qw422016.N().S(` `)
		//line template/build/show.qtpl:119
	}
	//line template/build/show.qtpl:119
	qw422016.N().S(` </td> </tr> </table> </div> </div> </div> `)
	//line template/build/show.qtpl:126
	for _, s := range p.Build.Stages {
		//line template/build/show.qtpl:126
		qw422016.N().S(` `)
		//line template/build/show.qtpl:127
		if len(s.Jobs) > 0 {
			//line template/build/show.qtpl:127
			qw422016.N().S(` <div class="panel"> <div class="panel-header"><h3>`)
			//line template/build/show.qtpl:129
			qw422016.E().S(s.Name)
			//line template/build/show.qtpl:129
			qw422016.N().S(`</h3></div> `)
			//line template/build/show.qtpl:130
			job.StreamRenderTable(qw422016, s.Jobs)
			//line template/build/show.qtpl:130
			qw422016.N().S(` </div> `)
			//line template/build/show.qtpl:132
		}
		//line template/build/show.qtpl:132
		qw422016.N().S(` `)
		//line template/build/show.qtpl:133
	}
	//line template/build/show.qtpl:133
	qw422016.N().S(` `)
//line template/build/show.qtpl:134
}

//line template/build/show.qtpl:134
func (p *ShowPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:134
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:134
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:134
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:134
}

//line template/build/show.qtpl:134
func (p *ShowPage) Body() string {
	//line template/build/show.qtpl:134
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:134
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:134
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:134
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:134
	return qs422016
//line template/build/show.qtpl:134
}

//line template/build/show.qtpl:136
func (p *ShowManifest) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:136
	qw422016.N().S(` <div class="panel"> <div class="panel-header"> <ul class="panel-actions"> <li> <a class="btn btn-primary" href="`)
	//line template/build/show.qtpl:141
	qw422016.E().S(p.Build.UIEndpoint("manifest", "raw"))
	//line template/build/show.qtpl:141
	qw422016.N().S(`"> `)
	//line template/build/show.qtpl:142
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
	//line template/build/show.qtpl:142
	qw422016.N().S(`<span>Raw</span> </a> </li> </ul> </div> `)
	//line template/build/show.qtpl:147
	template.StreamRenderCode(qw422016, p.Build.Manifest)
	//line template/build/show.qtpl:147
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:149
}

//line template/build/show.qtpl:149
func (p *ShowManifest) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:149
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:149
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:149
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:149
}

//line template/build/show.qtpl:149
func (p *ShowManifest) Body() string {
	//line template/build/show.qtpl:149
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:149
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:149
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:149
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:149
	return qs422016
//line template/build/show.qtpl:149
}

//line template/build/show.qtpl:151
func (p *ShowOutput) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:151
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:153
	if p.Build.Output.Valid {
		//line template/build/show.qtpl:153
		qw422016.N().S(` <div class="panel-header"> <ul class="panel-actions"> <li> <a class="btn btn-primary" href="`)
		//line template/build/show.qtpl:157
		qw422016.E().S(p.Build.UIEndpoint("output", "raw"))
		//line template/build/show.qtpl:157
		qw422016.N().S(`"> `)
		//line template/build/show.qtpl:158
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
		//line template/build/show.qtpl:158
		qw422016.N().S(`<span>Raw</span> </a> </li> </ul> </div> `)
		//line template/build/show.qtpl:163
		template.StreamRenderCode(qw422016, p.Build.Output.String)
		//line template/build/show.qtpl:163
		qw422016.N().S(` `)
		//line template/build/show.qtpl:164
	} else {
		//line template/build/show.qtpl:164
		qw422016.N().S(` <div class="panel-message muted">No build output has been produced.</div> `)
		//line template/build/show.qtpl:166
	}
	//line template/build/show.qtpl:166
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:168
}

//line template/build/show.qtpl:168
func (p *ShowOutput) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:168
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:168
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:168
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:168
}

//line template/build/show.qtpl:168
func (p *ShowOutput) Body() string {
	//line template/build/show.qtpl:168
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:168
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:168
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:168
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:168
	return qs422016
//line template/build/show.qtpl:168
}

//line template/build/show.qtpl:170
func (p *ShowObjects) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:170
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:172
	if len(p.Objects) == 0 {
		//line template/build/show.qtpl:172
		qw422016.N().S(` <div class="panel-message muted">No objects have been placed for this build.</div> `)
		//line template/build/show.qtpl:174
	} else {
		//line template/build/show.qtpl:174
		qw422016.N().S(` `)
		//line template/build/show.qtpl:175
		StreamRenderObjectTable(qw422016, p.Objects)
		//line template/build/show.qtpl:175
		qw422016.N().S(` `)
		//line template/build/show.qtpl:176
	}
	//line template/build/show.qtpl:176
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:178
}

//line template/build/show.qtpl:178
func (p *ShowObjects) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:178
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:178
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:178
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:178
}

//line template/build/show.qtpl:178
func (p *ShowObjects) Body() string {
	//line template/build/show.qtpl:178
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:178
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:178
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:178
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:178
	return qs422016
//line template/build/show.qtpl:178
}

//line template/build/show.qtpl:180
func (p *ShowVariables) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:180
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:182
	if len(p.Variables) == 0 {
		//line template/build/show.qtpl:182
		qw422016.N().S(` <div class="panel-message muted">No variables have been set for this build.</div> `)
		//line template/build/show.qtpl:184
	} else {
		//line template/build/show.qtpl:184
		qw422016.N().S(` `)
		//line template/build/show.qtpl:185
		StreamRenderVariableTable(qw422016, p.Variables)
		//line template/build/show.qtpl:185
		qw422016.N().S(` `)
		//line template/build/show.qtpl:186
	}
	//line template/build/show.qtpl:186
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:188
}

//line template/build/show.qtpl:188
func (p *ShowVariables) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:188
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:188
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:188
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:188
}

//line template/build/show.qtpl:188
func (p *ShowVariables) Body() string {
	//line template/build/show.qtpl:188
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:188
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:188
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:188
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:188
	return qs422016
//line template/build/show.qtpl:188
}

//line template/build/show.qtpl:190
func (p *ShowPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:190
	qw422016.N().S(` <a href="/" class="back">`)
	//line template/build/show.qtpl:191
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
	//line template/build/show.qtpl:191
	qw422016.N().S(`</a> Build #`)
	//line template/build/show.qtpl:191
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:191
	qw422016.N().S(` `)
	//line template/build/show.qtpl:191
	template.StreamRenderStatus(qw422016, p.Build.Status)
	//line template/build/show.qtpl:191
	qw422016.N().S(` `)
//line template/build/show.qtpl:192
}

//line template/build/show.qtpl:192
func (p *ShowPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:192
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:192
	p.StreamHeader(qw422016)
	//line template/build/show.qtpl:192
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:192
}

//line template/build/show.qtpl:192
func (p *ShowPage) Header() string {
	//line template/build/show.qtpl:192
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:192
	p.WriteHeader(qb422016)
	//line template/build/show.qtpl:192
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:192
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:192
	return qs422016
//line template/build/show.qtpl:192
}

//line template/build/show.qtpl:194
func (p *ShowPage) StreamActions(qw422016 *qt422016.Writer) {
//line template/build/show.qtpl:194
}

//line template/build/show.qtpl:194
func (p *ShowPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:194
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:194
	p.StreamActions(qw422016)
	//line template/build/show.qtpl:194
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:194
}

//line template/build/show.qtpl:194
func (p *ShowPage) Actions() string {
	//line template/build/show.qtpl:194
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:194
	p.WriteActions(qb422016)
	//line template/build/show.qtpl:194
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:194
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:194
	return qs422016
//line template/build/show.qtpl:194
}

//line template/build/show.qtpl:197
func (p *ShowPage) StreamNavigation(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:197
	qw422016.N().S(`<li><a href="`)
	//line template/build/show.qtpl:199
	qw422016.E().S(p.Build.UIEndpoint())
	//line template/build/show.qtpl:199
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:199
	qw422016.E().S(template.Active(p.Build.UIEndpoint() == p.URI))
	//line template/build/show.qtpl:199
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:200
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 9c1.641 0 3 1.359 3 3s-1.359 3-3 3-3-1.359-3-3 1.359-3 3-3zM12 17.016c2.766 0 5.016-2.25 5.016-5.016s-2.25-5.016-5.016-5.016-5.016 2.25-5.016 5.016 2.25 5.016 5.016 5.016zM12 4.5c5.016 0 9.281 3.094 11.016 7.5-1.734 4.406-6 7.5-11.016 7.5s-9.281-3.094-11.016-7.5c1.734-4.406 6-7.5 11.016-7.5z"></path>
</svg>
`)
	//line template/build/show.qtpl:200
	qw422016.N().S(`<span>Overview</span></a></li><li><a href="`)
	//line template/build/show.qtpl:204
	qw422016.E().S(p.Build.UIEndpoint("manifest"))
	//line template/build/show.qtpl:204
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:204
	qw422016.E().S(template.Active(p.Build.UIEndpoint("manifest") == p.URI))
	//line template/build/show.qtpl:204
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:205
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M6.984 12.984v-1.969h14.016v1.969h-14.016zM6.984 18.984v-1.969h14.016v1.969h-14.016zM6.984 5.016h14.016v1.969h-14.016v-1.969zM2.016 11.016v-1.031h3v0.938l-1.828 2.063h1.828v1.031h-3v-0.938l1.781-2.063h-1.781zM3 8.016v-3h-0.984v-1.031h1.969v4.031h-0.984zM2.016 17.016v-1.031h3v4.031h-3v-1.031h1.969v-0.469h-0.984v-1.031h0.984v-0.469h-1.969z"></path>
</svg>
`)
	//line template/build/show.qtpl:205
	qw422016.N().S(`<span>Manifest</span></a></li><a href="`)
	//line template/build/show.qtpl:208
	qw422016.E().S(p.Build.UIEndpoint("objects"))
	//line template/build/show.qtpl:208
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:208
	qw422016.E().S(template.Active(p.Build.UIEndpoint("objects") == p.URI))
	//line template/build/show.qtpl:208
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:209
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM9 15.984v-6h-3.984l6.984-6.984 6.984 6.984h-3.984v6h-6z"></path>
</svg>
`)
	//line template/build/show.qtpl:209
	qw422016.N().S(`<span>Objects</span></a></li><a href="`)
	//line template/build/show.qtpl:212
	qw422016.E().S(p.Build.UIEndpoint("artifacts"))
	//line template/build/show.qtpl:212
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:212
	qw422016.E().S(template.Active(p.Build.UIEndpoint("artifacts") == p.URI))
	//line template/build/show.qtpl:212
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:213
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM18.984 9l-6.984 6.984-6.984-6.984h3.984v-6h6v6h3.984z"></path>
</svg>
`)
	//line template/build/show.qtpl:213
	qw422016.N().S(`<span>Artifacts</span></a></li><a href="`)
	//line template/build/show.qtpl:216
	qw422016.E().S(p.Build.UIEndpoint("variables"))
	//line template/build/show.qtpl:216
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:216
	qw422016.E().S(template.Active(p.Build.UIEndpoint("variables") == p.URI))
	//line template/build/show.qtpl:216
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:217
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
	//line template/build/show.qtpl:217
	qw422016.N().S(`<span>Variables</span></a></li><a href="`)
	//line template/build/show.qtpl:220
	qw422016.E().S(p.Build.UIEndpoint("output"))
	//line template/build/show.qtpl:220
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:220
	qw422016.E().S(template.Active(p.Build.UIEndpoint("output") == p.URI))
	//line template/build/show.qtpl:220
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:221
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
	//line template/build/show.qtpl:221
	qw422016.N().S(`<span>Output</span></a></li><a href="`)
	//line template/build/show.qtpl:224
	qw422016.E().S(p.Build.UIEndpoint("tags"))
	//line template/build/show.qtpl:224
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:224
	qw422016.E().S(template.Active(p.Build.UIEndpoint("tags") == p.URI))
	//line template/build/show.qtpl:224
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:225
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.625 5.859l4.359 6.141-4.359 6.141c-0.375 0.516-0.984 0.844-1.641 0.844h-10.969c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h10.969c0.656 0 1.266 0.328 1.641 0.844z"></path>
</svg>
`)
	//line template/build/show.qtpl:225
	qw422016.N().S(`<span>Tags</span></a></li>`)
//line template/build/show.qtpl:228
}

//line template/build/show.qtpl:228
func (p *ShowPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:228
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:228
	p.StreamNavigation(qw422016)
	//line template/build/show.qtpl:228
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:228
}

//line template/build/show.qtpl:228
func (p *ShowPage) Navigation() string {
	//line template/build/show.qtpl:228
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:228
	p.WriteNavigation(qb422016)
	//line template/build/show.qtpl:228
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:228
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:228
	return qs422016
//line template/build/show.qtpl:228
}

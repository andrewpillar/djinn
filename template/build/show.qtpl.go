// This file is automatically generated by qtc from "show.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/show.qtpl:2
package build

//line template/build/show.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/runner"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/artifact"
	"github.com/andrewpillar/thrall/template/job"
)

//line template/build/show.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/show.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/show.qtpl:12
type ShowPage struct {
	template.BasePage

	Build *model.Build
	CSRF  string
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

type ShowArtifacts struct {
	ShowPage

	Search    string
	Artifacts []*model.Artifact
}

type ShowVariables struct {
	ShowPage

	Variables []*model.BuildVariable
}

type ShowKeys struct {
	ShowPage

	Keys []*model.BuildKey
}

type ShowTags struct {
	ShowPage

	CSRF string
	Tags []*model.Tag
}

//line template/build/show.qtpl:61
func (p *ShowPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:61
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:62
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:62
	qw422016.N().S(` - Thrall `)
//line template/build/show.qtpl:63
}

//line template/build/show.qtpl:63
func (p *ShowPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:63
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:63
}

//line template/build/show.qtpl:63
func (p *ShowPage) Title() string {
	//line template/build/show.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:63
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:63
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:63
	return qs422016
//line template/build/show.qtpl:63
}

//line template/build/show.qtpl:65
func (p *ShowManifest) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:65
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:66
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:66
	qw422016.N().S(` - Manifest `)
//line template/build/show.qtpl:67
}

//line template/build/show.qtpl:67
func (p *ShowManifest) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:67
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:67
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:67
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:67
}

//line template/build/show.qtpl:67
func (p *ShowManifest) Title() string {
	//line template/build/show.qtpl:67
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:67
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:67
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:67
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:67
	return qs422016
//line template/build/show.qtpl:67
}

//line template/build/show.qtpl:69
func (p *ShowOutput) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:69
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:70
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:70
	qw422016.N().S(` - Output `)
//line template/build/show.qtpl:71
}

//line template/build/show.qtpl:71
func (p *ShowOutput) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:71
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:71
}

//line template/build/show.qtpl:71
func (p *ShowOutput) Title() string {
	//line template/build/show.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:71
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:71
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:71
	return qs422016
//line template/build/show.qtpl:71
}

//line template/build/show.qtpl:73
func (p *ShowObjects) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:73
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:74
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:74
	qw422016.N().S(` - Objects `)
//line template/build/show.qtpl:75
}

//line template/build/show.qtpl:75
func (p *ShowObjects) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:75
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:75
}

//line template/build/show.qtpl:75
func (p *ShowObjects) Title() string {
	//line template/build/show.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:75
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:75
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:75
	return qs422016
//line template/build/show.qtpl:75
}

//line template/build/show.qtpl:77
func (p *ShowVariables) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:77
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:78
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:78
	qw422016.N().S(` - Variables `)
//line template/build/show.qtpl:79
}

//line template/build/show.qtpl:79
func (p *ShowVariables) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:79
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:79
}

//line template/build/show.qtpl:79
func (p *ShowVariables) Title() string {
	//line template/build/show.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:79
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:79
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:79
	return qs422016
//line template/build/show.qtpl:79
}

//line template/build/show.qtpl:81
func (p *ShowPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:81
	qw422016.N().S(` <div class="mb-10 overflow"> <div class="col-75 pr-5 left"> <div class="panel"> <div class="panel-header"> `)
	//line template/build/show.qtpl:86
	if p.Build.Namespace.IsZero() {
		//line template/build/show.qtpl:86
		qw422016.N().S(` <h3>Submitted by `)
		//line template/build/show.qtpl:87
		qw422016.E().S(p.Build.User.Username)
		//line template/build/show.qtpl:87
		qw422016.N().S(` &lt;`)
		//line template/build/show.qtpl:87
		qw422016.E().S(p.Build.User.Email)
		//line template/build/show.qtpl:87
		qw422016.N().S(`&gt;</h3> `)
		//line template/build/show.qtpl:88
	} else {
		//line template/build/show.qtpl:88
		qw422016.N().S(` <h3>Submitted by `)
		//line template/build/show.qtpl:89
		qw422016.E().S(p.Build.User.Username)
		//line template/build/show.qtpl:89
		qw422016.N().S(` &lt;`)
		//line template/build/show.qtpl:89
		qw422016.E().S(p.Build.User.Email)
		//line template/build/show.qtpl:89
		qw422016.N().S(`&gt; to <a href="`)
		//line template/build/show.qtpl:89
		qw422016.E().S(p.Build.Namespace.UIEndpoint())
		//line template/build/show.qtpl:89
		qw422016.N().S(`">`)
		//line template/build/show.qtpl:89
		qw422016.E().S(p.Build.Namespace.Path)
		//line template/build/show.qtpl:89
		qw422016.N().S(`</a></h3> `)
		//line template/build/show.qtpl:90
	}
	//line template/build/show.qtpl:90
	qw422016.N().S(` </div> <div class="panel-body"> `)
	//line template/build/show.qtpl:93
	if p.Build.Trigger.Comment != "" {
		//line template/build/show.qtpl:93
		qw422016.N().S(` <pre class="code">`)
		//line template/build/show.qtpl:94
		qw422016.E().S(p.Build.Trigger.Comment)
		//line template/build/show.qtpl:94
		qw422016.N().S(`</pre> `)
		//line template/build/show.qtpl:95
	} else {
		//line template/build/show.qtpl:95
		qw422016.N().S(` <em class="muted">No build comment.</em> `)
		//line template/build/show.qtpl:97
	}
	//line template/build/show.qtpl:97
	qw422016.N().S(` </div> <div class="panel-footer"> <div class="mb-10"> `)
	//line template/build/show.qtpl:101
	if len(p.Build.Tags) > 0 {
		//line template/build/show.qtpl:101
		qw422016.N().S(` `)
		//line template/build/show.qtpl:102
		for _, t := range p.Build.Tags {
			//line template/build/show.qtpl:102
			qw422016.N().S(` <a href="/?tag=`)
			//line template/build/show.qtpl:103
			qw422016.E().S(t.Name)
			//line template/build/show.qtpl:103
			qw422016.N().S(`" class="pill pill-light">`)
			//line template/build/show.qtpl:103
			qw422016.E().S(t.Name)
			//line template/build/show.qtpl:103
			qw422016.N().S(`</a> `)
			//line template/build/show.qtpl:104
		}
		//line template/build/show.qtpl:104
		qw422016.N().S(` `)
		//line template/build/show.qtpl:105
	} else {
		//line template/build/show.qtpl:105
		qw422016.N().S(` <em class="muted">No build tags.</em> `)
		//line template/build/show.qtpl:107
	}
	//line template/build/show.qtpl:107
	qw422016.N().S(` </div> </div> </div> </div> <div class="col-25 pl-5 right"> <div class="panel"> <table class="table"> <tr> <td>Started at:</td> <td class="align-right"> `)
	//line template/build/show.qtpl:118
	if p.Build.StartedAt.Valid {
		//line template/build/show.qtpl:118
		qw422016.N().S(` `)
		//line template/build/show.qtpl:119
		qw422016.E().S(p.Build.StartedAt.Time.Format("2006-01-02T15:04:05"))
		//line template/build/show.qtpl:119
		qw422016.N().S(` `)
		//line template/build/show.qtpl:120
	} else {
		//line template/build/show.qtpl:120
		qw422016.N().S(` <span class="muted">--</span> `)
		//line template/build/show.qtpl:122
	}
	//line template/build/show.qtpl:122
	qw422016.N().S(` </td> </tr> <tr> <td>Finished at:</td> <td class="align-right"> `)
	//line template/build/show.qtpl:128
	if p.Build.FinishedAt.Valid {
		//line template/build/show.qtpl:128
		qw422016.N().S(` `)
		//line template/build/show.qtpl:129
		qw422016.E().S(p.Build.FinishedAt.Time.Format("2006-01-02T15:04:05"))
		//line template/build/show.qtpl:129
		qw422016.N().S(` `)
		//line template/build/show.qtpl:130
	} else {
		//line template/build/show.qtpl:130
		qw422016.N().S(` <span class="muted">--</span> `)
		//line template/build/show.qtpl:132
	}
	//line template/build/show.qtpl:132
	qw422016.N().S(` </td> </tr> <tr> <td>Duration:</td> <td class="align-right"> `)
	//line template/build/show.qtpl:138
	if !p.Build.FinishedAt.Valid || !p.Build.StartedAt.Valid {
		//line template/build/show.qtpl:138
		qw422016.N().S(` <span class="muted">--</span> `)
		//line template/build/show.qtpl:140
	} else {
		//line template/build/show.qtpl:140
		qw422016.N().S(` `)
		//line template/build/show.qtpl:141
		qw422016.E().V(p.Build.FinishedAt.Time.Sub(p.Build.StartedAt.Time))
		//line template/build/show.qtpl:141
		qw422016.N().S(` `)
		//line template/build/show.qtpl:142
	}
	//line template/build/show.qtpl:142
	qw422016.N().S(` </td> </tr> </table> </div> </div> </div> `)
	//line template/build/show.qtpl:149
	for _, s := range p.Build.Stages {
		//line template/build/show.qtpl:149
		qw422016.N().S(` `)
		//line template/build/show.qtpl:150
		if len(s.Jobs) > 0 {
			//line template/build/show.qtpl:150
			qw422016.N().S(` <div class="panel"> <div class="panel-header"><h3>`)
			//line template/build/show.qtpl:152
			qw422016.E().S(s.Name)
			//line template/build/show.qtpl:152
			qw422016.N().S(`</h3></div> `)
			//line template/build/show.qtpl:153
			job.StreamRenderTable(qw422016, s.Jobs)
			//line template/build/show.qtpl:153
			qw422016.N().S(` </div> `)
			//line template/build/show.qtpl:155
		}
		//line template/build/show.qtpl:155
		qw422016.N().S(` `)
		//line template/build/show.qtpl:156
	}
	//line template/build/show.qtpl:156
	qw422016.N().S(` `)
//line template/build/show.qtpl:157
}

//line template/build/show.qtpl:157
func (p *ShowPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:157
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:157
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:157
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:157
}

//line template/build/show.qtpl:157
func (p *ShowPage) Body() string {
	//line template/build/show.qtpl:157
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:157
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:157
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:157
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:157
	return qs422016
//line template/build/show.qtpl:157
}

//line template/build/show.qtpl:159
func (p *ShowManifest) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:159
	qw422016.N().S(` <div class="panel"> <div class="panel-header"> <ul class="panel-actions"> <li> <a class="btn btn-primary" href="`)
	//line template/build/show.qtpl:164
	qw422016.E().S(p.Build.UIEndpoint("manifest", "raw"))
	//line template/build/show.qtpl:164
	qw422016.N().S(`"> `)
	//line template/build/show.qtpl:165
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
	//line template/build/show.qtpl:165
	qw422016.N().S(`<span>Raw</span> </a> </li> </ul> </div> `)
	//line template/build/show.qtpl:170
	template.StreamRenderCode(qw422016, p.Build.Manifest)
	//line template/build/show.qtpl:170
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:172
}

//line template/build/show.qtpl:172
func (p *ShowManifest) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:172
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:172
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:172
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:172
}

//line template/build/show.qtpl:172
func (p *ShowManifest) Body() string {
	//line template/build/show.qtpl:172
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:172
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:172
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:172
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:172
	return qs422016
//line template/build/show.qtpl:172
}

//line template/build/show.qtpl:174
func (p *ShowOutput) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:174
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:176
	if p.Build.Output.Valid {
		//line template/build/show.qtpl:176
		qw422016.N().S(` <div class="panel-header"> <ul class="panel-actions"> <li> <a class="btn btn-primary" href="`)
		//line template/build/show.qtpl:180
		qw422016.E().S(p.Build.UIEndpoint("output", "raw"))
		//line template/build/show.qtpl:180
		qw422016.N().S(`"> `)
		//line template/build/show.qtpl:181
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
		//line template/build/show.qtpl:181
		qw422016.N().S(`<span>Raw</span> </a> </li> </ul> </div> `)
		//line template/build/show.qtpl:186
		template.StreamRenderCode(qw422016, p.Build.Output.String)
		//line template/build/show.qtpl:186
		qw422016.N().S(` `)
		//line template/build/show.qtpl:187
	} else {
		//line template/build/show.qtpl:187
		qw422016.N().S(` <div class="panel-message muted">No build output has been produced.</div> `)
		//line template/build/show.qtpl:189
	}
	//line template/build/show.qtpl:189
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:191
}

//line template/build/show.qtpl:191
func (p *ShowOutput) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:191
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:191
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:191
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:191
}

//line template/build/show.qtpl:191
func (p *ShowOutput) Body() string {
	//line template/build/show.qtpl:191
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:191
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:191
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:191
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:191
	return qs422016
//line template/build/show.qtpl:191
}

//line template/build/show.qtpl:193
func (p *ShowObjects) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:193
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:195
	if len(p.Objects) == 0 {
		//line template/build/show.qtpl:195
		qw422016.N().S(` <div class="panel-message muted">No objects have been placed for this build.</div> `)
		//line template/build/show.qtpl:197
	} else {
		//line template/build/show.qtpl:197
		qw422016.N().S(` `)
		//line template/build/show.qtpl:198
		StreamRenderObjectTable(qw422016, p.Objects)
		//line template/build/show.qtpl:198
		qw422016.N().S(` `)
		//line template/build/show.qtpl:199
	}
	//line template/build/show.qtpl:199
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:201
}

//line template/build/show.qtpl:201
func (p *ShowObjects) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:201
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:201
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:201
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:201
}

//line template/build/show.qtpl:201
func (p *ShowObjects) Body() string {
	//line template/build/show.qtpl:201
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:201
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:201
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:201
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:201
	return qs422016
//line template/build/show.qtpl:201
}

//line template/build/show.qtpl:203
func (p *ShowArtifacts) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:203
	qw422016.N().S(` <div class="panel">`)
	//line template/build/show.qtpl:204
	artifact.StreamRenderIndex(qw422016, p.Artifacts, p.URL.Path, p.Search)
	//line template/build/show.qtpl:204
	qw422016.N().S(`</div> `)
//line template/build/show.qtpl:205
}

//line template/build/show.qtpl:205
func (p *ShowArtifacts) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:205
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:205
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:205
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:205
}

//line template/build/show.qtpl:205
func (p *ShowArtifacts) Body() string {
	//line template/build/show.qtpl:205
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:205
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:205
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:205
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:205
	return qs422016
//line template/build/show.qtpl:205
}

//line template/build/show.qtpl:207
func (p *ShowVariables) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:207
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:209
	if len(p.Variables) == 0 {
		//line template/build/show.qtpl:209
		qw422016.N().S(` <div class="panel-message muted">No variables have been set for this build.</div> `)
		//line template/build/show.qtpl:211
	} else {
		//line template/build/show.qtpl:211
		qw422016.N().S(` `)
		//line template/build/show.qtpl:212
		StreamRenderVariableTable(qw422016, p.Variables)
		//line template/build/show.qtpl:212
		qw422016.N().S(` `)
		//line template/build/show.qtpl:213
	}
	//line template/build/show.qtpl:213
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:215
}

//line template/build/show.qtpl:215
func (p *ShowVariables) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:215
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:215
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:215
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:215
}

//line template/build/show.qtpl:215
func (p *ShowVariables) Body() string {
	//line template/build/show.qtpl:215
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:215
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:215
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:215
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:215
	return qs422016
//line template/build/show.qtpl:215
}

//line template/build/show.qtpl:217
func (p *ShowKeys) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:217
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:219
	if len(p.Keys) == 0 {
		//line template/build/show.qtpl:219
		qw422016.N().S(` <div class="panel-message muted">No keys have been added to this build.</div> `)
		//line template/build/show.qtpl:221
	} else {
		//line template/build/show.qtpl:221
		qw422016.N().S(` `)
		//line template/build/show.qtpl:222
		StreamRenderKeyTable(qw422016, p.Keys)
		//line template/build/show.qtpl:222
		qw422016.N().S(` `)
		//line template/build/show.qtpl:223
	}
	//line template/build/show.qtpl:223
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:225
}

//line template/build/show.qtpl:225
func (p *ShowKeys) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:225
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:225
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:225
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:225
}

//line template/build/show.qtpl:225
func (p *ShowKeys) Body() string {
	//line template/build/show.qtpl:225
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:225
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:225
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:225
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:225
	return qs422016
//line template/build/show.qtpl:225
}

//line template/build/show.qtpl:227
func (p *ShowTags) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:227
	qw422016.N().S(` <div class="panel"> <div class="panel-header panel-body"> <form method="POST" action="`)
	//line template/build/show.qtpl:230
	qw422016.E().S(p.Build.UIEndpoint("tags"))
	//line template/build/show.qtpl:230
	qw422016.N().S(`"> `)
	//line template/build/show.qtpl:231
	qw422016.N().S(p.CSRF)
	//line template/build/show.qtpl:231
	qw422016.N().S(` <div class="form-field form-field-inline"> <input type="text" class="form-text" name="tags" placeholder="Tag this build..." autocomplete="off"/> <button type="submit" class="btn btn-primary">Tag</button> </div> </form> </div> `)
	//line template/build/show.qtpl:238
	if len(p.Tags) == 0 {
		//line template/build/show.qtpl:238
		qw422016.N().S(` <div class="panel-message muted">No tags have been set for this build.</div> `)
		//line template/build/show.qtpl:240
	} else {
		//line template/build/show.qtpl:240
		qw422016.N().S(` `)
		//line template/build/show.qtpl:241
		StreamRenderTagTable(qw422016, p.Tags, p.CSRF)
		//line template/build/show.qtpl:241
		qw422016.N().S(` `)
		//line template/build/show.qtpl:242
	}
	//line template/build/show.qtpl:242
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:244
}

//line template/build/show.qtpl:244
func (p *ShowTags) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:244
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:244
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:244
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:244
}

//line template/build/show.qtpl:244
func (p *ShowTags) Body() string {
	//line template/build/show.qtpl:244
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:244
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:244
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:244
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:244
	return qs422016
//line template/build/show.qtpl:244
}

//line template/build/show.qtpl:246
func (p *ShowPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:246
	qw422016.N().S(` <a href="/" class="back">`)
	//line template/build/show.qtpl:247
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
	//line template/build/show.qtpl:247
	qw422016.N().S(`</a> `)
	//line template/build/show.qtpl:248
	if !p.Build.Namespace.IsZero() {
		//line template/build/show.qtpl:248
		qw422016.N().S(` <a href="`)
		//line template/build/show.qtpl:249
		qw422016.E().S(p.Build.Namespace.UIEndpoint())
		//line template/build/show.qtpl:249
		qw422016.N().S(`">`)
		//line template/build/show.qtpl:249
		qw422016.E().S(p.Build.Namespace.Name)
		//line template/build/show.qtpl:249
		qw422016.N().S(`</a> / `)
		//line template/build/show.qtpl:250
	}
	//line template/build/show.qtpl:250
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:251
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:251
	qw422016.N().S(` `)
	//line template/build/show.qtpl:251
	template.StreamRenderStatus(qw422016, p.Build.Status)
	//line template/build/show.qtpl:251
	qw422016.N().S(` `)
//line template/build/show.qtpl:252
}

//line template/build/show.qtpl:252
func (p *ShowPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:252
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:252
	p.StreamHeader(qw422016)
	//line template/build/show.qtpl:252
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:252
}

//line template/build/show.qtpl:252
func (p *ShowPage) Header() string {
	//line template/build/show.qtpl:252
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:252
	p.WriteHeader(qb422016)
	//line template/build/show.qtpl:252
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:252
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:252
	return qs422016
//line template/build/show.qtpl:252
}

//line template/build/show.qtpl:254
func (p *ShowPage) StreamActions(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:254
	qw422016.N().S(` `)
	//line template/build/show.qtpl:255
	if p.User.ID == p.Build.UserID && p.Build.Status == runner.Running {
		//line template/build/show.qtpl:255
		qw422016.N().S(` <li> <form method="POST" action="`)
		//line template/build/show.qtpl:257
		qw422016.E().S(p.Build.UIEndpoint())
		//line template/build/show.qtpl:257
		qw422016.N().S(`"> `)
		//line template/build/show.qtpl:258
		qw422016.N().S(string(p.CSRF))
		//line template/build/show.qtpl:258
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"> <button type="submit" class="btn btn-danger">Kill</button> </form> </li> `)
		//line template/build/show.qtpl:263
	}
	//line template/build/show.qtpl:263
	qw422016.N().S(` `)
//line template/build/show.qtpl:264
}

//line template/build/show.qtpl:264
func (p *ShowPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:264
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:264
	p.StreamActions(qw422016)
	//line template/build/show.qtpl:264
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:264
}

//line template/build/show.qtpl:264
func (p *ShowPage) Actions() string {
	//line template/build/show.qtpl:264
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:264
	p.WriteActions(qb422016)
	//line template/build/show.qtpl:264
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:264
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:264
	return qs422016
//line template/build/show.qtpl:264
}

//line template/build/show.qtpl:267
func (p *ShowPage) StreamNavigation(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:267
	qw422016.N().S(`<li><a href="`)
	//line template/build/show.qtpl:269
	qw422016.E().S(p.Build.UIEndpoint())
	//line template/build/show.qtpl:269
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:269
	qw422016.E().S(template.Active(p.Build.UIEndpoint() == p.URL.Path))
	//line template/build/show.qtpl:269
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:270
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 9c1.641 0 3 1.359 3 3s-1.359 3-3 3-3-1.359-3-3 1.359-3 3-3zM12 17.016c2.766 0 5.016-2.25 5.016-5.016s-2.25-5.016-5.016-5.016-5.016 2.25-5.016 5.016 2.25 5.016 5.016 5.016zM12 4.5c5.016 0 9.281 3.094 11.016 7.5-1.734 4.406-6 7.5-11.016 7.5s-9.281-3.094-11.016-7.5c1.734-4.406 6-7.5 11.016-7.5z"></path>
</svg>
`)
	//line template/build/show.qtpl:270
	qw422016.N().S(`<span>Overview</span></a></li><li><a href="`)
	//line template/build/show.qtpl:274
	qw422016.E().S(p.Build.UIEndpoint("manifest"))
	//line template/build/show.qtpl:274
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:274
	qw422016.E().S(template.Active(p.Build.UIEndpoint("manifest") == p.URL.Path))
	//line template/build/show.qtpl:274
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:275
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M6.984 12.984v-1.969h14.016v1.969h-14.016zM6.984 18.984v-1.969h14.016v1.969h-14.016zM6.984 5.016h14.016v1.969h-14.016v-1.969zM2.016 11.016v-1.031h3v0.938l-1.828 2.063h1.828v1.031h-3v-0.938l1.781-2.063h-1.781zM3 8.016v-3h-0.984v-1.031h1.969v4.031h-0.984zM2.016 17.016v-1.031h3v4.031h-3v-1.031h1.969v-0.469h-0.984v-1.031h0.984v-0.469h-1.969z"></path>
</svg>
`)
	//line template/build/show.qtpl:275
	qw422016.N().S(`<span>Manifest</span></a></li><li><a href="`)
	//line template/build/show.qtpl:279
	qw422016.E().S(p.Build.UIEndpoint("objects"))
	//line template/build/show.qtpl:279
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:279
	qw422016.E().S(template.Active(p.Build.UIEndpoint("objects") == p.URL.Path))
	//line template/build/show.qtpl:279
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:280
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM9 15.984v-6h-3.984l6.984-6.984 6.984 6.984h-3.984v6h-6z"></path>
</svg>
`)
	//line template/build/show.qtpl:280
	qw422016.N().S(`<span>Objects</span></a></li><li><a href="`)
	//line template/build/show.qtpl:284
	qw422016.E().S(p.Build.UIEndpoint("artifacts"))
	//line template/build/show.qtpl:284
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:284
	qw422016.E().S(template.Active(p.Build.UIEndpoint("artifacts") == p.URL.Path))
	//line template/build/show.qtpl:284
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:285
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM18.984 9l-6.984 6.984-6.984-6.984h3.984v-6h6v6h3.984z"></path>
</svg>
`)
	//line template/build/show.qtpl:285
	qw422016.N().S(`<span>Artifacts</span></a></li><li><a href="`)
	//line template/build/show.qtpl:289
	qw422016.E().S(p.Build.UIEndpoint("variables"))
	//line template/build/show.qtpl:289
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:289
	qw422016.E().S(template.Active(p.Build.UIEndpoint("variables") == p.URL.Path))
	//line template/build/show.qtpl:289
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:290
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
	//line template/build/show.qtpl:290
	qw422016.N().S(`<span>Variables</span></a></li><li><a href="`)
	//line template/build/show.qtpl:294
	qw422016.E().S(p.Build.UIEndpoint("keys"))
	//line template/build/show.qtpl:294
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:294
	qw422016.E().S(template.Active(p.Build.UIEndpoint("keys") == p.URL.Path))
	//line template/build/show.qtpl:294
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:295
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M6.984 14.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-1.969 0.938-1.969 2.016 0.891 2.016 1.969 2.016zM12.656 9.984h10.359v4.031h-2.016v3.984h-3.984v-3.984h-4.359c-0.797 2.344-3.047 3.984-5.672 3.984-3.328 0-6-2.672-6-6s2.672-6 6-6c2.625 0 4.875 1.641 5.672 3.984z"></path>
</svg>
`)
	//line template/build/show.qtpl:295
	qw422016.N().S(`<span>Keys</span></a></li><li><a href="`)
	//line template/build/show.qtpl:299
	qw422016.E().S(p.Build.UIEndpoint("output"))
	//line template/build/show.qtpl:299
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:299
	qw422016.E().S(template.Active(p.Build.UIEndpoint("output") == p.URL.Path))
	//line template/build/show.qtpl:299
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:300
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
	//line template/build/show.qtpl:300
	qw422016.N().S(`<span>Output</span></a></li><li><a href="`)
	//line template/build/show.qtpl:304
	qw422016.E().S(p.Build.UIEndpoint("tags"))
	//line template/build/show.qtpl:304
	qw422016.N().S(`" class="`)
	//line template/build/show.qtpl:304
	qw422016.E().S(template.Active(p.Build.UIEndpoint("tags") == p.URL.Path))
	//line template/build/show.qtpl:304
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:305
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.625 5.859l4.359 6.141-4.359 6.141c-0.375 0.516-0.984 0.844-1.641 0.844h-10.969c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h10.969c0.656 0 1.266 0.328 1.641 0.844z"></path>
</svg>
`)
	//line template/build/show.qtpl:305
	qw422016.N().S(`<span>Tags</span></a></li>`)
//line template/build/show.qtpl:308
}

//line template/build/show.qtpl:308
func (p *ShowPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:308
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:308
	p.StreamNavigation(qw422016)
	//line template/build/show.qtpl:308
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:308
}

//line template/build/show.qtpl:308
func (p *ShowPage) Navigation() string {
	//line template/build/show.qtpl:308
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:308
	p.WriteNavigation(qb422016)
	//line template/build/show.qtpl:308
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:308
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:308
	return qs422016
//line template/build/show.qtpl:308
}

// This file is automatically generated by qtc from "show.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/show.qtpl:2
package build

//line template/build/show.qtpl:2
import (
	"fmt"

	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/artifact"
	"github.com/andrewpillar/thrall/template/job"
)

//line template/build/show.qtpl:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/show.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/show.qtpl:13
type ShowPage struct {
	template.Page

	Build *model.Build

	ShowManifest  bool
	ShowObjects   bool
	ShowArtifacts bool
	ShowVariables bool
	ShowOutput    bool
}

//line template/build/show.qtpl:27
func (p *ShowPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:27
	qw422016.N().S(` Build #`)
	//line template/build/show.qtpl:28
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:28
	qw422016.N().S(` - Thrall `)
//line template/build/show.qtpl:29
}

//line template/build/show.qtpl:29
func (p *ShowPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:29
	p.StreamTitle(qw422016)
	//line template/build/show.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:29
}

//line template/build/show.qtpl:29
func (p *ShowPage) Title() string {
	//line template/build/show.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:29
	p.WriteTitle(qb422016)
	//line template/build/show.qtpl:29
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:29
	return qs422016
//line template/build/show.qtpl:29
}

//line template/build/show.qtpl:31
func (p *ShowPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:31
	qw422016.N().S(` `)
	//line template/build/show.qtpl:32
	if p.ShowManifest {
		//line template/build/show.qtpl:32
		qw422016.N().S(` `)
		//line template/build/show.qtpl:33
		p.streamrenderManifest(qw422016)
		//line template/build/show.qtpl:33
		qw422016.N().S(` `)
		//line template/build/show.qtpl:34
	} else if p.ShowObjects {
		//line template/build/show.qtpl:34
		qw422016.N().S(` `)
		//line template/build/show.qtpl:35
		p.streamrenderObjects(qw422016)
		//line template/build/show.qtpl:35
		qw422016.N().S(` `)
		//line template/build/show.qtpl:36
	} else if p.ShowArtifacts {
		//line template/build/show.qtpl:36
		qw422016.N().S(` `)
		//line template/build/show.qtpl:37
		p.streamrenderArtifacts(qw422016)
		//line template/build/show.qtpl:37
		qw422016.N().S(` `)
		//line template/build/show.qtpl:38
	} else if p.ShowVariables {
		//line template/build/show.qtpl:38
		qw422016.N().S(` `)
		//line template/build/show.qtpl:39
		p.streamrenderVariables(qw422016)
		//line template/build/show.qtpl:39
		qw422016.N().S(` `)
		//line template/build/show.qtpl:40
	} else if p.ShowOutput {
		//line template/build/show.qtpl:40
		qw422016.N().S(` `)
		//line template/build/show.qtpl:41
		p.streamrenderOutput(qw422016)
		//line template/build/show.qtpl:41
		qw422016.N().S(` `)
		//line template/build/show.qtpl:42
	} else {
		//line template/build/show.qtpl:42
		qw422016.N().S(` <div class="mb-10 overflow"> <div class="col-75 pr-5 left"> <div class="panel"> <div class="panel-header"> `)
		//line template/build/show.qtpl:47
		if p.Build.Namespace.IsZero() {
			//line template/build/show.qtpl:47
			qw422016.N().S(` <h3>Submitted by `)
			//line template/build/show.qtpl:48
			qw422016.E().S(p.Build.User.Username)
			//line template/build/show.qtpl:48
			qw422016.N().S(` &lt;`)
			//line template/build/show.qtpl:48
			qw422016.E().S(p.Build.User.Email)
			//line template/build/show.qtpl:48
			qw422016.N().S(`&gt;</h3> `)
			//line template/build/show.qtpl:49
		} else {
			//line template/build/show.qtpl:49
			qw422016.N().S(` <h3>Submitted by `)
			//line template/build/show.qtpl:50
			qw422016.E().S(p.Build.User.Username)
			//line template/build/show.qtpl:50
			qw422016.N().S(` &lt;`)
			//line template/build/show.qtpl:50
			qw422016.E().S(p.Build.User.Email)
			//line template/build/show.qtpl:50
			qw422016.N().S(`&gt; to <a href="`)
			//line template/build/show.qtpl:50
			qw422016.E().S(p.Build.Namespace.UIEndpoint())
			//line template/build/show.qtpl:50
			qw422016.N().S(`">`)
			//line template/build/show.qtpl:50
			qw422016.E().S(p.Build.Namespace.Path)
			//line template/build/show.qtpl:50
			qw422016.N().S(`</a></h3> `)
			//line template/build/show.qtpl:51
		}
		//line template/build/show.qtpl:51
		qw422016.N().S(` </div> <div class="panel-body"> `)
		//line template/build/show.qtpl:54
		if p.Build.Trigger.Comment != "" {
			//line template/build/show.qtpl:54
			qw422016.N().S(` <pre class="code">`)
			//line template/build/show.qtpl:55
			qw422016.E().S(p.Build.Trigger.Comment)
			//line template/build/show.qtpl:55
			qw422016.N().S(`</pre> `)
			//line template/build/show.qtpl:56
		} else {
			//line template/build/show.qtpl:56
			qw422016.N().S(` <em class="muted">No build comment.</em> `)
			//line template/build/show.qtpl:58
		}
		//line template/build/show.qtpl:58
		qw422016.N().S(` </div> </div> </div> <div class="col-25 pl-5 right"> <div class="panel"> <table class="table"> <tr> <td>Started at:</td> <td class="align-right"> `)
		//line template/build/show.qtpl:68
		if p.Build.StartedAt.Valid {
			//line template/build/show.qtpl:68
			qw422016.N().S(` `)
			//line template/build/show.qtpl:69
			qw422016.E().S(p.Build.StartedAt.Time.Format("2006-01-02T15:04:05"))
			//line template/build/show.qtpl:69
			qw422016.N().S(` `)
			//line template/build/show.qtpl:70
		} else {
			//line template/build/show.qtpl:70
			qw422016.N().S(` <span class="muted">--</span> `)
			//line template/build/show.qtpl:72
		}
		//line template/build/show.qtpl:72
		qw422016.N().S(` </td> </tr> <tr> <td>Finished at:</td> <td class="align-right"> `)
		//line template/build/show.qtpl:78
		if p.Build.FinishedAt.Valid {
			//line template/build/show.qtpl:78
			qw422016.N().S(` `)
			//line template/build/show.qtpl:79
			qw422016.E().S(p.Build.FinishedAt.Time.Format("2006-01-02T15:04:05"))
			//line template/build/show.qtpl:79
			qw422016.N().S(` `)
			//line template/build/show.qtpl:80
		} else {
			//line template/build/show.qtpl:80
			qw422016.N().S(` <span class="muted">--</span> `)
			//line template/build/show.qtpl:82
		}
		//line template/build/show.qtpl:82
		qw422016.N().S(` </td> </tr> <tr> <td>Duration:</td> <td class="align-right"> `)
		//line template/build/show.qtpl:88
		if !p.Build.FinishedAt.Valid || !p.Build.StartedAt.Valid {
			//line template/build/show.qtpl:88
			qw422016.N().S(` <span class="muted">--</span> `)
			//line template/build/show.qtpl:90
		} else {
			//line template/build/show.qtpl:90
			qw422016.N().S(` `)
			//line template/build/show.qtpl:91
			qw422016.E().V(p.Build.FinishedAt.Time.Sub(p.Build.StartedAt.Time))
			//line template/build/show.qtpl:91
			qw422016.N().S(` `)
			//line template/build/show.qtpl:92
		}
		//line template/build/show.qtpl:92
		qw422016.N().S(` </td> </tr> </table> </div> </div> </div> `)
		//line template/build/show.qtpl:99
		for _, s := range p.Build.Stages {
			//line template/build/show.qtpl:99
			qw422016.N().S(` `)
			//line template/build/show.qtpl:100
			if len(s.Jobs) > 0 {
				//line template/build/show.qtpl:100
				qw422016.N().S(` <div class="panel"> <div class="panel-header"><h3>`)
				//line template/build/show.qtpl:102
				qw422016.E().S(s.Name)
				//line template/build/show.qtpl:102
				qw422016.N().S(`</h3></div> `)
				//line template/build/show.qtpl:103
				job.StreamRenderTable(qw422016, s.Jobs)
				//line template/build/show.qtpl:103
				qw422016.N().S(` </div> `)
				//line template/build/show.qtpl:105
			}
			//line template/build/show.qtpl:105
			qw422016.N().S(` `)
			//line template/build/show.qtpl:106
		}
		//line template/build/show.qtpl:106
		qw422016.N().S(` `)
		//line template/build/show.qtpl:107
	}
	//line template/build/show.qtpl:107
	qw422016.N().S(` `)
//line template/build/show.qtpl:108
}

//line template/build/show.qtpl:108
func (p *ShowPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:108
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:108
	p.StreamBody(qw422016)
	//line template/build/show.qtpl:108
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:108
}

//line template/build/show.qtpl:108
func (p *ShowPage) Body() string {
	//line template/build/show.qtpl:108
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:108
	p.WriteBody(qb422016)
	//line template/build/show.qtpl:108
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:108
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:108
	return qs422016
//line template/build/show.qtpl:108
}

//line template/build/show.qtpl:110
func (p *ShowPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:110
	qw422016.N().S(` <a href="/" class="back">`)
	//line template/build/show.qtpl:111
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
	//line template/build/show.qtpl:111
	qw422016.N().S(`</a> Build #`)
	//line template/build/show.qtpl:111
	qw422016.E().V(p.Build.ID)
	//line template/build/show.qtpl:111
	qw422016.N().S(` `)
	//line template/build/show.qtpl:111
	template.StreamRenderStatus(qw422016, p.Build.Status)
	//line template/build/show.qtpl:111
	qw422016.N().S(` `)
//line template/build/show.qtpl:112
}

//line template/build/show.qtpl:112
func (p *ShowPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:112
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:112
	p.StreamHeader(qw422016)
	//line template/build/show.qtpl:112
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:112
}

//line template/build/show.qtpl:112
func (p *ShowPage) Header() string {
	//line template/build/show.qtpl:112
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:112
	p.WriteHeader(qb422016)
	//line template/build/show.qtpl:112
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:112
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:112
	return qs422016
//line template/build/show.qtpl:112
}

//line template/build/show.qtpl:114
func (p *ShowPage) StreamActions(qw422016 *qt422016.Writer) {
//line template/build/show.qtpl:114
}

//line template/build/show.qtpl:114
func (p *ShowPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:114
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:114
	p.StreamActions(qw422016)
	//line template/build/show.qtpl:114
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:114
}

//line template/build/show.qtpl:114
func (p *ShowPage) Actions() string {
	//line template/build/show.qtpl:114
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:114
	p.WriteActions(qb422016)
	//line template/build/show.qtpl:114
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:114
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:114
	return qs422016
//line template/build/show.qtpl:114
}

//line template/build/show.qtpl:116
func (p *ShowPage) StreamNavigation(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:116
	qw422016.N().S(` `)
	//line template/build/show.qtpl:117
	qw422016.N().S(`<li>`)
	//line template/build/show.qtpl:118
	template.StreamRenderLink(qw422016, p.Build.UIEndpoint(), p.URI)
	//line template/build/show.qtpl:118
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 9c1.641 0 3 1.359 3 3s-1.359 3-3 3-3-1.359-3-3 1.359-3 3-3zM12 17.016c2.766 0 5.016-2.25 5.016-5.016s-2.25-5.016-5.016-5.016-5.016 2.25-5.016 5.016 2.25 5.016 5.016 5.016zM12 4.5c5.016 0 9.281 3.094 11.016 7.5-1.734 4.406-6 7.5-11.016 7.5s-9.281-3.094-11.016-7.5c1.734-4.406 6-7.5 11.016-7.5z"></path>
</svg>
`)
	//line template/build/show.qtpl:118
	qw422016.N().S(`<span>Overview</span></a></li><li>`)
	//line template/build/show.qtpl:119
	template.StreamRenderLink(qw422016, p.Build.UIEndpoint("manifest"), p.URI)
	//line template/build/show.qtpl:119
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M6.984 12.984v-1.969h14.016v1.969h-14.016zM6.984 18.984v-1.969h14.016v1.969h-14.016zM6.984 5.016h14.016v1.969h-14.016v-1.969zM2.016 11.016v-1.031h3v0.938l-1.828 2.063h1.828v1.031h-3v-0.938l1.781-2.063h-1.781zM3 8.016v-3h-0.984v-1.031h1.969v4.031h-0.984zM2.016 17.016v-1.031h3v4.031h-3v-1.031h1.969v-0.469h-0.984v-1.031h0.984v-0.469h-1.969z"></path>
</svg>
`)
	//line template/build/show.qtpl:119
	qw422016.N().S(`<span>Manifest</span></a></li><li>`)
	//line template/build/show.qtpl:120
	template.StreamRenderLink(qw422016, p.Build.UIEndpoint("objects"), p.URI)
	//line template/build/show.qtpl:120
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM9 15.984v-6h-3.984l6.984-6.984 6.984 6.984h-3.984v6h-6z"></path>
</svg>
`)
	//line template/build/show.qtpl:120
	qw422016.N().S(`<span>Objects</span></a></li><li>`)
	//line template/build/show.qtpl:121
	template.StreamRenderLink(qw422016, p.Build.UIEndpoint("artifacts"), p.URI)
	//line template/build/show.qtpl:121
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M5.016 18h13.969v2.016h-13.969v-2.016zM18.984 9l-6.984 6.984-6.984-6.984h3.984v-6h6v6h3.984z"></path>
</svg>
`)
	//line template/build/show.qtpl:121
	qw422016.N().S(`<span>Artifacts</span></a></li><li>`)
	//line template/build/show.qtpl:122
	template.StreamRenderLink(qw422016, p.Build.UIEndpoint("variables"), p.URI)
	//line template/build/show.qtpl:122
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M14.578 16.594l4.641-4.594-4.641-4.594 1.406-1.406 6 6-6 6zM9.422 16.594l-1.406 1.406-6-6 6-6 1.406 1.406-4.641 4.594z"></path>
</svg>
`)
	//line template/build/show.qtpl:122
	qw422016.N().S(`<span>Variables</span></a></li><li>`)
	//line template/build/show.qtpl:123
	template.StreamRenderLink(qw422016, p.Build.UIEndpoint("output"), p.URI)
	//line template/build/show.qtpl:123
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
	//line template/build/show.qtpl:123
	qw422016.N().S(`<span>Output</span></a></li>`)
	//line template/build/show.qtpl:124
	qw422016.N().S(` `)
//line template/build/show.qtpl:125
}

//line template/build/show.qtpl:125
func (p *ShowPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:125
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:125
	p.StreamNavigation(qw422016)
	//line template/build/show.qtpl:125
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:125
}

//line template/build/show.qtpl:125
func (p *ShowPage) Navigation() string {
	//line template/build/show.qtpl:125
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:125
	p.WriteNavigation(qb422016)
	//line template/build/show.qtpl:125
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:125
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:125
	return qs422016
//line template/build/show.qtpl:125
}

//line template/build/show.qtpl:127
func (p *ShowPage) streamrenderManifest(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:127
	qw422016.N().S(` <div class="panel"> <div class="panel-header"> <ul class="panel-actions"> <li><a class="btn btn-primary" href="`)
	//line template/build/show.qtpl:131
	qw422016.E().S(p.Build.UIEndpoint("manifest", "raw"))
	//line template/build/show.qtpl:131
	qw422016.N().S(`">`)
	//line template/build/show.qtpl:131
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
	//line template/build/show.qtpl:131
	qw422016.N().S(`<span>Raw</span></a></li> </ul> </div> `)
	//line template/build/show.qtpl:134
	template.StreamRenderCode(qw422016, p.Build.Manifest)
	//line template/build/show.qtpl:134
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:136
}

//line template/build/show.qtpl:136
func (p *ShowPage) writerenderManifest(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:136
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:136
	p.streamrenderManifest(qw422016)
	//line template/build/show.qtpl:136
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:136
}

//line template/build/show.qtpl:136
func (p *ShowPage) renderManifest() string {
	//line template/build/show.qtpl:136
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:136
	p.writerenderManifest(qb422016)
	//line template/build/show.qtpl:136
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:136
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:136
	return qs422016
//line template/build/show.qtpl:136
}

//line template/build/show.qtpl:138
func (p *ShowPage) streamrenderObjects(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:138
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:140
	if len(p.Build.Objects) > 0 {
		//line template/build/show.qtpl:140
		qw422016.N().S(` <table class="table"> <thead> <tr> <th>SOURCE</th> <th>NAME</th> <th>HASHES</th> <th></th> </tr> </thead> <tbody> `)
		//line template/build/show.qtpl:151
		for _, o := range p.Build.Objects {
			//line template/build/show.qtpl:151
			qw422016.N().S(` <tr> <td> `)
			//line template/build/show.qtpl:154
			if o.Object != nil {
				//line template/build/show.qtpl:154
				qw422016.N().S(` <a href="/objects/`)
				//line template/build/show.qtpl:155
				qw422016.E().V(o.Object.ID)
				//line template/build/show.qtpl:155
				qw422016.N().S(`">`)
				//line template/build/show.qtpl:155
				qw422016.E().S(o.Source)
				//line template/build/show.qtpl:155
				qw422016.N().S(`</a> `)
				//line template/build/show.qtpl:156
			} else {
				//line template/build/show.qtpl:156
				qw422016.N().S(` <a title="Object not found"><strike>`)
				//line template/build/show.qtpl:157
				qw422016.E().S(o.Source)
				//line template/build/show.qtpl:157
				qw422016.N().S(`</strike></a> `)
				//line template/build/show.qtpl:158
			}
			//line template/build/show.qtpl:158
			qw422016.N().S(` </td> <td><code>`)
			//line template/build/show.qtpl:160
			qw422016.E().S(o.Name)
			//line template/build/show.qtpl:160
			qw422016.N().S(`</code></td> <td> <div class="mb-10">MD5 <span class="code right">`)
			//line template/build/show.qtpl:162
			if o.Object != nil {
				//line template/build/show.qtpl:162
				qw422016.E().S(fmt.Sprintf("%x", o.Object.MD5))
				//line template/build/show.qtpl:162
			} else {
				//line template/build/show.qtpl:162
				qw422016.N().S(`--`)
				//line template/build/show.qtpl:162
			}
			//line template/build/show.qtpl:162
			qw422016.N().S(`</span></div> <div>SHA256 <span class="code right">`)
			//line template/build/show.qtpl:163
			if o.Object != nil {
				//line template/build/show.qtpl:163
				qw422016.E().S(fmt.Sprintf("%x", o.Object.SHA256))
				//line template/build/show.qtpl:163
			} else {
				//line template/build/show.qtpl:163
				qw422016.N().S(`--`)
				//line template/build/show.qtpl:163
			}
			//line template/build/show.qtpl:163
			qw422016.N().S(`</span></div> </td> <td class="align-right"> `)
			//line template/build/show.qtpl:166
			if o.Placed {
				//line template/build/show.qtpl:166
				qw422016.N().S(` <span class="pill pill-green">`)
				//line template/build/show.qtpl:167
				qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
				//line template/build/show.qtpl:167
				qw422016.N().S(` Placed</span> `)
				//line template/build/show.qtpl:168
			} else {
				//line template/build/show.qtpl:168
				qw422016.N().S(` <span class="pill pill-red">`)
				//line template/build/show.qtpl:169
				qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
				//line template/build/show.qtpl:169
				qw422016.N().S(` Not Placed</span> `)
				//line template/build/show.qtpl:170
			}
			//line template/build/show.qtpl:170
			qw422016.N().S(` </td> </tr> `)
			//line template/build/show.qtpl:173
		}
		//line template/build/show.qtpl:173
		qw422016.N().S(` </tbody> </table> `)
		//line template/build/show.qtpl:176
	} else {
		//line template/build/show.qtpl:176
		qw422016.N().S(` <div class="panel-message muted">No objects have been placed for this build.</div> `)
		//line template/build/show.qtpl:178
	}
	//line template/build/show.qtpl:178
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:180
}

//line template/build/show.qtpl:180
func (p *ShowPage) writerenderObjects(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:180
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:180
	p.streamrenderObjects(qw422016)
	//line template/build/show.qtpl:180
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:180
}

//line template/build/show.qtpl:180
func (p *ShowPage) renderObjects() string {
	//line template/build/show.qtpl:180
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:180
	p.writerenderObjects(qb422016)
	//line template/build/show.qtpl:180
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:180
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:180
	return qs422016
//line template/build/show.qtpl:180
}

//line template/build/show.qtpl:182
func (p *ShowPage) streamrenderArtifacts(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:182
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:184
	if len(p.Build.Artifacts) > 0 {
		//line template/build/show.qtpl:184
		qw422016.N().S(` `)
		//line template/build/show.qtpl:185
		artifact.StreamRenderTable(qw422016, p.Build.Artifacts)
		//line template/build/show.qtpl:185
		qw422016.N().S(` `)
		//line template/build/show.qtpl:186
	} else {
		//line template/build/show.qtpl:186
		qw422016.N().S(` <div class="panel-message muted">No artifacts have been collected from this build.</div> `)
		//line template/build/show.qtpl:188
	}
	//line template/build/show.qtpl:188
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:190
}

//line template/build/show.qtpl:190
func (p *ShowPage) writerenderArtifacts(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:190
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:190
	p.streamrenderArtifacts(qw422016)
	//line template/build/show.qtpl:190
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:190
}

//line template/build/show.qtpl:190
func (p *ShowPage) renderArtifacts() string {
	//line template/build/show.qtpl:190
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:190
	p.writerenderArtifacts(qb422016)
	//line template/build/show.qtpl:190
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:190
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:190
	return qs422016
//line template/build/show.qtpl:190
}

//line template/build/show.qtpl:192
func (p *ShowPage) streamrenderVariables(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:192
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:194
	if len(p.Build.Variables) > 0 {
		//line template/build/show.qtpl:194
		qw422016.N().S(` `)
		//line template/build/show.qtpl:195
		StreamRenderVariablesTable(qw422016, p.Build.Variables)
		//line template/build/show.qtpl:195
		qw422016.N().S(` `)
		//line template/build/show.qtpl:196
	} else {
		//line template/build/show.qtpl:196
		qw422016.N().S(` <div class="panel-message muted">No variables have been set for this build.</div> `)
		//line template/build/show.qtpl:198
	}
	//line template/build/show.qtpl:198
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:200
}

//line template/build/show.qtpl:200
func (p *ShowPage) writerenderVariables(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:200
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:200
	p.streamrenderVariables(qw422016)
	//line template/build/show.qtpl:200
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:200
}

//line template/build/show.qtpl:200
func (p *ShowPage) renderVariables() string {
	//line template/build/show.qtpl:200
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:200
	p.writerenderVariables(qb422016)
	//line template/build/show.qtpl:200
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:200
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:200
	return qs422016
//line template/build/show.qtpl:200
}

//line template/build/show.qtpl:202
func (p *ShowPage) streamrenderOutput(qw422016 *qt422016.Writer) {
	//line template/build/show.qtpl:202
	qw422016.N().S(` <div class="panel"> `)
	//line template/build/show.qtpl:204
	if p.Build.Output.Valid {
		//line template/build/show.qtpl:204
		qw422016.N().S(` <div class="panel-header"> <ul class="panel-actions"> <li><a class="btn btn-primary" href="`)
		//line template/build/show.qtpl:207
		qw422016.E().S(p.Build.UIEndpoint("output", "raw"))
		//line template/build/show.qtpl:207
		qw422016.N().S(`">`)
		//line template/build/show.qtpl:207
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
		//line template/build/show.qtpl:207
		qw422016.N().S(`<span>Raw</span></a></li> </ul> </div> `)
		//line template/build/show.qtpl:210
		template.StreamRenderCode(qw422016, p.Build.Output.String)
		//line template/build/show.qtpl:210
		qw422016.N().S(` `)
		//line template/build/show.qtpl:211
	} else {
		//line template/build/show.qtpl:211
		qw422016.N().S(` <div class="panel-message muted">No build output has been produced.</div> `)
		//line template/build/show.qtpl:213
	}
	//line template/build/show.qtpl:213
	qw422016.N().S(` </div> `)
//line template/build/show.qtpl:215
}

//line template/build/show.qtpl:215
func (p *ShowPage) writerenderOutput(qq422016 qtio422016.Writer) {
	//line template/build/show.qtpl:215
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/show.qtpl:215
	p.streamrenderOutput(qw422016)
	//line template/build/show.qtpl:215
	qt422016.ReleaseWriter(qw422016)
//line template/build/show.qtpl:215
}

//line template/build/show.qtpl:215
func (p *ShowPage) renderOutput() string {
	//line template/build/show.qtpl:215
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/show.qtpl:215
	p.writerenderOutput(qb422016)
	//line template/build/show.qtpl:215
	qs422016 := string(qb422016.B)
	//line template/build/show.qtpl:215
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/show.qtpl:215
	return qs422016
//line template/build/show.qtpl:215
}

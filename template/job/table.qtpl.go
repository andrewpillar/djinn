// This file is automatically generated by qtc from "table.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/job/table.qtpl:2
package job

//line template/job/table.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/job/table.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/job/table.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/job/table.qtpl:9
func StreamRenderTable(qw422016 *qt422016.Writer, jj []*model.Job) {
	//line template/job/table.qtpl:9
	qw422016.N().S(` <table class="table"> <thead> <tr> <th class="cell-pill">STATUS</th> <th>JOB</th> <th class="cell-date">DURATION</th> </tr> </thead> <tbody> `)
	//line template/job/table.qtpl:19
	for _, j := range jj {
		//line template/job/table.qtpl:19
		qw422016.N().S(` <tr> <td class="cell-pill">`)
		//line template/job/table.qtpl:21
		template.StreamRenderStatus(qw422016, j.Status)
		//line template/job/table.qtpl:21
		qw422016.N().S(`</td> <td><a href="`)
		//line template/job/table.qtpl:22
		qw422016.E().S(j.UIEndpoint())
		//line template/job/table.qtpl:22
		qw422016.N().S(`">`)
		//line template/job/table.qtpl:22
		qw422016.E().S(j.Name)
		//line template/job/table.qtpl:22
		qw422016.N().S(`</a></td> <td class="cell-date"> `)
		//line template/job/table.qtpl:24
		if !j.StartedAt.Valid || !j.FinishedAt.Valid {
			//line template/job/table.qtpl:24
			qw422016.N().S(` <span class="muted">--</span> `)
			//line template/job/table.qtpl:26
		} else {
			//line template/job/table.qtpl:26
			qw422016.N().S(` `)
			//line template/job/table.qtpl:27
			qw422016.E().V(j.FinishedAt.Time.Sub(j.StartedAt.Time))
			//line template/job/table.qtpl:27
			qw422016.N().S(` `)
			//line template/job/table.qtpl:28
		}
		//line template/job/table.qtpl:28
		qw422016.N().S(` </td> </tr> `)
		//line template/job/table.qtpl:31
	}
	//line template/job/table.qtpl:31
	qw422016.N().S(` </tbody> </table> `)
//line template/job/table.qtpl:34
}

//line template/job/table.qtpl:34
func WriteRenderTable(qq422016 qtio422016.Writer, jj []*model.Job) {
	//line template/job/table.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/job/table.qtpl:34
	StreamRenderTable(qw422016, jj)
	//line template/job/table.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line template/job/table.qtpl:34
}

//line template/job/table.qtpl:34
func RenderTable(jj []*model.Job) string {
	//line template/job/table.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/job/table.qtpl:34
	WriteRenderTable(qb422016, jj)
	//line template/job/table.qtpl:34
	qs422016 := string(qb422016.B)
	//line template/job/table.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/job/table.qtpl:34
	return qs422016
//line template/job/table.qtpl:34
}

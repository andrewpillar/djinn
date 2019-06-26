// This file is automatically generated by qtc from "table.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/artifact/table.qtpl:2
package artifact

//line template/artifact/table.qtpl:2
import (
	"fmt"

	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/artifact/table.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/artifact/table.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/artifact/table.qtpl:11
func StreamRenderTable(qw422016 *qt422016.Writer, aa []*model.Artifact) {
	//line template/artifact/table.qtpl:11
	qw422016.N().S(` <table class="table"> <thead> <tr> <th>SOURCE</th> <th>NAME</th> <tH>SIZE</th> <th>MD5</th> <th>SHA256</th> </tr> </thead> <tbody> `)
	//line template/artifact/table.qtpl:23
	for _, a := range aa {
		//line template/artifact/table.qtpl:23
		qw422016.N().S(` <tr> <td><span class="code">`)
		//line template/artifact/table.qtpl:25
		qw422016.E().S(a.Source)
		//line template/artifact/table.qtpl:25
		qw422016.N().S(`</span></td> <td><a href="`)
		//line template/artifact/table.qtpl:26
		qw422016.E().S(a.UIEndpoint("download", a.Name))
		//line template/artifact/table.qtpl:26
		qw422016.N().S(`">`)
		//line template/artifact/table.qtpl:26
		qw422016.E().S(a.Name)
		//line template/artifact/table.qtpl:26
		qw422016.N().S(`</a></td> <td>`)
		//line template/artifact/table.qtpl:27
		qw422016.E().S(template.RenderSize(a.Size.Int64))
		//line template/artifact/table.qtpl:27
		qw422016.N().S(`</td> <td> <span class="code"> `)
		//line template/artifact/table.qtpl:30
		if len(a.MD5) == 0 {
			//line template/artifact/table.qtpl:30
			qw422016.N().S(` -- `)
			//line template/artifact/table.qtpl:32
		} else {
			//line template/artifact/table.qtpl:32
			qw422016.N().S(` `)
			//line template/artifact/table.qtpl:33
			qw422016.E().S(fmt.Sprintf("%x", a.MD5))
			//line template/artifact/table.qtpl:33
			qw422016.N().S(` `)
			//line template/artifact/table.qtpl:34
		}
		//line template/artifact/table.qtpl:34
		qw422016.N().S(` </span> </td> <td> <span class="code"> `)
		//line template/artifact/table.qtpl:39
		if len(a.SHA256) == 0 {
			//line template/artifact/table.qtpl:39
			qw422016.N().S(` -- `)
			//line template/artifact/table.qtpl:41
		} else {
			//line template/artifact/table.qtpl:41
			qw422016.N().S(` `)
			//line template/artifact/table.qtpl:42
			qw422016.E().S(fmt.Sprintf("%x", a.SHA256))
			//line template/artifact/table.qtpl:42
			qw422016.N().S(` `)
			//line template/artifact/table.qtpl:43
		}
		//line template/artifact/table.qtpl:43
		qw422016.N().S(` </span> </td> </tr> `)
		//line template/artifact/table.qtpl:47
	}
	//line template/artifact/table.qtpl:47
	qw422016.N().S(` </tbody> </table> `)
//line template/artifact/table.qtpl:50
}

//line template/artifact/table.qtpl:50
func WriteRenderTable(qq422016 qtio422016.Writer, aa []*model.Artifact) {
	//line template/artifact/table.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/artifact/table.qtpl:50
	StreamRenderTable(qw422016, aa)
	//line template/artifact/table.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line template/artifact/table.qtpl:50
}

//line template/artifact/table.qtpl:50
func RenderTable(aa []*model.Artifact) string {
	//line template/artifact/table.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/artifact/table.qtpl:50
	WriteRenderTable(qb422016, aa)
	//line template/artifact/table.qtpl:50
	qs422016 := string(qb422016.B)
	//line template/artifact/table.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/artifact/table.qtpl:50
	return qs422016
//line template/artifact/table.qtpl:50
}
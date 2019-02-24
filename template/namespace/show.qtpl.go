// This file is automatically generated by qtc from "show.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/namespace/show.qtpl:2
package namespace

//line template/namespace/show.qtpl:2
import (
	"strings"

	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/namespace/show.qtpl:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/namespace/show.qtpl:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/namespace/show.qtpl:11
type ShowPage struct {
	*template.Page

	User      *model.User
	Namespace *model.Namespace
}

//line template/namespace/show.qtpl:20
func streamrenderFullName(qw422016 *qt422016.Writer, username, fullName string) {
	//line template/namespace/show.qtpl:22
	parts := strings.Split(fullName, "/")

	//line template/namespace/show.qtpl:24
	for i, p := range parts {
		//line template/namespace/show.qtpl:24
		qw422016.N().S(`<a href="/u/`)
		//line template/namespace/show.qtpl:25
		qw422016.E().S(username)
		//line template/namespace/show.qtpl:25
		qw422016.N().S(`/`)
		//line template/namespace/show.qtpl:25
		qw422016.E().S(strings.Join(parts[:i+1], "/"))
		//line template/namespace/show.qtpl:25
		qw422016.N().S(`">`)
		//line template/namespace/show.qtpl:25
		qw422016.E().S(p)
		//line template/namespace/show.qtpl:25
		qw422016.N().S(`</a>`)
		//line template/namespace/show.qtpl:26
		if i != len(parts)-1 {
			//line template/namespace/show.qtpl:26
			qw422016.N().S(`<span> / </span>`)
			//line template/namespace/show.qtpl:26
		}
		//line template/namespace/show.qtpl:27
	}
//line template/namespace/show.qtpl:28
}

//line template/namespace/show.qtpl:28
func writerenderFullName(qq422016 qtio422016.Writer, username, fullName string) {
	//line template/namespace/show.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/show.qtpl:28
	streamrenderFullName(qw422016, username, fullName)
	//line template/namespace/show.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/show.qtpl:28
}

//line template/namespace/show.qtpl:28
func renderFullName(username, fullName string) string {
	//line template/namespace/show.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/show.qtpl:28
	writerenderFullName(qb422016, username, fullName)
	//line template/namespace/show.qtpl:28
	qs422016 := string(qb422016.B)
	//line template/namespace/show.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/show.qtpl:28
	return qs422016
//line template/namespace/show.qtpl:28
}

//line template/namespace/show.qtpl:31
func (p *ShowPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/namespace/show.qtpl:31
	qw422016.N().S(`
`)
	//line template/namespace/show.qtpl:32
	qw422016.E().S(p.Namespace.FullName)
	//line template/namespace/show.qtpl:32
	qw422016.N().S(` - Thrall
`)
//line template/namespace/show.qtpl:33
}

//line template/namespace/show.qtpl:33
func (p *ShowPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/namespace/show.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/show.qtpl:33
	p.StreamTitle(qw422016)
	//line template/namespace/show.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/show.qtpl:33
}

//line template/namespace/show.qtpl:33
func (p *ShowPage) Title() string {
	//line template/namespace/show.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/show.qtpl:33
	p.WriteTitle(qb422016)
	//line template/namespace/show.qtpl:33
	qs422016 := string(qb422016.B)
	//line template/namespace/show.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/show.qtpl:33
	return qs422016
//line template/namespace/show.qtpl:33
}

//line template/namespace/show.qtpl:35
func (p *ShowPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/namespace/show.qtpl:35
	qw422016.N().S(`
<div class="dashboard-header">
	<h1>`)
	//line template/namespace/show.qtpl:37
	streamrenderFullName(qw422016, p.Namespace.User.Username, p.Namespace.FullName)
	//line template/namespace/show.qtpl:37
	qw422016.N().S(`</h1>
	<ul class="actions">
		<li><a href="/u/`)
	//line template/namespace/show.qtpl:39
	qw422016.E().S(p.User.Username)
	//line template/namespace/show.qtpl:39
	qw422016.N().S(`/`)
	//line template/namespace/show.qtpl:39
	qw422016.E().S(p.Namespace.FullName)
	//line template/namespace/show.qtpl:39
	qw422016.N().S(`/-/edit" class="button button-secondary">Edit</a></li>
		<li><a href="/namespaces/create?parent=`)
	//line template/namespace/show.qtpl:40
	qw422016.E().S(p.Namespace.FullName)
	//line template/namespace/show.qtpl:40
	qw422016.N().S(`" class="button button-primary">Create</a></li>
	</ul>
</div>
`)
//line template/namespace/show.qtpl:43
}

//line template/namespace/show.qtpl:43
func (p *ShowPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/namespace/show.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/show.qtpl:43
	p.StreamBody(qw422016)
	//line template/namespace/show.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/show.qtpl:43
}

//line template/namespace/show.qtpl:43
func (p *ShowPage) Body() string {
	//line template/namespace/show.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/show.qtpl:43
	p.WriteBody(qb422016)
	//line template/namespace/show.qtpl:43
	qs422016 := string(qb422016.B)
	//line template/namespace/show.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/show.qtpl:43
	return qs422016
//line template/namespace/show.qtpl:43
}

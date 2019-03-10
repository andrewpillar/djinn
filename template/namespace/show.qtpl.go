// This file is automatically generated by qtc from "show.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/namespace/show.qtpl:2
package namespace

//line template/namespace/show.qtpl:2
import (
	"strings"

	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/build"
)

//line template/namespace/show.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/namespace/show.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/namespace/show.qtpl:12
type ShowPage struct {
	*template.Page

	Namespace *model.Namespace
	Builds    []*model.Build
}

//line template/namespace/show.qtpl:21
func streamrenderFullName(qw422016 *qt422016.Writer, username, fullName string) {
	//line template/namespace/show.qtpl:21
	qw422016.N().S(` `)
	//line template/namespace/show.qtpl:23
	parts := strings.Split(fullName, "/")

	//line template/namespace/show.qtpl:24
	qw422016.N().S(` `)
	//line template/namespace/show.qtpl:25
	for i, p := range parts {
		//line template/namespace/show.qtpl:25
		qw422016.N().S(` <a href="/u/`)
		//line template/namespace/show.qtpl:26
		qw422016.E().S(username)
		//line template/namespace/show.qtpl:26
		qw422016.N().S(`/`)
		//line template/namespace/show.qtpl:26
		qw422016.E().S(strings.Join(parts[:i+1], "/"))
		//line template/namespace/show.qtpl:26
		qw422016.N().S(`">`)
		//line template/namespace/show.qtpl:26
		qw422016.E().S(p)
		//line template/namespace/show.qtpl:26
		qw422016.N().S(`</a> `)
		//line template/namespace/show.qtpl:27
		if i != len(parts)-1 {
			//line template/namespace/show.qtpl:27
			qw422016.N().S(`<span> / </span>`)
			//line template/namespace/show.qtpl:27
		}
		//line template/namespace/show.qtpl:27
		qw422016.N().S(` `)
		//line template/namespace/show.qtpl:28
	}
	//line template/namespace/show.qtpl:28
	qw422016.N().S(` `)
//line template/namespace/show.qtpl:29
}

//line template/namespace/show.qtpl:29
func writerenderFullName(qq422016 qtio422016.Writer, username, fullName string) {
	//line template/namespace/show.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/show.qtpl:29
	streamrenderFullName(qw422016, username, fullName)
	//line template/namespace/show.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/show.qtpl:29
}

//line template/namespace/show.qtpl:29
func renderFullName(username, fullName string) string {
	//line template/namespace/show.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/show.qtpl:29
	writerenderFullName(qb422016, username, fullName)
	//line template/namespace/show.qtpl:29
	qs422016 := string(qb422016.B)
	//line template/namespace/show.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/show.qtpl:29
	return qs422016
//line template/namespace/show.qtpl:29
}

//line template/namespace/show.qtpl:33
func (p *ShowPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/namespace/show.qtpl:33
	qw422016.N().S(` `)
	//line template/namespace/show.qtpl:34
	qw422016.E().S(p.Namespace.FullName)
	//line template/namespace/show.qtpl:34
	qw422016.N().S(` - Thrall `)
//line template/namespace/show.qtpl:35
}

//line template/namespace/show.qtpl:35
func (p *ShowPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/namespace/show.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/show.qtpl:35
	p.StreamTitle(qw422016)
	//line template/namespace/show.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/show.qtpl:35
}

//line template/namespace/show.qtpl:35
func (p *ShowPage) Title() string {
	//line template/namespace/show.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/show.qtpl:35
	p.WriteTitle(qb422016)
	//line template/namespace/show.qtpl:35
	qs422016 := string(qb422016.B)
	//line template/namespace/show.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/show.qtpl:35
	return qs422016
//line template/namespace/show.qtpl:35
}

//line template/namespace/show.qtpl:37
func (p *ShowPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/namespace/show.qtpl:37
	qw422016.N().S(` <div class="header"> <h1> `)
	//line template/namespace/show.qtpl:40
	if p.Namespace.Parent != nil {
		//line template/namespace/show.qtpl:40
		qw422016.N().S(` <a class="back" href="`)
		//line template/namespace/show.qtpl:41
		qw422016.E().S(p.Namespace.Parent.URI())
		//line template/namespace/show.qtpl:41
		qw422016.N().S(`">`)
		//line template/namespace/show.qtpl:41
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
		//line template/namespace/show.qtpl:41
		qw422016.N().S(`</a> `)
		//line template/namespace/show.qtpl:42
	} else {
		//line template/namespace/show.qtpl:42
		qw422016.N().S(` <a class="back" href="/namespaces">`)
		//line template/namespace/show.qtpl:43
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
		//line template/namespace/show.qtpl:43
		qw422016.N().S(`</a> `)
		//line template/namespace/show.qtpl:44
	}
	//line template/namespace/show.qtpl:44
	qw422016.N().S(` `)
	//line template/namespace/show.qtpl:45
	streamrenderFullName(qw422016, p.Namespace.User.Username, p.Namespace.FullName)
	//line template/namespace/show.qtpl:45
	qw422016.N().S(`<br/> <small>`)
	//line template/namespace/show.qtpl:46
	qw422016.E().S(p.Namespace.Description)
	//line template/namespace/show.qtpl:46
	qw422016.N().S(`</small> </h1> <ul class="actions"> <li><a href="`)
	//line template/namespace/show.qtpl:49
	qw422016.E().S(p.Namespace.URI())
	//line template/namespace/show.qtpl:49
	qw422016.N().S(`/-/edit" class="button button-secondary">Edit</a></li> <li><a href="/namespaces/create?parent=`)
	//line template/namespace/show.qtpl:50
	qw422016.E().S(p.Namespace.FullName)
	//line template/namespace/show.qtpl:50
	qw422016.N().S(`" class="button button-primary">Create</a></li> </ul> </div> <ul class="tabs"> <li><a href="`)
	//line template/namespace/show.qtpl:54
	qw422016.E().S(p.Namespace.URI())
	//line template/namespace/show.qtpl:54
	qw422016.N().S(`" `)
	//line template/namespace/show.qtpl:54
	if p.Namespace.URI() == p.URI {
		//line template/namespace/show.qtpl:54
		qw422016.N().S(`class="active" `)
		//line template/namespace/show.qtpl:54
	}
	//line template/namespace/show.qtpl:54
	qw422016.N().S(`>Builds</a></li> <li><a href="`)
	//line template/namespace/show.qtpl:55
	qw422016.E().S(p.Namespace.URI())
	//line template/namespace/show.qtpl:55
	qw422016.N().S(`/-/namespaces">Namespaces</a></li> </ul> <div class="body">`)
	//line template/namespace/show.qtpl:57
	build.StreamRenderBuilds(qw422016, p.Builds)
	//line template/namespace/show.qtpl:57
	qw422016.N().S(`</div> `)
//line template/namespace/show.qtpl:58
}

//line template/namespace/show.qtpl:58
func (p *ShowPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/namespace/show.qtpl:58
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/show.qtpl:58
	p.StreamBody(qw422016)
	//line template/namespace/show.qtpl:58
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/show.qtpl:58
}

//line template/namespace/show.qtpl:58
func (p *ShowPage) Body() string {
	//line template/namespace/show.qtpl:58
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/show.qtpl:58
	p.WriteBody(qb422016)
	//line template/namespace/show.qtpl:58
	qs422016 := string(qb422016.B)
	//line template/namespace/show.qtpl:58
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/show.qtpl:58
	return qs422016
//line template/namespace/show.qtpl:58
}

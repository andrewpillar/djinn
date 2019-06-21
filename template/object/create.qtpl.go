// This file is automatically generated by qtc from "create.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/object/create.qtpl:2
package object

//line template/object/create.qtpl:2
import (
	"github.com/andrewpillar/thrall/template"
)

//line template/object/create.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/object/create.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/object/create.qtpl:8
type CreatePage struct {
	template.Page
	template.Form
}

//line template/object/create.qtpl:15
func (p *CreatePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/object/create.qtpl:15
	qw422016.N().S(` Create Object - Thrall `)
//line template/object/create.qtpl:17
}

//line template/object/create.qtpl:17
func (p *CreatePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/object/create.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/object/create.qtpl:17
	p.StreamTitle(qw422016)
	//line template/object/create.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line template/object/create.qtpl:17
}

//line template/object/create.qtpl:17
func (p *CreatePage) Title() string {
	//line template/object/create.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/object/create.qtpl:17
	p.WriteTitle(qb422016)
	//line template/object/create.qtpl:17
	qs422016 := string(qb422016.B)
	//line template/object/create.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/object/create.qtpl:17
	return qs422016
//line template/object/create.qtpl:17
}

//line template/object/create.qtpl:19
func (p *CreatePage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/object/create.qtpl:19
	qw422016.N().S(` <div class="panel"> <form class="panel-body slim" method="POST" action="/objects" enctype="multipart/form-data"> `)
	//line template/object/create.qtpl:22
	qw422016.N().S(string(p.CSRF))
	//line template/object/create.qtpl:22
	qw422016.N().S(` `)
	//line template/object/create.qtpl:23
	if p.Errors.First("object") != "" {
		//line template/object/create.qtpl:23
		qw422016.N().S(` `)
		//line template/object/create.qtpl:24
		p.StreamError(qw422016, "object")
		//line template/object/create.qtpl:24
		qw422016.N().S(` `)
		//line template/object/create.qtpl:25
	}
	//line template/object/create.qtpl:25
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input class="form-text" type="text" id="name" name="name" value="`)
	//line template/object/create.qtpl:28
	qw422016.E().S(p.Form.Get("name"))
	//line template/object/create.qtpl:28
	qw422016.N().S(`" autocomplete="off"/> `)
	//line template/object/create.qtpl:29
	p.StreamError(qw422016, "name")
	//line template/object/create.qtpl:29
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="file">File</label> <input type="file" id="file" name="file"/> `)
	//line template/object/create.qtpl:34
	p.StreamError(qw422016, "file")
	//line template/object/create.qtpl:34
	qw422016.N().S(` </div> <div class="form-field"> <button type="submit" class="btn btn-primary">Create</button> </div> </form> </div> `)
//line template/object/create.qtpl:41
}

//line template/object/create.qtpl:41
func (p *CreatePage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/object/create.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/object/create.qtpl:41
	p.StreamBody(qw422016)
	//line template/object/create.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line template/object/create.qtpl:41
}

//line template/object/create.qtpl:41
func (p *CreatePage) Body() string {
	//line template/object/create.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/object/create.qtpl:41
	p.WriteBody(qb422016)
	//line template/object/create.qtpl:41
	qs422016 := string(qb422016.B)
	//line template/object/create.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/object/create.qtpl:41
	return qs422016
//line template/object/create.qtpl:41
}

//line template/object/create.qtpl:43
func (p *CreatePage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/object/create.qtpl:43
	qw422016.N().S(` <a href="/objects" class="back">`)
	//line template/object/create.qtpl:44
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
	//line template/object/create.qtpl:44
	qw422016.N().S(`</a> Create Object `)
//line template/object/create.qtpl:45
}

//line template/object/create.qtpl:45
func (p *CreatePage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/object/create.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/object/create.qtpl:45
	p.StreamHeader(qw422016)
	//line template/object/create.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line template/object/create.qtpl:45
}

//line template/object/create.qtpl:45
func (p *CreatePage) Header() string {
	//line template/object/create.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/object/create.qtpl:45
	p.WriteHeader(qb422016)
	//line template/object/create.qtpl:45
	qs422016 := string(qb422016.B)
	//line template/object/create.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/object/create.qtpl:45
	return qs422016
//line template/object/create.qtpl:45
}

//line template/object/create.qtpl:47
func (p *CreatePage) StreamActions(qw422016 *qt422016.Writer) {
//line template/object/create.qtpl:47
}

//line template/object/create.qtpl:47
func (p *CreatePage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/object/create.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/object/create.qtpl:47
	p.StreamActions(qw422016)
	//line template/object/create.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line template/object/create.qtpl:47
}

//line template/object/create.qtpl:47
func (p *CreatePage) Actions() string {
	//line template/object/create.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/object/create.qtpl:47
	p.WriteActions(qb422016)
	//line template/object/create.qtpl:47
	qs422016 := string(qb422016.B)
	//line template/object/create.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/object/create.qtpl:47
	return qs422016
//line template/object/create.qtpl:47
}

//line template/object/create.qtpl:48
func (p *CreatePage) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/object/create.qtpl:48
}

//line template/object/create.qtpl:48
func (p *CreatePage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/object/create.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/object/create.qtpl:48
	p.StreamNavigation(qw422016)
	//line template/object/create.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line template/object/create.qtpl:48
}

//line template/object/create.qtpl:48
func (p *CreatePage) Navigation() string {
	//line template/object/create.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/object/create.qtpl:48
	p.WriteNavigation(qb422016)
	//line template/object/create.qtpl:48
	qs422016 := string(qb422016.B)
	//line template/object/create.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/object/create.qtpl:48
	return qs422016
//line template/object/create.qtpl:48
}

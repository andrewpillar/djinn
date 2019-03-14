// This file is automatically generated by qtc from "create.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/create.qtpl:2
package build

//line template/build/create.qtpl:2
import (
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/template"
)

//line template/build/create.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/create.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/create.qtpl:9
type CreatePage struct {
	*template.Page

	Errors form.Errors
	Form   form.Form
}

//line template/build/create.qtpl:18
func (p *CreatePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/create.qtpl:18
	qw422016.N().S(` Submit Build - Thrall `)
//line template/build/create.qtpl:20
}

//line template/build/create.qtpl:20
func (p *CreatePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/create.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/create.qtpl:20
	p.StreamTitle(qw422016)
	//line template/build/create.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line template/build/create.qtpl:20
}

//line template/build/create.qtpl:20
func (p *CreatePage) Title() string {
	//line template/build/create.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/create.qtpl:20
	p.WriteTitle(qb422016)
	//line template/build/create.qtpl:20
	qs422016 := string(qb422016.B)
	//line template/build/create.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/create.qtpl:20
	return qs422016
//line template/build/create.qtpl:20
}

//line template/build/create.qtpl:22
func (p *CreatePage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/create.qtpl:22
	qw422016.N().S(` <form class="slim" method="POST" action="/builds"> `)
	//line template/build/create.qtpl:24
	if p.Errors.First("build") != "" {
		//line template/build/create.qtpl:24
		qw422016.N().S(` <div class="form-error">Failed to submit build: `)
		//line template/build/create.qtpl:25
		qw422016.E().S(p.Errors.First("build"))
		//line template/build/create.qtpl:25
		qw422016.N().S(`</div> `)
		//line template/build/create.qtpl:26
	}
	//line template/build/create.qtpl:26
	qw422016.N().S(` <div class="form-field"> <label class="label">Namespace <small>(optional)</small></label> <input class="text" type="text" name="namespace" autocomplete="off"/> </div> <div class="form-field"> <label class="label">Manifest</label> <textarea style="font-family: monospace;min-height: 250px;" class="text" name="manifest"></textarea> <div class="error">`)
	//line template/build/create.qtpl:34
	qw422016.E().S(p.Errors.First("manifest"))
	//line template/build/create.qtpl:34
	qw422016.N().S(`</div> </div> <div class="form-field"> <label class="label">Tags <small>(optional)</small></label> <input class="text" type="text" name="tags" autocomplete="off"/> </div> <div class="form-field"> <button type="submit" class="button button-primary">Submit</button> </div> </form> `)
//line template/build/create.qtpl:44
}

//line template/build/create.qtpl:44
func (p *CreatePage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/create.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/create.qtpl:44
	p.StreamBody(qw422016)
	//line template/build/create.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line template/build/create.qtpl:44
}

//line template/build/create.qtpl:44
func (p *CreatePage) Body() string {
	//line template/build/create.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/create.qtpl:44
	p.WriteBody(qb422016)
	//line template/build/create.qtpl:44
	qs422016 := string(qb422016.B)
	//line template/build/create.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/create.qtpl:44
	return qs422016
//line template/build/create.qtpl:44
}

//line template/build/create.qtpl:46
func (p *CreatePage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/build/create.qtpl:46
	qw422016.N().S(` <a href="/" class="back">`)
	//line template/build/create.qtpl:47
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
	//line template/build/create.qtpl:47
	qw422016.N().S(`</a> Submit Build `)
//line template/build/create.qtpl:48
}

//line template/build/create.qtpl:48
func (p *CreatePage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/build/create.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/create.qtpl:48
	p.StreamHeader(qw422016)
	//line template/build/create.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line template/build/create.qtpl:48
}

//line template/build/create.qtpl:48
func (p *CreatePage) Header() string {
	//line template/build/create.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/create.qtpl:48
	p.WriteHeader(qb422016)
	//line template/build/create.qtpl:48
	qs422016 := string(qb422016.B)
	//line template/build/create.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/create.qtpl:48
	return qs422016
//line template/build/create.qtpl:48
}

//line template/build/create.qtpl:50
func (p *CreatePage) StreamActions(qw422016 *qt422016.Writer) {
//line template/build/create.qtpl:50
}

//line template/build/create.qtpl:50
func (p *CreatePage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/build/create.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/create.qtpl:50
	p.StreamActions(qw422016)
	//line template/build/create.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line template/build/create.qtpl:50
}

//line template/build/create.qtpl:50
func (p *CreatePage) Actions() string {
	//line template/build/create.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/create.qtpl:50
	p.WriteActions(qb422016)
	//line template/build/create.qtpl:50
	qs422016 := string(qb422016.B)
	//line template/build/create.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/create.qtpl:50
	return qs422016
//line template/build/create.qtpl:50
}

//line template/build/create.qtpl:51
func (p *CreatePage) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/build/create.qtpl:51
}

//line template/build/create.qtpl:51
func (p *CreatePage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/build/create.qtpl:51
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/create.qtpl:51
	p.StreamNavigation(qw422016)
	//line template/build/create.qtpl:51
	qt422016.ReleaseWriter(qw422016)
//line template/build/create.qtpl:51
}

//line template/build/create.qtpl:51
func (p *CreatePage) Navigation() string {
	//line template/build/create.qtpl:51
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/create.qtpl:51
	p.WriteNavigation(qb422016)
	//line template/build/create.qtpl:51
	qs422016 := string(qb422016.B)
	//line template/build/create.qtpl:51
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/create.qtpl:51
	return qs422016
//line template/build/create.qtpl:51
}

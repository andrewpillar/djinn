// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/index.qtpl:2
package build

//line template/build/index.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/build/index.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/index.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/index.qtpl:9
type IndexPage struct {
	*template.Page

	Builds []*model.Build
}

//line template/build/index.qtpl:17
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/build/index.qtpl:17
	qw422016.N().S(` Builds - Thrall `)
//line template/build/index.qtpl:19
}

//line template/build/index.qtpl:19
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/build/index.qtpl:19
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/index.qtpl:19
	p.StreamTitle(qw422016)
	//line template/build/index.qtpl:19
	qt422016.ReleaseWriter(qw422016)
//line template/build/index.qtpl:19
}

//line template/build/index.qtpl:19
func (p *IndexPage) Title() string {
	//line template/build/index.qtpl:19
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/index.qtpl:19
	p.WriteTitle(qb422016)
	//line template/build/index.qtpl:19
	qs422016 := string(qb422016.B)
	//line template/build/index.qtpl:19
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/index.qtpl:19
	return qs422016
//line template/build/index.qtpl:19
}

//line template/build/index.qtpl:21
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/build/index.qtpl:21
	qw422016.N().S(` `)
	//line template/build/index.qtpl:22
	StreamRenderBuilds(qw422016, p.Builds)
	//line template/build/index.qtpl:22
	qw422016.N().S(` `)
//line template/build/index.qtpl:23
}

//line template/build/index.qtpl:23
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/build/index.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/index.qtpl:23
	p.StreamBody(qw422016)
	//line template/build/index.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line template/build/index.qtpl:23
}

//line template/build/index.qtpl:23
func (p *IndexPage) Body() string {
	//line template/build/index.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/index.qtpl:23
	p.WriteBody(qb422016)
	//line template/build/index.qtpl:23
	qs422016 := string(qb422016.B)
	//line template/build/index.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/index.qtpl:23
	return qs422016
//line template/build/index.qtpl:23
}

//line template/build/index.qtpl:25
func (p *IndexPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/build/index.qtpl:25
	qw422016.N().S(` Builds `)
//line template/build/index.qtpl:27
}

//line template/build/index.qtpl:27
func (p *IndexPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/build/index.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/index.qtpl:27
	p.StreamHeader(qw422016)
	//line template/build/index.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line template/build/index.qtpl:27
}

//line template/build/index.qtpl:27
func (p *IndexPage) Header() string {
	//line template/build/index.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/index.qtpl:27
	p.WriteHeader(qb422016)
	//line template/build/index.qtpl:27
	qs422016 := string(qb422016.B)
	//line template/build/index.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/index.qtpl:27
	return qs422016
//line template/build/index.qtpl:27
}

//line template/build/index.qtpl:29
func (p *IndexPage) StreamActions(qw422016 *qt422016.Writer) {
	//line template/build/index.qtpl:29
	qw422016.N().S(` <li><a href="/builds/create" class="button button-primary">Submit</a></li> `)
//line template/build/index.qtpl:31
}

//line template/build/index.qtpl:31
func (p *IndexPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/build/index.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/index.qtpl:31
	p.StreamActions(qw422016)
	//line template/build/index.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line template/build/index.qtpl:31
}

//line template/build/index.qtpl:31
func (p *IndexPage) Actions() string {
	//line template/build/index.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/index.qtpl:31
	p.WriteActions(qb422016)
	//line template/build/index.qtpl:31
	qs422016 := string(qb422016.B)
	//line template/build/index.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/index.qtpl:31
	return qs422016
//line template/build/index.qtpl:31
}

//line template/build/index.qtpl:33
func (p *IndexPage) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/build/index.qtpl:33
}

//line template/build/index.qtpl:33
func (p *IndexPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/build/index.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/index.qtpl:33
	p.StreamNavigation(qw422016)
	//line template/build/index.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line template/build/index.qtpl:33
}

//line template/build/index.qtpl:33
func (p *IndexPage) Navigation() string {
	//line template/build/index.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/index.qtpl:33
	p.WriteNavigation(qb422016)
	//line template/build/index.qtpl:33
	qs422016 := string(qb422016.B)
	//line template/build/index.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/index.qtpl:33
	return qs422016
//line template/build/index.qtpl:33
}

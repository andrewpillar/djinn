// This file is automatically generated by qtc from "create.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/namespace/create.qtpl:2
package namespace

//line template/namespace/create.qtpl:2
import (
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/namespace/create.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/namespace/create.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/namespace/create.qtpl:10
type CreatePage struct {
	*template.Page

	Errors form.Errors
	Form   form.Form
	Parent *model.Namespace
}

//line template/namespace/create.qtpl:20
func (p *CreatePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/namespace/create.qtpl:20
	qw422016.N().S(` Create Namespace - Thrall `)
//line template/namespace/create.qtpl:22
}

//line template/namespace/create.qtpl:22
func (p *CreatePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/namespace/create.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/create.qtpl:22
	p.StreamTitle(qw422016)
	//line template/namespace/create.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/create.qtpl:22
}

//line template/namespace/create.qtpl:22
func (p *CreatePage) Title() string {
	//line template/namespace/create.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/create.qtpl:22
	p.WriteTitle(qb422016)
	//line template/namespace/create.qtpl:22
	qs422016 := string(qb422016.B)
	//line template/namespace/create.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/create.qtpl:22
	return qs422016
//line template/namespace/create.qtpl:22
}

//line template/namespace/create.qtpl:24
func (p *CreatePage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/namespace/create.qtpl:24
	qw422016.N().S(` <div class="panel"> <form class="panel-body slim" method="POST" action="/namespaces"> `)
	//line template/namespace/create.qtpl:27
	if p.Parent != nil && !p.Parent.IsZero() {
		//line template/namespace/create.qtpl:27
		qw422016.N().S(` <input type="hidden" name="parent" value="`)
		//line template/namespace/create.qtpl:28
		qw422016.E().S(p.Parent.Path)
		//line template/namespace/create.qtpl:28
		qw422016.N().S(`"/> `)
		//line template/namespace/create.qtpl:29
	}
	//line template/namespace/create.qtpl:29
	qw422016.N().S(` `)
	//line template/namespace/create.qtpl:30
	if p.Errors.First("namespace") != "" {
		//line template/namespace/create.qtpl:30
		qw422016.N().S(` <div class="form-error">Failed to create namespace: `)
		//line template/namespace/create.qtpl:31
		qw422016.E().S(p.Errors.First("namespace"))
		//line template/namespace/create.qtpl:31
		qw422016.N().S(`</div> `)
		//line template/namespace/create.qtpl:32
	}
	//line template/namespace/create.qtpl:32
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input class="form-text" type="text" name="name" value="`)
	//line template/namespace/create.qtpl:35
	qw422016.E().S(p.Form.Get("name"))
	//line template/namespace/create.qtpl:35
	qw422016.N().S(`" autocomplete="off"/> <div class="form-error">`)
	//line template/namespace/create.qtpl:36
	qw422016.E().S(p.Errors.First("name"))
	//line template/namespace/create.qtpl:36
	qw422016.N().S(`</div> </div> <div class="form-field"> <label class="label" for="description">Description <small>(optional)</small></label> <input class="form-text" type="text" name="description" value="`)
	//line template/namespace/create.qtpl:40
	qw422016.E().S(p.Form.Get("description"))
	//line template/namespace/create.qtpl:40
	qw422016.N().S(`" autocomplete="off"/> <div class="error">`)
	//line template/namespace/create.qtpl:41
	qw422016.E().S(p.Errors.First("description"))
	//line template/namespace/create.qtpl:41
	qw422016.N().S(`</div> </div> <div class="form-field"> <label class="form-option"> <input class="form-selector" type="radio" name="visibility" value="private" checked="trust"/> <strong>Private</strong> `)
	//line template/namespace/create.qtpl:47
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.094 8.016v-2.016c0-1.688-1.406-3.094-3.094-3.094s-3.094 1.406-3.094 3.094v2.016h6.188zM12 17.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-2.016 0.938-2.016 2.016 0.938 2.016 2.016 2.016zM18 8.016c1.078 0 2.016 0.891 2.016 1.969v10.031c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h0.984v-2.016c0-2.766 2.25-5.016 5.016-5.016s5.016 2.25 5.016 5.016v2.016h0.984z"></path>
</svg>
`)
	//line template/namespace/create.qtpl:47
	qw422016.N().S(` <div class="form-desc">You choose who can view builds in the namespace.</div> </label> <label class="form-option"> <input class="form-selector" type="radio" name="visibility" value="internal"/> <strong>Internal</strong> `)
	//line template/namespace/create.qtpl:53
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.984l9 4.031v6c0 5.531-3.844 10.734-9 12-5.156-1.266-9-6.469-9-12v-6zM12 12v8.953c3.703-1.172 6.469-4.828 6.984-8.953h-6.984zM12 12v-8.813l-6.984 3.094v5.719h6.984z"></path>
</svg>
`)
	//line template/namespace/create.qtpl:53
	qw422016.N().S(` <div class="form-desc">Anyone with an account will be able to view builds in the namespace.</div> </label> <label class="form-option"> <input class="form-selector" type="radio" name="visibility" value="public"/> <strong>Public</strong> `)
	//line template/namespace/create.qtpl:59
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.906 17.391c1.313-1.406 2.109-3.328 2.109-5.391 0-3.328-2.063-6.234-5.016-7.406v0.422c0 1.078-0.938 1.969-2.016 1.969h-1.969v2.016c0 0.563-0.469 0.984-1.031 0.984h-1.969v2.016h6c0.563 0 0.984 0.422 0.984 0.984v3h0.984c0.891 0 1.641 0.609 1.922 1.406zM11.016 19.922v-1.922c-1.078 0-2.016-0.938-2.016-2.016v-0.984l-4.781-4.781c-0.141 0.563-0.234 1.172-0.234 1.781 0 4.078 3.094 7.453 7.031 7.922zM12 2.016c5.531 0 9.984 4.453 9.984 9.984s-4.453 9.984-9.984 9.984-9.984-4.453-9.984-9.984 4.453-9.984 9.984-9.984z"></path>
</svg>
`)
	//line template/namespace/create.qtpl:59
	qw422016.N().S(` <div class="form-desc">Anyone will be able to view builds in the namespace.</div> </label> </div> <div class="form-field"> <button type="submit" class="btn btn-primary">Create</button> </div> </form> </div> `)
//line template/namespace/create.qtpl:68
}

//line template/namespace/create.qtpl:68
func (p *CreatePage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/namespace/create.qtpl:68
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/create.qtpl:68
	p.StreamBody(qw422016)
	//line template/namespace/create.qtpl:68
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/create.qtpl:68
}

//line template/namespace/create.qtpl:68
func (p *CreatePage) Body() string {
	//line template/namespace/create.qtpl:68
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/create.qtpl:68
	p.WriteBody(qb422016)
	//line template/namespace/create.qtpl:68
	qs422016 := string(qb422016.B)
	//line template/namespace/create.qtpl:68
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/create.qtpl:68
	return qs422016
//line template/namespace/create.qtpl:68
}

//line template/namespace/create.qtpl:70
func (p *CreatePage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/namespace/create.qtpl:70
	qw422016.N().S(` `)
	//line template/namespace/create.qtpl:71
	if p.Parent != nil && !p.Parent.IsZero() {
		//line template/namespace/create.qtpl:71
		qw422016.N().S(` <a class="back" href="`)
		//line template/namespace/create.qtpl:72
		qw422016.E().S(p.Parent.UIEndpoint())
		//line template/namespace/create.qtpl:72
		qw422016.N().S(`">`)
		//line template/namespace/create.qtpl:72
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
		//line template/namespace/create.qtpl:72
		qw422016.N().S(`</a> `)
		//line template/namespace/create.qtpl:73
		streamrenderPath(qw422016, p.Parent.User.Username, p.Parent.Path)
		//line template/namespace/create.qtpl:73
		qw422016.N().S(` - Create Sub-namespace `)
		//line template/namespace/create.qtpl:74
	} else {
		//line template/namespace/create.qtpl:74
		qw422016.N().S(` <a class="back" href="/namespaces">`)
		//line template/namespace/create.qtpl:75
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
		//line template/namespace/create.qtpl:75
		qw422016.N().S(`</a> Create Namespace `)
		//line template/namespace/create.qtpl:76
	}
	//line template/namespace/create.qtpl:76
	qw422016.N().S(` `)
//line template/namespace/create.qtpl:77
}

//line template/namespace/create.qtpl:77
func (p *CreatePage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/namespace/create.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/create.qtpl:77
	p.StreamHeader(qw422016)
	//line template/namespace/create.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/create.qtpl:77
}

//line template/namespace/create.qtpl:77
func (p *CreatePage) Header() string {
	//line template/namespace/create.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/create.qtpl:77
	p.WriteHeader(qb422016)
	//line template/namespace/create.qtpl:77
	qs422016 := string(qb422016.B)
	//line template/namespace/create.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/create.qtpl:77
	return qs422016
//line template/namespace/create.qtpl:77
}

//line template/namespace/create.qtpl:79
func (p *CreatePage) StreamActions(qw422016 *qt422016.Writer) {
//line template/namespace/create.qtpl:79
}

//line template/namespace/create.qtpl:79
func (p *CreatePage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/namespace/create.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/create.qtpl:79
	p.StreamActions(qw422016)
	//line template/namespace/create.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/create.qtpl:79
}

//line template/namespace/create.qtpl:79
func (p *CreatePage) Actions() string {
	//line template/namespace/create.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/create.qtpl:79
	p.WriteActions(qb422016)
	//line template/namespace/create.qtpl:79
	qs422016 := string(qb422016.B)
	//line template/namespace/create.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/create.qtpl:79
	return qs422016
//line template/namespace/create.qtpl:79
}

//line template/namespace/create.qtpl:80
func (p *CreatePage) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/namespace/create.qtpl:80
}

//line template/namespace/create.qtpl:80
func (p *CreatePage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/namespace/create.qtpl:80
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/create.qtpl:80
	p.StreamNavigation(qw422016)
	//line template/namespace/create.qtpl:80
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/create.qtpl:80
}

//line template/namespace/create.qtpl:80
func (p *CreatePage) Navigation() string {
	//line template/namespace/create.qtpl:80
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/create.qtpl:80
	p.WriteNavigation(qb422016)
	//line template/namespace/create.qtpl:80
	qs422016 := string(qb422016.B)
	//line template/namespace/create.qtpl:80
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/create.qtpl:80
	return qs422016
//line template/namespace/create.qtpl:80
}

// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/key/index.qtpl:2
package key

//line template/key/index.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/key/index.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/key/index.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/key/index.qtpl:9
type IndexPage struct {
	template.BasePage

	CSRF   string
	Search string
	Keys   []*model.Key
}

//line template/key/index.qtpl:19
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/key/index.qtpl:19
	qw422016.N().S(` SSH Keys - Thrall `)
//line template/key/index.qtpl:21
}

//line template/key/index.qtpl:21
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/key/index.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/key/index.qtpl:21
	p.StreamTitle(qw422016)
	//line template/key/index.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line template/key/index.qtpl:21
}

//line template/key/index.qtpl:21
func (p *IndexPage) Title() string {
	//line template/key/index.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/key/index.qtpl:21
	p.WriteTitle(qb422016)
	//line template/key/index.qtpl:21
	qs422016 := string(qb422016.B)
	//line template/key/index.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/key/index.qtpl:21
	return qs422016
//line template/key/index.qtpl:21
}

//line template/key/index.qtpl:23
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/key/index.qtpl:23
	qw422016.N().S(` <div class="panel"> `)
	//line template/key/index.qtpl:25
	if len(p.Keys) == 0 && p.Search == "" {
		//line template/key/index.qtpl:25
		qw422016.N().S(` <div class="panel-message muted">SSH keys can allow build environments to access other environments.</div> `)
		//line template/key/index.qtpl:27
	} else {
		//line template/key/index.qtpl:27
		qw422016.N().S(` <div class="panel-header">`)
		//line template/key/index.qtpl:28
		template.StreamRenderSearch(qw422016, p.URI, p.Search, "Find an SSH key...")
		//line template/key/index.qtpl:28
		qw422016.N().S(`</div> `)
		//line template/key/index.qtpl:29
		if len(p.Keys) == 0 && p.Search != "" {
			//line template/key/index.qtpl:29
			qw422016.N().S(` <div class="panel-message muted">No results found.</div> `)
			//line template/key/index.qtpl:31
		} else {
			//line template/key/index.qtpl:31
			qw422016.N().S(` <table class="table"> <thead> <tr> <th>NAME</th> <th>NAMESPACE</th> <th></th> <th></th> </tr> </thead> <tbody> `)
			//line template/key/index.qtpl:42
			for _, k := range p.Keys {
				//line template/key/index.qtpl:42
				qw422016.N().S(` <tr> <td><a href="`)
				//line template/key/index.qtpl:44
				qw422016.E().S(k.UIEndpoint("edit"))
				//line template/key/index.qtpl:44
				qw422016.N().S(`">`)
				//line template/key/index.qtpl:44
				qw422016.E().S(k.Name)
				//line template/key/index.qtpl:44
				qw422016.N().S(`</a></td> <td> `)
				//line template/key/index.qtpl:46
				if k.Namespace != nil {
					//line template/key/index.qtpl:46
					qw422016.N().S(` <a href="`)
					//line template/key/index.qtpl:47
					qw422016.E().S(k.Namespace.UIEndpoint())
					//line template/key/index.qtpl:47
					qw422016.N().S(`">`)
					//line template/key/index.qtpl:47
					qw422016.E().S(k.Namespace.Path)
					//line template/key/index.qtpl:47
					qw422016.N().S(`</a> `)
					//line template/key/index.qtpl:48
				} else {
					//line template/key/index.qtpl:48
					qw422016.N().S(` <span class="muted">--</span> `)
					//line template/key/index.qtpl:50
				}
				//line template/key/index.qtpl:50
				qw422016.N().S(` </td> <td class="align-right"> `)
				//line template/key/index.qtpl:53
				if p.User.ID != k.UserID {
					//line template/key/index.qtpl:53
					qw422016.N().S(` <span class="muted">`)
					//line template/key/index.qtpl:54
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.984 12.984c2.344 0 7.031 1.172 7.031 3.516v2.484h-6v-2.484c0-1.5-0.797-2.625-1.969-3.469 0.328-0.047 0.656-0.047 0.938-0.047zM8.016 12.984c2.344 0 6.984 1.172 6.984 3.516v2.484h-14.016v-2.484c0-2.344 4.688-3.516 7.031-3.516zM8.016 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 2.953 1.359 2.953 3-1.313 3-2.953 3zM15.984 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 3 1.359 3 3-1.359 3-3 3z"></path>
</svg>
`)
					//line template/key/index.qtpl:54
					qw422016.N().S(`</span> `)
					//line template/key/index.qtpl:55
				}
				//line template/key/index.qtpl:55
				qw422016.N().S(` </td> <td class="align-right"> `)
				//line template/key/index.qtpl:58
				if p.User.ID == k.UserID || k.Namespace != nil && k.Namespace.UserID == p.User.ID {
					//line template/key/index.qtpl:58
					qw422016.N().S(` <form method="POST" action="`)
					//line template/key/index.qtpl:59
					qw422016.E().S(k.UIEndpoint())
					//line template/key/index.qtpl:59
					qw422016.N().S(`"> `)
					//line template/key/index.qtpl:60
					qw422016.N().S(p.CSRF)
					//line template/key/index.qtpl:60
					qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <button type="submit" class="btn btn-danger">Delete</button> </form> `)
					//line template/key/index.qtpl:64
				}
				//line template/key/index.qtpl:64
				qw422016.N().S(` </td> </tr> `)
				//line template/key/index.qtpl:67
			}
			//line template/key/index.qtpl:67
			qw422016.N().S(` </tbody> </table> `)
			//line template/key/index.qtpl:70
		}
		//line template/key/index.qtpl:70
		qw422016.N().S(` `)
		//line template/key/index.qtpl:71
	}
	//line template/key/index.qtpl:71
	qw422016.N().S(` </div> `)
//line template/key/index.qtpl:73
}

//line template/key/index.qtpl:73
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/key/index.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/key/index.qtpl:73
	p.StreamBody(qw422016)
	//line template/key/index.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line template/key/index.qtpl:73
}

//line template/key/index.qtpl:73
func (p *IndexPage) Body() string {
	//line template/key/index.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/key/index.qtpl:73
	p.WriteBody(qb422016)
	//line template/key/index.qtpl:73
	qs422016 := string(qb422016.B)
	//line template/key/index.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/key/index.qtpl:73
	return qs422016
//line template/key/index.qtpl:73
}

//line template/key/index.qtpl:75
func (p *IndexPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/key/index.qtpl:75
	qw422016.N().S(` SSH Keys `)
//line template/key/index.qtpl:77
}

//line template/key/index.qtpl:77
func (p *IndexPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/key/index.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/key/index.qtpl:77
	p.StreamHeader(qw422016)
	//line template/key/index.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line template/key/index.qtpl:77
}

//line template/key/index.qtpl:77
func (p *IndexPage) Header() string {
	//line template/key/index.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/key/index.qtpl:77
	p.WriteHeader(qb422016)
	//line template/key/index.qtpl:77
	qs422016 := string(qb422016.B)
	//line template/key/index.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/key/index.qtpl:77
	return qs422016
//line template/key/index.qtpl:77
}

//line template/key/index.qtpl:79
func (p *IndexPage) StreamActions(qw422016 *qt422016.Writer) {
	//line template/key/index.qtpl:79
	qw422016.N().S(` <li><a href="/keys/create" class="btn btn-primary">Create</a></li> `)
//line template/key/index.qtpl:81
}

//line template/key/index.qtpl:81
func (p *IndexPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/key/index.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/key/index.qtpl:81
	p.StreamActions(qw422016)
	//line template/key/index.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line template/key/index.qtpl:81
}

//line template/key/index.qtpl:81
func (p *IndexPage) Actions() string {
	//line template/key/index.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/key/index.qtpl:81
	p.WriteActions(qb422016)
	//line template/key/index.qtpl:81
	qs422016 := string(qb422016.B)
	//line template/key/index.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/key/index.qtpl:81
	return qs422016
//line template/key/index.qtpl:81
}

//line template/key/index.qtpl:83
func (p *IndexPage) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/key/index.qtpl:83
}

//line template/key/index.qtpl:83
func (p *IndexPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/key/index.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/key/index.qtpl:83
	p.StreamNavigation(qw422016)
	//line template/key/index.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line template/key/index.qtpl:83
}

//line template/key/index.qtpl:83
func (p *IndexPage) Navigation() string {
	//line template/key/index.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/key/index.qtpl:83
	p.WriteNavigation(qb422016)
	//line template/key/index.qtpl:83
	qs422016 := string(qb422016.B)
	//line template/key/index.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/key/index.qtpl:83
	return qs422016
//line template/key/index.qtpl:83
}

// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/namespace/index.qtpl:2
package namespace

//line template/namespace/index.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/model/types"
	"github.com/andrewpillar/thrall/template"
)

//line template/namespace/index.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/namespace/index.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/namespace/index.qtpl:10
type IndexPage struct {
	template.BasePage

	Paginator  model.Paginator
	Namespaces []*model.Namespace
	Search     string
}

//line template/namespace/index.qtpl:20
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/namespace/index.qtpl:20
	qw422016.N().S(` Namespaces - Thrall `)
//line template/namespace/index.qtpl:22
}

//line template/namespace/index.qtpl:22
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/namespace/index.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/index.qtpl:22
	p.StreamTitle(qw422016)
	//line template/namespace/index.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/index.qtpl:22
}

//line template/namespace/index.qtpl:22
func (p *IndexPage) Title() string {
	//line template/namespace/index.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/index.qtpl:22
	p.WriteTitle(qb422016)
	//line template/namespace/index.qtpl:22
	qs422016 := string(qb422016.B)
	//line template/namespace/index.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/index.qtpl:22
	return qs422016
//line template/namespace/index.qtpl:22
}

//line template/namespace/index.qtpl:24
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/namespace/index.qtpl:24
	qw422016.N().S(` <div class="panel"> `)
	//line template/namespace/index.qtpl:26
	if len(p.Namespaces) == 0 && p.Search == "" {
		//line template/namespace/index.qtpl:26
		qw422016.N().S(` <div class="panel-message muted">Namespaces allow you to group related build resources together.</div> `)
		//line template/namespace/index.qtpl:28
	} else {
		//line template/namespace/index.qtpl:28
		qw422016.N().S(` <div class="panel-header">`)
		//line template/namespace/index.qtpl:29
		template.StreamRenderSearch(qw422016, p.URL.Path, p.Search, "Find a namespace...")
		//line template/namespace/index.qtpl:29
		qw422016.N().S(`</div> `)
		//line template/namespace/index.qtpl:30
		if len(p.Namespaces) == 0 && p.Search != "" {
			//line template/namespace/index.qtpl:30
			qw422016.N().S(` <div class="panel-message muted">No results found.</div> `)
			//line template/namespace/index.qtpl:32
		} else {
			//line template/namespace/index.qtpl:32
			qw422016.N().S(` <table class="table"> <thead> <tr> <th>NAME</th> <th>LAST BUILD</th> <th></th> <th></th> </tr> </thead> <tbody> `)
			//line template/namespace/index.qtpl:43
			for _, n := range p.Namespaces {
				//line template/namespace/index.qtpl:43
				qw422016.N().S(` <tr> <td> <a href="`)
				//line template/namespace/index.qtpl:46
				qw422016.E().S(n.UIEndpoint())
				//line template/namespace/index.qtpl:46
				qw422016.N().S(`">`)
				//line template/namespace/index.qtpl:46
				qw422016.E().S(n.Path)
				//line template/namespace/index.qtpl:46
				qw422016.N().S(`</a> <div class="muted"> `)
				//line template/namespace/index.qtpl:48
				if n.Description == "" {
					//line template/namespace/index.qtpl:48
					qw422016.N().S(` <em>No description</em> `)
					//line template/namespace/index.qtpl:50
				} else {
					//line template/namespace/index.qtpl:50
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:51
					qw422016.E().S(n.Description)
					//line template/namespace/index.qtpl:51
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:52
				}
				//line template/namespace/index.qtpl:52
				qw422016.N().S(` </div> </td> <td class="muted"> `)
				//line template/namespace/index.qtpl:56
				if n.LastBuild != nil {
					//line template/namespace/index.qtpl:56
					qw422016.N().S(` <a href="`)
					//line template/namespace/index.qtpl:57
					qw422016.E().S(n.LastBuild.UIEndpoint())
					//line template/namespace/index.qtpl:57
					qw422016.N().S(`"> #`)
					//line template/namespace/index.qtpl:58
					qw422016.E().V(n.LastBuild.ID)
					//line template/namespace/index.qtpl:58
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:58
					qw422016.E().S(n.LastBuild.Trigger.CommentTitle())
					//line template/namespace/index.qtpl:58
					qw422016.N().S(` </a> `)
					//line template/namespace/index.qtpl:60
				} else {
					//line template/namespace/index.qtpl:60
					qw422016.N().S(` -- `)
					//line template/namespace/index.qtpl:62
				}
				//line template/namespace/index.qtpl:62
				qw422016.N().S(` </td> <td> `)
				//line template/namespace/index.qtpl:65
				if n.LastBuild != nil {
					//line template/namespace/index.qtpl:65
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:66
					template.StreamRenderStatus(qw422016, n.LastBuild.Status)
					//line template/namespace/index.qtpl:66
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:67
				}
				//line template/namespace/index.qtpl:67
				qw422016.N().S(` </td> <td class="align-right muted"> `)
				//line template/namespace/index.qtpl:70
				if p.User.ID != n.UserID {
					//line template/namespace/index.qtpl:70
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:71
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.984 12.984c2.344 0 7.031 1.172 7.031 3.516v2.484h-6v-2.484c0-1.5-0.797-2.625-1.969-3.469 0.328-0.047 0.656-0.047 0.938-0.047zM8.016 12.984c2.344 0 6.984 1.172 6.984 3.516v2.484h-14.016v-2.484c0-2.344 4.688-3.516 7.031-3.516zM8.016 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 2.953 1.359 2.953 3-1.313 3-2.953 3zM15.984 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 3 1.359 3 3-1.359 3-3 3z"></path>
</svg>
`)
					//line template/namespace/index.qtpl:71
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:72
				}
				//line template/namespace/index.qtpl:72
				qw422016.N().S(` `)
				//line template/namespace/index.qtpl:73
				switch n.Visibility {
				//line template/namespace/index.qtpl:74
				case types.Private:
					//line template/namespace/index.qtpl:74
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:75
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.094 8.016v-2.016c0-1.688-1.406-3.094-3.094-3.094s-3.094 1.406-3.094 3.094v2.016h6.188zM12 17.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-2.016 0.938-2.016 2.016 0.938 2.016 2.016 2.016zM18 8.016c1.078 0 2.016 0.891 2.016 1.969v10.031c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h0.984v-2.016c0-2.766 2.25-5.016 5.016-5.016s5.016 2.25 5.016 5.016v2.016h0.984z"></path>
</svg>
`)
					//line template/namespace/index.qtpl:75
					qw422016.N().S(` `)
				//line template/namespace/index.qtpl:76
				case types.Internal:
					//line template/namespace/index.qtpl:76
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:77
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.984l9 4.031v6c0 5.531-3.844 10.734-9 12-5.156-1.266-9-6.469-9-12v-6zM12 12v8.953c3.703-1.172 6.469-4.828 6.984-8.953h-6.984zM12 12v-8.813l-6.984 3.094v5.719h6.984z"></path>
</svg>
`)
					//line template/namespace/index.qtpl:77
					qw422016.N().S(` `)
				//line template/namespace/index.qtpl:78
				case types.Public:
					//line template/namespace/index.qtpl:78
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:79
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.906 17.391c1.313-1.406 2.109-3.328 2.109-5.391 0-3.328-2.063-6.234-5.016-7.406v0.422c0 1.078-0.938 1.969-2.016 1.969h-1.969v2.016c0 0.563-0.469 0.984-1.031 0.984h-1.969v2.016h6c0.563 0 0.984 0.422 0.984 0.984v3h0.984c0.891 0 1.641 0.609 1.922 1.406zM11.016 19.922v-1.922c-1.078 0-2.016-0.938-2.016-2.016v-0.984l-4.781-4.781c-0.141 0.563-0.234 1.172-0.234 1.781 0 4.078 3.094 7.453 7.031 7.922zM12 2.016c5.531 0 9.984 4.453 9.984 9.984s-4.453 9.984-9.984 9.984-9.984-4.453-9.984-9.984 4.453-9.984 9.984-9.984z"></path>
</svg>
`)
					//line template/namespace/index.qtpl:79
					qw422016.N().S(` `)
					//line template/namespace/index.qtpl:80
				}
				//line template/namespace/index.qtpl:80
				qw422016.N().S(` </td> </tr> `)
				//line template/namespace/index.qtpl:83
			}
			//line template/namespace/index.qtpl:83
			qw422016.N().S(` </tbody> </table> `)
			//line template/namespace/index.qtpl:86
		}
		//line template/namespace/index.qtpl:86
		qw422016.N().S(` `)
		//line template/namespace/index.qtpl:87
	}
	//line template/namespace/index.qtpl:87
	qw422016.N().S(` </div> `)
	//line template/namespace/index.qtpl:89
	template.StreamRenderPaginator(qw422016, p.URL.Path, p.Paginator)
	//line template/namespace/index.qtpl:89
	qw422016.N().S(` `)
//line template/namespace/index.qtpl:90
}

//line template/namespace/index.qtpl:90
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/namespace/index.qtpl:90
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/index.qtpl:90
	p.StreamBody(qw422016)
	//line template/namespace/index.qtpl:90
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/index.qtpl:90
}

//line template/namespace/index.qtpl:90
func (p *IndexPage) Body() string {
	//line template/namespace/index.qtpl:90
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/index.qtpl:90
	p.WriteBody(qb422016)
	//line template/namespace/index.qtpl:90
	qs422016 := string(qb422016.B)
	//line template/namespace/index.qtpl:90
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/index.qtpl:90
	return qs422016
//line template/namespace/index.qtpl:90
}

//line template/namespace/index.qtpl:92
func (p *IndexPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/namespace/index.qtpl:92
	qw422016.N().S(` Namespaces `)
//line template/namespace/index.qtpl:94
}

//line template/namespace/index.qtpl:94
func (p *IndexPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/namespace/index.qtpl:94
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/index.qtpl:94
	p.StreamHeader(qw422016)
	//line template/namespace/index.qtpl:94
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/index.qtpl:94
}

//line template/namespace/index.qtpl:94
func (p *IndexPage) Header() string {
	//line template/namespace/index.qtpl:94
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/index.qtpl:94
	p.WriteHeader(qb422016)
	//line template/namespace/index.qtpl:94
	qs422016 := string(qb422016.B)
	//line template/namespace/index.qtpl:94
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/index.qtpl:94
	return qs422016
//line template/namespace/index.qtpl:94
}

//line template/namespace/index.qtpl:96
func (p *IndexPage) StreamActions(qw422016 *qt422016.Writer) {
	//line template/namespace/index.qtpl:96
	qw422016.N().S(` <li><a href="/namespaces/create" class="btn btn-primary">Create</a></li> `)
//line template/namespace/index.qtpl:98
}

//line template/namespace/index.qtpl:98
func (p *IndexPage) WriteActions(qq422016 qtio422016.Writer) {
	//line template/namespace/index.qtpl:98
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/index.qtpl:98
	p.StreamActions(qw422016)
	//line template/namespace/index.qtpl:98
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/index.qtpl:98
}

//line template/namespace/index.qtpl:98
func (p *IndexPage) Actions() string {
	//line template/namespace/index.qtpl:98
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/index.qtpl:98
	p.WriteActions(qb422016)
	//line template/namespace/index.qtpl:98
	qs422016 := string(qb422016.B)
	//line template/namespace/index.qtpl:98
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/index.qtpl:98
	return qs422016
//line template/namespace/index.qtpl:98
}

//line template/namespace/index.qtpl:100
func (p *IndexPage) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/namespace/index.qtpl:100
}

//line template/namespace/index.qtpl:100
func (p *IndexPage) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/namespace/index.qtpl:100
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/index.qtpl:100
	p.StreamNavigation(qw422016)
	//line template/namespace/index.qtpl:100
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/index.qtpl:100
}

//line template/namespace/index.qtpl:100
func (p *IndexPage) Navigation() string {
	//line template/namespace/index.qtpl:100
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/index.qtpl:100
	p.WriteNavigation(qb422016)
	//line template/namespace/index.qtpl:100
	qs422016 := string(qb422016.B)
	//line template/namespace/index.qtpl:100
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/index.qtpl:100
	return qs422016
//line template/namespace/index.qtpl:100
}

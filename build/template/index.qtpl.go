// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line build/template/index.qtpl:2
package template

//line build/template/index.qtpl:2
import (
	"strings"

	"github.com/andrewpillar/thrall/build"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line build/template/index.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line build/template/index.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line build/template/index.qtpl:12
type Index struct {
	template.BasePage

	Paginator model.Paginator
	Builds    []*build.Build
	Search    string
	Status    string
	Tag       string
}

//line build/template/index.qtpl:24
func (p *Index) StreamTitle(qw422016 *qt422016.Writer) {
//line build/template/index.qtpl:24
	qw422016.N().S(` Builds - Thrall `)
//line build/template/index.qtpl:26
}

//line build/template/index.qtpl:26
func (p *Index) WriteTitle(qq422016 qtio422016.Writer) {
//line build/template/index.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line build/template/index.qtpl:26
	p.StreamTitle(qw422016)
//line build/template/index.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line build/template/index.qtpl:26
}

//line build/template/index.qtpl:26
func (p *Index) Title() string {
//line build/template/index.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line build/template/index.qtpl:26
	p.WriteTitle(qb422016)
//line build/template/index.qtpl:26
	qs422016 := string(qb422016.B)
//line build/template/index.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line build/template/index.qtpl:26
	return qs422016
//line build/template/index.qtpl:26
}

//line build/template/index.qtpl:28
func (p *Index) StreamBody(qw422016 *qt422016.Writer) {
//line build/template/index.qtpl:28
	qw422016.N().S(` <div class="panel"> `)
//line build/template/index.qtpl:30
	if len(p.Builds) == 0 && p.Search == "" && p.Status == "" {
//line build/template/index.qtpl:30
		qw422016.N().S(` <div class="panel-message muted">No builds have been submitted yet.</div> `)
//line build/template/index.qtpl:32
	} else {
//line build/template/index.qtpl:32
		qw422016.N().S(` <div class="panel-header"> `)
//line build/template/index.qtpl:34
		StreamRenderStatusNav(qw422016, p.URL.Path, p.Status)
//line build/template/index.qtpl:34
		qw422016.N().S(` `)
//line build/template/index.qtpl:35
		template.StreamRenderSearch(qw422016, p.URL.Path, p.Search, "Find a build...")
//line build/template/index.qtpl:35
		qw422016.N().S(` </div> `)
//line build/template/index.qtpl:37
		if len(p.Builds) == 0 && p.Search != "" {
//line build/template/index.qtpl:37
			qw422016.N().S(` <div class="panel-message muted">No results found.</div> `)
//line build/template/index.qtpl:39
		} else if len(p.Builds) == 0 && p.Status != "" {
//line build/template/index.qtpl:39
			qw422016.N().S(` <div class="panel-message muted">No `)
//line build/template/index.qtpl:40
			qw422016.E().S(strings.Replace(p.Status, "_", " ", -1))
//line build/template/index.qtpl:40
			qw422016.N().S(` builds.</div> `)
//line build/template/index.qtpl:41
		} else {
//line build/template/index.qtpl:41
			qw422016.N().S(` <table class="table"> <thead> <tr> <th>STATUS</th> <th>BUILD</th> <th>NAMESPACE</th> <th></th> <th></th> <th></th> </tr> </thead> <tbody> `)
//line build/template/index.qtpl:54
			for _, b := range p.Builds {
//line build/template/index.qtpl:54
				qw422016.N().S(` <tr> <td>`)
//line build/template/index.qtpl:56
				template.StreamRenderStatus(qw422016, b.Status)
//line build/template/index.qtpl:56
				qw422016.N().S(`</td> <td><a href="`)
//line build/template/index.qtpl:57
				qw422016.E().S(b.Endpoint())
//line build/template/index.qtpl:57
				qw422016.N().S(`">#`)
//line build/template/index.qtpl:57
				qw422016.E().V(b.ID)
//line build/template/index.qtpl:57
				if b.Trigger.Comment != "" {
//line build/template/index.qtpl:57
					qw422016.N().S(` - `)
//line build/template/index.qtpl:57
					qw422016.E().S(b.Trigger.CommentTitle())
//line build/template/index.qtpl:57
				}
//line build/template/index.qtpl:57
				qw422016.N().S(`</a></td> <td> `)
//line build/template/index.qtpl:59
				if b.Namespace != nil {
//line build/template/index.qtpl:59
					qw422016.N().S(` <a href="`)
//line build/template/index.qtpl:60
					qw422016.E().S(b.Namespace.Endpoint())
//line build/template/index.qtpl:60
					qw422016.N().S(`">`)
//line build/template/index.qtpl:60
					qw422016.E().V(b.Namespace.Values()["path"])
//line build/template/index.qtpl:60
					qw422016.N().S(`</a> `)
//line build/template/index.qtpl:61
				} else {
//line build/template/index.qtpl:61
					qw422016.N().S(` <span class="muted">--</span> `)
//line build/template/index.qtpl:63
				}
//line build/template/index.qtpl:63
				qw422016.N().S(` </td> <td class="align-right"> `)
//line build/template/index.qtpl:66
				for _, t := range b.Tags {
//line build/template/index.qtpl:66
					qw422016.N().S(` <a class="pill pill-light" href="?tag=`)
//line build/template/index.qtpl:67
					qw422016.E().S(t.Name)
//line build/template/index.qtpl:67
					qw422016.N().S(`">`)
//line build/template/index.qtpl:67
					qw422016.E().S(t.Name)
//line build/template/index.qtpl:67
					qw422016.N().S(`</a> `)
//line build/template/index.qtpl:68
				}
//line build/template/index.qtpl:68
				qw422016.N().S(` </td> <td class="align-right"> `)
//line build/template/index.qtpl:71
				if !b.FinishedAt.Valid || !b.StartedAt.Valid {
//line build/template/index.qtpl:71
					qw422016.N().S(` <span class="muted">--</span> `)
//line build/template/index.qtpl:73
				} else {
//line build/template/index.qtpl:73
					qw422016.N().S(` `)
//line build/template/index.qtpl:74
					qw422016.E().V(b.FinishedAt.Time.Sub(b.StartedAt.Time))
//line build/template/index.qtpl:74
					qw422016.N().S(` `)
//line build/template/index.qtpl:75
				}
//line build/template/index.qtpl:75
				qw422016.N().S(` </td> <td class="align-right"> `)
//line build/template/index.qtpl:78
				if p.User.ID != b.UserID {
//line build/template/index.qtpl:78
					qw422016.N().S(` <span class="muted">`)
//line build/template/index.qtpl:79
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.984 12.984c2.344 0 7.031 1.172 7.031 3.516v2.484h-6v-2.484c0-1.5-0.797-2.625-1.969-3.469 0.328-0.047 0.656-0.047 0.938-0.047zM8.016 12.984c2.344 0 6.984 1.172 6.984 3.516v2.484h-14.016v-2.484c0-2.344 4.688-3.516 7.031-3.516zM8.016 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 2.953 1.359 2.953 3-1.313 3-2.953 3zM15.984 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 3 1.359 3 3-1.359 3-3 3z"></path>
</svg>
`)
//line build/template/index.qtpl:79
					qw422016.N().S(`</span> `)
//line build/template/index.qtpl:80
				}
//line build/template/index.qtpl:80
				qw422016.N().S(` </td> </tr> `)
//line build/template/index.qtpl:83
			}
//line build/template/index.qtpl:83
			qw422016.N().S(` </tbody> </table> `)
//line build/template/index.qtpl:86
		}
//line build/template/index.qtpl:86
		qw422016.N().S(` `)
//line build/template/index.qtpl:87
	}
//line build/template/index.qtpl:87
	qw422016.N().S(` </div> `)
//line build/template/index.qtpl:89
	template.StreamRenderPaginator(qw422016, p.URL, p.Paginator)
//line build/template/index.qtpl:89
	qw422016.N().S(` `)
//line build/template/index.qtpl:90
}

//line build/template/index.qtpl:90
func (p *Index) WriteBody(qq422016 qtio422016.Writer) {
//line build/template/index.qtpl:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line build/template/index.qtpl:90
	p.StreamBody(qw422016)
//line build/template/index.qtpl:90
	qt422016.ReleaseWriter(qw422016)
//line build/template/index.qtpl:90
}

//line build/template/index.qtpl:90
func (p *Index) Body() string {
//line build/template/index.qtpl:90
	qb422016 := qt422016.AcquireByteBuffer()
//line build/template/index.qtpl:90
	p.WriteBody(qb422016)
//line build/template/index.qtpl:90
	qs422016 := string(qb422016.B)
//line build/template/index.qtpl:90
	qt422016.ReleaseByteBuffer(qb422016)
//line build/template/index.qtpl:90
	return qs422016
//line build/template/index.qtpl:90
}

//line build/template/index.qtpl:92
func (p *Index) StreamSection(qw422016 *qt422016.Writer) {
//line build/template/index.qtpl:92
	qw422016.N().S(` `)
//line build/template/index.qtpl:93
	p.StreamBody(qw422016)
//line build/template/index.qtpl:93
	qw422016.N().S(` `)
//line build/template/index.qtpl:94
}

//line build/template/index.qtpl:94
func (p *Index) WriteSection(qq422016 qtio422016.Writer) {
//line build/template/index.qtpl:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line build/template/index.qtpl:94
	p.StreamSection(qw422016)
//line build/template/index.qtpl:94
	qt422016.ReleaseWriter(qw422016)
//line build/template/index.qtpl:94
}

//line build/template/index.qtpl:94
func (p *Index) Section() string {
//line build/template/index.qtpl:94
	qb422016 := qt422016.AcquireByteBuffer()
//line build/template/index.qtpl:94
	p.WriteSection(qb422016)
//line build/template/index.qtpl:94
	qs422016 := string(qb422016.B)
//line build/template/index.qtpl:94
	qt422016.ReleaseByteBuffer(qb422016)
//line build/template/index.qtpl:94
	return qs422016
//line build/template/index.qtpl:94
}

//line build/template/index.qtpl:96
func (p *Index) StreamHeader(qw422016 *qt422016.Writer) {
//line build/template/index.qtpl:96
	qw422016.N().S(` Builds `)
//line build/template/index.qtpl:98
	if p.Tag != "" {
//line build/template/index.qtpl:98
		qw422016.N().S(` <span class="pill pill-light">`)
//line build/template/index.qtpl:99
		qw422016.E().S(p.Tag)
//line build/template/index.qtpl:99
		qw422016.N().S(`<a href="`)
//line build/template/index.qtpl:99
		qw422016.E().S(p.URL.Path)
//line build/template/index.qtpl:99
		qw422016.N().S(`">`)
//line build/template/index.qtpl:99
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
//line build/template/index.qtpl:99
		qw422016.N().S(`</a></span> `)
//line build/template/index.qtpl:100
	}
//line build/template/index.qtpl:100
	qw422016.N().S(` `)
//line build/template/index.qtpl:101
}

//line build/template/index.qtpl:101
func (p *Index) WriteHeader(qq422016 qtio422016.Writer) {
//line build/template/index.qtpl:101
	qw422016 := qt422016.AcquireWriter(qq422016)
//line build/template/index.qtpl:101
	p.StreamHeader(qw422016)
//line build/template/index.qtpl:101
	qt422016.ReleaseWriter(qw422016)
//line build/template/index.qtpl:101
}

//line build/template/index.qtpl:101
func (p *Index) Header() string {
//line build/template/index.qtpl:101
	qb422016 := qt422016.AcquireByteBuffer()
//line build/template/index.qtpl:101
	p.WriteHeader(qb422016)
//line build/template/index.qtpl:101
	qs422016 := string(qb422016.B)
//line build/template/index.qtpl:101
	qt422016.ReleaseByteBuffer(qb422016)
//line build/template/index.qtpl:101
	return qs422016
//line build/template/index.qtpl:101
}

//line build/template/index.qtpl:103
func (p *Index) StreamActions(qw422016 *qt422016.Writer) {
//line build/template/index.qtpl:103
	qw422016.N().S(` `)
//line build/template/index.qtpl:104
	if _, ok := p.User.Permissions["build:write"]; ok {
//line build/template/index.qtpl:104
		qw422016.N().S(` <li><a href="/builds/create" class="btn btn-primary">Submit</a></li> `)
//line build/template/index.qtpl:106
	}
//line build/template/index.qtpl:106
	qw422016.N().S(` `)
//line build/template/index.qtpl:107
}

//line build/template/index.qtpl:107
func (p *Index) WriteActions(qq422016 qtio422016.Writer) {
//line build/template/index.qtpl:107
	qw422016 := qt422016.AcquireWriter(qq422016)
//line build/template/index.qtpl:107
	p.StreamActions(qw422016)
//line build/template/index.qtpl:107
	qt422016.ReleaseWriter(qw422016)
//line build/template/index.qtpl:107
}

//line build/template/index.qtpl:107
func (p *Index) Actions() string {
//line build/template/index.qtpl:107
	qb422016 := qt422016.AcquireByteBuffer()
//line build/template/index.qtpl:107
	p.WriteActions(qb422016)
//line build/template/index.qtpl:107
	qs422016 := string(qb422016.B)
//line build/template/index.qtpl:107
	qt422016.ReleaseByteBuffer(qb422016)
//line build/template/index.qtpl:107
	return qs422016
//line build/template/index.qtpl:107
}

//line build/template/index.qtpl:109
func (p *Index) StreamNavigation(qw422016 *qt422016.Writer) {
//line build/template/index.qtpl:109
}

//line build/template/index.qtpl:109
func (p *Index) WriteNavigation(qq422016 qtio422016.Writer) {
//line build/template/index.qtpl:109
	qw422016 := qt422016.AcquireWriter(qq422016)
//line build/template/index.qtpl:109
	p.StreamNavigation(qw422016)
//line build/template/index.qtpl:109
	qt422016.ReleaseWriter(qw422016)
//line build/template/index.qtpl:109
}

//line build/template/index.qtpl:109
func (p *Index) Navigation() string {
//line build/template/index.qtpl:109
	qb422016 := qt422016.AcquireByteBuffer()
//line build/template/index.qtpl:109
	p.WriteNavigation(qb422016)
//line build/template/index.qtpl:109
	qs422016 := string(qb422016.B)
//line build/template/index.qtpl:109
	qt422016.ReleaseByteBuffer(qb422016)
//line build/template/index.qtpl:109
	return qs422016
//line build/template/index.qtpl:109
}

// This file is automatically generated by qtc from "render.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/render.qtpl:2
package build

//line template/build/render.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/build/render.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/render.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/render.qtpl:9
func StreamRenderTable(qw422016 *qt422016.Writer, builds []*model.Build, status, uri string) {
	//line template/build/render.qtpl:9
	qw422016.N().S(` <div class="panel"> <div class="panel-header"> `)
	//line template/build/render.qtpl:12
	qw422016.N().S(`<ul class="panel-nav"><li><a href="`)
	//line template/build/render.qtpl:14
	qw422016.E().S(uri)
	//line template/build/render.qtpl:14
	qw422016.N().S(`"`)
	//line template/build/render.qtpl:14
	if status == "" {
		//line template/build/render.qtpl:14
		qw422016.N().S(`class="active"`)
		//line template/build/render.qtpl:14
	}
	//line template/build/render.qtpl:14
	qw422016.N().S(`>`)
	//line template/build/render.qtpl:14
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M22.688 18.984c0.422 0.281 0.422 0.984-0.094 1.406l-2.297 2.297c-0.422 0.422-0.984 0.422-1.406 0l-9.094-9.094c-2.297 0.891-4.969 0.422-6.891-1.5-2.016-2.016-2.531-5.016-1.313-7.406l4.406 4.313 3-3-4.313-4.313c2.391-1.078 5.391-0.703 7.406 1.313 1.922 1.922 2.391 4.594 1.5 6.891z"></path>
</svg>
`)
	//line template/build/render.qtpl:14
	qw422016.N().S(`<span>All</span></a></li><li><a href="?status=queued"`)
	//line template/build/render.qtpl:15
	if status == "queued" {
		//line template/build/render.qtpl:15
		qw422016.N().S(`class="active"`)
		//line template/build/render.qtpl:15
	}
	//line template/build/render.qtpl:15
	qw422016.N().S(`>`)
	//line template/build/render.qtpl:15
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 11.484l3.984-3.984v-3.516h-7.969v3.516zM15.984 16.5l-3.984-3.984-3.984 3.984v3.516h7.969v-3.516zM6 2.016h12v6l-3.984 3.984 3.984 3.984v6h-12v-6l3.984-3.984-3.984-3.984v-6z"></path>
</svg>
`)
	//line template/build/render.qtpl:15
	qw422016.N().S(`<span>Queued</span></a></li><li><a href="?status=running"`)
	//line template/build/render.qtpl:16
	if status == "running" {
		//line template/build/render.qtpl:16
		qw422016.N().S(`class="active"`)
		//line template/build/render.qtpl:16
	}
	//line template/build/render.qtpl:16
	qw422016.N().S(`>`)
	//line template/build/render.qtpl:16
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M21.984 12c0 5.156-3.938 9.422-8.953 9.938v-2.016c3.938-0.516 6.984-3.891 6.984-7.922s-3.047-7.406-6.984-7.922v-2.016c5.016 0.516 8.953 4.781 8.953 9.938zM5.672 19.734l1.406-1.406c1.125 0.844 2.484 1.406 3.938 1.594v2.016c-2.016-0.188-3.844-0.984-5.344-2.203zM4.078 12.984c0.188 1.453 0.75 2.813 1.594 3.891l-1.406 1.453c-1.219-1.5-2.016-3.328-2.203-5.344h2.016zM5.672 7.078c-0.844 1.125-1.406 2.484-1.594 3.938h-2.016c0.188-2.016 0.984-3.844 2.203-5.344zM11.016 4.078c-1.453 0.188-2.813 0.75-3.938 1.594l-1.406-1.406c1.5-1.219 3.328-2.016 5.344-2.203v2.016zM13.031 9.797l2.953 2.203c-2.007 1.493-4.007 2.993-6 4.5z"></path>
</svg>
`)
	//line template/build/render.qtpl:16
	qw422016.N().S(`<span>Running</span></a></li><li><a href="?status=passed"`)
	//line template/build/render.qtpl:17
	if status == "passed" {
		//line template/build/render.qtpl:17
		qw422016.N().S(`class="active"`)
		//line template/build/render.qtpl:17
	}
	//line template/build/render.qtpl:17
	qw422016.N().S(`>`)
	//line template/build/render.qtpl:17
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
	//line template/build/render.qtpl:17
	qw422016.N().S(`<span>Passed</span></a></li><li><a href="?status=failed"`)
	//line template/build/render.qtpl:18
	if status == "failed" {
		//line template/build/render.qtpl:18
		qw422016.N().S(`class="active"`)
		//line template/build/render.qtpl:18
	}
	//line template/build/render.qtpl:18
	qw422016.N().S(`>`)
	//line template/build/render.qtpl:18
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
	//line template/build/render.qtpl:18
	qw422016.N().S(`<span>Failed</span></a></li></ul>`)
	//line template/build/render.qtpl:20
	qw422016.N().S(` <form class="panel-search"> <input type="text" name="search" class="text" placeholder="Find a build..." autocomplete="off"/> </form> </div> `)
	//line template/build/render.qtpl:25
	if len(builds) == 0 {
		//line template/build/render.qtpl:25
		qw422016.N().S(` <div class="panel-message muted"> `)
		//line template/build/render.qtpl:27
		if status == "" {
			//line template/build/render.qtpl:27
			qw422016.N().S(` No builds have been submitted yet. `)
			//line template/build/render.qtpl:29
		} else {
			//line template/build/render.qtpl:29
			qw422016.N().S(` No `)
			//line template/build/render.qtpl:30
			qw422016.E().S(status)
			//line template/build/render.qtpl:30
			qw422016.N().S(` builds. `)
			//line template/build/render.qtpl:31
		}
		//line template/build/render.qtpl:31
		qw422016.N().S(` </div> `)
		//line template/build/render.qtpl:33
	} else {
		//line template/build/render.qtpl:33
		qw422016.N().S(` <table class="table"> <thead> <tr> <th>STATUS</th> <th>BUILD</th> <th>NAMESPACE</th> <th></th> <th></th> </tr> </thead> <tbody> `)
		//line template/build/render.qtpl:45
		for _, b := range builds {
			//line template/build/render.qtpl:45
			qw422016.N().S(` <tr> <td>`)
			//line template/build/render.qtpl:47
			template.StreamRenderStatus(qw422016, b.Status)
			//line template/build/render.qtpl:47
			qw422016.N().S(`</td> <td><a href="`)
			//line template/build/render.qtpl:48
			qw422016.E().S(b.UIEndpoint())
			//line template/build/render.qtpl:48
			qw422016.N().S(`">#`)
			//line template/build/render.qtpl:48
			qw422016.E().V(b.ID)
			//line template/build/render.qtpl:48
			qw422016.N().S(`</a></td> <td> `)
			//line template/build/render.qtpl:50
			if b.Namespace != nil {
				//line template/build/render.qtpl:50
				qw422016.N().S(` <a href="`)
				//line template/build/render.qtpl:51
				qw422016.E().S(b.Namespace.UIEndpoint())
				//line template/build/render.qtpl:51
				qw422016.N().S(`">`)
				//line template/build/render.qtpl:51
				qw422016.E().S(b.Namespace.Path)
				//line template/build/render.qtpl:51
				qw422016.N().S(`</a> `)
				//line template/build/render.qtpl:52
			} else {
				//line template/build/render.qtpl:52
				qw422016.N().S(` <span class="muted">--</span> `)
				//line template/build/render.qtpl:54
			}
			//line template/build/render.qtpl:54
			qw422016.N().S(` </td> <td class="align-right"> `)
			//line template/build/render.qtpl:57
			for _, t := range b.Tags {
				//line template/build/render.qtpl:57
				qw422016.N().S(` <a class="pill pill-light" href="?tag=`)
				//line template/build/render.qtpl:58
				qw422016.E().S(t.Name)
				//line template/build/render.qtpl:58
				qw422016.N().S(`">`)
				//line template/build/render.qtpl:58
				qw422016.E().S(t.Name)
				//line template/build/render.qtpl:58
				qw422016.N().S(`</a> `)
				//line template/build/render.qtpl:59
			}
			//line template/build/render.qtpl:59
			qw422016.N().S(` </td> <td class="align-right muted"> `)
			//line template/build/render.qtpl:62
			if b.FinishedAt != nil && b.FinishedAt.Valid {
				//line template/build/render.qtpl:62
				qw422016.N().S(` `)
				//line template/build/render.qtpl:63
			} else {
				//line template/build/render.qtpl:63
				qw422016.N().S(` -- `)
				//line template/build/render.qtpl:65
			}
			//line template/build/render.qtpl:65
			qw422016.N().S(` </td> </tr> `)
			//line template/build/render.qtpl:68
		}
		//line template/build/render.qtpl:68
		qw422016.N().S(` </tbody> </table> `)
		//line template/build/render.qtpl:71
	}
	//line template/build/render.qtpl:71
	qw422016.N().S(` </div> `)
//line template/build/render.qtpl:73
}

//line template/build/render.qtpl:73
func WriteRenderTable(qq422016 qtio422016.Writer, builds []*model.Build, status, uri string) {
	//line template/build/render.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/render.qtpl:73
	StreamRenderTable(qw422016, builds, status, uri)
	//line template/build/render.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line template/build/render.qtpl:73
}

//line template/build/render.qtpl:73
func RenderTable(builds []*model.Build, status, uri string) string {
	//line template/build/render.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/render.qtpl:73
	WriteRenderTable(qb422016, builds, status, uri)
	//line template/build/render.qtpl:73
	qs422016 := string(qb422016.B)
	//line template/build/render.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/render.qtpl:73
	return qs422016
//line template/build/render.qtpl:73
}

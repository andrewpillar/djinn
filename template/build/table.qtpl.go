// This file is automatically generated by qtc from "table.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/build/table.qtpl:2
package build

//line template/build/table.qtpl:2
import (
	"fmt"
	htmltemplate "html/template"
	"strings"

	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/build/table.qtpl:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build/table.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build/table.qtpl:13
func StreamRenderStatusNav(qw422016 *qt422016.Writer, uri, status string) {
	//line template/build/table.qtpl:13
	qw422016.N().S(` `)
	//line template/build/table.qtpl:14
	qw422016.N().S(`<ul class="panel-nav"><li><a href="`)
	//line template/build/table.qtpl:17
	qw422016.E().S(uri)
	//line template/build/table.qtpl:17
	qw422016.N().S(`"`)
	//line template/build/table.qtpl:17
	if status == "" {
		//line template/build/table.qtpl:17
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:17
	}
	//line template/build/table.qtpl:17
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:18
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M22.688 18.984c0.422 0.281 0.422 0.984-0.094 1.406l-2.297 2.297c-0.422 0.422-0.984 0.422-1.406 0l-9.094-9.094c-2.297 0.891-4.969 0.422-6.891-1.5-2.016-2.016-2.531-5.016-1.313-7.406l4.406 4.313 3-3-4.313-4.313c2.391-1.078 5.391-0.703 7.406 1.313 1.922 1.922 2.391 4.594 1.5 6.891z"></path>
</svg>
`)
	//line template/build/table.qtpl:18
	qw422016.N().S(`<span>All</span></a></li><li><a href="?status=queued"`)
	//line template/build/table.qtpl:22
	if status == "queued" {
		//line template/build/table.qtpl:22
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:22
	}
	//line template/build/table.qtpl:22
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:23
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 11.484l3.984-3.984v-3.516h-7.969v3.516zM15.984 16.5l-3.984-3.984-3.984 3.984v3.516h7.969v-3.516zM6 2.016h12v6l-3.984 3.984 3.984 3.984v6h-12v-6l3.984-3.984-3.984-3.984v-6z"></path>
</svg>
`)
	//line template/build/table.qtpl:23
	qw422016.N().S(`<span>Queued</span></a></li><li><a href="?status=running"`)
	//line template/build/table.qtpl:27
	if status == "running" {
		//line template/build/table.qtpl:27
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:27
	}
	//line template/build/table.qtpl:27
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:28
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M21.984 12c0 5.156-3.938 9.422-8.953 9.938v-2.016c3.938-0.516 6.984-3.891 6.984-7.922s-3.047-7.406-6.984-7.922v-2.016c5.016 0.516 8.953 4.781 8.953 9.938zM5.672 19.734l1.406-1.406c1.125 0.844 2.484 1.406 3.938 1.594v2.016c-2.016-0.188-3.844-0.984-5.344-2.203zM4.078 12.984c0.188 1.453 0.75 2.813 1.594 3.891l-1.406 1.453c-1.219-1.5-2.016-3.328-2.203-5.344h2.016zM5.672 7.078c-0.844 1.125-1.406 2.484-1.594 3.938h-2.016c0.188-2.016 0.984-3.844 2.203-5.344zM11.016 4.078c-1.453 0.188-2.813 0.75-3.938 1.594l-1.406-1.406c1.5-1.219 3.328-2.016 5.344-2.203v2.016zM13.031 9.797l2.953 2.203c-2.007 1.493-4.007 2.993-6 4.5z"></path>
</svg>
`)
	//line template/build/table.qtpl:28
	qw422016.N().S(`<span>Running</span></a></li><li><a href="?status=passed"`)
	//line template/build/table.qtpl:32
	if status == "passed" {
		//line template/build/table.qtpl:32
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:32
	}
	//line template/build/table.qtpl:32
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:33
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
	//line template/build/table.qtpl:33
	qw422016.N().S(`<span>Passed</span></a></li><li><a href="?status=failed"`)
	//line template/build/table.qtpl:37
	if status == "failed" {
		//line template/build/table.qtpl:37
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:37
	}
	//line template/build/table.qtpl:37
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:38
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
	//line template/build/table.qtpl:38
	qw422016.N().S(`<span>Failed</span></a></li><li><a href="?status=killed"`)
	//line template/build/table.qtpl:42
	if status == "killed" {
		//line template/build/table.qtpl:42
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:42
	}
	//line template/build/table.qtpl:42
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:43
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 12.984v-6h-1.969v6h1.969zM12 17.297c0.703 0 1.313-0.609 1.313-1.313s-0.609-1.266-1.313-1.266-1.313 0.563-1.313 1.266 0.609 1.313 1.313 1.313zM15.75 3l5.25 5.25v7.5l-5.25 5.25h-7.5l-5.25-5.25v-7.5l5.25-5.25h7.5z"></path>
</svg>
`)
	//line template/build/table.qtpl:43
	qw422016.N().S(`<span>Killed</span></a></li><li><a href="?status=timed_out"`)
	//line template/build/table.qtpl:47
	if status == "timed_out" {
		//line template/build/table.qtpl:47
		qw422016.N().S(`class="active"`)
		//line template/build/table.qtpl:47
	}
	//line template/build/table.qtpl:47
	qw422016.N().S(`>`)
	//line template/build/table.qtpl:48
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 20.016c3.891 0 6.984-3.141 6.984-7.031s-3.094-6.984-6.984-6.984-6.984 3.094-6.984 6.984 3.094 7.031 6.984 7.031zM19.031 7.406c1.219 1.547 1.969 3.469 1.969 5.578 0 4.969-4.031 9-9 9s-9-4.031-9-9 4.031-9 9-9c2.109 0 4.078 0.797 5.625 2.016l1.406-1.453c0.516 0.422 0.984 0.891 1.406 1.406zM11.016 14.016v-6h1.969v6h-1.969zM15 0.984v2.016h-6v-2.016h6z"></path>
</svg>
`)
	//line template/build/table.qtpl:48
	qw422016.N().S(`<span>Timed Out</span></a></li></ul>`)
	//line template/build/table.qtpl:52
	qw422016.N().S(` `)
//line template/build/table.qtpl:53
}

//line template/build/table.qtpl:53
func WriteRenderStatusNav(qq422016 qtio422016.Writer, uri, status string) {
	//line template/build/table.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/table.qtpl:53
	StreamRenderStatusNav(qw422016, uri, status)
	//line template/build/table.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line template/build/table.qtpl:53
}

//line template/build/table.qtpl:53
func RenderStatusNav(uri, status string) string {
	//line template/build/table.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/table.qtpl:53
	WriteRenderStatusNav(qb422016, uri, status)
	//line template/build/table.qtpl:53
	qs422016 := string(qb422016.B)
	//line template/build/table.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/table.qtpl:53
	return qs422016
//line template/build/table.qtpl:53
}

//line template/build/table.qtpl:55
func StreamRenderIndex(qw422016 *qt422016.Writer, bb []*model.Build, uri, status, search string) {
	//line template/build/table.qtpl:55
	qw422016.N().S(` `)
	//line template/build/table.qtpl:56
	if len(bb) == 0 && search == "" && status == "" {
		//line template/build/table.qtpl:56
		qw422016.N().S(` <div class="panel-message muted">No builds have been submitted yet.</div> `)
		//line template/build/table.qtpl:58
	} else if len(bb) == 0 && search != "" {
		//line template/build/table.qtpl:58
		qw422016.N().S(` <div class="panel-header"> `)
		//line template/build/table.qtpl:60
		StreamRenderStatusNav(qw422016, uri, status)
		//line template/build/table.qtpl:60
		qw422016.N().S(` `)
		//line template/build/table.qtpl:61
		template.StreamRenderSearch(qw422016, uri, search, "Find a build...")
		//line template/build/table.qtpl:61
		qw422016.N().S(` </div> <div class="panel-message muted">No results found.</div> `)
		//line template/build/table.qtpl:64
	} else if len(bb) == 0 && status != "" {
		//line template/build/table.qtpl:64
		qw422016.N().S(` <div class="panel-header"> `)
		//line template/build/table.qtpl:66
		StreamRenderStatusNav(qw422016, uri, status)
		//line template/build/table.qtpl:66
		qw422016.N().S(` `)
		//line template/build/table.qtpl:67
		template.StreamRenderSearch(qw422016, uri, search, "Find a build...")
		//line template/build/table.qtpl:67
		qw422016.N().S(` </div> <div class="panel-message muted">No `)
		//line template/build/table.qtpl:69
		qw422016.E().S(strings.Replace(status, "_", " ", -1))
		//line template/build/table.qtpl:69
		qw422016.N().S(` builds.</div> `)
		//line template/build/table.qtpl:70
	} else {
		//line template/build/table.qtpl:70
		qw422016.N().S(` <div class="panel-header"> `)
		//line template/build/table.qtpl:72
		StreamRenderStatusNav(qw422016, uri, status)
		//line template/build/table.qtpl:72
		qw422016.N().S(` `)
		//line template/build/table.qtpl:73
		template.StreamRenderSearch(qw422016, uri, search, "Find a build...")
		//line template/build/table.qtpl:73
		qw422016.N().S(` </div> `)
		//line template/build/table.qtpl:75
		StreamRenderTable(qw422016, bb)
		//line template/build/table.qtpl:75
		qw422016.N().S(` `)
		//line template/build/table.qtpl:76
	}
	//line template/build/table.qtpl:76
	qw422016.N().S(` `)
//line template/build/table.qtpl:77
}

//line template/build/table.qtpl:77
func WriteRenderIndex(qq422016 qtio422016.Writer, bb []*model.Build, uri, status, search string) {
	//line template/build/table.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/table.qtpl:77
	StreamRenderIndex(qw422016, bb, uri, status, search)
	//line template/build/table.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line template/build/table.qtpl:77
}

//line template/build/table.qtpl:77
func RenderIndex(bb []*model.Build, uri, status, search string) string {
	//line template/build/table.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/table.qtpl:77
	WriteRenderIndex(qb422016, bb, uri, status, search)
	//line template/build/table.qtpl:77
	qs422016 := string(qb422016.B)
	//line template/build/table.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/table.qtpl:77
	return qs422016
//line template/build/table.qtpl:77
}

//line template/build/table.qtpl:79
func StreamRenderTable(qw422016 *qt422016.Writer, builds []*model.Build) {
	//line template/build/table.qtpl:79
	qw422016.N().S(` <table class="table"> <thead> <tr> <th>STATUS</th> <th>BUILD</th> <th>NAMESPACE</th> <th></th> <th></th> </tr> </thead> <tbody> `)
	//line template/build/table.qtpl:91
	for _, b := range builds {
		//line template/build/table.qtpl:91
		qw422016.N().S(` <tr> <td>`)
		//line template/build/table.qtpl:93
		template.StreamRenderStatus(qw422016, b.Status)
		//line template/build/table.qtpl:93
		qw422016.N().S(`</td> <td><a href="`)
		//line template/build/table.qtpl:94
		qw422016.E().S(b.UIEndpoint())
		//line template/build/table.qtpl:94
		qw422016.N().S(`">#`)
		//line template/build/table.qtpl:94
		qw422016.E().V(b.ID)
		//line template/build/table.qtpl:94
		qw422016.N().S(`</a></td> <td> `)
		//line template/build/table.qtpl:96
		if b.Namespace != nil {
			//line template/build/table.qtpl:96
			qw422016.N().S(` <a href="`)
			//line template/build/table.qtpl:97
			qw422016.E().S(b.Namespace.UIEndpoint())
			//line template/build/table.qtpl:97
			qw422016.N().S(`">`)
			//line template/build/table.qtpl:97
			qw422016.E().S(b.Namespace.Path)
			//line template/build/table.qtpl:97
			qw422016.N().S(`</a> `)
			//line template/build/table.qtpl:98
		} else {
			//line template/build/table.qtpl:98
			qw422016.N().S(` <span class="muted">--</span> `)
			//line template/build/table.qtpl:100
		}
		//line template/build/table.qtpl:100
		qw422016.N().S(` </td> <td class="align-right"> `)
		//line template/build/table.qtpl:103
		for _, t := range b.Tags {
			//line template/build/table.qtpl:103
			qw422016.N().S(` <a class="pill pill-light" href="?tag=`)
			//line template/build/table.qtpl:104
			qw422016.E().S(t.Name)
			//line template/build/table.qtpl:104
			qw422016.N().S(`">`)
			//line template/build/table.qtpl:104
			qw422016.E().S(t.Name)
			//line template/build/table.qtpl:104
			qw422016.N().S(`</a> `)
			//line template/build/table.qtpl:105
		}
		//line template/build/table.qtpl:105
		qw422016.N().S(` </td> <td class="align-right"> `)
		//line template/build/table.qtpl:108
		if !b.FinishedAt.Valid || !b.StartedAt.Valid {
			//line template/build/table.qtpl:108
			qw422016.N().S(` <span class="muted">--</span> `)
			//line template/build/table.qtpl:110
		} else {
			//line template/build/table.qtpl:110
			qw422016.N().S(` `)
			//line template/build/table.qtpl:111
			qw422016.E().V(b.FinishedAt.Time.Sub(b.StartedAt.Time))
			//line template/build/table.qtpl:111
			qw422016.N().S(` `)
			//line template/build/table.qtpl:112
		}
		//line template/build/table.qtpl:112
		qw422016.N().S(` </td> </tr> `)
		//line template/build/table.qtpl:115
	}
	//line template/build/table.qtpl:115
	qw422016.N().S(` </tbody> </table> `)
//line template/build/table.qtpl:118
}

//line template/build/table.qtpl:118
func WriteRenderTable(qq422016 qtio422016.Writer, builds []*model.Build) {
	//line template/build/table.qtpl:118
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/table.qtpl:118
	StreamRenderTable(qw422016, builds)
	//line template/build/table.qtpl:118
	qt422016.ReleaseWriter(qw422016)
//line template/build/table.qtpl:118
}

//line template/build/table.qtpl:118
func RenderTable(builds []*model.Build) string {
	//line template/build/table.qtpl:118
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/table.qtpl:118
	WriteRenderTable(qb422016, builds)
	//line template/build/table.qtpl:118
	qs422016 := string(qb422016.B)
	//line template/build/table.qtpl:118
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/table.qtpl:118
	return qs422016
//line template/build/table.qtpl:118
}

//line template/build/table.qtpl:120
func StreamRenderObjectTable(qw422016 *qt422016.Writer, oo []*model.BuildObject) {
	//line template/build/table.qtpl:120
	qw422016.N().S(` <table class="table"> <thead> <tr> <th>SOURCE</th> <th>NAME</th> <th>TYPE</th> <th>MD5</th> <th>SHA256</th> <th></th> </tr> </thead> <tbody> `)
	//line template/build/table.qtpl:133
	for _, o := range oo {
		//line template/build/table.qtpl:133
		qw422016.N().S(` <tr> <td> `)
		//line template/build/table.qtpl:136
		if o.Object != nil && o.Object.DeletedAt == nil {
			//line template/build/table.qtpl:136
			qw422016.N().S(` <a href="`)
			//line template/build/table.qtpl:137
			qw422016.E().S(o.Object.UIEndpoint())
			//line template/build/table.qtpl:137
			qw422016.N().S(`">`)
			//line template/build/table.qtpl:137
			qw422016.E().S(o.Source)
			//line template/build/table.qtpl:137
			qw422016.N().S(`</a> `)
			//line template/build/table.qtpl:138
		} else {
			//line template/build/table.qtpl:138
			qw422016.N().S(` <a title="Object not found"><strike>`)
			//line template/build/table.qtpl:139
			qw422016.E().S(o.Source)
			//line template/build/table.qtpl:139
			qw422016.N().S(`</strike></a> `)
			//line template/build/table.qtpl:140
		}
		//line template/build/table.qtpl:140
		qw422016.N().S(` </td> <td><span class="code">`)
		//line template/build/table.qtpl:142
		qw422016.E().S(o.Name)
		//line template/build/table.qtpl:142
		qw422016.N().S(`</span></td> <td> `)
		//line template/build/table.qtpl:144
		if o.Object != nil {
			//line template/build/table.qtpl:144
			qw422016.N().S(` <span class="code">`)
			//line template/build/table.qtpl:145
			qw422016.E().S(o.Object.Type)
			//line template/build/table.qtpl:145
			qw422016.N().S(`</span> `)
			//line template/build/table.qtpl:146
		} else {
			//line template/build/table.qtpl:146
			qw422016.N().S(` <span class="code">--</span> `)
			//line template/build/table.qtpl:148
		}
		//line template/build/table.qtpl:148
		qw422016.N().S(` </td> <td> `)
		//line template/build/table.qtpl:151
		if o.Object != nil {
			//line template/build/table.qtpl:151
			qw422016.N().S(` <span class="code">`)
			//line template/build/table.qtpl:152
			qw422016.E().S(fmt.Sprintf("%x", o.Object.MD5))
			//line template/build/table.qtpl:152
			qw422016.N().S(`</span> `)
			//line template/build/table.qtpl:153
		} else {
			//line template/build/table.qtpl:153
			qw422016.N().S(` <span class="code">--</span> `)
			//line template/build/table.qtpl:155
		}
		//line template/build/table.qtpl:155
		qw422016.N().S(` </td> <td> `)
		//line template/build/table.qtpl:158
		if o.Object != nil {
			//line template/build/table.qtpl:158
			qw422016.N().S(` <span class="code">`)
			//line template/build/table.qtpl:159
			qw422016.E().S(fmt.Sprintf("%x", o.Object.SHA256))
			//line template/build/table.qtpl:159
			qw422016.N().S(`</span> `)
			//line template/build/table.qtpl:160
		} else {
			//line template/build/table.qtpl:160
			qw422016.N().S(` <span class="code">--</span> `)
			//line template/build/table.qtpl:162
		}
		//line template/build/table.qtpl:162
		qw422016.N().S(` </td> <td class="align-right"> `)
		//line template/build/table.qtpl:165
		if o.Placed {
			//line template/build/table.qtpl:165
			qw422016.N().S(` <span class="pill pill-green">`)
			//line template/build/table.qtpl:166
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
			//line template/build/table.qtpl:166
			qw422016.N().S(` Placed</span> `)
			//line template/build/table.qtpl:167
		} else {
			//line template/build/table.qtpl:167
			qw422016.N().S(` <span class="pill pill-red">`)
			//line template/build/table.qtpl:168
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
			//line template/build/table.qtpl:168
			qw422016.N().S(` Not Placed</span> `)
			//line template/build/table.qtpl:169
		}
		//line template/build/table.qtpl:169
		qw422016.N().S(` </td> </tr> `)
		//line template/build/table.qtpl:172
	}
	//line template/build/table.qtpl:172
	qw422016.N().S(` </tbody> </table> `)
//line template/build/table.qtpl:175
}

//line template/build/table.qtpl:175
func WriteRenderObjectTable(qq422016 qtio422016.Writer, oo []*model.BuildObject) {
	//line template/build/table.qtpl:175
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/table.qtpl:175
	StreamRenderObjectTable(qw422016, oo)
	//line template/build/table.qtpl:175
	qt422016.ReleaseWriter(qw422016)
//line template/build/table.qtpl:175
}

//line template/build/table.qtpl:175
func RenderObjectTable(oo []*model.BuildObject) string {
	//line template/build/table.qtpl:175
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/table.qtpl:175
	WriteRenderObjectTable(qb422016, oo)
	//line template/build/table.qtpl:175
	qs422016 := string(qb422016.B)
	//line template/build/table.qtpl:175
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/table.qtpl:175
	return qs422016
//line template/build/table.qtpl:175
}

//line template/build/table.qtpl:177
func StreamRenderTagTable(qw422016 *qt422016.Writer, tt []*model.Tag, csrf htmltemplate.HTML) {
	//line template/build/table.qtpl:177
	qw422016.N().S(` <table class="table"> <thead> <tr> <th>TAG</th> <th>TAGGED BY</th> <th></th> </tr> </thead> <tbody> `)
	//line template/build/table.qtpl:187
	for _, t := range tt {
		//line template/build/table.qtpl:187
		qw422016.N().S(` <tr> <td><a href="/?tag=`)
		//line template/build/table.qtpl:189
		qw422016.E().S(t.Name)
		//line template/build/table.qtpl:189
		qw422016.N().S(`" class="pill pill-light">`)
		//line template/build/table.qtpl:189
		qw422016.E().S(t.Name)
		//line template/build/table.qtpl:189
		qw422016.N().S(`</a></td> <td>`)
		//line template/build/table.qtpl:190
		qw422016.E().S(t.User.Username)
		//line template/build/table.qtpl:190
		qw422016.N().S(` &lt;`)
		//line template/build/table.qtpl:190
		qw422016.E().S(t.User.Email)
		//line template/build/table.qtpl:190
		qw422016.N().S(`&gt;</td> <td class="align-right"> <form method="POST" action="`)
		//line template/build/table.qtpl:192
		qw422016.E().S(t.UIEndpoint())
		//line template/build/table.qtpl:192
		qw422016.N().S(`"> `)
		//line template/build/table.qtpl:193
		qw422016.N().S(string(csrf))
		//line template/build/table.qtpl:193
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"> <button type="submit" class="btn btn-danger">Delete</button> </form> </td> </tr> `)
		//line template/build/table.qtpl:199
	}
	//line template/build/table.qtpl:199
	qw422016.N().S(` </tbody> </table> `)
//line template/build/table.qtpl:202
}

//line template/build/table.qtpl:202
func WriteRenderTagTable(qq422016 qtio422016.Writer, tt []*model.Tag, csrf htmltemplate.HTML) {
	//line template/build/table.qtpl:202
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/table.qtpl:202
	StreamRenderTagTable(qw422016, tt, csrf)
	//line template/build/table.qtpl:202
	qt422016.ReleaseWriter(qw422016)
//line template/build/table.qtpl:202
}

//line template/build/table.qtpl:202
func RenderTagTable(tt []*model.Tag, csrf htmltemplate.HTML) string {
	//line template/build/table.qtpl:202
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/table.qtpl:202
	WriteRenderTagTable(qb422016, tt, csrf)
	//line template/build/table.qtpl:202
	qs422016 := string(qb422016.B)
	//line template/build/table.qtpl:202
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/table.qtpl:202
	return qs422016
//line template/build/table.qtpl:202
}

//line template/build/table.qtpl:204
func StreamRenderVariableTable(qw422016 *qt422016.Writer, vv []*model.BuildVariable) {
	//line template/build/table.qtpl:204
	qw422016.N().S(` <table class="table"> <thead> <tr> <th>KEY</th> <th>VALUE</th> <th>FROM MANIFEST</th> </tr> </thead> <tbody> `)
	//line template/build/table.qtpl:214
	for _, v := range vv {
		//line template/build/table.qtpl:214
		qw422016.N().S(` <tr> <td><span class="code">`)
		//line template/build/table.qtpl:216
		qw422016.E().S(v.Key)
		//line template/build/table.qtpl:216
		qw422016.N().S(`</span></td> <td><span class="code">`)
		//line template/build/table.qtpl:217
		qw422016.E().S(v.Value)
		//line template/build/table.qtpl:217
		qw422016.N().S(`</span></td> <td>`)
		//line template/build/table.qtpl:218
		if v.VariableID.Valid {
			//line template/build/table.qtpl:218
			qw422016.N().S(`False`)
			//line template/build/table.qtpl:218
		} else {
			//line template/build/table.qtpl:218
			qw422016.N().S(`True`)
			//line template/build/table.qtpl:218
		}
		//line template/build/table.qtpl:218
		qw422016.N().S(`</td> </tr> `)
		//line template/build/table.qtpl:220
	}
	//line template/build/table.qtpl:220
	qw422016.N().S(` </tbody> </table> `)
//line template/build/table.qtpl:223
}

//line template/build/table.qtpl:223
func WriteRenderVariableTable(qq422016 qtio422016.Writer, vv []*model.BuildVariable) {
	//line template/build/table.qtpl:223
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/build/table.qtpl:223
	StreamRenderVariableTable(qw422016, vv)
	//line template/build/table.qtpl:223
	qt422016.ReleaseWriter(qw422016)
//line template/build/table.qtpl:223
}

//line template/build/table.qtpl:223
func RenderVariableTable(vv []*model.BuildVariable) string {
	//line template/build/table.qtpl:223
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/build/table.qtpl:223
	WriteRenderVariableTable(qb422016, vv)
	//line template/build/table.qtpl:223
	qs422016 := string(qb422016.B)
	//line template/build/table.qtpl:223
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/build/table.qtpl:223
	return qs422016
//line template/build/table.qtpl:223
}

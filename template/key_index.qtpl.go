// Code generated by qtc from "key_index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/key_index.qtpl:2
package template

//line template/key_index.qtpl:2
import (
	"djinn-ci.com/key"
	"djinn-ci.com/template/form"
)

//line template/key_index.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/key_index.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/key_index.qtpl:9
type KeyIndex struct {
	*Paginator

	Keys []*key.Key
}

//line template/key_index.qtpl:17
func (p *KeyIndex) StreamTitle(qw422016 *qt422016.Writer) {
//line template/key_index.qtpl:17
	qw422016.N().S(`SSH Keys`)
//line template/key_index.qtpl:17
}

//line template/key_index.qtpl:17
func (p *KeyIndex) WriteTitle(qq422016 qtio422016.Writer) {
//line template/key_index.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:17
	p.StreamTitle(qw422016)
//line template/key_index.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:17
}

//line template/key_index.qtpl:17
func (p *KeyIndex) Title() string {
//line template/key_index.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:17
	p.WriteTitle(qb422016)
//line template/key_index.qtpl:17
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:17
	return qs422016
//line template/key_index.qtpl:17
}

//line template/key_index.qtpl:18
func (p *KeyIndex) StreamHeader(qw422016 *qt422016.Writer) {
//line template/key_index.qtpl:18
	p.StreamTitle(qw422016)
//line template/key_index.qtpl:18
}

//line template/key_index.qtpl:18
func (p *KeyIndex) WriteHeader(qq422016 qtio422016.Writer) {
//line template/key_index.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:18
	p.StreamHeader(qw422016)
//line template/key_index.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:18
}

//line template/key_index.qtpl:18
func (p *KeyIndex) Header() string {
//line template/key_index.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:18
	p.WriteHeader(qb422016)
//line template/key_index.qtpl:18
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:18
	return qs422016
//line template/key_index.qtpl:18
}

//line template/key_index.qtpl:20
func (p *KeyIndex) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/key_index.qtpl:20
}

//line template/key_index.qtpl:20
func (p *KeyIndex) WriteNavigation(qq422016 qtio422016.Writer) {
//line template/key_index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:20
	p.StreamNavigation(qw422016)
//line template/key_index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:20
}

//line template/key_index.qtpl:20
func (p *KeyIndex) Navigation() string {
//line template/key_index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:20
	p.WriteNavigation(qb422016)
//line template/key_index.qtpl:20
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:20
	return qs422016
//line template/key_index.qtpl:20
}

//line template/key_index.qtpl:21
func (p *KeyIndex) StreamFooter(qw422016 *qt422016.Writer) {
//line template/key_index.qtpl:21
}

//line template/key_index.qtpl:21
func (p *KeyIndex) WriteFooter(qq422016 qtio422016.Writer) {
//line template/key_index.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:21
	p.StreamFooter(qw422016)
//line template/key_index.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:21
}

//line template/key_index.qtpl:21
func (p *KeyIndex) Footer() string {
//line template/key_index.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:21
	p.WriteFooter(qb422016)
//line template/key_index.qtpl:21
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:21
	return qs422016
//line template/key_index.qtpl:21
}

//line template/key_index.qtpl:23
func (p *KeyIndex) StreamActions(qw422016 *qt422016.Writer) {
//line template/key_index.qtpl:23
	qw422016.N().S(` `)
//line template/key_index.qtpl:24
	if _, ok := p.User.Permissions["key:write"]; ok {
//line template/key_index.qtpl:24
		qw422016.N().S(` <li><a href="/keys/create" class="btn btn-primary">Create</a></li> `)
//line template/key_index.qtpl:26
	}
//line template/key_index.qtpl:26
	qw422016.N().S(` `)
//line template/key_index.qtpl:27
}

//line template/key_index.qtpl:27
func (p *KeyIndex) WriteActions(qq422016 qtio422016.Writer) {
//line template/key_index.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:27
	p.StreamActions(qw422016)
//line template/key_index.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:27
}

//line template/key_index.qtpl:27
func (p *KeyIndex) Actions() string {
//line template/key_index.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:27
	p.WriteActions(qb422016)
//line template/key_index.qtpl:27
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:27
	return qs422016
//line template/key_index.qtpl:27
}

//line template/key_index.qtpl:29
func (p *KeyIndex) streamrenderKeyItem(qw422016 *qt422016.Writer, k *key.Key) {
//line template/key_index.qtpl:29
	qw422016.N().S(` <tr> <td><a href="`)
//line template/key_index.qtpl:31
	qw422016.E().S(k.Endpoint("edit"))
//line template/key_index.qtpl:31
	qw422016.N().S(`">`)
//line template/key_index.qtpl:31
	qw422016.E().S(k.Name)
//line template/key_index.qtpl:31
	qw422016.N().S(`</a></td> <td> `)
//line template/key_index.qtpl:33
	if k.Namespace != nil {
//line template/key_index.qtpl:33
		qw422016.N().S(` <a href="`)
//line template/key_index.qtpl:34
		qw422016.E().S(k.Namespace.Endpoint())
//line template/key_index.qtpl:34
		qw422016.N().S(`">`)
//line template/key_index.qtpl:34
		qw422016.E().S(k.Namespace.Path)
//line template/key_index.qtpl:34
		qw422016.N().S(`</a> `)
//line template/key_index.qtpl:35
	} else {
//line template/key_index.qtpl:35
		qw422016.N().S(` <span class="muted">--</span> `)
//line template/key_index.qtpl:37
	}
//line template/key_index.qtpl:37
	qw422016.N().S(` </td> <td class="align-right"> `)
//line template/key_index.qtpl:40
	if p.User.ID != k.UserID {
//line template/key_index.qtpl:40
		qw422016.N().S(` <span class="muted">`)
//line template/key_index.qtpl:41
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.984 12.984c2.344 0 7.031 1.172 7.031 3.516v2.484h-6v-2.484c0-1.5-0.797-2.625-1.969-3.469 0.328-0.047 0.656-0.047 0.938-0.047zM8.016 12.984c2.344 0 6.984 1.172 6.984 3.516v2.484h-14.016v-2.484c0-2.344 4.688-3.516 7.031-3.516zM8.016 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 2.953 1.359 2.953 3-1.313 3-2.953 3zM15.984 11.016c-1.641 0-3-1.359-3-3s1.359-3 3-3 3 1.359 3 3-1.359 3-3 3z"></path>
</svg>
`)
//line template/key_index.qtpl:41
		qw422016.N().S(`</span> `)
//line template/key_index.qtpl:42
	}
//line template/key_index.qtpl:42
	qw422016.N().S(` </td> <td class="align-right"> `)
//line template/key_index.qtpl:45
	if p.User.ID == k.UserID || k.Namespace != nil && k.Namespace.UserID == p.User.ID {
//line template/key_index.qtpl:45
		qw422016.N().S(` <form method="POST" action="`)
//line template/key_index.qtpl:46
		qw422016.E().S(k.Endpoint())
//line template/key_index.qtpl:46
		qw422016.N().S(`"> `)
//line template/key_index.qtpl:47
		form.StreamMethod(qw422016, "DELETE")
//line template/key_index.qtpl:47
		qw422016.N().S(` `)
//line template/key_index.qtpl:48
		qw422016.N().V(p.CSRF)
//line template/key_index.qtpl:48
		qw422016.N().S(` <button type="submit" class="btn btn-danger">Delete</button> </form> `)
//line template/key_index.qtpl:51
	}
//line template/key_index.qtpl:51
	qw422016.N().S(` </td> </tr> `)
//line template/key_index.qtpl:54
}

//line template/key_index.qtpl:54
func (p *KeyIndex) writerenderKeyItem(qq422016 qtio422016.Writer, k *key.Key) {
//line template/key_index.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:54
	p.streamrenderKeyItem(qw422016, k)
//line template/key_index.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:54
}

//line template/key_index.qtpl:54
func (p *KeyIndex) renderKeyItem(k *key.Key) string {
//line template/key_index.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:54
	p.writerenderKeyItem(qb422016, k)
//line template/key_index.qtpl:54
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:54
	return qs422016
//line template/key_index.qtpl:54
}

//line template/key_index.qtpl:56
func (p *KeyIndex) StreamBody(qw422016 *qt422016.Writer) {
//line template/key_index.qtpl:56
	qw422016.N().S(` <div class="panel"> `)
//line template/key_index.qtpl:58
	if len(p.Keys) == 0 {
//line template/key_index.qtpl:58
		qw422016.N().S(` `)
//line template/key_index.qtpl:59
		if query := p.Query.Get("search"); query != "" {
//line template/key_index.qtpl:59
			qw422016.N().S(` <div class="panel-header">`)
//line template/key_index.qtpl:60
			p.StreamSearch(qw422016, "Find an SSH key...")
//line template/key_index.qtpl:60
			qw422016.N().S(`</div> <div class="panel-message muted">No results found.</div> `)
//line template/key_index.qtpl:62
		} else {
//line template/key_index.qtpl:62
			qw422016.N().S(` <div class="panel-message muted"> SSH keys can allow build environments to access other environments. </div> `)
//line template/key_index.qtpl:66
		}
//line template/key_index.qtpl:66
		qw422016.N().S(` `)
//line template/key_index.qtpl:67
	} else {
//line template/key_index.qtpl:67
		qw422016.N().S(` <table class="table"> <thead> <tr> <th>NAME</th> <th>NAMESPACE</th> <th></th> <th></th> </tr> </thead> <tbody> `)
//line template/key_index.qtpl:78
		for _, k := range p.Keys {
//line template/key_index.qtpl:78
			qw422016.N().S(` `)
//line template/key_index.qtpl:79
			p.streamrenderKeyItem(qw422016, k)
//line template/key_index.qtpl:79
			qw422016.N().S(` `)
//line template/key_index.qtpl:80
		}
//line template/key_index.qtpl:80
		qw422016.N().S(` </tbody> </table> `)
//line template/key_index.qtpl:83
	}
//line template/key_index.qtpl:83
	qw422016.N().S(` </div> `)
//line template/key_index.qtpl:85
	p.Paginator.StreamNavigation(qw422016)
//line template/key_index.qtpl:85
	qw422016.N().S(` `)
//line template/key_index.qtpl:86
}

//line template/key_index.qtpl:86
func (p *KeyIndex) WriteBody(qq422016 qtio422016.Writer) {
//line template/key_index.qtpl:86
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/key_index.qtpl:86
	p.StreamBody(qw422016)
//line template/key_index.qtpl:86
	qt422016.ReleaseWriter(qw422016)
//line template/key_index.qtpl:86
}

//line template/key_index.qtpl:86
func (p *KeyIndex) Body() string {
//line template/key_index.qtpl:86
	qb422016 := qt422016.AcquireByteBuffer()
//line template/key_index.qtpl:86
	p.WriteBody(qb422016)
//line template/key_index.qtpl:86
	qs422016 := string(qb422016.B)
//line template/key_index.qtpl:86
	qt422016.ReleaseByteBuffer(qb422016)
//line template/key_index.qtpl:86
	return qs422016
//line template/key_index.qtpl:86
}
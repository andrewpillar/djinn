// Code generated by qtc from "form.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line namespace/template/form.qtpl:2
package template

//line namespace/template/form.qtpl:2
import (
	"github.com/andrewpillar/thrall/namespace"
	"github.com/andrewpillar/thrall/template"
)

//line namespace/template/form.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line namespace/template/form.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line namespace/template/form.qtpl:9
type Form struct {
	template.BasePage
	template.Form

	Parent    *namespace.Namespace
	Namespace *namespace.Namespace
}

func (p *Form) action() string {
	if p.Namespace == nil {
		return "/namespaces"
	}
	return p.Namespace.Endpoint()
}

func (p *Form) Field(field string) string {
	old := p.Form.Fields[field]

	if p.Namespace != nil {
		if old != "" {
			return old
		}

		switch field {
		case "name":
			return p.Namespace.Name
		case "description":
			return p.Namespace.Description
		default:
			return ""
		}
	}
	return old
}

func (p *Form) checked(vis namespace.Visibility) string {
	if p.Namespace != nil {
		if p.Namespace.Visibility == vis {
			return `checked="true"`
		}
		return ""
	}

	if p.Parent != nil {
		if p.Parent.Visibility == vis {
			return `checked="true"`
		}
	}

	if vis == namespace.Private {
		return `checked="true"`
	}
	return ""
}

func (p *Form) disabled(vis namespace.Visibility) string {
	if p.Namespace != nil {
		if p.Namespace.ParentID.Valid && p.Namespace.Visibility != vis {
			return "disabled"
		}
		return ""
	}

	if p.Parent != nil {
		if p.Parent.Visibility != vis {
			return "disabled"
		}
	}
	return ""
}

//line namespace/template/form.qtpl:82
func (p *Form) StreamBody(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:82
	qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> <form method="POST" action="`)
//line namespace/template/form.qtpl:85
	qw422016.E().S(p.action())
//line namespace/template/form.qtpl:85
	qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:86
	qw422016.N().S(string(p.CSRF))
//line namespace/template/form.qtpl:86
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:87
	if p.Namespace != nil {
//line namespace/template/form.qtpl:87
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
//line namespace/template/form.qtpl:89
	}
//line namespace/template/form.qtpl:89
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:90
	if p.Parent != nil && !p.Parent.IsZero() {
//line namespace/template/form.qtpl:90
		qw422016.N().S(` <input type="hidden" name="parent" value="`)
//line namespace/template/form.qtpl:91
		qw422016.E().S(p.Parent.Path)
//line namespace/template/form.qtpl:91
		qw422016.N().S(`"/> `)
//line namespace/template/form.qtpl:92
	}
//line namespace/template/form.qtpl:92
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input class="form-text" type="text" name="name" value="`)
//line namespace/template/form.qtpl:95
	qw422016.E().S(p.Field("name"))
//line namespace/template/form.qtpl:95
	qw422016.N().S(`" autocomplete="off"/> `)
//line namespace/template/form.qtpl:96
	p.StreamError(qw422016, "name")
//line namespace/template/form.qtpl:96
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="name">Description</label> <input class="form-text" type="text" name="description" value="`)
//line namespace/template/form.qtpl:100
	qw422016.E().S(p.Field("description"))
//line namespace/template/form.qtpl:100
	qw422016.N().S(`" autocomplete="off"/> `)
//line namespace/template/form.qtpl:101
	p.StreamError(qw422016, "description")
//line namespace/template/form.qtpl:101
	qw422016.N().S(` </div> <div class="form-field"> <label class="form-option `)
//line namespace/template/form.qtpl:104
	qw422016.E().S(p.disabled(namespace.Private))
//line namespace/template/form.qtpl:104
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="private" `)
//line namespace/template/form.qtpl:105
	qw422016.E().S(p.checked(namespace.Private))
//line namespace/template/form.qtpl:105
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:105
	qw422016.E().S(p.disabled(namespace.Private))
//line namespace/template/form.qtpl:105
	qw422016.N().S(`/> <strong>Private</strong> `)
//line namespace/template/form.qtpl:107
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.094 8.016v-2.016c0-1.688-1.406-3.094-3.094-3.094s-3.094 1.406-3.094 3.094v2.016h6.188zM12 17.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-2.016 0.938-2.016 2.016 0.938 2.016 2.016 2.016zM18 8.016c1.078 0 2.016 0.891 2.016 1.969v10.031c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h0.984v-2.016c0-2.766 2.25-5.016 5.016-5.016s5.016 2.25 5.016 5.016v2.016h0.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:107
	qw422016.N().S(` <div class="form-desc">You choose who can view builds in the namespace.</div> </label> <label class="form-option `)
//line namespace/template/form.qtpl:110
	qw422016.E().S(p.disabled(namespace.Internal))
//line namespace/template/form.qtpl:110
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="internal" `)
//line namespace/template/form.qtpl:111
	qw422016.E().S(p.checked(namespace.Internal))
//line namespace/template/form.qtpl:111
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:111
	qw422016.E().S(p.disabled(namespace.Internal))
//line namespace/template/form.qtpl:111
	qw422016.N().S(`/> <strong>Internal</strong> `)
//line namespace/template/form.qtpl:113
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.984l9 4.031v6c0 5.531-3.844 10.734-9 12-5.156-1.266-9-6.469-9-12v-6zM12 12v8.953c3.703-1.172 6.469-4.828 6.984-8.953h-6.984zM12 12v-8.813l-6.984 3.094v5.719h6.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:113
	qw422016.N().S(` <div class="form-desc">Anyone with an account will be able to view builds in the namespace.</div> </label> <label class="form-option `)
//line namespace/template/form.qtpl:116
	qw422016.E().S(p.disabled(namespace.Public))
//line namespace/template/form.qtpl:116
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="public" `)
//line namespace/template/form.qtpl:117
	qw422016.E().S(p.checked(namespace.Public))
//line namespace/template/form.qtpl:117
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:117
	qw422016.E().S(p.disabled(namespace.Public))
//line namespace/template/form.qtpl:117
	qw422016.N().S(`/> <strong>Public</strong> `)
//line namespace/template/form.qtpl:119
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.906 17.391c1.313-1.406 2.109-3.328 2.109-5.391 0-3.328-2.063-6.234-5.016-7.406v0.422c0 1.078-0.938 1.969-2.016 1.969h-1.969v2.016c0 0.563-0.469 0.984-1.031 0.984h-1.969v2.016h6c0.563 0 0.984 0.422 0.984 0.984v3h0.984c0.891 0 1.641 0.609 1.922 1.406zM11.016 19.922v-1.922c-1.078 0-2.016-0.938-2.016-2.016v-0.984l-4.781-4.781c-0.141 0.563-0.234 1.172-0.234 1.781 0 4.078 3.094 7.453 7.031 7.922zM12 2.016c5.531 0 9.984 4.453 9.984 9.984s-4.453 9.984-9.984 9.984-9.984-4.453-9.984-9.984 4.453-9.984 9.984-9.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:119
	qw422016.N().S(` <div class="form-desc">Anyone will be able to view builds in the namespace.</div> </label> </div> <div class="form-field"> `)
//line namespace/template/form.qtpl:124
	if p.Namespace != nil {
//line namespace/template/form.qtpl:124
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line namespace/template/form.qtpl:126
	} else {
//line namespace/template/form.qtpl:126
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line namespace/template/form.qtpl:128
	}
//line namespace/template/form.qtpl:128
	qw422016.N().S(` </div> </form> `)
//line namespace/template/form.qtpl:131
	if p.Namespace != nil {
//line namespace/template/form.qtpl:131
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="`)
//line namespace/template/form.qtpl:133
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:133
		qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:134
		qw422016.N().S(string(p.CSRF))
//line namespace/template/form.qtpl:134
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <div class="overflow"> <div class="right"> <button type="submit" class="btn btn-danger">Delete</button> </div> <strong>Delete Namespace</strong><br/><p>Builds within the namespace will not be deleted.</p> </div> </form> `)
//line namespace/template/form.qtpl:143
	}
//line namespace/template/form.qtpl:143
	qw422016.N().S(` </div> </div> `)
//line namespace/template/form.qtpl:146
}

//line namespace/template/form.qtpl:146
func (p *Form) WriteBody(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:146
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:146
	p.StreamBody(qw422016)
//line namespace/template/form.qtpl:146
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:146
}

//line namespace/template/form.qtpl:146
func (p *Form) Body() string {
//line namespace/template/form.qtpl:146
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:146
	p.WriteBody(qb422016)
//line namespace/template/form.qtpl:146
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:146
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:146
	return qs422016
//line namespace/template/form.qtpl:146
}

//line namespace/template/form.qtpl:148
func (p *Form) StreamHeader(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:148
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:149
	if p.Namespace != nil {
//line namespace/template/form.qtpl:149
		qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:150
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:150
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:150
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:150
		qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:151
		streamrenderPath(qw422016, p.Namespace.User.Username, p.Namespace.Path)
//line namespace/template/form.qtpl:151
		qw422016.N().S(` - Edit `)
//line namespace/template/form.qtpl:152
	} else {
//line namespace/template/form.qtpl:152
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:153
		if p.Parent != nil && !p.Parent.IsZero() {
//line namespace/template/form.qtpl:153
			qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:154
			qw422016.E().S(p.Parent.Endpoint())
//line namespace/template/form.qtpl:154
			qw422016.N().S(`">`)
//line namespace/template/form.qtpl:154
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:154
			qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:155
			streamrenderPath(qw422016, p.Parent.User.Username, p.Parent.Path)
//line namespace/template/form.qtpl:155
			qw422016.N().S(` - Create Sub-namespace `)
//line namespace/template/form.qtpl:156
		} else {
//line namespace/template/form.qtpl:156
			qw422016.N().S(` <a class="back" href="/namespaces">`)
//line namespace/template/form.qtpl:157
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:157
			qw422016.N().S(`</a> Create Namespace `)
//line namespace/template/form.qtpl:158
		}
//line namespace/template/form.qtpl:158
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:159
	}
//line namespace/template/form.qtpl:159
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:160
}

//line namespace/template/form.qtpl:160
func (p *Form) WriteHeader(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:160
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:160
	p.StreamHeader(qw422016)
//line namespace/template/form.qtpl:160
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:160
}

//line namespace/template/form.qtpl:160
func (p *Form) Header() string {
//line namespace/template/form.qtpl:160
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:160
	p.WriteHeader(qb422016)
//line namespace/template/form.qtpl:160
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:160
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:160
	return qs422016
//line namespace/template/form.qtpl:160
}

//line namespace/template/form.qtpl:162
func (p *Form) StreamActions(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:162
}

//line namespace/template/form.qtpl:162
func (p *Form) WriteActions(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:162
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:162
	p.StreamActions(qw422016)
//line namespace/template/form.qtpl:162
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:162
}

//line namespace/template/form.qtpl:162
func (p *Form) Actions() string {
//line namespace/template/form.qtpl:162
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:162
	p.WriteActions(qb422016)
//line namespace/template/form.qtpl:162
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:162
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:162
	return qs422016
//line namespace/template/form.qtpl:162
}

//line namespace/template/form.qtpl:163
func (p *Form) StreamNavigation(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:163
}

//line namespace/template/form.qtpl:163
func (p *Form) WriteNavigation(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:163
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:163
	p.StreamNavigation(qw422016)
//line namespace/template/form.qtpl:163
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:163
}

//line namespace/template/form.qtpl:163
func (p *Form) Navigation() string {
//line namespace/template/form.qtpl:163
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:163
	p.WriteNavigation(qb422016)
//line namespace/template/form.qtpl:163
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:163
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:163
	return qs422016
//line namespace/template/form.qtpl:163
}
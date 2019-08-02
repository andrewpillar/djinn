// This file is automatically generated by qtc from "form.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/namespace/form.qtpl:2
package namespace

//line template/namespace/form.qtpl:2
import (
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/namespace/form.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/namespace/form.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/namespace/form.qtpl:9
type Form struct {
	template.Page
	template.Form

	Parent    *model.Namespace
	Namespace *model.Namespace
}

func (p *Form) action() string {
	if p.Namespace == nil {
		return "/namespaces"
	}

	return p.Namespace.UIEndpoint()
}

func (p *Form) Field(field string) string {
	old := p.Form.Field(field)

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

func (p *Form) checked(vis model.Visibility) string {
	if p.Namespace != nil {
		if p.Namespace.Visibility == vis {
			return `checked="true"`
		}

		return ``
	}

	if p.Parent != nil {
		if p.Parent.Visibility == vis {
			return `checked="true"`
		}
	}

	if vis == model.Private {
		return `checked="true"`
	}

	return ``
}

func (p *Form) disabled(vis model.Visibility) string {
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

//line template/namespace/form.qtpl:88
func (p *Form) StreamBody(qw422016 *qt422016.Writer) {
	//line template/namespace/form.qtpl:88
	qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> <form method="POST" action="`)
	//line template/namespace/form.qtpl:91
	qw422016.E().S(p.action())
	//line template/namespace/form.qtpl:91
	qw422016.N().S(`"> `)
	//line template/namespace/form.qtpl:92
	qw422016.N().S(string(p.CSRF))
	//line template/namespace/form.qtpl:92
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:93
	if p.Namespace != nil {
		//line template/namespace/form.qtpl:93
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
		//line template/namespace/form.qtpl:95
	}
	//line template/namespace/form.qtpl:95
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:96
	if p.Parent != nil && !p.Parent.IsZero() {
		//line template/namespace/form.qtpl:96
		qw422016.N().S(` <input type="hidden" name="parent" value="`)
		//line template/namespace/form.qtpl:97
		qw422016.E().S(p.Parent.Path)
		//line template/namespace/form.qtpl:97
		qw422016.N().S(`"/> `)
		//line template/namespace/form.qtpl:98
	}
	//line template/namespace/form.qtpl:98
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:99
	p.StreamError(qw422016, "namespace")
	//line template/namespace/form.qtpl:99
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input class="form-text" type="text" name="name" value="`)
	//line template/namespace/form.qtpl:102
	qw422016.E().S(p.Field("name"))
	//line template/namespace/form.qtpl:102
	qw422016.N().S(`" autocomplete="off"/> `)
	//line template/namespace/form.qtpl:103
	p.StreamError(qw422016, "name")
	//line template/namespace/form.qtpl:103
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="name">Description</label> <input class="form-text" type="text" name="description" value="`)
	//line template/namespace/form.qtpl:107
	qw422016.E().S(p.Field("description"))
	//line template/namespace/form.qtpl:107
	qw422016.N().S(`" autocomplete="off"/> `)
	//line template/namespace/form.qtpl:108
	p.StreamError(qw422016, "description")
	//line template/namespace/form.qtpl:108
	qw422016.N().S(` </div> <div class="form-field"> <label class="form-option `)
	//line template/namespace/form.qtpl:111
	qw422016.E().S(p.disabled(model.Private))
	//line template/namespace/form.qtpl:111
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="private" `)
	//line template/namespace/form.qtpl:112
	qw422016.E().S(p.checked(model.Private))
	//line template/namespace/form.qtpl:112
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:112
	qw422016.E().S(p.disabled(model.Private))
	//line template/namespace/form.qtpl:112
	qw422016.N().S(`/> <strong>Private</strong> `)
	//line template/namespace/form.qtpl:114
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.094 8.016v-2.016c0-1.688-1.406-3.094-3.094-3.094s-3.094 1.406-3.094 3.094v2.016h6.188zM12 17.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-2.016 0.938-2.016 2.016 0.938 2.016 2.016 2.016zM18 8.016c1.078 0 2.016 0.891 2.016 1.969v10.031c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h0.984v-2.016c0-2.766 2.25-5.016 5.016-5.016s5.016 2.25 5.016 5.016v2.016h0.984z"></path>
</svg>
`)
	//line template/namespace/form.qtpl:114
	qw422016.N().S(` <div class="form-desc">You choose who can view builds in the namespace.</div> </label> <label class="form-option `)
	//line template/namespace/form.qtpl:117
	qw422016.E().S(p.disabled(model.Internal))
	//line template/namespace/form.qtpl:117
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="internal" `)
	//line template/namespace/form.qtpl:118
	qw422016.E().S(p.checked(model.Internal))
	//line template/namespace/form.qtpl:118
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:118
	qw422016.E().S(p.disabled(model.Internal))
	//line template/namespace/form.qtpl:118
	qw422016.N().S(`/> <strong>Internal</strong> `)
	//line template/namespace/form.qtpl:120
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.984l9 4.031v6c0 5.531-3.844 10.734-9 12-5.156-1.266-9-6.469-9-12v-6zM12 12v8.953c3.703-1.172 6.469-4.828 6.984-8.953h-6.984zM12 12v-8.813l-6.984 3.094v5.719h6.984z"></path>
</svg>
`)
	//line template/namespace/form.qtpl:120
	qw422016.N().S(` <div class="form-desc">Anyone with an account will be able to view builds in the namespace.</div> </label> <label class="form-option `)
	//line template/namespace/form.qtpl:123
	qw422016.E().S(p.disabled(model.Public))
	//line template/namespace/form.qtpl:123
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="public" `)
	//line template/namespace/form.qtpl:124
	qw422016.E().S(p.checked(model.Public))
	//line template/namespace/form.qtpl:124
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:124
	qw422016.E().S(p.disabled(model.Public))
	//line template/namespace/form.qtpl:124
	qw422016.N().S(`/> <strong>Public</strong> `)
	//line template/namespace/form.qtpl:126
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.906 17.391c1.313-1.406 2.109-3.328 2.109-5.391 0-3.328-2.063-6.234-5.016-7.406v0.422c0 1.078-0.938 1.969-2.016 1.969h-1.969v2.016c0 0.563-0.469 0.984-1.031 0.984h-1.969v2.016h6c0.563 0 0.984 0.422 0.984 0.984v3h0.984c0.891 0 1.641 0.609 1.922 1.406zM11.016 19.922v-1.922c-1.078 0-2.016-0.938-2.016-2.016v-0.984l-4.781-4.781c-0.141 0.563-0.234 1.172-0.234 1.781 0 4.078 3.094 7.453 7.031 7.922zM12 2.016c5.531 0 9.984 4.453 9.984 9.984s-4.453 9.984-9.984 9.984-9.984-4.453-9.984-9.984 4.453-9.984 9.984-9.984z"></path>
</svg>
`)
	//line template/namespace/form.qtpl:126
	qw422016.N().S(` <div class="form-desc">Anyone will be able to view builds in the namespace.</div> </label> </div> <div class="form-field"> `)
	//line template/namespace/form.qtpl:131
	if p.Namespace != nil {
		//line template/namespace/form.qtpl:131
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
		//line template/namespace/form.qtpl:133
	} else {
		//line template/namespace/form.qtpl:133
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
		//line template/namespace/form.qtpl:135
	}
	//line template/namespace/form.qtpl:135
	qw422016.N().S(` </div> </form> `)
	//line template/namespace/form.qtpl:138
	if p.Namespace != nil {
		//line template/namespace/form.qtpl:138
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="`)
		//line template/namespace/form.qtpl:140
		qw422016.E().S(p.Namespace.UIEndpoint())
		//line template/namespace/form.qtpl:140
		qw422016.N().S(`"> `)
		//line template/namespace/form.qtpl:141
		qw422016.N().S(string(p.CSRF))
		//line template/namespace/form.qtpl:141
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <div class="overflow"> <div class="right"> <button type="submit" class="btn btn-danger">Delete</button> </div> <strong>Delete Namespace</strong><br/><p>Builds within the namespace will not be deleted.</p> </div> </form> `)
		//line template/namespace/form.qtpl:150
	}
	//line template/namespace/form.qtpl:150
	qw422016.N().S(` </div> </div> `)
//line template/namespace/form.qtpl:153
}

//line template/namespace/form.qtpl:153
func (p *Form) WriteBody(qq422016 qtio422016.Writer) {
	//line template/namespace/form.qtpl:153
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/form.qtpl:153
	p.StreamBody(qw422016)
	//line template/namespace/form.qtpl:153
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/form.qtpl:153
}

//line template/namespace/form.qtpl:153
func (p *Form) Body() string {
	//line template/namespace/form.qtpl:153
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/form.qtpl:153
	p.WriteBody(qb422016)
	//line template/namespace/form.qtpl:153
	qs422016 := string(qb422016.B)
	//line template/namespace/form.qtpl:153
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/form.qtpl:153
	return qs422016
//line template/namespace/form.qtpl:153
}

//line template/namespace/form.qtpl:155
func (p *Form) StreamHeader(qw422016 *qt422016.Writer) {
	//line template/namespace/form.qtpl:155
	qw422016.N().S(` `)
	//line template/namespace/form.qtpl:156
	if p.Namespace != nil {
		//line template/namespace/form.qtpl:156
		qw422016.N().S(` <a class="back" href="/u/`)
		//line template/namespace/form.qtpl:157
		qw422016.E().S(p.Namespace.User.Username)
		//line template/namespace/form.qtpl:157
		qw422016.N().S(`/`)
		//line template/namespace/form.qtpl:157
		qw422016.E().S(p.Namespace.Path)
		//line template/namespace/form.qtpl:157
		qw422016.N().S(`">`)
		//line template/namespace/form.qtpl:157
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
		//line template/namespace/form.qtpl:157
		qw422016.N().S(`</a> `)
		//line template/namespace/form.qtpl:158
		streamrenderPath(qw422016, p.Namespace.User.Username, p.Namespace.Path)
		//line template/namespace/form.qtpl:158
		qw422016.N().S(` - Edit `)
		//line template/namespace/form.qtpl:159
	} else {
		//line template/namespace/form.qtpl:159
		qw422016.N().S(` `)
		//line template/namespace/form.qtpl:160
		if p.Parent != nil && !p.Parent.IsZero() {
			//line template/namespace/form.qtpl:160
			qw422016.N().S(` <a class="back" href="`)
			//line template/namespace/form.qtpl:161
			qw422016.E().S(p.Parent.UIEndpoint())
			//line template/namespace/form.qtpl:161
			qw422016.N().S(`">`)
			//line template/namespace/form.qtpl:161
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
			//line template/namespace/form.qtpl:161
			qw422016.N().S(`</a> `)
			//line template/namespace/form.qtpl:162
			streamrenderPath(qw422016, p.Parent.User.Username, p.Parent.Path)
			//line template/namespace/form.qtpl:162
			qw422016.N().S(` - Create Sub-namespace `)
			//line template/namespace/form.qtpl:163
		} else {
			//line template/namespace/form.qtpl:163
			qw422016.N().S(` <a class="back" href="/namespaces">`)
			//line template/namespace/form.qtpl:164
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
			//line template/namespace/form.qtpl:164
			qw422016.N().S(`</a> Create Namespace `)
			//line template/namespace/form.qtpl:165
		}
		//line template/namespace/form.qtpl:165
		qw422016.N().S(` `)
		//line template/namespace/form.qtpl:166
	}
	//line template/namespace/form.qtpl:166
	qw422016.N().S(` `)
//line template/namespace/form.qtpl:167
}

//line template/namespace/form.qtpl:167
func (p *Form) WriteHeader(qq422016 qtio422016.Writer) {
	//line template/namespace/form.qtpl:167
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/form.qtpl:167
	p.StreamHeader(qw422016)
	//line template/namespace/form.qtpl:167
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/form.qtpl:167
}

//line template/namespace/form.qtpl:167
func (p *Form) Header() string {
	//line template/namespace/form.qtpl:167
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/form.qtpl:167
	p.WriteHeader(qb422016)
	//line template/namespace/form.qtpl:167
	qs422016 := string(qb422016.B)
	//line template/namespace/form.qtpl:167
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/form.qtpl:167
	return qs422016
//line template/namespace/form.qtpl:167
}

//line template/namespace/form.qtpl:169
func (p *Form) StreamActions(qw422016 *qt422016.Writer) {
//line template/namespace/form.qtpl:169
}

//line template/namespace/form.qtpl:169
func (p *Form) WriteActions(qq422016 qtio422016.Writer) {
	//line template/namespace/form.qtpl:169
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/form.qtpl:169
	p.StreamActions(qw422016)
	//line template/namespace/form.qtpl:169
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/form.qtpl:169
}

//line template/namespace/form.qtpl:169
func (p *Form) Actions() string {
	//line template/namespace/form.qtpl:169
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/form.qtpl:169
	p.WriteActions(qb422016)
	//line template/namespace/form.qtpl:169
	qs422016 := string(qb422016.B)
	//line template/namespace/form.qtpl:169
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/form.qtpl:169
	return qs422016
//line template/namespace/form.qtpl:169
}

//line template/namespace/form.qtpl:170
func (p *Form) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/namespace/form.qtpl:170
}

//line template/namespace/form.qtpl:170
func (p *Form) WriteNavigation(qq422016 qtio422016.Writer) {
	//line template/namespace/form.qtpl:170
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/form.qtpl:170
	p.StreamNavigation(qw422016)
	//line template/namespace/form.qtpl:170
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/form.qtpl:170
}

//line template/namespace/form.qtpl:170
func (p *Form) Navigation() string {
	//line template/namespace/form.qtpl:170
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/form.qtpl:170
	p.WriteNavigation(qb422016)
	//line template/namespace/form.qtpl:170
	qs422016 := string(qb422016.B)
	//line template/namespace/form.qtpl:170
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/form.qtpl:170
	return qs422016
//line template/namespace/form.qtpl:170
}

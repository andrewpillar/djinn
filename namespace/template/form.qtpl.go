// Code generated by qtc from "form.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line namespace/template/form.qtpl:2
package template

//line namespace/template/form.qtpl:2
import (
	"strings"

	"djinn-ci.com/namespace"
	"djinn-ci.com/template"
)

//line namespace/template/form.qtpl:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line namespace/template/form.qtpl:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line namespace/template/form.qtpl:11
type Form struct {
	template.BasePage
	template.Form

	Parent    *namespace.Namespace
	Namespace *namespace.Namespace
}

type WebhookForm struct {
	template.BasePage
	template.Form

	Namespace  *namespace.Namespace
	Webhook    *namespace.Webhook
	Deliveries []*namespace.WebhookDelivery
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

func (p *WebhookForm) action() string {
	if p.Webhook == nil {
		return p.Namespace.Endpoint("webhooks")
	}
	return p.Webhook.Endpoint()
}

func (p *WebhookForm) Field(field string) string {
	old := p.Form.Fields[field]

	if p.Webhook != nil {
		if old != "" {
			return old
		}

		if field == "payload_url" {
			return p.Webhook.PayloadURL.String()
		}
		return ""
	}
	return old
}

//line namespace/template/form.qtpl:116
func (p *Form) StreamTitle(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:116
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:117
	if p.Namespace == nil {
//line namespace/template/form.qtpl:117
		qw422016.N().S(` Create Namespace - Djinn CI `)
//line namespace/template/form.qtpl:119
	} else {
//line namespace/template/form.qtpl:119
		qw422016.N().S(` Edit Namespace - Djinn CI `)
//line namespace/template/form.qtpl:121
	}
//line namespace/template/form.qtpl:121
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:122
}

//line namespace/template/form.qtpl:122
func (p *Form) WriteTitle(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:122
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:122
	p.StreamTitle(qw422016)
//line namespace/template/form.qtpl:122
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:122
}

//line namespace/template/form.qtpl:122
func (p *Form) Title() string {
//line namespace/template/form.qtpl:122
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:122
	p.WriteTitle(qb422016)
//line namespace/template/form.qtpl:122
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:122
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:122
	return qs422016
//line namespace/template/form.qtpl:122
}

//line namespace/template/form.qtpl:124
func (p *Form) StreamBody(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:124
	qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> <form method="POST" action="`)
//line namespace/template/form.qtpl:127
	qw422016.E().S(p.action())
//line namespace/template/form.qtpl:127
	qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:128
	qw422016.N().S(string(p.CSRF))
//line namespace/template/form.qtpl:128
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:129
	if p.Namespace != nil {
//line namespace/template/form.qtpl:129
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
//line namespace/template/form.qtpl:131
	}
//line namespace/template/form.qtpl:131
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:132
	if p.Parent != nil && !p.Parent.IsZero() {
//line namespace/template/form.qtpl:132
		qw422016.N().S(` <input type="hidden" name="parent" value="`)
//line namespace/template/form.qtpl:133
		qw422016.E().S(p.Parent.Path)
//line namespace/template/form.qtpl:133
		qw422016.N().S(`"/> `)
//line namespace/template/form.qtpl:134
	}
//line namespace/template/form.qtpl:134
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:135
	if p.Namespace == nil {
//line namespace/template/form.qtpl:135
		qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input class="form-text" type="text" name="name" value="`)
//line namespace/template/form.qtpl:138
		qw422016.E().S(p.Field("name"))
//line namespace/template/form.qtpl:138
		qw422016.N().S(`" autocomplete="off"/> `)
//line namespace/template/form.qtpl:139
		p.StreamError(qw422016, "name")
//line namespace/template/form.qtpl:139
		qw422016.N().S(` </div> `)
//line namespace/template/form.qtpl:141
	}
//line namespace/template/form.qtpl:141
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Description</label> <input class="form-text" type="text" name="description" value="`)
//line namespace/template/form.qtpl:144
	qw422016.E().S(p.Field("description"))
//line namespace/template/form.qtpl:144
	qw422016.N().S(`" autocomplete="off"/> `)
//line namespace/template/form.qtpl:145
	p.StreamError(qw422016, "description")
//line namespace/template/form.qtpl:145
	qw422016.N().S(` </div> <div class="form-field"> <label class="form-option `)
//line namespace/template/form.qtpl:148
	qw422016.E().S(p.disabled(namespace.Private))
//line namespace/template/form.qtpl:148
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="private" `)
//line namespace/template/form.qtpl:149
	qw422016.E().S(p.checked(namespace.Private))
//line namespace/template/form.qtpl:149
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:149
	qw422016.E().S(p.disabled(namespace.Private))
//line namespace/template/form.qtpl:149
	qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:150
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.094 8.016v-2.016c0-1.688-1.406-3.094-3.094-3.094s-3.094 1.406-3.094 3.094v2.016h6.188zM12 17.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-2.016 0.938-2.016 2.016 0.938 2.016 2.016 2.016zM18 8.016c1.078 0 2.016 0.891 2.016 1.969v10.031c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h0.984v-2.016c0-2.766 2.25-5.016 5.016-5.016s5.016 2.25 5.016 5.016v2.016h0.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:150
	qw422016.N().S(` <div class="form-option-info"> <strong>Private</strong> <div class="form-desc">You choose who can view builds in the namespace.</div> </div> </label> <label class="form-option `)
//line namespace/template/form.qtpl:156
	qw422016.E().S(p.disabled(namespace.Internal))
//line namespace/template/form.qtpl:156
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="internal" `)
//line namespace/template/form.qtpl:157
	qw422016.E().S(p.checked(namespace.Internal))
//line namespace/template/form.qtpl:157
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:157
	qw422016.E().S(p.disabled(namespace.Internal))
//line namespace/template/form.qtpl:157
	qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:158
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.984l9 4.031v6c0 5.531-3.844 10.734-9 12-5.156-1.266-9-6.469-9-12v-6zM12 12v8.953c3.703-1.172 6.469-4.828 6.984-8.953h-6.984zM12 12v-8.813l-6.984 3.094v5.719h6.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:158
	qw422016.N().S(` <div class="form-option-info"> <strong>Internal</strong> <div class="form-desc">Anyone with an account will be able to view builds in the namespace.</div> </div> </label> <label class="form-option `)
//line namespace/template/form.qtpl:164
	qw422016.E().S(p.disabled(namespace.Public))
//line namespace/template/form.qtpl:164
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="public" `)
//line namespace/template/form.qtpl:165
	qw422016.E().S(p.checked(namespace.Public))
//line namespace/template/form.qtpl:165
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:165
	qw422016.E().S(p.disabled(namespace.Public))
//line namespace/template/form.qtpl:165
	qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:166
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.906 17.391c1.313-1.406 2.109-3.328 2.109-5.391 0-3.328-2.063-6.234-5.016-7.406v0.422c0 1.078-0.938 1.969-2.016 1.969h-1.969v2.016c0 0.563-0.469 0.984-1.031 0.984h-1.969v2.016h6c0.563 0 0.984 0.422 0.984 0.984v3h0.984c0.891 0 1.641 0.609 1.922 1.406zM11.016 19.922v-1.922c-1.078 0-2.016-0.938-2.016-2.016v-0.984l-4.781-4.781c-0.141 0.563-0.234 1.172-0.234 1.781 0 4.078 3.094 7.453 7.031 7.922zM12 2.016c5.531 0 9.984 4.453 9.984 9.984s-4.453 9.984-9.984 9.984-9.984-4.453-9.984-9.984 4.453-9.984 9.984-9.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:166
	qw422016.N().S(` <div class="form-option-info"> <strong>Public</strong> <div class="form-desc">Anyone will be able to view builds in the namespace.</div> </div> </label> </div> <div class="form-field"> `)
//line namespace/template/form.qtpl:174
	if p.Namespace != nil {
//line namespace/template/form.qtpl:174
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line namespace/template/form.qtpl:176
	} else {
//line namespace/template/form.qtpl:176
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line namespace/template/form.qtpl:178
	}
//line namespace/template/form.qtpl:178
	qw422016.N().S(` </div> </form> `)
//line namespace/template/form.qtpl:181
	if p.Namespace != nil {
//line namespace/template/form.qtpl:181
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="`)
//line namespace/template/form.qtpl:183
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:183
		qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:184
		qw422016.N().S(string(p.CSRF))
//line namespace/template/form.qtpl:184
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <div class="overflow"> <div class="right"> <button type="submit" class="btn btn-danger">Delete</button> </div> <strong>Delete Namespace</strong><br/><p>Builds within the namespace will not be deleted.</p> </div> </form> `)
//line namespace/template/form.qtpl:193
	}
//line namespace/template/form.qtpl:193
	qw422016.N().S(` </div> </div> `)
//line namespace/template/form.qtpl:196
}

//line namespace/template/form.qtpl:196
func (p *Form) WriteBody(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:196
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:196
	p.StreamBody(qw422016)
//line namespace/template/form.qtpl:196
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:196
}

//line namespace/template/form.qtpl:196
func (p *Form) Body() string {
//line namespace/template/form.qtpl:196
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:196
	p.WriteBody(qb422016)
//line namespace/template/form.qtpl:196
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:196
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:196
	return qs422016
//line namespace/template/form.qtpl:196
}

//line namespace/template/form.qtpl:198
func (p *Form) StreamHeader(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:198
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:199
	if p.Namespace != nil {
//line namespace/template/form.qtpl:199
		qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:200
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:200
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:200
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:200
		qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:201
		streamrenderPath(qw422016, p.Namespace.User.Username, p.Namespace.Path)
//line namespace/template/form.qtpl:201
		qw422016.N().S(` - Edit `)
//line namespace/template/form.qtpl:202
	} else {
//line namespace/template/form.qtpl:202
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:203
		if p.Parent != nil && !p.Parent.IsZero() {
//line namespace/template/form.qtpl:203
			qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:204
			qw422016.E().S(p.Parent.Endpoint())
//line namespace/template/form.qtpl:204
			qw422016.N().S(`">`)
//line namespace/template/form.qtpl:204
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:204
			qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:205
			streamrenderPath(qw422016, p.Parent.User.Username, p.Parent.Path)
//line namespace/template/form.qtpl:205
			qw422016.N().S(` - Create Sub-namespace `)
//line namespace/template/form.qtpl:206
		} else {
//line namespace/template/form.qtpl:206
			qw422016.N().S(` <a class="back" href="/namespaces">`)
//line namespace/template/form.qtpl:207
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:207
			qw422016.N().S(`</a> Create Namespace `)
//line namespace/template/form.qtpl:208
		}
//line namespace/template/form.qtpl:208
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:209
	}
//line namespace/template/form.qtpl:209
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:210
}

//line namespace/template/form.qtpl:210
func (p *Form) WriteHeader(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:210
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:210
	p.StreamHeader(qw422016)
//line namespace/template/form.qtpl:210
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:210
}

//line namespace/template/form.qtpl:210
func (p *Form) Header() string {
//line namespace/template/form.qtpl:210
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:210
	p.WriteHeader(qb422016)
//line namespace/template/form.qtpl:210
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:210
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:210
	return qs422016
//line namespace/template/form.qtpl:210
}

//line namespace/template/form.qtpl:212
func (p *Form) StreamActions(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:212
}

//line namespace/template/form.qtpl:212
func (p *Form) WriteActions(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:212
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:212
	p.StreamActions(qw422016)
//line namespace/template/form.qtpl:212
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:212
}

//line namespace/template/form.qtpl:212
func (p *Form) Actions() string {
//line namespace/template/form.qtpl:212
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:212
	p.WriteActions(qb422016)
//line namespace/template/form.qtpl:212
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:212
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:212
	return qs422016
//line namespace/template/form.qtpl:212
}

//line namespace/template/form.qtpl:213
func (p *Form) StreamNavigation(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:213
}

//line namespace/template/form.qtpl:213
func (p *Form) WriteNavigation(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:213
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:213
	p.StreamNavigation(qw422016)
//line namespace/template/form.qtpl:213
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:213
}

//line namespace/template/form.qtpl:213
func (p *Form) Navigation() string {
//line namespace/template/form.qtpl:213
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:213
	p.WriteNavigation(qb422016)
//line namespace/template/form.qtpl:213
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:213
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:213
	return qs422016
//line namespace/template/form.qtpl:213
}

//line namespace/template/form.qtpl:215
func (p *WebhookForm) StreamTitle(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:215
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:216
	if p.Webhook == nil {
//line namespace/template/form.qtpl:216
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:217
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:217
		qw422016.N().S(` - Create webhook - Djinn CI `)
//line namespace/template/form.qtpl:218
	} else {
//line namespace/template/form.qtpl:218
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:219
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:219
		qw422016.N().S(` - Edit webhook - Djinn CI `)
//line namespace/template/form.qtpl:220
	}
//line namespace/template/form.qtpl:220
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:221
}

//line namespace/template/form.qtpl:221
func (p *WebhookForm) WriteTitle(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:221
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:221
	p.StreamTitle(qw422016)
//line namespace/template/form.qtpl:221
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:221
}

//line namespace/template/form.qtpl:221
func (p *WebhookForm) Title() string {
//line namespace/template/form.qtpl:221
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:221
	p.WriteTitle(qb422016)
//line namespace/template/form.qtpl:221
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:221
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:221
	return qs422016
//line namespace/template/form.qtpl:221
}

//line namespace/template/form.qtpl:223
func (p *WebhookForm) StreamHeader(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:223
	qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:224
	qw422016.E().S(p.Namespace.Endpoint("webhooks"))
//line namespace/template/form.qtpl:224
	qw422016.N().S(`">`)
//line namespace/template/form.qtpl:224
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:224
	qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:225
	if p.Webhook == nil {
//line namespace/template/form.qtpl:225
		qw422016.N().S(` <a href="`)
//line namespace/template/form.qtpl:226
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:226
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:226
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:226
		qw422016.N().S(`</a> / Create webhook `)
//line namespace/template/form.qtpl:227
	} else {
//line namespace/template/form.qtpl:227
		qw422016.N().S(` <a href="`)
//line namespace/template/form.qtpl:228
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:228
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:228
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:228
		qw422016.N().S(`</a> / Edit webhook `)
//line namespace/template/form.qtpl:229
	}
//line namespace/template/form.qtpl:229
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:230
}

//line namespace/template/form.qtpl:230
func (p *WebhookForm) WriteHeader(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:230
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:230
	p.StreamHeader(qw422016)
//line namespace/template/form.qtpl:230
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:230
}

//line namespace/template/form.qtpl:230
func (p *WebhookForm) Header() string {
//line namespace/template/form.qtpl:230
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:230
	p.WriteHeader(qb422016)
//line namespace/template/form.qtpl:230
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:230
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:230
	return qs422016
//line namespace/template/form.qtpl:230
}

//line namespace/template/form.qtpl:232
func (p *WebhookForm) StreamBody(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:232
	qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> <form method="POST" action="`)
//line namespace/template/form.qtpl:235
	qw422016.E().S(p.action())
//line namespace/template/form.qtpl:235
	qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:236
	if p.Webhook != nil {
//line namespace/template/form.qtpl:236
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
//line namespace/template/form.qtpl:238
	}
//line namespace/template/form.qtpl:238
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:239
	qw422016.N().V(p.CSRF)
//line namespace/template/form.qtpl:239
	qw422016.N().S(` <div class="form-field"> <label class="label" for="payload_url">Payload URL</label> <input type="text" class="form-text" id="payload_url" name="payload_url" placeholder="https://example.com" autocomplete="off" value="`)
//line namespace/template/form.qtpl:242
	qw422016.E().S(p.Field("payload_url"))
//line namespace/template/form.qtpl:242
	qw422016.N().S(`"/> `)
//line namespace/template/form.qtpl:243
	p.StreamError(qw422016, "payload_url")
//line namespace/template/form.qtpl:243
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="secret">Secret</label> <input type="password" class="form-text" id="secret" name="secret" autocomplete="off"/> </div> <div class="form-field"> <label class="form-option"> <input class="form-selector" type="checkbox" name="ssl" checked="true"/> <div class="form-option-info"> <strong>SSL</strong> <div class="form-desc">Verify SSL certificates when sending payloads</div> </div> </label> </div> <div class="form-field"> <label class="form-option"> <input class="form-selector" type="checkbox" name="active" checked="true"/	> <div class="form-option-info"> <strong>Active</strong> <div class="form-desc">Send events for this webhook</div> </div> </label> </div> <div class="form-field"> <strong>Events</strong><br/> `)
//line namespace/template/form.qtpl:269
	for _, event := range namespace.WebhookEvents {
//line namespace/template/form.qtpl:269
		qw422016.N().S(` <label class="hook-event"> `)
//line namespace/template/form.qtpl:271
		if p.Webhook != nil && p.Webhook.Events.Has(event) {
//line namespace/template/form.qtpl:271
			qw422016.N().S(` <input checked="true" class="form-selector" type="checkbox" name="events[]" value="`)
//line namespace/template/form.qtpl:272
			qw422016.E().S(event.String())
//line namespace/template/form.qtpl:272
			qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:272
			qw422016.E().S(strings.Replace(strings.Title(event.String()), "_", " ", -1))
//line namespace/template/form.qtpl:272
			qw422016.N().S(` `)
//line namespace/template/form.qtpl:273
		} else {
//line namespace/template/form.qtpl:273
			qw422016.N().S(` <input class="form-selector" type="checkbox" name="events[]" value="`)
//line namespace/template/form.qtpl:274
			qw422016.E().S(event.String())
//line namespace/template/form.qtpl:274
			qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:274
			qw422016.E().S(strings.Replace(strings.Title(event.String()), "_", " ", -1))
//line namespace/template/form.qtpl:274
			qw422016.N().S(` `)
//line namespace/template/form.qtpl:275
		}
//line namespace/template/form.qtpl:275
		qw422016.N().S(` </label> `)
//line namespace/template/form.qtpl:277
	}
//line namespace/template/form.qtpl:277
	qw422016.N().S(` </div> <div class="form-field"> `)
//line namespace/template/form.qtpl:280
	if p.Webhook == nil {
//line namespace/template/form.qtpl:280
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line namespace/template/form.qtpl:282
	} else {
//line namespace/template/form.qtpl:282
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line namespace/template/form.qtpl:284
	}
//line namespace/template/form.qtpl:284
	qw422016.N().S(` </div> </form> `)
//line namespace/template/form.qtpl:287
	if p.Webhook != nil {
//line namespace/template/form.qtpl:287
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="`)
//line namespace/template/form.qtpl:289
		qw422016.E().S(p.Webhook.Endpoint())
//line namespace/template/form.qtpl:289
		qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:290
		qw422016.N().V(p.CSRF)
//line namespace/template/form.qtpl:290
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <div class="overflow"> <div class="right"> <button class="btn btn-danger" type="submit">Delete</button> </div> </div> </form> `)
//line namespace/template/form.qtpl:298
	}
//line namespace/template/form.qtpl:298
	qw422016.N().S(` </div> </div> `)
//line namespace/template/form.qtpl:301
	if p.Webhook != nil {
//line namespace/template/form.qtpl:301
		qw422016.N().S(` <div class="panel"> <div class="panel-header"> <h3>Recent deliveries</h3> </div> `)
//line namespace/template/form.qtpl:306
		if len(p.Deliveries) == 0 {
//line namespace/template/form.qtpl:306
			qw422016.N().S(` <div class="panel-message muted">No recent deliveries.</div> `)
//line namespace/template/form.qtpl:308
		} else {
//line namespace/template/form.qtpl:308
			qw422016.N().S(` <table class="table"> <tbody> `)
//line namespace/template/form.qtpl:311
			for _, d := range p.Deliveries {
//line namespace/template/form.qtpl:311
				qw422016.N().S(` <tr> `)
//line namespace/template/form.qtpl:313
				if d.ResponseCode >= 200 && d.ResponseCode < 300 {
//line namespace/template/form.qtpl:313
					qw422016.N().S(` <td>`)
//line namespace/template/form.qtpl:314
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
//line namespace/template/form.qtpl:314
					qw422016.N().S(`</td> `)
//line namespace/template/form.qtpl:315
				} else {
//line namespace/template/form.qtpl:315
					qw422016.N().S(` <td>`)
//line namespace/template/form.qtpl:316
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
//line namespace/template/form.qtpl:316
					qw422016.N().S(`</td> `)
//line namespace/template/form.qtpl:317
				}
//line namespace/template/form.qtpl:317
				qw422016.N().S(` <td><code><a href="">`)
//line namespace/template/form.qtpl:318
				qw422016.E().S(d.DeliveryID)
//line namespace/template/form.qtpl:318
				qw422016.N().S(`</a></code></td> <td>`)
//line namespace/template/form.qtpl:319
				qw422016.E().S(d.Duration.String())
//line namespace/template/form.qtpl:319
				qw422016.N().S(`</td> </tr> `)
//line namespace/template/form.qtpl:321
			}
//line namespace/template/form.qtpl:321
			qw422016.N().S(` </tbody> </table> `)
//line namespace/template/form.qtpl:324
		}
//line namespace/template/form.qtpl:324
		qw422016.N().S(` </div> `)
//line namespace/template/form.qtpl:326
	}
//line namespace/template/form.qtpl:326
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:327
}

//line namespace/template/form.qtpl:327
func (p *WebhookForm) WriteBody(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:327
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:327
	p.StreamBody(qw422016)
//line namespace/template/form.qtpl:327
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:327
}

//line namespace/template/form.qtpl:327
func (p *WebhookForm) Body() string {
//line namespace/template/form.qtpl:327
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:327
	p.WriteBody(qb422016)
//line namespace/template/form.qtpl:327
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:327
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:327
	return qs422016
//line namespace/template/form.qtpl:327
}

//line namespace/template/form.qtpl:329
func (p *WebhookForm) StreamActions(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:329
}

//line namespace/template/form.qtpl:329
func (p *WebhookForm) WriteActions(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:329
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:329
	p.StreamActions(qw422016)
//line namespace/template/form.qtpl:329
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:329
}

//line namespace/template/form.qtpl:329
func (p *WebhookForm) Actions() string {
//line namespace/template/form.qtpl:329
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:329
	p.WriteActions(qb422016)
//line namespace/template/form.qtpl:329
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:329
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:329
	return qs422016
//line namespace/template/form.qtpl:329
}

//line namespace/template/form.qtpl:330
func (p *WebhookForm) StreamNavigation(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:330
}

//line namespace/template/form.qtpl:330
func (p *WebhookForm) WriteNavigation(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:330
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:330
	p.StreamNavigation(qw422016)
//line namespace/template/form.qtpl:330
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:330
}

//line namespace/template/form.qtpl:330
func (p *WebhookForm) Navigation() string {
//line namespace/template/form.qtpl:330
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:330
	p.WriteNavigation(qb422016)
//line namespace/template/form.qtpl:330
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:330
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:330
	return qs422016
//line namespace/template/form.qtpl:330
}

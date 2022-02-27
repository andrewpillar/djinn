// Code generated by qtc from "form.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line namespace/template/form.qtpl:2
package template

//line namespace/template/form.qtpl:2
import (
	"strings"

	"djinn-ci.com/event"
	"djinn-ci.com/namespace"
	"djinn-ci.com/template"
)

//line namespace/template/form.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line namespace/template/form.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line namespace/template/form.qtpl:12
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

//line namespace/template/form.qtpl:117
func (p *Form) StreamTitle(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:117
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:118
	if p.Namespace == nil {
//line namespace/template/form.qtpl:118
		qw422016.N().S(` Create Namespace - Djinn CI `)
//line namespace/template/form.qtpl:120
	} else {
//line namespace/template/form.qtpl:120
		qw422016.N().S(` Edit Namespace - Djinn CI `)
//line namespace/template/form.qtpl:122
	}
//line namespace/template/form.qtpl:122
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:123
}

//line namespace/template/form.qtpl:123
func (p *Form) WriteTitle(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:123
	p.StreamTitle(qw422016)
//line namespace/template/form.qtpl:123
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:123
}

//line namespace/template/form.qtpl:123
func (p *Form) Title() string {
//line namespace/template/form.qtpl:123
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:123
	p.WriteTitle(qb422016)
//line namespace/template/form.qtpl:123
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:123
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:123
	return qs422016
//line namespace/template/form.qtpl:123
}

//line namespace/template/form.qtpl:125
func (p *Form) StreamBody(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:125
	qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> <form method="POST" action="`)
//line namespace/template/form.qtpl:128
	qw422016.E().S(p.action())
//line namespace/template/form.qtpl:128
	qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:129
	qw422016.N().S(string(p.CSRF))
//line namespace/template/form.qtpl:129
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:130
	if p.Namespace != nil {
//line namespace/template/form.qtpl:130
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
//line namespace/template/form.qtpl:132
	}
//line namespace/template/form.qtpl:132
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:133
	if p.Parent != nil {
//line namespace/template/form.qtpl:133
		qw422016.N().S(` <input type="hidden" name="parent" value="`)
//line namespace/template/form.qtpl:134
		qw422016.E().S(p.Parent.Path)
//line namespace/template/form.qtpl:134
		qw422016.N().S(`"/> `)
//line namespace/template/form.qtpl:135
	}
//line namespace/template/form.qtpl:135
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:136
	if p.Namespace == nil {
//line namespace/template/form.qtpl:136
		qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Name</label> <input class="form-text" type="text" name="name" value="`)
//line namespace/template/form.qtpl:139
		qw422016.E().S(p.Field("name"))
//line namespace/template/form.qtpl:139
		qw422016.N().S(`" autocomplete="off"/> `)
//line namespace/template/form.qtpl:140
		p.StreamError(qw422016, "name")
//line namespace/template/form.qtpl:140
		qw422016.N().S(` </div> `)
//line namespace/template/form.qtpl:142
	}
//line namespace/template/form.qtpl:142
	qw422016.N().S(` <div class="form-field"> <label class="label" for="name">Description</label> <input class="form-text" type="text" name="description" value="`)
//line namespace/template/form.qtpl:145
	qw422016.E().S(p.Field("description"))
//line namespace/template/form.qtpl:145
	qw422016.N().S(`" autocomplete="off"/> `)
//line namespace/template/form.qtpl:146
	p.StreamError(qw422016, "description")
//line namespace/template/form.qtpl:146
	qw422016.N().S(` </div> <div class="form-field"> <label class="form-option `)
//line namespace/template/form.qtpl:149
	qw422016.E().S(p.disabled(namespace.Private))
//line namespace/template/form.qtpl:149
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="private" `)
//line namespace/template/form.qtpl:150
	qw422016.E().S(p.checked(namespace.Private))
//line namespace/template/form.qtpl:150
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:150
	qw422016.E().S(p.disabled(namespace.Private))
//line namespace/template/form.qtpl:150
	qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:151
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M15.094 8.016v-2.016c0-1.688-1.406-3.094-3.094-3.094s-3.094 1.406-3.094 3.094v2.016h6.188zM12 17.016c1.078 0 2.016-0.938 2.016-2.016s-0.938-2.016-2.016-2.016-2.016 0.938-2.016 2.016 0.938 2.016 2.016 2.016zM18 8.016c1.078 0 2.016 0.891 2.016 1.969v10.031c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969v-10.031c0-1.078 0.938-1.969 2.016-1.969h0.984v-2.016c0-2.766 2.25-5.016 5.016-5.016s5.016 2.25 5.016 5.016v2.016h0.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:151
	qw422016.N().S(` <div class="form-option-info"> <strong>Private</strong> <div class="form-desc">You choose who can view builds in the namespace.</div> </div> </label> <label class="form-option `)
//line namespace/template/form.qtpl:157
	qw422016.E().S(p.disabled(namespace.Internal))
//line namespace/template/form.qtpl:157
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="internal" `)
//line namespace/template/form.qtpl:158
	qw422016.E().S(p.checked(namespace.Internal))
//line namespace/template/form.qtpl:158
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:158
	qw422016.E().S(p.disabled(namespace.Internal))
//line namespace/template/form.qtpl:158
	qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:159
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 0.984l9 4.031v6c0 5.531-3.844 10.734-9 12-5.156-1.266-9-6.469-9-12v-6zM12 12v8.953c3.703-1.172 6.469-4.828 6.984-8.953h-6.984zM12 12v-8.813l-6.984 3.094v5.719h6.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:159
	qw422016.N().S(` <div class="form-option-info"> <strong>Internal</strong> <div class="form-desc">Anyone with an account will be able to view builds in the namespace.</div> </div> </label> <label class="form-option `)
//line namespace/template/form.qtpl:165
	qw422016.E().S(p.disabled(namespace.Public))
//line namespace/template/form.qtpl:165
	qw422016.N().S(`"> <input class="form-selector" type="radio" name="visibility" value="public" `)
//line namespace/template/form.qtpl:166
	qw422016.E().S(p.checked(namespace.Public))
//line namespace/template/form.qtpl:166
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:166
	qw422016.E().S(p.disabled(namespace.Public))
//line namespace/template/form.qtpl:166
	qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:167
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M17.906 17.391c1.313-1.406 2.109-3.328 2.109-5.391 0-3.328-2.063-6.234-5.016-7.406v0.422c0 1.078-0.938 1.969-2.016 1.969h-1.969v2.016c0 0.563-0.469 0.984-1.031 0.984h-1.969v2.016h6c0.563 0 0.984 0.422 0.984 0.984v3h0.984c0.891 0 1.641 0.609 1.922 1.406zM11.016 19.922v-1.922c-1.078 0-2.016-0.938-2.016-2.016v-0.984l-4.781-4.781c-0.141 0.563-0.234 1.172-0.234 1.781 0 4.078 3.094 7.453 7.031 7.922zM12 2.016c5.531 0 9.984 4.453 9.984 9.984s-4.453 9.984-9.984 9.984-9.984-4.453-9.984-9.984 4.453-9.984 9.984-9.984z"></path>
</svg>
`)
//line namespace/template/form.qtpl:167
	qw422016.N().S(` <div class="form-option-info"> <strong>Public</strong> <div class="form-desc">Anyone will be able to view builds in the namespace.</div> </div> </label> </div> <div class="form-field"> `)
//line namespace/template/form.qtpl:175
	if p.Namespace != nil {
//line namespace/template/form.qtpl:175
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line namespace/template/form.qtpl:177
	} else {
//line namespace/template/form.qtpl:177
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line namespace/template/form.qtpl:179
	}
//line namespace/template/form.qtpl:179
	qw422016.N().S(` </div> </form> `)
//line namespace/template/form.qtpl:182
	if p.Namespace != nil {
//line namespace/template/form.qtpl:182
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="`)
//line namespace/template/form.qtpl:184
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:184
		qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:185
		qw422016.N().S(string(p.CSRF))
//line namespace/template/form.qtpl:185
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <div class="overflow"> <div class="right"> <button type="submit" class="btn btn-danger">Delete</button> </div> <strong>Delete Namespace</strong><br/><p>Builds within the namespace will not be deleted.</p> </div> </form> `)
//line namespace/template/form.qtpl:194
	}
//line namespace/template/form.qtpl:194
	qw422016.N().S(` </div> </div> `)
//line namespace/template/form.qtpl:197
}

//line namespace/template/form.qtpl:197
func (p *Form) WriteBody(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:197
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:197
	p.StreamBody(qw422016)
//line namespace/template/form.qtpl:197
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:197
}

//line namespace/template/form.qtpl:197
func (p *Form) Body() string {
//line namespace/template/form.qtpl:197
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:197
	p.WriteBody(qb422016)
//line namespace/template/form.qtpl:197
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:197
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:197
	return qs422016
//line namespace/template/form.qtpl:197
}

//line namespace/template/form.qtpl:199
func (p *Form) StreamHeader(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:199
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:200
	if p.Namespace != nil {
//line namespace/template/form.qtpl:200
		qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:201
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:201
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:201
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:201
		qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:202
		streamrenderPath(qw422016, p.Namespace.User.Username, p.Namespace.Path)
//line namespace/template/form.qtpl:202
		qw422016.N().S(` - Edit `)
//line namespace/template/form.qtpl:203
	} else {
//line namespace/template/form.qtpl:203
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:204
		if p.Parent != nil {
//line namespace/template/form.qtpl:204
			qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:205
			qw422016.E().S(p.Parent.Endpoint())
//line namespace/template/form.qtpl:205
			qw422016.N().S(`">`)
//line namespace/template/form.qtpl:205
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:205
			qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:206
			streamrenderPath(qw422016, p.Parent.User.Username, p.Parent.Path)
//line namespace/template/form.qtpl:206
			qw422016.N().S(` - Create Sub-namespace `)
//line namespace/template/form.qtpl:207
		} else {
//line namespace/template/form.qtpl:207
			qw422016.N().S(` <a class="back" href="/namespaces">`)
//line namespace/template/form.qtpl:208
			qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:208
			qw422016.N().S(`</a> Create Namespace `)
//line namespace/template/form.qtpl:209
		}
//line namespace/template/form.qtpl:209
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:210
	}
//line namespace/template/form.qtpl:210
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:211
}

//line namespace/template/form.qtpl:211
func (p *Form) WriteHeader(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:211
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:211
	p.StreamHeader(qw422016)
//line namespace/template/form.qtpl:211
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:211
}

//line namespace/template/form.qtpl:211
func (p *Form) Header() string {
//line namespace/template/form.qtpl:211
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:211
	p.WriteHeader(qb422016)
//line namespace/template/form.qtpl:211
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:211
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:211
	return qs422016
//line namespace/template/form.qtpl:211
}

//line namespace/template/form.qtpl:213
func (p *Form) StreamActions(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:213
}

//line namespace/template/form.qtpl:213
func (p *Form) WriteActions(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:213
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:213
	p.StreamActions(qw422016)
//line namespace/template/form.qtpl:213
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:213
}

//line namespace/template/form.qtpl:213
func (p *Form) Actions() string {
//line namespace/template/form.qtpl:213
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:213
	p.WriteActions(qb422016)
//line namespace/template/form.qtpl:213
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:213
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:213
	return qs422016
//line namespace/template/form.qtpl:213
}

//line namespace/template/form.qtpl:214
func (p *Form) StreamNavigation(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:214
}

//line namespace/template/form.qtpl:214
func (p *Form) WriteNavigation(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:214
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:214
	p.StreamNavigation(qw422016)
//line namespace/template/form.qtpl:214
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:214
}

//line namespace/template/form.qtpl:214
func (p *Form) Navigation() string {
//line namespace/template/form.qtpl:214
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:214
	p.WriteNavigation(qb422016)
//line namespace/template/form.qtpl:214
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:214
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:214
	return qs422016
//line namespace/template/form.qtpl:214
}

//line namespace/template/form.qtpl:216
func (p *WebhookForm) StreamTitle(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:216
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:217
	if p.Webhook == nil {
//line namespace/template/form.qtpl:217
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:218
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:218
		qw422016.N().S(` - Create webhook - Djinn CI `)
//line namespace/template/form.qtpl:219
	} else {
//line namespace/template/form.qtpl:219
		qw422016.N().S(` `)
//line namespace/template/form.qtpl:220
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:220
		qw422016.N().S(` - Edit webhook - Djinn CI `)
//line namespace/template/form.qtpl:221
	}
//line namespace/template/form.qtpl:221
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:222
}

//line namespace/template/form.qtpl:222
func (p *WebhookForm) WriteTitle(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:222
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:222
	p.StreamTitle(qw422016)
//line namespace/template/form.qtpl:222
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:222
}

//line namespace/template/form.qtpl:222
func (p *WebhookForm) Title() string {
//line namespace/template/form.qtpl:222
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:222
	p.WriteTitle(qb422016)
//line namespace/template/form.qtpl:222
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:222
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:222
	return qs422016
//line namespace/template/form.qtpl:222
}

//line namespace/template/form.qtpl:224
func (p *WebhookForm) StreamHeader(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:224
	qw422016.N().S(` <a class="back" href="`)
//line namespace/template/form.qtpl:225
	qw422016.E().S(p.Namespace.Endpoint("webhooks"))
//line namespace/template/form.qtpl:225
	qw422016.N().S(`">`)
//line namespace/template/form.qtpl:225
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/form.qtpl:225
	qw422016.N().S(`</a> `)
//line namespace/template/form.qtpl:226
	if p.Webhook == nil {
//line namespace/template/form.qtpl:226
		qw422016.N().S(` <a href="`)
//line namespace/template/form.qtpl:227
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:227
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:227
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:227
		qw422016.N().S(`</a> / Create webhook `)
//line namespace/template/form.qtpl:228
	} else {
//line namespace/template/form.qtpl:228
		qw422016.N().S(` <a href="`)
//line namespace/template/form.qtpl:229
		qw422016.E().S(p.Namespace.Endpoint())
//line namespace/template/form.qtpl:229
		qw422016.N().S(`">`)
//line namespace/template/form.qtpl:229
		qw422016.E().S(p.Namespace.Name)
//line namespace/template/form.qtpl:229
		qw422016.N().S(`</a> / Edit webhook `)
//line namespace/template/form.qtpl:230
	}
//line namespace/template/form.qtpl:230
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:231
}

//line namespace/template/form.qtpl:231
func (p *WebhookForm) WriteHeader(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:231
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:231
	p.StreamHeader(qw422016)
//line namespace/template/form.qtpl:231
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:231
}

//line namespace/template/form.qtpl:231
func (p *WebhookForm) Header() string {
//line namespace/template/form.qtpl:231
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:231
	p.WriteHeader(qb422016)
//line namespace/template/form.qtpl:231
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:231
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:231
	return qs422016
//line namespace/template/form.qtpl:231
}

//line namespace/template/form.qtpl:233
func (p *WebhookForm) StreamBody(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:233
	qw422016.N().S(` <div class="panel"> <div class="panel-body slim"> <form method="POST" action="`)
//line namespace/template/form.qtpl:236
	qw422016.E().S(p.action())
//line namespace/template/form.qtpl:236
	qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:237
	if p.Webhook != nil {
//line namespace/template/form.qtpl:237
		qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> `)
//line namespace/template/form.qtpl:239
	}
//line namespace/template/form.qtpl:239
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:240
	qw422016.N().V(p.CSRF)
//line namespace/template/form.qtpl:240
	qw422016.N().S(` <div class="form-field"> <label class="label" for="payload_url">Payload URL</label> <input type="text" class="form-text" id="payload_url" name="payload_url" placeholder="https://example.com" autocomplete="off" value="`)
//line namespace/template/form.qtpl:243
	qw422016.E().S(p.Field("payload_url"))
//line namespace/template/form.qtpl:243
	qw422016.N().S(`"/> `)
//line namespace/template/form.qtpl:244
	p.StreamError(qw422016, "payload_url")
//line namespace/template/form.qtpl:244
	qw422016.N().S(` </div> <div class="form-field"> <label class="label" for="secret">Secret</label> <input type="password" class="form-text" id="secret" name="secret" autocomplete="off"/> </div> `)
//line namespace/template/form.qtpl:250
	if p.Webhook != nil {
//line namespace/template/form.qtpl:250
		qw422016.N().S(` <div class="form-field"> <label class="form-option"> <input class="form-selector" type="checkbox" name="remove_secret"/> <div class="form-option-info"> <strong>Remove secret</strong> <div class="form-desc">Leave the secret blank to remove it</div> </div> </label> </div> `)
//line namespace/template/form.qtpl:260
	}
//line namespace/template/form.qtpl:260
	qw422016.N().S(` <div class="form-field"> <label class="form-option"> `)
//line namespace/template/form.qtpl:263
	if p.Webhook != nil {
//line namespace/template/form.qtpl:263
		qw422016.N().S(` <input class="form-selector" type="checkbox" name="ssl" `)
//line namespace/template/form.qtpl:264
		if p.Webhook.SSL {
//line namespace/template/form.qtpl:264
			qw422016.N().S(`checked="true"`)
//line namespace/template/form.qtpl:264
		}
//line namespace/template/form.qtpl:264
		qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:265
	} else {
//line namespace/template/form.qtpl:265
		qw422016.N().S(` <input class="form-selector" type="checkbox" name="ssl" checked="true"/> `)
//line namespace/template/form.qtpl:267
	}
//line namespace/template/form.qtpl:267
	qw422016.N().S(` <div class="form-option-info"> <strong>SSL</strong> <div class="form-desc">Verify SSL certificates when sending payloads</div> </div> </label> </div> <div class="form-field"> <label class="form-option"> `)
//line namespace/template/form.qtpl:276
	if p.Webhook != nil {
//line namespace/template/form.qtpl:276
		qw422016.N().S(` <input class="form-selector" type="checkbox" name="active" `)
//line namespace/template/form.qtpl:277
		if p.Webhook.Active {
//line namespace/template/form.qtpl:277
			qw422016.N().S(`checked="true"`)
//line namespace/template/form.qtpl:277
		}
//line namespace/template/form.qtpl:277
		qw422016.N().S(`/> `)
//line namespace/template/form.qtpl:278
	} else {
//line namespace/template/form.qtpl:278
		qw422016.N().S(` <input class="form-selector" type="checkbox" name="active" checked="true"/> `)
//line namespace/template/form.qtpl:280
	}
//line namespace/template/form.qtpl:280
	qw422016.N().S(` <div class="form-option-info"> <strong>Active</strong> <div class="form-desc">Send events for this webhook</div> </div> </label> </div> <div class="form-field"> <strong>Events</strong><br/> `)
//line namespace/template/form.qtpl:289
	for _, event := range event.Types {
//line namespace/template/form.qtpl:289
		qw422016.N().S(` <label class="hook-event"> `)
//line namespace/template/form.qtpl:291
		if p.Webhook != nil && p.Webhook.Events.Has(event) {
//line namespace/template/form.qtpl:291
			qw422016.N().S(` <input checked="true" class="form-selector" type="checkbox" name="events[]" value="`)
//line namespace/template/form.qtpl:292
			qw422016.E().S(event.String())
//line namespace/template/form.qtpl:292
			qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:292
			qw422016.E().S(strings.Replace(strings.Title(event.String()), "_", " ", -1))
//line namespace/template/form.qtpl:292
			qw422016.N().S(` `)
//line namespace/template/form.qtpl:293
		} else {
//line namespace/template/form.qtpl:293
			qw422016.N().S(` <input class="form-selector" type="checkbox" name="events[]" value="`)
//line namespace/template/form.qtpl:294
			qw422016.E().S(event.String())
//line namespace/template/form.qtpl:294
			qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:294
			qw422016.E().S(strings.Replace(strings.Title(event.String()), "_", " ", -1))
//line namespace/template/form.qtpl:294
			qw422016.N().S(` `)
//line namespace/template/form.qtpl:295
		}
//line namespace/template/form.qtpl:295
		qw422016.N().S(` </label> `)
//line namespace/template/form.qtpl:297
	}
//line namespace/template/form.qtpl:297
	qw422016.N().S(` </div> <div class="form-field"> `)
//line namespace/template/form.qtpl:300
	if p.Webhook == nil {
//line namespace/template/form.qtpl:300
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Create</button> `)
//line namespace/template/form.qtpl:302
	} else {
//line namespace/template/form.qtpl:302
		qw422016.N().S(` <button type="submit" class="btn btn-primary">Save</button> `)
//line namespace/template/form.qtpl:304
	}
//line namespace/template/form.qtpl:304
	qw422016.N().S(` </div> </form> `)
//line namespace/template/form.qtpl:307
	if p.Webhook != nil {
//line namespace/template/form.qtpl:307
		qw422016.N().S(` <div class="separator"></div> <form method="POST" action="`)
//line namespace/template/form.qtpl:309
		qw422016.E().S(p.Webhook.Endpoint())
//line namespace/template/form.qtpl:309
		qw422016.N().S(`"> `)
//line namespace/template/form.qtpl:310
		qw422016.N().V(p.CSRF)
//line namespace/template/form.qtpl:310
		qw422016.N().S(` <input type="hidden" name="_method" value="DELETE"/> <div class="overflow"> <div class="right"> <button class="btn btn-danger" type="submit">Delete</button> </div> </div> </form> `)
//line namespace/template/form.qtpl:318
	}
//line namespace/template/form.qtpl:318
	qw422016.N().S(` </div> </div> `)
//line namespace/template/form.qtpl:321
	if p.Webhook != nil {
//line namespace/template/form.qtpl:321
		qw422016.N().S(` <div class="panel"> <div class="panel-header"> <h3>Recent deliveries</h3> </div> `)
//line namespace/template/form.qtpl:326
		if len(p.Deliveries) == 0 {
//line namespace/template/form.qtpl:326
			qw422016.N().S(` <div class="panel-message muted">No recent deliveries.</div> `)
//line namespace/template/form.qtpl:328
		} else {
//line namespace/template/form.qtpl:328
			qw422016.N().S(` <table class="table"> <tbody> `)
//line namespace/template/form.qtpl:331
			for _, d := range p.Deliveries {
//line namespace/template/form.qtpl:331
				qw422016.N().S(` <tr> `)
//line namespace/template/form.qtpl:333
				if d.ResponseCode.Int64 >= 200 && d.ResponseCode.Int64 < 300 {
//line namespace/template/form.qtpl:333
					qw422016.N().S(` <td class="hook-status hook-status-ok">`)
//line namespace/template/form.qtpl:334
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
//line namespace/template/form.qtpl:334
					qw422016.N().S(`</td> `)
//line namespace/template/form.qtpl:335
				} else {
//line namespace/template/form.qtpl:335
					qw422016.N().S(` <td class="hook-status hook-status-err">`)
//line namespace/template/form.qtpl:336
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
//line namespace/template/form.qtpl:336
					qw422016.N().S(`</td> `)
//line namespace/template/form.qtpl:337
				}
//line namespace/template/form.qtpl:337
				qw422016.N().S(` <td>`)
//line namespace/template/form.qtpl:338
				if d.Redelivery {
//line namespace/template/form.qtpl:338
					qw422016.N().S(`<span class="muted" title="Redelivery">`)
//line namespace/template/form.qtpl:338
					qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 18v-3l3.984 3.984-3.984 4.031v-3c-4.406 0-8.016-3.609-8.016-8.016 0-1.547 0.469-3.047 1.266-4.266l1.453 1.453c-0.469 0.844-0.703 1.781-0.703 2.813 0 3.328 2.672 6 6 6zM12 3.984c4.406 0 8.016 3.609 8.016 8.016 0 1.547-0.469 3.047-1.266 4.266l-1.453-1.453c0.469-0.844 0.703-1.781 0.703-2.813 0-3.328-2.672-6-6-6v3l-3.984-3.984 3.984-4.031v3z"></path>
</svg>
`)
//line namespace/template/form.qtpl:338
					qw422016.N().S(`</span>`)
//line namespace/template/form.qtpl:338
				}
//line namespace/template/form.qtpl:338
				qw422016.N().S(` <code><a href="`)
//line namespace/template/form.qtpl:338
				qw422016.E().S(d.Endpoint())
//line namespace/template/form.qtpl:338
				qw422016.N().S(`">`)
//line namespace/template/form.qtpl:338
				qw422016.E().V(d.EventID)
//line namespace/template/form.qtpl:338
				qw422016.N().S(`</a></code></td> <td>`)
//line namespace/template/form.qtpl:339
				qw422016.E().S(d.Duration.String())
//line namespace/template/form.qtpl:339
				qw422016.N().S(`</td> </tr> `)
//line namespace/template/form.qtpl:341
			}
//line namespace/template/form.qtpl:341
			qw422016.N().S(` </tbody> </table> `)
//line namespace/template/form.qtpl:344
		}
//line namespace/template/form.qtpl:344
		qw422016.N().S(` </div> `)
//line namespace/template/form.qtpl:346
	}
//line namespace/template/form.qtpl:346
	qw422016.N().S(` `)
//line namespace/template/form.qtpl:347
}

//line namespace/template/form.qtpl:347
func (p *WebhookForm) WriteBody(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:347
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:347
	p.StreamBody(qw422016)
//line namespace/template/form.qtpl:347
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:347
}

//line namespace/template/form.qtpl:347
func (p *WebhookForm) Body() string {
//line namespace/template/form.qtpl:347
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:347
	p.WriteBody(qb422016)
//line namespace/template/form.qtpl:347
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:347
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:347
	return qs422016
//line namespace/template/form.qtpl:347
}

//line namespace/template/form.qtpl:349
func (p *WebhookForm) StreamActions(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:349
}

//line namespace/template/form.qtpl:349
func (p *WebhookForm) WriteActions(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:349
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:349
	p.StreamActions(qw422016)
//line namespace/template/form.qtpl:349
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:349
}

//line namespace/template/form.qtpl:349
func (p *WebhookForm) Actions() string {
//line namespace/template/form.qtpl:349
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:349
	p.WriteActions(qb422016)
//line namespace/template/form.qtpl:349
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:349
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:349
	return qs422016
//line namespace/template/form.qtpl:349
}

//line namespace/template/form.qtpl:350
func (p *WebhookForm) StreamNavigation(qw422016 *qt422016.Writer) {
//line namespace/template/form.qtpl:350
}

//line namespace/template/form.qtpl:350
func (p *WebhookForm) WriteNavigation(qq422016 qtio422016.Writer) {
//line namespace/template/form.qtpl:350
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/form.qtpl:350
	p.StreamNavigation(qw422016)
//line namespace/template/form.qtpl:350
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/form.qtpl:350
}

//line namespace/template/form.qtpl:350
func (p *WebhookForm) Navigation() string {
//line namespace/template/form.qtpl:350
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/form.qtpl:350
	p.WriteNavigation(qb422016)
//line namespace/template/form.qtpl:350
	qs422016 := string(qb422016.B)
//line namespace/template/form.qtpl:350
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/form.qtpl:350
	return qs422016
//line namespace/template/form.qtpl:350
}

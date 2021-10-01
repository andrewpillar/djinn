// Code generated by qtc from "delivery.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line namespace/template/delivery.qtpl:2
package template

//line namespace/template/delivery.qtpl:2
import (
	htmltemplate "html/template"

	"djinn-ci.com/namespace"
	"djinn-ci.com/template"
)

//line namespace/template/delivery.qtpl:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line namespace/template/delivery.qtpl:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line namespace/template/delivery.qtpl:11
type ShowDelivery struct {
	template.BasePage

	CSRF     htmltemplate.HTML
	Delivery *namespace.WebhookDelivery
}

//line namespace/template/delivery.qtpl:20
func (p *ShowDelivery) StreamTitle(qw422016 *qt422016.Writer) {
//line namespace/template/delivery.qtpl:20
	qw422016.N().S(` `)
//line namespace/template/delivery.qtpl:21
	qw422016.E().S(p.Delivery.Webhook.Namespace.Name)
//line namespace/template/delivery.qtpl:21
	qw422016.N().S(` - Webhook Delivery `)
//line namespace/template/delivery.qtpl:22
}

//line namespace/template/delivery.qtpl:22
func (p *ShowDelivery) WriteTitle(qq422016 qtio422016.Writer) {
//line namespace/template/delivery.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/delivery.qtpl:22
	p.StreamTitle(qw422016)
//line namespace/template/delivery.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/delivery.qtpl:22
}

//line namespace/template/delivery.qtpl:22
func (p *ShowDelivery) Title() string {
//line namespace/template/delivery.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/delivery.qtpl:22
	p.WriteTitle(qb422016)
//line namespace/template/delivery.qtpl:22
	qs422016 := string(qb422016.B)
//line namespace/template/delivery.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/delivery.qtpl:22
	return qs422016
//line namespace/template/delivery.qtpl:22
}

//line namespace/template/delivery.qtpl:24
func (p *ShowDelivery) StreamBody(qw422016 *qt422016.Writer) {
//line namespace/template/delivery.qtpl:24
	qw422016.N().S(` `)
//line namespace/template/delivery.qtpl:25
	if p.Delivery.Error.Valid {
//line namespace/template/delivery.qtpl:25
		qw422016.N().S(` <div class="alert alert-danger mb-10"> Failed to deliver payload: `)
//line namespace/template/delivery.qtpl:27
		qw422016.E().S(p.Delivery.Error.String)
//line namespace/template/delivery.qtpl:27
		qw422016.N().S(` </div> `)
//line namespace/template/delivery.qtpl:29
	}
//line namespace/template/delivery.qtpl:29
	qw422016.N().S(` <div class="overflow"> <div class="col-50 col-left"> <div class="panel"> <div class="panel-header"><h3>Request Headers</h3></div> `)
//line namespace/template/delivery.qtpl:34
	if !p.Delivery.RequestHeaders.Valid {
//line namespace/template/delivery.qtpl:34
		qw422016.N().S(` <div class="panel-message muted">None</div> `)
//line namespace/template/delivery.qtpl:36
	} else {
//line namespace/template/delivery.qtpl:36
		qw422016.N().S(` <div class="panel-body"> <pre class="code">`)
//line namespace/template/delivery.qtpl:38
		qw422016.E().S(p.Delivery.RequestHeaders.String)
//line namespace/template/delivery.qtpl:38
		qw422016.N().S(`</pre> </div> `)
//line namespace/template/delivery.qtpl:40
	}
//line namespace/template/delivery.qtpl:40
	qw422016.N().S(` </div> <div class="panel"> <div class="panel-header"><h3>Request Body</h3></div> `)
//line namespace/template/delivery.qtpl:44
	if !p.Delivery.RequestHeaders.Valid {
//line namespace/template/delivery.qtpl:44
		qw422016.N().S(` <div class="panel-message muted">None</div> `)
//line namespace/template/delivery.qtpl:46
	} else {
//line namespace/template/delivery.qtpl:46
		qw422016.N().S(` <div class="panel-body"> <pre class="code">`)
//line namespace/template/delivery.qtpl:48
		qw422016.E().S(template.IndentJSON(p.Delivery.RequestBody.String))
//line namespace/template/delivery.qtpl:48
		qw422016.N().S(`</pre> </div> `)
//line namespace/template/delivery.qtpl:50
	}
//line namespace/template/delivery.qtpl:50
	qw422016.N().S(` </div> </div> <div class="col-50 col-right"> `)
//line namespace/template/delivery.qtpl:54
	if !p.Delivery.Error.Valid {
//line namespace/template/delivery.qtpl:54
		qw422016.N().S(` <div class="panel"> <table class="table"> <tr> <td><strong>Status</strong></td> <td class="right"><code class="code">`)
//line namespace/template/delivery.qtpl:59
		qw422016.E().S(p.Delivery.ResponseStatus)
//line namespace/template/delivery.qtpl:59
		qw422016.N().S(`</code></td> </tr> </table> </div> `)
//line namespace/template/delivery.qtpl:63
	}
//line namespace/template/delivery.qtpl:63
	qw422016.N().S(` <div class="panel"> <div class="panel-header"><h3>Response Headers</h3></div> `)
//line namespace/template/delivery.qtpl:66
	if !p.Delivery.ResponseHeaders.Valid {
//line namespace/template/delivery.qtpl:66
		qw422016.N().S(` <div class="panel-message muted">None</div> `)
//line namespace/template/delivery.qtpl:68
	} else {
//line namespace/template/delivery.qtpl:68
		qw422016.N().S(` <div class="panel-body"> <pre class="code">`)
//line namespace/template/delivery.qtpl:70
		qw422016.E().S(p.Delivery.ResponseHeaders.String)
//line namespace/template/delivery.qtpl:70
		qw422016.N().S(`</pre> </div> `)
//line namespace/template/delivery.qtpl:72
	}
//line namespace/template/delivery.qtpl:72
	qw422016.N().S(` </div> <div class="panel"> <div class="panel-header"><h3>Response Body</h3></div> `)
//line namespace/template/delivery.qtpl:76
	if !p.Delivery.ResponseBody.Valid {
//line namespace/template/delivery.qtpl:76
		qw422016.N().S(` <div class="panel-message muted">None</div> `)
//line namespace/template/delivery.qtpl:78
	} else {
//line namespace/template/delivery.qtpl:78
		qw422016.N().S(` <div class="panel-body"> <pre class="code">`)
//line namespace/template/delivery.qtpl:80
		qw422016.E().S(p.Delivery.ResponseBody.String)
//line namespace/template/delivery.qtpl:80
		qw422016.N().S(`</pre> </div> `)
//line namespace/template/delivery.qtpl:82
	}
//line namespace/template/delivery.qtpl:82
	qw422016.N().S(` </div> </div> </div> `)
//line namespace/template/delivery.qtpl:86
}

//line namespace/template/delivery.qtpl:86
func (p *ShowDelivery) WriteBody(qq422016 qtio422016.Writer) {
//line namespace/template/delivery.qtpl:86
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/delivery.qtpl:86
	p.StreamBody(qw422016)
//line namespace/template/delivery.qtpl:86
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/delivery.qtpl:86
}

//line namespace/template/delivery.qtpl:86
func (p *ShowDelivery) Body() string {
//line namespace/template/delivery.qtpl:86
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/delivery.qtpl:86
	p.WriteBody(qb422016)
//line namespace/template/delivery.qtpl:86
	qs422016 := string(qb422016.B)
//line namespace/template/delivery.qtpl:86
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/delivery.qtpl:86
	return qs422016
//line namespace/template/delivery.qtpl:86
}

//line namespace/template/delivery.qtpl:88
func (p *ShowDelivery) StreamHeader(qw422016 *qt422016.Writer) {
//line namespace/template/delivery.qtpl:88
	qw422016.N().S(` <a class="back" href="`)
//line namespace/template/delivery.qtpl:89
	qw422016.E().S(p.Delivery.Webhook.Endpoint())
//line namespace/template/delivery.qtpl:89
	qw422016.N().S(`">`)
//line namespace/template/delivery.qtpl:89
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M20.016 11.016v1.969h-12.188l5.578 5.625-1.406 1.406-8.016-8.016 8.016-8.016 1.406 1.406-5.578 5.625h12.188z"></path>
</svg>
`)
//line namespace/template/delivery.qtpl:89
	qw422016.N().S(`</a> `)
//line namespace/template/delivery.qtpl:90
	if p.Delivery.ResponseCode.Int64 >= 200 && p.Delivery.ResponseCode.Int64 < 300 {
//line namespace/template/delivery.qtpl:90
		qw422016.N().S(` <span class="hook-status hook-status-ok">`)
//line namespace/template/delivery.qtpl:91
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
//line namespace/template/delivery.qtpl:91
		qw422016.N().S(`</span> `)
//line namespace/template/delivery.qtpl:92
	} else {
//line namespace/template/delivery.qtpl:92
		qw422016.N().S(` <span class="hook-status hook-status-err">`)
//line namespace/template/delivery.qtpl:93
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
//line namespace/template/delivery.qtpl:93
		qw422016.N().S(`</span> `)
//line namespace/template/delivery.qtpl:94
	}
//line namespace/template/delivery.qtpl:94
	qw422016.N().S(` <code>`)
//line namespace/template/delivery.qtpl:95
	qw422016.E().S(p.Delivery.EventID.String())
//line namespace/template/delivery.qtpl:95
	qw422016.N().S(`</code> `)
//line namespace/template/delivery.qtpl:96
	if p.Delivery.Redelivery {
//line namespace/template/delivery.qtpl:96
		qw422016.N().S(` <span class="muted" title="Redelivery">`)
//line namespace/template/delivery.qtpl:97
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 18v-3l3.984 3.984-3.984 4.031v-3c-4.406 0-8.016-3.609-8.016-8.016 0-1.547 0.469-3.047 1.266-4.266l1.453 1.453c-0.469 0.844-0.703 1.781-0.703 2.813 0 3.328 2.672 6 6 6zM12 3.984c4.406 0 8.016 3.609 8.016 8.016 0 1.547-0.469 3.047-1.266 4.266l-1.453-1.453c0.469-0.844 0.703-1.781 0.703-2.813 0-3.328-2.672-6-6-6v3l-3.984-3.984 3.984-4.031v3z"></path>
</svg>
`)
//line namespace/template/delivery.qtpl:97
		qw422016.N().S(`</span> `)
//line namespace/template/delivery.qtpl:98
	}
//line namespace/template/delivery.qtpl:98
	qw422016.N().S(` `)
//line namespace/template/delivery.qtpl:99
}

//line namespace/template/delivery.qtpl:99
func (p *ShowDelivery) WriteHeader(qq422016 qtio422016.Writer) {
//line namespace/template/delivery.qtpl:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/delivery.qtpl:99
	p.StreamHeader(qw422016)
//line namespace/template/delivery.qtpl:99
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/delivery.qtpl:99
}

//line namespace/template/delivery.qtpl:99
func (p *ShowDelivery) Header() string {
//line namespace/template/delivery.qtpl:99
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/delivery.qtpl:99
	p.WriteHeader(qb422016)
//line namespace/template/delivery.qtpl:99
	qs422016 := string(qb422016.B)
//line namespace/template/delivery.qtpl:99
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/delivery.qtpl:99
	return qs422016
//line namespace/template/delivery.qtpl:99
}

//line namespace/template/delivery.qtpl:101
func (p *ShowDelivery) StreamNavigation(qw422016 *qt422016.Writer) {
//line namespace/template/delivery.qtpl:101
}

//line namespace/template/delivery.qtpl:101
func (p *ShowDelivery) WriteNavigation(qq422016 qtio422016.Writer) {
//line namespace/template/delivery.qtpl:101
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/delivery.qtpl:101
	p.StreamNavigation(qw422016)
//line namespace/template/delivery.qtpl:101
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/delivery.qtpl:101
}

//line namespace/template/delivery.qtpl:101
func (p *ShowDelivery) Navigation() string {
//line namespace/template/delivery.qtpl:101
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/delivery.qtpl:101
	p.WriteNavigation(qb422016)
//line namespace/template/delivery.qtpl:101
	qs422016 := string(qb422016.B)
//line namespace/template/delivery.qtpl:101
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/delivery.qtpl:101
	return qs422016
//line namespace/template/delivery.qtpl:101
}

//line namespace/template/delivery.qtpl:103
func (p *ShowDelivery) StreamActions(qw422016 *qt422016.Writer) {
//line namespace/template/delivery.qtpl:103
	qw422016.N().S(` <li> <form method="POST" action="`)
//line namespace/template/delivery.qtpl:105
	qw422016.E().S(p.Delivery.Endpoint())
//line namespace/template/delivery.qtpl:105
	qw422016.N().S(`"> `)
//line namespace/template/delivery.qtpl:106
	qw422016.N().V(p.CSRF)
//line namespace/template/delivery.qtpl:106
	qw422016.N().S(` <input type="hidden" name="_method" value="PATCH"/> <button class="btn btn-primary">Redeliver</button> </form> </li> `)
//line namespace/template/delivery.qtpl:111
}

//line namespace/template/delivery.qtpl:111
func (p *ShowDelivery) WriteActions(qq422016 qtio422016.Writer) {
//line namespace/template/delivery.qtpl:111
	qw422016 := qt422016.AcquireWriter(qq422016)
//line namespace/template/delivery.qtpl:111
	p.StreamActions(qw422016)
//line namespace/template/delivery.qtpl:111
	qt422016.ReleaseWriter(qw422016)
//line namespace/template/delivery.qtpl:111
}

//line namespace/template/delivery.qtpl:111
func (p *ShowDelivery) Actions() string {
//line namespace/template/delivery.qtpl:111
	qb422016 := qt422016.AcquireByteBuffer()
//line namespace/template/delivery.qtpl:111
	p.WriteActions(qb422016)
//line namespace/template/delivery.qtpl:111
	qs422016 := string(qb422016.B)
//line namespace/template/delivery.qtpl:111
	qt422016.ReleaseByteBuffer(qb422016)
//line namespace/template/delivery.qtpl:111
	return qs422016
//line namespace/template/delivery.qtpl:111
}

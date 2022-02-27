// Code generated by qtc from "template.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/template.qtpl:2
package template

//line template/template.qtpl:2
import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/url"
	"regexp"
	"strconv"

	"djinn-ci.com/database"
	"djinn-ci.com/user"

	"github.com/andrewpillar/webutil"
)

//line template/template.qtpl:18
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/template.qtpl:18
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/template.qtpl:18
type Page interface {
//line template/template.qtpl:18
	Title() string
//line template/template.qtpl:18
	StreamTitle(qw422016 *qt422016.Writer)
//line template/template.qtpl:18
	WriteTitle(qq422016 qtio422016.Writer)
//line template/template.qtpl:18
	Body() string
//line template/template.qtpl:18
	StreamBody(qw422016 *qt422016.Writer)
//line template/template.qtpl:18
	WriteBody(qq422016 qtio422016.Writer)
//line template/template.qtpl:18
	Footer() string
//line template/template.qtpl:18
	StreamFooter(qw422016 *qt422016.Writer)
//line template/template.qtpl:18
	WriteFooter(qq422016 qtio422016.Writer)
//line template/template.qtpl:18
}

//line template/template.qtpl:28
type BasePage struct {
	URL  *url.URL
	User *user.User
}

type Form struct {
	CSRF   template.HTML
	Errors webutil.ValidationErrors
	Fields map[string]string
}

func pattern(name string) string { return "(^\\/" + name + "\\/?[a-z0-9\\/?]*$)" }

func LinkToQuery(url *url.URL, query map[string]string) string {
	q := url.Query()

	for k, v := range query {
		if v == "" {
			delete(q, k)
			continue
		}
		q.Set(k, v)
	}

	url2 := (*url)
	url2.RawQuery = q.Encode()

	return url2.String()
}

func Active(condition bool) string {
	if condition {
		return "active"
	}
	return ""
}

func Match(uri, pattern string) bool {
	matched, err := regexp.Match(pattern, []byte(uri))

	if err != nil {
		return false
	}
	return matched
}

func IndentJSON(s string) string {
	var buf bytes.Buffer

	json.Indent(&buf, []byte(s), "", "    ")
	return buf.String()
}

func pageURL(url *url.URL, page int64) string {
	q := url.Query()
	q.Set("page", strconv.FormatInt(page, 10))

	url.RawQuery = q.Encode()
	return url.String()
}

//line template/template.qtpl:91
func StreamFileInput(qw422016 *qt422016.Writer, f Form) {
//line template/template.qtpl:91
	qw422016.N().S(` <div class="form-field"> <label class="label" for="file">File</label> <input type="file" id="file" name="file"/> `)
//line template/template.qtpl:95
	f.StreamError(qw422016, "file")
//line template/template.qtpl:95
	qw422016.N().S(` </div> `)
//line template/template.qtpl:97
}

//line template/template.qtpl:97
func WriteFileInput(qq422016 qtio422016.Writer, f Form) {
//line template/template.qtpl:97
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:97
	StreamFileInput(qw422016, f)
//line template/template.qtpl:97
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:97
}

//line template/template.qtpl:97
func FileInput(f Form) string {
//line template/template.qtpl:97
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:97
	WriteFileInput(qb422016, f)
//line template/template.qtpl:97
	qs422016 := string(qb422016.B)
//line template/template.qtpl:97
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:97
	return qs422016
//line template/template.qtpl:97
}

//line template/template.qtpl:99
func StreamRender(qw422016 *qt422016.Writer, p Page) {
//line template/template.qtpl:99
	qw422016.N().S(` <!DOCTYPE HTML> <html lang="en"> <head> <meta charset="utf-8"> <meta content="width=device-width, initial-scale=1" name="viewport"> <title>`)
//line template/template.qtpl:105
	p.StreamTitle(qw422016)
//line template/template.qtpl:105
	qw422016.N().S(`</title> <style type="text/css">`)
//line template/template.qtpl:106
	qw422016.N().S(`*{margin:0;padding:0}body{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";font-size:14px;background:#eee;color:#444}a{color:#146de0;cursor:pointer;text-decoration:none}a:hover{text-decoration:underline}button{cursor:pointer}h1,h2,h3,h4,h5,h6{font-weight:400}.btn{border:none;border-radius:3px;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";padding:10px;padding-left:15px;padding-right:15px;color:#fff}.btn svg{fill:#fff}.btn:hover{text-decoration:none}.btn:disabled{cursor:not-allowed;background:#b2b2b2!important}.btn-primary{background:#61a0ea}.btn-primary:hover{background:#5090d9}.btn-danger{background:#de4141}.btn-danger:hover{background:#cd3030}.code{padding:3px;border-radius:3px;background:#e6f0f5;font-family:monospace;white-space:pre-wrap}.code-wrap{overflow-x:auto;overflow-y:hidden}table.code{background:#272b39;color:#fff;width:100%;font-family:monospace;font-size:12px;border-collapse:collapse;border-spacing:0}table.code .line-number{text-align:right;-moz-user-select:none;-ms-user-select:none;-webkit-user-select:none;min-width:30px;width:1%;padding-left:10px;padding-right:10px;line-height:20px}table.code .line-number a{display:block;color:rgba(255,255,255,.3)}table.code .line{padding-left:10px;padding-right:10px;line-height:20px;white-space:pre;overflow:visible;word-wrap:normal}table.code .line:target{background:#383e51}.col-75{width:75%;box-sizing:border-box}.col-25{width:25%;box-sizing:border-box}.col-50{width:50%;box-sizing:border-box}.col-left{float:left;padding-right:5px}.col-right{float:right;padding-left:5px}@media (max-width:1100px){.col-75{margin-bottom:10px;width:100%}.col-25{margin-bottom:10px;width:100%}.col-50{margin-bottom:10px;width:100%}.col-left{padding-right:0;float:none}.col-right{padding-left:0;float:none}}.dashboard .sidebar{position:fixed;top:0;left:0;height:100%;width:225px;background:#383e51;overflow:auto}.dashboard .sidebar .sidebar-header{color:#fff;padding:20px;background:#272b39}.dashboard .sidebar .sidebar-header .logo{margin-top:-5px;margin-right:30px;display:inline-block;vertical-align:middle;width:0}.dashboard .sidebar .sidebar-header .logo .handle{margin-left:-3px;border-style:solid;border-width:2px 0 8px 7px;border-color:transparent transparent transparent #fff}.dashboard .sidebar .sidebar-header .logo .lid{margin-bottom:-20px;margin-left:13px;border-style:solid;border-width:5px 0 7px 5px;border-color:transparent transparent transparent #fff}.dashboard .sidebar .sidebar-header .logo .lantern{margin-left:-5px;border-style:solid;border-width:15px 15px 35px 0;border-color:transparent #fff transparent transparent}.dashboard .sidebar .sidebar-header h2{display:inline-block}.dashboard .sidebar .sidebar-auth a{display:block;color:rgba(255,255,255,.5);padding:15px;text-align:center}.dashboard .sidebar .sidebar-auth a.active,.dashboard .sidebar .sidebar-auth a:hover,.dashboard .sidebar .sidebar-auth button:hover{text-decoration:none;background:#272b39;color:#fff}.dashboard .sidebar .sidebar-nav{list-style:none}.dashboard .sidebar .sidebar-nav li{display:block}.dashboard .sidebar .sidebar-nav li a,.dashboard .sidebar .sidebar-nav li button{display:block;color:rgba(255,255,255,.5);padding:15px}.dashboard .sidebar .sidebar-nav li a svg,.dashboard .sidebar .sidebar-nav li button svg{margin-right:3px;display:inline-block;vertical-align:middle;fill:rgba(255,255,255,.5);width:15px}.dashboard .sidebar .sidebar-nav li a span,.dashboard .sidebar .sidebar-nav li button span{margin-top:2px;display:inline-block;vertical-align:middle}.dashboard .sidebar .sidebar-nav li button{width:100%;border:none;text-align:left;background:rgba(0,0,0,0)}.dashboard .sidebar .sidebar-nav li a.active,.dashboard .sidebar .sidebar-nav li a:hover,.dashboard .sidebar .sidebar-nav li button:hover{text-decoration:none;background:#272b39;color:#fff}.dashboard .sidebar .sidebar-nav li a.active svg,.dashboard .sidebar .sidebar-nav li a:hover svg,.dashboard .sidebar .sidebar-nav li button:hover svg{fill:#fff}.dashboard .sidebar .sidebar-nav li.sidebar-nav-header{padding:15px;font-weight:700;color:#fff}.dashboard-header{margin-bottom:10px}.dashboard-header h1{float:left}.dashboard-header h1 .back{margin-top:2px;display:inline-block;vertical-align:middle}.dashboard-header h1 .back svg{fill:#7f7f7f}.dashboard-header h1 .back:hover{text-decoration:none}.dashboard-header h1 .back:hover svg{fill:#444}.dashboard-header h1 small{margin-top:10px;display:block;font-size:16px;color:rgba(0,0,0,.5)}.dashboard-header .pill{margin-top:-5px;margin-left:10px}.dashboard-header .dashboard-actions{float:right;list-style:none}.dashboard-header .dashboard-actions li{display:inline}.dashboard-header .dashboard-actions li form{display:inline-block}.dashboard-header .dashboard-actions li a{cursor:pointer;display:inline-block}.dashboard-nav{list-style:none}.dashboard-nav li{display:inline}.dashboard-nav li a{display:inline-block;padding:15px;color:#9f9f9f}.dashboard-nav li a svg{margin-right:3px;width:20px;vertical-align:middle;display:inline-block;fill:#9f9f9f}.dashboard-nav li a span{margin-top:2px;display:inline-block;vertical-align:middle}.dashboard-nav li a.active,.dashboard-nav li a:hover{text-decoration:none;color:#272b39}.dashboard-nav li a.active svg,.dashboard-nav li a:hover svg{fill:#272b39}.dashboard-content{margin-left:225px}.dashboard-content .alert{overflow:auto;padding:15px}.dashboard-content .alert .alert-message{float:left;color:rgba(0,0,0,.6)}.dashboard-content .alert a.alert-close{float:right;display:inline-block}.dashboard-content .alert a.alert-close svg{width:15px;height:15px;fill:rgba(0,0,0,.4)}.dashboard-content .alert a.alert-close:hover svg{fill:rgba(0,0,0,.5)}.dashboard-content .alert-success{background:#caf5ca;border:solid 1px #a0dfa0}.dashboard-content .alert-warn{background:#fff3cd;border:solid 1px #d9c995}.dashboard-content .alert-danger{background:#ffd4d4;border:solid 1px #e19e9e}.dashboard-content .dashboard-wrap{margin:0 auto;max-width:1300px;padding:20px}@media (max-width:1500px){.dashboard .sidebar{width:70px}.dashboard .sidebar .sidebar-header{padding:15px}.dashboard .sidebar .sidebar-header .logo{margin-top:0;margin-right:0;margin-left:12px}.dashboard .sidebar .sidebar-header h2{display:none}.dashboard .sidebar .sidebar-nav .sidebar-nav-header{display:none}.dashboard .sidebar .sidebar-nav li a,.dashboard .sidebar .sidebar-nav li button{text-align:center}.dashboard .sidebar .sidebar-nav li a span,.dashboard .sidebar .sidebar-nav li button span{display:none}.dashboard .dashboard-content{margin-left:70px}}@media (max-width:1000px){.dashboard .dashboard-content .dashboard-header .dashboard-nav li a span{display:none}}.form-field+.form-field{margin-top:15px}.form-field{overflow:auto}.form-field .label{margin-bottom:5px;display:block;font-weight:700}.form-field .label small{color:rgba(0,0,0,.5)}.form-field .form-error{margin-top:5px;color:#ff4343;min-height:20px}.form-field .form-text{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";font-size:14px;padding:10px;outline:0;border-radius:3px;box-sizing:border-box;width:100%;border:solid 1px #e4e4e4}.form-field .form-text:focus{border:solid 1px #c2c2c2}.form-field .form-code{min-height:250px;font-family:monospace}.form-field textarea.form-text{min-width:100%;max-width:100%}.form-field .form-option+.form-option{margin-top:10px}.form-field .form-option{display:block;cursor:pointer;overflow:auto}.form-field .form-option .form-selector{margin-right:5px;outline:0}.form-field .form-option .form-option-info{margin-right:5px;display:inline-block}.form-field .form-option svg{margin-right:5px;fill:rgba(0,0,0,.4)}.form-field .hook-event{cursor:pointer;display:inline-block;width:250px;padding:10px 0 10px 0}.form-field .disabled{color:#aaa;cursor:not-allowed}.form-field .disabled svg{fill:#aaa}.form-search{float:right;padding:7px}.form-search .form-text{width:auto}.form-search a svg{margin-top:-3px;fill:#e4e4e4;width:20px;vertical-align:middle;display:inline-block}.form-search a:hover svg{fill:#c2c2c2}.form-field-inline .form-text{display:inline-block;width:auto}.form-field-inline .form-error{display:inline-block}form h2{margin-bottom:15px}.panel+.panel{margin-top:15px}.panel{background:#fff;border-radius:3px;box-shadow:0 2px 4px 0 rgba(0,0,0,.1)}.panel .panel-body{padding:15px}.panel .panel-message{font-size:20px;padding:150px;text-align:center}.panel .panel-footer{border-top:solid 1px #e4e4e4;padding:15px}.panel table.code{border-radius:0 0 3px 3px}.panel-header{border-bottom:solid 1px #e4e4e4;overflow:auto}.panel-header h3{float:left;padding:15px;font-weight:700}.panel-header .panel-nav{list-style:none;float:left}.panel-header .panel-nav li{display:inline}.panel-header .panel-nav li a{display:inline-block;padding:15px;padding-left:17px;padding-right:17px;color:rgba(0,0,0,.4)}.panel-header .panel-nav li a svg{margin-right:3px;width:15px;vertical-align:middle;display:inline-block;fill:rgba(0,0,0,.4)}.panel-header .panel-nav li a span{margin-top:2px;vertical-align:middle;display:inline-block}.panel-header .panel-nav li a.active,.panel-header .panel-nav li a:hover{text-decoration:none;border-bottom:solid 2px #383e51;color:#383e51}.panel-header .panel-nav li a.active svg,.panel-header .panel-nav li a:hover svg{fill:#383e51}.panel-header .panel-actions{float:right;list-style:none;padding:7px}.panel-header .panel-actions .btn{padding:5px;padding-left:12px;padding-right:12px}.panel-header .panel-actions li{display:inline}.panel-header .panel-actions li a{display:inline-block}.panel-header .panel-actions li a svg{margin-right:3px;width:15px;vertical-align:middle;display:inline-block}.panel-header .panel-actions li a span{margin-top:2px;vertical-align:middle;display:inline-block}@media (max-width:1100px){.panel .panel-header .panel-nav li a span{display:none}}@media (max-width:700px){.panel .panel-header .form-search{display:none}}.pill{display:inline-block;text-align:center;padding:3px;padding-left:10px;padding-right:10px;border-radius:25px;color:#fff;font-size:14px;vertical-align:middle}.pill a{text-decoration:none}.pill svg{margin-top:-2px;display:inline-block;vertical-align:middle;width:15px;fill:#fff}.pill-bubble{margin-right:5px;border-radius:100%;width:25px;height:25px;text-align:center;display:inline-block}.pill-bubble svg{width:15px;fill:#fff;vertical-align:middle}a.pill:hover{text-decoration:none}.pill-light{background:#61a0ea}a.pill-light:hover{background:#5090d9}.pill-gray{background:#6a7393}.pill-dark{background:#272b39}.pill-red{background:#c64242}.pill-green{background:#269326}.pill-blue{background:#61a0ea}.pill-orange{background:#ff7400}@media (max-width:950px){.pill{width:25px!important}.pill span{display:none}}.providers{margin-top:15px;margin-bottom:15px}.provider-btn{display:inline-block;border-radius:3px;color:#fff;border:none;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";padding:10px;padding-left:15px;padding-right:15px}.provider-btn svg{margin-right:5px;fill:#fff;vertical-align:middle}.provider-btn span{display:inline-block;vertical-align:middle}.provider-btn:hover{text-decoration:none}.provider-github{background:#24292e}.provider-github:hover{background:#353a3f}.provider-gitlab{background:#fa7035}.provider-gitlab:hover{background:#e65328}.table{border-collapse:collapse;width:100%}.table td,.table th{padding:10px}.table td svg,.table th svg{width:15px;display:inline-block;vertical-align:middle}.table td form,.table th form{display:inline-block}.table td.success{color:#269326}.table td.success svg{fill:#269326}.table td.warning{color:#ff7400}.table td.warning svg{fill:#ff7400}.table td.error{color:#c64242}.table td.error svg{fill:#c64242}.table th{background:rgba(0,0,0,.03);border-bottom:solid 1px #e4e4e4;text-align:left;color:rgba(0,0,0,.5);font-weight:400}.table th.align-right{text-align:right}.table tr{border-bottom:solid 1px #e4e4e4}.table tr:last-child{border-bottom:none}.table .cell-pill{width:100px}.table .cell-date{text-align:right!important;width:250px}@media (max-width:900px){th.hide-mobile{display:none}td.hide-mobile{display:none}}.overflow{overflow:auto;padding-bottom:5px}.muted{color:#9f9f9f}.muted svg{fill:#9f9f9f}.hook-status{width:10px;text-align:center}.hook-status-err svg{fill:#c64242}.hook-status-none svg{fill:#6a7393}.hook-status-ok svg{fill:#269326}.align-center{text-align:center}.align-right{text-align:right}.inline-block{display:inline-block}.separator{margin-top:20px;margin-bottom:20px;border-bottom:solid 1px #cfcfcf}.slim{margin:0 auto;max-width:600px}.left{float:left}.right{float:right}.w-90{width:90px}.mt-5{margin-top:5px}.middle{vertical-align:middle}.mb-10{margin-bottom:10px}.pr-5{padding-right:5px}.pl-5{padding-left:5px}.progress-wrap .progress-bg{padding:3px;border-radius:3px;width:100%;background:#e4e4e4}.progress-wrap .progress{margin-top:-6px;padding:3px;border-radius:3px;background:#61a0ea}.svg-red svg{fill:#c64242}.svg-green svg{fill:#269326}.paginator{margin:0 auto;list-style:none;max-width:250px}.paginator li{display:inline}.paginator li a{display:inline-block;box-sizing:border-box;text-align:center;padding:10px;width:50%}.paginator li a.disabled{cursor:not-allowed;color:rgba(0,0,0,.5)}.paginator li a:hover{text-decoration:none}.paginator li .prev:hover{border-radius:3px 0 0 3px;background:#61a0ea;color:#fff}.paginator li .next:hover{border-radius:0 3px 3px 0;background:#61a0ea;color:#fff}.scope-list h3{margin-bottom:15px}.scope-list .scope-item{margin-top:15px;overflow:auto;border-top:solid 1px #cfcfcf;padding:15px}.scope-list .scope-item svg{display:inline-block;margin-right:15px;float:left;fill:rgba(0,0,0,.4)}.scope-list .scope-item span{display:inline-block}.scope-list .scope-item span strong{display:block}`)
//line template/template.qtpl:106
	qw422016.N().S(`</style> </head> <body>`)
//line template/template.qtpl:108
	p.StreamBody(qw422016)
//line template/template.qtpl:108
	qw422016.N().S(`</body> <footer>`)
//line template/template.qtpl:109
	p.StreamFooter(qw422016)
//line template/template.qtpl:109
	qw422016.N().S(`</footer> </html> `)
//line template/template.qtpl:111
}

//line template/template.qtpl:111
func WriteRender(qq422016 qtio422016.Writer, p Page) {
//line template/template.qtpl:111
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:111
	StreamRender(qw422016, p)
//line template/template.qtpl:111
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:111
}

//line template/template.qtpl:111
func Render(p Page) string {
//line template/template.qtpl:111
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:111
	WriteRender(qb422016, p)
//line template/template.qtpl:111
	qs422016 := string(qb422016.B)
//line template/template.qtpl:111
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:111
	return qs422016
//line template/template.qtpl:111
}

//line template/template.qtpl:114
func StreamRenderPaginator(qw422016 *qt422016.Writer, url *url.URL, p database.Paginator) {
//line template/template.qtpl:115
	if len(p.Pages) > 1 {
//line template/template.qtpl:115
		qw422016.N().S(`<ul class="paginator panel">`)
//line template/template.qtpl:117
		if p.Page == p.Prev {
//line template/template.qtpl:117
			qw422016.N().S(`<li><a class="disabled">Previous</a></li>`)
//line template/template.qtpl:119
		} else {
//line template/template.qtpl:119
			qw422016.N().S(`<li><a href="`)
//line template/template.qtpl:120
			qw422016.E().S(pageURL(url, p.Prev))
//line template/template.qtpl:120
			qw422016.N().S(`" class="prev">Previous</a></li>`)
//line template/template.qtpl:121
		}
//line template/template.qtpl:122
		if p.Page == p.Next {
//line template/template.qtpl:122
			qw422016.N().S(`<li><a class="disabled">Next</a></li>`)
//line template/template.qtpl:124
		} else {
//line template/template.qtpl:124
			qw422016.N().S(`<li><a href="`)
//line template/template.qtpl:125
			qw422016.E().S(pageURL(url, p.Next))
//line template/template.qtpl:125
			qw422016.N().S(`" class="next">Next</a></li>`)
//line template/template.qtpl:126
		}
//line template/template.qtpl:126
		qw422016.N().S(`</ul>`)
//line template/template.qtpl:128
	}
//line template/template.qtpl:129
}

//line template/template.qtpl:129
func WriteRenderPaginator(qq422016 qtio422016.Writer, url *url.URL, p database.Paginator) {
//line template/template.qtpl:129
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:129
	StreamRenderPaginator(qw422016, url, p)
//line template/template.qtpl:129
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:129
}

//line template/template.qtpl:129
func RenderPaginator(url *url.URL, p database.Paginator) string {
//line template/template.qtpl:129
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:129
	WriteRenderPaginator(qb422016, url, p)
//line template/template.qtpl:129
	qs422016 := string(qb422016.B)
//line template/template.qtpl:129
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:129
	return qs422016
//line template/template.qtpl:129
}

//line template/template.qtpl:132
func (f Form) StreamError(qw422016 *qt422016.Writer, field string) {
//line template/template.qtpl:132
	qw422016.N().S(` <div class="form-error"> `)
//line template/template.qtpl:134
	if err := f.Errors.First(field); err != "" {
//line template/template.qtpl:134
		qw422016.E().S(err)
//line template/template.qtpl:134
	}
//line template/template.qtpl:134
	qw422016.N().S(` </div> `)
//line template/template.qtpl:136
}

//line template/template.qtpl:136
func (f Form) WriteError(qq422016 qtio422016.Writer, field string) {
//line template/template.qtpl:136
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:136
	f.StreamError(qw422016, field)
//line template/template.qtpl:136
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:136
}

//line template/template.qtpl:136
func (f Form) Error(field string) string {
//line template/template.qtpl:136
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:136
	f.WriteError(qb422016, field)
//line template/template.qtpl:136
	qs422016 := string(qb422016.B)
//line template/template.qtpl:136
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:136
	return qs422016
//line template/template.qtpl:136
}

//line template/template.qtpl:138
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line template/template.qtpl:138
	qw422016.N().S(` Djinn CI `)
//line template/template.qtpl:138
}

//line template/template.qtpl:138
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line template/template.qtpl:138
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:138
	p.StreamTitle(qw422016)
//line template/template.qtpl:138
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:138
}

//line template/template.qtpl:138
func (p *BasePage) Title() string {
//line template/template.qtpl:138
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:138
	p.WriteTitle(qb422016)
//line template/template.qtpl:138
	qs422016 := string(qb422016.B)
//line template/template.qtpl:138
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:138
	return qs422016
//line template/template.qtpl:138
}

//line template/template.qtpl:139
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line template/template.qtpl:139
}

//line template/template.qtpl:139
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line template/template.qtpl:139
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:139
	p.StreamBody(qw422016)
//line template/template.qtpl:139
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:139
}

//line template/template.qtpl:139
func (p *BasePage) Body() string {
//line template/template.qtpl:139
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:139
	p.WriteBody(qb422016)
//line template/template.qtpl:139
	qs422016 := string(qb422016.B)
//line template/template.qtpl:139
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:139
	return qs422016
//line template/template.qtpl:139
}

//line template/template.qtpl:140
func (p *BasePage) StreamFooter(qw422016 *qt422016.Writer) {
//line template/template.qtpl:140
	qw422016.N().S(` `)
//line template/template.qtpl:140
}

//line template/template.qtpl:140
func (p *BasePage) WriteFooter(qq422016 qtio422016.Writer) {
//line template/template.qtpl:140
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:140
	p.StreamFooter(qw422016)
//line template/template.qtpl:140
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:140
}

//line template/template.qtpl:140
func (p *BasePage) Footer() string {
//line template/template.qtpl:140
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:140
	p.WriteFooter(qb422016)
//line template/template.qtpl:140
	qs422016 := string(qb422016.B)
//line template/template.qtpl:140
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:140
	return qs422016
//line template/template.qtpl:140
}

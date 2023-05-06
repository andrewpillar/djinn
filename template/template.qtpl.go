// Code generated by qtc from "template.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/template.qtpl:2
package template

//line template/template.qtpl:2
import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"djinn-ci.com/auth"
	"djinn-ci.com/errors"
	"djinn-ci.com/runner"

	"github.com/gorilla/csrf"
)

//line template/template.qtpl:20
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/template.qtpl:20
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/template.qtpl:20
type Template interface {
//line template/template.qtpl:20
	Title() string
//line template/template.qtpl:20
	StreamTitle(qw422016 *qt422016.Writer)
//line template/template.qtpl:20
	WriteTitle(qq422016 qtio422016.Writer)
//line template/template.qtpl:20
	Body() string
//line template/template.qtpl:20
	StreamBody(qw422016 *qt422016.Writer)
//line template/template.qtpl:20
	WriteBody(qq422016 qtio422016.Writer)
//line template/template.qtpl:20
	Footer() string
//line template/template.qtpl:20
	StreamFooter(qw422016 *qt422016.Writer)
//line template/template.qtpl:20
	WriteFooter(qq422016 qtio422016.Writer)
//line template/template.qtpl:20
}

//line template/template.qtpl:30
type Page struct {
	User *auth.User
	CSRF template.HTML
	URL  *url.URL
}

func NewPage(u *auth.User, r *http.Request) *Page {
	return &Page{
		User: u,
		CSRF: csrf.TemplateField(r),
		URL:  r.URL,
	}
}

func (p *Page) Href(vals url.Values) string {
	q := p.URL.Query()

	for k := range vals {
		v := vals.Get(k)

		if v == "" {
			delete(q, k)
			continue
		}
		q.Set(k, v)
	}

	url := (*p.URL)
	url.RawQuery = q.Encode()

	return url.String()
}

func HumanSize(n int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	i := 0

	for ; n >= 1024; i++ {
		n /= 1024
	}
	return fmt.Sprintf("%d %s", n, units[i])
}

func JSON(s string) string {
	var buf bytes.Buffer
	json.Indent(&buf, []byte(s), "", "    ")
	return buf.String()
}

type Error struct {
	Code    int
	Message string
	Error   error
}

type FatalError struct {
	Error

	Stack string
}

//line template/template.qtpl:93
func StreamCode(qw422016 *qt422016.Writer, code string) {
//line template/template.qtpl:93
	qw422016.N().S(` <div class="code-wrap"> <table class="code"> <tbody> `)
//line template/template.qtpl:97
	for i, line := range strings.Split(code, "\n") {
//line template/template.qtpl:97
		qw422016.N().S(` <tr> <td class="line-number"><a href="#L`)
//line template/template.qtpl:99
		qw422016.E().V(i + 1)
//line template/template.qtpl:99
		qw422016.N().S(`">`)
//line template/template.qtpl:99
		qw422016.E().V(i + 1)
//line template/template.qtpl:99
		qw422016.N().S(`</a></td> <td class="line" id="L`)
//line template/template.qtpl:100
		qw422016.E().V(i + 1)
//line template/template.qtpl:100
		qw422016.N().S(`">`)
//line template/template.qtpl:100
		qw422016.E().S(line)
//line template/template.qtpl:100
		qw422016.N().S(`</td> </tr> `)
//line template/template.qtpl:102
	}
//line template/template.qtpl:102
	qw422016.N().S(` </tbody> </table> </div> `)
//line template/template.qtpl:106
}

//line template/template.qtpl:106
func WriteCode(qq422016 qtio422016.Writer, code string) {
//line template/template.qtpl:106
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:106
	StreamCode(qw422016, code)
//line template/template.qtpl:106
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:106
}

//line template/template.qtpl:106
func Code(code string) string {
//line template/template.qtpl:106
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:106
	WriteCode(qb422016, code)
//line template/template.qtpl:106
	qs422016 := string(qb422016.B)
//line template/template.qtpl:106
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:106
	return qs422016
//line template/template.qtpl:106
}

//line template/template.qtpl:108
func StreamStatus(qw422016 *qt422016.Writer, s runner.Status) {
//line template/template.qtpl:108
	qw422016.N().S(` `)
//line template/template.qtpl:109
	switch s {
//line template/template.qtpl:110
	case runner.Queued:
//line template/template.qtpl:110
		qw422016.N().S(` <span class="pill w-90 pill-dark">`)
//line template/template.qtpl:111
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 11.484l3.984-3.984v-3.516h-7.969v3.516zM15.984 16.5l-3.984-3.984-3.984 3.984v3.516h7.969v-3.516zM6 2.016h12v6l-3.984 3.984 3.984 3.984v6h-12v-6l3.984-3.984-3.984-3.984v-6z"></path>
</svg>
`)
//line template/template.qtpl:111
		qw422016.N().S(` <span>Queued</span></span> `)
//line template/template.qtpl:112
	case runner.Running:
//line template/template.qtpl:112
		qw422016.N().S(` <span class="pill w-90 pill-blue">`)
//line template/template.qtpl:113
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M21.984 12c0 5.156-3.938 9.422-8.953 9.938v-2.016c3.938-0.516 6.984-3.891 6.984-7.922s-3.047-7.406-6.984-7.922v-2.016c5.016 0.516 8.953 4.781 8.953 9.938zM5.672 19.734l1.406-1.406c1.125 0.844 2.484 1.406 3.938 1.594v2.016c-2.016-0.188-3.844-0.984-5.344-2.203zM4.078 12.984c0.188 1.453 0.75 2.813 1.594 3.891l-1.406 1.453c-1.219-1.5-2.016-3.328-2.203-5.344h2.016zM5.672 7.078c-0.844 1.125-1.406 2.484-1.594 3.938h-2.016c0.188-2.016 0.984-3.844 2.203-5.344zM11.016 4.078c-1.453 0.188-2.813 0.75-3.938 1.594l-1.406-1.406c1.5-1.219 3.328-2.016 5.344-2.203v2.016zM13.031 9.797l2.953 2.203c-2.007 1.493-4.007 2.993-6 4.5z"></path>
</svg>
`)
//line template/template.qtpl:113
		qw422016.N().S(` <span>Running</span></span> `)
//line template/template.qtpl:114
	case runner.Passed:
//line template/template.qtpl:114
		qw422016.N().S(` <span class="pill w-90 pill-green">`)
//line template/template.qtpl:115
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
//line template/template.qtpl:115
		qw422016.N().S(` <span>Passed</span></span> `)
//line template/template.qtpl:116
	case runner.PassedWithFailures:
//line template/template.qtpl:116
		qw422016.N().S(` <span class="pill w-90 pill-orange">`)
//line template/template.qtpl:117
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 14.016v-4.031h-1.969v4.031h1.969zM12.984 18v-2.016h-1.969v2.016h1.969zM0.984 21l11.016-18.984 11.016 18.984h-22.031z"></path>
</svg>
`)
//line template/template.qtpl:117
		qw422016.N().S(` <span>Passed</span></span> `)
//line template/template.qtpl:118
	case runner.Failed:
//line template/template.qtpl:118
		qw422016.N().S(` <span class="pill w-90 pill-red">`)
//line template/template.qtpl:119
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
//line template/template.qtpl:119
		qw422016.N().S(` <span>Failed</span></span> `)
//line template/template.qtpl:120
	case runner.Killed:
//line template/template.qtpl:120
		qw422016.N().S(` <span class="pill w-90 pill-red">`)
//line template/template.qtpl:121
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 12.984v-6h-1.969v6h1.969zM12 17.297c0.703 0 1.313-0.609 1.313-1.313s-0.609-1.266-1.313-1.266-1.313 0.563-1.313 1.266 0.609 1.313 1.313 1.313zM15.75 3l5.25 5.25v7.5l-5.25 5.25h-7.5l-5.25-5.25v-7.5l5.25-5.25h7.5z"></path>
</svg>
`)
//line template/template.qtpl:121
		qw422016.N().S(` <span>Killed</span></span> `)
//line template/template.qtpl:122
	case runner.TimedOut:
//line template/template.qtpl:122
		qw422016.N().S(` <span class="pill w-90 pill-gray">`)
//line template/template.qtpl:123
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 20.016c3.891 0 6.984-3.141 6.984-7.031s-3.094-6.984-6.984-6.984-6.984 3.094-6.984 6.984 3.094 7.031 6.984 7.031zM19.031 7.406c1.219 1.547 1.969 3.469 1.969 5.578 0 4.969-4.031 9-9 9s-9-4.031-9-9 4.031-9 9-9c2.109 0 4.078 0.797 5.625 2.016l1.406-1.453c0.516 0.422 0.984 0.891 1.406 1.406zM11.016 14.016v-6h1.969v6h-1.969zM15 0.984v2.016h-6v-2.016h6z"></path>
</svg>
`)
//line template/template.qtpl:123
		qw422016.N().S(` <span>Timed Out</span></span> `)
//line template/template.qtpl:124
	}
//line template/template.qtpl:124
	qw422016.N().S(` `)
//line template/template.qtpl:125
}

//line template/template.qtpl:125
func WriteStatus(qq422016 qtio422016.Writer, s runner.Status) {
//line template/template.qtpl:125
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:125
	StreamStatus(qw422016, s)
//line template/template.qtpl:125
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:125
}

//line template/template.qtpl:125
func Status(s runner.Status) string {
//line template/template.qtpl:125
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:125
	WriteStatus(qb422016, s)
//line template/template.qtpl:125
	qs422016 := string(qb422016.B)
//line template/template.qtpl:125
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:125
	return qs422016
//line template/template.qtpl:125
}

//line template/template.qtpl:127
func StreamIconStatus(qw422016 *qt422016.Writer, s runner.Status) {
//line template/template.qtpl:127
	qw422016.N().S(` `)
//line template/template.qtpl:128
	switch s {
//line template/template.qtpl:129
	case runner.Queued:
//line template/template.qtpl:129
		qw422016.N().S(` <span class="pill-bubble pill-dark">`)
//line template/template.qtpl:130
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 11.484l3.984-3.984v-3.516h-7.969v3.516zM15.984 16.5l-3.984-3.984-3.984 3.984v3.516h7.969v-3.516zM6 2.016h12v6l-3.984 3.984 3.984 3.984v6h-12v-6l3.984-3.984-3.984-3.984v-6z"></path>
</svg>
`)
//line template/template.qtpl:130
		qw422016.N().S(`</span> `)
//line template/template.qtpl:131
	case runner.Running:
//line template/template.qtpl:131
		qw422016.N().S(` <span class="pill-bubble pill-blue">`)
//line template/template.qtpl:132
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M21.984 12c0 5.156-3.938 9.422-8.953 9.938v-2.016c3.938-0.516 6.984-3.891 6.984-7.922s-3.047-7.406-6.984-7.922v-2.016c5.016 0.516 8.953 4.781 8.953 9.938zM5.672 19.734l1.406-1.406c1.125 0.844 2.484 1.406 3.938 1.594v2.016c-2.016-0.188-3.844-0.984-5.344-2.203zM4.078 12.984c0.188 1.453 0.75 2.813 1.594 3.891l-1.406 1.453c-1.219-1.5-2.016-3.328-2.203-5.344h2.016zM5.672 7.078c-0.844 1.125-1.406 2.484-1.594 3.938h-2.016c0.188-2.016 0.984-3.844 2.203-5.344zM11.016 4.078c-1.453 0.188-2.813 0.75-3.938 1.594l-1.406-1.406c1.5-1.219 3.328-2.016 5.344-2.203v2.016zM13.031 9.797l2.953 2.203c-2.007 1.493-4.007 2.993-6 4.5z"></path>
</svg>
`)
//line template/template.qtpl:132
		qw422016.N().S(`</span> `)
//line template/template.qtpl:133
	case runner.Passed:
//line template/template.qtpl:133
		qw422016.N().S(` <span class="pill-bubble pill-green">`)
//line template/template.qtpl:134
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M9 16.172l10.594-10.594 1.406 1.406-12 12-5.578-5.578 1.406-1.406z"></path>
</svg>
`)
//line template/template.qtpl:134
		qw422016.N().S(`</span> `)
//line template/template.qtpl:135
	case runner.PassedWithFailures:
//line template/template.qtpl:135
		qw422016.N().S(` <span class="pill-bubble pill-orange">`)
//line template/template.qtpl:136
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 14.016v-4.031h-1.969v4.031h1.969zM12.984 18v-2.016h-1.969v2.016h1.969zM0.984 21l11.016-18.984 11.016 18.984h-22.031z"></path>
</svg>
`)
//line template/template.qtpl:136
		qw422016.N().S(`</span> `)
//line template/template.qtpl:137
	case runner.Failed:
//line template/template.qtpl:137
		qw422016.N().S(` <span class="pill-bubble pill-red">`)
//line template/template.qtpl:138
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M18.984 6.422l-5.578 5.578 5.578 5.578-1.406 1.406-5.578-5.578-5.578 5.578-1.406-1.406 5.578-5.578-5.578-5.578 1.406-1.406 5.578 5.578 5.578-5.578z"></path>
</svg>
`)
//line template/template.qtpl:138
		qw422016.N().S(`</span> `)
//line template/template.qtpl:139
	case runner.Killed:
//line template/template.qtpl:139
		qw422016.N().S(` <span class="pill-bubble pill-red">`)
//line template/template.qtpl:140
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 12.984v-6h-1.969v6h1.969zM12 17.297c0.703 0 1.313-0.609 1.313-1.313s-0.609-1.266-1.313-1.266-1.313 0.563-1.313 1.266 0.609 1.313 1.313 1.313zM15.75 3l5.25 5.25v7.5l-5.25 5.25h-7.5l-5.25-5.25v-7.5l5.25-5.25h7.5z"></path>
</svg>
`)
//line template/template.qtpl:140
		qw422016.N().S(`</span> `)
//line template/template.qtpl:141
	case runner.TimedOut:
//line template/template.qtpl:141
		qw422016.N().S(` <span class="pill-bubble pill-gray">`)
//line template/template.qtpl:142
		qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12 20.016c3.891 0 6.984-3.141 6.984-7.031s-3.094-6.984-6.984-6.984-6.984 3.094-6.984 6.984 3.094 7.031 6.984 7.031zM19.031 7.406c1.219 1.547 1.969 3.469 1.969 5.578 0 4.969-4.031 9-9 9s-9-4.031-9-9 4.031-9 9-9c2.109 0 4.078 0.797 5.625 2.016l1.406-1.453c0.516 0.422 0.984 0.891 1.406 1.406zM11.016 14.016v-6h1.969v6h-1.969zM15 0.984v2.016h-6v-2.016h6z"></path>
</svg>
`)
//line template/template.qtpl:142
		qw422016.N().S(`</span> `)
//line template/template.qtpl:143
	}
//line template/template.qtpl:143
	qw422016.N().S(` `)
//line template/template.qtpl:144
}

//line template/template.qtpl:144
func WriteIconStatus(qq422016 qtio422016.Writer, s runner.Status) {
//line template/template.qtpl:144
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:144
	StreamIconStatus(qw422016, s)
//line template/template.qtpl:144
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:144
}

//line template/template.qtpl:144
func IconStatus(s runner.Status) string {
//line template/template.qtpl:144
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:144
	WriteIconStatus(qb422016, s)
//line template/template.qtpl:144
	qs422016 := string(qb422016.B)
//line template/template.qtpl:144
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:144
	return qs422016
//line template/template.qtpl:144
}

//line template/template.qtpl:146
func StreamLogo(qw422016 *qt422016.Writer) {
//line template/template.qtpl:146
	qw422016.N().S(` <div class="logo"> <div class="handle"></div> <div class="lid"></div> <div class="lantern"></div> </div> `)
//line template/template.qtpl:152
}

//line template/template.qtpl:152
func WriteLogo(qq422016 qtio422016.Writer) {
//line template/template.qtpl:152
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:152
	StreamLogo(qw422016)
//line template/template.qtpl:152
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:152
}

//line template/template.qtpl:152
func Logo() string {
//line template/template.qtpl:152
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:152
	WriteLogo(qb422016)
//line template/template.qtpl:152
	qs422016 := string(qb422016.B)
//line template/template.qtpl:152
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:152
	return qs422016
//line template/template.qtpl:152
}

//line template/template.qtpl:154
func StreamRender(qw422016 *qt422016.Writer, tmpl Template) {
//line template/template.qtpl:154
	qw422016.N().S(` <!DOCTYPE HTML> <html lang="en"> <head> <meta charset="utf-8"> <meta content="width=device-width, initial-scale=1" name="viewport"> <title>`)
//line template/template.qtpl:160
	tmpl.StreamTitle(qw422016)
//line template/template.qtpl:160
	qw422016.N().S(` - Djinn CI</title> <style type="text/css">`)
//line template/template.qtpl:161
	qw422016.N().S(`*{margin:0;padding:0}body{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";font-size:14px;background:#eee;color:#444}a{color:#146de0;cursor:pointer;text-decoration:none}a:hover{text-decoration:underline}button{cursor:pointer}h1,h2,h3,h4,h5,h6{font-weight:400}.btn{border:none;border-radius:3px;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";padding:10px;padding-left:15px;padding-right:15px;color:#fff}.btn svg{fill:#fff}.btn:hover{text-decoration:none}.btn:disabled{cursor:not-allowed;background:#b2b2b2!important}.btn-primary{background:#61a0ea}.btn-primary:hover{background:#5090d9}.btn-danger{background:#de4141}.btn-danger:hover{background:#cd3030}span.code{padding:3px;border-radius:3px;background:#e6f0f5;font-family:monospace;white-space:pre-wrap}pre.code{background:#272b39;border-radius:0 0 3px 3px;box-sizing:border-box;color:#fff;font-family:monospace;font-size:12px;overflow:auto;padding:15px;width:100%}td.code{border-radius:0 3px 3px 0;text-align:right;width:50%}.code-wrap{overflow:scroll}table.code{background:#272b39;color:#fff;width:100%;font-family:monospace;font-size:12px;border-collapse:collapse;border-spacing:0}table.code .line-number{text-align:right;-moz-user-select:none;-ms-user-select:none;-webkit-user-select:none;min-width:30px;width:1%;padding-left:10px;padding-right:10px;line-height:20px}table.code .line-number a{display:block;color:rgba(255,255,255,.3)}table.code .line{padding-left:10px;padding-right:10px;line-height:20px;white-space:pre;word-wrap:normal}table.code .line:target{background:#383e51}.col-75{width:75%;box-sizing:border-box}.col-25{width:25%;box-sizing:border-box}.col-50{width:50%;box-sizing:border-box}.col-left{float:left;padding-right:5px}.col-right{float:right;padding-left:5px}@media (max-width:1100px){.col-75{margin-bottom:10px;width:100%}.col-25{margin-bottom:10px;width:100%}.col-50{margin-bottom:10px;width:100%}.col-left{padding-right:0;float:none}.col-right{padding-left:0;float:none}}.dashboard .sidebar{position:fixed;top:0;left:0;height:100%;width:225px;background:#383e51;overflow:auto}.dashboard .sidebar .sidebar-header{color:#fff;padding:20px;background:#272b39}.dashboard .sidebar .sidebar-header .logo{margin-top:-5px;margin-right:30px;display:inline-block;vertical-align:middle;width:0}.dashboard .sidebar .sidebar-header .logo .handle{margin-left:-3px;border-style:solid;border-width:2px 0 8px 7px;border-color:transparent transparent transparent #fff}.dashboard .sidebar .sidebar-header .logo .lid{margin-bottom:-20px;margin-left:13px;border-style:solid;border-width:5px 0 7px 5px;border-color:transparent transparent transparent #fff}.dashboard .sidebar .sidebar-header .logo .lantern{margin-left:-5px;border-style:solid;border-width:15px 15px 35px 0;border-color:transparent #fff transparent transparent}.dashboard .sidebar .sidebar-header h2{display:inline-block}.dashboard .sidebar .sidebar-auth a{display:block;color:rgba(255,255,255,.5);padding:15px;text-align:center}.dashboard .sidebar .sidebar-auth a.active,.dashboard .sidebar .sidebar-auth a:hover,.dashboard .sidebar .sidebar-auth button:hover{text-decoration:none;background:#272b39;color:#fff}.dashboard .sidebar .sidebar-nav{list-style:none}.dashboard .sidebar .sidebar-nav li{display:block}.dashboard .sidebar .sidebar-nav li a,.dashboard .sidebar .sidebar-nav li button{display:block;color:rgba(255,255,255,.5);padding:15px}.dashboard .sidebar .sidebar-nav li a svg,.dashboard .sidebar .sidebar-nav li button svg{margin-right:3px;display:inline-block;vertical-align:middle;fill:rgba(255,255,255,.5);width:15px}.dashboard .sidebar .sidebar-nav li a span,.dashboard .sidebar .sidebar-nav li button span{margin-top:2px;display:inline-block;vertical-align:middle}.dashboard .sidebar .sidebar-nav li button{width:100%;border:none;text-align:left;background:rgba(0,0,0,0)}.dashboard .sidebar .sidebar-nav li a.active,.dashboard .sidebar .sidebar-nav li a:hover,.dashboard .sidebar .sidebar-nav li button:hover{text-decoration:none;background:#272b39;color:#fff}.dashboard .sidebar .sidebar-nav li a.active svg,.dashboard .sidebar .sidebar-nav li a:hover svg,.dashboard .sidebar .sidebar-nav li button:hover svg{fill:#fff}.dashboard .sidebar .sidebar-nav li.sidebar-nav-header{padding:15px;font-weight:700;color:#fff}.dashboard-header{margin-bottom:10px}.dashboard-header h1{float:left}.dashboard-header h1 .back{margin-top:2px;display:inline-block;vertical-align:middle}.dashboard-header h1 .back svg{fill:#7f7f7f}.dashboard-header h1 .back:hover{text-decoration:none}.dashboard-header h1 .back:hover svg{fill:#444}.dashboard-header h1 small{margin-top:10px;display:block;font-size:16px;color:rgba(0,0,0,.5)}.dashboard-header .pill{margin-top:-5px;margin-left:10px}.dashboard-header .dashboard-actions{float:right;list-style:none}.dashboard-header .dashboard-actions li{display:inline}.dashboard-header .dashboard-actions li form{display:inline-block}.dashboard-header .dashboard-actions li a{cursor:pointer;display:inline-block}.dashboard-nav{list-style:none}.dashboard-nav li{display:inline}.dashboard-nav li a{display:inline-block;padding:15px;color:#9f9f9f}.dashboard-nav li a svg{margin-right:3px;width:20px;vertical-align:middle;display:inline-block;fill:#9f9f9f}.dashboard-nav li a span{margin-top:2px;display:inline-block;vertical-align:middle}.dashboard-nav li a.active,.dashboard-nav li a:hover{text-decoration:none;color:#272b39}.dashboard-nav li a.active svg,.dashboard-nav li a:hover svg{fill:#272b39}.dashboard-content{margin-left:225px}.dashboard-content .alert{overflow:auto;padding:15px}.dashboard-content .alert .alert-message{float:left;color:rgba(0,0,0,.6)}.dashboard-content .alert a.alert-close{float:right;display:inline-block}.dashboard-content .alert a.alert-close svg{width:15px;height:15px;fill:rgba(0,0,0,.4)}.dashboard-content .alert a.alert-close:hover svg{fill:rgba(0,0,0,.5)}.dashboard-content .alert-success{background:#caf5ca;border:solid 1px #a0dfa0}.dashboard-content .alert-warn{background:#fff3cd;border:solid 1px #d9c995}.dashboard-content .alert-danger{background:#ffd4d4;border:solid 1px #e19e9e}.dashboard-content .dashboard-wrap{margin:0 auto;max-width:1300px;padding:20px}@media (max-width:1500px){.dashboard .sidebar{width:70px}.dashboard .sidebar .sidebar-header{padding:15px}.dashboard .sidebar .sidebar-header .logo{margin-top:0;margin-right:0;margin-left:12px}.dashboard .sidebar .sidebar-header h2{display:none}.dashboard .sidebar .sidebar-nav .sidebar-nav-header{display:none}.dashboard .sidebar .sidebar-nav li a,.dashboard .sidebar .sidebar-nav li button{text-align:center}.dashboard .sidebar .sidebar-nav li a span,.dashboard .sidebar .sidebar-nav li button span{display:none}.dashboard .dashboard-content{margin-left:70px}}@media (max-width:1000px){.dashboard .dashboard-content .dashboard-header .dashboard-nav li a span{display:none}}.form-field+.form-field{margin-top:15px}.form-field{overflow:auto}.form-field .label{margin-bottom:5px;display:block;font-weight:700}.form-field .label small{color:rgba(0,0,0,.5)}.form-field .form-error{margin-top:5px;color:#ff4343;min-height:20px}.form-field .form-text{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";font-size:14px;padding:10px;outline:0;border-radius:3px;box-sizing:border-box;width:100%;border:solid 1px #e4e4e4}.form-field .form-text:focus{border:solid 1px #c2c2c2}.form-field .form-code{min-height:250px;font-family:monospace}.form-field textarea.form-text{min-width:100%;max-width:100%}.form-field .form-option+.form-option{margin-top:10px}.form-field .form-option{display:block;cursor:pointer;overflow:auto}.form-field .form-option .form-selector{margin-right:5px;outline:0}.form-field .form-option .form-option-info{margin-right:5px;display:inline-block}.form-field .form-option svg{margin-right:5px;fill:rgba(0,0,0,.4)}.form-field .hook-event{cursor:pointer;display:inline-block;width:250px;padding:10px 0 10px 0}.form-field .disabled{color:#aaa;cursor:not-allowed}.form-field .disabled svg{fill:#aaa}.form-search{float:right;padding:7px}.form-search .form-text{width:auto}.form-search a svg{margin-top:-3px;fill:#e4e4e4;width:20px;vertical-align:middle;display:inline-block}.form-search a:hover svg{fill:#c2c2c2}.form-field-inline .form-text{display:inline-block;width:auto}.form-field-inline .form-error{display:inline-block}form h2{margin-bottom:15px}.panel+.panel{margin-top:15px}.panel{background:#fff;border-radius:3px;box-shadow:0 2px 4px 0 rgba(0,0,0,.1)}.panel .panel-body{padding:15px}.panel .panel-message{font-size:20px;padding:150px;text-align:center}.panel .panel-footer{border-top:solid 1px #e4e4e4;padding:15px}.panel table.code{border-radius:0 0 3px 3px}.panel-header{border-bottom:solid 1px #e4e4e4;overflow:auto}.panel-header h3{float:left;padding:15px;font-weight:700}.panel-header .panel-nav{list-style:none;float:left}.panel-header .panel-nav li{display:inline}.panel-header .panel-nav li a{display:inline-block;padding:15px;padding-left:17px;padding-right:17px;color:rgba(0,0,0,.4)}.panel-header .panel-nav li a svg{margin-right:3px;width:15px;vertical-align:middle;display:inline-block;fill:rgba(0,0,0,.4)}.panel-header .panel-nav li a span{margin-top:2px;vertical-align:middle;display:inline-block}.panel-header .panel-nav li a.active,.panel-header .panel-nav li a:hover{text-decoration:none;border-bottom:solid 2px #383e51;color:#383e51}.panel-header .panel-nav li a.active svg,.panel-header .panel-nav li a:hover svg{fill:#383e51}.panel-header .panel-actions{float:right;list-style:none;padding:7px}.panel-header .panel-actions .btn{padding:5px;padding-left:12px;padding-right:12px}.panel-header .panel-actions li{display:inline}.panel-header .panel-actions li a{display:inline-block}.panel-header .panel-actions li a svg{margin-right:3px;width:15px;vertical-align:middle;display:inline-block}.panel-header .panel-actions li a span{margin-top:2px;vertical-align:middle;display:inline-block}@media (max-width:1100px){.panel .panel-header .panel-nav li a span{display:none}}@media (max-width:700px){.panel .panel-header .form-search{display:none}}.pill{display:inline-block;text-align:center;padding:3px;padding-left:10px;padding-right:10px;border-radius:25px;color:#fff;font-size:14px;vertical-align:middle}.pill a{text-decoration:none}.pill svg{margin-top:-2px;display:inline-block;vertical-align:middle;width:15px;fill:#fff}.pill-bubble{margin-right:5px;border-radius:100%;width:25px;height:25px;text-align:center;display:inline-block}.pill-bubble svg{width:15px;fill:#fff;vertical-align:middle}a.pill:hover{text-decoration:none}.pill-light{background:#61a0ea}a.pill-light:hover{background:#5090d9}.pill-gray{background:#6a7393}.pill-dark{background:#272b39}.pill-red{background:#c64242}.pill-green{background:#269326}.pill-blue{background:#61a0ea}.pill-orange{background:#ff7400}@media (max-width:950px){.pill{width:25px!important}.pill span{display:none}}.providers{margin-top:15px;margin-bottom:15px}.provider-btn{display:inline-block;border-radius:3px;color:#fff;border:none;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Sego UI Symbol";padding:10px;padding-left:15px;padding-right:15px}.provider-btn svg{margin-right:5px;fill:#fff;vertical-align:middle}.provider-btn span{display:inline-block;vertical-align:middle}.provider-btn:hover{text-decoration:none}.provider-github{background:#24292e}.provider-github:hover{background:#353a3f}.provider-gitlab{background:#fa7035}.provider-gitlab:hover{background:#e65328}.table{border-collapse:collapse;width:100%}.table td,.table th{padding:10px}.table td svg,.table th svg{width:15px;display:inline-block;vertical-align:middle}.table td form,.table th form{display:inline-block}.table td.success{color:#269326}.table td.success svg{fill:#269326}.table td.warning{color:#ff7400}.table td.warning svg{fill:#ff7400}.table td.error{color:#c64242}.table td.error svg{fill:#c64242}.table th{background:rgba(0,0,0,.03);border-bottom:solid 1px #e4e4e4;text-align:left;color:rgba(0,0,0,.5);font-weight:400}.table th.align-right{text-align:right}.table tr{border-bottom:solid 1px #e4e4e4}.table tr:last-child{border-bottom:none}.table .cell-pill{width:100px}.table .cell-date{text-align:right!important;width:250px}@media (max-width:900px){th.hide-mobile{display:none}td.hide-mobile{display:none}}.overflow{overflow:auto;padding-bottom:5px}.muted{color:#9f9f9f}.muted svg{fill:#9f9f9f}a.active,a.muted:hover{color:#272b39}a.active svg,a.muted:hover svg{fill:#272b39}.hook-status{width:10px;text-align:center}.hook-status-err svg{fill:#c64242}.hook-status-none svg{fill:#6a7393}.hook-status-ok svg{fill:#269326}.align-center{text-align:center}.align-right{text-align:right}.inline-block{display:inline-block}.separator{margin-top:20px;margin-bottom:20px;border-bottom:solid 1px #cfcfcf}.slim{margin:0 auto;max-width:600px}.left{float:left}.right{float:right}.w-90{width:90px}.mt-5{margin-top:5px}.middle{vertical-align:middle}.mb-10{margin-bottom:10px}.pr-5{padding-right:5px}.pl-5{padding-left:5px}.progress-wrap .progress-bg{padding:3px;border-radius:3px;width:100%;background:#e4e4e4}.progress-wrap .progress{margin-top:-6px;padding:3px;border-radius:3px;background:#61a0ea}.svg-red svg{fill:#c64242}.svg-green svg{fill:#269326}.paginator{margin:0 auto;list-style:none;max-width:250px}.paginator li{display:inline}.paginator li a{display:inline-block;box-sizing:border-box;text-align:center;padding:10px;width:50%}.paginator li a.disabled{cursor:not-allowed;color:rgba(0,0,0,.5)}.paginator li a:hover{text-decoration:none}.paginator li .prev:hover{border-radius:3px 0 0 3px;background:#61a0ea;color:#fff}.paginator li .next:hover{border-radius:0 3px 3px 0;background:#61a0ea;color:#fff}.scope-list h3{margin-bottom:15px}.scope-list .scope-item{margin-top:15px;overflow:auto;border-top:solid 1px #cfcfcf;padding:15px}.scope-list .scope-item svg{display:inline-block;margin-right:15px;float:left;fill:rgba(0,0,0,.4)}.scope-list .scope-item span{display:inline-block}.scope-list .scope-item span strong{display:block}`)
//line template/template.qtpl:161
	qw422016.N().S(`</style> </head> <body>`)
//line template/template.qtpl:163
	tmpl.StreamBody(qw422016)
//line template/template.qtpl:163
	qw422016.N().S(`</body> <footer>`)
//line template/template.qtpl:164
	tmpl.StreamFooter(qw422016)
//line template/template.qtpl:164
	qw422016.N().S(`</footer> </html> `)
//line template/template.qtpl:166
}

//line template/template.qtpl:166
func WriteRender(qq422016 qtio422016.Writer, tmpl Template) {
//line template/template.qtpl:166
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:166
	StreamRender(qw422016, tmpl)
//line template/template.qtpl:166
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:166
}

//line template/template.qtpl:166
func Render(tmpl Template) string {
//line template/template.qtpl:166
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:166
	WriteRender(qb422016, tmpl)
//line template/template.qtpl:166
	qs422016 := string(qb422016.B)
//line template/template.qtpl:166
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:166
	return qs422016
//line template/template.qtpl:166
}

//line template/template.qtpl:168
func (p Error) StreamTitle(qw422016 *qt422016.Writer) {
//line template/template.qtpl:168
	qw422016.N().S(`Error`)
//line template/template.qtpl:168
}

//line template/template.qtpl:168
func (p Error) WriteTitle(qq422016 qtio422016.Writer) {
//line template/template.qtpl:168
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:168
	p.StreamTitle(qw422016)
//line template/template.qtpl:168
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:168
}

//line template/template.qtpl:168
func (p Error) Title() string {
//line template/template.qtpl:168
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:168
	p.WriteTitle(qb422016)
//line template/template.qtpl:168
	qs422016 := string(qb422016.B)
//line template/template.qtpl:168
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:168
	return qs422016
//line template/template.qtpl:168
}

//line template/template.qtpl:170
func (p Error) StreamBody(qw422016 *qt422016.Writer) {
//line template/template.qtpl:170
	qw422016.N().S(` <div class="error"> `)
//line template/template.qtpl:172
	StreamLogo(qw422016)
//line template/template.qtpl:172
	qw422016.N().S(` <h1>`)
//line template/template.qtpl:173
	qw422016.E().V(p.Code)
//line template/template.qtpl:173
	qw422016.N().S(`</h1> <h2>`)
//line template/template.qtpl:174
	qw422016.E().S(p.Message)
//line template/template.qtpl:174
	qw422016.N().S(`</h2> <br/> <a href="/">Back</a> <br/><br/> `)
//line template/template.qtpl:178
	if p.Error != nil {
//line template/template.qtpl:178
		qw422016.N().S(` <textarea readonly>`)
//line template/template.qtpl:179
		qw422016.E().S(errors.Format(p.Error))
//line template/template.qtpl:179
		qw422016.N().S(`</textarea> `)
//line template/template.qtpl:180
	}
//line template/template.qtpl:180
	qw422016.N().S(` </div> `)
//line template/template.qtpl:182
}

//line template/template.qtpl:182
func (p Error) WriteBody(qq422016 qtio422016.Writer) {
//line template/template.qtpl:182
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:182
	p.StreamBody(qw422016)
//line template/template.qtpl:182
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:182
}

//line template/template.qtpl:182
func (p Error) Body() string {
//line template/template.qtpl:182
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:182
	p.WriteBody(qb422016)
//line template/template.qtpl:182
	qs422016 := string(qb422016.B)
//line template/template.qtpl:182
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:182
	return qs422016
//line template/template.qtpl:182
}

//line template/template.qtpl:184
func (p Error) StreamFooter(qw422016 *qt422016.Writer) {
//line template/template.qtpl:184
	qw422016.N().S(` <style type="text/css">`)
//line template/template.qtpl:185
	qw422016.N().S(`*{margin:0;padding:0}a{color:#66c9ff;cursor:pointer;text-decoration:none}body{font-family:sans-serif;font-size:14px;background:#383e51;color:#fff}h1,h2{font-weight:400}.error{margin:0 auto;margin-top:250px;padding:20px;text-align:center}.error .logo{margin:0 auto;margin-bottom:20px;width:0}.error .logo .handle{margin-left:-20px;border-style:solid;border-width:5px 0 12px 10px;border-color:transparent transparent transparent #fff}.error .logo .lid{margin-bottom:-30px;margin-left:5px;border-style:solid;border-width:5px 0 12px 10px;border-color:transparent transparent transparent #fff}.error .logo .lantern{margin-left:-25px;border-style:solid;border-width:25px 25px 75px 0;border-color:transparent #fff transparent transparent}.error h2{margin-top:20px}textarea{font-family:monospace;box-sizing:border-box;min-width:100%;max-width:100%;min-width:700px;min-height:300px;border:solid 1px rgba(255,255,255,.3);border-radius:3px;background:rgba(0,0,0,.3);color:#fff;white-space:pre}textarea:focus{border:solid 1px rgba(255,255,255,.5)}`)
//line template/template.qtpl:185
	qw422016.N().S(`</style> `)
//line template/template.qtpl:186
}

//line template/template.qtpl:186
func (p Error) WriteFooter(qq422016 qtio422016.Writer) {
//line template/template.qtpl:186
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:186
	p.StreamFooter(qw422016)
//line template/template.qtpl:186
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:186
}

//line template/template.qtpl:186
func (p Error) Footer() string {
//line template/template.qtpl:186
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:186
	p.WriteFooter(qb422016)
//line template/template.qtpl:186
	qs422016 := string(qb422016.B)
//line template/template.qtpl:186
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:186
	return qs422016
//line template/template.qtpl:186
}

//line template/template.qtpl:188
func (p FatalError) StreamBody(qw422016 *qt422016.Writer) {
//line template/template.qtpl:188
	qw422016.N().S(` <div class="error"> `)
//line template/template.qtpl:190
	StreamLogo(qw422016)
//line template/template.qtpl:190
	qw422016.N().S(` <h1>`)
//line template/template.qtpl:191
	qw422016.E().V(p.Code)
//line template/template.qtpl:191
	qw422016.N().S(`</h1> <h2>`)
//line template/template.qtpl:192
	qw422016.E().S(p.Message)
//line template/template.qtpl:192
	qw422016.N().S(`</h2> <br/> <a href="/">Back</a> <br/><br/> <textarea readonly>`)
//line template/template.qtpl:196
	qw422016.E().S(p.Stack)
//line template/template.qtpl:196
	qw422016.N().S(`</textarea> </div> `)
//line template/template.qtpl:198
}

//line template/template.qtpl:198
func (p FatalError) WriteBody(qq422016 qtio422016.Writer) {
//line template/template.qtpl:198
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/template.qtpl:198
	p.StreamBody(qw422016)
//line template/template.qtpl:198
	qt422016.ReleaseWriter(qw422016)
//line template/template.qtpl:198
}

//line template/template.qtpl:198
func (p FatalError) Body() string {
//line template/template.qtpl:198
	qb422016 := qt422016.AcquireByteBuffer()
//line template/template.qtpl:198
	p.WriteBody(qb422016)
//line template/template.qtpl:198
	qs422016 := string(qb422016.B)
//line template/template.qtpl:198
	qt422016.ReleaseByteBuffer(qb422016)
//line template/template.qtpl:198
	return qs422016
//line template/template.qtpl:198
}

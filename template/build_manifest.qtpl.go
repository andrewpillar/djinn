// Code generated by qtc from "build_manifest.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/build_manifest.qtpl:1
package template

//line template/build_manifest.qtpl:1
import "djinn-ci.com/build"

//line template/build_manifest.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/build_manifest.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/build_manifest.qtpl:4
type BuildManifest struct {
	Build *build.Build
}

//line template/build_manifest.qtpl:10
func (p *BuildManifest) StreamTitle(qw422016 *qt422016.Writer) {
//line template/build_manifest.qtpl:10
	qw422016.N().S(`Manifest`)
//line template/build_manifest.qtpl:10
}

//line template/build_manifest.qtpl:10
func (p *BuildManifest) WriteTitle(qq422016 qtio422016.Writer) {
//line template/build_manifest.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/build_manifest.qtpl:10
	p.StreamTitle(qw422016)
//line template/build_manifest.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line template/build_manifest.qtpl:10
}

//line template/build_manifest.qtpl:10
func (p *BuildManifest) Title() string {
//line template/build_manifest.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line template/build_manifest.qtpl:10
	p.WriteTitle(qb422016)
//line template/build_manifest.qtpl:10
	qs422016 := string(qb422016.B)
//line template/build_manifest.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line template/build_manifest.qtpl:10
	return qs422016
//line template/build_manifest.qtpl:10
}

//line template/build_manifest.qtpl:12
func (p *BuildManifest) StreamHeader(qw422016 *qt422016.Writer) {
//line template/build_manifest.qtpl:12
}

//line template/build_manifest.qtpl:12
func (p *BuildManifest) WriteHeader(qq422016 qtio422016.Writer) {
//line template/build_manifest.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/build_manifest.qtpl:12
	p.StreamHeader(qw422016)
//line template/build_manifest.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line template/build_manifest.qtpl:12
}

//line template/build_manifest.qtpl:12
func (p *BuildManifest) Header() string {
//line template/build_manifest.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
//line template/build_manifest.qtpl:12
	p.WriteHeader(qb422016)
//line template/build_manifest.qtpl:12
	qs422016 := string(qb422016.B)
//line template/build_manifest.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
//line template/build_manifest.qtpl:12
	return qs422016
//line template/build_manifest.qtpl:12
}

//line template/build_manifest.qtpl:13
func (p *BuildManifest) StreamActions(qw422016 *qt422016.Writer) {
//line template/build_manifest.qtpl:13
}

//line template/build_manifest.qtpl:13
func (p *BuildManifest) WriteActions(qq422016 qtio422016.Writer) {
//line template/build_manifest.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/build_manifest.qtpl:13
	p.StreamActions(qw422016)
//line template/build_manifest.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line template/build_manifest.qtpl:13
}

//line template/build_manifest.qtpl:13
func (p *BuildManifest) Actions() string {
//line template/build_manifest.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line template/build_manifest.qtpl:13
	p.WriteActions(qb422016)
//line template/build_manifest.qtpl:13
	qs422016 := string(qb422016.B)
//line template/build_manifest.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line template/build_manifest.qtpl:13
	return qs422016
//line template/build_manifest.qtpl:13
}

//line template/build_manifest.qtpl:14
func (p *BuildManifest) StreamNavigation(qw422016 *qt422016.Writer) {
//line template/build_manifest.qtpl:14
}

//line template/build_manifest.qtpl:14
func (p *BuildManifest) WriteNavigation(qq422016 qtio422016.Writer) {
//line template/build_manifest.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/build_manifest.qtpl:14
	p.StreamNavigation(qw422016)
//line template/build_manifest.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line template/build_manifest.qtpl:14
}

//line template/build_manifest.qtpl:14
func (p *BuildManifest) Navigation() string {
//line template/build_manifest.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line template/build_manifest.qtpl:14
	p.WriteNavigation(qb422016)
//line template/build_manifest.qtpl:14
	qs422016 := string(qb422016.B)
//line template/build_manifest.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line template/build_manifest.qtpl:14
	return qs422016
//line template/build_manifest.qtpl:14
}

//line template/build_manifest.qtpl:15
func (p *BuildManifest) StreamFooter(qw422016 *qt422016.Writer) {
//line template/build_manifest.qtpl:15
}

//line template/build_manifest.qtpl:15
func (p *BuildManifest) WriteFooter(qq422016 qtio422016.Writer) {
//line template/build_manifest.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/build_manifest.qtpl:15
	p.StreamFooter(qw422016)
//line template/build_manifest.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line template/build_manifest.qtpl:15
}

//line template/build_manifest.qtpl:15
func (p *BuildManifest) Footer() string {
//line template/build_manifest.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line template/build_manifest.qtpl:15
	p.WriteFooter(qb422016)
//line template/build_manifest.qtpl:15
	qs422016 := string(qb422016.B)
//line template/build_manifest.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line template/build_manifest.qtpl:15
	return qs422016
//line template/build_manifest.qtpl:15
}

//line template/build_manifest.qtpl:17
func (p *BuildManifest) StreamBody(qw422016 *qt422016.Writer) {
//line template/build_manifest.qtpl:17
	qw422016.N().S(` <div class="panel"> <div class="panel-header"> <ul class="panel-actions"> <li> <a class="btn btn-primary" href="`)
//line template/build_manifest.qtpl:22
	qw422016.E().S(p.Build.Endpoint("manifest", "raw"))
//line template/build_manifest.qtpl:22
	qw422016.N().S(`"> `)
//line template/build_manifest.qtpl:23
	qw422016.N().S(`<!-- Generated by IcoMoon.io -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M12.984 9h5.531l-5.531-5.484v5.484zM15.984 14.016v-2.016h-7.969v2.016h7.969zM15.984 18v-2.016h-7.969v2.016h7.969zM14.016 2.016l6 6v12c0 1.078-0.938 1.969-2.016 1.969h-12c-1.078 0-2.016-0.891-2.016-1.969l0.047-16.031c0-1.078 0.891-1.969 1.969-1.969h8.016z"></path>
</svg>
`)
//line template/build_manifest.qtpl:23
	qw422016.N().S(`<span>Raw</span> </a> </li> </ul> </div> `)
//line template/build_manifest.qtpl:28
	StreamCode(qw422016, p.Build.Manifest.String())
//line template/build_manifest.qtpl:28
	qw422016.N().S(` </div> `)
//line template/build_manifest.qtpl:30
}

//line template/build_manifest.qtpl:30
func (p *BuildManifest) WriteBody(qq422016 qtio422016.Writer) {
//line template/build_manifest.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/build_manifest.qtpl:30
	p.StreamBody(qw422016)
//line template/build_manifest.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line template/build_manifest.qtpl:30
}

//line template/build_manifest.qtpl:30
func (p *BuildManifest) Body() string {
//line template/build_manifest.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line template/build_manifest.qtpl:30
	p.WriteBody(qb422016)
//line template/build_manifest.qtpl:30
	qs422016 := string(qb422016.B)
//line template/build_manifest.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line template/build_manifest.qtpl:30
	return qs422016
//line template/build_manifest.qtpl:30
}
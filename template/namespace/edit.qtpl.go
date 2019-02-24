// This file is automatically generated by qtc from "edit.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line template/namespace/edit.qtpl:2
package namespace

//line template/namespace/edit.qtpl:2
import (
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
)

//line template/namespace/edit.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/namespace/edit.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/namespace/edit.qtpl:10
type EditPage struct {
	*template.Page

	Errors    form.Errors
	Namespace *model.Namespace
}

//line template/namespace/edit.qtpl:18
func (p *EditPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line template/namespace/edit.qtpl:18
	qw422016.N().S(`
`)
	//line template/namespace/edit.qtpl:19
	qw422016.E().S(p.Namespace.Name)
	//line template/namespace/edit.qtpl:19
	qw422016.N().S(` - Edit
`)
//line template/namespace/edit.qtpl:20
}

//line template/namespace/edit.qtpl:20
func (p *EditPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line template/namespace/edit.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/edit.qtpl:20
	p.StreamTitle(qw422016)
	//line template/namespace/edit.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/edit.qtpl:20
}

//line template/namespace/edit.qtpl:20
func (p *EditPage) Title() string {
	//line template/namespace/edit.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/edit.qtpl:20
	p.WriteTitle(qb422016)
	//line template/namespace/edit.qtpl:20
	qs422016 := string(qb422016.B)
	//line template/namespace/edit.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/edit.qtpl:20
	return qs422016
//line template/namespace/edit.qtpl:20
}

//line template/namespace/edit.qtpl:22
func (p *EditPage) StreamBody(qw422016 *qt422016.Writer) {
	//line template/namespace/edit.qtpl:22
	qw422016.N().S(`
<div class="dashboard-header">
	<h1>Edit Namespace</h1>
</div>
<div class="dashboard-body">
	<div class="panel panel-slim">
		<form method="POST" action="/u/`)
	//line template/namespace/edit.qtpl:28
	qw422016.E().S(p.Namespace.User.Username)
	//line template/namespace/edit.qtpl:28
	qw422016.N().S(`/`)
	//line template/namespace/edit.qtpl:28
	qw422016.E().S(p.Namespace.Name)
	//line template/namespace/edit.qtpl:28
	qw422016.N().S(`">
			<input type="hidden" name="_method" value="PATCH"/>
			`)
	//line template/namespace/edit.qtpl:30
	if p.Errors.First("namespace") != "" {
		//line template/namespace/edit.qtpl:30
		qw422016.N().S(`
				<div class="form-error">Failed to create namespace: `)
		//line template/namespace/edit.qtpl:31
		qw422016.E().S(p.Errors.First("namespace"))
		//line template/namespace/edit.qtpl:31
		qw422016.N().S(`</div>
			`)
		//line template/namespace/edit.qtpl:32
	}
	//line template/namespace/edit.qtpl:32
	qw422016.N().S(`
			<div class="input-field">
				<label class="input-field-label">Name</label>
				<input class="input-text" type="text" name="name" value="`)
	//line template/namespace/edit.qtpl:35
	qw422016.E().S(p.Namespace.Name)
	//line template/namespace/edit.qtpl:35
	qw422016.N().S(`" autocomplete="off"/>
				<span class="error">`)
	//line template/namespace/edit.qtpl:36
	qw422016.E().S(p.Errors.First("name"))
	//line template/namespace/edit.qtpl:36
	qw422016.N().S(`</span>
			</div>
			<div class="input-field">
				<label class="input-field-label">Description</label>
				<textarea class="input-text" name="description">`)
	//line template/namespace/edit.qtpl:40
	qw422016.E().S(p.Namespace.Description)
	//line template/namespace/edit.qtpl:40
	qw422016.N().S(`</textarea>
			</div>
			<div class="input-field">
				<label class="input-field-label">Visibility</label>
				<label class="input-option">
					<input class="input-option-selector" type="radio" name="visibility" value="private" `)
	//line template/namespace/edit.qtpl:45
	if p.Namespace.Visibility == model.Private {
		//line template/namespace/edit.qtpl:45
		qw422016.N().S(`checked="true"`)
		//line template/namespace/edit.qtpl:45
	}
	//line template/namespace/edit.qtpl:45
	qw422016.N().S(`/>
					<div class="input-option-description">
						Private<br/>
						Only you will be able to view builds in the namespace.
					</div>
				</label>
				<label class="input-option">
					<input class="input-option-selector" type="radio" name="visibility" value="internal" `)
	//line template/namespace/edit.qtpl:52
	if p.Namespace.Visibility == model.Internal {
		//line template/namespace/edit.qtpl:52
		qw422016.N().S(`checked="true"`)
		//line template/namespace/edit.qtpl:52
	}
	//line template/namespace/edit.qtpl:52
	qw422016.N().S(`/>
					<div class="input-option-description">
						Internal<br/>
						Anyone with an account will be able to view builds in the namespace
					</div>
				</label>
				<label class="input-option">
					<input class="input-option-selector" type="radio" name="visibility" value="public" `)
	//line template/namespace/edit.qtpl:59
	if p.Namespace.Visibility == model.Public {
		//line template/namespace/edit.qtpl:59
		qw422016.N().S(`checked="true"`)
		//line template/namespace/edit.qtpl:59
	}
	//line template/namespace/edit.qtpl:59
	qw422016.N().S(`/>
					<div class="input-option-description">
						Public<br/>
						Anyone will be able to view builds in the namespace.
					</div>
				</label>
			</div>
			<div class="input-field">
				<button type="submit" class="button button-primary">Save</button>
			</div>
		</form>
	</div>
</div>
`)
//line template/namespace/edit.qtpl:72
}

//line template/namespace/edit.qtpl:72
func (p *EditPage) WriteBody(qq422016 qtio422016.Writer) {
	//line template/namespace/edit.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line template/namespace/edit.qtpl:72
	p.StreamBody(qw422016)
	//line template/namespace/edit.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line template/namespace/edit.qtpl:72
}

//line template/namespace/edit.qtpl:72
func (p *EditPage) Body() string {
	//line template/namespace/edit.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
	//line template/namespace/edit.qtpl:72
	p.WriteBody(qb422016)
	//line template/namespace/edit.qtpl:72
	qs422016 := string(qb422016.B)
	//line template/namespace/edit.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
	//line template/namespace/edit.qtpl:72
	return qs422016
//line template/namespace/edit.qtpl:72
}

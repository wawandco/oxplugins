package resource

// templateIndexTmpl is a variable for the [template_name]/index.plush.html
var templateIndexTmpl string = `<div class="py-4 mb-2">
	<h3 class="d-inline-block">{{ .Name.Group }}</h3>
	<div class="float-right">
		<%= linkTo(new{{ .Name.Resource }}Path(), {class: "btn btn-primary"}) { %>
		Create New {{ .Name.Proper }}
		<% } %>
	</div>
</div>

<table class="table table-hover table-bordered">
	<thead class="thead-light">
		{{ range $p := .Model.Attrs -}}
		{{ if ne $p.CommonType "text" -}}
		<th>{{ $p.Name.Pascalize }}</th>
		{{ end }}
		{{- end -}}
		<th>&nbsp;</th>
	</thead>
	<tbody>
		<%= for ({{ .Name.VarCaseSingle }}) in {{ .Name.VarCasePlural }} { %>
		<tr>
			{{ range $mp := .Model.Attrs -}}
			{{- if ne $mp.CommonType "text" -}}
			<td class="align-middle"><%= {{ $.Name.VarCaseSingle }}.{{ $mp.Name.Pascalize }} %></td>
			{{ end }}
			{{- end -}}
			<td>
				<div class="float-right">
					<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-info", body: "View"}) %>
					<%= linkTo(edit{{ .Name.Proper }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-warning", body: "Edit"}) %>
					<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>
				</div>
			</td>
		</tr>
		<% } %>
	</tbody>
</table>

<div class="text-center">
	<%= paginator(pagination) %>
</div>`

// templateNewTmpl is a variable for the [template_name]/new.plush.html
var templateNewTmpl string = `<div class="py-4 mb-2">
  	<h3 class="d-inline-block">New {{ .Name.Proper }}</h3>
</div>

<%= formFor({{.Name.VarCaseSingle}}, {action: {{ .Name.VarCasePlural }}Path(), method: "POST"}) { %>
	<%= partial("{{ .Name.Folder.Pluralize }}/form.html") %>
	<%= linkTo({{ .Name.VarCasePlural }}Path(), {class: "btn btn-warning", "data-confirm": "Are you sure?", body: "Cancel"}) %>
<% } %>`

// templateEditTmpl is a variable for the [template_name]/edit.plush.html
var templateEditTmpl string = `<div class="py-4 mb-2">
<h3 class="d-inline-block">Edit {{.Name.Proper}}</h3>
</div>

<%= formFor({{.Name.VarCaseSingle}}, {action: {{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), method: "PUT"}) { %>
	<%= partial("{{ .Name.Folder.Pluralize }}/form.html") %>
<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-warning", "data-confirm": "Are you sure?", body: "Cancel"}) %>
<% } %>`

// templateShowTmpl is a variable for the [template_name]/show.plush.html
var templateShowTmpl string = `<div class="py-4 mb-2">
	<h3 class="d-inline-block">{{ .Name.Proper }} Details</h3>
	<div class="float-right">
		<%= linkTo({{ .Name.VarCasePlural }}Path(), {class: "btn btn-info"}) { %>
		Back to all {{ .Name.Group }}
		<% } %>
		<%= linkTo(edit{{ .Name.Proper }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-warning", body: "Edit"}) %>
		<%= linkTo({{ .Name.VarCaseSingle }}Path({ {{ .Name.ParamID }}: {{ .Name.VarCaseSingle }}.ID }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>
	</div>
</div>

<ul class="list-group mb-2 ">
{{- range $p := .Model.Attrs }}
	<li class="list-group-item pb-1">
		<label class="small d-block">{{ $p.Name.Pascalize }}</label>
		<p class="d-inline-block"><%= {{$.Name.VarCaseSingle}}.{{$p.Name.Pascalize}} %></p>
	</li>
{{- end }}
</ul>`

// templateFormTmpl is a variable for the [template_name]/form.plush.html
var templateFormTmpl string = `{{ range $p := .Model.Attrs -}}
{{ if eq $p.CommonType "text" -}}
<%= f.TextAreaTag("{{$p.Name.Pascalize}}", {rows: 10}) %>
{{ else -}}
{{ if eq $p.CommonType "bool" -}}
<%= f.CheckboxTag("{{$p.Name.Pascalize}}", {unchecked: false}) %>
{{ else -}}
<%= f.InputTag("{{$p.Name.Pascalize}}") %>
{{ end -}}
{{ end -}}
{{ end -}}

<button class="btn btn-success" role="submit">Save</button>`

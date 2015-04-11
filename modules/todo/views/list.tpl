<div class="row" style="height:20px;">
</div>
<div class="row">
<div class="col-md-12">
{{ if .flash.notice}}
<div class="alert alert-info" role="alert">
	{{.flash.notice}}	
</div>
{{end}}
<ul class="list-group">
{{ range $todo := .todos}}
	<li class="list-group-item">
		<a href="/todo/{{$todo.Id}}">{{$todo.Title}}</a>
		<a href="/todo/delete/{{$todo.Id}}" style="float:right; margin-top:-4px;" class="btn btn-danger btn-sm">Delete</a>
	</li>
{{end}}  
</ul>
</div>
</div>

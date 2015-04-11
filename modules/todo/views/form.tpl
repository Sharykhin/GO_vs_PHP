<div class="row">
<div class="col-md-12">
<form method="post" action="/todo">
  <div class="form-group">
  <input type="hidden" name="id" value="{{if .todo }}{{.todo.Id}}{{end}}" />
    <label for="exampleInputEmail1">Todo</label>
    <input type="text" name="title" class="form-control" id="todo" placeholder="Write todo" value="{{if .todo }} {{.todo.Title}} {{end}}">
  </div> 
  <button type="submit" class="btn btn-primary">Add</button>
</form>
</div>
</div>
<?php
use yii\helpers\Html;
?>
<div class="row">
    <div class="col-md-12">
        <form method="post" action="/todo/create">
            <div class="form-group">
                <input type="hidden" name="id" value="<?php if(isset($todo)) { echo $todo->id; } ?>" />
                <label for="exampleInputEmail1">Todo</label>
                <input type="text" name="title" class="form-control" id="todo" placeholder="Write todo" value="<?php if(isset($todo)) { echo $todo->title; } ?>">
            </div>
            <button type="submit" class="btn btn-primary">Add</button>
        </form>
    </div>
</div>
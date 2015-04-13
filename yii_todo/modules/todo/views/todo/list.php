<?php include_once 'form.php'; ?>
<div class="row" style="height:20px;">
</div>
<div class="row">
    <div class="col-md-12">
        <?php if(Yii::$app->session->hasFlash('notice')): ?>
        <div class="alert alert-info" role="alert">

            <?= Yii::$app->session->getFlash('notice') ?>

        </div>
        <?php endif; ?>

        <ul class="list-group">
            <?php foreach($todos as $key=>$todo) { ?>
            <li class="list-group-item">
                <a href="/todo/<?php echo $todo->id ?>"><?php echo $todo->title ?></a>
                <a href="/todo/delete/<?php echo $todo->id ?>" style="float:right; margin-top:-4px;" class="btn btn-danger btn-sm">Delete</a>
            </li>
            <?php } ?>
        </ul>
    </div>
</div>
<?php

namespace app\modules\todo\controllers;

use app\modules\todo\models\Todo;
use yii\web\Controller;

class TodoController extends Controller
{
    public $enableCsrfValidation = false;
    public function actionIndex()
    {

        $todos = Todo::find()->all();
        return $this->render('list',['todos'=>$todos]);
    }

    public function actionCreate()
    {
        $todo = new Todo();
        $request = \Yii::$app->request->post();
        if (isset($request['id']) && !empty($request['id'])) {
            $id = intval($request['id']);
            $todo = Todo::find()->where(['id'=>$id])->one();
            $todo->title = $request['title'];
            $todo->save(false);
            \Yii::$app->getSession()->setFlash('notice', 'Todo has been updated');
            return $this->redirect(['index']);
        }
        $todo->title=$request['title'];
        $todo->isdone=false;
        $todo->save(false);
        \Yii::$app->getSession()->setFlash('notice', 'Todo has been created');
        return $this->redirect(['index']);
    }

    public function actionDelete($id)
    {
        $todo  = Todo::find()
        ->where(['id' => $id])
        ->one();
        $todo->delete();
        \Yii::$app->getSession()->setFlash('notice', 'Todo has been deleted');
        return $this->redirect(['index']);
    }

    public function actionUpdate($id)
    {
        $todo  = Todo::find()
            ->where(['id' => $id])
            ->one();
        return $this->render('form',['todo'=>$todo]);
    }
}

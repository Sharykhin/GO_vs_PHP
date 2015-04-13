<?php

namespace app\modules\todo\models;

use Yii;

/**
 * This is the model class for table "todo".
 *
 * @property integer $id
 * @property string $title
 * @property integer $isdone
 */
class Todo extends \yii\db\ActiveRecord
{
    /**
     * @inheritdoc
     */
    public static function tableName()
    {
        return 'todo';
    }

    /**
     * @inheritdoc
     */
    public function rules()
    {
        return [
            [['title', 'isdone'], 'required'],
            [['isdone'], 'integer'],
            [['title'], 'string', 'max' => 255],
            [['title', 'isdone'], 'unique', 'targetAttribute' => ['title', 'isdone'], 'message' => 'The combination of Title and Isdone has already been taken.']
        ];
    }

    /**
     * @inheritdoc
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'title' => 'Title',
            'isdone' => 'Isdone',
        ];
    }
}

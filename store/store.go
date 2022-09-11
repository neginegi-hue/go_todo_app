package store

import (
	"errors"

	"github.com/neginegi-hue/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	//動作確認用の仮実績なのであえてexportしている。
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

//Allはソート済みのタスク一覧を返す
func (ts *TaskStore) All() entity.Tasks {
	Tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks{
		Tasks[i-1] = t
	}
	return Tasks
}

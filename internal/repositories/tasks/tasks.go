package tasks

import (
	"fmt"
	"log"
	"time"

	taskModel "github.com/angmeng/task_app/internal/models/task"
	"github.com/angmeng/task_app/pkg/queryparser"
	"github.com/upper/db/v4"
)

type Repository struct {
	Session db.Session
}

func NewRepository(sess db.Session) *Repository {
	return &Repository{Session: sess}
}

func (repo *Repository) All(rq *queryparser.Query) (taskModel.Pagination, error) {
	log.Printf("query: %#v", rq)

	var (
		tasks      []taskModel.Task
		pagination taskModel.Pagination
		sortBy     = "due_date"
	)

	if rq.Sort != "" {
		sortBy = rq.Sort
	}

	q := repo.Session.SQL().
		SelectFrom("tasks").
		Where(rq.Filter).
		Limit(rq.Size).
		Offset(rq.Size * (rq.Page - 1)).
		OrderBy(sortBy)

	pg := q.Paginate(uint(rq.Size))
	err := pg.Page(uint(rq.Page)).All(&tasks)
	if err != nil {
		return pagination, err
	}

	pagination.Data = tasks
	pagination.Meta.CurrentPage = uint(rq.Page)
	pagination.Meta.TotalResults, err = pg.TotalEntries()
	if err != nil {
		return pagination, err
	}

	pagination.Meta.LastPage, err = pg.TotalPages()
	if err != nil {
		return pagination, err
	}

	return pagination, nil
}

func (repo *Repository) Create(task taskModel.Task) (taskModel.Task, error) {
	table := repo.Session.Collection("tasks")

	if err := table.InsertReturning(&task); err != nil {
		return taskModel.Task{}, err
	}

	return task, nil
}

func (repo *Repository) Update(taskParams taskModel.Task) (taskModel.Task, error) {
	var task taskModel.Task
	table := repo.Session.Collection("tasks")
	res := table.Find(taskParams.ID)
	err := res.One(&task)

	if err != nil {
		return taskModel.Task{}, err
	}

	task.Name = taskParams.Name
	task.Description = taskParams.Description
	task.DueDate = taskParams.DueDate
	task.UpdatedAt = time.Now().UTC()

	if err = res.Update(task); err != nil {
		return taskModel.Task{}, err
	}

	return task, nil
}

func (repo *Repository) Delete(id int) error {

	var task taskModel.Task
	table := repo.Session.Collection("tasks")
	res := table.Find(id)
	err := res.One(&task)

	if err != nil {
		return err
	}

	q := repo.Session.SQL().
		DeleteFrom("tasks").
		Where("id", task.ID)

	_, _ = q.Exec()

	if err = res.Delete(); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetUserTask(userID, taskID int) (taskModel.Task, error) {
	var task taskModel.Task
	query := fmt.Sprintf(`SELECT t.id, t.name, t.description, t.tser_id, t.statts_id, t.dte_date, t.created_at, t.updated_at
                        FROM tasks t WHERE t.id = %d AND t.user_id=%d`, taskID, userID)

	row, err := repo.Session.SQL().Query(query)
	if err != nil {
		return taskModel.Task{}, err
	}

	iter := repo.Session.SQL().NewIterator(row)
	err = iter.One(&task)
	if err != nil {
		return taskModel.Task{}, err
	}

	return task, nil
}

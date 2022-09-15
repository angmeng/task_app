package tasks

import (
	"fmt"

	"github.com/a8m/rql"
	"github.com/angmeng/task_app/internal/models/task"
	taskModel "github.com/angmeng/task_app/internal/models/task"
	taskRepo "github.com/angmeng/task_app/internal/repositories/tasks"
	"github.com/angmeng/task_app/pkg/queryparser"
	"github.com/angmeng/task_app/pkg/validator"
	"github.com/angmeng/task_app/stores"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

var (
	// QueryParam is the name of the query string key.
	QueryParam = "query"
	// MustNewParser panics if the configuration is invalid.
	QueryParser = rql.MustNewParser(rql.Config{
		Model:         taskModel.Task{},
		FieldSep:      ".",
		LimitMaxValue: 50,
	})
)

func Index(c *fiber.Ctx) error {
	sess := stores.ConnectPG()
	defer sess.Close()
	repo := taskRepo.NewRepository(sess)

	rq, err := queryparser.GetDBQuery(c, taskModel.Task{})
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": fmt.Sprintf("failed to getDBQuery, %v", err),
		})
	}

	// should pass the authenticated user ID to find all the belongs records.
	result, err := repo.All(rq)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(result)
}

func Create(c *fiber.Ctx) error {
	var taskParams taskModel.Task

	sess := stores.ConnectPG()
	defer sess.Close()

	repo := taskRepo.NewRepository(sess)

	if err := c.BodyParser(&taskParams); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(taskParams); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			// "message": err.Error(),
			"message": govalidator.ErrorsByField(err),
		})
	}

	// a middleware will verify the access token from client to authenticated the valid user.
	// we will skip this for now.
	// taskParams.UserID = authenticatedUser.ID
	taskParams.StatusID = task.NOT_URGENT
	task, err := repo.Create(taskParams)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(task)
}

func Update(c *fiber.Ctx) error {
	var taskParams taskModel.Task
	sess := stores.ConnectPG()
	defer sess.Close()

	id, _ := c.ParamsInt("id")

	repo := taskRepo.NewRepository(sess)

	if err := c.BodyParser(&taskParams); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(taskParams); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			// "message": err.Error(),
			"message": govalidator.ErrorsByField(err),
			//"message": govalidator.Errors{err},
		})
	}

	// its better to find the record by the authenticated user ID.
	// we will skip this for now.
	// task, err := repo.GetUserTask(authenticatedUser.ID, id)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	taskParams.ID = uint(id)
	task, err := repo.Update(taskParams)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(task)
}

func Delete(c *fiber.Ctx) error {
	sess := stores.ConnectPG()
	defer sess.Close()

	repo := taskRepo.NewRepository(sess)
	id, _ := c.ParamsInt("id")

	// its better to find the record by the authenticated user ID.
	// we will skip this for now.
	// task, err := repo.GetUserTask(authenticatedUser.ID, id)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	err := repo.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task has been deleted successfully",
	})
}

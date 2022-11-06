package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SendTask struct {
	Title    string    `json:"title" db:"title"`
	Status   string    `json:"status" db:"status"`
	Created  time.Time `json:"created" db:"created" gorm:"type:datetime(6)"`
	Modified time.Time `json:"modified" db:"modified" gorm:"type:datetime(6)"`
}

func GetTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := []Task{}
		db.Table("task").Find(&task)
		return c.JSON(http.StatusOK, task)
	}
}

func AddTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := time.Now()
		var st SendTask
		if err := c.Bind(&st); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		task := &SendTask{
			Title:    st.Title,
			Status:   "todo",
			Created:  t,
			Modified: t,
		}

		db.Table("task").Create(&task)
		return c.JSON(http.StatusOK, task)
	}
}

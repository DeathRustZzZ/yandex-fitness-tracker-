package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("ожидалось 3 параметра, получено %d", len(data))
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность тренировки должна быть больше нуля")
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}

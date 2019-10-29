package gitHub

import (
	"bufio"
	"fmt"
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
	"time"
)

// getMilestoneModelJSON Подготавливает JSON для создания milestone
func getMilestoneModelJSON(scanner *bufio.Scanner) (io.Reader, error) {
	milestone := milestoneModel{}
	// Title
	milestone.Title = consoleIO.ReadString("Введите название", scanner)
	// State
	milestone.State = "open"
	// Description
	description, err := consoleIO.ReadByEditor("Сюда введите описание")
	if err != nil {
		return nil, err
	}

	milestone.Description = description
	// Due on
	dueOn := ""
	for {
		dueOn = consoleIO.ReadString(
			fmt.Sprintf("Введите дату и время окончания в формате %q:", "2019-10-21 19:25:00"),
			scanner,
		)

		var date time.Time
		date, err = time.Parse("2006-01-02 15:04:05", dueOn)
		if err == nil {
			dueOn = date.Format("2006-01-02T15:04:05Z")

			break
		}

		log.Printf("Введена некорректная дата. Неободимый формат: %q", "2019-10-21 19:25:00")
	}

	milestone.DueOn = dueOn

	return convertStructToJSON(milestone)
}

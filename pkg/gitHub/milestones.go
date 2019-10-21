package gitHub

import (
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
)

// getMilestoneModelJSON Подготавливает JSON для создания milestone
func getMilestoneModelJSON(context *Context) (io.Reader, error) {
	milestone := MilestoneModel{}

	log.Println("Начинаю наполнять  milestone")
	// Title
	milestone.Title = consoleIO.ReadString("Введите название", context.Scanner)
	// State
	state := ""
	for {
		state = consoleIO.ReadString("Введите статус: open/close:", context.Scanner)
		if state == "open" || state == "close" {
			break
		}

		log.Printf("Введен некорректный статус. возможны только %q и %q\n", "open", "close")
	}

	milestone.State = state
	// Description
	description, err := consoleIO.ReadByEditor(context.Cmd, "Сюда введите содержимое")
	if err != nil {
		return nil, err
	}

	milestone.Description = description

	return convertStructToJSON(milestone)
}

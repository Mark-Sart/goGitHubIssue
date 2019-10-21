package gitHub

import (
	"bufio"
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
	"os/exec"
)

// getMilestoneModelJSON Подготавливает JSON для создания milestone
func getMilestoneModelJSON(scanner *bufio.Scanner, cmd *exec.Cmd) (io.Reader, error) {
	milestone := MilestoneModel{}

	log.Println("Начинаю наполнять  milestone")
	// Title
	milestone.Title = consoleIO.ReadString("Введите название", scanner)
	// State
	state := ""
	for {
		state = consoleIO.ReadString("Введите статус: open/close:", scanner)
		if state == "open" || state == "close" {
			break
		}

		log.Printf("Введен некорректный статус. возможны только %q и %q\n", "open", "close")
	}

	milestone.State = state
	// Description
	description, err := consoleIO.ReadByEditor(cmd, "Сюда введите содержимое")
	if err != nil {
		return nil, err
	}

	milestone.Description = description

	return convertStructToJSON(milestone)
}

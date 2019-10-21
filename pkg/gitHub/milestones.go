package gitHub

import (
	"bufio"
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
)

// getMilestoneModelJSON Подготавливает JSON для создания milestone
func getMilestoneModelJSON(scanner *bufio.Scanner) (io.Reader, error) {
	milestone := MilestoneModel{}

	log.Println("Начинаю наполнять  milestone")
	// Title
	milestone.Title = consoleIO.ReadString("Введите название", scanner)
	// State
	state := ""
	for state != "open" && state != "close" {
		state = consoleIO.ReadString("Введите статус: open/close:", scanner)
	}

	milestone.State = state

	return convertStructToJSON(milestone)
}

package gitHub

import (
	"bufio"
	"fmt"
	"goGitHubIssue/pkg/consoleIO"
	"io"
)

// getMilestoneModelJSON Подготавливает JSON для создания milestone
func getMilestoneModelJSON(scanner *bufio.Scanner) (io.Reader, error) {
	milestone := MilestoneModel{}

	fmt.Println("Начинаю наполнять  milestone")
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

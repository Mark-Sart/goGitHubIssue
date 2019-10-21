package gitHub

import (
	"bufio"
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
	"strings"
)

// getCreateIssueModelJSON Подготавливает JSON для создания issue
func getCreateIssueModelJSON(credentials Credentials, scanner *bufio.Scanner) (io.Reader, error) {
	// Переменные
	var title, body string
	var milestone int
	var milestoneOK, milestoneNeed bool
	var labels, assignees []string

	// Редактор
	cmd, err := consoleIO.InitEditor()
	if err != nil {
		return nil, err
	}

	// Title
	title = consoleIO.ReadString("Введите заголовок:", scanner)
	// Body
	body, err = consoleIO.ReadByEditor(cmd, "Сюда введите содержимое")
	if err != nil {
		return nil, err
	}
	// Milestone
	milestone, err = consoleIO.ReadInt("Введите milestone:", scanner)
	if err == nil {
		milestoneNeed, milestoneOK = true, true
	} else if err == io.EOF {
		milestoneNeed, milestoneOK = false, false
	} else {
		milestoneNeed, milestoneOK = true, false
		log.Println(err)
	}
	// Labels
	labels = consoleIO.ReadList("Введите labels через запятую:", scanner)
	// Assignees
	assignees = consoleIO.ReadList("Введите assignees через запятую:", scanner)

	// Чекаем milestone, если он нужен
	if milestoneNeed && milestoneOK {
		milestoneOK, err = checkMilestone(credentials, milestone)
		if err != nil {
			return nil, err
		}
	}

	// Если не существует milestone, а он нужен, то создаем
	if milestoneNeed && !milestoneOK {
		answer := consoleIO.ReadString("Данного milestone не существует. Создать новый? (Y/n)", scanner)
		answer = strings.ToLower(answer)
		if answer == "y" || answer == "" {
			milestone, err = createMilestone(credentials)
			if err != nil {
				return nil, err
			}

			milestoneOK = true
		} else {
			milestoneNeed = false
		}
	}

	if milestoneNeed {
		createIssueModel := CreateIssueWithMilestoneModel{}
		createIssueModel.Title = title
		createIssueModel.Body = body
		createIssueModel.Labels = labels
		createIssueModel.Assignees = assignees
		createIssueModel.Milestone = milestone

		return convertStructToJSON(createIssueModel)
	} else {
		createIssueModel := CreateIssueModel{}
		createIssueModel.Title = title
		createIssueModel.Body = body
		createIssueModel.Labels = labels
		createIssueModel.Assignees = assignees

		return convertStructToJSON(createIssueModel)
	}
}

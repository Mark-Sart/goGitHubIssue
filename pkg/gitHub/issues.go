package gitHub

import (
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
	"strings"
)

// getCreateIssueModelJSON Подготавливает JSON для создания issue
func getCreateIssueModelJSON(context *Context) (io.Reader, error) {
	// Переменные
	var title, description string
	var milestone int
	var milestoneOK, milestoneNeed bool
	var labels, assignees []string

	// Title
	title = consoleIO.ReadString("Введите заголовок:", context.Scanner)
	// Body
	description, err := consoleIO.ReadByEditor(context.Cmd, "Сюда введите содержимое")
	if err != nil {
		return nil, err
	}
	// Milestone
	milestone, err = consoleIO.ReadInt("Введите milestone:", context.Scanner)
	if err == nil {
		milestoneNeed, milestoneOK = true, true
	} else if err == io.EOF {
		milestoneNeed, milestoneOK = false, false
	} else {
		milestoneNeed, milestoneOK = true, false
		log.Println(err)
	}
	// Labels
	labels = consoleIO.ReadList("Введите labels через запятую:", context.Scanner)
	// Assignees
	assignees = consoleIO.ReadList("Введите assignees через запятую:", context.Scanner)

	// Чекаем milestone, если он нужен
	if milestoneNeed && milestoneOK {
		milestoneOK, err = checkMilestone(context, milestone)
		if err != nil {
			return nil, err
		}
	}

	// Если не существует milestone, а он нужен, то создаем
	if milestoneNeed && !milestoneOK {
		answer := consoleIO.ReadString("Данного milestone не существует. Создать новый? (Y/n)", context.Scanner)
		answer = strings.ToLower(answer)
		if answer == "y" || answer == "" {
			milestone, err = createMilestone(context)
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
		createIssueModel.Description = description
		createIssueModel.Labels = labels
		createIssueModel.Assignees = assignees
		createIssueModel.Milestone = milestone

		return convertStructToJSON(createIssueModel)
	} else {
		createIssueModel := CreateIssueModel{}
		createIssueModel.Title = title
		createIssueModel.Description = description
		createIssueModel.Labels = labels
		createIssueModel.Assignees = assignees

		return convertStructToJSON(createIssueModel)
	}
}

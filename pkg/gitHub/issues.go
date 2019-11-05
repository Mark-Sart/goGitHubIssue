package gitHub

import (
	"bufio"
	"fmt"
	"goGitHubIssue/pkg/consoleIO"
	"io"
	"log"
	"strings"
)

// getCreateIssueModelJSON Подготавливает JSON для создания issue
func getCreateIssueModelJSON(credentials Credentials, scanner *bufio.Scanner) (io.Reader, error) {
	// Переменные
	var title, description string
	var milestone int
	var milestoneOK, milestoneNeed bool
	var labels []string
	var collaborators, correctCollaborators []string

	// Заголовок
	title = consoleIO.ReadString("Введите заголовок:", scanner)

	// Описание
	description, err := consoleIO.ReadByEditor("Сюда введите описание")
	if err != nil {
		return nil, err
	}
	// Labels
	labels = consoleIO.ReadList("Введите labels через запятую:", scanner)
	// Assignees
	assignees = consoleIO.ReadList("Введите assignees через запятую:", scanner)

	// Milestone
	for {
		milestone, err = consoleIO.ReadInt("Введите номер спринта:", scanner)
		if err == nil {
			milestoneNeed, milestoneOK = true, true
		} else if err == io.EOF {
			milestoneNeed, milestoneOK = false, false
		} else {
			milestoneNeed, milestoneOK = true, false
			log.Println(err)
		}

		// Чекаем спринт, если он нужен
		if milestoneNeed && milestoneOK {
			log.Printf("Проверяю спринт № %d\n", milestone)
			milestoneOK, err = checkMilestone(credentials, milestone)
			if err != nil {
				return nil, err
			}
		}

		// Если спринт не существует, а он нужен, то создаем
		if milestoneNeed && !milestoneOK {
			answer := consoleIO.ReadString(
				fmt.Sprintf("Спринт № %d не существует. Создать новый? (Y/n)", milestone),
				scanner,
			)
			answer = strings.ToLower(answer)
			if answer == "y" || answer == "" {
				milestone, err = createMilestone(credentials, scanner)
				if err != nil {
					return nil, err
				}

				log.Printf("Создан спринт № %d", milestone)
				milestoneOK = true
			} else {
				answer = consoleIO.ReadString("Использовать другой спринт? (Y/n)", scanner)
				answer = strings.ToLower(answer)
				if answer == "n" {
					milestoneNeed = false
				}
			}
		}

		if !milestoneNeed || milestoneOK {
			break
		}
	}

	if milestoneNeed && milestoneOK {
		issue := createIssueWithMilestoneModel{}
		issue.Title = title
		issue.Description = description
		issue.Labels = labels
		issue.Assignees = assignees
		issue.Milestone = milestone

		return convertStructToJSON(issue)
	} else {
		issue := createIssueModel{}
		issue.Title = title
		issue.Description = description
		issue.Labels = labels
		issue.Assignees = assignees

		return convertStructToJSON(issue)
	}
}

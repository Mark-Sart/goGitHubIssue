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

	// Метки
	labels = consoleIO.ReadList("Введите метки через запятую:", scanner)

	// Ответственные
	collaborators = consoleIO.ReadList("Введите коллабораторов через запятую:", scanner)
	if len(collaborators) > 0 {
		var users, correctUsers, incorrectUsers []string

		log.Println("Начинаю проверять коллабораторов")
		correctCollaborators, err = checkCollaborators(credentials, collaborators)
		users = getSliceDiff(collaborators, correctCollaborators)

		if len(users) > 0 {
			log.Printf("Не найдены следующие коллаборатры: %s\n", users)
			answer := consoleIO.ReadString("Назначить данных юзеров коллабораторами? (Y/n)", scanner)
			answer = strings.ToLower(answer)
			if answer == "y" || answer == "" {
				var usersToCollaborators []string

				for len(users) > 0 {
					log.Println("Начинаю проверять юзеров")
					correctUsers, err = checkUsers(credentials, users)
					incorrectUsers = getSliceDiff(users, correctUsers)
					usersToCollaborators = append(usersToCollaborators, correctUsers...)

					log.Printf("Следующие юзеры будут назначены коллабораторами: %s\n", usersToCollaborators)

					if len(incorrectUsers) > 0 {
						log.Printf("Не найдены следующие users: %s\n", incorrectUsers)

						answer = consoleIO.ReadString("Назначить других users коллабораторами? (Y/n)", scanner)
						answer = strings.ToLower(answer)
						if answer == "y" || answer == "" {
							users = consoleIO.ReadList("Введите юзеров через запятую", scanner)
						} else {
							break
						}
					}
				}

				var newCollaborators []string
				newCollaborators, err = assignCollaborators(credentials, usersToCollaborators)
				if err != nil {
					return nil, err
				}

				if len(newCollaborators) != len(usersToCollaborators) {
					return nil, fmt.Errorf("что-то пошло не так")
				}

				return nil, fmt.Errorf("отправлены приглашения юзерам %s стать коллабораторами, но создание "+
					"issue невозможно, пока не будут приняты все приглашения", newCollaborators)
			}
		}

		log.Println("Все коллабораторы корректны")
	}

	// Спринт
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
		issue.Assignees = correctCollaborators
		issue.Milestone = milestone

		return convertStructToJSON(issue)
	} else {
		issue := createIssueModel{}
		issue.Title = title
		issue.Description = description
		issue.Labels = labels
		issue.Assignees = correctCollaborators

		return convertStructToJSON(issue)
	}
}

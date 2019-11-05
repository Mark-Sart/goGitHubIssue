package gitHub

import (
	"fmt"
	"io"
	"log"
	"sort"
)

// Проверка коллаборатора
const checkCollaboratorConst = 0

// Проверка юзера
const checkUserConst = 1

// Назначение юзера коллаборатором
const assignCollaboratorConst = 2

// checkUsers Чекает каждого юзера и возвращает корректных
func checkUsers(credentials Credentials, users []string) ([]string, error) {
	return multiRequests(credentials, users, checkUserConst)
}

// checkCollaborators Чекает каждого коллаборатора и возвращает корректных
func checkCollaborators(credentials Credentials, collaborators []string) ([]string, error) {
	return multiRequests(credentials, collaborators, checkCollaboratorConst)
}

// assignCollaborators Назначает юзеров коллабораторами
func assignCollaborators(credentials Credentials, users []string) ([]string, error) {
	return multiRequests(credentials, users, assignCollaboratorConst)
}

// multiRequests Мультизапросы
func multiRequests(credentials Credentials, users []string, operation int) ([]string, error) {
	ch := make(chan operationCollaboratorsModel)
	var correctUsers []string

	for _, user := range users {
		go multiRequestHandler(ch, credentials, user, operation)
	}

	for range users {
		result := <-ch

		if result.err != nil {
			return correctUsers, result.err
		}

		if result.status {
			correctUsers = append(correctUsers, result.name)
		}
	}

	return correctUsers, nil
}

// multiRequestHandler Обработчик мультизапросов
func multiRequestHandler(ch chan operationCollaboratorsModel, credentials Credentials, user string, operation int) {
	log.Printf("Проверяю %s\n", user)
	operationStatus := operationCollaboratorsModel{
		name: user,
	}

	var status bool
	var err error

	switch operation {
	case checkCollaboratorConst:
		if user == credentials.Owner {
			status, err = true, nil
		} else {
			status, err = checkCollaborator(credentials, user)
		}

	case checkUserConst:
		status, err = checkUser(credentials, user)

	case assignCollaboratorConst:
		var body io.Reader
		body, err = getCollaboratorModel()
		if err != nil {
			break
		}

		status, err = assignCollaborator(credentials, user, body)

	default:
		err = fmt.Errorf("Неизвестная операция %d\n", operation)
	}

	if err != nil {
		operationStatus.err = err
		ch <- operationStatus

		return
	}

	operationStatus.status = status
	ch <- operationStatus
}

// getCollaboratorModel Подготавливает JSON для назначения юзера коллаборатором
func getCollaboratorModel() (io.Reader, error) {
	collaborator := collaboratorModel{
		permission: "push",
	}

	return convertStructToJSON(collaborator)
}

// getSliceDiff Возвращает отсутствующие во втором слайсе элементы первого слайса
func getSliceDiff(slice1 []string, slice2 []string) []string {
	var diff []string
	countSlice1 := len(slice2)

	sort.Strings(slice2)

	for _, item := range slice1 {
		if sort.SearchStrings(slice2, item) == countSlice1 {
			diff = append(diff, item)
		}
	}

	return diff
}

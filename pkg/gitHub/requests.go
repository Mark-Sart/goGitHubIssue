package gitHub

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://api.github.com"
const repoURL = "repos"
const userURL = "users"

// CreateIssue Создает issue
func CreateIssue(credentials Credentials, scanner *bufio.Scanner) (int, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/issues", baseURL, repoURL, credentials.Owner, credentials.Repo)

	log.Println("Начинаю наполнять модель issue")
	body, err := getCreateIssueModelJSON(credentials, scanner)
	if err != nil {
		return 0, err
	}

	log.Println("Создаю issue")
	response, err := doRequest(http.MethodPost, url, credentials.Token, body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		issueNumber := &numberModel{}
		err = json.NewDecoder(response.Body).Decode(issueNumber)
		if err != nil {
			return 0, err
		}

		return issueNumber.Value, nil
	}

	return 0, fmt.Errorf(response.Status)
}

// checkMilestone Чекает наличия milestone у репозитория
func checkMilestone(credentials Credentials, milestone int) (bool, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/milestones/%d", baseURL, repoURL, credentials.Owner, credentials.Repo, milestone)

	response, err := doRequest(http.MethodGet, url, credentials.Token, nil)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

// createMilestone Создает milestone
func createMilestone(credentials Credentials, scanner *bufio.Scanner) (int, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/milestones", baseURL, repoURL, credentials.Owner, credentials.Repo)

	log.Println("Начинаю наполнять модель milestone")
	body, err := getMilestoneModelJSON(scanner)
	if err != nil {
		return 0, err
	}

	log.Println("Создаю milestone")
	response, err := doRequest(http.MethodPost, url, credentials.Token, body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		milestoneNumber := &numberModel{}
		err = json.NewDecoder(response.Body).Decode(milestoneNumber)
		if err != nil {
			return 0, err
		}

		return milestoneNumber.Value, nil
	}

	return 0, fmt.Errorf(response.Status)
}

// checkCollaborator Чекает наличие collaborator
func checkCollaborator(credentials Credentials, collaborator string) (bool, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", baseURL, repoURL, credentials.Owner, credentials.Repo, collaborator)

	response, err := doRequest(http.MethodGet, url, credentials.Token, nil)
	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusNoContent {
		return true, nil
	}

	return false, nil
}

// checkUser Чекает наличие user
func checkUser(credentials Credentials, user string) (bool, error) {
	url := fmt.Sprintf("%s/%s/%s", baseURL, userURL, user)

	response, err := doRequest(http.MethodGet, url, credentials.Token, nil)
	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

// assignCollaborator Добавляет user в collaborator
func assignCollaborator(credentials Credentials, user string, body io.Reader) (bool, error) {
	url := fmt.Sprintf(
		"%s/%s/%s/%s/collaborators/%s",
		baseURL,
		repoURL,
		credentials.Owner,
		credentials.Repo,
		user,
	)

	response, err := doRequest(http.MethodPut, url, credentials.Token, body)
	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusNoContent || response.StatusCode == http.StatusCreated {
		return true, nil
	}

	return false, nil
}

// doRequest Выполняет запрос
func doRequest(method, url, token string, body io.Reader) (*http.Response, error) {
	// Собираем запрос
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Устанавливаем токен
	request.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	// Отсылаем запрос
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// convertStructToJSON Конвертирует структуру в JSON-reader
func convertStructToJSON(structure interface{}) (io.Reader, error) {
	// Кодируем в JSON
	issueJSON, err := json.Marshal(structure)
	if err != nil {
		return nil, err
	}

	// Создаем ридер
	reader := bytes.NewReader(issueJSON)

	return reader, nil
}

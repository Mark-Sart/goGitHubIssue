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

const baseURL = "https://api.github.com/repos"

// CreateIssue Создает issue
func CreateIssue(credentials Credentials, scanner *bufio.Scanner) (int, error) {
	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, credentials.Owner, credentials.Repo)

	body, err := getCreateIssueModelJSON(credentials, scanner)
	if err != nil {
		return 0, err
	}

	response, err := doRequest(http.MethodPost, url, credentials.Token, body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		issueNumber := &number{}
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
	url := fmt.Sprintf("%s/%s/%s/milestones/%d", baseURL, credentials.Owner, credentials.Repo, milestone)

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
	url := fmt.Sprintf("%s/%s/%s/milestones", baseURL, credentials.Owner, credentials.Repo)

	body, err := getMilestoneModelJSON(scanner)
	if err != nil {
		return 0, err
	}

	response, err := doRequest(http.MethodPost, url, credentials.Token, body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		milestoneNumber := &number{}
		err = json.NewDecoder(response.Body).Decode(milestoneNumber)
		if err != nil {
			return 0, err
		}

		return milestoneNumber.Value, nil
	}

	return 0, fmt.Errorf(response.Status)
}

// doRequest Выполняет запрос
func doRequest(method, url, token string, body io.Reader) (*http.Response, error) {
	log.Println("Подготавливаю запрос")
	// Собираем запрос
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Устанавливаем токен
	request.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	// Отсылаем запрос
	log.Println("Отсылаю запрос")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	log.Println("Ответ получен")

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

package gitHub

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://api.github.com/repos"

// CreateIssue Создает issue
func CreateIssue(credentials Credentials, scanner *bufio.Scanner) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, credentials.Owner, credentials.Repo)

	body, err := getCreateIssueModelJSON(credentials, scanner)
	if err != nil {
		return nil, err
	}

	return doRequest(http.MethodPost, url, credentials.Token, body)
}

// checkMilestone Чекает наличия milestone у репозитория
func checkMilestone(credentials Credentials, milestone int) (bool, error) {
	url := fmt.Sprintf("%s/%s/%s/milestones/%d", baseURL, credentials.Owner, credentials.Repo, milestone)

	response, err := doRequest(http.MethodGet, url, credentials.Token, nil)
	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

// createMilestone Создает milestone
func createMilestone(credentials Credentials) (int, error) {
	return 0, nil
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

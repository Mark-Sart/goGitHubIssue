package gitHub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://api.github.com/repos"

// CreateIssue Создает issue
func CreateIssue(context *Context) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, context.Credentials.Owner, context.Credentials.Repo)

	body, err := getCreateIssueModelJSON(context)
	if err != nil {
		return nil, err
	}

	return doRequest(http.MethodPost, url, context.Credentials.Token, body)
}

// checkMilestone Чекает наличия milestone у репозитория
func checkMilestone(context *Context, milestone int) (bool, error) {
	url := fmt.Sprintf(
		"%s/%s/%s/milestones/%d",
		baseURL,
		context.Credentials.Owner,
		context.Credentials.Repo,
		milestone,
	)

	response, err := doRequest(http.MethodGet, url, context.Credentials.Token, nil)
	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

// createMilestone Создает milestone
func createMilestone(context *Context) (int, error) {
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

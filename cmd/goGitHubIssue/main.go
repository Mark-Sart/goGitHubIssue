package main

import (
	"bufio"
	"flag"
	"fmt"
	"goGitHubIssue/pkg/consoleIO"
	"goGitHubIssue/pkg/gitHub"
	"log"
	"net/http"
	"os"
)

var mode = flag.String("mode", "", "Операция с issue")
var owner = flag.String("owner", "", "Владелец репозитория")
var repo = flag.String("repo", "", "Название репозитория")

func main() {
	// Парсим флаги
	flag.Parse()

	// Проверяем аргументов
	if *mode == "" {
		log.Fatal("Не указана операция. Используйте -h для получения подробной информации.")
	}

	if *owner == "" {
		log.Fatal("Не указан владелец репозитория. Используйте -h для получения подробной информации.")
	}

	if *repo == "" {
		log.Fatal("Не указан репозиторий. Используйте -h для получения подробной информации.")
	}

	// Сканер
	scanner := bufio.NewScanner(os.Stdin)

	// Редактор
	cmd, err := consoleIO.InitEditor()
	if err != nil {
		log.Fatal(err)
	}

	// Access-token
	token := consoleIO.ReadString("Введите access-token:", scanner)
	if token == "" {
		log.Fatal("Не был введен access-token")
	}

	// Данные о репозитории
	credentials := gitHub.Credentials{
		Owner: *owner,
		Repo:  *repo,
		Token: token,
	}

	// Ответ
	var response *http.Response

	// Логика
	switch *mode {
	case "create":
		response, err = gitHub.CreateIssue(credentials, scanner, cmd)
		if err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println("Доступны только следующие операции:")
		fmt.Printf("%-10s%-s", "create", "создать issue\n")

		return
	}

	// Печатаем результат
	fmt.Println(response)
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"goGitHubIssue/pkg/consoleIO"
	"goGitHubIssue/pkg/gitHub"
	"log"
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
		err := fmt.Errorf("не указана операция")
		errHandler(err)
	}

	if *owner == "" {
		err := fmt.Errorf("не указан владелец репозитория")
		errHandler(err)
	}

	if *repo == "" {
		err := fmt.Errorf("не указан репозиторий")
		errHandler(err)
	}

	// Сканер
	scanner := bufio.NewScanner(os.Stdin)

	// Access-token
	token := consoleIO.ReadString("Введите access-token:", scanner)
	if token == "" {
		err := fmt.Errorf("не был введен access-token")
		errHandler(err)
	}

	// Пользовательские данные
	credentials := gitHub.Credentials{
		Owner: *owner,
		Repo:  *repo,
		Token: token,
	}

	// Логика
	switch *mode {
	case "create":
		number, err := gitHub.CreateIssue(credentials, scanner)
		errHandler(err)

		log.Printf("Создан issue № %d", number)

	default:
		fmt.Println("Доступны только следующие операции:")
		fmt.Printf("%-10s%-s", "create", "создать issue\n")

		return
	}
}

func errHandler(err error) {
	if err != nil {
		log.Fatalf("Ошибка: %s\n", err)
	}
}

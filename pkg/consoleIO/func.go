package consoleIO

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const defaultEditor = "nano"
const filename = ".tmp.txt"

// ReadString Считывает с консоли строку
func ReadString(label string, scanner *bufio.Scanner) string {
	fmt.Println(label)

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

// ReadInt Считывает с консоли число
func ReadInt(label string, scanner *bufio.Scanner) (int, error) {
	fmt.Println(label)

	if scanner.Scan() {
		input := scanner.Text()

		if input != "" {
			milestone, err := strconv.Atoi(input)
			if err != nil {
				return 0, err
			}

			return milestone, nil
		}
	}

	return 0, io.EOF
}

// ReadList Считывает массив строк через запятую
func ReadList(label string, scanner *bufio.Scanner) []string {
	fmt.Println(label)

	list := make([]string, 0, 0)

	if scanner.Scan() {
		input := scanner.Text()
		if input != "" {
			input = strings.Trim(scanner.Text(), ",")
			list = strings.Split(input, ",")

			for idx, item := range list {
				list[idx] = strings.Trim(item, " ")
			}
		}
	}

	return list
}

// InitEditor Настраивает редактор
func InitEditor() (*exec.Cmd, error) {
	// Выбираем редактор
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}

	// Достаем путь до редактора
	path, err := exec.LookPath(editor)
	if err != nil {
		return nil, err
	}

	// Настраиваем редактора
	cmd := exec.Command(path, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil
}

// ReadByEditor Считывает текст через редактор
func ReadByEditor(cmd *exec.Cmd, label string) (string, error) {
	// Создаем временный файл
	err := ioutil.WriteFile(filename, []byte(label), 0755)
	if err != nil {
		return "", err
	}

	// Запускаем редактор
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	// Читаем файл
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Удаляем файла
	err = os.Remove(filename)
	if err != nil {
		return "", err
	}

	return string(inputBytes), nil
}

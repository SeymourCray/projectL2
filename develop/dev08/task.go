package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	builtInCommands = map[string]func([]string){
		"cd":   cd,
		"pwd":  pwd,
		"echo": echo,
		"kill": kill,
		"ps":   ps,
	}
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Ошибка при чтении входных данных:", err.Error())
			continue
		}

		input = strings.TrimSpace(input)

		// проверяем, используются ли конвейеры
		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			runPipelines(commands)
			continue
		}

		// если не используются конвейеры, то ищем встроенные команды
		parts := strings.Split(input, " ")
		command := parts[0]
		args := parts[1:]

		if cmdFunc, ok := builtInCommands[command]; ok {
			cmdFunc(args)
			continue
		}

		// если введенная команда не является встроенной, запускаем ее в качестве внешней команды
		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Команду не удалось выполнить:", err)
		}
	}
}

func runPipelines(commands []string) {
	var previousCmd, cmd *exec.Cmd
	var err error

	for _, command := range commands {
		parts := strings.Split(strings.TrimSpace(command), " ")
		cmd = exec.Command(parts[0], parts[1:]...)

		if previousCmd != nil {

			// Направляем вывод предыдущей команды на ввод текущей
			cmd.Stdin, err = previousCmd.StdoutPipe()
			if err != nil {
				fmt.Println("Ошибка при создании канала вывода:", err)
				return
			}

			if err = previousCmd.Start(); err != nil {
				fmt.Println("Ошибка при запуске команды:", err)
				return
			}
		}

		previousCmd = cmd
	}

	// Запускаем последнюю команду в пайплайне
	cmd.Stdout = os.Stdout
	if err = cmd.Start(); err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}

	if err = cmd.Wait(); err != nil {
		fmt.Println("Ошибка при завершении команды:", err)
		return
	}
}

func cd(args []string) {
	// Если аргументов нет, то значение по умолчанию - домашняя папка
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Не удалось получить домашнюю директорию пользователя:", err)
			return
		}
		if err := os.Chdir(homeDir); err != nil {
			fmt.Println("Не удалось перейти в домашнюю директорию пользователя:", err)
			return
		}
	} else {
		if err := os.Chdir(args[0]); err != nil {
			fmt.Println("Не удалось перейти в заданную директорию:", err)
			return
		}
	}
}

func pwd(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Не удалось получить текущую директорию:", err)
		return
	}
	fmt.Println(dir)
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func kill(args []string) {
	if len(args) == 0 {
		fmt.Println("Необходимо указать ID процесса для завершения")
		return
	}

	pid := args[0]
	if match, _ := regexp.MatchString("^[0-9]+$", pid); !match {
		fmt.Println("ID процесса должен быть числом")
		return
	}

	if err := exec.Command("kill", "-9", pid).Run(); err != nil {
		fmt.Println("Не удалось завершить процесс:", err)
		return
	}

	fmt.Println("Процесс успешно завершен")
}

func ps(args []string) {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Не удалось выполнить команду ps:", err)
	}
}

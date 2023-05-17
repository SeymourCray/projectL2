package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
// для указания размера сообщения при получении
const bufferSize = 100

func main() {
	// ключ таймаута
	timeout := flag.Duration("timeout", time.Second*10, "set timeout for connection")
	flag.Parse()
	// получаем хост и порт из команды
	var host, port string
	if flag.NArg() != 2 {
		log.Fatalln("wrong format")
	} else {
		host, port = flag.Args()[0], flag.Args()[1]
	}
	// подключаемся к серверу
	con, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()
	// контексты для завершения работы через SIGQUIT и завершения работы сервера
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGQUIT)
	defer stop()

	wCtx, wCancel := context.WithCancel(ctx)
	defer wCancel()

	rCtx, rCancel := context.WithCancel(ctx)
	defer rCancel()

	var wg sync.WaitGroup
	wg.Add(2)
	// горутины для чтения и записи
	go write(wCtx, con, &wg)
	go read(rCtx, rCancel, con, &wg)

	select {
	case <-ctx.Done():
		stop()
		wg.Wait()
	case <-rCtx.Done():
		stop()
		wg.Wait()
	}
}

// функция читает из Stdin и отправляет сообщение на сервер
func write(ctx context.Context, c net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			msg := scanner.Bytes()
			if _, err := c.Write(msg); err != nil {
				log.Println("error writing to connection:", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("error reading from stdin:", err)
	}
}

// функция считывает сообщение от сервера и выводит в Stdout
func read(ctx context.Context, cancel context.CancelFunc, c net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(os.Stdout)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg := make([]byte, bufferSize)
			c.SetReadDeadline(time.Now().Add(time.Second * 10))
			n, err := reader.Read(msg)
			if err == io.EOF {
				log.Println("connection closed by server")
				cancel()
				return
			} else if err != nil {
				log.Println("error reading from connection:", err)
				return
			}
			if _, err := writer.Write(append(msg[:n], '\n')); err != nil {
				log.Println("error writing to stdout:", err)
			}
			if err := writer.Flush(); err != nil {
				log.Println("error flushing stdout:", err)
			}
		}
	}
}

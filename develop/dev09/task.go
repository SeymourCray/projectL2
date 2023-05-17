package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

var (
	crawledPages = map[string]bool{}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <url>")
		os.Exit(1)
	}

	urlStr := os.Args[1]
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Некорректный URL:", err)
		os.Exit(1)
	}

	if urlObj.Scheme == "" {
		urlStr = "http://" + urlStr
		urlObj, _ = url.Parse(urlStr)
	}

	parentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		os.Exit(1)
	}

	if err := crawlPage(urlObj, parentDir); err != nil {
		fmt.Println("Ошибка при загрузке сайта:", err)
		os.Exit(1)
	}

	fmt.Println("Сайт успешно загружен:", urlObj.String())
}

func crawlPage(urlObj *url.URL, parentDir string) error {
	// игнорируем уже загруженные страницы
	if crawledPages[urlObj.String()] {
		return nil
	}
	crawledPages[urlObj.String()] = true

	resp, err := http.Get(urlObj.String())
	if err != nil {
		return fmt.Errorf("не удалось получить ответ от сервера: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка при получении данных с сервера: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка при чтении тела ответа: %w", err)
	}

	// формируем путь для сохранения файла
	filePath := filepath.Join(parentDir, urlObj.Host, urlObj.Path)
	if !strings.HasSuffix(filePath, "/") {
		filePath += "/"
	}

	filePath = filepath.Join(filePath, "index.html")

	// создаем директорию для сохранения файла, если она не существует
	if err := os.Mkdir(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("не удалось создать директорию: %w", err)
	}

	// создаем файл для сохранения данных
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	// записываем данные в файл
	if _, err := file.Write(body); err != nil {
		return fmt.Errorf("не удалось записать данные в файл: %w", err)
	}

	// парсим HTML-документ, чтобы найти ссылки
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("ошибка при парсинге HTML-документа: %w", err)
	}

	// обходим все теги <a> и рекурсивно загружаем все ссылки, которые ссылаются на текущий хост
	var crawlErr error
	var visitNode func(*html.Node)
	visitNode = func(node *html.Node) {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				href := attr.Val
				refObj, err := url.Parse(href)
				if err != nil {
					continue
				}

				// игнорируем внешние ссылки и ссылки на ресурсы
				if refObj.Scheme != "http" && refObj.Scheme != "https" ||
					refObj.Host != urlObj.Host ||
					strings.HasPrefix(refObj.Path, "/static/") ||
					strings.HasPrefix(href, "#") {
					continue
				}

				// загружаем ссылку
				if err := crawlPage(refObj, filepath.Dir(filePath)); err != nil {
					crawlErr = err
					return
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visitNode(child)
			if crawlErr != nil {
				return
			}
		}
	}
	visitNode(doc)

	if crawlErr != nil {
		return crawlErr
	}

	return nil
}

/*
Задание
Реализуй конструктор объекта LogSystem с функциональными опциями. Конструктор должен принимать различные опции для настройки логгера, который будет использоваться в системе логирования.

package main

import (
	"fmt"
	"os"
	"io"
)

// Logger interface
type Logger interface {
	Log(message string)
}
// FileLogger struct
// ...

// ConsoleLogger struct
type ConsoleLogger struct{
	out io.ReadWriter
}

// LogOption functional option type
type LogOption func(*LogSystem)

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}
Критерии завершенности:
- Конструктор NewLogSystem принимает функциональные опции и настраивет логгер в соответствии с переданными опциями.
- Опция WithLogger устанавливает переданный логгер в поле logger структуры LogSystem.
- Реализация логгера FileLogger записывет сообщение в файл.
- Метод Log структуры LogSystem вызывает метод Log установленного логгера и передает ему сообщение для логирования.
- При выполнении кода в функции main создается файл log.txt, в этот файл записывается переданное сообщение.
- Покрытие тестами 100%.

*/

package main

import (
	"fmt"
	"io"
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

type LogSystem struct {
	logger Logger
}

func (ls *LogSystem) Log(message string) {
	if ls.logger == nil {
		fmt.Fprintln(os.Stderr, "Ошибка: логгер не определен")
		return
	}
	ls.logger.Log(message)
}

// FileLogger struct
type FileLogger struct {
	file io.Writer
}

func (f *FileLogger) Log(message string) {
	mess := []byte(message + "\n")
	if _, err := f.file.Write(mess); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}

// ConsoleLogger struct
type ConsoleLogger struct {
	out io.ReadWriter
}

func (c *ConsoleLogger) Log(message string) {
	mess := []byte(message + "\n")
	if _, err := c.out.Write(mess); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}

// LogOption functional option type
type LogOption func(*LogSystem)

func NewLogSystem(options ...LogOption) *LogSystem {
	ls := &LogSystem{
		logger: nil,
	}

	for _, option := range options {
		option(ls)
	}
	return ls
}

func WithLogger(logType Logger) LogOption {
	return func(log *LogSystem) {
		log.logger = logType
	}
}

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := &FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}

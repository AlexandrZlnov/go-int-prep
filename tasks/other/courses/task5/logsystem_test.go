package main

import (
	"bytes"
	"os"
	"testing"
)

// Тест FileLogger.Log
func TestFileLogger_Log(t *testing.T) {
	// Используем bytes.Buffer вместо настоящего файла
	// Это позволяет проверить запись без создания файла на диске
	buf := &bytes.Buffer{}
	logger := &FileLogger{file: buf}

	message := "Hello test"
	logger.Log(message)

	expected := message + "\n" // Мы добавляем \n в Log
	if buf.String() != expected {
		t.Errorf("FileLogger.Log записал '%s', ожидается '%s'", buf.String(), expected)
	}
}

// Тест ConsoleLogger.Log
func TestConsoleLogger_Log(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := &ConsoleLogger{out: buf}

	message := "Console test"
	logger.Log(message)

	expected := message + "\n"
	if buf.String() != expected {
		t.Errorf("ConsoleLogger.Log записал '%s', ожидается '%s'", buf.String(), expected)
	}
}

// Тест NewLogSystem и WithLogger
func TestNewLogSystem_WithLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	fileLogger := &FileLogger{file: buf}

	logSys := NewLogSystem(WithLogger(fileLogger))
	if logSys.logger == nil {
		t.Fatal("логгер не был установлен через WithLogger")
	}

	message := "Test LogSystem"
	logSys.Log(message)

	expected := message + "\n"
	if buf.String() != expected {
		t.Errorf("LogSystem.Log записал '%s', ожидается '%s'", buf.String(), expected)
	}
}

// Тест nil-safe LogSystem.Log
func TestLogSystem_Log_NilLogger(t *testing.T) {
	// Создаем LogSystem без логгера
	logSys := NewLogSystem()

	// Перехватываем stderr
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	message := "Test nil logger"
	logSys.Log(message)

	// Закрываем pipe и читаем stderr
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stderr = oldStderr

	if buf.Len() == 0 {
		t.Error("Ожидалось сообщение об ошибке в stderr при nil logger")
	}
}

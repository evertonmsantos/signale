package signale

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Logger é a estrutura principal para o logger personalizado
type Logger struct {
	UseEmoji     bool // Define se os emojis serão usados
	UseTimestamp bool // Define se o timestamp será exibido
}

// New cria uma nova instância do Logger
func New(useEmoji, useTimestamp bool) *Logger {
	return &Logger{UseEmoji: useEmoji, UseTimestamp: useTimestamp}
}

// formatHeader formata o cabeçalho com base nas configurações do Logger
func (l *Logger) formatHeader() string {
	if l.UseTimestamp {
		return fmt.Sprintf("[%s] ", time.Now().Format("15:04:05"))
	}
	return ""
}

// Success imprime uma mensagem de sucesso com verde mais claro
func (l *Logger) Success(msg string) {
	emoji := ""
	if l.UseEmoji {
		emoji = "✅ "
	}
	lightGreen := color.New(color.FgHiGreen, color.Underline).SprintFunc()
	fmt.Printf("%s%s%s %s\n", l.formatHeader(), emoji, lightGreen("success"), msg)
}

// Error imprime uma mensagem de erro com vermelho mais escuro
func (l *Logger) Error(msg string) {
	emoji := ""
	if l.UseEmoji {
		emoji = "❌ "
	}
	darkRed := color.New(color.FgHiRed, color.Underline).SprintFunc()
	fmt.Printf("%s%s%s %s\n", l.formatHeader(), emoji, darkRed("error"), msg)
}

// Warning imprime uma mensagem de aviso com amarelo mais claro
func (l *Logger) Warning(msg string) {
	emoji := ""
	if l.UseEmoji {
		emoji = "⚠️ "
	}
	lightYellow := color.New(color.FgHiYellow, color.Underline).SprintFunc()
	fmt.Printf("%s%s%s %s\n", l.formatHeader(), emoji, lightYellow("warning"), msg)
}

// Info imprime uma mensagem informativa com azul mais claro
func (l *Logger) Info(msg string) {
	emoji := ""
	if l.UseEmoji {
		emoji = "ℹ️ "
	}
	lightBlue := color.New(color.FgHiBlue, color.Underline).SprintFunc()
	fmt.Printf("%s%s%s %s\n", l.formatHeader(), emoji, lightBlue("info"), msg)
}

// Custom imprime uma mensagem personalizada com cor e emoji definidos
func (l *Logger) Custom(msgType, emoji string, colorCode color.Attribute, msg string) {
	prefix := ""
	if l.UseEmoji {
		prefix = emoji + " "
	}
	customColor := color.New(colorCode, color.Underline).SprintFunc()
	fmt.Printf("%s%s%s %s\n", l.formatHeader(), prefix, customColor(msgType), msg)
}

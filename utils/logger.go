package utils

import (
	"github.com/fatih/color"
	"log"
)

func Info(text string) {
	white := color.WhiteString("\tINFO\t")
	log.Println(white, text)
}

func FatalError(text string) {
	red := color.RedString("\tERROR\t")
	log.Fatalln(red, text)
}

func Error(text string) {
	red := color.RedString("\tERROR\t")
	log.Println(red, text)
}

func Success(text string) {
	green := color.GreenString("\tSUCCESS\t")
	log.Println(green, text)
}

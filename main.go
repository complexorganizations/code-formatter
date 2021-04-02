package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var codePath string

func init() {
	if len(os.Args) > 1 {
		codePath = os.Args[1]
	} else {
		os.Exit(0)
	}
}

func main() {
	if fileExists(codePath) {
		formatFile()
	} else if folderExists(codePath) {
		formatDirectory()
	}
}

func formatFile() {
	switch filepath.Ext(codePath) {
	case ".go":
		installCheck("go")
		formatGoFiles(codePath)
	case ".sh", ".bash":
		installCheck("shfmt")
		formatShellScriptFiles(codePath)
	case ".html":
		installCheck("html-minifier")
		formatHTMLFiles(codePath)
	case ".css":
		installCheck("csso-cli")
		formatCSSFiles(codePath)
	case ".js":
		installCheck("uglifyjs")
		formatJSFiles(codePath)
	default:
		log.Println("Error:", codePath)
	}
}

func formatDirectory() {
	filepath.Walk(codePath, func(path string, info os.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".go":
			installCheck("go")
			formatGoFiles(path)
		case ".sh", ".bash":
			installCheck("shfmt")
			formatShellScriptFiles(codePath)
		case ".html":
			installCheck("html-minifier")
			formatHTMLFiles(codePath)
		case ".css":
			installCheck("csso-cli")
			formatCSSFiles(codePath)
		case ".js":
			installCheck("uglifyjs")
			formatJSFiles(codePath)
		default:
			log.Println("Error:", codePath)
		}
		return nil
	})
}

func formatGoFiles(filePath string) {
	cmd := exec.Command("go", "fmt", filePath)
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func formatShellScriptFiles(filePath string) {
	cmd := exec.Command("shfmt", "-l -w", filePath)
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func formatHTMLFiles(filePath string) {
	cmd := exec.Command("html-minifier")
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func formatCSSFiles(filePath string) {
	cmd := exec.Command("csso-cli")
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func formatJSFiles(filePath string) {
	cmd := exec.Command("uglifyjs")
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func installCheck(appName string) {
	if !commandExists(appName) {
		log.Fatalf("The application %s was not found in the system. \n", appName)
	}
}

func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func commandExists(cmd string) bool {
	appName, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = appName // Var declared and not used
	return true
}

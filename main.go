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
	// Golang
	case ".go":
		installCheck("go")
		formatGoFiles(codePath)
	// Shell Script
	case ".sh":
		installCheck("shfmt")
		formatShellScriptFiles(codePath)
	// HTML
	case ".html":
		installCheck("html-minifier")
		formatHTMLFiles(codePath)
	// CSS
	case ".css":
		installCheck("cssnano")
		formatCSSFiles(codePath)
	// JavaScript
	case ".js":
		installCheck("uglifyjs")
		formatJSFiles(codePath)
	// Python
	case ".py":
		//
	// Java
	case ".java":
		//
	// C++
	case ".cpp":
		//
	// C
	case ".c":
		//
	// TypeScript
	case ".ts":
		//
	// PHP
	case ".php":
		//
	// Kotlin
	case ".kts":
		//
	// Ruby
	case ".rb":
		//
	// Visual Basic
	case ".vba":
		//
	// Swift
	case ".swift":
		//
	// Rust
	case ".rs":
		//
	// Scala
	case ".scala":
		//
	// Dart
	case ".dart":
		//
	// PowerShell
	case ".ps1":
		//
	// JSON
	case ".json":
		//
	// Markdown
	case ".md":
		//
	// Julia
	case ".jl":
		//
	// Yaml
	case ".yaml":
		//
	}
}

func formatDirectory() {
	filepath.Walk(codePath, func(path string, info os.FileInfo, err error) error {
		switch filepath.Ext(path) {
		// Golang
		case ".go":
			installCheck("go")
			formatGoFiles(codePath)
		// Shell Script
		case ".sh":
			installCheck("shfmt")
			formatShellScriptFiles(codePath)
		// HTML
		case ".html":
			installCheck("html-minifier")
			formatHTMLFiles(codePath)
		// CSS
		case ".css":
			installCheck("cssnano")
			formatCSSFiles(codePath)
		// JavaScript
		case ".js":
			installCheck("uglifyjs")
			formatJSFiles(codePath)
		// Python
		case ".py":
			//
		// Java
		case ".java":
			//
		// C++
		case ".cpp":
			//
		// C
		case ".c":
			//
		// TypeScript
		case ".ts":
			//
		// PHP
		case ".php":
			//
		// Kotlin
		case ".kts":
			//
		// Ruby
		case ".rb":
			//
		// Visual Basic
		case ".vba":
			//
		// Swift
		case ".swift":
			//
		// Rust
		case ".rs":
			//
		// Scala
		case ".scala":
			//
		// Dart
		case ".dart":
			//
		// PowerShell
		case ".ps1":
			//
		// JSON
		case ".json":
			//
		// Markdown
		case ".md":
			//
		// Julia
		case ".jl":
			//
		// Yaml
		case ".yaml":
			//
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
	cmd := exec.Command("html-minifier", filePath, "-o", filePath)
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func formatCSSFiles(filePath string) {
	cmd := exec.Command("cssnano", filePath, filePath)
	cmd.Run()
	fmt.Println("Enhancing:", filePath)
}

func formatJSFiles(filePath string) {
	cmd := exec.Command("uglifyjs", "-b", "--", filePath, "-o", filePath)
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
	_ = appName // variable declared and not used
	return true
}

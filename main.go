package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	codePath string
	err      error
)

func init() {
	if len(os.Args) > 1 {
		tempCodePath := flag.String("path", "/user/example/folder/file", "The file's location in the system.")
		flag.Parse()
		codePath = *tempCodePath
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
	// System path
	if len(codePath) < 1 || codePath == "/user/example/folder/file" {
		log.Fatal("Error: The system path has not been given.")
	}
}

func main() {
	if fileExists(codePath) {
		// Format a single file
		formatFile()
	} else if folderExists(codePath) {
		// Format a whole folder
		formatDirectory()
	}
}

func formatFile() {
	switch filepath.Ext(codePath) {
	// Golang
	case ".go":
		formatGoFiles(codePath)
	// Shell Script
	case ".sh":
		formatShellScriptFiles(codePath)
	// HTML
	case ".html":
		formatHTMLFiles(codePath)
	// CSS
	case ".css":
		formatCSSFiles(codePath)
	// JavaScript
	case ".js":
		formatJSFiles(codePath)
	// TypeScript
	case ".ts":
		formatTypeScriptFiles(codePath)
	// Python
	case ".py":
		formatPythonFiles(codePath)
	// Java
	case ".java":
		formatJavaFiles(codePath)
	// c##
	case ".cs":
		formatCSFiles(codePath)
	// C++
	case ".cpp":
		formatCPPFiles(codePath)
	// C
	case ".c":
		formatCFiles(codePath)
	// PHP
	case ".php":
		formatPHPFiles(codePath)
	// Kotlin
	case ".kts":
		formatKotlinFiles(codePath)
	// Ruby
	case ".rb":
		formatRubyFiles(codePath)
	// Swift
	case ".swift":
		formatSwiftFiles(codePath)
	// Rust
	case ".rs":
		formatRustFiles(codePath)
	// Scala
	case ".scala":
		formatScalaFiles(codePath)
	// Dart
	case ".dart":
		formatDartFiles(codePath)
	// PowerShell
	case ".ps1":
		formatPowerShellFiles(codePath)
	// Yaml
	case ".yaml":
		formatYamlFiles(codePath)
	// JSON
	case ".json":
		formatJSONFiles(codePath)
	// Markdown
	case ".md":
		formatMarkdownFiles(codePath)
	default:
		log.Println("Error:", codePath)
	}
}

func formatDirectory() {
	filepath.Walk(codePath, func(path string, info os.FileInfo, err error) error {
		switch filepath.Ext(path) {
		// Golang
		case ".go":
			formatGoFiles(path)
		// Shell Script
		case ".sh":
			formatShellScriptFiles(path)
		// HTML
		case ".html":
			formatHTMLFiles(path)
		// CSS
		case ".css":
			formatCSSFiles(path)
		// JavaScript
		case ".js":
			formatJSFiles(path)
		// TypeScript
		case ".ts":
			formatTypeScriptFiles(path)
		// Python
		case ".py":
			formatPythonFiles(path)
		// Java
		case ".java":
			formatJavaFiles(path)
		// c##
		case ".cs":
			formatCSFiles(path)
		// C++
		case ".cpp":
			formatCPPFiles(path)
		// C
		case ".c":
			formatCFiles(path)
		// PHP
		case ".php":
			formatPHPFiles(path)
		// Kotlin
		case ".kts":
			formatKotlinFiles(path)
		// Ruby
		case ".rb":
			formatRubyFiles(path)
		// Swift
		case ".swift":
			formatSwiftFiles(path)
		// Rust
		case ".rs":
			formatRustFiles(path)
		// Scala
		case ".scala":
			formatScalaFiles(path)
		// Dart
		case ".dart":
			formatDartFiles(path)
		// PowerShell
		case ".ps1":
			formatPowerShellFiles(path)
		// Yaml
		case ".yaml":
			formatYamlFiles(path)
		// JSON
		case ".json":
			formatJSONFiles(path)
		// Markdown
		case ".md":
			formatMarkdownFiles(path)
		default:
			log.Println("Error:", path)
		}
		return nil
	})
}

// Golang
func formatGoFiles(filePath string) {
	installCheck("go")
	cmd := exec.Command("go", "fmt", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Shell Script
func formatShellScriptFiles(filePath string) {
	installCheck("shfmt")
	cmd := exec.Command("shfmt", "-l -w", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// HTML
func formatHTMLFiles(filePath string) {
	installCheck("html-minifier")
	cmd := exec.Command("html-minifier", filePath, "-o", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// CSS
func formatCSSFiles(filePath string) {
	installCheck("cssnano")
	cmd := exec.Command("cssnano", filePath, filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// JS
func formatJSFiles(filePath string) {
	installCheck("uglifyjs")
	cmd := exec.Command("uglifyjs", "-b", "--", filePath, "-o", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Python
func formatPythonFiles(filePath string) {
	installCheck("yapf")
	cmd := exec.Command("yapf", "--in-place", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Java
func formatJavaFiles(filePath string) {
	installCheck("google-java-format")
	cmd := exec.Command("google-java-format", "--replace", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// C++
func formatCPPFiles(filePath string) {
	installCheck("clang-format")
	cmd := exec.Command("clang-format", "-style=Google", "-i", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// C##
func formatCSFiles(filePath string) {
	installCheck("dotnet-format")
	cmd := exec.Command("dotnet-format", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// C
func formatCFiles(filePath string) {
	installCheck("clang-format")
	cmd := exec.Command("clang-format", "-style=Google", "-i", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// TypeScript
func formatTypeScriptFiles(filePath string) {
	installCheck("gts")
	cmd := exec.Command("gts", "fix", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// PHP
func formatPHPFiles(filePath string) {
	installCheck("phptidy")
	cmd := exec.Command("phptidy", "replace", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Kotlin
func formatKotlinFiles(filePath string) {
	installCheck("ktlint")
	cmd := exec.Command("ktlint", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Ruby
func formatRubyFiles(filePath string) {
	installCheck("rufo")
	cmd := exec.Command("rufo", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Swift
func formatSwiftFiles(filePath string) {
	installCheck("swift-format")
	cmd := exec.Command("swift-format", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Rust
func formatRustFiles(filePath string) {
	installCheck("rustfmt")
	cmd := exec.Command("rustfmt", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Scala
func formatScalaFiles(filePath string) {
	installCheck("scalafmt")
	cmd := exec.Command("scalafmt", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Dart
func formatDartFiles(filePath string) {
	installCheck("dart")
	cmd := exec.Command("dart", "format", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// PowerShell
func formatPowerShellFiles(filePath string) {
	installCheck("Edit-DTWBeautifyScript")
	cmd := exec.Command("Edit-DTWBeautifyScript", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// JSON
func formatJSONFiles(filePath string) {
	installCheck("json-format")
	cmd := exec.Command("json-format", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Markdown
func formatMarkdownFiles(filePath string) {
	installCheck("mdformat")
	cmd := exec.Command("mdformat", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Yaml
func formatYamlFiles(filePath string) {
	installCheck("yamllint")
	cmd := exec.Command("yamllint", filePath)
	err = cmd.Run()
	handleErrors(err)
	log.Println("Optimizing:", filePath)
}

// Application Check
func installCheck(appName string) {
	if !commandExists(appName) {
		log.Printf("The application %s was not found in the system. \n", appName)
	}
}

// Directory Check
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// File Check
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Application Check
func commandExists(cmd string) bool {
	appName, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = appName
	return true
}

// Handle errors
func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}

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
	// Visual Basic
	case ".vba":
		formatVBAFiles(codePath)
	// Swift
	case ".swift":
		formatSwiftFiles(codePath)
	// R
	case ".r":
		formatRFiles(codePath)
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
	// Julia
	case ".jl":
		formatJuliaFiles(codePath)
	// Yaml
	case ".yaml":
		formatYamlFiles(codePath)
	// JSON
	case ".json":
		formatJSONFiles(codePath)
	// Markdown
	case ".md":
		formatMarkdownFiles(codePath)
	}
}

func formatDirectory() {
	filepath.Walk(codePath, func(path string, info os.FileInfo, err error) error {
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
		// Visual Basic
		case ".vba":
			formatVBAFiles(codePath)
		// Swift
		case ".swift":
			formatSwiftFiles(codePath)
		// R
		case ".r":
			formatRFiles(codePath)
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
		// Julia
		case ".jl":
			formatJuliaFiles(codePath)
		// Yaml
		case ".yaml":
			formatYamlFiles(codePath)
		// JSON
		case ".json":
			formatJSONFiles(codePath)
		// Markdown
		case ".md":
			formatMarkdownFiles(codePath)
		}
		return nil
	})
}

func formatGoFiles(filePath string) {
	installCheck("go")
	cmd := exec.Command("go", "fmt", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatShellScriptFiles(filePath string) {
	installCheck("shfmt")
	cmd := exec.Command("shfmt", "-l -w", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatHTMLFiles(filePath string) {
	installCheck("html-minifier")
	cmd := exec.Command("html-minifier", filePath, "-o", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatCSSFiles(filePath string) {
	installCheck("cssnano")
	cmd := exec.Command("cssnano", filePath, filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatJSFiles(filePath string) {
	installCheck("uglifyjs")
	cmd := exec.Command("uglifyjs", "-b", "--", filePath, "-o", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatPythonFiles(filePath string) {
	installCheck("yapf")
	cmd := exec.Command("yapf", "--in-place", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatJavaFiles(filePath string) {
	installCheck("google-java-format")
	cmd := exec.Command("google-java-format", "--replace", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatCPPFiles(filePath string) {
	installCheck("dotnet-format")
	cmd := exec.Command("dotnet-format", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatCSFiles(filePath string) {
	//
}

func formatCFiles(filePath string) {
	installCheck("clang-format")
	cmd := exec.Command("clang-format", "-style=Google", "-i", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatTypeScriptFiles(filePath string) {
	installCheck("gts")
	cmd := exec.Command("gts", "fix", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatPHPFiles(filePath string) {
	//
}

func formatKotlinFiles(filePath string) {
	//
}

func formatRubyFiles(filePath string) {
	installCheck("rufo")
	cmd := exec.Command("rufo", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatVBAFiles(filePath string) {
	//
}

func formatSwiftFiles(filePath string) {
	installCheck("swift-format")
	cmd := exec.Command("swift-format", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatRFiles(filePath string) {
	//
}

func formatRustFiles(filePath string) {
	installCheck("rustfmt")
	cmd := exec.Command("rustfmt", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatScalaFiles(filePath string) {
	installCheck("scalafmt")
	cmd := exec.Command("scalafmt", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatDartFiles(filePath string) {
	installCheck("dart")
	cmd := exec.Command("dart", "format", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatPowerShellFiles(filePath string) {
	//
}

func formatJSONFiles(filePath string) {
	installCheck("json-format")
	cmd := exec.Command("json-format", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatMarkdownFiles(filePath string) {
	installCheck("mdformat")
	cmd := exec.Command("mdformat", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
}

func formatJuliaFiles(filePath string) {
	//
}

func formatYamlFiles(filePath string) {
	installCheck("yamllint")
	cmd := exec.Command("yamllint", filePath)
	cmd.Run()
	fmt.Println("Optimizing:", filePath)
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

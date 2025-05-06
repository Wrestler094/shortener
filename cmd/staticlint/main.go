// Package main запускает кастомный multichecker для статического анализа проекта.
package main

import (
	"strings"

	"github.com/Wrestler094/shortener/cmd/staticlint/noosexit"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"

	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilness"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unusedresult"

	"honnef.co/go/tools/staticcheck"
)

const (
	saPrefix = "SA"
	st1000   = "ST1000"
)

func main() {
	analyzers := getAnalyzers()
	multichecker.Main(analyzers...)
}

func getAnalyzers() []*analysis.Analyzer {
	var analyzers []*analysis.Analyzer

	// Добавляем стандартные анализаторы
	analyzers = append(analyzers, getStandardAnalyzers()...)

	// Добавляем staticcheck анализаторы
	analyzers = append(analyzers, getStaticcheckAnalyzers()...)

	// Добавляем кастомный анализатор
	analyzers = append(analyzers, noosexit.Analyzer)

	return analyzers
}

func getStandardAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		// Анализаторы для проверки безопасности
		copylock.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,

		// Анализаторы для проверки форматирования и структуры
		printf.Analyzer,
		structtag.Analyzer,

		// Анализаторы для проверки логики
		unreachable.Analyzer,
		unmarshal.Analyzer,
		assign.Analyzer,
		nilness.Analyzer,
		shadow.Analyzer,
		unusedresult.Analyzer,
	}
}

func getStaticcheckAnalyzers() []*analysis.Analyzer {
	var analyzers []*analysis.Analyzer

	for _, a := range staticcheck.Analyzers {
		if a.Analyzer == nil {
			continue
		}
		if strings.HasPrefix(a.Analyzer.Name, saPrefix) || a.Analyzer.Name == st1000 {
			analyzers = append(analyzers, a.Analyzer)
		}
	}

	return analyzers
}

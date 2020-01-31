package structural

import (
	"fmt"
)

// LegacyPrinter exported to be used globally
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter exported to be used globally
type MyLegacyPrinter struct{}

// Print exported to be used globally
func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s\n", s)
	return newMsg
}

// NewPrinter exported to be used globally
type NewPrinter interface {
	PrintStored() string
}

// PrinterAdapter exported to be used globally
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

// Print exported to be used globally
func (p *PrinterAdapter) Print() string {
	if p.OldPrinter != nil {
		newMsg := fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
		return newMsg
	}
	return p.Msg
}

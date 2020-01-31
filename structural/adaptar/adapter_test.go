package structural

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.Print()
	if returnedMsg == "Legacy Printer: Adapter: Hello World!\n" {
		t.Log("Used legacy printer to print the message")
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.Print()
	if returnedMsg == "Hello World!" {
		t.Log("Used new printer to print the message")
	}
}

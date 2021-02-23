package main

import (
	"fmt"
	two "github.com/gkarthiks/import-proto-module/src"
	one "github.com/gkarthiks/moduled-proto/src"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	importedModuleOne := &one.MessageOne{
		Message: "Message set in imported module one proto from module two",
	}

	messageTwo := &two.MessageTwo{
		MessageTwo:         "Direct message in module 2",
		ImportedMessageOne: importedModuleOne,
	}

	fmt.Println("Printing the module two's proto")
	fmt.Println(protojson.Format(messageTwo))
}

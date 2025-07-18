// cmd/artificialeaselainext/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaselainext/internal/artificialeaselainext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaselainext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}

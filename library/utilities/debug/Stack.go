package debug

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Stack() {
	str := fmt.Sprintf("%s", debug.Stack())
	fmt.Println("[STACK TRACE]", strings.Replace(str, "\n", "|", -1))
}

package tictacgo

import "fmt"

const RESET = "\033[0m"
const RED = "\033[31m"
const BLUE = "\033[34m"
const SAVE_POSITION = "\033[s"
const RESTORE_POSITION = "\033[u"
const ERASE_LINE = "\033[2K"
const ERASE_LINE_TO_END = "\033[K"
const MOVE_UP_N = "\033[%dA"
const MOVE_FORWARD_N = "\033[%dC"

func ansiUpN(n int) string {
	return fmt.Sprintf(MOVE_UP_N, n)
}

func ansiForwardN(n int) string {
	return fmt.Sprintf(MOVE_FORWARD_N, n)
}

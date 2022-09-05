package readline

// Extreme gratitude and total acknowledgement
// to GitHub user `gohxs`, who created this work as per
// https://github.com/gohxs/readline/commit/26cb607ab5e014ba59864d69f4cde26d28f064ce
// We have simply transcribed this work.

import (
	"bytes"
	"fmt"
	"runtime"
)

func (r *RuneBuffer) cursorPosition() []byte {

	if runtime.GOOS != "darwin" && runtime.GOOS != "windows" {
		return r.getBackspaceSequence()
	}

	fullWidth := runes.WidthAll(r.buf) + r.promptLen()             // full line Width
	lineCount := LineCount(r.width-1, fullWidth) - 1               // Total line count
	cursorWidth := (runes.WidthAll(r.buf[:r.idx])) + r.promptLen() // cursor line Width
	desiredLine := (cursorWidth / r.width)                         // get Line position starting from prompt
	desiredCol := cursorWidth % r.width                            // get column position starting from prompt
	tbuf := bytes.NewBuffer(nil)

	if lineCount > 0 || fullWidth == r.width {
		tbuf.WriteString(fmt.Sprintf("\033[%dA", lineCount))
	}
	tbuf.WriteString("\r") // go back anyway
	// Position cursor
	if desiredLine > 0 {
		tbuf.WriteString(fmt.Sprintf("\033[%dB", desiredLine)) // Reset
	}
	if desiredCol > 0 {
		tbuf.WriteString(fmt.Sprintf("\033[%dC", desiredCol)) // Reset
	}

	return tbuf.Bytes()

}

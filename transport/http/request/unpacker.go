package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	httpErrs "github.com/storybuilder/storybuilder/transport/http/errors"
	"github.com/storybuilder/storybuilder/transport/http/request/unpackers"
)

// Pre-compiled regexes for message formatting — compiled once at startup.
var (
	reNewLine     = regexp.MustCompile(`[\r\n]+`)
	reSpecialChar = regexp.MustCompile(`[\t"']*`)
)

// Unpack the request in to the given unpacker struct.
func Unpack(r *http.Request, unpacker unpackers.Unpacker) error {
	err := json.NewDecoder(r.Body).Decode(unpacker)
	if err != nil {
		return httpErrs.NewValidationError(formatUnpackerMessage(unpacker.RequiredFormat()))
	}
	return nil
}

// formatUnpackerMessage removes any special characters from the message string.
func formatUnpackerMessage(p string) string {
	m := reSpecialChar.ReplaceAllString(reNewLine.ReplaceAllString(p, " "), "")
	return fmt.Sprintf("Required format: %s", m)
}

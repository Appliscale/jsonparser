package jsonparser

// Errors
var (
	keyPathNotFoundError       = 1
	unknownValueTypeError      = 2
	malformedJsonError         = 3
	malformedStringError       = 4
	malformedArrayError        = 5
	malformedObjectError       = 6
	malformedValueError        = 7
	malformedStringEscapeError = 8
)

type ParserError struct {
	offset int
	s string
	errorType int
}

func (err *ParserError) Error() string {
	return err.s
}

func (err *ParserError) Offset() int {
	return err.offset
}

func MalformedObjectError(offset int) error {
	return &ParserError{s: "Value looks like object, but can't find closing '}' symbol", offset: offset, errorType: malformedObjectError}
}

func UnknownValueTypeError(offset int) error {
	return &ParserError{s:"Unknown value type", offset: offset, errorType: unknownValueTypeError}
}

func KeyPathNotFoundError(offset int) error {
	return &ParserError{s: "Key path not found", offset: offset, errorType: keyPathNotFoundError}
}

func MalformedJsonError(offset int) error {
	return &ParserError{s: "Malformed JSON error", offset: offset, errorType: malformedJsonError}
}

func MalformedStringError(offset int) error {
	return &ParserError{s: "Value is string, but can't find closing '\"' symbol", offset: offset, errorType: malformedStringError}
}

func MalformedArrayError(offset int) error {
	return &ParserError{s: "Value is array, but can't find closing ']' symbol", offset: offset, errorType: malformedArrayError}
}

func MalformedValueError(offset int) error {
	return &ParserError{s: "Value looks like Number/Boolean/None, but can't find its end: ',' or '}' symbol", offset: offset, errorType: malformedValueError}
}

func MalformedStringEscapeError(offset int) error {
	return &ParserError{s: "Encountered an invalid escape sequence in a string", offset: offset, errorType: malformedStringEscapeError}
}

func addToErrorOffset(err error, offset int) error {
	if parserError, ok := err.(*ParserError); ok {
		parserError.offset = parserError.offset + offset
	}
	return err
}

func isKeyPathNotFoundError(err error) bool {
	return isParserErrorOfType(err, keyPathNotFoundError)
}

func isParserErrorOfType(err error, errorType int) bool {
	parserError, ok := err.(*ParserError)
	return ok && parserError.errorType == errorType
}
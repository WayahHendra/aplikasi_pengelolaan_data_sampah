package utils

const (
	Success = "\033[32m[SUCCESS]\033[0m "
	Info    = "\033[34m[INFO]\033[0m "
	Warning = "\033[33m[WARNING]\033[0m "
	Error   = "\033[31m[ERROR]\033[0m "
)

var Types = map[string]string{
	"success": Success,
	"info":    Info,
	"warning": Warning,
	"error":   Error,
}

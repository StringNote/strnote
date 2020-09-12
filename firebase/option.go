package firebase

import "google.golang.org/api/option"

var (
	opt option.ClientOption
)

// GetOption は admin.json を読んでOptionを生成します。
func GetOption() option.ClientOption {
	if opt == nil {
		opt = option.WithCredentialsFile("admin.json")
	}
	return opt
}

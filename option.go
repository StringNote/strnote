package main

import "google.golang.org/api/option"

var (
	opt option.ClientOption
)

func getOption() option.ClientOption {
	if opt == nil {
		opt = option.WithCredentialsFile("admin.json")
	}
	return opt
}

package slogAttr

import "log/slog"

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func OpInfo(op string) slog.Attr {
	return slog.Attr{
		Key:   "operation",
		Value: slog.StringValue(op),
	}
}

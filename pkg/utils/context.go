package utils

import "context"

type contextKey struct{ name string }

var contextKeyName = &contextKey{ConstContextRequestID}

func SetContextRequestID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextKeyName, value)
}

func GetContextRequestID(ctx context.Context) string {
	v, _ := ctx.Value(contextKeyName).(string)
	return v
}

package jsonrpc

import "context"

type RequestContext struct {
	HeaderMap map[string]string
}

const LogKeyRequestHeader = "request_header"

func WithContextRequestHeader(ctx context.Context, headerMap map[string]string) context.Context {
	return context.WithValue(ctx, LogKeyRequestHeader, headerMap)
}

func GetContextRequestHeader(ctx context.Context) map[string]string {
	val := ctx.Value(LogKeyRequestHeader)
	if val != nil {
		if headerMap, ok := val.(map[string]string); ok {
			return headerMap
		}
	}
	return nil
}

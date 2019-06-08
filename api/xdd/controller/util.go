package controller

import (
	"context"
	"strings"

	"github.com/micro/go-micro/metadata"

	"microx/api/xdd/middleware/tracer"
	mxcontext "microx/pkg/context"
)

func toContext(mctx *mxcontext.MxContext) context.Context {
	ctx, _ := tracer.ContextWithSpan(mctx.Context())

	mda, _ := metadata.FromContext(ctx)
	md := metadata.Copy(mda)

	// set headers
	for k, v := range mctx.Context().Request.Header {
		if _, ok := md[k]; !ok {
			md[k] = strings.Join(v, ",")
		}
	}

	ctx = metadata.NewContext(ctx, md)
	return ctx
}

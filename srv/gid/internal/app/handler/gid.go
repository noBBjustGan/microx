package handler

import (
	"context"
	"microx/srv/gid/internal/infra/idgen"

	"microx/pkg/log"
	gid "microx/srv/gid/api"
)

type Gid struct{}

func (this *Gid) GetOne(ctx context.Context, req *gid.Request, rsp *gid.Response) error {
	log.Info("Received Gid.Create request")
	rsp.Id = idgen.GetOne()
	return nil
}

func (this *Gid) GetMulti(ctx context.Context, req *gid.MultiRequest, rsp *gid.MultiResponse) error {
	log.Info("Received Gid.Create request")
	rsp.Ids = idgen.GetMulti(int(req.Count))

	return nil
}

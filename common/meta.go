package common

import (
	"context"
	"errors"
	"strconv"

	"github.com/micro/go-micro/metadata"

	"microx/common/typ"
	"microx/pkg/log"
)

func GetHeaderFromContext(ctx context.Context) (*typ.Header, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		log.Error("metadata.FromContext error")
		return nil, errors.New("metadata.FromContext error")
	}

	var err error
	header := typ.Header{}
	header.Token = md["Token"]
	if header.UserId, err = strconv.ParseInt(md["User-Id"], 10, 64); err != nil {
		log.Error(err)
		return nil, err
	}
	header.AppId, err = strconv.Atoi(md["App-Id"])
	if err != nil {
		log.Error(err)
		return nil, err
	}
	header.AppVersion = md["App-Version"]
	header.OsType = md["Os-Type"]
	header.OsVersion = md["Os-Version"]
	header.Resolution = md["Resolution"]
	header.Model = md["Model"]
	header.Channel = md["Channel"]
	header.Net = md["Net"]
	header.DeviceId = md["Device-Id"]

	return &header, nil
}

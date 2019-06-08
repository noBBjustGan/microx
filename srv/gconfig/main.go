package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"
	"google.golang.org/grpc"
)

var (
	apps    []string
	runMode string
)

type Service struct{}

func init() {
	flag.StringVar(&runMode, "run_mode", "dev", "Run mode")
}

func main() {
	flag.Parse()
	// load config files
	err := loadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	// new service
	service := grpc.NewServer()
	proto.RegisterSourceServer(service, new(Service))
	ts, err := net.Listen("tcp", ":9600")
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("configServer started")
	err = service.Serve(ts)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) Read(ctx context.Context, req *proto.ReadRequest) (rsp *proto.ReadResponse, err error) {
	appName := parsePath(req.Path)
	rsp = &proto.ReadResponse{
		ChangeSet: getConfig(appName),
	}
	return
}

func (s Service) Watch(req *proto.WatchRequest, server proto.Source_WatchServer) (err error) {
	appName := parsePath(req.Path)
	rsp := &proto.WatchResponse{
		ChangeSet: getConfig(appName),
	}
	if err = server.Send(rsp); err != nil {
		log.Logf("[Watch] watch files error，%s", err)
		return err
	}

	return
}

func loadConfigFile() (err error) {
	if err := config.Load(
		file.NewSource(file.WithPath("./conf/app.yml")),
	); err != nil {
		return err
	}
	if runMode == "" {
		runMode = config.Get("run_mode").String("")
	}

	apps = config.Get("apps").StringSlice([]string{})
	for _, app := range apps {
		if err := config.Load(file.NewSource(
			file.WithPath("./conf/" + runMode + "/" + app + ".yml"),
		)); err != nil {
			log.Fatalf("[loadConfigFile] load files error，%s", err)
			return err
		}
	}

	// watch changes
	watcher, err := config.Watch()
	if err != nil {
		log.Fatalf("[loadConfigFile] start watching files error，%s", err)
		return err
	}

	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadConfigFile] watch files error，%s", err)
				return
			}

			log.Logf("[loadConfigFile] file change， %s", string(v.Bytes()))
		}
	}()

	return
}

func getConfig(appName string) *proto.ChangeSet {
	bytes := config.Get(appName).Bytes()

	log.Logf("[getConfig] appName: %s", appName)
	return &proto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix()}
}

func parsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}

package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	v1 "grpc_demo/api/service/v1"
	"grpc_demo/server"
)

type Config struct {
	GRPCPort            string
	DataStoreDBHost     string
	DataStoreDBUser     string
	DataStoreDBPort     string
	DataStoreDBPassword string
	DataStoreDBSchema   string
}

func RunServer() error {
	ctx := context.Background()
	var cfg Config

	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DataStoreDBHost, "db-host", "", "db host")
	flag.StringVar(&cfg.DataStoreDBUser, "db-user", "", "db user")
	flag.StringVar(&cfg.DataStoreDBPort, "db-port", "", "db port")
	flag.StringVar(&cfg.DataStoreDBPassword, "db-password", "", "db password")
	flag.StringVar(&cfg.DataStoreDBSchema, "db-schema", "", "db schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server:%s ", cfg.GRPCPort)
	}

	params := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.DataStoreDBUser, cfg.DataStoreDBPassword, cfg.DataStoreDBHost, cfg.DataStoreDBPort, cfg.DataStoreDBSchema, params)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("数据库连接失败: %s", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)
	return server.RunServer(ctx, v1API, cfg.GRPCPort)

}

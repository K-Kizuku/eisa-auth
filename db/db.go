package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/K-Kizuku/eisa-auth/db/sql/query"
	env "github.com/K-Kizuku/eisa-auth/pkg/config"
	"github.com/jackc/pgx/v5"
)

func New() *query.Queries {
	db, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	queries := query.New(db)
	return queries
}

func Init() (*pgx.Conn, error) {
	if env.Mode == "dev" {
		conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@db:5432/example?sslmode=disable")
		if err != nil {
			return nil, err
		}
		return conn, err
	} else if env.Mode == "prod" {
		// d, err := cloudsqlconn.NewDialer(context.Background())
		// if err != nil {
		// 	return nil, err
		// }

		// var opts []cloudsqlconn.DialOption

		// mysql.RegisterDialContext("cloudsqlconn",
		// 	func(ctx context.Context, addr string) (net.Conn, error) {
		// 		return d.Dial(ctx, config.InstanceConnectionName, opts...)
		// 	},
		// )

		dsn := fmt.Sprintf("user=%s password=%s database=%s", env.DBUser, env.DBPassword, env.DBName)
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			return nil, err
		}
		var opts []cloudsqlconn.Option
		// if usePrivate != "" {
		// 	opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
		// }
		opts = append(opts, cloudsqlconn.WithCredentialsFile("./key.json"))
		d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
		if err != nil {
			return nil, err
		}
		config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
			return d.Dial(ctx, env.InstanceConnectionName)
		}
		// dbURI := stdlib.RegisterConnConfig(config)
		// dbPool, err := sql.Open("pgx", dbURI)
		// if err != nil {
		// 	return nil, fmt.Errorf("sql.Open: %w", err)
		// }

		pgxConn, err := pgx.ConnectConfig(context.Background(), config)
		if err != nil {
			return nil, fmt.Errorf("pgx.ConnectConfig: %w", err)
		}
		return pgxConn, err

	}
	return nil, errors.New("mode is invalid")
}

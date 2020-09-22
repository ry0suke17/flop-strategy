package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yneee/flop-strategy/transport/openapi/api"
	"github.com/yneee/flop-strategy/transport/server"

	"github.com/yneee/flop-strategy/domain/service"

	"github.com/yneee/flop-strategy/domain/repository/flsdb/flspostgres"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/infra/flsflag"
)

var (
	postgresDatabaseURL     = flag.String("postgres_database_url", "", "PostgreSQL のデータベース URL")
	postgresMaxOpenConns    = flag.Int("postgres_max_open_conns", 0, "PostgreSQL へのオープン接続の最大数を表す")
	postgresMaxIdleConns    = flag.Int("postgres_max_idle_conns", 5, "PostgreSQL へのアイドル接続の最大数を表す")
	postgresConnMaxLifetime = flag.Duration("postgres_conn_max_lifetime", 60*time.Second, "PostgreSQL への接続が再利用される最大時間を表す")
)

func flagParse() {
	flsflag.Parse("FLS", "_")
}

func do() (err error) {
	flagParse()

	// 依存をセットアップする {
	db, err := flspostgres.NewClient(
		*postgresDatabaseURL,
		*postgresMaxOpenConns,
		*postgresMaxIdleConns,
		*postgresConnMaxLifetime,
	)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := db.Close()
		if closeErr != nil {
			if err == nil {
				err = flserr.Wrap(closeErr)
			} else {
				err = flserr.Wrapf(err, "%v", closeErr)
			}
		}
	}()
	// }

	// サービスを生成する {
	flopStrtategyService := service.NewFlopStrtategyService(db)
	// }

	// サーバーを起動する {
	flopStrtategyServer := server.NewFlopStrategyServer(flopStrtategyService)
	router := api.NewRouter(flopStrtategyServer)
	// The web process must listen for HTTP traffic on $PORT
	// https://devcenter.heroku.com/articles/container-registry-and-runtime#dockerfile-commands-and-runtime {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// }
	return http.ListenAndServe(":"+port, router)
	// }
}

func main() {
	if err := do(); err != nil {
		log.Printf("err=%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

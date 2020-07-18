package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yneee/flop-strategy/app/flopsituationapp"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/repository/flsdb/flspostgres"
)

var (
	postgresSourceName      = flag.String("postgres_source_name", "", "PostgreSQL データソース名")
	postgresMaxOpenConns    = flag.Int("postgres_max_open_conns", 0, "PostgreSQL へのオープン接続の最大数を表す")
	postgresMaxIdleConns    = flag.Int("postgres_max_idle_conns", 5, "PostgreSQL へのアイドル接続の最大数を表す")
	postgresConnMaxLifetime = flag.Duration("postgres_conn_max_lifetime", 60*time.Second, "PostgreSQL への接続が再利用される最大時間を表す")
)

func do() (err error) {
	flag.Parse()

	// 依存をセットアップする {
	db, err := flspostgres.NewClient(
		*postgresSourceName,
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

	// アプリケーションを生成する {
	flopSituationApp := flopsituationapp.NewApp(db)
	// }

	// テスト的にデータ取得 {
	list, err := flopSituationApp.ListFlopSituations(context.Background())
	if err != nil {
		log.Printf("err=%v\n", err)
		os.Exit(1)
	}
	fmt.Println(list)
	// }

	return nil
}

func main() {
	if err := do(); err != nil {
		log.Printf("err=%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

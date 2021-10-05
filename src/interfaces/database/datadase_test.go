package database_test

import (
	"testing"
	_ "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	_ "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	dbInfra "server/infrastructure/database"
	dbInter "server/interfaces/database"
)

//insert {duckduckgo.com, true} , {example.com, false}

func TestExists(t *testing.T) {^M
	urls := map[string]struct {
			url     string
			want    bool
	}{
			"duckExist": {url: "duckduckgo.com", want: true},
			"examExist": {url: "example.com", want: true},
			"failNotExist" : {url: "failed.failed", want: false},
	}

	sqlh := dbInfra.NewSqlHandler()
	repo := dbInter.NewSiteRepository(sqlh)

	for state, data := range urls {
			t.Run(state, func(t *testing.T) {
					b, err := repo.Exists(data.url)

					if err != nil {
							t.Errorf("unexpected Error %v", err)
					}
					if data.want != b {
							t.Errorf("url %s expect %v but get %v", data.url, data.want, b)
					}
			})
	}
}

func TestFindByURL(t *testing.T) {^M
	urls := map[string]struct {
			url     string
			gpc     bool
	}{
			"duckExist": {url: "duckduckgo.com", gpc: true},
			"examExist": {url: "example.com", gpc: false},
	}

	sqlh := dbInfra.NewSqlHandler()
	repo := dbInter.NewSiteRepository(sqlh)

	for state, data := range urls {
			t.Run(state, func(t *testing.T) {
					s, err := repo.FindByURL(data.url)
					if err != nil {
							t.Errorf("unexpected Error %v", err)
					}
					if s.GPC().Enable != data.gpc {
							t.Errorf("url %s expect %v but get %v", data.url, data.gpc, s.GPC().Enable)
					}
			})
	}
}

func TestStore(t *testing.T) {
	url := "insert.insert"
	flag := false
	sites, _ := site.NewSite(url, gpc.NewGpc(flag))

	sqlh := dbInfra.NewSqlHandler()
	repo := database.NewSiteRepository(sqlh)

	err := repo.Store(sites)
	if err != nil {
			t.Errorf("unexpected Error %v", err)
	}
	check, _ := repo.Exists(url)
	if check != true {
			t.Errorf("can't insert %s", url)
	}
	find, _ := repo.FindByURL(url)
	if find.GPC().Enable != flag {
			t.Errorf("%s insert data is %v but get %v", url, flag, find.GPC().Enable)
	}

}
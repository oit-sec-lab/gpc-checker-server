package database

import (
    "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
    "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
)

type SiteRepository struct {
    SqlHandler
}

func (repo *SiteRepository) Store(s site.Site) (err error) {
    _, err = repo.Execute(
        "INSERT INTO sites (url, gpc) VALUES (?,?)", s.URL(), s.GPC(),
    )
    if err != nil {
        return
    }
    return
}

func (repo *SiteRepository) FindByURL(identifier string) (s site.Site, err error) {
    row, err := repo.Query("SELECT id, url, gpc FROM sites WHERE url = ?", identifier)
    defer row.Close()
    if err != nil {
        return
    }
    var id int
    var url string
    var gpc gpc.Gpc
    row.Next()
    if err = row.Scan(&id, &url, &gpc); err != nil {
        return
    }

    s, err = site.NewSite(id, url, gpc)
    if err != nil {
        return
    }
    return
}

func (repo *SiteRepository) Exists(identifier string) (find bool, err error) {
	row, err := repo.Query("SELECT id FROM sites WHERE url = ?", identifier)
    if err != nil {
        return
    }
    defer row.Close()

    //if result had row, row.Next() would be true
	if row.Next() {
        find = true
    } else {
        find = false
    }
    return
}
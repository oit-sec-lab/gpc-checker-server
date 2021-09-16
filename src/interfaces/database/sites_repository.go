package database

import (
    "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
    "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
)

type SiteRepository struct {
    SqlHandler
}

func (repo *SiteRepository) Store(s site.Site) (err error) {
    tmpgpc := s.GPC()
    inurl := s.URL()
    ingpc := tmpgpc.Enable

    _, err = repo.Execute("INSERT INTO sites (url, gpc) VALUES (?,?)", inurl, ingpc,)
    if err != nil {
        return
    }
    return
}

func (repo *SiteRepository) FindByURL(identifier string) (s site.Site, err error) {
    row, err := repo.Query("SELECT url, gpc FROM sites WHERE url = ?", identifier)
    if err != nil {
        return
    }
    defer row.Close()

    var inurl string
    var ingpc gpc.Gpc
    var tmpgpc bool
    row.Next()
    row.Scan(&inurl, &tmpgpc)

    if err != nil {
        return
    }

    ingpc = gpc.NewGpc(tmpgpc)
    s, err = site.NewSite( inurl, ingpc)
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
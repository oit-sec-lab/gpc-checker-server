package database

import "../../domain"


type SiteRepository struct {
    SqlHandler
}

func (repo *SiteRepository) Store(s domain.Site) (err error) {
    result, err := repo.Execute(
        "INSERT INTO sites (url, gpc) VALUES (?,?)", s.url, s.gpc,
    )
    if err != nil {
        return
    }
    return
}

func (repo *SiteRepository) FindByURL(identifier string) (site domain.Site, err error) {
    row, err := repo.Query("SELECT id, url, gpc FROM sites WHERE url = ?", identifier)
    defer row.Close()
    if err != nil {
        return
    }
    var id int
    var url string
    var gpc bool
    row.Next()
    if err = row.Scan(&id, &url, &gpc); err != nil {
        return
    }
    site.id = id
    site.url = url
    site.gpc = gpc
    return
}
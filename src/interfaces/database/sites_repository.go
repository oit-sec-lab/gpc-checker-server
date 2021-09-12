package database

import "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"

type SiteRepository struct {
    SqlHandler
}

func (repo *SiteRepository) Store(s site.Site) (err error) {
    result, err := repo.Execute(
        "INSERT INTO sites (url, gpc) VALUES (?,?)", s.url, s.gpc,
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
    var gpc bool
    row.Next()
    if err = row.Scan(&id, &url, &gpc); err != nil {
        return
    }
    s.id = id
    s.url = url
    s.gpc = gpc
    return
}

func (repo *SiteRepository) Exist(identifier string) (find bool, err error) {
	var id int
	err := repo.QueryRow("SELECT id FROM sites WHERE url = ?", identifier).Scan(&id)
	switch {
	case err != nil :
		return
	case err == sql.ErrNoRows :
		find = false
	default :
		find = true
	}
	return find
}
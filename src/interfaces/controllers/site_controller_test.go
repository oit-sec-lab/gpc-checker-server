package controllers

import(
	"github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptest"
    "testing"
	"strings"
    sqlInfra "github.com/oit-sec-lab/gpc-checker-server/src/infrastructure/database"
    httpInfra "github.com/oit-sec-lab/gpc-checker-server/src/infrastructure/network"
)

const URLArrayString string = `[{"id":1,"url":"https://example.net/"},{"id":2,"url":"https://example.org/"},{"id":3,"url":"https://duckduckgo.com/"}]`

func TestNewSiteController(t *testing.T){
    site := NewSiteController(sqlInfra.NewSqlHandler(),httpInfra.NewHttpHandler())
    if site == nil{
        t.Fatalf("SiteController is empty.")
    }
}

func TestSiteController(t *testing.T){
    ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
    req, _ := http.NewRequest("POST", "/sites", strings.NewReader(URLArrayString))
    req.Header.Set("Content-Type", "application/json")
    ginContext.Request = req
    site := NewSiteController(sqlInfra.NewSqlHandler(),httpInfra.NewHttpHandler())
    _, err := site.VerifyGPC(ginContext)
    if err != nil{
        t.Fatal(err)
    }
}

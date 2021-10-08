package json

import(
    entitiesJson "github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/json"
)

func GenerateURLJsonArray() ([]entitiesJson.URL_Json){
    var r []entitiesJson.URL_Json
    return r
}

func GenerateRetJson() (entitiesJson.Ret_Json){
    var r entitiesJson.Ret_Json
    return r
}

func GenerateRetJsonArray() ([]entitiesJson.Ret_Json){
    var r []entitiesJson.Ret_Json
    return r
}

func MakeRetJson(id int, url string, gpc bool)(entitiesJson.Ret_Json){
    var r entitiesJson.Ret_Json
    r.Id = id
    r.Url = url
    r.Gpc = gpc
    return r
}

func GetURL(j entitiesJson.URL_Json)(u string){
    return entitiesJson.URL(j)
}

func GetID(j entitiesJson.URL_Json)(i int){
    return entitiesJson.ID(j)
}

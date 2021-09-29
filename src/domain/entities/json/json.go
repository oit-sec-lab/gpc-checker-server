package json

type URL_Json struct{
    Id int `json:"id"`
    Url string `json:"url"`
}

type Ret_Json struct{
    Id int `json:"id"`
    Url string `json:"url"`
    Gpc bool `json:"gpc"`
}

type URL_Json_array []URL_Json

type Ret_Json_array []Ret_Json

func New_Ret_Json_array() Ret_Json_array{
	return Ret_Json_array{}
}

func ID(j URL_Json) int{
    return j.Id
}

func URL(j URL_Json) string{
    return j.Url
}

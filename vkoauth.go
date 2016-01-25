package vkoauth

import (
	"encoding/json"
	"net/http"
)

type VkAuth struct {
	Link         string //Ccылка на авторизацию
	ClientId     string //Идентификатор Вашего приложения.
	RedirectUri  string //Адрес, на который будет переадресован пользователь после прохождения авторизации (домен указанного адреса должен соответствовать основному домену в настройках приложения и перечисленным значениям в списке доверенных redirect uri - адреса сравниваются вплоть до path-части).
	Display      string //Указывает тип отображения страницы авторизации
	Scope        string //Битовая маска настроек доступа приложения, которые необходимо проверить при авторизации пользователя и запросить, в случае отсутствия необходимых.
	ResponseType string //Тип ответа, который Вы хотите получить. Укажите code, чтобы осуществляеть запросы со стороннего сервера.
	Vers         string //Версия API
	SecretApp    string //Секретный ключ приложения
}
type VkClient struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	UserId      int64  `json:"user_id"`
}

var Vk *VkAuth

func NewVkAuth(w http.ResponseWriter, r *http.Request) *VkClient {
	errorValue := r.FormValue("error")
	code := r.FormValue("code")
	if len(errorValue) > 0 {
		panic(errorValue)
		w.Write([]byte(errorValue))
		return nil
	}
	if len(code) > 0 {
		url_token := "https://oauth.vk.com/access_token?client_id=" + Vk.ClientId + "&client_secret=" + Vk.SecretApp + "&redirect_uri=" + Vk.RedirectUri + "&code=" + code
		resp, err := http.Get(url_token)
		if err != nil {
			w.Write([]byte(err.Error()))
			return nil
		}
		var Client VkClient
		err = json.NewDecoder(resp.Body).Decode(&Client)
		if err != nil {
			w.Write([]byte(err.Error()))
			return nil
		}
		if Client.UserId == 0 {
			return nil
		}
		return &Client
	} else {
		http.Redirect(w, r, Vk.Link+"client_id="+Vk.ClientId+"&display="+Vk.Display+"&redirect_uri="+Vk.RedirectUri+"&scope="+Vk.Scope+"&response_type="+Vk.ResponseType+"&v="+Vk.Vers, 302)
		return nil
	}
}

func AssignVk(vk *VkAuth) *VkAuth {
	Vk = vk
	return Vk
}

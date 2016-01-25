# go-vkweb
VkOAuth пример использования API вконтакте на вашем сайте, авторизация, получение access_token и отправка запросов на API VK

Пример использования со стандартным веб сервером
```
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"vk/vkoauth"
)

var vk *vkoauth.VkAuth

func UsersGet(AccessToken string) string {
	ApiUrl := "https://api.vk.com/method/users.get?v=" + vk.Vers + "&access_token=" + AccessToken
	resp, err := http.Get(ApiUrl)
	if err != nil {
		log.Println(err.Error())
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	return string(body)
}
func main() {
	vk = vkoauth.AssignVk(&vkoauth.VkAuth{Link: "https://oauth.vk.com/authorize?", SecretApp: "fpeivftcJjRMSVGYtmBp", ClientId: "5240839", RedirectUri: "http://localhost:8080/vk", Display: "page", Scope: "status,friends", ResponseType: "code", Vers: "5.44"})
	http.HandleFunc("/vk", loginHandler)
	http.ListenAndServe(":8080", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	vkClient := vkoauth.NewVkAuth(w, r)
	if vkClient != nil {
		w.Write([]byte(UsersGet(vkClient.AccessToken)))
	} else {
		w.Write([]byte("Error auth"))
	}

}
```

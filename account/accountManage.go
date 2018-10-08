package account

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJpc3MiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJuYmYiOjE1MzU0NDE5OTUsInN1YiI6InVzZXIiLCJ1c2VybmFtZSI6ImFkbWluIn0.-agKC7UzJL6fOordtJb7qVBwbGeUa8TW__C0LHlniXw"
	IPAddressOfAPIServer = "192.168.3.109:8080"
)

type BelongOrg struct {
	OrganizationID string `json:"organizationID"`
	IsAdmin bool `json:"isAdmin"`
}

type UserToBeCreated struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	SessionTTL int `json:"sessionTTL"`
	IsAdmin bool `json:"isAdmin"`
	IsActive bool `json:"isActive"`
	Email string `json:"email"` //this field has to be created
	Note string `json:"note,omitempty"`
	Organizations []BelongOrg `json:"organizations,omitempty"`
}

type UserForUpdate struct {
	ID string `json:"id,omitempty"`
	UserName string `json:"username"` //this filed has to be created
	SessionTTL int `json:"sessionTTL,omitempty"`
	IsAdmin bool `json:"isAdmin,omitempty"`
	IsActive bool `json:"isActive,omitempty"`
	Email string `json:"email,omitempty"`
	Note string `json:"note,omitempty"`
}

type UserPasswordChange struct {
	ID string `json:"id"`
	Password string `json:"password"`
}

type externalUser struct {
	ID string
	UserName string
	SessionTTL int
	IsAdmin bool
	IsActive bool
	CreatedAt string
	UpdatedAt string
	Email string
	Note string
}

type UserList struct {
	TotalCount int
	Result []externalUser
}

func makeRequest(url, method string, body io.Reader) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{Transport:tr}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Grpc-Metadata-Authorization", token)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	return response
}

func CreateUser(newUser UserToBeCreated) string {
	b, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range b {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	body := bytes.NewBuffer(b)
	url := "https://" + IPAddressOfAPIServer + "/api/users"
	resp := makeRequest(url, "POST", body)
	defer resp.Body.Close()

	var userID struct{
		ID string
	}
	if resp.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyByte, &userID)
	}else {
		fmt.Println(resp.Status)
	}

	return userID.ID
	//if false, how to return and how to handle ?

}

func DeleteUser(userID string) bool {
	url := "https://" + IPAddressOfAPIServer + "/api/users/" + userID
	resp := makeRequest(url, "DELETE", nil)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	}else {
		return false
	}
}

func UpdateUser(updatedUser UserForUpdate, userID string) bool {
	b, err := json.Marshal(updatedUser)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range b {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	body := bytes.NewBuffer(b)
	url := "https://" + IPAddressOfAPIServer + "/api/users/" + userID
	resp := makeRequest(url, "PUT", body)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	}else {
		return false
	}
}

func ChangeUserPassword(userID string, newPasswd UserPasswordChange) bool {
	b, err := json.Marshal(newPasswd)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewBuffer(b)
	url := "https://" + IPAddressOfAPIServer + "/api/users/" + userID + "/password"
	resp := makeRequest(url, "PUT", body)
	if resp.StatusCode == 200 {
		return true
	}else {
		return false
	}
}

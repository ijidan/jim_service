package pkg

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var path = "https://sm.ms/api/v2/"

type Token struct {
	Token string `json:"token"`
}
type TokenRsp struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data   Token `json:"data"`
	RequestId string `json:"RequestId"`
}
type UserProfile struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	GroupExpire   string `json:"group_expire"`
	EmailVerified int    `json:"email_verified"`
	DiskUsage     string `json:"disk_usage"`
	DiskLimit     string `json:"disk_limit"`
	DiskUsageRaw  int    `json:"disk_usage_raw"`
	DiskLimitRaw  int64  `json:"disk_limit_raw"`
}
type UserProfileRsp struct {
	Success   bool        `json:"success"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      UserProfile `json:"data"`
	RequestId string      `json:"RequestId"`
}

type Image struct {
	FileId    int    `json:"file_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	Storename string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
	Page      string `json:"page"`
}

type UploadImageRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Data      Image  `json:"data"`
	RequestId string `json:"RequestId"`
}

type SmMs struct {
}

func (s *SmMs) GetToken(userName string, password string) (string, error) {
	client := resty.New()
	url := fmt.Sprintf("%s%s", path, "token")
	tokenRsp := &TokenRsp{}
	_, err := client.R().
		SetHeaders(map[string]string{"Accept": "application/json"}).
		ForceContentType("application/json").
		SetFormData(map[string]string{"username": userName, "password": password}).
		SetResult(tokenRsp).
		Post(url)
	if err != nil {
		return "", err
	}
	if tokenRsp.Success ==false{
		return "",errors.New(tokenRsp.Message)
	}
	return tokenRsp.Data.Token, nil
}

func (s *SmMs) GetUserProfile(token string) (UserProfile, error) {
	client := resty.New()
	url := fmt.Sprintf("%s%s", path, "profile")
	userProfileRsp := &UserProfileRsp{}
	_, err := client.R().
		SetHeaders(map[string]string{"Accept": "application/json", "Content-Type": "multipart/form-data"}).
		ForceContentType("application/json").
		SetAuthScheme("Basic").
		SetAuthToken(token).
		SetResult(userProfileRsp).Post(url)
	if err != nil {
		return UserProfile{}, err
	}
	if userProfileRsp.Success ==false{
		return UserProfile{},errors.New(userProfileRsp.Message)
	}
	return userProfileRsp.Data, nil
}

func (s *SmMs) UploadImage(token string, filePath string) (Image, error) {
	client := resty.New()
	url := fmt.Sprintf("%s%s", path, "upload")
	uploadImageRsp := &UploadImageRsp{}
	_, err := client.R().
		SetHeaders(map[string]string{"Accept": "application/json", "Content-Type": "multipart/form-data"}).
		ForceContentType("application/json").
		SetAuthScheme("Basic").
		SetAuthToken(token).
		SetFile("smfile",filePath).
		SetResult(uploadImageRsp).
		Post(url)
	if err != nil {
		return Image{}, err
	}
	if uploadImageRsp.Success==false{
		return Image{},errors.New(uploadImageRsp.Message)
	}
	return uploadImageRsp.Data, nil
}

func NewSmMs() *SmMs  {
	smMs:=&SmMs{}
	return smMs
}


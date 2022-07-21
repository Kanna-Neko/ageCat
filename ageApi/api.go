package ageapi

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

func (handle *AgeHandle) login() error {
	doc, err := goquery.NewDocument("https://www.agemys.cc/accounts/login/")
	if err != nil {
		return err
	}
	var csrf, _ = doc.Find(".account_form>input").First().Attr("value")
	_, err = handle.client.R().SetFormData(map[string]string{
		"csrfmiddlewaretoken": csrf,
		"username":            handle.account,
		"password":            handle.password,
	}).Post("https://www.agemys.cc/accounts/login/")
	if err != nil {
		return err
	}
	return nil
}

func (handle *AgeHandle) query() ([]AgeInfo, error) {
	var data AgeQuery
	res, err := handle.client.R().SetResult(&data).Get("https://www.agemys.cc/collect_get?pageindex=0&pagesize=99999999")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New("Query error")
	}
	return data.Data.List, nil
}

func (handle *AgeHandle) UpdateData() ([]AgeInfo, error) {
	err := handle.login()
	if err != nil {
		return nil, err
	}
	data, err := handle.query()
	if err != nil {
		return nil, err
	}
	var updateInfo []AgeInfo
	for i := 0; i < len(data); i++ {
		val, isExist := handle.jaxleof[data[i].AID]
		if !isExist || val != data[i].UpTime {
			updateInfo = append(updateInfo, data[i])
		}
		handle.jaxleof[data[i].AID] = data[i].UpTime
	}
	return updateInfo,nil
}

type AgeHandle struct {
	account  string
	password string
	client   *resty.Client
	jaxleof  map[string]int
}

func AgeHandleConstructor(account string, password string) *AgeHandle {
	handle := new(AgeHandle)
	handle.account = account
	handle.password = password
	handle.client = resty.New()
	handle.jaxleof = make(map[string]int)
	return handle
}

type AgeQuery struct {
	Status  int
	Message string
	Data    AgeInfoList
}

type AgeInfo struct {
	AID      string
	Time     int
	UpTime   int
	Title    string
	NewTitle string
	PicSmall string
	Href     string
}

type AgeInfoList struct {
	List    []AgeInfo
	Allsize int
}

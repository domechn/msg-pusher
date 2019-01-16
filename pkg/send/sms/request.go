/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : request.go
#   Created       : 2019/1/7 16:06
#   Last Modified : 2019/1/7 16:06
#   Describe      :
#
# ====================================================*/
package sms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/satori/go.uuid"
	"net/url"
	"sort"
	"strings"
	"time"
)

// Request implements send.Message
type Request struct {
	// system parameters
	AccessKeyId      string
	Timestamp        string
	Format           string
	SignatureMethod  string
	SignatureVersion string
	SignatureNonce   string
	Signature        string

	// business parameters
	Action          string
	Version         string
	RegionId        string
	PhoneNumbers    string
	SignName        string
	TemplateCode    string
	TemplateParam   string
	SmsUpExtendCode string
	OutId           string
}

func (r *Request) Content() []byte {
	return []byte(r.TemplateParam)
}

func (r *Request) To() string {
	return r.PhoneNumbers
}

type kv struct {
	key   string
	value string
}

type kvs []kv

func (k kvs) Len() int {
	return len(k)
}

func (k kvs) Less(i, j int) bool {
	return k[i].key < k[j].key
}

func (k kvs) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

// NewRequest implements send.Message,这里没有赋值AccessKeyId，
// 而是在send的时候才赋的该值，如果需要单独使用，请记得给该值赋值
func NewRequest(phoneNumbers, signName, templateCode, templateParam, outID string) *Request {
	req := new(Request)
	now := time.Now()
	local, _ := time.LoadLocation("GMT")
	// 2006-01-02T15:04:05Z
	req.Timestamp = now.In(local).Format("2006-01-02T15:04:05Z")
	req.Format = "json"
	req.SignatureMethod = "HMAC-SHA1"
	req.SignatureVersion = "1.0"
	uid := uuid.NewV4()
	req.SignatureNonce = uid.String()

	req.Action = "SendSms"
	req.Version = "2017-05-25"
	req.RegionId = "cn-hangzhou"
	req.PhoneNumbers = phoneNumbers
	req.SignName = signName
	req.TemplateCode = templateCode
	req.TemplateParam = templateParam
	req.SmsUpExtendCode = "1234567"
	req.OutId = outID
	return req
}

func (r *Request) isValid() error {
	if len(r.AccessKeyId) == 0 {
		return fmt.Errorf("AccessKeyId required")
	}

	if len(r.PhoneNumbers) == 0 {
		return fmt.Errorf("PhoneNumbers required")
	}

	if len(r.SignName) == 0 {
		return fmt.Errorf("SignName required")
	}

	if len(r.TemplateCode) == 0 {
		return fmt.Errorf("TemplateCode required")
	}

	if len(r.TemplateParam) == 0 {
		return fmt.Errorf("TemplateParam required")
	}

	return nil
}

func (r *Request) Encode(accessKeySecret, gatewayUrl string) (string, error) {
	if err := r.isValid(); err != nil {
		return "", err
	}
	var k kvs
	k = append(k, kv{
		key:   "SignatureMethod",
		value: r.SignatureMethod,
	}, kv{
		key:   "SignatureNonce",
		value: r.SignatureNonce,
	}, kv{
		key:   "AccessKeyId",
		value: r.AccessKeyId,
	}, kv{
		key:   "SignatureVersion",
		value: r.SignatureVersion,
	}, kv{
		key:   "Timestamp",
		value: r.Timestamp,
	}, kv{
		key:   "Format",
		value: r.Format,
	}, kv{
		key:   "Action",
		value: r.Action,
	}, kv{
		key:   "Version",
		value: r.Version,
	}, kv{
		key:   "RegionId",
		value: r.RegionId,
	}, kv{
		key:   "PhoneNumbers",
		value: r.PhoneNumbers,
	}, kv{
		key:   "SignName",
		value: r.SignName,
	}, kv{
		key:   "TemplateParam",
		value: r.TemplateParam,
	}, kv{
		key:   "TemplateCode",
		value: r.TemplateCode,
	}, kv{
		key:   "SmsUpExtendCode",
		value: r.SmsUpExtendCode,
	}, kv{
		key:   "OutId",
		value: r.OutId,
	})

	sort.Sort(k)

	var temp []string
	for _, v := range k {
		temp = append(temp, specialUrlEncode(v.key)+"="+specialUrlEncode(v.value))
	}
	sortQueryString := strings.Join(temp, "&")
	stringToSign := "GET" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortQueryString)
	sign := sign(accessKeySecret+"&", stringToSign)
	signature := specialUrlEncode(sign)
	return gatewayUrl + "?Signature=" + signature + "&" + sortQueryString, nil
}

func specialUrlEncode(value string) string {
	rstValue := url.QueryEscape(value)
	rstValue = strings.Replace(rstValue, "+", "%20", -1)
	rstValue = strings.Replace(rstValue, "*", "%2A", -1)
	rstValue = strings.Replace(rstValue, "%7E", "~", -1)
	return rstValue
}

func sign(accessKeySecret, sortquerystring string) string {
	h := hmac.New(sha1.New, []byte(accessKeySecret))
	h.Write([]byte(sortquerystring))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

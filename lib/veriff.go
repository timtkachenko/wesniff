package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/errors"
	"io"
	"io/ioutil"
	"os"
	"time"
	"wesniff/infra"
	"wesniff/wesniff/operations/users"
)

const TimeFormat = "2006-01-02T15:04:05.000Z"
const HTTP_CLIENT_TIMEOUT = time.Millisecond * 991000
const baseUrl = "https://stationapi.veriff.com/v1"
const callback = "https://stationapi.veriff.com"

type Veriff struct {
}

//curl -X POST \
//  --url '/v1/sessions/' \
//  -H 'Content-Type: application/json' \
//  -H 'X-AUTH-CLIENT: Your-API-KEY' \
//  -d '{
//    "verification": {
//        "callback": "https://veriff.com",
//        "person": {
//            "firstName": "John",
//            "lastName": "Smith",
//            "idNumber": "123456789"
//        },
//        "document": {
//            "number": "B01234567",
//            "type": "PASSPORT",
//            "country": "EE"
//        },
//        "vendorData": "11111111",
//        "lang": "en",
//        "timestamp": "2016-05-19T08:30:25.597Z"
//    }
//}'
func (v *Veriff) Create(params users.CreateUserParams) (session *Verification, err error) {
	payload := &CreatePayload{
		Verification: VerifficationPost{
			Callback: callback,
			Person: Person{
				FirstName: "John",
				LastName:  "Smith",
				IdNumber:  "123456789",
			},
			Document: Document{
				Number:  "B01234567",
				Type:    "PASSPORT",
				Country: "EE",
			},
			VendorData: params.Email,
			Lang:       "en",
			Timestamp:  time.Now().Format(TimeFormat),
		},
	}
	client := infra.NewApiClient(baseUrl+"/sessions", HTTP_CLIENT_TIMEOUT)
	resp, err := client.Execute(infra.ApiClientOptions{
		Method:     "POST",
		Headers:    map[string]string{"X-AUTH-CLIENT": os.Getenv("VERIFF_PUBLIC_KEY")},
		EncodeJson: true,
	}, payload)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New(400, "cannot create")
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(400, fmt.Sprintf("cannot create: StatusCode %d", resp.StatusCode))
	}
	var result SessionResponse
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, errors.New(400, err.Error())
	}
	return &result.Verification, nil
}

//curl -X POST \
//  --url '/v1/sessions/aea9ba6d-1b47-47fc-a4fc-f72b6d3584a7/media' \
//  -H 'Content-Type: application/json' \
//  -H 'X-AUTH-CLIENT: Your-API-KEY' \
//  -H 'X-SIGNATURE: 034c6da2bb31fd9e6892516c6d7b90ebe10f79b47cfb3d155d77b4d9b66e1d53' \
//  -d '{
//    "image": {
//    "context": "document-front",
//      "content": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA+.../9fgAEAKcxisFjVfn0AAAAASUVORK5CYII=",
//      "timestamp": "2019-10-29T06:30:25.597Z"
//    }
//}'
func (v *Veriff) Upload(sesson *Verification, image io.ReadCloser) (*UploadResponse, error) {
	Content, err := ioutil.ReadAll(image)
	base64Encoding:= "data:image/png;base64,"+ base64.StdEncoding.EncodeToString(Content)
	//fmt.Printf("%s\n", base64Encoding)
	payload := &UploadPayload{
		Image: Image{
			Context:   "document-front",
			Content:   base64Encoding,
			Timestamp: time.Now().Format(TimeFormat),
		},
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := infra.NewApiClient(baseUrl+"/sessions/"+ sesson.Id+"/media", HTTP_CLIENT_TIMEOUT)
	resp, err := client.Execute(infra.ApiClientOptions{
		Method:     "POST",
		Headers:    map[string]string{"X-AUTH-CLIENT": os.Getenv("VERIFF_PUBLIC_KEY"),"X-HMAC-SIGNATURE": buildSignature([]byte(os.Getenv("VERIFF_PRIVATE_KEY")), jsonPayload)},
	}, string(jsonPayload))
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New(400, "cannot create")
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(400, fmt.Sprintf("cannot create: StatusCode %d", resp.StatusCode))
	}
	var result UploadResponse
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, errors.New(400, err.Error())
	}
	return &result, nil
}

//curl -X PATCH \
//  --url '/v1/sessions/fd5c1563-1d23-4b1a-ae46-7ba429927ed8' \
//  -H 'Content-Type: application/json' \
//  -H 'X-AUTH-CLIENT: Your-API-KEY' \
//  -H 'X-SIGNATURE: dd994f70b1150ae012f9c1d6d20adf7ed69780044835d39de20b00ffae0660a0' \
//  -d '{
//    "verification": {
//      "status": "submitted",
//      "timestamp": "2019-10-29T06:30:25.597Z"
//    }
//}'
func (v *Veriff) Update(sesson *Verification) (*UpdateResponse, error) {
	payload := &UpdatePayload{
		Verification: VerificationUpdate{
			Status:    "submitted",
			Timestamp: time.Now().Format(TimeFormat),
		},
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	client := infra.NewApiClient(baseUrl+"/sessions/"+ sesson.Id, HTTP_CLIENT_TIMEOUT)
	resp, err := client.Execute(infra.ApiClientOptions{
		Method:     "PATCH",
		Headers:    map[string]string{"X-AUTH-CLIENT": os.Getenv("VERIFF_PUBLIC_KEY"),"X-HMAC-SIGNATURE": buildSignature([]byte(os.Getenv("VERIFF_PRIVATE_KEY")), jsonPayload)},
	}, string(jsonPayload))
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New(400, "cannot create")
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(400, fmt.Sprintf("cannot create: StatusCode %d", resp.StatusCode))
	}
	var result UpdateResponse
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, errors.New(400, err.Error())
	}
	return &result, nil
}

func buildSignature(key []byte, payload []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}

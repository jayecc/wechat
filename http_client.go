package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

var DefaultHttpClient *http.Client

func init() {
	client := &http.Client{}
	client.Timeout = time.Second * 5
	DefaultHttpClient = client
}

// encodeURL 编码url
func encodeURL(baseURL string, params map[string]string) (string, error) {

	parse, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := parse.Query()
	for k, v := range params {
		query.Set(k, v)
	}

	parse.RawQuery = query.Encode()
	return parse.String(), nil
}

// struct2Map 结构体转map
func struct2Map(in interface{}, out map[string]string) error {

	elem := reflect.ValueOf(in).Elem()
	_type := elem.Type()
	for i := 0; i < _type.NumField(); i++ {
		switch _type.Field(i).Type.Kind() {
		case reflect.Int, reflect.Uint, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
			out[_type.Field(i).Tag.Get("json")] = strconv.Itoa(int(elem.Field(i).Int()))
		case reflect.String:
			out[_type.Field(i).Tag.Get("json")] = elem.Field(i).String()
		default:
			return fmt.Errorf("data %v unresolved type %v", _type.Field(i).Name, _type.Field(i).Type.Kind())
		}
	}
	return nil
}

// httpGetJSON http get request
func httpGetJSON(clt *http.Client, baseUrl string, request interface{}, response interface{}) error {

	params := make(map[string]string)
	if err := struct2Map(request, params); err != nil {
		return errors.Wrap(err, "params error")
	}

	u, err := encodeURL(baseUrl, params)
	if err != nil {
		return errors.Wrap(err, "url encode error")
	}

	httpResp, err := clt.Get(u)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return decodeJSONHttpResponse(httpResp.Body, response)
}

// httpPostJSON http post request
func httpPostJSON(clt *http.Client, u string, request interface{}, response interface{}) error {

	buffer := textBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer textBufferPool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(request); err != nil {
		return err
	}

	httpResp, err := clt.Post(u, "application/json; charset=utf-8", buffer)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return decodeJSONHttpResponse(httpResp.Body, response)
}

type MultipartFormField struct {
	IsFile   bool
	Name     string
	FileName string
	Value    io.Reader
}

// httpPostMultipartForm http post request
func httpPostMultipartForm(clt *http.Client, u string, fields []MultipartFormField, response interface{}) error {

	buffer := mediaBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer mediaBufferPool.Put(buffer)

	multipartWriter := multipart.NewWriter(buffer)
	for i := 0; i < len(fields); i++ {
		if field := &fields[i]; field.IsFile {
			partWriter, err := multipartWriter.CreateFormFile(field.Name, field.FileName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(partWriter, field.Value); err != nil {
				return err
			}
		} else {
			partWriter, err := multipartWriter.CreateFormField(field.Name)
			if err != nil {
				return err
			}
			if _, err = io.Copy(partWriter, field.Value); err != nil {
				return err
			}
		}
	}
	if err := multipartWriter.Close(); err != nil {
		return err
	}

	httpResp, err := clt.Post(u, multipartWriter.FormDataContentType(), buffer)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return decodeJSONHttpResponse(httpResp.Body, response)
}

// decodeJSONHttpResponse http json response decode
func decodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

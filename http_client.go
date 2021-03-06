package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// DefaultHTTPClient 默认http客户端
var DefaultHTTPClient *http.Client

// queryParams map
type queryParams map[string]string

func init() {
	client := &http.Client{}
	client.Timeout = time.Second * 5
	DefaultHTTPClient = client
}

// encodeURL 编码url
func encodeURL(baseURL string, params queryParams) (string, error) {

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
func httpGetJSON(clt *http.Client, URL string, request interface{}, response interface{}) error {

	params := make(map[string]string)
	if err := struct2Map(request, params); err != nil {
		return errors.Wrap(err, "params error")
	}

	u, err := encodeURL(URL, params)
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
	return decodeJSONResponse(httpResp.Body, response)
}

// httpPostJSON http post request
func httpPostJSON(clt *http.Client, URL string, request interface{}, response interface{}) error {

	buffer := textBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer textBufferPool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(request); err != nil {
		return err
	}

	httpResp, err := clt.Post(URL, "application/json; charset=utf-8", buffer)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return decodeJSONResponse(httpResp.Body, response)
}

// MultipartFormField 文件
type MultipartFormField struct {
	IsFile   bool
	Name     string
	FileName string
	Value    io.Reader
}

// httpPostMultipartForm http post request
func httpPostMultipartForm(clt *http.Client, URL string, fields []MultipartFormField, response interface{}) error {

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

	httpResp, err := clt.Post(URL, multipartWriter.FormDataContentType(), buffer)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return decodeJSONResponse(httpResp.Body, response)
}

// decodeJSONHttpResponse http json response decode
func decodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// decodeJSONResponse 响应状态码判断
func decodeJSONResponse(r io.Reader, response interface{}) error {

	var errCode Error

	if err := decodeJSONHttpResponse(r, &errCode); err != nil {
		return err
	}

	if errCode.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(errCode.Error()), "http response code error")
	}

	return decodeJSONHttpResponse(r, response)
}

package utils

import (
	"testing"
)

func TestTriggerStart(t *testing.T) {
	type args struct {
		w WEBHook
	}
	headerSlice := make([]Headers, 0)
	headerSlice = append(headerSlice, Headers{Name: "Content-Type", Val: "application/json"})
	headerSlice = append(headerSlice, Headers{Name: "cookie", Val: "uuid=d1f0ecf1-26df-41e6-9e38-1e2f83a6fb79"})
	headerSlice = append(headerSlice, Headers{Name: "Authorization", Val: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXV"})

	reqBody := make(map[string]interface{})
	reqBody["id"] = 11
	reqBody["name"] = "ling"
	reqBody["age"] = 22
	// bodyInfo, _ := json.Marshal(reqBody)

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{

		{"post", args{WEBHook{
			Type:    "before",
			Method:  "post",
			URL:     "http://httpbin.org/post",
			Headers: headerSlice,
			Body:    reqBody,
		}}, "", false},
		{"get", args{WEBHook{
			Type:    "before",
			Method:  "get",
			URL:     "http://httpbin.org/get",
			Headers: headerSlice,
			Body:    reqBody,
		}}, "", false},
		{"delete", args{WEBHook{
			Type:    "before",
			Method:  "delete",
			URL:     "http://httpbin.org/delete",
			Headers: headerSlice,
			Body:    reqBody,
		}}, "", false},
		{"put", args{WEBHook{
			Type:    "before",
			Method:  "put",
			URL:     "http://httpbin.org/put",
			Headers: headerSlice,
			Body:    reqBody,
		}}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TriggerStart(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("TriggerStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TriggerStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

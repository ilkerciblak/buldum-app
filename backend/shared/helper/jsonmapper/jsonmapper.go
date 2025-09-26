package jsonmapper

import (
	"encoding/json"
	"net/http"
)

// TODO: any yerine IRequestBody olsun belki? ama sacma aq xd ya da degil Validate method zorunlu tutariz
func DecodeRequestBody[T any](r *http.Request) (target T, err error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&target); err != nil {
		return target, err
	}
	return target, nil
}

func EncodeObjectToJson[T any](payload T) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func EncodeDecodeMapToStruct[T interface{}](m map[string]interface{}) (target T, err error) {
	marsh, err := json.Marshal(m)
	if err != nil {
		return target, err
	}

	err = json.Unmarshal(marsh, &target)
	if err != nil {
		return target, err
	}

	return target, nil
}

func DecodeRequestBodyWithTarget(r *http.Request, target any) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&target); err != nil {
		return err
	}
	return nil
}

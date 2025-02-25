package omap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// MarshalJSON implements json.Marshaler interface
// to marshall a sorted list of key-value pairs into an object.
func (om Map[K, V]) MarshalJSON() ([]byte, error) {
	if om == nil {
		return []byte(`null`), nil
	}

	if len(om) == 0 {
		return []byte(`{}`), nil
	}

	var buf bytes.Buffer
	if err := buf.WriteByte('{'); err != nil {
		return nil, err
	}

	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	for i, p := range om {
		if i != 0 {
			if err := buf.WriteByte(','); err != nil {
				return nil, err
			}
		}

		if err := enc.Encode(p.Key); err != nil {
			return nil, err
		}

		if err := buf.WriteByte(':'); err != nil {
			return nil, err
		}

		if err := enc.Encode(p.Val); err != nil {
			return nil, err
		}
	}

	if err := buf.WriteByte('}'); err != nil {
		return nil, err
	}

	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
// to unmarshal an object into a sorted list of key-value pairs.
func (om *Map[K, V]) UnmarshalJSON(data []byte) error {
	var m map[K]V
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	t, err := dec.Token()
	if err != nil {
		return err
	}

	if t != json.Delim('{') {
		return errors.New("expected start of object")
	}

	for {
		t, err := dec.Token()
		if err != nil {
			return err
		}

		if t == json.Delim('}') {
			break
		}

		key, ok := t.(K)
		if !ok {
			return fmt.Errorf("expected object key to be %T, got %T", key, t)
		}

		*om = append(*om, Pair[K, V]{
			Key: key,
			Val: m[key],
		})

		// ignored value
		if err := skipJSONValue(dec); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}
	}

	return nil
}

var errJSONEnd = errors.New("invalid end of json array or object")

func skipJSONValue(dec *json.Decoder) error {
	t, err := dec.Token()
	if err != nil {
		return err
	}

	switch t {
	case json.Delim('['), json.Delim('{'):
		for {
			if err := skipJSONValue(dec); err != nil {
				if errors.Is(err, errJSONEnd) {
					break
				}
				return err
			}
		}
	case json.Delim(']'), json.Delim('}'):
		return errJSONEnd
	}

	return nil
}

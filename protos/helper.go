package protos

import (
	"bytes"
	"fmt"
	"strings"
)

func (x ServiceType_Enum) Suit(m ServiceEncryptionMethod_Enum) bool {
	switch x {
	case ServiceType_SS_LIBEV:
		switch m {
		case ServiceEncryptionMethod_AES_256_CFB, ServiceEncryptionMethod_AES_128_GCM, ServiceEncryptionMethod_AES_192_GCM, ServiceEncryptionMethod_AES_256_GCM, ServiceEncryptionMethod_CHACHA20_IETF_POLY1305:
			return true
		}
	case ServiceType_SSR:
		switch m {
		case ServiceEncryptionMethod_AES_256_CFB, ServiceEncryptionMethod_AES_256_CTR, ServiceEncryptionMethod_CHACHA20, ServiceEncryptionMethod_CHACHA20_IETF:
			return true

		}

	}

	return false
}

func (x ServiceEncryptionMethod_Enum) ValidMethod() string {
	return strings.ReplaceAll(strings.ToLower(x.String()), "_", "-")
}

func (x *ServiceEncryptionMethod_Enum) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, "\"")
	value, ok := ServiceEncryptionMethod_Enum_value[string(data)]
	if !ok {
		return fmt.Errorf("unmarshal \"%s\" to %T failed", data, x)
	}

	*x = ServiceEncryptionMethod_Enum(value)
	return nil
}

func (x ServiceEncryptionMethod_Enum) MarshalJSON() ([]byte, error) {
	var buffer []byte
	buffer = append(buffer, '"')
	buffer = append(buffer, x.ValidMethod()...)
	buffer = append(buffer, '"')
	return buffer, nil
}

func (x *ServiceType_Enum) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, "\"")
	value, ok := ServiceType_Enum_value[string(data)]
	if !ok {
		return fmt.Errorf("unmarshal \"%s\" to %T failed", data, x)
	}

	*x = ServiceType_Enum(value)
	return nil
}

func (x ServiceType_Enum) MarshalJSON() ([]byte, error) {
	var buffer []byte
	buffer = append(buffer, '"')
	buffer = append(buffer, x.String()...)
	buffer = append(buffer, '"')
	return buffer, nil
}

func (x ServiceType_Enum) ImageName() string {
	if x == ServiceType_NOT_SET {
		return ""
	}

	return fmt.Sprintf("goignite/%s:latest", strings.ReplaceAll(strings.ToLower(x.String()), "_", "-"))
}

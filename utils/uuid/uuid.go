package uuid

import uuid "github.com/satori/go.uuid"

func Generate () string {
	u, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return u.String()
}
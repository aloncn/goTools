package guuid

import uuid "github.com/satori/go.uuid"

func Gen () string {
	u, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return u.String()
}
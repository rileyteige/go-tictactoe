package server

import (
	"github.com/nu7hatch/gouuid"
	"log"
)

type Guid uuid.UUID

func emptyGuid() Guid {
	const empty = "00000000-0000-0000-0000-000000000000"
	id, _ := parseGuid(empty)
	return id
}

func generateGuid() Guid {
	id, err := uuid.NewV4()
	if err == nil {
		log.Fatal(err)
	}

	return Guid(*id)
}

func parseGuid(s string) (Guid, error) {
	id, err := uuid.ParseHex(s)
	if err != nil {
		return generateGuid(), err
	}

	return Guid(*id), nil
}

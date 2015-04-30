package server

import (
	"github.com/nu7hatch/gouuid"
	"log"
)

type Guid string

const EmptyGuid = Guid("00000000-0000-0000-0000-000000000000")

func generateGuid() Guid {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	return Guid(id.String())
}

func parseGuid(s string) (Guid, error) {
	id, err := uuid.ParseHex(s)
	if err != nil {
		return generateGuid(), err
	}

	return Guid(id.String()), nil
}

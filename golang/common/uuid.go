package common

import "github.com/google/uuid"

func GetNewDocId () string {
	docId := uuid.New()
	return docId.String()
}


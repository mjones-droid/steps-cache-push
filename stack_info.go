package main

import (
	"encoding/json"
	"fmt"
)

func stackVersionData(stackID string) ([]byte, error) {
	type archiveInfo struct {
		StackID string `json:"stack_id,omitempty"`
	}
	log.Donef("Marshaling Stack Version Data JSON")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	stackData, err := json.Marshal(archiveInfo{
		StackID: stackID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data, error: %s", err)
	}
	return stackData, nil
}

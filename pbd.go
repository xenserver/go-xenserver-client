package client

import (
	"github.com/nilshell/xmlrpc"
)

type PBD XenAPIObject

func (self *PBD) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "PBD.get_record", self.Ref)
	if err != nil {
		return record, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return
}

func (self *PBD) GetUUID() (uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PBD.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	uuid = result.Value.(string)
	return
}

func (self *PBD) GetSR() (sr *SR, err error) {
	sr = new(SR)
	result := APIResult{}
	err = self.Client.APICall(&result, "PBD.get_SR", self.Ref)
	if err != nil {
		return nil, err
	}
	sr.Ref = result.Value.(string)
	sr.Client = self.Client
	return
}

package client

import (
	"github.com/nilshell/xmlrpc"
)

type PIF XenAPIObject

func (self *PIF) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_record", self.Ref)
	if err != nil {
		return record, err
	}
	if result.Value != nil {
		for k, v := range result.Value.(xmlrpc.Struct) {
			record[k] = v
		}
	}
	return
}

func (self *PIF) GetUUID() (uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		uuid = result.Value.(string)
	}
	return
}

func (self *PIF) GetIP() (ip string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_IP", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		ip = result.Value.(string)
	}

	return
}

func (self *PIF) GetMAC() (mac string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_MAC", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		mac = result.Value.(string)
	}

	return
}

func (self *PIF) GetIsAttached() (value bool, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_currently_attached", self.Ref)
	if err != nil {
		return false, err
	}
	if result.Value != nil {
		value = result.Value.(bool)
	}
	return
}

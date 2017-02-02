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
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return record, nil
}

func (self *PIF) GetIP() (ip string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_IP", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value == nil {
		return "", nil
	}
	ip = result.Value.(string)
	return ip, nil
}

func (self *PIF) GetMAC() (mac string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_MAC", self.Ref)
	if err != nil {
		return "", err
	}
	mac = result.Value.(string)
	return mac, nil
}

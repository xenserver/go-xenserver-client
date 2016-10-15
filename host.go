package client

import (
	"github.com/nilshell/xmlrpc"
)

type Host XenAPIObject

func (self *Host) CallPlugin(plugin, method string, params map[string]string) (response string, err error) {
	result := APIResult{}
	params_rec := make(xmlrpc.Struct)
	for key, value := range params {
		params_rec[key] = value
	}
	err = self.Client.APICall(&result, "host.call_plugin", self.Ref, plugin, method, params_rec)
	if err != nil {
		return "", err
	}
	response = result.Value.(string)
	return
}

func (self *Host) GetAddress() (address string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_address", self.Ref)
	if err != nil {
		return "", err
	}
	address = result.Value.(string)
	return address, nil
}

func (self *Host) GetHostname() (hostname string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_hostname", self.Ref)
	if err != nil {
		return "", err
	}
	hostname = result.Value.(string)
	return hostname, nil
}

func (self *Host) GetNameLabel() (nameLabel string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_name_label", self.Ref)
	if err != nil {
		return "", err
	}
	nameLabel = result.Value.(string)
	return nameLabel, nil
}

func (self *Host) GetSoftwareVersion() (versions map[string]interface{}, err error) {
	versions = make(map[string]interface{})

	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_software_version", self.Ref)
	if err != nil {
		return nil, err
	}

	for k, v := range result.Value.(xmlrpc.Struct) {
		versions[k] = v.(string)
	}
	return
}

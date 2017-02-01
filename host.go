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

func (self *Host) GetCPUInfo() (cpus map[string]interface{}, err error) {
	cpus = make(map[string]interface{})

	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_cpu_info", self.Ref)
	if err != nil {
		return nil, err
	}

	for k, v := range result.Value.(xmlrpc.Struct) {
		cpus[k] = v.(string)
	}
	return
}

func (self *Host) GetCPUs() (hcpus []*Host_CPU, err error) {
	hcpus = make([]*Host_CPU, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_host_CPUs", self.Ref)
	if err != nil {
		return hcpus, err
	}

	for _, elem := range result.Value.([]interface{}) {
		vm := new(Host_CPU)
		vm.Ref = elem.(string)
		vm.Client = self.Client
		hcpus = append(hcpus, vm)
	}

	return hcpus, nil
}
//todo: get utilisation, get_host_CPUs
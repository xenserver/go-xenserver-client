package client

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/nilshell/xmlrpc"
)

type XenAPIClient struct {
	Session  interface{}
	Host     string
	Url      string
	Username string
	Password string
	RPC      *xmlrpc.Client
}

type APIResult struct {
	Status           string
	Value            interface{}
	ErrorDescription string
}

type XenAPIObject struct {
	Ref    string
	Client *XenAPIClient
}

func (c *XenAPIClient) RPCCall(result interface{}, method string, params []interface{}) (err error) {
	log.Debugf("RPCCall method=%v params=%v\n", method, params)
	p := new(xmlrpc.Params)
	p.Params = params
	err = c.RPC.Call(method, *p, result)
	return err
}

func (client *XenAPIClient) Login() (err error) {
	//Do loging call
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)
	params[0] = client.Username
	params[1] = client.Password

	err = client.RPCCall(&result, "session.login_with_password", params)
	if err == nil {
		// err might not be set properly, so check the reference
		if result["Value"] == nil {
			return errors.New("Invalid credentials supplied")
		}
	}
	client.Session = result["Value"]
	return err
}

func (client *XenAPIClient) NewVIF() (new_vif *VIF) {
	new_vif = &VIF{
		Client: client,
	}
	return
}

func (client *XenAPIClient) NewNetwork() (new_network *Network) {
	new_network = &Network{
		Client: client,
	}
	return
}

func (client *XenAPIClient) NewSR() (new_sr *SR) {
	new_sr = &SR{
		Client: client,
	}
	return
}

func (client *XenAPIClient) NewVBD() (new_vbd *VBD) {
	new_vbd = &VBD{
		Client: client,
	}
	return
}

func (client *XenAPIClient) NewVDI() (new_vdi *VDI) {
	new_vdi = &VDI{
		Client: client,
	}
	return
}

func (client *XenAPIClient) APICall(result *APIResult, method string, params ...interface{}) (err error) {
	if client.Session == nil {
		log.Errorf("no session\n")
		return fmt.Errorf("No session. Unable to make call")
	}

	//Make a params slice which will include the session
	p := make([]interface{}, len(params)+1)
	p[0] = client.Session

	if params != nil {
		for idx, element := range params {
			p[idx+1] = element
		}
	}

	res := xmlrpc.Struct{}

	err = client.RPCCall(&res, method, p)

	if err != nil {
		return err
	}

	result.Status = res["Status"].(string)

	if result.Status != "Success" {
		log.Errorf("Encountered an API error: %v %v", result.Status, res["ErrorDescription"])
		return fmt.Errorf("API Error: %s", res["ErrorDescription"])
	} else {
		result.Value = res["Value"]
	}
	return
}

func (client *XenAPIClient) GetHosts() (hosts []*Host, err error) {
	hosts = make([]*Host, 0)
	result := APIResult{}
	err = client.APICall(&result, "host.get_all")
	if err != nil {
		return hosts, err
	}
	for _, elem := range result.Value.([]interface{}) {
		host := new(Host)
		host.Ref = elem.(string)
		host.Client = client
		hosts = append(hosts, host)
	}
	return hosts, nil
}

func (client *XenAPIClient) GetPools() (pools []*Pool, err error) {
	pools = make([]*Pool, 0)
	result := APIResult{}
	err = client.APICall(&result, "pool.get_all")
	if err != nil {
		return pools, err
	}

	for _, elem := range result.Value.([]interface{}) {
		pool := new(Pool)
		pool.Ref = elem.(string)
		pool.Client = client
		pools = append(pools, pool)
	}

	return pools, nil
}

func (client *XenAPIClient) GetDefaultSR() (sr *SR, err error) {
	pools, err := client.GetPools()

	if err != nil {
		return nil, err
	}

	pool_rec, err := pools[0].GetRecord()

	if err != nil {
		return nil, err
	}

	if pool_rec["default_SR"] == "" {
		return nil, errors.New("No default_SR specified for the pool.")
	}

	sr = new(SR)
	sr.Ref = pool_rec["default_SR"].(string)
	sr.Client = client

	return sr, nil
}

func (client *XenAPIClient) GetVMByUuid(vm_uuid string) (vm *VM, err error) {
	vm = new(VM)
	result := APIResult{}
	err = client.APICall(&result, "VM.get_by_uuid", vm_uuid)
	if err != nil {
		return nil, err
	}
	vm.Ref = result.Value.(string)
	vm.Client = client
	return
}

func (client *XenAPIClient) GetHostByUuid(host_uuid string) (host *Host, err error) {
	host = new(Host)
	result := APIResult{}
	err = client.APICall(&result, "host.get_by_uuid", host_uuid)
	if err != nil {
		return nil, err
	}
	host.Ref = result.Value.(string)
	host.Client = client
	return
}

func (client *XenAPIClient) GetVMByNameLabel(name_label string) (vms []*VM, err error) {
	vms = make([]*VM, 0)
	result := APIResult{}
	err = client.APICall(&result, "VM.get_by_name_label", name_label)
	if err != nil {
		return vms, err
	}

	for _, elem := range result.Value.([]interface{}) {
		vm := new(VM)
		vm.Ref = elem.(string)
		vm.Client = client
		vms = append(vms, vm)
	}

	return vms, nil
}

func (client *XenAPIClient) GetVMAll() (vms []*VM, err error) {
	vms = make([]*VM, 0)
	result := APIResult{}
	err = client.APICall(&result, "VM.get_all")
	if err != nil {
		return vms, err
	}

	for _, elem := range result.Value.([]interface{}) {
		vm := new(VM)
		vm.Ref = elem.(string)
		vm.Client = client
		vms = append(vms, vm)
	}

	return vms, nil
}

func (client *XenAPIClient) GetHostByNameLabel(name_label string) (hosts []*Host, err error) {
	hosts = make([]*Host, 0)
	result := APIResult{}
	err = client.APICall(&result, "host.get_by_name_label", name_label)
	if err != nil {
		return hosts, err
	}

	for _, elem := range result.Value.([]interface{}) {
		host := new(Host)
		host.Ref = elem.(string)
		host.Client = client
		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (client *XenAPIClient) GetSRByNameLabel(name_label string) (srs []*SR, err error) {
	srs = make([]*SR, 0)
	result := APIResult{}
	err = client.APICall(&result, "SR.get_by_name_label", name_label)
	if err != nil {
		return srs, err
	}

	for _, elem := range result.Value.([]interface{}) {
		sr := new(SR)
		sr.Ref = elem.(string)
		sr.Client = client
		srs = append(srs, sr)
	}

	return srs, nil
}

func (client *XenAPIClient) GetNetworks() (networks []*Network, err error) {
	networks = make([]*Network, 0)
	result := APIResult{}
	err = client.APICall(&result, "network.get_all")
	if err != nil {
		return nil, err
	}

	for _, elem := range result.Value.([]interface{}) {
		network := new(Network)
		network.Ref = elem.(string)
		network.Client = client
		networks = append(networks, network)
	}

	return networks, nil
}

func (client *XenAPIClient) GetNetworkByUuid(network_uuid string) (network *Network, err error) {
	network = new(Network)
	result := APIResult{}
	err = client.APICall(&result, "network.get_by_uuid", network_uuid)
	if err != nil {
		return nil, err
	}
	network.Ref = result.Value.(string)
	network.Client = client
	return
}

func (client *XenAPIClient) GetVIFByUuid(vfi_uuid string) (vif *VIF, err error) {
	vif = client.NewVIF()
	result := APIResult{}
	err = client.APICall(&result, "VIF.get_by_uuid", vfi_uuid)
	if err != nil {
		return nil, err
	}
	vif.Ref = result.Value.(string)
	return
}

func (client *XenAPIClient) GetNetworkByNameLabel(name_label string) (networks []*Network, err error) {
	networks = make([]*Network, 0)
	result := APIResult{}
	err = client.APICall(&result, "network.get_by_name_label", name_label)
	if err != nil {
		return networks, err
	}

	for _, elem := range result.Value.([]interface{}) {
		network := new(Network)
		network.Ref = elem.(string)
		network.Client = client
		networks = append(networks, network)
	}

	return networks, nil
}

func (client *XenAPIClient) GetVdiByNameLabel(name_label string) (vdis []*VDI, err error) {
	vdis = make([]*VDI, 0)
	result := APIResult{}
	err = client.APICall(&result, "VDI.get_by_name_label", name_label)
	if err != nil {
		return vdis, err
	}

	for _, elem := range result.Value.([]interface{}) {
		vdi := new(VDI)
		vdi.Ref = elem.(string)
		vdi.Client = client
		vdis = append(vdis, vdi)
	}

	return vdis, nil
}

func (client *XenAPIClient) GetSRByUuid(sr_uuid string) (sr *SR, err error) {
	sr = new(SR)
	result := APIResult{}
	err = client.APICall(&result, "SR.get_by_uuid", sr_uuid)
	if err != nil {
		return nil, err
	}
	sr.Ref = result.Value.(string)
	sr.Client = client
	return
}

func (client *XenAPIClient) GetVdiByUuid(vdi_uuid string) (vdi *VDI, err error) {
	vdi = new(VDI)
	result := APIResult{}
	err = client.APICall(&result, "VDI.get_by_uuid", vdi_uuid)
	if err != nil {
		return nil, err
	}
	vdi.Ref = result.Value.(string)
	vdi.Client = client
	return
}

func (client *XenAPIClient) GetPIFs() (pifs []*PIF, err error) {
	pifs = make([]*PIF, 0)
	result := APIResult{}
	err = client.APICall(&result, "PIF.get_all")
	if err != nil {
		return pifs, err
	}
	for _, elem := range result.Value.([]interface{}) {
		pif := new(PIF)
		pif.Ref = elem.(string)
		pif.Client = client
		pifs = append(pifs, pif)
	}

	return pifs, nil
}

func (client *XenAPIClient) GetVIFs() (vifs []*VIF, err error) {
	vifs = make([]*VIF, 0)
	result := APIResult{}
	err = client.APICall(&result, "VIF.get_all")
	if err != nil {
		return nil, err
	}
	for _, elem := range result.Value.([]interface{}) {
		vif := new(VIF)
		vif.Ref = elem.(string)
		vif.Client = client
		vifs = append(vifs, vif)
	}

	return vifs, nil
}

func (client *XenAPIClient) CreateTask() (task *Task, err error) {
	result := APIResult{}
	err = client.APICall(&result, "task.create", "packer-task", "Packer task")

	if err != nil {
		return
	}

	task = new(Task)
	task.Ref = result.Value.(string)
	task.Client = client
	return
}

func (client *XenAPIClient) CreateNetwork(name_label string, name_description string, bridge string) (network *Network, err error) {
	network = new(Network)

	net_rec := make(xmlrpc.Struct)
	net_rec["name_label"] = name_label
	net_rec["name_description"] = name_description
	net_rec["bridge"] = bridge
	net_rec["other_config"] = make(xmlrpc.Struct)

	result := APIResult{}
	err = client.APICall(&result, "network.create", net_rec)
	if err != nil {
		return nil, err
	}
	network.Ref = result.Value.(string)
	network.Client = client

	return network, nil
}

func NewXenAPIClient(host, username, password string) (client XenAPIClient) {
	client.Host = host
	client.Url = fmt.Sprintf("http://%s", host)
	client.Username = username
	client.Password = password
	client.RPC, _ = xmlrpc.NewClient(client.Url, nil)
	return
}

func (client *XenAPIClient) Close() error {
	return client.RPC.Close()
}

func (client *XenAPIClient) CreateVbd(vm_ref, vdi_ref, vbdType, mode string, bootable bool) (*VBD, error) {
	vbd := new(VBD)
	oc := make(xmlrpc.Struct)
	vbd_rec := make(xmlrpc.Struct)
	vbd_rec["other_config"] = oc
	vbd_rec["status_code"] = "0"
	vbd_rec["VM"] = vm_ref
	vbd_rec["unplugabble"] = "0"
	vbd_rec["VDI"] = vdi_ref //empty for CDs
	vbd_rec["qos_algorithm_type"] = ""
	vbd_rec["bootable"] = bootable
	vbd_rec["storage_lock"] = "0"
	vbd_rec["currently_attached"] = "0"
	vbd_rec["mode"] = mode

	vm := new(VM)
	vm.Client = client
	vm.Ref = vm_ref
	vbds, err := vm.GetVBDs()
	if err != nil {
		return nil ,err
	}

	vbd_rec["userdevice"] = fmt.Sprintf("%d", len(vbds))
	vbd_rec["qos_algorithm_params"] = oc
	vbd_rec["type"] = vbdType //CD or Disk
	if vbdType == "CD" {
		vbd_rec["empty"] = true
	}else if vbdType == "Disk"{
		vbd_rec["empty"] = false
	} else {
		vbd_rec["empty"] = true
	}
	vbd.Client = client

	result := APIResult{}
	err = client.APICall(&result, "VBD.create", vbd_rec)
	if err != nil {
		return nil, fmt.Errorf("apicall error %+v", err)
	}
	vbd.Ref = result.Value.(string)
	return vbd, nil
}

func (client *XenAPIClient) GetVBDByUuid(vbd_uuid string) (vbd *VBD, err error) {
	vbd = new(VBD)
	result := APIResult{}
	err = client.APICall(&result, "VBD.get_by_uuid", vbd_uuid)
	if err != nil {
		return nil, err
	}
	vbd.Ref = result.Value.(string)
	vbd.Client = client
	return
}

func (client *XenAPIClient) CreateVM(config VMConfig) (new_instance *VM, err error) {
	templates, err := client.GetVMByNameLabel(config.GuestOS)
	if err != nil || len(templates) == 0 {
		return nil, fmt.Errorf(`no template exist for guestOS "%s". %+v`, config.GuestOS, err)
	}

	clone := new(VM)
	clone.Client = client
	clone.Ref = templates[0].Ref

	if new_instance, err = clone.Clone(config.Name_label); err != nil {
		return nil, err
	}

	if config.Other_config != nil {
		if err = new_instance.SetOtherConfig(config.Other_config); err != nil {
			return nil, err
		}
	}

	if err = new_instance.Provision(); err != nil {
		return nil, err
	}


	return new_instance, nil
}

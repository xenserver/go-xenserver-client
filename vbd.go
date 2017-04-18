package client

import (
	"github.com/nilshell/xmlrpc"
)

type VBD XenAPIObject

func (self *VBD) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.get_record", self.Ref)
	if err != nil {
		return record, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return record, nil
}

func (self *VBD) GetVDI() (vdi *VDI, err error) {
	vbd_rec, err := self.GetRecord()
	if err != nil {
		return nil, err
	}

	vdi = new(VDI)
	vdi.Ref = vbd_rec["VDI"].(string)
	vdi.Client = self.Client

	return vdi, nil
}

func (self *VBD) Eject() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "Async.VBD.eject", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VBD) Unplug() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "Async.VBD.unplug", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VBD) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.destroy", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VBD) Plug() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "Async.VBD.plug", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VBD) GetUuid() (vbd_uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	vbd_uuid = result.Value.(string)
	return vbd_uuid, nil
}

func (self *VBD) Insert(vdiRef string) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "Async.VBD.insert", self.Ref, vdiRef)
	if err != nil {
		return err
	}
	return nil
}

func (self *VBD) GetType() (vbd_type string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.get_type", self.Ref)
	if err != nil {
		return "", err
	}
	vbd_type = result.Value.(string)
	return vbd_type, nil
}

func (self *VBD) GetVM() (vm *VM, err error) {
	vbd_rec, err := self.GetRecord()
	if err != nil {
		return nil, err
	}

	vm = new(VM)
	vm.Ref = vbd_rec["VM"].(string)
	vm.Client = self.Client

	return vm, nil
}

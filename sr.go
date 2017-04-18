package client

import (
	"fmt"
	"github.com/nilshell/xmlrpc"
)

type SR XenAPIObject

func (self *SR) GetUuid() (uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	uuid = result.Value.(string)
	return
}

func (self *SR) CreateVdi(name_label, sr_ref string, size int64) (vdi_uuid string, err error) {
	vdi := new(VDI)

	vdi_rec := make(xmlrpc.Struct)
	vdi_rec["name_label"] = name_label
	vdi_rec["name_description"] = "VirtualHardDisk"
	vdi_rec["SR"] = sr_ref
	vdi_rec["virtual_size"] = fmt.Sprintf("%d", size)
	vdi_rec["type"] = "user"
	vdi_rec["sharable"] = false
	vdi_rec["read_only"] = false

	oc := make(xmlrpc.Struct)
	oc["temp"] = "temp"
	vdi_rec["other_config"] = oc
	vdi_rec["xenstore_data"] = oc

	vdi.Client = self.Client

	result := APIResult{}

	err = self.Client.APICall(&result, "VDI.create", vdi_rec)

	if err != nil {
		return "", err
	}
	vdi.Ref = result.Value.(string)
	vdi_uuid, err = vdi.GetUuid()

	return
}

func (self *SR) GetNameLabel() (name_label string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_name_label", self.Ref)
	if err != nil {
		return "", err
	}
	name_label = result.Value.(string)
	return
}

func (self *SR) GetNameDescription() (description string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_name_description", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		description = result.Value.(string)
	}
	return
}

func (self *SR) GetPhysicalSize() (size string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_physical_size", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		size = result.Value.(string)
	}

	return
}

func (self *SR) GetPhysicalUtilisation() (utilisation string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_physical_utilisation", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		utilisation = result.Value.(string)
	}
	return
}

func (self *SR) GetVirtualAllocation() (vallocation string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_virtual_allocation", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		vallocation = result.Value.(string)
	}
	return
}

func (self *SR) GetType() (sr_type string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_type", self.Ref)
	if err != nil {
		return "", err
	}
	if result.Value != nil {
		sr_type = result.Value.(string)
	}
	return
}

func (self *SR) GetVDIs() (vdis []*VDI, err error) {
	vdis = make([]*VDI, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_VDIs", self.Ref)
	if err != nil {
		return nil, err
	}
	for _, elem := range result.Value.([]interface{}) {
		vdi := new(VDI)
		vdi.Ref = elem.(string)
		vdi.Client = self.Client
		vdis = append(vdis, vdi)
	}

	return
}

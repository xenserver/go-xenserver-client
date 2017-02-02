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
	return uuid, nil
}

func (self *SR) CreateVdi(name_label, sr_ref string, size int64) (vdi_uuid string, err error) {
	vdi := new(VDI)

	vdi_rec := make(xmlrpc.Struct)
	vdi_rec["name_label"] = name_label
	vdi_rec["name_description"] = name_label
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

	return "", nil
}

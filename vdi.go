package client

type VDI XenAPIObject

type VDIType int

const (
	_ VDIType = iota
	Disk
	CD
	Floppy
)

func (self *VDI) GetUuid() (vdi_uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	vdi_uuid = result.Value.(string)
	return vdi_uuid, nil
}

func (self *VDI) GetVBDs() (vbds []VBD, err error) {
	vbds = make([]VBD, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.get_VBDs", self.Ref)
	if err != nil {
		return vbds, err
	}
	for _, elem := range result.Value.([]interface{}) {
		vbd := VBD{}
		vbd.Ref = elem.(string)
		vbd.Client = self.Client
		vbds = append(vbds, vbd)
	}

	return vbds, nil
}

func (self *VDI) GetVirtualSize() (virtual_size string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.get_virtual_size", self.Ref)
	if err != nil {
		return "", err
	}
	virtual_size = result.Value.(string)  
	return virtual_size, nil
}

func (self *VDI) Forget() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.forget", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VDI) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.destroy", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VDI) SetNameLabel(name_label string) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.set_name_label", self.Ref, name_label)
	if err != nil {
		return err
	}
	return
}

func (self *VDI) SetReadOnly(value bool) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.set_read_only", self.Ref, value)
	if err != nil {
		return err
	}
	return
}

func (self *VDI) SetSharable(value bool) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.set_sharable", self.Ref, value)
	if err != nil {
		return err
	}
	return
}

package client

type VIF XenAPIObject

func (self *VIF) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.destroy", self.Ref)
	return err
}

func (self *VIF) GetNetwork() (network *Network, err error) {

	network = new(Network)
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.get_network", self.Ref)

	if err != nil {
		return nil, err
	}
	network.Ref = result.Value.(string)
	network.Client = self.Client
	return
}

func (self *VIF) GetUuid() (vif_uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	vif_uuid = result.Value.(string)
	return
}

func (self *VIF) GetMAC() (vif_mac string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.get_MAC", self.Ref)
	if err != nil {
		return "", err
	}
	vif_mac = result.Value.(string)
	return
}

func (self *VIF) Plug() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.plug", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VIF) Unplug() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.unplug", self.Ref)
	return err
}

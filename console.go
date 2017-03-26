package client

type Console XenAPIObject

//get protocol

func (self *Console) GetProtocol() (protocol string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_protocol", self.Ref)
	if err != nil {
		return "", err
	}
	protocol = result.Value.(string)
	return protocol, nil
}

func (self *Console) GetLocation() (location string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_location", self.Ref)
	if err != nil {
		return "", err
	}
	location = result.Value.(string)
	return location, nil
}
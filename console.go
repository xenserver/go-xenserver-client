package client

type Console XenAPIObject

func (self *Console) GetLocation() (location string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_location", self.Ref)
	if err != nil {
		return "", err
	}
	location = result.Value.(string)

	return location, nil
}

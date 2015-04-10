package client

type VIF XenAPIObject

func (self *VIF) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.destroy", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

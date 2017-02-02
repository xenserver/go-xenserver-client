package client

type Host_CPU XenAPIObject

func (self *Host_CPU) GetUtilisation() (utilisation float64, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host_cpu.get_utilisation", self.Ref)
	if err != nil {
		return 0, err
	}
	utilisation = result.Value.(float64)
	return utilisation, nil
}
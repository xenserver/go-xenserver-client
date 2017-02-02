package client

type HostMetrics XenAPIObject

func (self *HostMetrics) GetTotalMemory() (count string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host_metrics.get_memory_total", self.Ref)
	if err != nil {
		return "", err
	}
	count = result.Value.(string)
	return count, nil
}

func (self *HostMetrics) GetFreeMemory() (count string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host_metrics.get_memory_free", self.Ref)
	if err != nil {
		return "", err
	}
	count = result.Value.(string)
	return count, nil
}

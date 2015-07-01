package client

import (
        "github.com/nilshell/xmlrpc"
)

type Console XenAPIObject

func (self *Console) GetRecord() (record map[string]interface{}, err error) {
        record = make(map[string]interface{})
        result := APIResult{}
        err = self.Client.APICall(&result, "console.get_record", self.Ref)
        if err != nil {
                return record, err
        }
        for k, v := range result.Value.(xmlrpc.Struct) {
                record[k] = v
        }
        return record, nil
}


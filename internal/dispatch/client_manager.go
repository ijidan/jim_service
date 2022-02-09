package dispatch

import (
	"fmt"
	"github.com/spf13/cast"
	"sync"
)

var ClientIdGatewayIdMap = sync.Map{}
var GatewayIdSendMessageMap = sync.Map{}
var GatewayIdPushToAllMap = sync.Map{}

func Dump() string  {
	str1:=DumpClientIdGatewayIdMap()
	str2:=DumpGatewayIdSendMessageMap()
	str3:=DumpGatewayIdPushToAllMap()
	str:=fmt.Sprintf("%s\n%s\n%s",str1,str2,str3)
	return str
}
func DumpClientIdGatewayIdMap() string {
	var clientIdGatewayIdMapList []string
	ClientIdGatewayIdMap.Range(func(key, value interface{}) bool {
		item:=fmt.Sprintf("[%s:%s]",key,value.(string))
		clientIdGatewayIdMapList=append(clientIdGatewayIdMapList,item)
		return true
	})
	clientIdGatewayIdMapStr:=cast.ToString(clientIdGatewayIdMapList)
	return clientIdGatewayIdMapStr
}

func DumpGatewayIdSendMessageMap() string {
	var GatewayIdSendMessageMapList []string
	GatewayIdSendMessageMap.Range(func(key, value interface{}) bool {
		item:=fmt.Sprintf("[%s:%s]",key,"111")
		GatewayIdSendMessageMapList=append(GatewayIdSendMessageMapList,item)
		return true
	})
	gatewayIdSendMessageMapStr:=cast.ToString(GatewayIdSendMessageMapList)
	return gatewayIdSendMessageMapStr
}

func DumpGatewayIdPushToAllMap() string {
	var GatewayIdPushToAllMapList []string
	GatewayIdPushToAllMap.Range(func(key, value interface{}) bool {
		item:=fmt.Sprintf("[%s:%s]",key,"222")
		GatewayIdPushToAllMapList=append(GatewayIdPushToAllMapList,item)
		return true
	})
	gatewayIdPushToAllMapStr:=cast.ToString(GatewayIdPushToAllMapList)
	return gatewayIdPushToAllMapStr
}



package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"jim_service/global"
	"jim_service/pkg"
	"testing"
)

func TestGetServerList(t *testing.T) {
	defer global.Close()
	serviceDiscovery := pkg.NewServiceDiscovery(global.ClientV3, global.Config.App.Name)
	serverList := serviceDiscovery.GetServerList("ping_service")
	if len(serverList)==0{
		assert.Nil(t, errors.New("server list is empty"))
	}
	t.Log(serverList)
}

func BenchmarkGetServerList(b *testing.B) {
	defer global.Close()
	for i := 0; i < b.N; i++ {
		serviceDiscovery := pkg.NewServiceDiscovery(global.ClientV3, global.Config.App.Name)
		_ = serviceDiscovery.GetServerList("ping_service")
	}
}

package naming_grpc

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jun3372/nacos-sdk-go/clients/naming_client/naming_proxy"
	"github.com/jun3372/nacos-sdk-go/util"
)

func TestRedoSubscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProxy := naming_proxy.NewMockINamingProxy(ctrl)
	evListener := NewConnectionEventListener(mockProxy)

	cases := []struct {
		serviceName string
		groupName   string
		clusters    string
	}{
		{"service-a", "group-a", ""},
		{"service-b", "group-b", "cluster-b"},
	}

	for _, v := range cases {
		fullServiceName := util.GetGroupName(v.serviceName, v.groupName)
		evListener.CacheSubscriberForRedo(fullServiceName, v.clusters)
		mockProxy.EXPECT().Subscribe(v.serviceName, v.groupName, v.clusters)
		evListener.redoSubscribe()
		evListener.RemoveSubscriberForRedo(fullServiceName, v.clusters)
	}
}

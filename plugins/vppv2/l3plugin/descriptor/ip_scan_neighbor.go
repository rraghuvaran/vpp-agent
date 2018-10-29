//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package descriptor

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ligato/cn-infra/logging"
	scheduler "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/l3plugin/descriptor/adapter"
	"github.com/ligato/vpp-agent/plugins/vppv2/l3plugin/vppcalls"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/l3"
)

const (
	// IPScanNeighbor is the name of the descriptor.
	IPScanNeighborDescriptorName = "vpp-ip-scan-neighbor"
)

var defaultIPScanNeighbor = &l3.IPScanNeighbor{
	Mode: l3.IPScanNeighbor_DISABLED,
}

// IPScanNeighborDescriptor teaches KVScheduler how to configure VPP proxy ARPs.
type IPScanNeighborDescriptor struct {
	log       logging.Logger
	ipNeigh   vppcalls.IPNeighVppAPI
	scheduler scheduler.KVScheduler
}

// NewIPScanNeighborDescriptor creates a new instance of the IPScanNeighborDescriptor.
func NewIPScanNeighborDescriptor(scheduler scheduler.KVScheduler,
	proxyArpHandler vppcalls.IPNeighVppAPI, log logging.PluginLogger) *IPScanNeighborDescriptor {

	return &IPScanNeighborDescriptor{
		scheduler: scheduler,
		ipNeigh:   proxyArpHandler,
		log:       log,
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (d *IPScanNeighborDescriptor) GetDescriptor() *adapter.IPScanNeighborDescriptor {
	return &adapter.IPScanNeighborDescriptor{
		Name: IPScanNeighborDescriptorName,
		KeySelector: func(key string) bool {
			return key == l3.IPScanNeighborKey
		},
		ValueTypeName:      proto.MessageName(&l3.IPScanNeighbor{}),
		ValueComparator:    d.EquivalentIPScanNeighbors,
		NBKeyPrefix:        l3.IPScanNeighborKey,
		Add:                d.Add,
		Modify:             d.Modify,
		Delete:             d.Delete,
		IsRetriableFailure: d.IsRetriableFailure,
		Dump:               d.Dump,
	}
}

func (d *IPScanNeighborDescriptor) EquivalentIPScanNeighbors(key string, oldValue, newValue *l3.IPScanNeighbor) bool {
	return proto.Equal(oldValue, newValue)
}

func (d *IPScanNeighborDescriptor) Add(key string, value *l3.IPScanNeighbor) (metadata interface{}, err error) {
	return d.Modify(key, defaultIPScanNeighbor, value, nil)
}

func (d *IPScanNeighborDescriptor) Delete(key string, value *l3.IPScanNeighbor, metadata interface{}) error {
	_, err := d.Modify(key, value, defaultIPScanNeighbor, metadata)
	return err
}

func (d *IPScanNeighborDescriptor) Modify(key string, oldValue, newValue *l3.IPScanNeighbor, oldMetadata interface{}) (newMetadata interface{}, err error) {
	if err := d.ipNeigh.SetIPScanNeighbor(newValue); err != nil {
		return nil, err
	}
	return nil, nil
}

func (d *IPScanNeighborDescriptor) IsRetriableFailure(err error) bool {
	return false
}

func (d *IPScanNeighborDescriptor) Dump(correlate []adapter.IPScanNeighborKVWithMetadata) (dumps []adapter.IPScanNeighborKVWithMetadata, err error) {

	// TODO: implement dump

	return dumps, nil
}

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricMainChannelEtcdRaftOptionsApplyConfiguration represents an declarative configuration of the FabricMainChannelEtcdRaftOptions type for use
// with apply.
type FabricMainChannelEtcdRaftOptionsApplyConfiguration struct {
	TickInterval         *string `json:"tickInterval,omitempty"`
	ElectionTick         *uint32 `json:"electionTick,omitempty"`
	HeartbeatTick        *uint32 `json:"heartbeatTick,omitempty"`
	MaxInflightBlocks    *uint32 `json:"maxInflightBlocks,omitempty"`
	SnapshotIntervalSize *uint32 `json:"snapshotIntervalSize,omitempty"`
}

// FabricMainChannelEtcdRaftOptionsApplyConfiguration constructs an declarative configuration of the FabricMainChannelEtcdRaftOptions type for use with
// apply.
func FabricMainChannelEtcdRaftOptions() *FabricMainChannelEtcdRaftOptionsApplyConfiguration {
	return &FabricMainChannelEtcdRaftOptionsApplyConfiguration{}
}

// WithTickInterval sets the TickInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TickInterval field is set to the value of the last call.
func (b *FabricMainChannelEtcdRaftOptionsApplyConfiguration) WithTickInterval(value string) *FabricMainChannelEtcdRaftOptionsApplyConfiguration {
	b.TickInterval = &value
	return b
}

// WithElectionTick sets the ElectionTick field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ElectionTick field is set to the value of the last call.
func (b *FabricMainChannelEtcdRaftOptionsApplyConfiguration) WithElectionTick(value uint32) *FabricMainChannelEtcdRaftOptionsApplyConfiguration {
	b.ElectionTick = &value
	return b
}

// WithHeartbeatTick sets the HeartbeatTick field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeartbeatTick field is set to the value of the last call.
func (b *FabricMainChannelEtcdRaftOptionsApplyConfiguration) WithHeartbeatTick(value uint32) *FabricMainChannelEtcdRaftOptionsApplyConfiguration {
	b.HeartbeatTick = &value
	return b
}

// WithMaxInflightBlocks sets the MaxInflightBlocks field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxInflightBlocks field is set to the value of the last call.
func (b *FabricMainChannelEtcdRaftOptionsApplyConfiguration) WithMaxInflightBlocks(value uint32) *FabricMainChannelEtcdRaftOptionsApplyConfiguration {
	b.MaxInflightBlocks = &value
	return b
}

// WithSnapshotIntervalSize sets the SnapshotIntervalSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SnapshotIntervalSize field is set to the value of the last call.
func (b *FabricMainChannelEtcdRaftOptionsApplyConfiguration) WithSnapshotIntervalSize(value uint32) *FabricMainChannelEtcdRaftOptionsApplyConfiguration {
	b.SnapshotIntervalSize = &value
	return b
}

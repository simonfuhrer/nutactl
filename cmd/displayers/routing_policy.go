// Copyright © 2020 Simon Fuhrer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package displayers

import (
	"fmt"
	"io"
	"strings"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// RoutingPolicies wraps a nutanix RoutingPolicyListIntent.
type RoutingPolicies struct {
	schema.RoutingPolicyListIntent
}

// var _ Displayable = &Projects{}

func (o RoutingPolicies) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o RoutingPolicies) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o RoutingPolicies) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o RoutingPolicies) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o RoutingPolicies) header() []string {
	return []string{
		"UUID",
		"VPC",
		"Priority",
		"Source",
		"Action",
		"Bidirectional",
		"Destination",
		"Traffic",
		"UpdatedAt",
		"CreatedAt",
	}
}

func (o RoutingPolicies) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, r := range o.Entities {
		source := r.Status.Resources.Source.AddressType
		if source == "" {
			source = fmt.Sprintf("%v/%v", r.Status.Resources.Source.IPSubnet.IP, r.Status.Resources.Source.IPSubnet.PrefixLength)
		}

		target := r.Status.Resources.Destination.AddressType
		if target == "" {
			target = fmt.Sprintf("%v/%v", r.Status.Resources.Destination.IPSubnet.IP, r.Status.Resources.Destination.IPSubnet.PrefixLength)
		}
		t := r.Status.Resources.ProtocolType

		if t != "ALL" {
			destRange := "Any"
			srcRange := "Any"
			var srcRangeList []string
			var destRangeList []string
			if r.Status.Resources.ProtocolParameters.Icmp != nil {
				t = fmt.Sprintf("%v ICMPCode: %v ICMPType: %v", t, r.Status.Resources.ProtocolParameters.Icmp.IcmpCode, r.Status.Resources.ProtocolParameters.Icmp.IcmpType)
			} else if r.Status.Resources.ProtocolParameters.TCP != nil {
				if len(r.Status.Resources.ProtocolParameters.TCP.SourcePortRangeList) > 0 {
					srcRangeList = append(srcRangeList, fmt.Sprintf("%v - %v", r.Status.Resources.ProtocolParameters.TCP.SourcePortRange.StartPort, r.Status.Resources.ProtocolParameters.TCP.SourcePortRange.EndPort))
				}
				if len(r.Status.Resources.ProtocolParameters.TCP.DestinationPortRangeList) > 0 {
					for _, v := range r.Status.Resources.ProtocolParameters.TCP.DestinationPortRangeList {
						destRangeList = append(destRangeList, fmt.Sprintf("%v-%v", v.StartPort, v.EndPort))
					}
				}

			} else if r.Status.Resources.ProtocolParameters.UDP != nil {
				if len(r.Status.Resources.ProtocolParameters.UDP.SourcePortRangeList) > 0 {
					srcRangeList = append(srcRangeList, fmt.Sprintf("%v - %v", r.Status.Resources.ProtocolParameters.UDP.SourcePortRange.StartPort, r.Status.Resources.ProtocolParameters.UDP.SourcePortRange.EndPort))
				}
				if len(r.Status.Resources.ProtocolParameters.UDP.DestinationPortRangeList) > 0 {
					for _, v := range r.Status.Resources.ProtocolParameters.UDP.DestinationPortRangeList {
						destRangeList = append(destRangeList, fmt.Sprintf("%v-%v", v.StartPort, v.EndPort))
					}
				}
			}

			if len(destRangeList) > 0 {
				destRange = strings.Join(destRangeList, ",")
			}

			if len(srcRangeList) > 0 {
				srcRange = strings.Join(srcRangeList, ",")
			}
			t = fmt.Sprintf("%v Src: %v Dest: %v", t, srcRange, destRange)
		}

		data[i] = []string{
			r.Metadata.UUID,
			r.Status.Resources.VirtualNetworkReference.Name,
			fmt.Sprintf("%v", r.Status.Resources.Priority),
			source,
			r.Status.Resources.Action.Action,
			fmt.Sprintf("%v", r.Status.Resources.IsBidirectional),
			target,
			t,
			RenderTime(r.Metadata.LastUpdateTime),
			RenderTime(r.Metadata.CreationTime),
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o RoutingPolicies) Text(w io.Writer) error {
	return nil
}

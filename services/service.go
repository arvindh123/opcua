// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"github.com/wmnsk/gopcua/datatypes"
	"github.com/wmnsk/gopcua/errors"
	"github.com/wmnsk/gopcua/ua"
)

// ServiceType definitions.
const (
	ServiceTypeFindServersRequest           uint16 = 422
	ServiceTypeFindServersResponse          uint16 = 425
	ServiceTypeGetEndpointsRequest          uint16 = 428
	ServiceTypeGetEndpointsResponse         uint16 = 431
	ServiceTypeOpenSecureChannelRequest     uint16 = 446
	ServiceTypeOpenSecureChannelResponse    uint16 = 449
	ServiceTypeCloseSecureChannelRequest    uint16 = 452
	ServiceTypeCloseSecureChannelResponse   uint16 = 455
	ServiceTypeCreateSessionRequest         uint16 = 461
	ServiceTypeCreateSessionResponse        uint16 = 464
	ServiceTypeActivateSessionRequest       uint16 = 467
	ServiceTypeActivateSessionResponse      uint16 = 470
	ServiceTypeCloseSessionRequest          uint16 = 473
	ServiceTypeCloseSessionResponse         uint16 = 476
	ServiceTypeCancelRequest                uint16 = 479
	ServiceTypeCancelResponse               uint16 = 482
	ServiceTypeReadRequest                  uint16 = 631
	ServiceTypeReadResponse                 uint16 = 634
	ServiceTypeWriteRequest                 uint16 = 673
	ServiceTypeWriteResponse                uint16 = 676
	ServiceTypeCreateSubscriptionRequest    uint16 = 787
	ServiceTypeCreateSubscriptionResponse   uint16 = 790
	ServiceTypeFindServersOnNetworkRequest  uint16 = 12208
	ServiceTypeFindServersOnNetworkResponse uint16 = 12211
)

type Service interface {
	ServiceType() uint16
}

// Decode decodes given bytes into Service, depending on the type of service.
func Decode(b []byte) (Service, error) {
	// peek at the type id without stripping it from the buffer to determine
	// the type of service. The type id will be decoded again with the reflective
	// decoder.
	typeID := new(datatypes.ExpandedNodeID)
	_, err := typeID.Decode(b)
	if err != nil {
		return nil, errors.NewErrUnsupported(typeID, "cannot decode TypeID.")
	}
	if typeID.NodeID.Type() != datatypes.NodeIDTypeFourByte {
		return nil, errors.NewErrUnsupported(typeID.NodeID, "should be FourByteNodeID.")
	}

	var s Service
	id := uint16(typeID.NodeID.IntID())
	// log.Printf("Decode: id:%d %d bytes %x", id, len(b), b)
	switch id {
	case ServiceTypeFindServersRequest:
		s = &FindServersRequest{}
	case ServiceTypeFindServersResponse:
		s = &FindServersResponse{}
	case ServiceTypeGetEndpointsRequest:
		s = &GetEndpointsRequest{}
	case ServiceTypeGetEndpointsResponse:
		s = &GetEndpointsResponse{}
	case ServiceTypeOpenSecureChannelRequest:
		s = &OpenSecureChannelRequest{}
	case ServiceTypeOpenSecureChannelResponse:
		s = &OpenSecureChannelResponse{}
	case ServiceTypeCloseSecureChannelRequest:
		s = &CloseSecureChannelRequest{}
	case ServiceTypeCloseSecureChannelResponse:
		s = &CloseSecureChannelResponse{}
	case ServiceTypeCreateSessionRequest:
		s = &CreateSessionRequest{}
	case ServiceTypeCreateSessionResponse:
		s = &CreateSessionResponse{}
	case ServiceTypeActivateSessionRequest:
		s = &ActivateSessionRequest{}
	case ServiceTypeActivateSessionResponse:
		s = &ActivateSessionResponse{}
	case ServiceTypeCloseSessionRequest:
		s = &CloseSessionRequest{}
	case ServiceTypeCloseSessionResponse:
		s = &CloseSessionResponse{}
	case ServiceTypeCancelRequest:
		s = &CancelRequest{}
	case ServiceTypeCancelResponse:
		s = &CancelResponse{}
	case ServiceTypeReadRequest:
		s = &ReadRequest{}
	case ServiceTypeReadResponse:
		s = &ReadResponse{}
	case ServiceTypeWriteRequest:
		s = &WriteRequest{}
	case ServiceTypeWriteResponse:
		s = &WriteResponse{}
	case ServiceTypeCreateSubscriptionRequest:
		s = &CreateSubscriptionRequest{}
	case ServiceTypeCreateSubscriptionResponse:
		s = &CreateSubscriptionResponse{}
	case ServiceTypeFindServersOnNetworkRequest:
		s = &FindServersOnNetworkRequest{}
	case ServiceTypeFindServersOnNetworkResponse:
		s = &FindServersOnNetworkResponse{}
	default:
		return nil, errors.NewErrUnsupported(id, "unsupported or not implemented yet.")
	}

	if err := ua.Decode(b, s); err != nil {
		return nil, err
	}
	return s, nil
}

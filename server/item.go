package server

import (
	pb "github.com/WiggidyW/eve-item-parser-server-go/proto"
)

type ParsedItem struct {
	Name     string
	Quantity int64
}

func (pi *ParsedItem) IntoKnown(type_id uint32) *pb.KnownItem {
	return &pb.KnownItem{
		Name:     pi.Name,
		Quantity: pi.Quantity,
		TypeId:   type_id,
	}
}

func (pi *ParsedItem) IntoUnknown() *pb.UnknownItem {
	return &pb.UnknownItem{
		Name:     pi.Name,
		Quantity: pi.Quantity,
	}
}

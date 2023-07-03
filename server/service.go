package server

import (
	"context"

	pb "github.com/WiggidyW/eve-item-parser-server-go/proto"
)

type Service struct {
	db *Db
}

func NewService(db_url string, db_max_readers int) (*Service, error) {
	db, err := NewDb(db_url, db_max_readers)
	if err != nil {
		return nil, err
	} else {
		return &Service{db}, nil
	}
}

func (sv *Service) Parse(
	ctx context.Context,
	req *pb.ParseReq,
) (*pb.ParseRep, error) {
	parsed_items := Parse(req.Text)
	return sv.db.QueryNames(ctx, parsed_items, req.Language)
}

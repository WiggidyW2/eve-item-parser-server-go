package server

import (
	"context"
	"database/sql"

	pb "github.com/WiggidyW/eve-item-parser-server-go/proto"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	max_readers int
	*sql.DB
}

type queryNameResult struct {
	KnownItem   *pb.KnownItem
	UnknownItem *pb.UnknownItem
	Err         error
}

func NewDb(url string, max_readers int) (*Db, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, err
	} else {
		return &Db{max_readers, db}, nil
	}
}

func (db *Db) QueryNames(
	ctx context.Context,
	items []*ParsedItem,
) (*pb.ParseRep, error) {
	items_len := len(items)
	parse_rep := &pb.ParseRep{
		KnownItems:   make([]*pb.KnownItem, 0, items_len),
		UnknownItems: make([]*pb.UnknownItem, 0, items_len),
	}
	chn := make(chan *queryNameResult, db.max_readers)

	var item_idx = 0
	var reader_count = 0
	for {
		if reader_count == db.max_readers || item_idx == items_len {
			query_rep := <-chn
			err := updateParseRep(parse_rep, query_rep)
			if err != nil {
				return nil, err
			}
			reader_count--
		}
		if item_idx < items_len {
			go db.queryName(ctx, items[item_idx], chn)
			item_idx++
			reader_count++
		}
		if reader_count == 0 {
			return parse_rep, nil
		}
	}
}

func updateParseRep(
	parse_rep *pb.ParseRep,
	query_rep *queryNameResult,
) error {
	if query_rep.KnownItem != nil {
		parse_rep.KnownItems = append(
			parse_rep.KnownItems,
			query_rep.KnownItem,
		)
	} else if query_rep.UnknownItem != nil {
		parse_rep.UnknownItems = append(
			parse_rep.UnknownItems,
			query_rep.UnknownItem,
		)
	} else /* if query_rep.Err != nil */ {
		return query_rep.Err
	}
	return nil
}

func (db *Db) queryName(
	ctx context.Context,
	item *ParsedItem,
	chn chan *queryNameResult,
) {
	var type_id uint32
	var db_result = new(queryNameResult)

	err := db.QueryRowContext(
		ctx,
		"SELECT type_id FROM types WHERE name = ?",
		item.Name,
	).Scan(&type_id)

	if err != nil {
		db_result.Err = err
	} else if type_id != 0 {
		db_result.KnownItem = item.IntoKnown(type_id)
	} else /* if type_id == 0 */ {
		db_result.UnknownItem = item.IntoUnknown()
	}

	chn <- db_result
}

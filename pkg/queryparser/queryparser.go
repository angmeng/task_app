package queryparser

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/a8m/rql"
	"github.com/gofiber/fiber/v2"
)

type Params struct {
	Filter map[string]interface{} `json:"filter,omitempty"`
	Page   int                    `json:"page"`
	Size   int                    `json:"size"`
	Sort   []string               `json:"sort"`
}

type Query struct {
	Filter []interface{}
	Page   int
	Size   int
	Sort   string
}

func GetDBQuery(c *fiber.Ctx, model interface{}) (*Query, error) {
	var (
		b      []byte
		err    error
		params = Params{}
		rq     = Query{}
		// QueryParam is the name of the query string key.
		QueryParam = "query"
		// MustNewParser panics if the configuration is invalid.
		QueryParser = rql.MustNewParser(rql.Config{
			Model:         model,
			FieldSep:      ".",
			LimitMaxValue: 50,
		})
	)

	if v := c.Query(QueryParam); v != "" {
		b, err = base64.StdEncoding.DecodeString(v)
	} else {
		return &rq, nil
	}

	if err != nil {
		return &rq, err
	}

	if err := json.Unmarshal(b, &params); err != nil {
		return &rq, err
	}

	if params.Filter == nil {
		params.Filter = map[string]interface{}{}
	}

	q, err := QueryParser.ParseQuery(&rql.Query{Filter: params.Filter, Sort: params.Sort})
	if err != nil {
		return &rq, err
	}

	if q.FilterExp != "" {
		q.FilterExp = strings.Replace(q.FilterExp, "LIKE", "ILIKE", -1)
		rq.Filter = parse(q.FilterExp, q.FilterArgs)
	}

	if params.Size <= 0 {
		rq.Size = 10
	} else {
		rq.Size = params.Size
	}

	rq.Page = params.Page
	rq.Sort = q.Sort

	fmt.Println("q", rq)
	return &rq, err
}

func parse(expr string, args []interface{}) []interface{} {
	if expr == "" {
		return nil
	}
	qry := make([]interface{}, 0, len(args)+1)
	qry = append(qry, expr)
	qry = append(qry, args...)
	return qry
}

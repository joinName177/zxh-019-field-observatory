package observatory

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func WriteCSV(w io.Writer, l Layer) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"x", "y"}); err != nil {
		return err
	}
	for _, p := range l.Points {
		if err := c.Write([]string{strconv.FormatFloat(p.X, 'f', 6, 64), strconv.FormatFloat(p.Y, 'f', 6, 64)}); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
func ReadCSV(r io.Reader) ([]Point, error) {
	c := csv.NewReader(r)
	if _, err := c.Read(); err != nil {
		return nil, err
	}
	out := []Point{}
	for {
		row, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(row) != 2 {
			return nil, fmt.Errorf("invalid coordinate row")
		}
		x, e := strconv.ParseFloat(strings.TrimSpace(row[0]), 64)
		if e != nil {
			return nil, e
		}
		y, e := strconv.ParseFloat(strings.TrimSpace(row[1]), 64)
		if e != nil {
			return nil, e
		}
		out = append(out, Point{x, y})
	}
	return out, nil
}

package builder

import "errors"

type date struct {
	year  int
	month int
	day   int
}

type builder struct {
	year  int
	month int
	day   int
}

func (b builder) new() (*date, error) {
	if b.day > 31 {
		return nil, errors.New("invalid date")
	}
	switch b.month {
	case 2:
		if b.day > 28 {
			if !((b.year%4 == 0) && (b.day == 29)) {
				return nil, errors.New("invalid date")
			}
		}
	case 4, 6, 9, 11:
		if b.day > 30 {
			return nil, errors.New("invalid date")
		}
	}
	d := date{}
	d.year = b.year
	d.month = b.month
	d.day = b.day
	return &d, nil
}

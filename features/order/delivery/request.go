package delivery

import "altaproject2/domain"

type OrderFormat struct {
	Userid int
	Total  int
}

func (i *OrderFormat) ToModel() domain.Order {
	return domain.Order{
		Userid: i.Userid,
		Total:  i.Total,
	}
}

package builder

import "testing"

func TestBuilder(t *testing.T) {
	o := NewOrm()
	o.QueryTable("user").Where("user_id = ", 1).All()
}

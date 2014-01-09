package rel

import (
	"testing"
)

func TestAttributeNotEqSql(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" != 10"
	if sql != expected {
		t.Logf("TestAttributeNotEqSql sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" != 1 OR \"users\".\"id\" != 2)"
	if sql != expected {
		t.Logf("TestAttributeNotEqAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqNil(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEq(nil))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IS NOT NULL"
	if sql != expected {
		t.Logf("TestAttributeNotEqNil sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" != 1 AND \"users\".\"id\" != 2)"
	if sql != expected {
		t.Logf("TestAttributeNotEqAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGt(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Gt(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" > 10"
	if sql != expected {
		t.Logf("TestAttributeGt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtEq(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" >= 10"
	if sql != expected {
		t.Logf("TestAttributeGtEq sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtEqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" >= 1 OR \"users\".\"id\" >= 2)"
	if sql != expected {
		t.Logf("TestAttributeGtEqAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtEqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" >= 1 AND \"users\".\"id\" >= 2)"
	if sql != expected {
		t.Logf("TestAttributeGtEqAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" > 1 AND \"users\".\"id\" > 2)"
	if sql != expected {
		t.Logf("TestAttributeGtAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" > 1 OR \"users\".\"id\" > 2)"
	if sql != expected {
		t.Logf("TestAttributeGtAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLt(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Lt(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" < 10"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtEq(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" <= 10"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtEqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" <= 1 OR \"users\".\"id\" <= 2)"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtEqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" <= 1 AND \"users\".\"id\" <= 2)"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" < 1 OR \"users\".\"id\" < 2)"
	if sql != expected {
		t.Logf("TestAttributeLtAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" < 1 AND \"users\".\"id\" < 2)"
	if sql != expected {
		t.Logf("TestAttributeLtAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeCount(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id").Count())
	sql := mgr.ToSql()
	expected := "SELECT COUNT(\"users\".\"id\") FROM \"users\""
	if sql != expected {
		t.Logf("TestAttributeCount sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeEq(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Eq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" = 10"
	if sql != expected {
		t.Logf("TestAttributeEq sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeEqNil(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Eq(nil))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IS NULL"
	if sql != expected {
		t.Logf("TestAttributeEqNil sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").EqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" = 1 OR \"users\".\"id\" = 2)"
	if sql != expected {
		t.Logf("TestAttributeEqAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").EqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" = 1 AND \"users\".\"id\" = 2)"
	if sql != expected {
		t.Logf("TestAttributeEqAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeMatches(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("name").Matches(Sql("%bacon%")))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"name\" LIKE '%bacon%'"
	if sql != expected {
		t.Logf("TestAttributeMatches sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeMatchesAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("name").MatchesAny(Sql("%chunky%"), Sql("%bacon%")))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" LIKE '%chunky%' OR \"users\".\"name\" LIKE '%bacon%')"
	if sql != expected {
		t.Logf("TestAttributeMatchesAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeMatchesAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("name").MatchesAll(Sql("%chunky%"), Sql("%bacon%")))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" LIKE '%chunky%' AND \"users\".\"name\" LIKE '%bacon%')"
	if sql != expected {
		t.Logf("TestAttributeMatchesAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeDoesNotMatch(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("name").DoesNotMatch(Sql("%bacon%")))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"name\" NOT LIKE '%bacon%'"
	if sql != expected {
		t.Logf("TestAttributeDoesNotMatch sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeDoesNotMatchAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("name").DoesNotMatchAny(Sql("%chunky%"), Sql("%bacon%")))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" NOT LIKE '%chunky%' OR \"users\".\"name\" NOT LIKE '%bacon%')"
	if sql != expected {
		t.Logf("TestAttributeDoesNotMatchAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeDoesNotMatchAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("name").DoesNotMatchAll(Sql("%chunky%"), Sql("%bacon%")))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"name\" NOT LIKE '%chunky%' AND \"users\".\"name\" NOT LIKE '%bacon%')"
	if sql != expected {
		t.Logf("TestAttributeDoesNotMatchAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

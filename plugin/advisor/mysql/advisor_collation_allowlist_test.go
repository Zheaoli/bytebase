package mysql

// Framework code is generated by the generator.

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytebase/bytebase/plugin/advisor"
)

func TestCollationAllowlist(t *testing.T) {
	tests := []advisor.TestCase{
		{
			Statement: `CREATE TABLE t(a int) COLLATE utf8mb4_polish_ci`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			Statement: `CREATE TABLE t(a varchar(255))`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			Statement: `CREATE TABLE t(a int) COLLATE latin1_bin`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"CREATE TABLE t(a int) COLLATE latin1_bin\" used disabled collation 'latin1_bin'",
					Line:    1,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t COLLATE latin1_bin`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"ALTER TABLE t COLLATE latin1_bin\" used disabled collation 'latin1_bin'",
					Line:    3,
				},
			},
		},
		{
			Statement: `ALTER DATABASE test COLLATE latin1_bin`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"ALTER DATABASE test COLLATE latin1_bin\" used disabled collation 'latin1_bin'",
					Line:    1,
				},
			},
		},
		{
			Statement: `CREATE TABLE t(a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"CREATE TABLE t(a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin)\" used disabled collation 'latin1_bin'",
					Line:    1,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(b int);
				ALTER TABLE t ADD COLUMN a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"ALTER TABLE t ADD COLUMN a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin\" used disabled collation 'latin1_bin'",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t MODIFY COLUMN a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"ALTER TABLE t MODIFY COLUMN a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin\" used disabled collation 'latin1_bin'",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t CHANGE COLUMN a a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DisabledCollation,
					Title:   "collation.allowlist",
					Content: "\"ALTER TABLE t CHANGE COLUMN a a varchar(255) CHARACTER SET latin1 COLLATE latin1_bin\" used disabled collation 'latin1_bin'",
					Line:    3,
				},
			},
		},
	}

	payload, err := json.Marshal(advisor.AllowlistRulePayload{
		Allowlist: []string{"utf8mb4_polish_ci"},
	})
	require.NoError(t, err)
	advisor.RunSQLReviewRuleTests(t, tests, &CollationAllowlistAdvisor{}, &advisor.SQLReviewRule{
		Type:    advisor.SchemaRuleCollationAllowlist,
		Level:   advisor.SchemaRuleLevelWarning,
		Payload: string(payload),
	}, advisor.MockMySQLDatabase)
}

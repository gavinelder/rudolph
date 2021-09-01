package machinerules

import (
	"time"

	"github.com/airbnb/rudolph/pkg/dynamodb"
	"github.com/airbnb/rudolph/pkg/types"
	"github.com/pkg/errors"
)

// @deprecated
func UpdateMachineRule(client dynamodb.UpdateItemAPI, machineID string, sha256 string, ruleType types.RuleType, rulePolicy types.Policy, expires time.Time) (err error) {
	pk := machineRulePK(machineID)
	sk := machineRuleSK(sha256, ruleType)

	_, err = client.UpdateItem(
		dynamodb.PrimaryKey{
			PartitionKey: pk,
			SortKey:      sk,
		},
		updateRulePolicyRequest{
			Policy:       rulePolicy,
			ExpiresAfter: expires.Unix(),
		},
	)

	if err != nil {
		err = errors.Wrapf(err, "failed to update machine rule")
		return
	}

	return
}

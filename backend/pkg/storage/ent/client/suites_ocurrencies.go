package client

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/redhat-appstudio/quality-studio/api/apis/prow/v1alpha1"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowsuites"
)

func (d *Database) GetSuitesFailureFrequency(startDate, endDate string) ([]v1alpha1.SuitesFailureFrequency, error) {
	var suitesFailure []v1alpha1.SuitesFailureFrequency

	err := d.client.ProwSuites.Query().
		Where(prowsuites.Status("failed")).
		Where(func(s *sql.Selector) { // "merged_at BETWEEN ? AND 2022-08-17", "2022-08-16"
			s.Where(sql.ExprP(fmt.Sprintf("created_at BETWEEN '%s' AND '%s'", "2023-10-25", fmt.Sprintf("%s:23:59:59", "2023-10-31"))))
		}). // myDate >= '20090101 00:00:00' AND myDate < '20090201 00:00:00'  --CORRECT!
		GroupBy(prowsuites.FieldName, prowsuites.FieldStatus, prowsuites.FieldErrorMessage).
		Aggregate(db.Count()).
		Scan(context.Background(), &suitesFailure)

	fmt.Println(err)

	fmt.Println(suitesFailure)

	return suitesFailure, nil
}

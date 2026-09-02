package cmd

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/whughesiii2187/wch-projects/Go/budgetcli/internal/models"
)

var updateCmd = &cobra.Command{
	Use:  "update",
	Long: "Make updates to bills that have been prepped or activate / deactivate bills",
	Run:  runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) {
	formOptions := huh.NewForm(
		huh.NewGroup(
				huh.NewSelect[string]().
					Title("Which update method to you choose?"). 
					Options(
					huh.NewOption("Update upcoming payment", "payment"),
					huh.NewOption("Deactivate a bill", "deactivate"),
			
					
				)
			)
		)
}

func pullUpdates() ([]models.UpdatePayments, error) {
	rows, err := DB.Query(
		context.Background(),
		fmt.Sprintln("select a.payment_id,a.bill_id,b.bill_name,a.due_date,a.amount_due,a.paid from bill_history as a join bills as b on a.bill_id = b.bill_id WHERE a.paid = false"),
	)
	if err != nil {
		fmt.Println("Error querying bills:", err)
		return nil, err
	}
	defer rows.Close()

	var update []models.UpdatePayments

	for rows.Next() {
		var updates models.UpdatePayments

		err := rows.Scan(
			&updates.PaymentId,
			&updates.BillId,
			&updates.BillName,
			&updates.DueDate,
			&updates.AmountDue,
			&updates.Paid,
		)
		if err != nil {
			fmt.Println("Error reading rows:", err)
			return nil, err
		}
		update = append(update, updates)
	}
	return update, nil
}

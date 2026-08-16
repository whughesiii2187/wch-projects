package cmd

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
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

type model struct {
	tbl table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.tbl, cmd = m.tbl.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.tbl.View()
}

func runUpdate(cmd *cobra.Command, args []string) {
	updates, _ := pullUpdates()

	[]table.Column{
		{Title: "Bill Name", Width: 20},
		{Title: "Due Date", Width: 12},
		{Title: "Amount Due", Width: 12},
		{Title: "Paid", Width: 8},
	}

	fmt.Printf("Rows: %v", updates)
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

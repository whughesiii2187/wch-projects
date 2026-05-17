package cmd

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/whughesiii2187/wch-projects/Go/budgetcli/internal/models"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a budget item",
	Run:   runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) {
	var billType string
	var balanceInput string
	var payPeriodInput string
	var dueAmountInput string
	var dueRecurringDateInput string
	var bill models.Bill

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter Bill Name:").
				Prompt("?").
				Value(&bill.BillName),
			huh.NewSelect[string]().
				Title("Select the Bill Type:").
				Options(
					huh.NewOption("Recurring", "recurring"),
					huh.NewOption("Credit/Loan", "credit/loan"),
				).
				Value(&billType),
			huh.NewConfirm().
				Title("Is this bill an AutoPay?").
				Affirmative("Auto Pay").
				Negative("Manual Pay").
				Value(&bill.IsAutoPay),
			huh.NewSelect[string]().
				Title("Which Pay Period?").
				Options(
					huh.NewOption("A", "A"),
					huh.NewOption("B", "B"),
				).
				Value(&payPeriodInput),
			huh.NewInput().
				Title("Enter current due amount, if any").
				Prompt("$").
				Value(&dueAmountInput).
				Validate(func(s string) error {
					_, err := strconv.ParseFloat(s, 64)
					if err != nil {
						return fmt.Errorf("please enter a valid number")
					}
					return nil
				}),
			huh.NewInput().
				Title("Enter the date due (day of the month)").
				Prompt(":").
				Value(&dueRecurringDateInput).
				Validate(func(s string) error {
					_, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("please enter a valid date")
					}
					return nil
				}),
			huh.NewInput().
				Title("Additional Notes:").
				Prompt("?").
				Value(&bill.Notes),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the current balance due:").
				Prompt("$").
				Value(&balanceInput).
				Validate(func(s string) error {
					_, err := strconv.ParseFloat(s, 64)
					if err != nil {
						return fmt.Errorf("please enter a valid amount")
					}
					return nil
				}),
		).WithHideFunc(func() bool { return billType != "credit/loan" }),
	)

	err := form.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	bill.PayPeriodPaid = models.PayPeriod(payPeriodInput)

	if balanceInput != "" {
		num, err := strconv.ParseFloat(balanceInput, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		bill.DueBalance = &num
	}
	bill.DueRecurringDate, _ = strconv.Atoi(dueRecurringDateInput)
	bill.DueAmount, _ = strconv.ParseFloat(dueAmountInput, 64)

	fmt.Printf("%+v\n", bill)
}

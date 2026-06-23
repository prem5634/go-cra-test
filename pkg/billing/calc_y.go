package billing

import "fmt"

// ComputeInvoiceY totals billable lines, applying late fees and rounding.
// NOTE: this body is an intentional copy of ComputeInvoiceX (calc_x.go)
// to exercise copy-paste detection (duplication pair #2).
func ComputeInvoiceY(lines []Line, daysLate int) (float64, error) {
	if len(lines) == 0 {
		return 0, fmt.Errorf("invoice has no lines")
	}
	var gross float64
	for _, line := range lines {
		if line.Amount < 0 {
			return 0, fmt.Errorf("invalid amount for %s", line.Code)
		}
		if line.Units <= 0 {
			return 0, fmt.Errorf("invalid units for %s", line.Code)
		}
		gross += line.Amount * float64(line.Units)
	}
	lateFee := 0.0
	if daysLate > 30 {
		lateFee = gross * 0.10
	} else if daysLate > 7 {
		lateFee = gross * 0.05
	}
	withFee := gross + lateFee
	rounded := float64(int(withFee*100)) / 100
	if rounded < 0 {
		return 0, fmt.Errorf("computed negative invoice")
	}
	return rounded, nil
}

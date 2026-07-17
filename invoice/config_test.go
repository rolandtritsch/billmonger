package invoice

import (
	"strings"
	"testing"
)

func TestVATAmount(t *testing.T) {
	bill := BillDetails{VATTaxRate: 23}
	got := bill.VATAmount(9000)
	if got != 2070 {
		t.Fatalf("VATAmount(9000) = %v, want 2070", got)
	}
}

func TestVATAmountRoundsToNearestCent(t *testing.T) {
	bill := BillDetails{VATTaxRate: 23}
	got := bill.VATAmount(1439.16)
	if got != 331.01 {
		t.Fatalf("VATAmount(1439.16) = %v, want 331.01", got)
	}
}

func TestZeroVATAmount(t *testing.T) {
	bill := BillDetails{VATTaxID: "-", VATTaxRate: 0}
	got := bill.VATAmount(9000)
	if got != 0 {
		t.Fatalf("VATAmount(9000) = %v, want 0", got)
	}
}

func TestValidateVATRejectsBlankID(t *testing.T) {
	bill := BillDetails{VATTaxRate: 0}
	err := bill.ValidateVAT()
	if err == nil || !strings.Contains(err.Error(), "cannot be blank") {
		t.Fatalf("ValidateVAT() error = %v, want blank VAT ID error", err)
	}
}

func TestValidateVATRejectsRateForDashID(t *testing.T) {
	bill := BillDetails{VATTaxID: "-", VATTaxRate: 23}
	err := bill.ValidateVAT()
	if err == nil || !strings.Contains(err.Error(), "must be 0") {
		t.Fatalf("ValidateVAT() error = %v, want zero-rate error", err)
	}
}

func TestValidateVATAllowsDashIDWithZeroRate(t *testing.T) {
	bill := BillDetails{VATTaxID: "-", VATTaxRate: 0}
	if err := bill.ValidateVAT(); err != nil {
		t.Fatalf("ValidateVAT() error = %v, want nil", err)
	}
}

func TestNiceFloatStrRoundsToNearestCent(t *testing.T) {
	got := niceFloatStr(1770.1668)
	if got != "1,770.17" {
		t.Fatalf("niceFloatStr(1770.1668) = %q, want %q", got, "1,770.17")
	}
}

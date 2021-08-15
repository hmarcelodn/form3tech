package model

type AccountBuilderGB struct {
	bankID        string
	bic           string
	bankIDCode    string
	accountNumber string
	iban          string
	name          string
}

func (ab AccountBuilderGB) SetBankID(bankID string) iAccountBuilder {
	ab.bankID = bankID
	return ab
}

func (ab AccountBuilderGB) SetBic(bic string) iAccountBuilder {
	ab.bic = bic
	return ab
}

func (ab AccountBuilderGB) SetAccountNumber(accountNumber string) iAccountBuilder {
	ab.accountNumber = accountNumber
	return ab
}

func (ab AccountBuilderGB) SetIban(iban string) iAccountBuilder {
	ab.iban = iban
	return ab
}

func (ab AccountBuilderGB) SetName(name string) iAccountBuilder {
	ab.name = name
	return ab
}

func (ab AccountBuilderGB) Build() AccountAttributes {
	country := "GB"
	var name []string
	name = append(name, ab.name)
	return AccountAttributes{
		Country:       &country,
		BankID:        ab.bankID,
		Bic:           ab.bic,
		BankIDCode:    "GBDSC",
		AccountNumber: ab.accountNumber,
		Iban:          ab.iban,
		Name:          name,
	}
}

package model

type AccountBuilderAU struct {
	bankID        string
	bic           string
	bankIDCode    string
	accountNumber string
	iban          string
	name          string
}

func (ab AccountBuilderAU) SetBankID(bankID string) iAccountBuilder {
	ab.bankID = bankID
	return ab
}

func (ab AccountBuilderAU) SetBic(bic string) iAccountBuilder {
	ab.bic = bic
	return ab
}

func (ab AccountBuilderAU) SetAccountNumber(accountNumber string) iAccountBuilder {
	ab.accountNumber = accountNumber
	return ab
}

func (ab AccountBuilderAU) SetIban(iban string) iAccountBuilder {
	ab.iban = iban
	return ab
}

func (ab AccountBuilderAU) SetName(name string) iAccountBuilder {
	ab.name = name
	return ab
}

func (ab AccountBuilderAU) Build() AccountAttributes {
	country := "AU"
	var name []string
	name = append(name, ab.name)
	return AccountAttributes{
		Country:       &country,
		BankID:        ab.bankID,
		Bic:           ab.bic,
		BankIDCode:    "AUBSB",
		AccountNumber: ab.accountNumber,
		Iban:          ab.iban,
		Name:          name,
	}
}

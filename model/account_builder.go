package model

type AccountBuilder struct {
	country       string
	bankID        string
	bic           string
	bankIDCode    string
	accountNumber string
	iban          string
	name          string
}

func (ab AccountBuilder) SetCountry(country string) AccountBuilder {
	ab.country = country
	return ab
}

func (ab AccountBuilder) SetBankID(bankID string) AccountBuilder {
	ab.bankID = bankID
	return ab
}

func (ab AccountBuilder) SetBic(bic string) AccountBuilder {
	ab.bic = bic
	return ab
}

func (ab AccountBuilder) SetAccountNumber(accountNumber string) AccountBuilder {
	ab.accountNumber = accountNumber
	return ab
}

func (ab AccountBuilder) SetIban(iban string) AccountBuilder {
	ab.iban = iban
	return ab
}

func (ab AccountBuilder) SetName(name string) AccountBuilder {
	ab.name = name
	return ab
}

func (ab AccountBuilder) Build() AccountAttributes {
	country := ab.country
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

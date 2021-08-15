package model

type iAccountBuilder interface {
	SetBankID(bankID string) iAccountBuilder
	SetBic(bic string) iAccountBuilder
	SetAccountNumber(accountNumber string) iAccountBuilder
	SetIban(iban string) iAccountBuilder
	SetName(name string) iAccountBuilder
	Build() AccountAttributes
}

func GetAccountBuilder(builderType string) iAccountBuilder {
	if builderType == "GB" {
		return AccountBuilderGB{}
	}

	if builderType == "AU" {
		return AccountBuilderAU{}
	}

	return nil
}

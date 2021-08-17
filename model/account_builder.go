package model

type AccountBuilder struct {
	country                string
	bankID                 string
	bic                    string
	bankIDCode             string
	accountNumber          string
	iban                   string
	name                   string
	customerID             string
	processingService      string
	userDefinedInformation string
	validationType         string
	referenceMask          string
	acceptanceQualifier    string
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

func (ab AccountBuilder) SetCustomerID(customerID string) AccountBuilder {
	ab.customerID = customerID
	return ab
}

func (ab AccountBuilder) SetProcessingService(processingService string) AccountBuilder {
	ab.processingService = processingService
	return ab
}

func (ab AccountBuilder) SetUserDefinedInformation(userDefinedInformation string) AccountBuilder {
	ab.userDefinedInformation = userDefinedInformation
	return ab
}

func (ab AccountBuilder) SetValidationType(validationType string) AccountBuilder {
	ab.validationType = validationType
	return ab
}

func (ab AccountBuilder) SetReferenceMask(referenceMask string) AccountBuilder {
	ab.referenceMask = referenceMask
	return ab
}

func (ab AccountBuilder) SetAcceptanceQualifier(acceptanceQualifier string) AccountBuilder {
	ab.acceptanceQualifier = acceptanceQualifier
	return ab
}

func (ab AccountBuilder) Build() AccountAttributes {
	country := ab.country
	var name []string
	name = append(name, ab.name)
	return AccountAttributes{
		Country:                &country,
		BankID:                 ab.bankID,
		Bic:                    ab.bic,
		BankIDCode:             ab.bankIDCode,
		AccountNumber:          ab.accountNumber,
		Iban:                   ab.iban,
		Name:                   name,
		CustomerID:             ab.customerID,
		ProcessingService:      ab.processingService,
		UserDefinedInformation: ab.userDefinedInformation,
		ValidationType:         ab.validationType,
		ReferenceMask:          ab.referenceMask,
		AcceptanceQualifier:    ab.acceptanceQualifier,
	}
}

package authdomain

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// Contact Email Value Object
type ContactEmail string

const (
	ERRORS_CONTACT_EMAIL_EMPTY shareddomain.DomainError = "ERRORS_CONTACT_EMAIL_EMPTY"
	ERRORS_CONTACT_EMAIL_MIN   shareddomain.DomainError = "ERRORS_CONTACT_EMAIL_MIN"
)

func CreateEmail(email string) (ContactEmail, error) {
	if len(email) == 0 {
		return ContactEmail(""), errors.New(string(ERRORS_CONTACT_EMAIL_EMPTY))
	}

	// TODO: Validate email format

	return ContactEmail(email), nil
}

func CreateEmailList(min uint8, emails ...string) ([]ContactEmail, error) {
	if min == 0 && len(emails) == 1 && len(emails[0]) == 0 {
		return []ContactEmail{}, nil
	}

	if len(emails) < int(min) {
		return []ContactEmail{}, errors.New(string(ERRORS_CONTACT_EMAIL_MIN))
	}

	emailList := make([]ContactEmail, len(emails))

	for i, v := range emails {
		email, err := CreateEmail(v)

		if err != nil {
			return []ContactEmail{}, err
		}

		emailList[i] = email
	}

	return emailList, nil
}

// Contact Phone Value Object
type ContactPhone struct {
	Country Country `json:"country"`
	Phone   string  `json:"phone"`
}

const (
	ERRORS_CONTACT_PHONE_FORMAT shareddomain.DomainError = "ERRORS_CONTACT_PHONE_FORMAT"
	ERRORS_CONTACT_PHONE_MIN    shareddomain.DomainError = "ERRORS_CONTACT_PHONE_MIN"
)

func CreatePhone(country Country, phone string) (ContactPhone, error) {
	validate := validator.New()

	if err := validate.Var(phone, "min=7,max=10,numeric,excludes=.,excludes=0x2C"); err != nil {
		return ContactPhone{}, errors.New(string(ERRORS_CONTACT_PHONE_FORMAT))
	}

	return ContactPhone{
		Country: country,
		Phone:   phone,
	}, nil
}

type ContactPhoneRequest struct {
	Country uint8
	Phone   string
}

func CreatePhoneList(min uint8, locationRepo LocationRepository, phones ...ContactPhoneRequest) ([]ContactPhone, error) {
	if min == 0 && len(phones) == 1 && phones[0].Country == 0 && len(phones[0].Phone) == 0 {
		return []ContactPhone{}, nil
	}

	if len(phones) < int(min) {
		return []ContactPhone{}, errors.New(string(ERRORS_CONTACT_PHONE_MIN))
	}

	phoneList := make([]ContactPhone, len(phones))

	for i, v := range phones {
		country, countryErr := locationRepo.GetCountryById(v.Country, false, false)

		if countryErr != nil {
			return []ContactPhone{}, countryErr
		}

		phone, phoneErr := CreatePhone(country, v.Phone)

		if phoneErr != nil {
			return []ContactPhone{}, phoneErr
		}

		phoneList[i] = phone
	}

	return phoneList, nil
}

// Contact Address Value Object
type ContactAddress struct {
	Municipality Municipality `json:"municipality"`
	Address      string       `json:"address"`
}

const (
	ERRORS_CONTACT_ADDRESS_EMPTY shareddomain.DomainError = "ERRORS_CONTACT_ADDRESS_EMPTY"
	ERRORS_CONTACT_ADDRESS_MIN   shareddomain.DomainError = "ERRORS_CONTACT_ADDRESS_MIN"
)

func CreateAddress(municipality Municipality, address string) (ContactAddress, error) {
	if len(address) == 0 {
		return ContactAddress{}, errors.New(string(ERRORS_CONTACT_ADDRESS_EMPTY))
	}

	return ContactAddress{
		Municipality: municipality,
		Address:      address,
	}, nil
}

type ContactAddressRequest struct {
	Municipality uint16
	Address      string
}

func CreateAddressList(min uint8, locationRepo LocationRepository, addresses ...ContactAddressRequest) ([]ContactAddress, error) {
	if min == 0 && len(addresses) == 1 && addresses[0].Municipality == 0 && len(addresses[0].Address) == 0 {
		return []ContactAddress{}, nil
	}

	if len(addresses) < int(min) {
		return []ContactAddress{}, errors.New(string(ERRORS_CONTACT_ADDRESS_MIN))
	}

	addressList := make([]ContactAddress, len(addresses))

	for i, v := range addresses {
		municipality, municipalityErr := locationRepo.GetMunicipalityById(v.Municipality)

		if municipalityErr != nil {
			return []ContactAddress{}, municipalityErr
		}

		address, addressErr := CreateAddress(municipality, v.Address)

		if addressErr != nil {
			return []ContactAddress{}, addressErr
		}

		addressList[i] = address
	}

	return addressList, nil
}

// Contacts Value Object
type Contacts struct {
	Emails    []ContactEmail   `json:"emails"`
	Phones    []ContactPhone   `json:"phones"`
	Addresses []ContactAddress `json:"addresses"`
}

// Contacts Value Object Factory
func CreateContacts(emails []ContactEmail, phones []ContactPhone, addresses []ContactAddress) Contacts {
	return Contacts{
		Emails:    emails,
		Phones:    phones,
		Addresses: addresses,
	}
}

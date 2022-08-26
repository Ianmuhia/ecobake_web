// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
	GetID() string
}

type AccountDelete struct {
	Errors []*AccountError `json:"errors"`
	User   *User           `json:"user"`
}

type AccountError struct {
	Field   *string          `json:"field"`
	Message *string          `json:"message"`
	Code    AccountErrorCode `json:"code"`
}

type AccountInput struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type AccountRegister struct {
	RequiresConfirmation *bool           `json:"requiresConfirmation"`
	Errors               []*AccountError `json:"errors"`
	User                 *User           `json:"user"`
}

type AccountRegisterInput struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	RedirectURL *string `json:"redirectUrl"`
	Channel     *string `json:"channel"`
}

type AccountRegisterResponse struct {
	Errors []*AccountError `json:"errors"`
	User   *User           `json:"user"`
}

type AccountRequestDeletion struct {
	Errors []*AccountError `json:"errors"`
}

type AccountUpdate struct {
	Errors []*AccountError `json:"errors"`
	User   *User           `json:"user"`
}

type Address struct {
	ID                       string  `json:"id"`
	FirstName                string  `json:"firstName"`
	LastName                 string  `json:"lastName"`
	CompanyName              string  `json:"companyName"`
	StreetAddress1           string  `json:"streetAddress1"`
	StreetAddress2           string  `json:"streetAddress2"`
	City                     string  `json:"city"`
	CityArea                 string  `json:"cityArea"`
	PostalCode               string  `json:"postalCode"`
	CountryArea              string  `json:"countryArea"`
	Phone                    *string `json:"phone"`
	IsDefaultShippingAddress *bool   `json:"isDefaultShippingAddress"`
	IsDefaultBillingAddress  *bool   `json:"isDefaultBillingAddress"`
}

func (Address) IsNode()            {}
func (this Address) GetID() string { return this.ID }

type Categories struct {
	Categories []*Category           `json:"categories"`
	Errors     []ListEntityErrorCode `json:"errors"`
}

type CheckoutError struct {
	Field    *string           `json:"field"`
	Message  *string           `json:"message"`
	Code     CheckoutErrorCode `json:"code"`
	Variants []string          `json:"variants"`
	Lines    []string          `json:"lines"`
}

type ConfirmAccount struct {
	User   *User           `json:"user"`
	Errors []*AccountError `json:"errors"`
}

type ConfirmEmailChange struct {
	User   *User           `json:"user"`
	Errors []*AccountError `json:"errors"`
}

type CreateCategory struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CreateToken struct {
	Token        *string         `json:"token"`
	RefreshToken *string         `json:"refreshToken"`
	CsrfToken    *string         `json:"csrfToken"`
	User         *User           `json:"user"`
	Errors       []*AccountError `json:"errors"`
}

type DeactivateAllUserTokens struct {
	Errors []*AccountError `json:"errors"`
}

type File struct {
	URL         string  `json:"url"`
	ContentType *string `json:"contentType"`
}

type FileUpload struct {
	UploadedFile *File          `json:"uploadedFile"`
	Errors       []*UploadError `json:"errors"`
}

type Image struct {
	URL string  `json:"url"`
	Alt *string `json:"alt"`
}

type LoginResp struct {
	User    *User   `json:"user"`
	Refresh *string `json:"refresh"`
	Access  *string `json:"access"`
}

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type NewUser struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type PasswordChange struct {
	User          *User           `json:"user"`
	AccountErrors []*AccountError `json:"accountErrors"`
	Errors        []*AccountError `json:"errors"`
}

type RefreshToken struct {
	Token  *string         `json:"token"`
	User   *User           `json:"user"`
	Errors []*AccountError `json:"errors"`
}

type RequestEmailChange struct {
	User   *User           `json:"user"`
	Errors []*AccountError `json:"errors"`
}

type RequestPasswordReset struct {
	Errors []*AccountError `json:"errors"`
}

type SetPassword struct {
	Token        *string         `json:"token"`
	RefreshToken *string         `json:"refreshToken"`
	CsrfToken    *string         `json:"csrfToken"`
	User         *User           `json:"user"`
	Errors       []*AccountError `json:"errors"`
}

type UploadError struct {
	Field   *string         `json:"field"`
	Message *string         `json:"message"`
	Code    UploadErrorCode `json:"code"`
}

type UserAvatarDelete struct {
	User   *User           `json:"user"`
	Errors []*AccountError `json:"errors"`
}

type UserAvatarUpdate struct {
	User   *User           `json:"user"`
	Errors []*AccountError `json:"errors"`
}

type Users struct {
	Users  []*User               `json:"users"`
	Errors []ListEntityErrorCode `json:"errors"`
}

type VerifyToken struct {
	User    *User           `json:"user"`
	IsValid bool            `json:"isValid"`
	Payload *string         `json:"payload"`
	Errors  []*AccountError `json:"errors"`
}

type AccountErrorCode string

const (
	AccountErrorCodeActivateOwnAccount          AccountErrorCode = "ACTIVATE_OWN_ACCOUNT"
	AccountErrorCodeActivateSuperuserAccount    AccountErrorCode = "ACTIVATE_SUPERUSER_ACCOUNT"
	AccountErrorCodeDuplicatedInputItem         AccountErrorCode = "DUPLICATED_INPUT_ITEM"
	AccountErrorCodeDeactivateOwnAccount        AccountErrorCode = "DEACTIVATE_OWN_ACCOUNT"
	AccountErrorCodeDeactivateSuperuserAccount  AccountErrorCode = "DEACTIVATE_SUPERUSER_ACCOUNT"
	AccountErrorCodeDeleteNonStaffUser          AccountErrorCode = "DELETE_NON_STAFF_USER"
	AccountErrorCodeDeleteOwnAccount            AccountErrorCode = "DELETE_OWN_ACCOUNT"
	AccountErrorCodeDeleteStaffAccount          AccountErrorCode = "DELETE_STAFF_ACCOUNT"
	AccountErrorCodeDeleteSuperuserAccount      AccountErrorCode = "DELETE_SUPERUSER_ACCOUNT"
	AccountErrorCodeGraphqlError                AccountErrorCode = "GRAPHQL_ERROR"
	AccountErrorCodeInactive                    AccountErrorCode = "INACTIVE"
	AccountErrorCodeInvalid                     AccountErrorCode = "INVALID"
	AccountErrorCodeInvalidPassword             AccountErrorCode = "INVALID_PASSWORD"
	AccountErrorCodeLeftNotManageablePermission AccountErrorCode = "LEFT_NOT_MANAGEABLE_PERMISSION"
	AccountErrorCodeInvalidCredentials          AccountErrorCode = "INVALID_CREDENTIALS"
	AccountErrorCodeNotFound                    AccountErrorCode = "NOT_FOUND"
	AccountErrorCodeOutOfScopeUser              AccountErrorCode = "OUT_OF_SCOPE_USER"
	AccountErrorCodeOutOfScopeGroup             AccountErrorCode = "OUT_OF_SCOPE_GROUP"
	AccountErrorCodeOutOfScopePermission        AccountErrorCode = "OUT_OF_SCOPE_PERMISSION"
	AccountErrorCodePasswordEntirelyNumeric     AccountErrorCode = "PASSWORD_ENTIRELY_NUMERIC"
	AccountErrorCodePasswordTooCommon           AccountErrorCode = "PASSWORD_TOO_COMMON"
	AccountErrorCodePasswordTooShort            AccountErrorCode = "PASSWORD_TOO_SHORT"
	AccountErrorCodePasswordTooSimilar          AccountErrorCode = "PASSWORD_TOO_SIMILAR"
	AccountErrorCodeRequired                    AccountErrorCode = "REQUIRED"
	AccountErrorCodeUnique                      AccountErrorCode = "UNIQUE"
	AccountErrorCodeJwtSignatureExpired         AccountErrorCode = "JWT_SIGNATURE_EXPIRED"
	AccountErrorCodeJwtInvalidToken             AccountErrorCode = "JWT_INVALID_TOKEN"
	AccountErrorCodeJwtDecodeError              AccountErrorCode = "JWT_DECODE_ERROR"
	AccountErrorCodeJwtMissingToken             AccountErrorCode = "JWT_MISSING_TOKEN"
	AccountErrorCodeJwtInvalidCsrfToken         AccountErrorCode = "JWT_INVALID_CSRF_TOKEN"
	AccountErrorCodeChannelInactive             AccountErrorCode = "CHANNEL_INACTIVE"
	AccountErrorCodeMissingChannelSlug          AccountErrorCode = "MISSING_CHANNEL_SLUG"
	AccountErrorCodeAccountNotConfirmed         AccountErrorCode = "ACCOUNT_NOT_CONFIRMED"
)

var AllAccountErrorCode = []AccountErrorCode{
	AccountErrorCodeActivateOwnAccount,
	AccountErrorCodeActivateSuperuserAccount,
	AccountErrorCodeDuplicatedInputItem,
	AccountErrorCodeDeactivateOwnAccount,
	AccountErrorCodeDeactivateSuperuserAccount,
	AccountErrorCodeDeleteNonStaffUser,
	AccountErrorCodeDeleteOwnAccount,
	AccountErrorCodeDeleteStaffAccount,
	AccountErrorCodeDeleteSuperuserAccount,
	AccountErrorCodeGraphqlError,
	AccountErrorCodeInactive,
	AccountErrorCodeInvalid,
	AccountErrorCodeInvalidPassword,
	AccountErrorCodeLeftNotManageablePermission,
	AccountErrorCodeInvalidCredentials,
	AccountErrorCodeNotFound,
	AccountErrorCodeOutOfScopeUser,
	AccountErrorCodeOutOfScopeGroup,
	AccountErrorCodeOutOfScopePermission,
	AccountErrorCodePasswordEntirelyNumeric,
	AccountErrorCodePasswordTooCommon,
	AccountErrorCodePasswordTooShort,
	AccountErrorCodePasswordTooSimilar,
	AccountErrorCodeRequired,
	AccountErrorCodeUnique,
	AccountErrorCodeJwtSignatureExpired,
	AccountErrorCodeJwtInvalidToken,
	AccountErrorCodeJwtDecodeError,
	AccountErrorCodeJwtMissingToken,
	AccountErrorCodeJwtInvalidCsrfToken,
	AccountErrorCodeChannelInactive,
	AccountErrorCodeMissingChannelSlug,
	AccountErrorCodeAccountNotConfirmed,
}

func (e AccountErrorCode) IsValid() bool {
	switch e {
	case AccountErrorCodeActivateOwnAccount, AccountErrorCodeActivateSuperuserAccount, AccountErrorCodeDuplicatedInputItem, AccountErrorCodeDeactivateOwnAccount, AccountErrorCodeDeactivateSuperuserAccount, AccountErrorCodeDeleteNonStaffUser, AccountErrorCodeDeleteOwnAccount, AccountErrorCodeDeleteStaffAccount, AccountErrorCodeDeleteSuperuserAccount, AccountErrorCodeGraphqlError, AccountErrorCodeInactive, AccountErrorCodeInvalid, AccountErrorCodeInvalidPassword, AccountErrorCodeLeftNotManageablePermission, AccountErrorCodeInvalidCredentials, AccountErrorCodeNotFound, AccountErrorCodeOutOfScopeUser, AccountErrorCodeOutOfScopeGroup, AccountErrorCodeOutOfScopePermission, AccountErrorCodePasswordEntirelyNumeric, AccountErrorCodePasswordTooCommon, AccountErrorCodePasswordTooShort, AccountErrorCodePasswordTooSimilar, AccountErrorCodeRequired, AccountErrorCodeUnique, AccountErrorCodeJwtSignatureExpired, AccountErrorCodeJwtInvalidToken, AccountErrorCodeJwtDecodeError, AccountErrorCodeJwtMissingToken, AccountErrorCodeJwtInvalidCsrfToken, AccountErrorCodeChannelInactive, AccountErrorCodeMissingChannelSlug, AccountErrorCodeAccountNotConfirmed:
		return true
	}
	return false
}

func (e AccountErrorCode) String() string {
	return string(e)
}

func (e *AccountErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountErrorCode", str)
	}
	return nil
}

func (e AccountErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CheckoutErrorCode string

const (
	CheckoutErrorCodeBillingAddressNotSet          CheckoutErrorCode = "BILLING_ADDRESS_NOT_SET"
	CheckoutErrorCodeCheckoutNotFullyPaid          CheckoutErrorCode = "CHECKOUT_NOT_FULLY_PAID"
	CheckoutErrorCodeProductNotPublished           CheckoutErrorCode = "PRODUCT_NOT_PUBLISHED"
	CheckoutErrorCodeProductUnavailableForPurchase CheckoutErrorCode = "PRODUCT_UNAVAILABLE_FOR_PURCHASE"
	CheckoutErrorCodeInsufficientStock             CheckoutErrorCode = "INSUFFICIENT_STOCK"
	CheckoutErrorCodeInvalid                       CheckoutErrorCode = "INVALID"
	CheckoutErrorCodeInvalidShippingMethod         CheckoutErrorCode = "INVALID_SHIPPING_METHOD"
	CheckoutErrorCodeNotFound                      CheckoutErrorCode = "NOT_FOUND"
	CheckoutErrorCodePaymentError                  CheckoutErrorCode = "PAYMENT_ERROR"
	CheckoutErrorCodeQuantityGreaterThanLimit      CheckoutErrorCode = "QUANTITY_GREATER_THAN_LIMIT"
	CheckoutErrorCodeRequired                      CheckoutErrorCode = "REQUIRED"
	CheckoutErrorCodeShippingAddressNotSet         CheckoutErrorCode = "SHIPPING_ADDRESS_NOT_SET"
	CheckoutErrorCodeShippingMethodNotApplicable   CheckoutErrorCode = "SHIPPING_METHOD_NOT_APPLICABLE"
	CheckoutErrorCodeDeliveryMethodNotApplicable   CheckoutErrorCode = "DELIVERY_METHOD_NOT_APPLICABLE"
	CheckoutErrorCodeShippingMethodNotSet          CheckoutErrorCode = "SHIPPING_METHOD_NOT_SET"
	CheckoutErrorCodeShippingNotRequired           CheckoutErrorCode = "SHIPPING_NOT_REQUIRED"
	CheckoutErrorCodeTaxError                      CheckoutErrorCode = "TAX_ERROR"
	CheckoutErrorCodeUnique                        CheckoutErrorCode = "UNIQUE"
	CheckoutErrorCodeVoucherNotApplicable          CheckoutErrorCode = "VOUCHER_NOT_APPLICABLE"
	CheckoutErrorCodeGiftCardNotApplicable         CheckoutErrorCode = "GIFT_CARD_NOT_APPLICABLE"
	CheckoutErrorCodeZeroQuantity                  CheckoutErrorCode = "ZERO_QUANTITY"
	CheckoutErrorCodeMissingChannelSlug            CheckoutErrorCode = "MISSING_CHANNEL_SLUG"
	CheckoutErrorCodeChannelInactive               CheckoutErrorCode = "CHANNEL_INACTIVE"
	CheckoutErrorCodeUnavailableVariantInChannel   CheckoutErrorCode = "UNAVAILABLE_VARIANT_IN_CHANNEL"
	CheckoutErrorCodeEmailNotSet                   CheckoutErrorCode = "EMAIL_NOT_SET"
	CheckoutErrorCodeNoLines                       CheckoutErrorCode = "NO_LINES"
)

var AllCheckoutErrorCode = []CheckoutErrorCode{
	CheckoutErrorCodeBillingAddressNotSet,
	CheckoutErrorCodeCheckoutNotFullyPaid,
	CheckoutErrorCodeProductNotPublished,
	CheckoutErrorCodeProductUnavailableForPurchase,
	CheckoutErrorCodeInsufficientStock,
	CheckoutErrorCodeInvalid,
	CheckoutErrorCodeInvalidShippingMethod,
	CheckoutErrorCodeNotFound,
	CheckoutErrorCodePaymentError,
	CheckoutErrorCodeQuantityGreaterThanLimit,
	CheckoutErrorCodeRequired,
	CheckoutErrorCodeShippingAddressNotSet,
	CheckoutErrorCodeShippingMethodNotApplicable,
	CheckoutErrorCodeDeliveryMethodNotApplicable,
	CheckoutErrorCodeShippingMethodNotSet,
	CheckoutErrorCodeShippingNotRequired,
	CheckoutErrorCodeTaxError,
	CheckoutErrorCodeUnique,
	CheckoutErrorCodeVoucherNotApplicable,
	CheckoutErrorCodeGiftCardNotApplicable,
	CheckoutErrorCodeZeroQuantity,
	CheckoutErrorCodeMissingChannelSlug,
	CheckoutErrorCodeChannelInactive,
	CheckoutErrorCodeUnavailableVariantInChannel,
	CheckoutErrorCodeEmailNotSet,
	CheckoutErrorCodeNoLines,
}

func (e CheckoutErrorCode) IsValid() bool {
	switch e {
	case CheckoutErrorCodeBillingAddressNotSet, CheckoutErrorCodeCheckoutNotFullyPaid, CheckoutErrorCodeProductNotPublished, CheckoutErrorCodeProductUnavailableForPurchase, CheckoutErrorCodeInsufficientStock, CheckoutErrorCodeInvalid, CheckoutErrorCodeInvalidShippingMethod, CheckoutErrorCodeNotFound, CheckoutErrorCodePaymentError, CheckoutErrorCodeQuantityGreaterThanLimit, CheckoutErrorCodeRequired, CheckoutErrorCodeShippingAddressNotSet, CheckoutErrorCodeShippingMethodNotApplicable, CheckoutErrorCodeDeliveryMethodNotApplicable, CheckoutErrorCodeShippingMethodNotSet, CheckoutErrorCodeShippingNotRequired, CheckoutErrorCodeTaxError, CheckoutErrorCodeUnique, CheckoutErrorCodeVoucherNotApplicable, CheckoutErrorCodeGiftCardNotApplicable, CheckoutErrorCodeZeroQuantity, CheckoutErrorCodeMissingChannelSlug, CheckoutErrorCodeChannelInactive, CheckoutErrorCodeUnavailableVariantInChannel, CheckoutErrorCodeEmailNotSet, CheckoutErrorCodeNoLines:
		return true
	}
	return false
}

func (e CheckoutErrorCode) String() string {
	return string(e)
}

func (e *CheckoutErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckoutErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckoutErrorCode", str)
	}
	return nil
}

func (e CheckoutErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ListEntityErrorCode string

const (
	ListEntityErrorCodeGraphqlError ListEntityErrorCode = "GRAPHQL_ERROR"
	ListEntityErrorCodeNotFound     ListEntityErrorCode = "NOT_FOUND"
)

var AllListEntityErrorCode = []ListEntityErrorCode{
	ListEntityErrorCodeGraphqlError,
	ListEntityErrorCodeNotFound,
}

func (e ListEntityErrorCode) IsValid() bool {
	switch e {
	case ListEntityErrorCodeGraphqlError, ListEntityErrorCodeNotFound:
		return true
	}
	return false
}

func (e ListEntityErrorCode) String() string {
	return string(e)
}

func (e *ListEntityErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ListEntityErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ListEntityErrorCode", str)
	}
	return nil
}

func (e ListEntityErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UploadErrorCode string

const (
	UploadErrorCodeGraphqlError UploadErrorCode = "GRAPHQL_ERROR"
)

var AllUploadErrorCode = []UploadErrorCode{
	UploadErrorCodeGraphqlError,
}

func (e UploadErrorCode) IsValid() bool {
	switch e {
	case UploadErrorCodeGraphqlError:
		return true
	}
	return false
}

func (e UploadErrorCode) String() string {
	return string(e)
}

func (e *UploadErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UploadErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UploadErrorCode", str)
	}
	return nil
}

func (e UploadErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
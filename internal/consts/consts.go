package consts

const (
	FailedValidation                = 1001
	RegexAlphanumericSimbols string = `^[a-zA-Z]?[a-zA-Z0-9\-\.\,\~\@\!\#\%\&\^\*\$(\)\/\s]+$`
	RegexAlphanumericPetik   string = `^[a-zA-Z0-9'\s]*[a-zA-Z][a-zA-Z0-9'\s]*$`

	AlphanumericSimbols      string = "Kolom harus dengan huruf, angka dan spasi."
	AlphanumericPetik        string = "Kolom hanya bisa dengan huruf, angka, spasi, dan petik serta harus minimal 1 huruf"
	AlphanumericSimbolsLogin string = "Password harus dengan huruf, angka, simbols, dan minimal 12 karakter"
	Digit                    string = "Kolom harus dengan angka."
	IsEmail                  string = "kolom harus sesuai kaidah email"
	MaxMinChar17             string = "Panjang data harus 17 character."
	RequiredField            string = "Kolom harus di isi"

	// config const redis db
	Token string = "token"

	// config response message
	SuccessRegister string = "successfully register"
)

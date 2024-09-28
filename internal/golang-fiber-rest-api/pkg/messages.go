package pkg

// Messages, uygulamada kullanılan sabit mesajları tutar.
type Messages struct {
	UserNotFoundByID      string
	UserNotFoundByEmail   string
	NoUsersInDatabase     string
	NotFoundError         string
	NotFoundUsers         string
	UserIdCannotBeEmpty   string
	EmailCannotBeEmpty    string
	UserCreatedSuccess    string
	JsonBindingError      string
	SaveStarted           string
	GetUserByIdStarted    string
	GetUserStarted        string
	GetUserByEmailStarted string
	UserAlreadyExists     string
}

var Msg = Messages{
	UserNotFoundByID:      "Verilen kimliğe sahip kullanıcı bulunamadı: %s\n",
	UserNotFoundByEmail:   "Verilen e-posta adresine sahip kullanıcı bulunamadı: %s\n",
	NoUsersInDatabase:     "Veritabanında kullanıcı bulunamadı\n",
	NotFoundError:         "not found error",
	NotFoundUsers:         "not found users",
	UserIdCannotBeEmpty:   "userId boş olamaz",
	EmailCannotBeEmpty:    "email boş olamaz",
	UserCreatedSuccess:    "Kullanıcı Başarıyla Oluşturuldu",
	JsonBindingError:      "JSON bağlama sırasında bir hata oluştu - ERROR: %v\n",
	SaveStarted:           "userController.Save STARTED with request: %#v\n",
	GetUserByIdStarted:    "userController.GetUserById STARTED with userId: %s\n",
	GetUserStarted:        "userController.GetUser INFO - Başladı \n",
	GetUserByEmailStarted: "userController.GetUserByEmail STARTED with email: %s\n",
	UserAlreadyExists:     "User Already Exist for given email: %s",
}

package data

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/salar-hafezi/go-taskmanager/models"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) error {
	objId := bson.NewObjectId()
	user.Id = objId
	hPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hPass
	// clear the real incoming pass
	user.Password = ""
	err = r.C.Insert(&user)
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	// validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}

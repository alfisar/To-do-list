package repository

import (
	"testing"
	"todolist/config"
	"todolist/domain"

	"github.com/stretchr/testify/require"
)

var (
	poolData = config.AppConfig
	repo     = NewUserRepo()
)

func TestCreate(t *testing.T) {
	data := domain.User{
		Name:     "Alfisar",
		Email:    "alfisartest@gmail.com",
		Username: "alfisar",
		Password: "coba doang",
		Photo:    "lululul.jpeg",
	}
	err := repo.Create(poolData.DBSql, data)
	require.Nil(t, err)
}

func TestCreateFail(t *testing.T) {
	data := domain.User{
		Name:     "Alfisar",
		Email:    "alfisartest@gmail.com",
		Username: "alfisar",
		Password: "coba doang",
		Photo:    "lululul.jpeg",
	}
	err := repo.Create(poolData.DBSql, data)
	require.NotNil(t, err)
}

func TestCreateFailDB(t *testing.T) {
	data := domain.User{
		Name:     "Alfisar",
		Email:    "alfisartest@gmail.com",
		Username: "alfisar",
		Password: "coba doang",
		Photo:    "lululul.jpeg",
	}
	err := repo.Create(nil, data)
	require.NotNil(t, err)
}

func TestGet(t *testing.T) {
	where := map[string]any{
		"email": "alfisartest@gmail.com",
	}
	data, err := repo.Get(poolData.DBSql, where)
	require.Nil(t, err)
	require.NotEqual(t, domain.User{}, data)
}

func TestGetFail(t *testing.T) {
	where := map[string]any{
		"email": "alfisartest1@gmail.com",
	}
	data, err := repo.Get(poolData.DBSql, where)
	require.NotNil(t, err)
	require.Equal(t, domain.User{}, data)
}

func TestGetFailDB(t *testing.T) {
	where := map[string]any{}
	data, err := repo.Get(nil, where)
	require.NotNil(t, err)
	require.Equal(t, domain.User{}, data)
}

func TestUpdate(t *testing.T) {
	where := map[string]any{
		"email": "alfisartest@gmail.com",
	}
	update := map[string]any{
		"email": "alfisar@gmail.com",
	}
	err := repo.Update(poolData.DBSql, update, where)
	require.Nil(t, err)

}

func TestUpdateFail(t *testing.T) {
	where := map[string]any{
		"email": "alfisartest1@gmail.com",
	}

	update := map[string]any{
		"email": "alfisar@gmail.com",
	}
	err := repo.Update(poolData.DBSql, update, where)
	require.NotNil(t, err)
}

func TestUpdateFailWhere(t *testing.T) {
	where := map[string]any{}
	update := map[string]any{
		"email": "alfisar@gmail.com",
	}
	err := repo.Update(poolData.DBSql, update, where)
	require.NotNil(t, err)

}

func TestUpdateFailDB(t *testing.T) {
	where := map[string]any{}
	update := map[string]any{
		"email": "alfisar@gmail.com",
	}
	err := repo.Update(nil, update, where)
	require.NotNil(t, err)

}

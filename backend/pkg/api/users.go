package api

import (
	"fmt"
	"net/http"

	types "backend/pkg/types"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

  _, req, err := GetBodyData[types.LoginRequest](r)
  if err != nil {
    return err
	}

	acc, err := s.store.GetUserByUsername(req.Username)
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJWT(acc)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
    ID: acc.ID,
		Token:  token,
		Username: acc.Username,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleSignUp(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
    _, req, err := GetBodyData[types.CreateUserRequest](r)
    if err != nil {
      return err
    }

    user, err := types.NewUser(req.Username, req.Password)
    if err != nil {
      return err
    }
    
    id, err := s.store.CreateUser(user)
    if err != nil {
      return err
    }

    user.ID = id

    return WriteJSON(w, http.StatusOK, user)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	username := getID(r, "username")

	if err := s.store.DeleteUser(username); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": username})
}

func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	id := getID(r, "id")
	
	user, err := s.store.GetUserByID(id)
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUserByUsername(w http.ResponseWriter, r *http.Request) error {
	username := getID(r, "username")
	
	user, err := s.store.GetUserByUsername(username)
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleUpdateUser(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {	
	_, req, err := GetBodyData[types.UpdateUserRequest](r)
	if err != nil {
		return err
	}

	if err = s.store.UpdateUser(req); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, req)
}

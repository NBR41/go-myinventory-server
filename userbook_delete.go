package main

import (
	"net/http"
)

func deleteUserBook(w http.ResponseWriter, r *http.Request) {
	ids, err := getURLIDs(r, urlUserIDField, urlBookIDField)
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.close() }()

	err = m.DeleteUserBook(ids[0], ids[1])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, nil)
}

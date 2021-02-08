package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Outerwolf/bookstore_items-api/domain/items"
	"github.com/Outerwolf/bookstore_items-api/services"
	"github.com/Outerwolf/bookstore_items-api/utils/http_utils"
	"github.com/Outerwolf/bookstore_oauth-go/oauth"
	"github.com/Outerwolf/bookstore_utils-go/rest_errors"
)

var (
	ItemsCoontroller itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AutheneticateRequest(r); err != nil {
		//http_utils.RespondError(w, *err)
		return
	}

	sellerId := oauth.GetCallerId(r)

	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("Invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil{
		http_utils.RespondError(w,err)
	}
	http_utils.RespondError(w, http.StatusOK, item)
	
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr :=  rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query. queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, items)
}

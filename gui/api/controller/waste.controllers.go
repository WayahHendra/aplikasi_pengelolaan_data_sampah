package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"trash-app/core"
	services "trash-app/gui/api/service"
	logger "trash-app/utils"
)

type TrashControllers struct {
	service *services.TrashServices
}

func NewTrashControllers(service *services.TrashServices) *TrashControllers {
	return &TrashControllers{
		service: service,
	}
}

func (c *TrashControllers) CreateWaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Print(logger.Types["error"], "Method not allowed on /waste/add request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	var newTrash core.Waste
	err := json.NewDecoder(r.Body).Decode(&newTrash)
	if err != nil {
		log.Print(logger.Types["warning"], "Invalid JSON format on /waste/add request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": err.Error(),
		})

		return
	}

	err = c.service.CreateWaste(newTrash)
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/add request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	log.Print(logger.Types["success"], "Trash data added successfully on /waste request")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data added successfully",
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) GetAllWaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Print(logger.Types["error"], "Method not allowed on /waste/get/all request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	wasteData, err := c.service.GetAllWaste()
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/get/all request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	log.Print(logger.Types["success"], "Trash data retrieved successfully on /waste/get/all request")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data retrieved successfully",
		"data":    wasteData,
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) GetWasteByDate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	date := parts[3]
	
	if r.Method != http.MethodGet {
		log.Print(logger.Types["error"], "Method not allowed on /waste/get/"+ date +" request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	if len(parts) < 4 {
		log.Print(logger.Types["warning"], "Invalid format date on /waste/get/"+ date +" request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": "Invalid format date or data with date does not exist",
		})

		return
	}

	wasteData, err := c.service.GetWasteByDate(date)
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/get/"+ date +" request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		
		return
	}

	log.Print(logger.Types["success"], "Trash data retrieved successfully on /waste/get/"+ date +" request")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data with date ("+ date +") retrieved successfully",
		"data":    wasteData,
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) UpdateWasteById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id := parts[3]
	
	if r.Method != http.MethodPut {
		log.Print(logger.Types["error"], "Method not allowed on /waste/update/"+ id +" request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	var updatedWaste core.Waste
	err := json.NewDecoder(r.Body).Decode(&updatedWaste)
	if err != nil {
		log.Print(logger.Types["warning"], "Invalid JSON format on /waste/update/"+ id +" request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": err.Error(),
		})

		return
	}

	if len(parts) < 4 {
		log.Print(logger.Types["warning"], "Invalid id or data with id does not exist on /waste/update/"+ id +" request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": "Invalid id or data with id does not exist",
		})

		return
	}

	wasteData, err := c.service.UpdateWasteById(id, updatedWaste)
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/update/"+ id +" request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	log.Print(logger.Types["success"], "Trash data successfully updated on /waste/update/"+ id +" request")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data with id ("+ id +") successfully updated",
		"data":    wasteData,
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) DeleteWasteById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id := parts[3]

	if r.Method != http.MethodDelete {
		log.Print(logger.Types["error"], "Method not allowed on /waste/delete/"+ id +" request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	if len(parts) < 4 {
		log.Print(logger.Types["warning"], "Invalid id or data with id does not exist on /waste/delete/"+ id +" request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": "Invalid id or data with id does not exist",
		})

		return
	}

	err := c.service.DeleteWasteById(id)
    if err != nil {
        log.Print(logger.Types["error"], err.Error()+" on /waste/delete/"+ id +" request")
        w.WriteHeader(http.StatusInternalServerError)

        json.NewEncoder(w).Encode(map[string]string{
            "status":  "error",
            "message": err.Error(),
        })

        return
    }

	log.Print(logger.Types["success"], "Trash data successfully deleted on /waste/update/"+ id +" request")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data with id ("+ id +") successfully deleted",
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}
package records

import (
	"challengeFinal/auth"
	"challengeFinal/database"
	"challengeFinal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRecordsHandler(c *gin.Context) {
	var APIRecords []models.APIRecord
	database.CreateDbConnection()
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userConfig, err := auth.GetUserConfig(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user config"})
		return
	}

	if userConfig == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No display configuration found for user"})
		return
	}

	

	filters, ok := userConfig["filters"].(map[string]interface{})
	if !ok {
		filters = make(map[string]interface{})
	}
	pageParam, ok := c.GetQuery("page")
	if !ok || pageParam == "" {
		if userPage, exists := userConfig["page"]; exists {
			pageParam = fmt.Sprintf("%v", userPage)
		} else {
			pageParam = "1"
		}
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 { // Asegurarse de que sea un número válido
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset := (page - 1) * limit
	sortBy := c.DefaultQuery("sort_by", fmt.Sprintf("%v", userConfig["sort_by"]))
	order := c.DefaultQuery("order", fmt.Sprintf("%v", userConfig["order"]))

	allowedSortFields := map[string]bool{
		"age":            true,
		"education":      true,
		"marital_status": true,
		"occupation":     true,
		"income":         true,
		"id":             true,
	}
	if !allowedSortFields[sortBy] {
		sortBy = "id"
	}
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	query := database.DBConn.Model(&models.Record{}).
		Select("id", "age", "education", "marital_status", "occupation", "income", "hours_per_week").
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s %s", sortBy, order))

	queryParams := map[string]string{
		"education":      "education = ?",
		"marital_status": "marital_status = ?",
		"occupation":     "occupation = ?",
		"age":            "age = ?",
		"income":         "income = ?",
	}

	for param, condition := range queryParams {
		value := c.DefaultQuery(param, "")
		if value == "" {
			if savedValue, exists := filters[param]; exists && savedValue != "" {
				value = fmt.Sprintf("%v", savedValue)
			}
		}

		if value != "" {
			query = query.Where(condition, value)
		}
	}
	if err := query.Find(&APIRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	countQuery := database.DBConn.Model(&models.Record{})
	for param, condition := range queryParams {
		value := c.DefaultQuery(param, "")
		if value == "" {
			if savedValue, exists := filters[param]; exists && savedValue != "" {
				value = fmt.Sprintf("%v", savedValue)
			}
		}

		if value != "" {
			countQuery = countQuery.Where(condition, value)
		}
	}

	var totalRecords int64
	if err := countQuery.Count(&totalRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (int(totalRecords) + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"data":         APIRecords,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}

func SaveUserConfig(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var config map[string]interface{}
	if err := c.BindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid config format"})
		return
	}

	if err := auth.UpdateUserDisplayConfig(userID.(uint), config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save config"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Configuration saved successfully"})
}

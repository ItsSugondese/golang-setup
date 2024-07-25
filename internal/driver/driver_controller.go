package driver

import (
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// post /driver/create
func CreateDriver(c *gin.Context) {
	var driverDto DriverDto
	if err := c.ShouldBind(&driverDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}
	driverDto.Image = file

	// calling service

	driver, err := CreateDriverService(driverDto)
	if err != nil {

		if strings.Contains(err.Error(), "email") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "email already exists", "type": "error"})
			return
		}
		if strings.Contains(err.Error(), "phone") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "phone Number already exists", "type": "error"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "create success", "data": driver, "type": "success"})
}

// put /driver/update
func UpdateDriver(c *gin.Context) {
	// Extracting id from URL
	id := c.Param("id")

	// Handle file upload (optional)
	var file *multipart.FileHeader
	var err error
	file, err = c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error uploading file", "type": "error"})
		return
	}

	// Binding payload with Go struct
	var driverDto DriverDto
	if err := c.ShouldBind(&driverDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	driverDto.Image = file

	// Update driver service
	driver, err := UpdateDriverService(driverDto, id)
	if err != nil {

		if strings.Contains(err.Error(), "email") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "email problem", "type": "error"})
			return
		}
		if strings.Contains(err.Error(), "phone") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "phone Number problem", "type": "error"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "update success", "data": driver, "type": "success"})
}

// get /driver/get-all
func GetDrivers(c *gin.Context) {
	drivers, err := GetAlldriversService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drivers})
}

// get /driver/get/:id
func GetDriverById(c *gin.Context) {
	id := c.Param("id")
	//converting string to integer
	driverId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	driver, err := GetdriverByIdService(driverId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": driver})

}

// delete /driver/delete/:id
func DeleteDriver(c *gin.Context) {
	id := c.Param("id")
	//converting string to integer
	driverId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	err = DeleteDriverService(driverId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver deleted successfully", "type": "success"})

}

package pack

import (
	"net/http"
	"strconv"

	config "packs-per-order/config"
	backtracking "packs-per-order/lib/pack-strategies/backtracking"

	"github.com/gin-gonic/gin"
)

func HandlePackRequest(c *gin.Context) {
	// Get the order from query parameter
	orderStr := c.Query("order")
	order, err := strconv.Atoi(orderStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order value"})
		return
	}

	// Get pack sizes from configuration
	config, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	packSizes := config.PackSizes

	// Create a new backtracking strategy with the configured pack sizes
	strategy := backtracking.NewPackageStrategy(packSizes)

	// Find the ideal pack
	result := strategy.FindIdealPack(order)

	// Return the result
	c.JSON(http.StatusOK, result)
}

package seed

import (
	"fmt"
	"time"

	"github.com/asishshaji/thalir-backend/models"
	"gorm.io/gorm"
)

var products = []models.Product{
	{
		ID:        1,
		Name:      "Onion",
		Type:      "R",
		BuyPrice:  200.0,
		SellPrice: 300.0,
		Units:     100,
	},
	{
		ID:        2,
		Name:      "Potato",
		Type:      "R",
		BuyPrice:  100.0,
		SellPrice: 200.0,
		Units:     200,
	},
	{
		ID:        3,
		Name:      "Mushroom",
		Type:      "P",
		BuyPrice:  400.0,
		SellPrice: 500.0,
		Units:     600,
	},
	{
		ID:        4,
		Name:      "Carrot",
		Type:      "P",
		BuyPrice:  400.0,
		SellPrice: 500.0,
		Units:     600,
	},
	{
		ID:        5,
		Name:      "Eggplant",
		Type:      "P",
		BuyPrice:  400.0,
		SellPrice: 500.0,
		Units:     600,
	},
	{
		ID:        6,
		Name:      "Broccoli",
		Type:      "P",
		BuyPrice:  400.0,
		SellPrice: 500.0,
		Units:     600,
	},
}

var orders = []models.Order{
	{
		ID:          1,
		PhoneNumber: "9400367283",
		OrderItems: []models.OrderItem{
			{

				ProductId:    1,
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
			{
				ProductId:    2,
				Units:        5,
				Profit:       500,
				BuyingPrice:  100.0,
				SellingPrice: 200.0,
			},
		},
	},
	{
		ID:          2,
		PhoneNumber: "4400367283",
		OrderItems: []models.OrderItem{
			{
				ProductId:    1,
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
			{
				ProductId:    2,
				Units:        5,
				Profit:       500,
				BuyingPrice:  100.0,
				SellingPrice: 200.0,
			},
		},
	},
	{
		ID:          3,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 29, 1, 1, 1, 21, &time.Location{}),
		OrderItems: []models.OrderItem{
			{
				ProductId: 1,
				CreatedAt: time.Date(2021, 11, 29, 1, 1, 1, 21, &time.Location{}),

				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
			{
				ProductId:    2,
				CreatedAt:    time.Date(2021, 12, 1, 1, 1, 1, 21, &time.Location{}),
				Units:        5,
				Profit:       500,
				BuyingPrice:  100.0,
				SellingPrice: 200.0,
			},
		},
	},
	{
		ID:          4,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 28, 1, 1, 1, 21, &time.Location{}),
		OrderItems: []models.OrderItem{
			{
				ProductId:    1,
				CreatedAt:    time.Date(2021, 11, 28, 1, 1, 1, 21, &time.Location{}),
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          5,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 28, 1, 1, 1, 21, &time.Location{}),
		OrderItems: []models.OrderItem{
			{
				ProductId:    1,
				CreatedAt:    time.Date(2021, 11, 28, 1, 1, 1, 21, &time.Location{}),
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          6,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 27, 1, 1, 1, 21, &time.Location{}),
		OrderItems: []models.OrderItem{
			{
				ProductId:    1,
				CreatedAt:    time.Date(2021, 11, 27, 1, 1, 1, 21, &time.Location{}),
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          7,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 26, 1, 1, 1, 21, &time.Location{}),

		OrderItems: []models.OrderItem{
			{
				ProductId: 1,
				Units:     5,
				CreatedAt: time.Date(2021, 11, 26, 1, 1, 1, 21, &time.Location{}),

				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          8,
		PhoneNumber: "4400367283",
		OrderItems: []models.OrderItem{
			{
				ProductId:    1,
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          9,
		PhoneNumber: "4400367283",
		OrderItems: []models.OrderItem{
			{
				ProductId:    1,
				Units:        5,
				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          10,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 25, 1, 1, 1, 21, &time.Location{}),

		OrderItems: []models.OrderItem{
			{
				ProductId: 1,
				Units:     5,
				CreatedAt: time.Date(2021, 11, 25, 1, 1, 1, 21, &time.Location{}),

				Profit:       500,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
	{
		ID:          11,
		PhoneNumber: "4400367283",
		CreatedAt:   time.Date(2021, 11, 26, 1, 1, 1, 21, &time.Location{}),

		OrderItems: []models.OrderItem{
			{
				ProductId: 1,
				Units:     10,
				CreatedAt: time.Date(2021, 11, 26, 1, 1, 1, 21, &time.Location{}),

				Profit:       1000,
				BuyingPrice:  200.0,
				SellingPrice: 300.0,
			},
		},
	},
}

func Load(db *gorm.DB) {

	prod := models.Product{}
	order := models.Order{}
	orderDetail := models.OrderItem{}
	db.Migrator().DropTable(&prod, &order, &orderDetail)
	fmt.Println("Dropping tables")
	// if db.Migrator().HasTable(&prod) && db.Migrator().HasTable(&order) && db.Migrator().HasTable(&orderDetail) {
	// 	db.Migrator().DropTable(&prod, &order, &orderDetail)
	// }

	db.AutoMigrate(&prod, &order, &orderDetail)

	// err := db.Create(&products).Error
	// if err != nil {
	// 	log.Fatalln("Cannot seed products table", err)
	// }

	// err = db.Create(&orders).Error
	// if err != nil {
	// 	log.Fatalln("Cannot seed products table", err)
	// }
}

package repo

import (
	"context"
	"database/sql"
	"platform/config"
	"platform/logging"
)

type SQLRepository struct {
	config.Configuration
	logging.Logger
	Commands SQLCommands
	*sql.DB
	context.Context
}

type SQLCommands struct {
	Init,
	Seed,
	GetProduct,
	GetProducts,
	GetCategories,
	GetPage,
	GetPageCount,
	GetCategoryPage,
	GetCategoryPageCount,
	GetOrder,
	GetOrderLines,
	GetOrders,
	GetOrdersLines,
	SaveOrder,
	SaveOrderLine,
	UpdateOrder,
	SaveProduct,
	UpdateProduct,
	SaveCategory,
	UpdateCategory *sql.Stmt
}

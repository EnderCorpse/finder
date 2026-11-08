package database

import (
	"context"
	"time"

	"github.com/ComputerScienceHouse/finder/internal/logging"
	"github.com/ComputerScienceHouse/finder/internal/models"
)

func InitItems() error {
	if _, err := db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS items (
			item_id 		UUID PRIMARY KEY,
			location 		TEXT,
			name 			TEXT,
			quantity 		INT,
			description 	TEXT,
			current_owner 	TEXT,
			time_in 		TIMESTAMPTZ,
			time_out 		TIMESTAMPTZ
		)
	`); err != nil {
		return err
	}

	return nil
}

func GetAllItems(ctx context.Context) ([]models.ItemObject, error) {
	items := make([]models.ItemObject, 0)
	//_, err := db.Exec(ctx, `
	//INSERT INTO items (item_id) VALUES ('01a1ba49-50c1-709f-bd77-0932abfa5c64')
	//`)
	rows, err := db.Query(ctx, `
		SELECT * FROM items
		WHERE 1=1
		`)
	if err != nil {
		logging.Logger.Fatalf("GetAllItems: %s", err)
		return nil, err
	}
	logging.Logger.Info("GetAllItems: %s", rows)
	defer rows.Close()

	for rows.Next() {
		var itemID string
		var location string
		var name string
		var quantity int
		var description string
		var currentOwner string
		var timeIn time.Time
		var timeOut time.Time

		if err := rows.Scan(&itemID, &location, &name, &quantity, &description, &currentOwner, &timeIn, &timeOut); err != nil {
			return nil, err
		}
		logging.Logger.Infof("GetAllItems: %s %s", itemID, location)
		items = append(items, models.ItemObject{
			ID:          itemID,
			Location:    location,
			Name:        name,
			Quantity:    quantity,
			Description: description,
			Requester:   currentOwner,
			RequestDate: timeIn.Format(time.RFC3339),
			ReturnDate:  timeOut.Format(time.RFC3339),
		})
	}

	return items, nil
}

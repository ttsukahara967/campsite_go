package main

import (
    "database/sql"
)

func GetAllCampsites(db *sql.DB) ([]Campsite, error) {
    rows, err := db.Query("SELECT * FROM campsites")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var campsites []Campsite
    for rows.Next() {
        var c Campsite
        err := rows.Scan(&c.ID, &c.Name, &c.Address, &c.Description, &c.Facilities, &c.Price, &c.ImageURL, &c.Latitude, &c.Longitude, &c.CreatedAt)
        if err != nil {
            return nil, err
        }
        campsites = append(campsites, c)
    }
    return campsites, nil
}

func GetCampsiteByID(db *sql.DB, id int) (*Campsite, error) {
    var c Campsite
    err := db.QueryRow("SELECT * FROM campsites WHERE id = ?", id).Scan(
        &c.ID, &c.Name, &c.Address, &c.Description, &c.Facilities, &c.Price, &c.ImageURL, &c.Latitude, &c.Longitude, &c.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return &c, nil
}

// 他にCreate/Update/Deleteも実装できます（必要なら追加でサンプル書きます）


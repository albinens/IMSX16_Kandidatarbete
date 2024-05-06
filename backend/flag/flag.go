package flag

import (
	"example.com/m/v2/database"
)

type Flag struct {
	Name  string
	Value string
}

func (f Flag) IsSet() bool {
	if f.Value == "" {
		return false
	}
	return true
}

func (f Flag) Set() error {
	db := database.GetDB()

	_, err := db.Exec("INSERT INTO flags (name, value) VALUES ($1, $2)", f.Name, "true")
	if err != nil {
		return err
	}

	return nil
}

func Get(name string) (Flag, error) {
	db := database.GetDB()

	res, err := db.Queryx("SELECT value FROM flags WHERE name = $1", name)
	if err != nil {
		return Flag{}, err
	}
	defer res.Close()

	var value string
	for res.Next() {
		err = res.Scan(&value)
		if err != nil {
			return Flag{}, err
		}
	}

	return Flag{
		Name:  name,
		Value: value,
	}, nil
}

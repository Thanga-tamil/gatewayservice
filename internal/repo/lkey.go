package repo

import "gateway/internal/config"

func FetchLicenseKey() string {
	var licensekey string

	if err := config.Psql.Raw(`SELECT licensekey FROM chat_settings;`).Scan(&licensekey).Error; err != nil {
		panic(err)
	}

	return licensekey
}

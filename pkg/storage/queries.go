package storage

const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS urls (
    					id SERIAL PRIMARY KEY,
    					original_url TEXT NOT NULL,
    					short_url TEXT NOT NULL UNIQUE
													);`
	setNewValuesInDB = `INSERT INTO urls (original_url, short_url)
						VALUES ($1, $2)
						ON CONFLICT (short_url)
						DO UPDATE
						SET short_url = EXCLUDED.short_url;`
	getShortUrl = `SELECT original_url FROM urls WHERE short_url = $1 LIMIT 1`
)

package storage

// Все sql-запросы к БД
const (
	preCreateTableQuery = `drop table if exists urls;`
	createTableQuery    = `CREATE TABLE IF NOT EXISTS urls (
    					id SERIAL PRIMARY KEY,
    					original_url TEXT NOT NULL,
    					short_url TEXT NOT NULL UNIQUE,
    					user_id TEXT NOT NULL,
                        is_deleted BOOLEAN DEFAULT False
													);`
	createOriginalURLIndexQuery = `CREATE UNIQUE INDEX IF NOT EXISTS original_url_idx ON urls (original_url);`
	setNewValueInDB             = `INSERT INTO urls (original_url, short_url, user_id)
						VALUES ($1, $2, $3)
						ON CONFLICT (short_url)
						DO UPDATE
						SET short_url = EXCLUDED.short_url;`
	setDeleteBatch   = `UPDATE urls SET is_deleted=true WHERE short_url=$1 and user_id=$2;`
	getBatchByUserID = `SELECT original_url, short_url from urls WHERE user_id = $1`
	getOriginalURL   = `SELECT original_url, is_deleted FROM urls WHERE short_url = $1 LIMIT 1;`
	getShortURL      = `SELECT short_url FROM urls WHERE original_url = $1 LIMIT 1;`
	getCountOfURLs   = `select count(*) from urls;`
	getCountOfUsers  = `select count(DISTINCT user_id) from urls`
)

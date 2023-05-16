package migration

func init() {
	sourceDriver.append(
		20230417172902,
		`
			CREATE TABLE IF NOT EXISTS users (
				id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				email varchar(255) NOT NULL,
				password varchar(255) NOT NULL,
				created_at timestamp NULL,
				updated_at timestamp NULL,
				CONSTRAINT users_pk PRIMARY KEY (id),
				CONSTRAINT users_email_uk UNIQUE (email)
			);
		`,
		`
			DROP TABLE IF EXISTS users;
		`,
	)
}

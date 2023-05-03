package migration

func init() {
	sourceDriver.append(
		20230503102317,
		`
			CREATE TABLE IF NOT EXISTS categories (
				id char(36) NOT NULL,
				user_id char(36) NULL,
				name varchar(255) NOT NULL,
				is_global boolean NOT NULL,
				is_expense boolean NOT NULL,
				created_at timestamp NULL,
				updated_at timestamp NULL,
				CONSTRAINT categories_pk PRIMARY KEY (id),
				CONSTRAINT categories_users_fk FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`,
		`
			DROP TABLE IF EXISTS categories;
		`,
	)
}

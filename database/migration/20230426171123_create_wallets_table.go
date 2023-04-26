package migration

func init() {
	sourceDriver.append(
		20230426171123,
		`
			CREATE TABLE IF NOT EXISTS wallets (
				id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				total_amount decimal(16,2) NOT NULL,
				logo_type varchar(50) NULL,
				created_at timestamp NULL,
				updated_at timestamp NULL,
				CONSTRAINT wallets_pk PRIMARY KEY (id),
				CONSTRAINT wallets_users_fk FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`,
		`
			DROP TABLE IF EXISTS wallets;
		`,
	)
}

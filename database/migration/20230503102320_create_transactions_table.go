package migration

func init() {
	sourceDriver.append(
		20230503102320,
		`
			CREATE TABLE IF NOT EXISTS transactions (
				id char(36) NOT NULL,
				category_id char(36) NOT NULL,
				wallet_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				amount decimal(16,2) NOT NULL,
				date date NOT NULL,
				created_at timestamp NULL,
				updated_at timestamp NULL,
				CONSTRAINT transactions_pk PRIMARY KEY (id),
				CONSTRAINT transaction_categories_fk FOREIGN KEY (category_id) REFERENCES categories(id),
				CONSTRAINT transaction_wallets_fk FOREIGN KEY (wallet_id) REFERENCES wallets(id),
				CONSTRAINT transaction_users_fk FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`,
		`
			DROP TABLE IF EXISTS transactions;
		`,
	)
}

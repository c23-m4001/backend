package migration

func init() {
	sourceDriver.append(
		20230423142227,
		`
			CREATE TABLE IF NOT EXISTS user_access_tokens (
				id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				revoked boolean NOT NULL,
				expired_at timestamp NOT NULL,
				ip_address varchar(50) NULL,
				longitude double precision NULL,
				latitude double precision NULL,
				location_name varchar(255) NULL,
				created_at timestamp NULL,
				updated_at timestamp NULL,
				CONSTRAINT user_access_tokens_pk PRIMARY KEY (id),
				CONSTRAINT user_access_tokens_users_fk FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`,
		`
			DROP TABLE IF EXISTS user_access_tokens;
		`,
	)
}

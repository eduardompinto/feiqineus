package internal

type VerifiedMessageDB struct {
	db Database
}

// NewVerifiedMessageDB return a new VerifiedMessageDB
func NewVerifiedMessageDB(db Database) VerifiedMessageDB {
	return VerifiedMessageDB{db: db}
}

func (r VerifiedMessageDB) Get(message string) *VerifiedMessage {
	r.db.Initialize()
	q := `
		SELECT 
			id, first_appear, text, link, is_fake, explanation, checked, text_normalized
		FROM 
			verified_messages 
		WHERE 
			SIMILARITY(text_normalized, $1) > 0.6
		ORDER BY SIMILARITY(text_normalized, $1) DESC
		LIMIT 1;
	`
	row := r.db.Pool.QueryRow(q, message)
	if row == nil {
		return nil
	}
	var msg VerifiedMessage
	err := row.Scan(
		&msg.ID,
		&msg.FirstAppear,
		&msg.Text,
		&msg.Link,
		&msg.IsFake,
		&msg.Explanation,
		&msg.Checked,
		&msg.TextNormalized,
	)
	if err != nil {
		return nil
	}
	return &msg
}

func (r VerifiedMessageDB) Save(vm *VerifiedMessage) error {
	r.db.Initialize()
	q := `
		INSERT INTO 
			verified_messages (id, first_appear, text, link, is_fake, explanation, checked, text_normalized) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Pool.Query(
		q,
		vm.ID,
		vm.FirstAppear,
		vm.Text,
		vm.Link,
		vm.IsFake,
		vm.Explanation,
		vm.Checked,
		vm.TextNormalized,
	)
	return err
}

package versions

func init() {

}

type V1 struct{}

func (v V1) Name() string {
	return "v1"
}

func (v V1) Migration() string {
	return `
		CREATE DATABASE example (
			id int UNSIGNED NOT NULL AUTO_INCREMENT,
			message varchar(255) NOT NULL,
			last_modified timestamp,

			PRIMARY KEY (id)
		)
	`
}

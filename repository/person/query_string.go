package personrepo

const (

	getPersonWithPersonID = `
	SELECT 
		PersonID,
		FirstName,
		LastName,
		Address,
		City
	FROM
	  Persons 
	WHERE 
	  PersonID = ?;
	`
)
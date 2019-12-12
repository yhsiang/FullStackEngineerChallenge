package models

import (
	"errors"
	"time"

	"github.com/yhsiang/review360/database"
)

type Employee struct {
	ID        int64     `json:"id" uri:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (e Employee) Type() ResponseType {
	return EmployeeType
}

func (e Employee) FindAll(db database.DB) (employees Employees, err error) {
	rows, err := db.Query("SELECT id, name FROM employees")
	if err != nil {
		return employees, err
	}
	for rows.Next() {
		var employee Employee
		err = rows.Scan(&employee.ID, &employee.Name)
		if err != nil {
			return employees, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e Employee) Find(db database.DB) (Employee, error) {
	err := db.QueryRow("SELECT name FROM employees WHERE id = ?", e.ID).Scan(&e.Name)
	if err != nil {
		return e, err
	}
	return e, nil
}

func (em Employee) Save(db database.DB) (Employee, error) {
	if em.ID > 0 {
		em.UpdatedAt = time.Now()
		sqlStatement := `UPDATE employees SET name = ?, updated_at = ? WHERE id = ?`
		_, err := db.Exec(sqlStatement, em.Name, em.UpdatedAt, em.ID)
		if err != nil {
			return em, err
		}
	} else {
		em.UpdatedAt = time.Now()
		em.CreatedAt = time.Now()
		sqlStatement := `INSERT INTO employees (name, updated_at, created_at) VALUES (?)`
		result, err := db.Exec(sqlStatement, em.Name, em.UpdatedAt, em.CreatedAt)
		if err != nil {
			return em, err
		}
		em.ID, err = result.LastInsertId()
		if err != nil {
			return em, err
		}
	}

	return em, nil
}

func (em Employee) Remove(db database.DB) error {

	if em.ID == 0 {
		return errors.New("employee id should not be zero")
	}

	sqlStatement := `DELETE FROM employees WHERE id = ?`
	_, err := db.Exec(sqlStatement, em.ID)
	if err != nil {
		return err
	}

	return nil
}

type Employees []Employee

func (ems Employees) Type() ResponseType {
	return EmployeesType
}

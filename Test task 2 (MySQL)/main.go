package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dept_manager struct {
	first_name string
	last_name  string
	salary     int
	title      string
}

type employee struct {
	department string
	title      string
	first_name string
	last_name  string
	hire_date  string
	experience int
}

type department struct {
	department       string
	cur_emp_count    int
	cur_salary_count int
}

func main() {

	firstQuery()
	//secondQuery()
	//thirdQuery()

}

func firstQuery() {
	db, err := sql.Open("mysql", "root:@/employees")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT first_name,last_name,salary, title FROM employees e, dept_manager d,salaries s, titles t WHERE e.emp_no = d.emp_no AND e.emp_no = s.emp_no AND e.emp_no = t.emp_no AND CURDATE() BETWEEN d.from_date AND d.to_date AND CURDATE() BETWEEN s.from_date AND s.to_date AND CURDATE() BETWEEN t.from_date AND t.to_date")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	dept_managers := []dept_manager{}

	for rows.Next() {
		d := dept_manager{}
		err := rows.Scan(&d.first_name, &d.last_name, &d.salary, &d.title)

		if err != nil {
			fmt.Println(err)
			continue
		}
		dept_managers = append(dept_managers, d)
	}
	for _, d := range dept_managers {
		fmt.Println(d.first_name, d.last_name, d.salary, d.title)
	}
}

func secondQuery() {
	db, err := sql.Open("mysql", "root:@/employees")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT first_name, last_name, title, dept_name,TIMESTAMPDIFF(YEAR,e.hire_date,CURDATE()) AS experience FROM employees e, dept_emp de, departments d, titles t WHERE e.emp_no = de.emp_no AND de.dept_no = d.dept_no AND e.emp_no = t.emp_no AND CURDATE() BETWEEN de.from_date AND de.to_date AND CURDATE() BETWEEN t.from_date AND t.to_date")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	employees := []employee{}

	for rows.Next() {
		d := employee{}
		err := rows.Scan(&d.first_name, &d.last_name, &d.title, &d.department, &d.experience)

		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, d)
	}
	for _, d := range employees {
		fmt.Println(d.first_name, d.last_name, d.hire_date, d.title, d.department, d.experience)
	}
}

func thirdQuery() {

	db, err := sql.Open("mysql", "root:@/employees")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT dept_name, count(de.emp_no), sum(s.salary) FROM departments d, dept_emp de, salaries s WHERE d.dept_no = de.dept_no AND de.emp_no = s.emp_no AND curdate() BETWEEN s.from_date AND s.to_date AND curdate() BETWEEN de.from_date AND de.to_date GROUP BY dept_name")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	departments := []department{}

	for rows.Next() {
		d := department{}
		err := rows.Scan(&d.department, &d.cur_emp_count, &d.cur_salary_count)

		if err != nil {
			fmt.Println(err)
			continue
		}
		departments = append(departments, d)
	}
	for _, d := range departments {
		fmt.Println(d.department, d.cur_emp_count, d.cur_salary_count)
	}
}

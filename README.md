# Test tasks for Back End Internship at Readdle 2020!

>## Test task 1 (HTTP, APIs, time)
>
>Use 3rd-party JSON API: https://date.nager.at/PublicHoliday/Country/UA
>Write a console application that prints if it’s a holiday today (and the name of it). If today isn’t a holiday, the application should print the next closest holiday. 
>
>Additionally, if the holiday is adjacent to a weekend (so that amount of non-working days is extended), the application should print this information. I.e. the next holiday is May 1, Friday, and it’s adjacent to Saturday (May 2) and Sunday (May 3), so the application should print something like: “The next holiday is International Workers' Day, May 1, and the weekend will last 3 days: May 1 - May 3”.

### My approach
1. check if today is a holiday
2. determine how many days a holiday lasts
3. determine the start and end date of the holiday


>## Test task 2 (MySQL)
>Download and install the Employee sample database (https://dev.mysql.com/doc/employee/en/employees-installation.html).
>
>Structure: https://dev.mysql.com/doc/employee/en/sakila-st ructure.html.

### Queries
1. Find all current managers of each department and display his/her title, first name, last name, current salary.
```php
SELECT first_name,last_name,salary, title
FROM employees e, dept_manager d,salaries s, titles t
WHERE e.emp_no = d.emp_no
AND e.emp_no = s.emp_no
AND e.emp_no = t.emp_no
AND CURDATE() BETWEEN d.from_date AND d.to_date
AND CURDATE() BETWEEN s.from_date AND s.to_date
AND CURDATE() BETWEEN t.from_date AND t.to_date
```
2. Find all employees (department, title, first name, last name, hire date, how many years they have been working) to congratulate them on their hire anniversary this month.
```php
SELECT first_name, last_name, title, dept_name,TIMESTAMPDIFF(YEAR,e.hire_date,CURDATE()) AS experience
FROM employees e, dept_emp de, departments d, titles t
WHERE e.emp_no = de.emp_no
AND de.dept_no = d.dept_no
AND e.emp_no = t.emp_no
AND CURDATE() BETWEEN de.from_date AND de.to_date
AND CURDATE() BETWEEN t.from_date AND t.to_date
```
3. Find all departments, their current employee count, their current sum salary.
```php
SELECT dept_name, count(de.emp_no), sum(s.salary)
FROM departments d, dept_emp de, salaries s
WHERE d.dept_no = de.dept_no
AND de.emp_no = s.emp_no
AND curdate() between s.from_date AND s.to_date
AND curdate() between de.from_date AND de.to_date
GROUP BY dept_name
```

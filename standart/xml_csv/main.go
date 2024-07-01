package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// начало решения
type Organization struct {
	XMLName    xml.Name     `xml:"organization"`
	Department []Department `xml:"department"`
}

type Department struct {
	Employees struct {
		Employee []Employee `xml:"employee"`
	} `xml:"employees"`
	Code string `xml:"code"`
}

type Employee struct {
	Id     int    `xml:"id,attr"`
	Name   string `xml:"name"`
	City   string `xml:"city"`
	Salary int    `xml:"salary"`
}

func ReadData(inXML io.Reader) ([]byte, error) {
	res := make([]byte, 0)
	buf := make([]byte, 5)
	for {
		_, err := inXML.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// fmt.Printf("read %d bytes: %q\n", n, buf[:n])
		res = append(res, buf...)
	}

	return res, nil
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	var organization Organization
	data, err := ReadData(inXML)
	if err != nil {
		return err
	}
	if err = xml.Unmarshal(data, &organization); err != nil {
		return err
	}
	w := csv.NewWriter(outCSV)
	if err := w.Write([]string{"id", "name", "city", "department", "salary"}); err != nil {
		return err
	}
	for _, dep := range organization.Department {
		department := dep.Code
		for _, emp := range dep.Employees.Employee {
			if err = w.Write([]string{
				fmt.Sprintf("%d", emp.Id),
				emp.Name,
				emp.City,
				department,
				fmt.Sprintf("%d", emp.Salary),
			}); err != nil {
				return err
			}
			w.Flush()
			if err := w.Error(); err != nil {
				return err
			}
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

// конец решения

// решение преподавателя
// // Organization описывает организацию
// type Organization struct {
// 	Departments []Department `xml:"department"`
// }

// // Department описывает департамент организации
// type Department struct {
// 	Code      string     `xml:"code"`
// 	Employees []Employee `xml:"employees>employee"`
// }

// // Employee описывает сотрудника департамента
// type Employee struct {
// 	Id     int     `xml:"id,attr"`
// 	Name   string  `xml:"name"`
// 	City   string  `xml:"city"`
// 	Salary float64 `xml:"salary"`
// }

// // decodeOrganization декодирует организацию из XML-документа
// func decodeOrganization(in io.Reader) (Organization, error) {
// 	var org Organization
// 	decoder := xml.NewDecoder(in)
// 	err := decoder.Decode(&org)
// 	return org, err
// }

// // employeeWriter записывает сотрудников в CSV
// type employeeWriter struct {
// 	w   *csv.Writer
// 	err error
// }

// // writeHeader записывает заголовок
// func (ew *employeeWriter) writeHeader() {
// 	if ew.err != nil {
// 		return
// 	}
// 	header := []string{"id", "name", "city", "department", "salary"}
// 	ew.err = ew.w.Write(header)
// }

// // writeEmployee записывает сотрудника в строку
// func (ew *employeeWriter) writeEmployee(depCode string, emp Employee) {
// 	if ew.err != nil {
// 		return
// 	}
// 	fields := []string{
// 		strconv.Itoa(emp.Id),
// 		emp.Name,
// 		emp.City,
// 		depCode,
// 		strconv.FormatFloat(emp.Salary, 'f', -1, 64),
// 	}
// 	ew.err = ew.w.Write(fields)
// }

// // flush финализирует данные
// func (ew *employeeWriter) flush() error {
// 	ew.w.Flush()
// 	if ew.err == nil {
// 		ew.err = ew.w.Error()
// 	}
// 	return ew.err
// }

// // newEmployeeWriter создает нового писателя сотрудников в CSV
// func newEmployeeWriter(w io.Writer) *employeeWriter {
// 	return &employeeWriter{w: csv.NewWriter(w)}
// }

// // ConvertEmployees преобразует XML-документ с информацией об организации
// // в плоский CSV-документ с информацией о сотрудниках
// func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
// 	org, err := decodeOrganization(inXML)
// 	if err != nil {
// 		return fmt.Errorf("failed to parse xml: %w", err)
// 	}

// 	w := newEmployeeWriter(outCSV)
// 	w.writeHeader()

// 	for _, dep := range org.Departments {
// 		for _, emp := range dep.Employees {
// 			w.writeEmployee(dep.Code, emp)
// 		}
// 	}

// 	if err := w.flush(); err != nil {
// 		return fmt.Errorf("failed writing csv: %w", err)
// 	}

// 	return nil
// }



func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary></salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city></city>
                <salary></salary>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	ConvertEmployees(out, in)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}

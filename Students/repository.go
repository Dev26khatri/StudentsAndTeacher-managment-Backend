package students

//Repository knows the SQL. It means repo doing all executions of your querys like
//CREATE , UPDATE , READ(retrive)-Showing data by ID and List , DELETE
//All Querys of SQL which is done by the repository
import (
	dto "GOGIN/DTO"
	"database/sql"

	"gorm.io/gorm"
)

type Repository interface {
	Create(student Student) (*Student, error)
	GetAll() ([]Student, error)
	GetById(id int) (Student, error)
	Update(student dto.UpdateStudentRequest, id int) error
	Delete(id int) error
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{
		db: db,
	}
}

func (r *gormRepository) Create(student Student) (*Student, error) {
	// query := `INSERT INTO students(name , email , age) VALUES($1,$2,$3);`
	err := r.db.Create(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, err
}
func (r *gormRepository) GetAll() ([]Student, error) {
	// query := `SELECT *FROM students`

	// rows, err := r.db.Query(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var students []models.Student
	// for rows.Next() {
	// 	var student models.Student
	// 	err := rows.Scan(
	// 		&student.ID,
	// 		&student.Name,
	// 		&student.Email,
	// 		&student.Age,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// }
	// rows, err := r.db.Query(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var students []models.Student

	// for rows.Next() {

	// 	var student models.Student

	// 	err := rows.Scan(
	// 		&student.ID,
	// 		&student.Name,
	// 		&student.Email,
	// 		&student.Age,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	students = append(students, student)
	// }

	// return students, nil
	var student []Student
	err := r.db.Find(&student).Error
	return student, err
}
func (r *gormRepository) GetById(id int) (Student, error) {
	// query := `SELECT id , name , email ,age FROM students WHERE id = $1`

	var student Student
	// err := r.db.QueryRow(query, id).Scan(
	// 	&student.ID,
	// 	&student.Name,
	// 	&student.Email,
	// 	&student.Age,
	// )
	err := r.db.First(&student, id).Error
	if err != nil {
		return Student{}, err
	}
	return student, err
}

func (r *gormRepository) Update(student dto.UpdateStudentRequest, id int) error {
	// query := `UPDATE students SET name = $1 , email = $2 , age = $3 WHERE id = $4`

	// result, err := r.db.Exec(query, student.Name, student.Email, student.Age, id)

	result := r.db.Model(&student).Where("id = ?", id).Updates(&student)
	if result.Error != nil {
		return result.Error
	}
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func (r *gormRepository) Delete(id int) error {
	// query := `DELETE FROM students WHERE id = $1`

	// result, err := r.db.Exec(query, id)
	result := r.db.Delete(&Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return result.Error
}

func (r *gormRepository) CreateWithTx(tx *gorm.DB, student *Student) error {

	return tx.Create(student).Error
}


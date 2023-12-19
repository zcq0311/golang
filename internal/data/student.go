package data

import (
	"context"
	"gorm.io/gorm"
	"student/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
	db   *gorm.DB
}

// 初始化 studentRepo
func NewStudentRepo(db *gorm.DB, data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		db:   db,
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
	var stu biz.Student
	r.data.gormDB.Where("id = ?", id).First(&stu) // 这里使用了 gorm
	r.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
	return &biz.Student{
		ID:        stu.ID,
		Name:      stu.Name,
		Status:    stu.Status,
		Info:      stu.Info,
		UpdatedAt: stu.UpdatedAt,
		CreatedAt: stu.CreatedAt,
	}, nil
}

// CreateStudent 在数据库中插入新的学生记录
func (r *studentRepo) CreateStudent(ctx context.Context, s *biz.Student) (*biz.Student, error) {
	// 这里是 GORM 创建记录的示例
	if result := r.db.WithContext(ctx).Create(s); result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

// UpdateStudent修改学生信息
func (r *studentRepo) UpdateStudent(ctx context.Context, s *biz.Student) (*biz.Student, error) {
	if err := r.db.WithContext(ctx).Model(&biz.Student{}).Where("id = ?", s.ID).Updates(s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteStudent删除学生信息
func (r *studentRepo) DeleteStudent(ctx context.Context, id int32) error {
	result := r.db.WithContext(ctx).Delete(&biz.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

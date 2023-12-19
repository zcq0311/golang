package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Student is a Student model.
type Student struct {
	ID        int32
	Name      string
	Info      string
	Status    int32
	UpdatedAt time.Time
	CreatedAt time.Time
}

// 定义 Student 的操作接口
type StudentRepo interface {
	// 根据 id 获取学生信息
	GetStudent(context.Context, int32) (*Student, error)
	//增加学生信息
	CreateStudent(ctx context.Context, s *Student) (*Student, error)
	//修改学生信息
	UpdateStudent(ctx context.Context, s *Student) (*Student, error)
	//删除学生信息
	DeleteStudent(ctx context.Context, id int32) error
}

type StudentUsecase struct {
	repo StudentRepo
	log  *log.Helper
}

// 初始化 StudentUsecase
func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 查数据：通过 id 获取 student 信息
func (uc *StudentUsecase) Get(ctx context.Context, id int32) (*Student, error) {
	uc.log.WithContext(ctx).Infof("biz.Get: %d", id)
	return uc.repo.GetStudent(ctx, id)
}

// Create 创建新的学生记录
func (uc *StudentUsecase) Create(ctx context.Context, name string, info string) (*Student, error) {
	// 创建一个新的 Student 结构体实例
	s := &Student{
		Name: name,
		Info: info,
		// 初始化其他必要的字段
	}

	// 调用 repo 的 CreateStudent 方法来保存学生记录
	createdStudent, err := uc.repo.CreateStudent(ctx, s)
	if err != nil {
		// 处理错误
		return nil, err
	}

	// 返回创建的学生记录
	return createdStudent, nil
}

// Update 更新学生信息
func (uc *StudentUsecase) Update(ctx context.Context, s *Student) (*Student, error) {
	uc.log.WithContext(ctx).Infof("biz.Update: %d", s.ID)
	return uc.repo.UpdateStudent(ctx, s)
}

// Delete 删除学生记录
func (uc *StudentUsecase) Delete(ctx context.Context, id int32) error {
	uc.log.WithContext(ctx).Infof("biz.Delete: %d", id)
	return uc.repo.DeleteStudent(ctx, id)
}

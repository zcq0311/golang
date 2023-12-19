package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "student/api/helloworld/v1"
	"student/internal/biz"
)

// 引入 biz.StudentUsecase
type StudentService struct {
	pb.UnimplementedStudentServer

	student *biz.StudentUsecase
	log     *log.Helper
}

// 初始化
func NewStudentService(stu *biz.StudentUsecase, logger log.Logger) *StudentService {
	return &StudentService{
		student: stu,
		log:     log.NewHelper(logger),
	}
}

// 创建新的学生信息
func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	// 从请求中提取Name
	name := req.Name
	info := req.Info

	// 调用业务逻辑层的 Create 方法来创建学生
	student, err := s.student.Create(ctx, name, info)
	if err != nil {
		// 如果创建过程中发生错误，返回错误
		s.log.Errorf("Failed to create student: %v", err)
		return nil, err
	}

	// 如果创建成功，构造响应消息
	return &pb.CreateStudentReply{
		Id:   student.ID,   // 返回数据库生成的学生ID
		Name: student.Name, // 返回学生的名字
		Info: student.Info, //返回学生的信息
	}, nil
}

// 修改学生信息
func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	// 构造一个 biz.Student 对象
	studentToUpdate := &biz.Student{
		ID:     req.Id,
		Name:   req.Name,
		Info:   req.Info,
		Status: req.Status,
		// 确保其他必要的字段也被正确设置
	}

	updatedStudent, err := s.student.Update(ctx, studentToUpdate)
	if err != nil {
		s.log.Errorf("Failed to update student: %v", err)
		return nil, err
	}

	return &pb.UpdateStudentReply{
		Success: true,
		Message: "Student updated successfully",
		UpdatedStudent: &pb.Stu{
			Id:     updatedStudent.ID,
			Name:   updatedStudent.Name,
			Info:   updatedStudent.Info,
			Status: updatedStudent.Status,
			// 设置其他 Stu 类型字段...
		},
	}, nil
}

func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	err := s.student.Delete(ctx, req.Id)
	if err != nil {
		s.log.Errorf("Failed to delete student: %v", err)
		return &pb.DeleteStudentReply{
			Success: false, // 显式地设置为 false 当出现错误
		}, err
	}

	return &pb.DeleteStudentReply{
		Success: true, // 显式地设置为 true 当成功删除
	}, nil
}

// 获取学生信息
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	stu, err := s.student.Get(ctx, req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.GetStudentReply{
		Id:     stu.ID,
		Name:   stu.Name,
		Status: stu.Status,
		Info:   stu.Info,
	}, nil
}
func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	return &pb.ListStudentReply{}, nil
}

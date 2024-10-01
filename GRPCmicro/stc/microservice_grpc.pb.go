// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: microservice.proto

package GRPCmicro_stc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Teacher_SayTeacher_FullMethodName = "/Teacher/SayTeacher"
)

// TeacherClient is the client API for Teacher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeacherClient interface {
	SayTeacher(ctx context.Context, in *TeacherRequest, opts ...grpc.CallOption) (*TeacherReply, error)
}

type teacherClient struct {
	cc grpc.ClientConnInterface
}

func NewTeacherClient(cc grpc.ClientConnInterface) TeacherClient {
	return &teacherClient{cc}
}

func (c *teacherClient) SayTeacher(ctx context.Context, in *TeacherRequest, opts ...grpc.CallOption) (*TeacherReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeacherReply)
	err := c.cc.Invoke(ctx, Teacher_SayTeacher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeacherServer is the server API for Teacher service.
// All implementations must embed UnimplementedTeacherServer
// for forward compatibility.
type TeacherServer interface {
	SayTeacher(context.Context, *TeacherRequest) (*TeacherReply, error)
	mustEmbedUnimplementedTeacherServer()
}

// UnimplementedTeacherServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTeacherServer struct{}

func (UnimplementedTeacherServer) SayTeacher(context.Context, *TeacherRequest) (*TeacherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayTeacher not implemented")
}
func (UnimplementedTeacherServer) mustEmbedUnimplementedTeacherServer() {}
func (UnimplementedTeacherServer) testEmbeddedByValue()                 {}

// UnsafeTeacherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeacherServer will
// result in compilation errors.
type UnsafeTeacherServer interface {
	mustEmbedUnimplementedTeacherServer()
}

func RegisterTeacherServer(s grpc.ServiceRegistrar, srv TeacherServer) {
	// If the following call pancis, it indicates UnimplementedTeacherServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Teacher_ServiceDesc, srv)
}

func _Teacher_SayTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServer).SayTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Teacher_SayTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServer).SayTeacher(ctx, req.(*TeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Teacher_ServiceDesc is the grpc.ServiceDesc for Teacher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Teacher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Teacher",
	HandlerType: (*TeacherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayTeacher",
			Handler:    _Teacher_SayTeacher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservice.proto",
}

const (
	Student_SayStudent_FullMethodName = "/Student/SayStudent"
)

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClient interface {
	SayStudent(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*StudentReply, error)
}

type studentClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClient(cc grpc.ClientConnInterface) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) SayStudent(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*StudentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StudentReply)
	err := c.cc.Invoke(ctx, Student_SayStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServer is the server API for Student service.
// All implementations must embed UnimplementedStudentServer
// for forward compatibility.
type StudentServer interface {
	SayStudent(context.Context, *StudentRequest) (*StudentReply, error)
	mustEmbedUnimplementedStudentServer()
}

// UnimplementedStudentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStudentServer struct{}

func (UnimplementedStudentServer) SayStudent(context.Context, *StudentRequest) (*StudentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayStudent not implemented")
}
func (UnimplementedStudentServer) mustEmbedUnimplementedStudentServer() {}
func (UnimplementedStudentServer) testEmbeddedByValue()                 {}

// UnsafeStudentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServer will
// result in compilation errors.
type UnsafeStudentServer interface {
	mustEmbedUnimplementedStudentServer()
}

func RegisterStudentServer(s grpc.ServiceRegistrar, srv StudentServer) {
	// If the following call pancis, it indicates UnimplementedStudentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Student_ServiceDesc, srv)
}

func _Student_SayStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).SayStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Student_SayStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).SayStudent(ctx, req.(*StudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Student_ServiceDesc is the grpc.ServiceDesc for Student service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Student_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayStudent",
			Handler:    _Student_SayStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservice.proto",
}

const (
	Course_SayCourse_FullMethodName = "/Course/SayCourse"
)

// CourseClient is the client API for Course service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseClient interface {
	SayCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseReply, error)
}

type courseClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseClient(cc grpc.ClientConnInterface) CourseClient {
	return &courseClient{cc}
}

func (c *courseClient) SayCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourseReply)
	err := c.cc.Invoke(ctx, Course_SayCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServer is the server API for Course service.
// All implementations must embed UnimplementedCourseServer
// for forward compatibility.
type CourseServer interface {
	SayCourse(context.Context, *CourseRequest) (*CourseReply, error)
	mustEmbedUnimplementedCourseServer()
}

// UnimplementedCourseServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCourseServer struct{}

func (UnimplementedCourseServer) SayCourse(context.Context, *CourseRequest) (*CourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayCourse not implemented")
}
func (UnimplementedCourseServer) mustEmbedUnimplementedCourseServer() {}
func (UnimplementedCourseServer) testEmbeddedByValue()                {}

// UnsafeCourseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServer will
// result in compilation errors.
type UnsafeCourseServer interface {
	mustEmbedUnimplementedCourseServer()
}

func RegisterCourseServer(s grpc.ServiceRegistrar, srv CourseServer) {
	// If the following call pancis, it indicates UnimplementedCourseServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Course_ServiceDesc, srv)
}

func _Course_SayCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).SayCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Course_SayCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).SayCourse(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Course_ServiceDesc is the grpc.ServiceDesc for Course service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Course_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Course",
	HandlerType: (*CourseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayCourse",
			Handler:    _Course_SayCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservice.proto",
}
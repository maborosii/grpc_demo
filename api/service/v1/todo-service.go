package v1

import (
	"context"
	"database/sql"
	"fmt"
	v1 "grpc_demo/api/proto/v1"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// 封装数据库连接配置
type ToDoServiceServer struct {
	db *sql.DB
	// 新版grpc需要继承一个向下兼容的结构体
	v1.UnimplementedToDoServiceServer
}

// simple factory
func NewToDoServiceServer(db *sql.DB) *ToDoServiceServer {
	return &ToDoServiceServer{db: db}
}

// 校验api版本
func (t *ToDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Error(codes.Unimplemented, fmt.Sprintf("unsupported API version:service implements API version '%s',but given '%s'", apiVersion, api))
		}
	}
	return nil
}

// 获取数据库连接
func (t *ToDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := t.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "连接数据库失败"+err.Error())
	}
	return c, nil
}

// 新增记录
func (t *ToDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := t.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := t.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// 校验请求时间戳
	_, err = ptypes.Timestamp(req.Todo.Reminder)
	// _, err = ts.CheckValid(req.Todo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "参数错误"+err.Error())
	}
	res, err := c.ExecContext(ctx, "insert into ToDo(`Titile`,`Description`,`Reminder`) values (?,?,?)", req.Todo.Title, req.Todo.Description, req.Todo.Reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "添加 Todo失败"+err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "获取最新 ID失败"+err.Error())
	}
	return &v1.CreateResponse{Api: apiVersion, Id: id}, nil

}

// 查看记录
func (t *ToDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := t.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := t.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// // 校验请求时间戳
	// _, err = ptypes.Timestamp(req.Todo.Reminder)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "参数错误"+err.Error())
	// }
	rows, err := c.QueryContext(ctx, "select `ID`,`Title`,`Description`,`Reminder` From ToDo where ID = ?", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "查询失败"+err.Error())
	}
	// 关闭行扫描，注意defer的执行顺序---栈
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "获取记录失败"+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ID='%d'找不到", req.Id))
	}
	var td v1.ToDo
	var reminder time.Time

	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
		return nil, status.Error(codes.Unknown, "查找数据失败"+err.Error())
	}
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("查询到多条记录ID=%d", req.Id))
	}
	return &v1.ReadResponse{Api: apiVersion, Todo: &td}, nil

}

// 删除记录
func (t *ToDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := t.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := t.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "delete From ToDo where ID = ?", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "删除记录失败"+err.Error())
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return nil, status.Error(codes.Unknown, "无法获取删除状态"+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ID=%d的记录找不到", req.Id))
	}

	return &v1.DeleteResponse{Api: apiVersion, Status: true}, nil

}

// 更新记录
func (t *ToDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := t.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := t.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	// 校验请求时间戳
	_, err = ptypes.Timestamp(req.Todo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder参数错误"+err.Error())
	}

	res, err := c.ExecContext(ctx, "update ToDo set `Title` = ?,`Reminder` = ?where `ID` = ?", req.Todo.Title, req.Todo.Reminder, req.Todo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "更新记录失败"+err.Error())
	}
	rows, err := res.RowsAffected()

	if err != nil {
		return nil, status.Error(codes.Unknown, "无法获取更新记录状态"+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ID=%d的记录找不到", req.Todo.Id))
	}
	return &v1.UpdateResponse{Api: apiVersion, Status: true}, nil
}

// 查看全部记录
func (t *ToDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := t.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := t.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "select `ID`,`Title`,`Description`,`Reminder` From ToDo")
	if err != nil {
		return nil, status.Error(codes.Unknown, "查询失败"+err.Error())
	}
	defer rows.Close()

	todoList := []*v1.ToDo{}
	var reminder time.Time

	for rows.Next() {
		td := new(v1.ToDo)
		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, "获取记录失败"+err.Error())
		}
		// time.Time类型转换为timstamppb.Timestamp
		// td.Reminder = timestamppb.New(reminder)
		td.Reminder, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "reminder 无效"+err.Error())
		}
		todoList = append(todoList, td)

	}
	if rows.Err() != nil {
		return nil, status.Error(codes.Unknown, "获取记录失败"+err.Error())

	}
	return &v1.ReadAllResponse{Api: apiVersion, Todos: todoList}, nil

}

// func (t *ToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {
// 	fmt.Println("向下兼容")
// }

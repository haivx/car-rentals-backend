## Tạo dự án và cấu trúc thư mục - package

```
.
├── README.md
├── controller   -> các hàm xử lý với APi
├── ddl          -> chứa các câu lệnh DDL để tạo bảng
├── go.mod
├── go.sum
├── main.go      -> File chạy chính
├── model        -> định nghĩa model - entity
├── repo         -> Chứa hàm thao tác với cơ sở dữ liệu
├── router       -> Định nghĩa các router
├── test         -> Viết các hàm unit test để kiểm tra repo
├── test_post.rest -> Kiểm thử API post
├── test_user.rest -> Kiểm thử API user
└── tmp
```
## Demo Login/logout && Authentication JWT
## Tạo dự án và cấu trúc thư mục - package

```
.
├── README.md
├── controller     -> các hàm xử lý với API
├── services       -> các hàm xử lý logic theo chức năng riêng
├── util           -> các hàm xử lý logic chung
├── db             -> Chứa logic thao tác với database
├── go.mod
├── go.sum
├── main.go        -> File chạy chính
├── model          -> định nghĩa model - entity
├── repo           -> Chứa hàm thao tác với cơ sở dữ liệu
├── router         -> Định nghĩa các router
├── test           -> Viết các hàm unit test để kiểm tra repo
├── config         -> Viết các hàm config
├── config         -> Viết các script khởi tạo application
└── tmp
```
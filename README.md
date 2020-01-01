## Truy cập API
http://34.80.70.61/api/demo
## API Doc
https://app.swaggerhub.com/apis-docs/khuongnguyenBlue/vine/1.0.0
## Thiết kế hệ thống
### Biểu đồ Usecase
![Usecase Diagram](https://i.ibb.co/JyxhZNB/vin-exam.png)


Các usecase màu trắng là các chức năng có khả năng thực hiện xong trong thời gian 1 tháng
### Kịch bản sử dụng
**Role Trainee**

Code | Chức năng | Mô tả | APIs
--- | --- | --- | ---
0 | Đăng nhập | | `POST /api/login`
1 | Xem danh sách môn học | Hiển thị sau khi đăng nhập | `GET /api/subjects`
2 | Xem danh sách bài thi | Hiển thị sau khi chọn môn học | `GET /api/subjects/:id/exams`
3 | Ôn tập | Hiển thị sau khi ấn `Ôn tập` tại `#2` | định nghĩa sau
4 | Thi thử | Hiển thị sau khi ấn `Thi thử` tại `#2` | định nghĩa sau
5 | Xem bài thi | Hiển thị sau khi chọn bài thi tại `#2` | `GET /api/subjects/:id/exams/:id`
6 | Thi chính thức | Hiển thị sau khi ấn `Làm bài` tại `#5` | `GET /api/exams/:id/test`
7 | Nộp bài | Bắt đầu khi submit bài thi tại `#6`| `POST /api/exams/:id`
8 | Xem kết quả thi | Redirect tại `#7` | `/api/exams/:id/result`, `/api/exams/:id/ranking`

**Role Trainer**

Code | Chức năng | Mô tả | APIs
--- | --- | --- | ---
0 | Đăng nhập | | Định nghĩa sau
1 | Xem danh sách bài thi |  | Định nghĩa sau
2 | Tạo bài thi | | Định nghĩa sau
2 | Sửa bài thi |  | Định nghĩa sau
3 | Xóa bài thi |  | Định nghĩa sau
4 | Xem danh sách câu hỏi |  | Định nghĩa sau
5 | Tạo câu hỏi | | Định nghĩa sau
6 | Import câu hỏi | | Định nghĩa sau
7 | Export câu hỏi | | Định nghĩa sau
6 | Sửa câu hỏi |  | Định nghĩa sau
7 | Xóa câu hỏi |  | Định nghĩa sau

### Thiết kế DL

![DB Diagram](https://i.ibb.co/k8TVKw7/vin-exam-db.png)

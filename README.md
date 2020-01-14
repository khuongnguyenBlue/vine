## Truy cập API
http://34.80.70.61/api/demo
## API Doc
https://app.swaggerhub.com/apis-docs/khuongnguyenBlue/vine/1.0.0
https://www.getpostman.com/collections/f82ddeffd31c322e175b
## Thiết kế hệ thống
### Biểu đồ Usecase
![Usecase Diagram](https://s2.upanh.pro/2020/01/14/vin.exam-usecase.png)


Các usecase màu trắng là các chức năng có dự kiến sẽ hoàn thành
### Kịch bản sử dụng
**Role Trainee**

Code | Chức năng | Mô tả | APIs
--- | --- | --- | ---
1 | Đăng nhập | | `POST /api/login`
2 | Đăng ký | | `POST /api/registration`
3 | Xem danh sách môn học | Hiển thị sau khi đăng nhập | `GET /api/subjects`
4 | Xem danh sách bài thi | Hiển thị sau khi chọn môn học | `GET /api/subjects/:id/exams`
5 | Ôn tập | Hiển thị sau khi ấn `Ôn tập` tại `#4` | định nghĩa sau
6 | Thi thử | Hiển thị sau khi ấn `Thi thử` tại `#4` | định nghĩa sau
7 | Xem bài thi | Hiển thị sau khi chọn bài thi tại `#4` | `GET /api/exams/:id`
8 | Thi chính thức | Hiển thị sau khi ấn `Làm bài` tại `#7` | `GET /api/exams/:id/test`
9 | Nộp bài | Bắt đầu khi submit bài thi tại `#8`| `POST /api/exams/:id`
10 | Xem kết quả thi | Redirect tại `#9 | `/api/exams/:id/review`, `/api/exams/:id/ranking`

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

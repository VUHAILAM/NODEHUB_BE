# JOB4E_BE

This is backend source from JOB4E final project.

##Structure Project
* cmd: Chứa main của các services.
* endpoints(Controllers): Xử lý về nhận request và trả respsonse
* transports(Router): Define các endpoints, ví dụ như gắn middleware hoặc define các methods
* services: implement các hàm, các chức năng của các service
* pkg: chứa các khỏi tạo config, constant, các hàm init những cái dùng chung như redis, mysql, kafka,...
* middlewares: chứa phần implement các middleware như authen, logging, tracing,...
* deployments: chứa file scripts để deploy service như docker hay kubectl,....
* migrates: chứa các file scripts chạy sql
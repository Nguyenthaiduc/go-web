**Samesite**
<p>SameSite là một thuộc tính của Set-Cookie trong HTTP response header giúp bạn ngăn cấm một site khác truy cập vào cookie hoặc cho phép trang chụ thể. SameSite nhận giá trị thuộc 1 trong 3 (None, Strict, LaX)</p>

***Đặt SameSite = Lax***
<p>Để tăng sự tiện dụng nhưng vẫn duy trì tính an toàn nhất định, hãy đặt SameSite = Lax. Với cài đặt này, trình duyệt sẽ cho phép chia sẻ cookie giữa các trang web có cùng tên miền bắt chéo với nhau bắt nguồn từ yêu cầu GET cấp cao nhất.

Như vậy 2 trang web có thể chia sẻ dữ liệu cookie cho nhau miễn nó thuộc cùng một miền chéo, ví dụ hai web có cùng domain chính với nhau.</p>

<div style = "margin 0 auto">
<img src = "https://uploads-ssl.webflow.com/5e3a6b6029e8b285ad11a875/5f13eb661237837b32dfecc1_samesite%20cookie%20with%20compatiable%20browser.png"/>
</div>
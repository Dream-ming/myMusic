// web/js/login.js
document.getElementById("loginForm").addEventListener("submit", function (e) {
    e.preventDefault(); // 阻止默认提交行为

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    // 模拟请求，可替换成真实后端接口
    fetch("/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
    })
        .then(res => res.json())
        .then(data => {
            if (data.success) {
                alert("登录成功！");
                // 跳转到首页或其他页面
                window.location.href = "index.html";
            } else {
                alert("用户名或密码错误");
            }
        })
        .catch(err => {
            console.error("登录失败：", err);
        });
});

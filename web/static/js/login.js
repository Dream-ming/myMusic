document.getElementById("loginForm").addEventListener("submit", function (e) {
    e.preventDefault(); // 阻止表单默认提交行为

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

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
                // 登录成功，跳转到 music.html 页面
                window.location.href = "/music";
            } else {
                alert("登录失败：" + data.message);
            }
        })
        .catch(err => {
            alert("请求失败：" + err);
        });
});

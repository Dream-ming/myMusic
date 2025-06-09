function register() {
    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirm-password").value;

    if (!username || !password || !confirmPassword) {
        alert("请填写所有字段");
        return;
    }

    if (password !== confirmPassword) {
        alert("两次密码不一致");
        return;
    }

    
    fetch("/api/user/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
    })
        .then(res => res.json())
        .then(data => {
            if (data.success) {
                alert("注册成功！欢迎你，" + username);
                window.location.href = "/login";
            } else {
                alert("注册失败：" + (data.message || data.error || "未知错误"));
            }
        })
        .catch(err => {
            alert("请求失败：" + err);
        });
}

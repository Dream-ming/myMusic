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

    // 示例：这里将来你会用 fetch 调用后端注册 API
    alert("注册成功！欢迎你，" + username);
    window.location.href = "/login";
}

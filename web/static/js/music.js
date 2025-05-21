// 侧边栏开关
document.getElementById("sidebarToggle").addEventListener("click", () => {
    const sidebar = document.getElementById("sidebar");
    sidebar.classList.toggle("hidden"); // 切换隐藏/显示
});

// 加载用户信息
fetch("/api/user/info")
    .then(res => res.json())
    .then(data => {
        document.getElementById("userId").textContent = data.id;
        document.getElementById("userAge").textContent = data.age;
    })
    .catch(err => {
        console.error("加载用户信息失败:", err);
    });

// // 退出登录
// document.getElementById("logoutBtn").addEventListener("click", () => {
//     fetch("/api/logout", { method: "POST" })
//         .then(() => window.location.href = "/login")
//         .catch(err => {
//             console.error("退出登录失败:", err);
//         });
// });

document.getElementById('logoutBtn').addEventListener('click', function () {
    window.location.href = '/login';
});

// 搜索歌曲
document.getElementById("searchBtn").addEventListener("click", () => {
    const query = document.getElementById("searchInput").value.trim();
    if (query) {
        window.location.href = `/search?query=${encodeURIComponent(query)}`;
    }
});

// 加载历史最热歌曲
fetch("/api/songs/history_top")
    .then(res => res.json())
    .then(data => {
        const container = document.getElementById("historyTopSongs");
        container.innerHTML = ""; // 清空容器
        data.forEach((song, index) => {
            const li = document.createElement("li");
            li.innerHTML = `
                <a href="/player?id=${song.id}" target="_blank">
                    <span class="song-number">${index + 1}</span>
                    <span class="song-title">${song.name}</span>
                    <span class="song-artist">${song.artist}</span>
                </a>`;
            container.appendChild(li);
        });
    })
    .catch(err => {
        console.error("加载历史最热歌曲失败:", err);
    });

// 加载今日最热歌曲
fetch("/api/songs/today_top")
    .then(res => res.json())
    .then(data => {
        const container = document.getElementById("todayTopSongs");
        container.innerHTML = ""; // 清空容器
        data.forEach((song, index) => {
            const li = document.createElement("li");
            li.innerHTML = `
                <a href="/player?id=${song.id}" target="_blank">
                    <span class="song-number">${index + 1}</span>
                    <span class="song-title">${song.name}</span>
                    <span class="song-artist">${song.artist}</span>
                </a>`;
            container.appendChild(li);
        });
    })
    .catch(err => {
        console.error("加载今日最热歌曲失败:", err);
    });
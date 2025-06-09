// 获取 URL 中的 ?id= 参数
const urlParams = new URLSearchParams(window.location.search);
const songId = urlParams.get("id");

const playPauseBtn = document.getElementById("play-pause-btn");
const currentTimeElem = document.getElementById("current-time");
const totalTimeElem = document.getElementById("total-time");
const progressBarFill = document.querySelector(".progress-bar-fill");

const token = localStorage.getItem('token');
if (!token) {
    alert("请先登录后再操作！");
    window.location.href = "/login"; // 跳转到登录页面
}

const writeCommentBtn = document.getElementById("write-comment-btn");
const commentModal = document.getElementById("comment-modal");
const commentInput = document.getElementById("comment-input");
const submitCommentBtn = document.getElementById("submit-comment-btn");
const cancelCommentBtn = document.getElementById("cancel-comment-btn");
const commentsContainer = document.getElementById("comments");

writeCommentBtn.addEventListener("click", () => {
    commentModal.classList.remove("hidden");
    commentInput.focus();
});

cancelCommentBtn.addEventListener("click", () => {
    commentModal.classList.add("hidden");
    commentInput.value = ""; // 清空输入框
});

// 提交评论
submitCommentBtn.addEventListener("click", () => {
    const commentText = commentInput.value.trim();
    if (!commentText) {
        alert("评论内容不能为空！");
        return;
    }

    if (!songId) {
        alert("未能获取歌曲 ID，请刷新页面重试！");
        return;
    }

    // 发送评论请求
    fetch("/api/comment/push_comment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
            song_id: parseInt(songId, 10),
            content: commentText
        })
    })
        .then((res) => {
            if (!res.ok) {
                return res.json().then((data) => {
                    throw new Error(data.error || "提交评论失败");
                });
            }
            return res.json();
        })
        .then(() => {
            commentModal.classList.add("hidden"); // 隐藏评论输入框
            commentInput.value = ""; // 清空输入框
            loadComments(); // 重新加载评论
        })
        .catch((err) => {
            console.error("提交评论失败：", err);
            alert(err.message || "提交评论失败，请稍后重试！");
        });
});

// 渲染评论列表
function renderComments(comments) {
    commentsContainer.innerHTML = ""; // 清空现有评论
    comments.forEach((comment) => {
        const commentElem = document.createElement("div");
        commentElem.className = "comment";
        commentElem.innerHTML = `
            <span class="username">${comment.user_name || "匿名用户"}</span>
            <span class="content">${comment.content}</span>
            <span class="time">${formatTimeForComments(comment.created_at)}</span>
        `;
        commentsContainer.appendChild(commentElem);
    });
}

// 格式化评论时间
function formatTimeForComments(timestamp) {
    const date = new Date(timestamp);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    const hours = String(date.getHours()).padStart(2, "0");
    const minutes = String(date.getMinutes()).padStart(2, "0");
    return `${year}-${month}-${day} ${hours}:${minutes}`;
}

// 加载评论
function loadComments() {
    if (!songId) {
        alert("未能加载歌曲评论，请返回重试！");
        return;
    }

    fetch("/api/comment/query_comment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
            song_id: parseInt(songId, 10)
        })
    })
        .then((res) => {
            if (!res.ok) {
                throw new Error("加载评论失败");
            }
            return res.json();
        })
        .then((comments) => {
            if (!Array.isArray(comments)) {
                console.error("评论数据格式不正确：", comments);
                alert("加载评论失败，请稍后再试！");
                return;
            }
            renderComments(comments); // 渲染评论列表
        })
        .catch((err) => {
            console.error("加载评论失败：", err);
            alert("加载评论失败，请稍后再试！");
        });
}

// 页面加载时获取评论
document.addEventListener("DOMContentLoaded", () => {
    loadComments(); // 加载评论
});

// 音频对象
const audio = new Audio();

// 播放/暂停功能
playPauseBtn.addEventListener("click", () => {
    if (audio.paused) {
        audio.play();
        playPauseBtn.textContent = "⏸️ 暂停";
    } else {
        audio.pause();
        playPauseBtn.textContent = "▶️ 播放";
    }
});

// 更新播放时间和进度条
audio.addEventListener("timeupdate", () => {
    const current = audio.currentTime;
    const duration = audio.duration || 1;
    currentTimeElem.textContent = formatTime(current);
    const percent = (current / duration) * 100;
    progressBarFill.style.width = `${percent}%`;
});

// 加载完成时设置总时长
audio.addEventListener("loadedmetadata", () => {
    totalTimeElem.textContent = formatTime(audio.duration);
});

// 点击进度条跳转播放位置
document.querySelector(".progress-bar").addEventListener("click", (e) => {
    const rect = e.currentTarget.getBoundingClientRect();
    const clickX = e.clientX - rect.left;
    const percentage = clickX / rect.width;
    audio.currentTime = percentage * audio.duration;
});

// 格式化时间为 mm:ss
function formatTime(seconds) {
    const minutes = Math.floor(seconds / 60);
    const secs = Math.floor(seconds % 60);
    return `${String(minutes).padStart(2, "0")}:${String(secs).padStart(2, "0")}`;
}

// 请求歌曲信息
if (!songId) {
    alert("未能加载歌曲信息，请返回重试！");
} else {
    fetch(`/api/player?id=${songId}`)
        .then((res) => res.json())
        .then((data) => {
            if (data.error) {
                alert(data.error);
                return;
            }

            document.querySelector(".song-title").textContent = data.name || "未知歌曲";
            document.querySelector(".song-author").textContent = `演唱：${data.artist || "未知歌手"}`;
            document.querySelector(".cover-image").src = data.cover || "/static/images/default-cover.jpg";

            // 这里修正为 data.oss_url
            if (data.oss_url) {
                audio.src = data.oss_url;
            } else {
                alert("音频链接不存在！");
            }

            // 歌词处理
            const lyricsContainer = document.getElementById("lyrics");
            lyricsContainer.innerHTML = "";
            // 修正为 data.lyric_path && data.lyric_path.Valid && data.lyric_path.String
            if (data.lyric_path && data.lyric_path.Valid && data.lyric_path.String) {
                fetch(data.lyric_path.String)
                    .then((lyricsRes) => lyricsRes.ok ? lyricsRes.text() : Promise.reject("无法加载歌词"))
                    .then((lyrics) => {
                        lyrics.split("\n").forEach((line) => {
                            const p = document.createElement("p");
                            p.textContent = line;
                            lyricsContainer.appendChild(p);
                        });
                    })
                    .catch((err) => {
                        console.error("加载歌词失败：", err);
                        lyricsContainer.innerHTML = "<p>歌词加载失败！</p>";
                    });
            } else {
                lyricsContainer.innerHTML = "<p>暂无歌词</p>";
            }

            // 歌曲时长显示（单位可能是秒，需要转 mm:ss）
            if (data.duration && data.duration.Valid) {
                totalTimeElem.textContent = formatTime(data.duration.Int64);
            } else {
                totalTimeElem.textContent = "00:00";
            }
        })
        .catch((err) => {
            console.error("加载歌曲信息失败：", err);
            alert("加载歌曲信息失败，请稍后再试！");
        });
}
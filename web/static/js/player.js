// 获取 URL 中的 ?id= 参数
const urlParams = new URLSearchParams(window.location.search);
const songId = urlParams.get("id");

// DOM 元素
const playPauseBtn = document.getElementById("play-pause-btn");
const currentTimeElem = document.getElementById("current-time");
const totalTimeElem = document.getElementById("total-time");
const progressBarFill = document.querySelector(".progress-bar-fill");

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
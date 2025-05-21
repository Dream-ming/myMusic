// 获取 URL 中的 ?id= 参数
const urlParams = new URLSearchParams(window.location.search);
const songId = urlParams.get("id");

// 如果没有歌曲 ID，提示用户并返回
if (!songId) {
    alert("未能加载歌曲信息，请返回重试！");
} else {
    // 请求歌曲数据
    fetch(`/api/player?id=${songId}`)
        .then((res) => res.json())
        .then((data) => {
            if (data.error) {
                alert(data.error); // 显示错误信息
                return;
            }

            // 更新歌曲信息
            document.querySelector(".song-title").textContent = data.name || "未知歌曲";
            document.querySelector(".song-author").textContent = `演唱：${data.artist || "未知歌手"}`;
            document.querySelector(".cover-image").src = data.cover || "/static/images/default-cover.jpg";

            // 更新歌词
            const lyricsContainer = document.getElementById("lyrics");
            lyricsContainer.innerHTML = ""; // 清空原始内容

            // 检查歌词路径是否有效
            if (data.lyric_path && data.lyric_path.Valid) {
                fetch(data.lyric_path.String)
                    .then((lyricsRes) => {
                        if (!lyricsRes.ok) {
                            throw new Error("无法加载歌词文件");
                        }
                        return lyricsRes.text(); // 确保以纯文本格式解析内容
                    })
                    .then((lyrics) => {
                        lyrics.split("\n").forEach((line) => {
                            const p = document.createElement("p");
                            p.textContent = line;
                            lyricsContainer.appendChild(p);
                        });
                    })
                    .catch((err) => {
                        console.error("加载歌词失败:", err);
                        lyricsContainer.innerHTML = "<p>歌词加载失败！</p>";
                    });
            } else {
                lyricsContainer.innerHTML = "<p>暂无歌词</p>";
            }
        })
        .catch((err) => {
            console.error("加载歌曲信息失败:", err);
            alert("加载歌曲信息失败，请稍后再试！");
        });
}
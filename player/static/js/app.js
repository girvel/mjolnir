const player = document.getElementById("player");
let playlist;

let playlistIndex = -1;
const playNextTrack = () => {
    playlistIndex++;
    const name = playlist.items[playlistIndex];
    player.src = playlist.prefix + name;
    player.autoplay = true;

    if ('mediaSession' in navigator) {
        navigator.mediaSession.metadata = new MediaMetadata({
            title: name,
            //artist: "Vibe Chemistry",
        });
    }
};

player.addEventListener("ended", playNextTrack);
if ('mediaSession' in navigator) {
    navigator.mediaSession.setActionHandler("nexttrack", () => playNextTrack());
}

axios.get("/api/playlist").then((response) => {
    playlist = response.data;
    playNextTrack();
});

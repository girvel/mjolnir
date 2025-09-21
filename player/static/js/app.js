const removeExtension = (filename) => {
  let index = filename.lastIndexOf(".");
  if (index == -1) return filename;
  return filename.substring(0, index);
}

const player = document.getElementById("player");
let playlist;

let playlistIndex = -1;
const playNextTrack = () => {
  playlistIndex++;
  const filename = playlist.items[playlistIndex];
  player.src = playlist.prefix + filename;
  player.autoplay = true;

  if ('mediaSession' in navigator) {
    const [artist, title] = removeExtension(filename).split(" - ", 2);
    navigator.mediaSession.metadata = new MediaMetadata({
      artist: artist,
      title: title,
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

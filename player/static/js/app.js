const removeExtension = (filename) => {
  let index = filename.lastIndexOf(".");
  if (index == -1) return filename;
  return filename.substring(0, index);
}

const mod = (a, b) => {
  return ((a % b) + b) % b;
}

const player = document.getElementById("player");
let playlist;

let playlistIndex = -1;
const shiftTrack = (offset) => {
  playlistIndex = mod(playlistIndex + offset, playlist.items.length);
  const filename = playlist.items[playlistIndex];

  player.src = playlist.prefix + filename;
  player.play();

  if ('mediaSession' in navigator) {
    const [artist, title] = removeExtension(filename).split(" - ", 2);
    navigator.mediaSession.metadata = new MediaMetadata({
      artist: artist,
      title: title,
    });
  }
};

player.addEventListener("ended", () => shiftTrack(1));
if ('mediaSession' in navigator) {
  navigator.mediaSession.setActionHandler("nexttrack", () => shiftTrack(1));
  navigator.mediaSession.setActionHandler("previoustrack", () => shiftTrack(-1));
}

axios.get("/api/playlist").then((response) => {
  playlist = response.data;
  shiftTrack(1);
});

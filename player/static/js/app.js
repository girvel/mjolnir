const removeExtension = (filename) => {
  let index = filename.lastIndexOf(".");
  if (index == -1) return filename;
  return filename.substring(0, index);
}

const mod = (a, b) => {
  return ((a % b) + b) % b;
}

const player = {
  dom_element: document.getElementById("player"),
  playlist: undefined,
  playlistIndex: -1,

  init: function() {
    this.dom_element.addEventListener("ended", () => this.shiftTrack(1));
    if ('mediaSession' in navigator) {
      navigator.mediaSession.setActionHandler("nexttrack", () => this.shiftTrack(1));
      navigator.mediaSession.setActionHandler("previoustrack", () => this.shiftTrack(-1));
    }

    axios.get("/api/playlist").then((response) => {
      this.playlist = response.data;
      this.shiftTrack(1);
    });
  },

  shiftTrack: function(offset) {
    this.playlistIndex = mod(this.playlistIndex + offset, this.playlist.items.length);
    const filename = this.playlist.items[this.playlistIndex];

    this.dom_element.src = this.playlist.prefix + filename;
    this.dom_element.play();

    if ('mediaSession' in navigator) {
      const [artist, title] = removeExtension(filename).split(" - ", 2);
      navigator.mediaSession.metadata = new MediaMetadata({
        artist: artist,
        title: title,
      });
    }
  },
}

player.init();

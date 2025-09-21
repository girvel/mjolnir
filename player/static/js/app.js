const removeExtension = (filename) => {
  let index = filename.lastIndexOf(".");
  if (index == -1) return filename;
  return filename.substring(0, index);
}

const mod = (a, b) => {
  return ((a % b) + b) % b;
}

const player = {
  audio_dom: document.getElementById("player"),
  tracklist_dom: document.getElementById("tracklist"),
  playlist: undefined,
  playlistIndex: -1,

  init: function() {
    this.audio_dom.addEventListener("ended", () => this.shiftTrack(1));
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

    let tracklist_render = "";
    for (const [i, filename] of this.playlist.items.entries()) {
      if (i > 0) tracklist_render += "\n";
      tracklist_render += (i == this.playlistIndex ? ">" : "-") + " " + filename;
    }
    this.tracklist_dom.textContent = tracklist_render;

    const filename = this.playlist.items[this.playlistIndex];

    this.audio_dom.src = this.playlist.prefix + filename;
    this.audio_dom.play();

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

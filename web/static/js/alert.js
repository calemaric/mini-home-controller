let numberOfPlays = 1;

if (window.audio && window.audio.isPlaying) {
        window.audio.stop();
}

let audio = new Audio('/static/sounds/alert.mp3');
audio.play();

audio.addEventListener('ended', function () {
        if (numberOfPlays >= 2) {
                numberOfPlays = 1
                return
        }

        this.currentTime = 0;
        this.play();
        numberOfPlays++
}, false);




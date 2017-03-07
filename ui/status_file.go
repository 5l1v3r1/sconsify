package ui

import (
	"github.com/fabiofalci/sconsify/sconsify"
	"text/template"
	"io/ioutil"
	"bytes"
)

func ToStatusFile(fileName string, text string) {
	toFileEvents := sconsify.InitialiseEvents()

	t := template.Must(template.New("statusTemplate").Parse(text))

	for {
		select {
		case track := <-toFileEvents.TrackPausedUpdates():
			var b bytes.Buffer
			t.Execute(&b, track)
			ioutil.WriteFile(fileName, b.Bytes(), 0644)
		case track := <-toFileEvents.TrackPlayingUpdates():
			var b bytes.Buffer
			t.Execute(&b, track)
			ioutil.WriteFile(fileName, b.Bytes(), 0644)
		case <-toFileEvents.ShutdownEngineUpdates():
			break
		case <-toFileEvents.TrackNotAvailableUpdates():
		case <-toFileEvents.PlayTokenLostUpdates():
		case <-toFileEvents.NextPlayUpdates():
		case <-toFileEvents.PlaylistsUpdates():
		case <-toFileEvents.ArtistAlbumsUpdates():
		case <-toFileEvents.NewTrackLoadedUpdate():
		case <-toFileEvents.ShutdownSpotifyUpdates():
		case <-toFileEvents.SearchUpdates():
		case <-toFileEvents.PlayUpdates():
		case <-toFileEvents.ReplayUpdates():
		case <-toFileEvents.PauseUpdates():
		case <-toFileEvents.PlayPauseToggleUpdates():
		case <-toFileEvents.GetArtistAlbumsUpdates():
		}
	}
}


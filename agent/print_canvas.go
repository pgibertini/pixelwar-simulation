package agent

import (
	"fmt"
	"html/template"
	"net/http"
)

const canvasTemplate = `
<html>
  <head>
    <title>Canvas</title>
  </head>
  <body>
    <canvas id="canvas" width="{{.Width}}" height="{{.Height}}"></canvas>
    <script>
      setInterval(function() {
        // Send a request to the server to get the current state of the canvas
        fetch('/get_canvas', {
          method: 'POST',
          body: JSON.stringify({placeID: '{{.PlaceID}}'}),
          headers: {
            'Content-Type': 'application/json',
          },
        })
          .then(response => response.json())
          .then(canvas => {
            // Update the displayed image on the canvas
            const ctx = canvas.getContext('2d');
            for (let y = 0; y < canvas.height; y++) {
              for (let x = 0; x < canvas.width; x++) {
                ctx.fillStyle = canvas.grid[y][x];
                ctx.fillRect(x, y, 1, 1);
              }
            }
          });
      }, 1000); // Update the canvas every second
    </script>
  </body>
</html>
`

func (srv *Server) doCanvas(w http.ResponseWriter, r *http.Request) {
	if !srv.checkMethod("GET", w, r) {
		return
	}

	// Parse the place ID from the request
	placeID := r.URL.Query().Get("placeID")
	if placeID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing placeID parameter")
		return
	}

	// Retrieve the canvas from the server's map of places
	srv.Lock()
	place, ok := srv.places[placeID]
	srv.Unlock()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid placeID")
		return
	}

	// Pass the canvas data to the template
	tpl := template.Must(template.New("canvas").Parse(canvasTemplate))
	err := tpl.Execute(w, map[string]interface{}{
		"PlaceID": placeID,
		"Width":   place.canvas.GetWidth(),
		"Height":  place.canvas.GetHeight(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to render template: %v", err)
		return
	}
}

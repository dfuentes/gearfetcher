package main

import (
	"html/template"
	"time"
)

var homeTemplate *template.Template = template.Must(template.New("homepage").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Gear Fetcher</title>
  </head>
  <body>
    <form id="search" action="/search" method="GET" role="search">
      <center>
        <select name="server">
          <option value="Anathema">Anathema</option>
          <option value="Arcanite Reaper">Arcanite Reaper</option>
          <option value="Arugal">Arugal</option>
          <option value="Ashkandi">Ashkandi</option>
          <option value="Atiesh">Atiesh</option>
          <option value="Azuresong">Azuresong</option>
          <option value="Benediction">Benediction</option>
          <option value="Bigglesworth">Bigglesworth</option>
          <option value="Blaumeux">Blaumeux</option>
          <option value="Bloodsail Buccaneers">Bloodsail Buccaneers</option>
          <option value="Deviate Delight">Deviate Delight</option>
          <option value="Earthfury">Earthfury</option>
          <option value="Faerlina">Faerlina</option>
          <option value="Fairbanks">Fairbanks</option>
          <option value="Felstriker">Felstriker</option>
          <option value="Grobbulus">Grobbulus</option>
          <option value="Heartseeker">Heartseeker</option>
          <option value="Herod">Herod</option>
          <option value="Incendius">Incendius</option>
          <option value="Kirtonos">Kirtonos</option>
          <option value="Kromcrush">Kromcrush</option>
          <option value="Kurinnaxx">Kurinnaxx</option>
          <option value="Mankrik">Mankrik</option>
          <option value="Myzrael">Myzrael</option>
          <option value="Netherwind">Netherwind</option>
          <option value="Old Blanchy">Old Blanchy</option>
          <option value="Pagle">Pagle</option>
          <option value="Rattlegore">Rattlegore</option>
          <option value="Remulos">Remulos</option>
          <option value="Skeram">Skeram</option>
          <option value="Smolderweb">Smolderweb</option>
          <option value="Stalagg">Stalagg</option>
          <option value="Sulfuras">Sulfuras</option>
          <option value="Thalnos">Thalnos</option>
          <option value="Thunderfury">Thunderfury</option>
          <option value="Westfall">Westfall</option>
          <option value="Whitemane">Whitemane</option>
          <option value="Windseeker">Windseeker</option>
          <option value="Yojamba">Yojamba</option>
        </select>
        <input type="text" name="name" placeholder="Character name..." autofocus="on">
        <br><br>
        <input type="submit" value="Search">
      </center>
    </form>
  </body>
</html>
`))

var resultsTemplate *template.Template = template.Must(template.New("results").Funcs(template.FuncMap{"formatDate": FormatDate}).Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Gear Fetcher</title>
  </head>
  <body>
  {{range .}}
  <p> {{formatDate .StartTime}} - {{.EncounterName}} - {{.Percentile}} - <a href={{.GetGearLink}}>Gear</a></p>
  {{else}}
  No parses found
  {{end}}
  </body>
</html>
`))

func FormatDate(ts int) string {
	t := time.Unix(0, int64(ts)*1e6)
	return t.Format("01/02/06")
}

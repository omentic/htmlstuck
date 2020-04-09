<img align='middle' src='https://j-james.me/htmlstuck/assets/htmlstuck.png'>

## working through issues - roadmap
- [ ] fix polyfill
	- [x] css either loads on j-james.me/htmlstuck or j-james.me/htmlstuck/story, never both
		- `../` accesses different locations from `htmlstuck` and `htmlstuck/story`.
		- fix: refer to assets and css with absolute links (`j-james.me/htmlstuck/css/main.css`)
	- [x] [cross-origin requests are blocked](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS/Errors/CORSMissingAllowOrigin)
		-	kludge: host swfs locally (this is not scalable)
	- [ ] response has unsupported mime type
		- [known bug](https://github.com/ruffle-rs/ruffle/issues/400)
		- todo: send `application/wasm` as a mime type
		- is that possible on github pages?
	- [ ] failed to compile: http code is not ok
		- chromium produces this
		- this is either a fancy mime error or its own thing
		- nothing's turned up about this, i might open an issue
- [ ] `404.html` should load a random image from https://mspaintadventures.fandom.com/wiki/Scribble_Mode
	- should be easily accomplishable in javascript
- [ ] deal with ruffle better
	- currently `ruffle_web_latest.zip` is downloaded by me and thrown into `ruffle`
	- fetching the latest version on push should be possible with github actions
	- although, if it doesn't update automatically, maybe that's not super useful
- [x] make a nice logo

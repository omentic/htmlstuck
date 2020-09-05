<img align='middle' src='https://j-james.me/htmlstuck/assets/htmlstuck.png'>

<!-- (psst - htmlstuck.com is available) (for pretty cheap too) -->

<!-- ## What is this?

With Adobe Flash's end-of-life set for [December 31, 2020](https://www.adobe.com/products/flashplayer/end-of-life.html), major web browsers have begun to drop support entirely.

## How can I use this on my website? -->

## Roadmap
- [x] Fix polyfill
	- [x] CSS either loads on <https://j-james.me/htmlstuck> or <https://j-james.me/htmlstuck/story>, never both
		- `../` accesses different locations from `htmlstuck` and `htmlstuck/story`.
		- Fix: Refer to assets and css with absolute links (`j-james.me/htmlstuck/css/main.css`)
	- [x] [cross-origin requests are blocked](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS/Errors/CORSMissingAllowOrigin)
		- Kludge: Host SWFs locally
	- [x] Response has unsupported MIME type
		- [known bug](https://github.com/ruffle-rs/ruffle/issues/400)
		- Fix: Send `application/wasm` as a mime type
		- GitHub Pages now does this by default
	- [x] Failed to compile: HTTP code is not ok
    	- Fixed by latest Ruffle version
- [x] Make a nice logo
- [x] 404.html should load a random image <!-- from https://mspaintadventures.fandom.com/wiki/Scribble_Mode -->
- [x] Don't link directly to HTML pages
- [x] Accommodate special Flash pages - [S] Cascade, [S] GAME OVER, and [S] Terezi: Remem8er come to mind
- [x] Deal with Ruffle updates better
	- Currently, `ruffle_web_latest.zip` is downloaded by me and thrown into `ruffle`
	- Is the latest version hosted somewhere?
    - Honestly, keeping this as is is probably best. Homestuck flashes use a small subset of Flash functionality, so Ruffle updates mostly affect them.
- [x] Track down missing early Flashes - mspa.com?
	- [cdn.mspaintadventures.com](https://cdn.mspaintadventures.com) still hosts old Flashes
- [ ] Update with Flash non-[S] pages
- [ ] Nicer landing page - something like this README, with a big link to the lists
- [ ] Test browser / OS compatibility
  - [ ] Makin: bad compression / crushed pixels (Firefox)
  - [ ] Bambosh: wack diagonal seam (Firefox)
- [ ] **Autoplay Flashes upon load**

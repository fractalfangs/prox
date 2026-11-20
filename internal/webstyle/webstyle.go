// Package webstyle holds the shared stylesheet for the small HTML pages served
// to browsers (OAuth callbacks, safe-mode notices). The look is inspired by
// lio.cat: monospace, lowercase, amber and cobalt, no cards or shadows.
package webstyle

// CSS is embedded inside a <style> element by every page.
//
// It must not contain a percent sign, because some pages are rendered through
// fmt.Sprintf.
const CSS = `
:root{--bg:#fbbf24;--fg:#121212;--accent:#1447e6;--muted:rgba(18,18,18,.6);--rule:rgba(18,18,18,.25);color-scheme:light}
@media (prefers-color-scheme:dark){:root{--bg:#121212;--fg:#ffffff;--accent:#ffb900;--muted:rgba(255,255,255,.5);--rule:rgba(255,255,255,.15);color-scheme:dark}}
*{box-sizing:border-box}
html{background:var(--bg)}
body{margin:0;min-height:100vh;display:flex;justify-content:center;padding:0 24px;background:var(--bg);color:var(--fg);font-family:"JetBrains Mono","JetBrains Mono Variable",ui-monospace,SFMono-Regular,Menlo,Consolas,monospace;font-size:16px;line-height:1.6;text-transform:lowercase;-webkit-font-smoothing:antialiased}
main{display:flex;flex-direction:column;flex:0 1 48rem;min-width:0;padding:20px 8px 0}
.brand{display:block;font-size:1.5rem;font-weight:700;letter-spacing:-.02em;padding:24px 0 16px;color:var(--fg)}
.tags{color:var(--muted)}
.tags b{color:var(--accent);font-weight:500}
h1{margin:1.25rem 0 1rem;font-size:1.25rem;font-weight:500}
p{margin:0 0 1rem}
ul{margin:0 0 1rem;padding-left:1.25rem}
code{font-family:inherit;text-transform:none;color:var(--accent);word-break:break-all}
a{position:relative;padding-right:1em;color:var(--accent);text-decoration:none}
a::after{content:"\2197";position:absolute;top:0;right:.1em;line-height:1}
a:hover,a:focus-visible,.link:hover,.link:focus-visible{text-decoration:underline;text-underline-offset:4px;outline:none}
.nav{display:flex;flex-wrap:wrap;gap:1rem;margin:.5rem 0 1rem;font-size:1.125rem}
.link{padding:0;border:0;background:none;color:var(--accent);font:inherit;text-transform:inherit;cursor:pointer}
.notice{margin:1rem 0;padding-left:1rem;border-left:2px solid var(--accent)}
.notice h3{margin:0 0 .25rem;font-size:1rem;font-weight:500}
.muted{color:var(--muted)}
.error{color:#b91c1c}
@media (prefers-color-scheme:dark){.error{color:#fca5a5}}
footer{display:flex;flex-wrap:wrap;gap:.25rem;padding:40px 0;opacity:.35}
`

package handler

import (
	"html/template"
	"net/http"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>GoSleep</title>
    <link
      href="https://fonts.googleapis.com/css2?family=Geist:wght@400;500;600;700&display=swap"
      rel="stylesheet"
    />
    <style>
      * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
      }

      body {
        font-family: "Geist", -apple-system, BlinkMacSystemFont, "Segoe UI",
          "Roboto", sans-serif;
        background: linear-gradient(rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.7)),
          url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 20"><defs><pattern id="grass" x="0" y="0" width="10" height="20" patternUnits="userSpaceOnUse"><path d="M0 20 Q2 15 5 20 Q7 15 10 20" fill="none" stroke="%2334d399" stroke-width="0.5"/><path d="M1 20 Q3 12 6 20 Q8 12 11 20" fill="none" stroke="%2310b981" stroke-width="0.3"/><path d="M2 20 Q4 10 7 20 Q9 10 12 20" fill="none" stroke="%2359f287" stroke-width="0.2"/></pattern></defs><rect width="100" height="20" fill="%23065f46"/><rect width="100" height="20" fill="url(%23grass)"/></svg>')
            repeat;
        color: #fff;
        min-height: 100vh;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;
        padding: 2rem;
      }

      .main-text {
        font-size: clamp(4rem, 12vw, 8rem);
        font-weight: 700;
        line-height: 1.1;
        background: linear-gradient(135deg, #fff 0%, #888 100%);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        background-clip: text;
        margin-bottom: 2rem;
      }

      .subtitle {
        font-size: clamp(2rem, 6vw, 4rem);
        font-weight: 600;
        margin-bottom: 3rem;
      }

      .go-text {
        color: #00add8;
        text-shadow: 0 0 20px rgba(0, 173, 216, 0.3);
      }

      .sleep-text {
        color: white;
      }

      .learn-text {
        font-size: clamp(1.2rem, 3vw, 2rem);
        font-weight: 500;
        margin-bottom: 2rem;
        color: #e5e7eb;
      }

      .learn-text a {
        color: #00add8;
        text-decoration: none;
        transition: all 0.3s ease;
      }

      .learn-text a:hover {
        color: #0ea5e9;
        text-shadow: 0 0 10px rgba(0, 173, 216, 0.5);
      }

      .nextjs-link {
        color: #fff !important;
        background: linear-gradient(135deg, #000 0%, #333 100%);
        padding: 0.3rem 0.6rem;
        border-radius: 6px;
        border: 1px solid #444;
        transition: all 0.3s ease;
      }

      .nextjs-link:hover {
        background: linear-gradient(135deg, #111 0%, #444 100%);
        border-color: #666;
        transform: translateY(-1px);
      }

      .built-by {
        position: fixed;
        bottom: 2rem;
        right: 2rem;
        font-size: 0.9rem;
        color: #888;
        background: rgba(0, 0, 0, 0.8);
        padding: 0.5rem 1rem;
        border-radius: 8px;
        backdrop-filter: blur(10px);
        border: 1px solid #333;
      }

      .built-by a {
        color: #00add8;
        text-decoration: none;
        transition: color 0.3s ease;
      }

      .built-by a:hover {
        color: #0ea5e9;
      }

      @media (max-width: 768px) {
        .main-text {
          font-size: clamp(2.5rem, 10vw, 6rem);
        }

        .subtitle {
          font-size: clamp(1.5rem, 5vw, 3rem);
        }

        .learn-text {
          font-size: clamp(1rem, 2.5vw, 1.5rem);
        }

        .built-by {
          position: static;
          margin-top: 3rem;
          font-size: 0.8rem;
        }
      }

      /* Grass animation */
      @keyframes sway {
        0%,
        100% {
          transform: translateX(0);
        }
        50% {
          transform: translateX(2px);
        }
      }

      body::before {
        content: "";
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: inherit;
        animation: sway 3s ease-in-out infinite;
        z-index: -1;
      }
    </style>
  </head>
  <body>
    <div class="main-text">seriously disappointing</div>
    <div class="subtitle">
      <span class="go-text">go</span><span class="sleep-text">sleep</span>
    </div>
    <div class="learn-text">
	 U always begging for help, tf? See
      <a href="https://chatgpt.com" target="_blank">chatgpt.com</a>,
      <a href="https://google.com" target="_blank">google.com</a>,
      <a href="https://claude.ai" target="_blank">claude.ai</a>,
      <a href="https://gemini.google.com" target="_blank">gemini.google.com</a>,
      <a href="https://youtube.com" target="_blank">youtube.com</a>,
      <a href="https://stackoverflow.com" target="_blank">stackoverflow.com</a>,
       GTFO.
	   <br /><br />
      u should learn golang and
      <a
        href="https://idolodevblog.vercel.app/blogs/you-should-be-using-nextjs"
        target="_blank"
        class="nextjs-link"
        >nextjs</a
      >
      btw.  it's the best.
    </div>

    <div class="built-by">
      built by
      <a href="https://treasureuzoma.netlify.app" target="_blank"
        >treasure uzoma</a
      >
    </div>
  </body>
</html>
`

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("page").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}
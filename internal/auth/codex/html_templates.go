package codex

import "github.com/router-for-me/CLIProxyAPI/v8/internal/webstyle"

// LoginSuccessHTML is the HTML template for the page shown after a successful
// OAuth2 authentication with Codex. It informs the user that the authentication
// was successful and provides a countdown timer to automatically close the window.
const LoginSuccessHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Authentication Successful - Codex</title>
    <link rel="icon" type="image/svg+xml" href="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23ffb900'%3E%3Ccircle cx='12' cy='12' r='10'/%3E%3C/svg%3E">
    <style>` + webstyle.CSS + `</style>
</head>
<body>
    <main>
        <span class="brand">cliproxyapi</span>
        <p class="tags"><b>codex</b> ⋅ oauth ⋅ signed in</p>
        <h1>Authentication Successful!</h1>
        <p>You have successfully authenticated with Codex. You can now close this window and return to your terminal to continue.</p>

        {{SETUP_NOTICE}}

        <nav class="nav">
            <button id="close-window" class="link" onclick="window.close()">Close Window</button>
            <a href="{{PLATFORM_URL}}" target="_blank">Open Platform</a>
        </nav>

        <p class="muted">This window will close automatically in <span id="countdown">10</span> seconds</p>

        <footer>
            <span>cliproxyapi</span><span>⋅</span><span>codex</span>
        </footer>
    </main>

    <script>
        let countdown = 10;
        const countdownElement = document.getElementById('countdown');

        const timer = setInterval(() => {
            countdown--;
            countdownElement.textContent = countdown;

            if (countdown <= 0) {
                clearInterval(timer);
                window.close();
            }
        }, 1000);

        // Close window when user presses Escape
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape') {
                window.close();
            }
        });

        // Focus the close button for keyboard accessibility
        document.getElementById('close-window').focus();
    </script>
</body>
</html>`

// SetupNoticeHTML is the HTML template for the section that provides instructions
// for additional setup. This is displayed on the success page when further actions
// are required from the user.
const SetupNoticeHtml = `
        <div class="notice">
            <h3>Additional Setup Required</h3>
            <p>To complete your setup, please visit the <a href="{{PLATFORM_URL}}" target="_blank">Codex</a> to configure your account.</p>
        </div>`

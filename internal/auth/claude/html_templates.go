// Package claude provides authentication and token management functionality
// for Anthropic's Claude AI services. It handles OAuth2 token storage, serialization,
// and retrieval for maintaining authenticated sessions with the Claude API.
package claude

import "github.com/router-for-me/CLIProxyAPI/v8/internal/webstyle"

// LoginSuccessHtml is the HTML template displayed to users after successful OAuth authentication.
// This template provides a user-friendly success page with options to close the window
// or navigate to the Claude platform. It includes automatic window closing functionality
// and keyboard accessibility features.
const LoginSuccessHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Authentication Successful - Claude</title>
    <link rel="icon" type="image/svg+xml" href="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23ffb900'%3E%3Ccircle cx='12' cy='12' r='10'/%3E%3C/svg%3E">
    <style>` + webstyle.CSS + `</style>
</head>
<body>
    <main>
        <span class="brand">cliproxyapi</span>
        <p class="tags"><b>claude</b> ⋅ oauth ⋅ signed in</p>
        <h1>Authentication Successful!</h1>
        <p>You have successfully authenticated with Claude. You can now close this window and return to your terminal to continue.</p>

        {{SETUP_NOTICE}}

        <nav class="nav">
            <button id="close-window" class="link" onclick="window.close()">Close Window</button>
            <a href="{{PLATFORM_URL}}" target="_blank">Open Platform</a>
        </nav>

        <p class="muted">This window will close automatically in <span id="countdown">10</span> seconds</p>

        <footer>
            <span>cliproxyapi</span><span>⋅</span><span>claude</span>
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

// SetupNoticeHtml is the HTML template for the setup notice section.
// This template is embedded within the success page to inform users about
// additional setup steps required to complete their Claude account configuration.
const SetupNoticeHtml = `
        <div class="notice">
            <h3>Additional Setup Required</h3>
            <p>To complete your setup, please visit the <a href="{{PLATFORM_URL}}" target="_blank">Claude</a> to configure your account.</p>
        </div>`

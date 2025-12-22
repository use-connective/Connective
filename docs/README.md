****<div align="center">
<h2>Connective - Integration Infrastructure for every product</h2>
</div>

**Connective** is an open-source, developer-friendly platform that enables SaaS teams to offer plug-and-play third-party integrations directly inside their product with minimal engineering effort.

### See how it works:
Suppose you want to integrate **Slack** into your SaaS to send messages to your users.

1. Select **Slack** from the **Connective Dashboard**.  
   <br>
   ![Dashboard](./../assets/Dashboard.png)
   <br><br>
2. Add Slack OAuth **Client ID** & **Client Secret**.  
   <br>
   ![OAuth](./../assets/oAuth.png)
   <br><br>

3. Use the **Connective SDK** inside your product:
    ```typescript
    const sdk = createConnective({
        projectId: "",
        userId: "",
        baseURL: "http://localhost:8082", // Backend URL
        projectSecret: "",
    });
    
    const handleConnect = (provider: string) => {
        sdk.open({ provider });
    };
    ```

4. This opens the Connect Portal, where users can authenticate & connect Slack.
   <br><br>
   ![Dashboard](./../assets/ConnectPortal.png)
   <br> <br>

5. Once connected, you can perform actions like sending messages, viewing channels, etc.
   Connective handles OAuth, token refresh & re-authentication automatically.
   Other integrations work the same way.


### 🌐 More Details
You may visit [useconnective.tech](https://useconnective.tech) for more details.

### ❗IMPORTANT
This project is under rapid development and hence not ready for product use. You can join waitlist by visiting project website.

### 🎉 Be Part of the Community

Help shape the future of the project!
Whether you're contributing code, suggesting features, or just curious — our Discord is the place to be.

[👉 Join Discord](https://discord.gg/9uwbKse6)

### Current Implementation
**Connective** is currently in it's very early development phase. As of now it only supports few integrations. Also actions can't be performed only integration can be done. **It is not ready for production use at this time.**

### Future Planning
I'm planning to add following features in the future:
* Actions - Call 3rd party integration APIs directly from SDK.
* Workflow Builder - Automate Stuff
* App Events - Send custom events from you app and trigger actions.
* Webhook Triggering - Do stuff when something happens in 3rd party integration.


### Contribute / Feedback

I'm building this in public and would love feedback!
You can help by:

⭐ Starring the repo <br>
🧩 Opening issues with suggestions <br>
💬 Reviewing architecture/pr design <br>
🛠 Contributing connectors or improvements
